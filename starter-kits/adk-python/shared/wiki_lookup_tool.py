#!/usr/bin/env python3
"""Minimal wiki lookup helper for ADK runtime context building.

This helper reads the local LLM wiki, finds a small set of relevant pages,
and returns a compact evidence bundle that a Wiki Knowledge Enricher agent can
turn into ``learning_context``.
"""

from __future__ import annotations

import argparse
import json
import re
from dataclasses import asdict, dataclass
from pathlib import Path
from typing import Iterable, Sequence

EXCLUDED_FILE_NAMES = {"README.md", "index.md", "log.md"}
EXCLUDED_DIRECTORIES = {"lint", "queries", "source-notes"}
TOKEN_PATTERN = re.compile(r"[A-Za-z]{2,}|[가-힣]{2,}")


@dataclass
class WikiMatch:
    path: str
    score: float
    excerpt: str


@dataclass
class WikiContext:
    target_profile: str
    wiki_root: str
    query_keywords: list[str]
    index_excerpt: list[str]
    matches: list[WikiMatch]
    usage_notes: list[str]


def resolve_wiki_root(start_path: str | Path | None = None) -> Path:
    """Resolve the local wiki directory from cwd or a provided hint path."""

    start = Path(start_path).resolve() if start_path else Path.cwd().resolve()
    current = start if start.is_dir() else start.parent

    while True:
        direct_wiki = current / "wiki"
        if (direct_wiki / "index.md").exists():
            return direct_wiki

        if current.name == "wiki" and (current / "index.md").exists():
            return current

        if current.parent == current:
            break
        current = current.parent

    raise FileNotFoundError("Could not locate wiki/index.md from the provided path.")


def _strip_frontmatter(text: str) -> str:
    lines = text.splitlines()
    if not lines or lines[0].strip() != "---":
        return text

    for index, line in enumerate(lines[1:], start=1):
        if line.strip() == "---":
            return "\n".join(lines[index + 1 :]).strip()

    return text


def _normalize_keywords(
    problem_text: str, approved_solution_summary: str, concept_candidates: Sequence[str]
) -> list[str]:
    keywords: list[str] = []
    seen: set[str] = set()

    def add_keyword(raw_value: str) -> None:
        value = raw_value.strip()
        key = value.casefold()
        if not value or key in seen:
            return
        if value.isdigit():
            return
        seen.add(key)
        keywords.append(value)

    for concept in concept_candidates:
        add_keyword(concept)

    for token in TOKEN_PATTERN.findall(f"{problem_text} {approved_solution_summary}"):
        add_keyword(token)
        if len(keywords) >= 12:
            break

    return keywords


def _relative_page_path(page_path: Path, wiki_root: Path) -> str:
    repo_root = wiki_root.parent
    return str(page_path.relative_to(repo_root))


def _collect_index_excerpt(index_text: str, keywords: Sequence[str], limit: int = 6) -> list[str]:
    excerpts: list[str] = []
    lowered_keywords = [keyword.casefold() for keyword in keywords]

    for raw_line in index_text.splitlines():
        line = raw_line.strip()
        if not line:
            continue
        lowered_line = line.casefold()
        if any(keyword in lowered_line for keyword in lowered_keywords):
            excerpts.append(line)
        if len(excerpts) >= limit:
            break

    return excerpts


def _extract_excerpt(body_text: str, keywords: Sequence[str], limit: int = 280) -> str:
    lowered_keywords = [keyword.casefold() for keyword in keywords]
    lines = [line.strip() for line in body_text.splitlines() if line.strip()]

    for index, line in enumerate(lines):
        lowered_line = line.casefold()
        if any(keyword in lowered_line for keyword in lowered_keywords):
            if line.startswith("#") and index + 1 < len(lines):
                candidate = f"{line} {lines[index + 1]}"
                return candidate[:limit]
            return line[:limit]

    for line in lines:
        if not line.startswith("#"):
            return line[:limit]

    return ""


def _score_page(page_path: Path, body_text: str, keywords: Sequence[str]) -> float:
    lowered_body = body_text.casefold()
    lowered_name = page_path.stem.replace("-", " ").casefold()
    path_parts = set(page_path.parts)
    score = 0.0

    for keyword in keywords:
        lowered_keyword = keyword.casefold()
        match_count = min(lowered_body.count(lowered_keyword), 5)
        score += float(match_count)
        if lowered_keyword in lowered_name:
            score += 3.0

    if "syntheses" in path_parts:
        score += 2.0
    if "components" in path_parts or "overview" in path_parts:
        score += 0.5

    return score


def _iter_candidate_pages(wiki_root: Path) -> Iterable[Path]:
    for page_path in wiki_root.rglob("*.md"):
        if page_path.name in EXCLUDED_FILE_NAMES:
            continue
        if any(part in EXCLUDED_DIRECTORIES for part in page_path.parts):
            continue
        yield page_path


def lookup_wiki_context(
    problem_text: str,
    concept_candidates: Sequence[str] | None = None,
    target_profile: str = "KR-2022-middle",
    approved_solution_summary: str = "",
    max_pages: int = 4,
    wiki_root: str | Path | None = None,
) -> dict:
    """Return a compact evidence bundle for a post-review wiki enrichment step."""

    concept_candidates = concept_candidates or []
    wiki_root_path = resolve_wiki_root(wiki_root)
    keywords = _normalize_keywords(problem_text, approved_solution_summary, concept_candidates)
    index_text = (wiki_root_path / "index.md").read_text(encoding="utf-8")

    scored_matches: list[WikiMatch] = []
    for page_path in _iter_candidate_pages(wiki_root_path):
        raw_text = page_path.read_text(encoding="utf-8")
        body_text = _strip_frontmatter(raw_text)
        score = _score_page(page_path, body_text, keywords)
        if score <= 0:
            continue
        scored_matches.append(
            WikiMatch(
                path=_relative_page_path(page_path, wiki_root_path),
                score=round(score, 2),
                excerpt=_extract_excerpt(body_text, keywords),
            )
        )

    scored_matches.sort(key=lambda item: (-item.score, item.path))
    context = WikiContext(
        target_profile=target_profile,
        wiki_root=str(wiki_root_path),
        query_keywords=keywords,
        index_excerpt=_collect_index_excerpt(index_text, keywords),
        matches=scored_matches[:max_pages],
        usage_notes=[
            "Run this lookup only after reviewer approval.",
            "Use the result to build a compact learning_context, not to guess the math answer.",
            "If the matches are weak or sparse, mark curriculum mapping and difficulty as tentative.",
        ],
    )
    return asdict(context)


def build_learning_context_seed(
    problem_text: str,
    concept_candidates: Sequence[str] | None = None,
    target_profile: str = "KR-2022-middle",
    approved_solution_summary: str = "",
    max_pages: int = 4,
    wiki_root: str | Path | None = None,
) -> dict:
    """Alias for teams that prefer context-builder naming."""

    return lookup_wiki_context(
        problem_text=problem_text,
        concept_candidates=concept_candidates,
        target_profile=target_profile,
        approved_solution_summary=approved_solution_summary,
        max_pages=max_pages,
        wiki_root=wiki_root,
    )


def main() -> None:
    parser = argparse.ArgumentParser(description="Build a compact wiki context seed.")
    parser.add_argument("problem_text", help="Original math problem text.")
    parser.add_argument(
        "--concept",
        action="append",
        default=[],
        help="Concept candidates extracted by the problem interpreter.",
    )
    parser.add_argument(
        "--approved-solution-summary",
        default="",
        help="Optional short summary of the approved solution.",
    )
    parser.add_argument(
        "--target-profile",
        default="KR-2022-middle",
        help="Curriculum profile used for the enrichment step.",
    )
    parser.add_argument(
        "--max-pages",
        type=int,
        default=4,
        help="Maximum number of wiki pages to keep in the result.",
    )
    parser.add_argument(
        "--wiki-root",
        default=None,
        help="Optional path hint to the repo root or the wiki directory.",
    )
    args = parser.parse_args()

    result = lookup_wiki_context(
        problem_text=args.problem_text,
        concept_candidates=args.concept,
        target_profile=args.target_profile,
        approved_solution_summary=args.approved_solution_summary,
        max_pages=args.max_pages,
        wiki_root=args.wiki_root,
    )
    print(json.dumps(result, ensure_ascii=False, indent=2))


if __name__ == "__main__":
    main()

