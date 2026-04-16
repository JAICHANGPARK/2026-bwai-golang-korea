package wikilookup

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

var tokenPattern = regexp.MustCompile(`[A-Za-z]{2,}|[가-힣]{2,}`)

var excludedFileNames = map[string]bool{
	"README.md": true,
	"index.md":  true,
	"log.md":    true,
}

var excludedDirectories = map[string]bool{
	"lint":         true,
	"queries":      true,
	"source-notes": true,
}

// WikiMatch is one ranked wiki page candidate for post-review enrichment.
type WikiMatch struct {
	Path    string  `json:"path"`
	Score   float64 `json:"score"`
	Excerpt string  `json:"excerpt"`
}

// WikiContext is a compact evidence bundle that a Wiki Knowledge Enricher
// agent can turn into learning_context.
type WikiContext struct {
	TargetProfile string      `json:"target_profile"`
	WikiRoot      string      `json:"wiki_root"`
	QueryKeywords []string    `json:"query_keywords"`
	IndexExcerpt  []string    `json:"index_excerpt"`
	Matches       []WikiMatch `json:"matches"`
	UsageNotes    []string    `json:"usage_notes"`
}

// ResolveWikiRoot finds the local wiki directory from a starting path.
func ResolveWikiRoot(startPath string) (string, error) {
	base := startPath
	if base == "" {
		cwd, err := os.Getwd()
		if err != nil {
			return "", err
		}
		base = cwd
	}

	absoluteBase, err := filepath.Abs(base)
	if err != nil {
		return "", err
	}

	info, err := os.Stat(absoluteBase)
	if err == nil && !info.IsDir() {
		absoluteBase = filepath.Dir(absoluteBase)
	}

	current := absoluteBase
	for {
		directWiki := filepath.Join(current, "wiki", "index.md")
		if _, err := os.Stat(directWiki); err == nil {
			return filepath.Join(current, "wiki"), nil
		}

		if filepath.Base(current) == "wiki" {
			directIndex := filepath.Join(current, "index.md")
			if _, err := os.Stat(directIndex); err == nil {
				return current, nil
			}
		}

		parent := filepath.Dir(current)
		if parent == current {
			break
		}
		current = parent
	}

	return "", fmt.Errorf("could not locate wiki/index.md from %q", startPath)
}

// LookupWikiContext reads the local wiki and returns a compact evidence bundle.
func LookupWikiContext(
	problemText string,
	approvedSolutionSummary string,
	conceptCandidates []string,
	targetProfile string,
	maxPages int,
	startPath string,
) (WikiContext, error) {
	if maxPages <= 0 {
		maxPages = 4
	}
	if targetProfile == "" {
		targetProfile = "KR-2022-middle"
	}

	wikiRoot, err := ResolveWikiRoot(startPath)
	if err != nil {
		return WikiContext{}, err
	}

	indexBytes, err := os.ReadFile(filepath.Join(wikiRoot, "index.md"))
	if err != nil {
		return WikiContext{}, err
	}

	keywords := normalizeKeywords(problemText, approvedSolutionSummary, conceptCandidates)
	matches, err := collectMatches(wikiRoot, keywords)
	if err != nil {
		return WikiContext{}, err
	}
	if len(matches) > maxPages {
		matches = matches[:maxPages]
	}

	return WikiContext{
		TargetProfile: targetProfile,
		WikiRoot:      wikiRoot,
		QueryKeywords: keywords,
		IndexExcerpt:  collectIndexExcerpt(string(indexBytes), keywords, 6),
		Matches:       matches,
		UsageNotes: []string{
			"Run this lookup only after reviewer approval.",
			"Use the result to build a compact learning_context, not to guess the math answer.",
			"If the matches are weak or sparse, mark curriculum mapping and difficulty as tentative.",
		},
	}, nil
}

func normalizeKeywords(problemText string, approvedSolutionSummary string, conceptCandidates []string) []string {
	keywords := make([]string, 0, 12)
	seen := map[string]bool{}

	addKeyword := func(value string) {
		trimmed := strings.TrimSpace(value)
		lowered := strings.ToLower(trimmed)
		if trimmed == "" || seen[lowered] {
			return
		}
		if isDigitsOnly(trimmed) {
			return
		}
		seen[lowered] = true
		keywords = append(keywords, trimmed)
	}

	for _, concept := range conceptCandidates {
		addKeyword(concept)
	}

	for _, token := range tokenPattern.FindAllString(problemText+" "+approvedSolutionSummary, -1) {
		addKeyword(token)
		if len(keywords) >= 12 {
			break
		}
	}

	return keywords
}

func collectMatches(wikiRoot string, keywords []string) ([]WikiMatch, error) {
	matches := []WikiMatch{}
	repoRoot := filepath.Dir(wikiRoot)

	err := filepath.WalkDir(wikiRoot, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			if excludedDirectories[d.Name()] {
				return filepath.SkipDir
			}
			return nil
		}
		if filepath.Ext(path) != ".md" || excludedFileNames[filepath.Base(path)] {
			return nil
		}

		rawBytes, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		body := stripFrontmatter(string(rawBytes))
		score := scorePage(path, body, keywords)
		if score <= 0 {
			return nil
		}

		relativePath, err := filepath.Rel(repoRoot, path)
		if err != nil {
			return err
		}
		matches = append(matches, WikiMatch{
			Path:    filepath.ToSlash(relativePath),
			Score:   score,
			Excerpt: extractExcerpt(body, keywords, 280),
		})
		return nil
	})
	if err != nil {
		return nil, err
	}

	sort.Slice(matches, func(i, j int) bool {
		if matches[i].Score == matches[j].Score {
			return matches[i].Path < matches[j].Path
		}
		return matches[i].Score > matches[j].Score
	})

	return matches, nil
}

func collectIndexExcerpt(indexText string, keywords []string, limit int) []string {
	excerpts := []string{}
	loweredKeywords := make([]string, 0, len(keywords))
	for _, keyword := range keywords {
		loweredKeywords = append(loweredKeywords, strings.ToLower(keyword))
	}

	for _, rawLine := range strings.Split(indexText, "\n") {
		line := strings.TrimSpace(rawLine)
		if line == "" {
			continue
		}
		loweredLine := strings.ToLower(line)
		for _, keyword := range loweredKeywords {
			if strings.Contains(loweredLine, keyword) {
				excerpts = append(excerpts, line)
				break
			}
		}
		if len(excerpts) >= limit {
			break
		}
	}

	return excerpts
}

func stripFrontmatter(text string) string {
	lines := strings.Split(text, "\n")
	if len(lines) == 0 || strings.TrimSpace(lines[0]) != "---" {
		return text
	}

	for index := 1; index < len(lines); index++ {
		if strings.TrimSpace(lines[index]) == "---" {
			return strings.TrimSpace(strings.Join(lines[index+1:], "\n"))
		}
	}

	return text
}

func scorePage(path string, body string, keywords []string) float64 {
	loweredBody := strings.ToLower(body)
	loweredName := strings.ToLower(strings.ReplaceAll(strings.TrimSuffix(filepath.Base(path), filepath.Ext(path)), "-", " "))
	pathParts := strings.Split(filepath.ToSlash(path), "/")
	pathSet := map[string]bool{}
	for _, part := range pathParts {
		pathSet[part] = true
	}

	score := 0.0
	for _, keyword := range keywords {
		loweredKeyword := strings.ToLower(keyword)
		score += float64(min(strings.Count(loweredBody, loweredKeyword), 5))
		if strings.Contains(loweredName, loweredKeyword) {
			score += 3.0
		}
	}

	if pathSet["syntheses"] {
		score += 2.0
	}
	if pathSet["components"] || pathSet["overview"] {
		score += 0.5
	}

	return score
}

func extractExcerpt(body string, keywords []string, limit int) string {
	loweredKeywords := make([]string, 0, len(keywords))
	for _, keyword := range keywords {
		loweredKeywords = append(loweredKeywords, strings.ToLower(keyword))
	}

	lines := []string{}
	for _, rawLine := range strings.Split(body, "\n") {
		line := strings.TrimSpace(rawLine)
		if line != "" {
			lines = append(lines, line)
		}
	}

	for index, line := range lines {
		loweredLine := strings.ToLower(line)
		for _, keyword := range loweredKeywords {
			if strings.Contains(loweredLine, keyword) {
				if strings.HasPrefix(line, "#") && index+1 < len(lines) {
					return trimToLimit(line+" "+lines[index+1], limit)
				}
				return trimToLimit(line, limit)
			}
		}
	}

	for _, line := range lines {
		if !strings.HasPrefix(line, "#") {
			return trimToLimit(line, limit)
		}
	}

	return ""
}

func trimToLimit(value string, limit int) string {
	if limit <= 0 {
		return value
	}
	runes := []rune(value)
	if len(runes) <= limit {
		return value
	}
	return string(runes[:limit])
}

func isDigitsOnly(value string) bool {
	for _, char := range value {
		if char < '0' || char > '9' {
			return false
		}
	}
	return value != ""
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
