---
title: Wiki Log
type: overview
status: active
updated: 2026-04-15
source_docs:
  - docs/hands-on-math-agent-session-guide.md
  - docs/prompt-pack/README.md
  - docs/prompt-pack/02a-wiki-enricher-extension.md
  - docs/llm-wiki-usage-guide.md
  - docs/antigravity-math-agent-step-by-step.md
  - docs/gemini-cli-math-agent-step-by-step.md
  - docs/antigravity-workspace-rules-draft.md
  - docs/antigravity-workflow-usage-guide.md
  - docs/workspace-distribution-guide.md
  - docs/math-agent-system-spec.md
  - docs/math-agent-knowledge-plan.md
  - docs/internal-knowledge-base-build.md
  - docs/math-curriculum-research/README.md
  - docs/math-curriculum-research/korea.md
  - docs/math-curriculum-research/comparison.md
  - docs/math-curriculum-research/japan.md
  - docs/math-curriculum-research/china.md
  - docs/math-curriculum-research/us.md
  - docs/math-concept-encyclopedia/README.md
  - docs/math-concept-encyclopedia/korea.md
  - docs/math-concept-encyclopedia/japan.md
  - docs/math-concept-encyclopedia/china.md
  - docs/math-concept-encyclopedia/us.md
  - docs/math-concept-encyclopedia/multilingual-glossary.md
  - docs/math-concept-encyclopedia/formula-examples.md
  - docs/math-concept-encyclopedia/unit-practice-book.md
  - docs/math-concept-encyclopedia/comparative-problem-book.md
tags:
  - log
  - chronology
---

# Wiki Log

## [2026-04-15] refine | step guides mention local wiki lookup helper
- Sources: docs/antigravity-math-agent-step-by-step.md, docs/gemini-cli-math-agent-step-by-step.md
- Updated: wiki/source-notes/antigravity-math-agent-step-by-step.md, wiki/source-notes/gemini-cli-math-agent-step-by-step.md, wiki/components/local-llm-wiki-ops.md, wiki/log.md
- Notes: 두 step guide의 `wiki enricher` 확장에 prompt-only file reading 대신 local `wiki lookup helper`나 `context builder`를 실제 ADK 앱 코드에 넣어야 한다는 점을 명시했다.

## [2026-04-15] refine | step by step wiki enricher extension
- Sources: docs/antigravity-math-agent-step-by-step.md, docs/gemini-cli-math-agent-step-by-step.md
- Updated: wiki/source-notes/antigravity-math-agent-step-by-step.md, wiki/source-notes/gemini-cli-math-agent-step-by-step.md, wiki/components/local-llm-wiki-ops.md, wiki/log.md
- Notes: 두 step-by-step 문서에도 `review approved 이후`에만 `LLM Wiki`를 읽는 `wiki enricher` 선택 확장 단계를 추가해 prompt pack과 hands-on 흐름이 어긋나지 않게 맞췄다.

## [2026-04-15] ingest | wiki enricher extension for prompt pack
- Sources: docs/prompt-pack/02a-wiki-enricher-extension.md, docs/prompt-pack/README.md, docs/prompt-pack/02-agent-design.md, docs/prompt-pack/03-system-prompts.md, docs/prompt-pack/06-integration-spec.md
- Added: wiki/source-notes/prompt-pack-wiki-enricher-extension.md, wiki/components/wiki-enricher-agent.md
- Updated: wiki/components/prompt-pack.md, wiki/components/internal-knowledge-base.md, wiki/profiles/hands-on-profile.md, wiki/source-notes/prompt-pack-readme.md, wiki/index.md, wiki/log.md
- Notes: 기본 two-agent 흐름은 유지하되, reviewer 승인 이후 `LLM Wiki`를 읽어 `learning_context`를 만드는 `wiki enricher` 확장 에이전트와 ADK orchestration 흐름을 prompt pack에 추가했다.

## [2026-04-15] refine | readme event positioning
- Sources: README.md
- Updated: wiki/overview/project-map.md, wiki/index.md, wiki/log.md
- Notes: 저장소의 최상위 정체성을 `Build with AI Golang Korea 2026` 핸즈온 세션 workspace로 명확히 드러내도록 README와 관련 overview 문구를 보정했다.

## [2026-04-15] ingest | shared workspace packaging guide
- Sources: docs/workspace-distribution-guide.md, docs/hands-on-assets/README.md, docs/hands-on-math-agent-session-guide.md
- Added: wiki/source-notes/workspace-distribution-guide.md
- Updated: wiki/source-notes/hands-on-math-agent-session-guide.md, wiki/components/local-llm-wiki-ops.md, wiki/profiles/hands-on-profile.md, wiki/index.md, wiki/log.md
- Notes: 공유 workspace 기준으로 public asset pack과 runnable starter kit을 분리하고, facilitator 전용 handout은 `docs/hands-on-assets/`에서 제외했다.

## [2026-04-15] ingest | antigravity workspace rules and workflows
- Sources: docs/antigravity-workspace-rules-draft.md, docs/antigravity-workflow-usage-guide.md
- Added: wiki/source-notes/antigravity-workspace-rules-draft.md, wiki/source-notes/antigravity-workflow-usage-guide.md, wiki/components/antigravity-workspace-automation.md
- Updated: wiki/components/local-llm-wiki-ops.md, wiki/profiles/hands-on-profile.md, wiki/index.md, wiki/log.md
- Notes: Antigravity의 workspace-level `Rules` 초안과 `Workflows` 사용 가이드를 분리해 추가하고, slash-command처럼 쓸 workflow 파일들을 `.agent/workflows/`에 정리했다.

## [2026-04-15] ingest | gemini cli math agent step by step
- Sources: docs/gemini-cli-math-agent-step-by-step.md
- Added: wiki/source-notes/gemini-cli-math-agent-step-by-step.md
- Updated: wiki/components/local-llm-wiki-ops.md, wiki/profiles/hands-on-profile.md, wiki/index.md, wiki/log.md
- Notes: Gemini CLI 사용자가 `AGENTS.md` 기반 context file과 단계별 프롬프트로 `Phase 1 -> Phase 2 -> Phase 3` 핸즈온을 진행하는 가이드를 편입했다.

## [2026-04-15] ingest | antigravity math agent step by step
- Sources: docs/antigravity-math-agent-step-by-step.md
- Added: wiki/source-notes/antigravity-math-agent-step-by-step.md
- Updated: wiki/components/local-llm-wiki-ops.md, wiki/profiles/hands-on-profile.md, wiki/index.md, wiki/log.md
- Notes: Antigravity에서 스킬 두 개를 함께 쓰며 `Phase 1 -> Phase 2 -> Phase 3`로 진행하는 참가자용 단계별 프롬프트 가이드를 편입했다.

## [2026-04-15] ingest | llm wiki usage guide
- Sources: docs/llm-wiki-usage-guide.md
- Added: wiki/source-notes/llm-wiki-usage-guide.md, wiki/components/local-llm-wiki-ops.md
- Updated: wiki/profiles/hands-on-profile.md, wiki/index.md, wiki/log.md
- Notes: clone 후 `Gemini CLI` 또는 `Antigravity`로 이 저장소를 wiki-first 방식으로 활용하는 운영 가이드를 편입했다.

## [2026-04-06] scaffold | initialize llm wiki
- Sources: docs/hands-on-math-agent-session-guide.md, docs/prompt-pack/README.md, docs/math-agent-system-spec.md, docs/math-agent-knowledge-plan.md, docs/internal-knowledge-base-build.md
- Added: AGENTS.md, wiki/README.md, wiki/index.md, wiki/log.md
- Added: wiki/overview/project-map.md, wiki/profiles/, wiki/components/, wiki/syntheses/, wiki/source-notes/
- Notes: 현재 문서군을 `raw sources + persistent wiki` 구조로 부트스트랩했다.

## [2026-04-06] ingest | hands-on profile bootstrap
- Sources: docs/hands-on-math-agent-session-guide.md, docs/prompt-pack/README.md
- Added: wiki/source-notes/hands-on-math-agent-session-guide.md, wiki/source-notes/prompt-pack-readme.md
- Updated: wiki/profiles/hands-on-profile.md, wiki/components/prompt-pack.md, wiki/overview/project-map.md, wiki/index.md
- Notes: 1시간 워크숍 범위와 prompt pack 재사용 흐름을 분리해 정리했다.

## [2026-04-06] ingest | system spec bootstrap
- Sources: docs/math-agent-system-spec.md
- Added: wiki/source-notes/math-agent-system-spec.md
- Updated: wiki/profiles/product-expansion-profile.md, wiki/components/solver-explainer-agent.md, wiki/components/expert-verifier-agent.md, wiki/overview/project-map.md, wiki/index.md
- Notes: solve first, verify second, retrieve for explanation 흐름을 제품형 기준으로 정리했다.

## [2026-04-06] ingest | knowledge plan bootstrap
- Sources: docs/math-agent-knowledge-plan.md
- Added: wiki/source-notes/math-agent-knowledge-plan.md
- Updated: wiki/profiles/product-expansion-profile.md, wiki/syntheses/knowledge-roadmap.md, wiki/overview/project-map.md, wiki/index.md
- Notes: 데이터 수집 우선순위와 교육과정 비교 관점을 위키에 편입했다.

## [2026-04-06] ingest | internal knowledge base bootstrap
- Sources: docs/internal-knowledge-base-build.md
- Added: wiki/source-notes/internal-knowledge-base-build.md
- Updated: wiki/components/internal-knowledge-base.md, wiki/syntheses/knowledge-roadmap.md, wiki/syntheses/profile-boundary.md, wiki/index.md
- Notes: 원천 데이터, 정규화 데이터, 검색용 청크를 분리하는 설계를 고정했다.

## [2026-04-06] query | math curriculum graph design
- Sources: docs/math-agent-knowledge-plan.md, docs/math-curriculum-research/korea.md
- Added: wiki/queries/2026-04-06-math-curriculum-graph.md
- Updated: wiki/index.md
- Notes: 중1~고3 수학 지식을 학년 허브, 계통 축, 개념 카드, 평가 레이어로 나누는 옵시디언형 구조안을 정리했다.

## [2026-04-06] query | math curriculum graph samples
- Sources: docs/math-agent-knowledge-plan.md, docs/math-curriculum-research/korea.md
- Added: wiki/queries/math-curriculum-graph/middle-1-hub.md, wiki/queries/math-curriculum-graph/middle-2-hub.md, wiki/queries/math-curriculum-graph/functions-strand.md, wiki/queries/math-curriculum-graph/linear-function.md
- Updated: wiki/queries/2026-04-06-math-curriculum-graph.md, wiki/index.md
- Notes: 학년 허브 2개와 함수 계통, 일차함수 개념 카드 샘플을 실제 페이지로 추가했다.

## [2026-04-06] query | math curriculum graph extended samples
- Sources: docs/math-agent-knowledge-plan.md, docs/math-curriculum-research/korea.md, docs/math-curriculum-research/comparison.md
- Added: wiki/queries/math-curriculum-graph/middle-3-hub.md, wiki/queries/math-curriculum-graph/high-1-hub.md, wiki/queries/math-curriculum-graph/statistics-and-probability-strand.md, wiki/queries/math-curriculum-graph/quadratic-function.md, wiki/queries/math-curriculum-graph/csat-2028.md
- Updated: wiki/queries/math-curriculum-graph/middle-2-hub.md, wiki/queries/math-curriculum-graph/functions-strand.md, wiki/queries/math-curriculum-graph/linear-function.md, wiki/queries/2026-04-06-math-curriculum-graph.md, wiki/index.md
- Notes: 중3, 고1, 확통 계통, 이차함수, 2028 수능 평가 레이어를 추가해 샘플 그래프를 세로와 가로 축으로 확장했다.

## [2026-04-06] ingest | research folder overviews
- Sources: docs/math-curriculum-research/README.md, docs/math-curriculum-research/korea.md, docs/math-curriculum-research/comparison.md, docs/math-concept-encyclopedia/README.md, docs/math-concept-encyclopedia/korea.md
- Added: wiki/source-notes/math-curriculum-research-readme.md, wiki/source-notes/math-curriculum-research-korea.md, wiki/source-notes/math-curriculum-research-comparison.md, wiki/source-notes/math-concept-encyclopedia-readme.md, wiki/source-notes/math-concept-encyclopedia-korea.md
- Updated: wiki/index.md
- Notes: 리서치 폴더와 한국 개념 백과사전의 폴더 목적, 핵심 주장, 현재 위키 반영 지점을 source note로 등록했다.

## [2026-04-06] query | research folder coverage audit and graph expansion
- Sources: docs/math-curriculum-research/README.md, docs/math-concept-encyclopedia/README.md, docs/math-curriculum-research/korea.md, docs/math-concept-encyclopedia/korea.md
- Added: wiki/queries/2026-04-06-research-folder-coverage-audit.md, wiki/queries/math-curriculum-graph/high-2-hub.md, wiki/queries/math-curriculum-graph/high-3-hub.md, wiki/queries/math-curriculum-graph/geometry-strand.md, wiki/queries/math-curriculum-graph/trigonometric-ratio.md, wiki/queries/math-curriculum-graph/probability-and-statistics-course.md
- Updated: wiki/queries/math-curriculum-graph/high-1-hub.md, wiki/queries/math-curriculum-graph/statistics-and-probability-strand.md, wiki/queries/2026-04-06-math-curriculum-graph.md, wiki/index.md
- Notes: 고2·고3·기하·삼각비·확통 샘플을 추가하고 두 리서치 폴더의 현재 반영 현황을 감사 노트로 정리했다.

## [2026-04-06] ingest | concept card support docs
- Sources: docs/math-concept-encyclopedia/multilingual-glossary.md, docs/math-concept-encyclopedia/formula-examples.md, docs/math-concept-encyclopedia/unit-practice-book.md, docs/math-concept-encyclopedia/comparative-problem-book.md
- Added: wiki/source-notes/math-concept-encyclopedia-multilingual-glossary.md, wiki/source-notes/math-concept-encyclopedia-formula-examples.md, wiki/source-notes/math-concept-encyclopedia-unit-practice-book.md, wiki/source-notes/math-concept-encyclopedia-comparative-problem-book.md
- Updated: wiki/source-notes/math-concept-encyclopedia-readme.md, wiki/source-notes/math-concept-encyclopedia-korea.md, wiki/source-notes/math-curriculum-research-readme.md, wiki/source-notes/math-curriculum-research-korea.md, wiki/index.md
- Notes: 개념 카드용 alias, 예제, 연습문제, 나라별 비교 문제 자산을 source note 계층으로 편입했다.

## [2026-04-06] ingest | persistent concept cards
- Sources: docs/math-agent-knowledge-plan.md, docs/internal-knowledge-base-build.md, docs/math-curriculum-research/README.md, docs/math-curriculum-research/korea.md, docs/math-concept-encyclopedia/README.md, docs/math-concept-encyclopedia/korea.md, docs/math-concept-encyclopedia/multilingual-glossary.md, docs/math-concept-encyclopedia/formula-examples.md, docs/math-concept-encyclopedia/unit-practice-book.md, docs/math-concept-encyclopedia/comparative-problem-book.md
- Added: wiki/syntheses/concept-card-model.md, wiki/syntheses/concept-cards/linear-equation.md, wiki/syntheses/concept-cards/proportion.md, wiki/syntheses/concept-cards/linear-function.md, wiki/syntheses/concept-cards/trigonometric-ratio.md, wiki/syntheses/concept-cards/probability-distribution.md
- Updated: wiki/components/internal-knowledge-base.md, wiki/syntheses/knowledge-roadmap.md, wiki/overview/project-map.md, wiki/profiles/product-expansion-profile.md, wiki/index.md, wiki/log.md
- Notes: 학년 허브 위에 개념별 지속 문서를 올리는 구조를 도입하고, 정의·수식·증명 스케치·예제를 포함한 첫 개념 카드 다섯 개를 추가했다.

## [2026-04-06] ingest | japan china us curriculum and concept layers
- Sources: docs/math-curriculum-research/japan.md, docs/math-curriculum-research/china.md, docs/math-curriculum-research/us.md, docs/math-concept-encyclopedia/japan.md, docs/math-concept-encyclopedia/china.md, docs/math-concept-encyclopedia/us.md
- Added: wiki/source-notes/math-curriculum-research-japan.md, wiki/source-notes/math-curriculum-research-china.md, wiki/source-notes/math-curriculum-research-us.md, wiki/source-notes/math-concept-encyclopedia-japan.md, wiki/source-notes/math-concept-encyclopedia-china.md, wiki/source-notes/math-concept-encyclopedia-us.md
- Updated: wiki/syntheses/concept-card-model.md, wiki/syntheses/concept-cards/linear-equation.md, wiki/syntheses/concept-cards/proportion.md, wiki/syntheses/concept-cards/linear-function.md, wiki/syntheses/concept-cards/trigonometric-ratio.md, wiki/syntheses/concept-cards/probability-distribution.md, wiki/source-notes/math-curriculum-research-readme.md, wiki/source-notes/math-concept-encyclopedia-readme.md, wiki/queries/2026-04-06-research-folder-coverage-audit.md, wiki/index.md, wiki/log.md
- Notes: 일본·중국·미국 문서를 source-note 계층에 편입하고, 다섯 개 핵심 개념 카드에 국가별 배치 스냅샷을 추가했다.

## [2026-04-09] ingest | concept cards expansion batch 2
- Sources: docs/math-concept-encyclopedia/korea.md, docs/math-concept-encyclopedia/japan.md, docs/math-concept-encyclopedia/china.md, docs/math-concept-encyclopedia/us.md, docs/math-curriculum-research/korea.md, docs/math-curriculum-research/japan.md, docs/math-curriculum-research/china.md, docs/math-curriculum-research/us.md, docs/math-concept-encyclopedia/multilingual-glossary.md, docs/math-concept-encyclopedia/formula-examples.md, docs/math-concept-encyclopedia/unit-practice-book.md, docs/math-concept-encyclopedia/comparative-problem-book.md
- Added: wiki/syntheses/concept-cards/simultaneous-equations.md, wiki/syntheses/concept-cards/quadratic-function.md, wiki/syntheses/concept-cards/counting-principles.md, wiki/syntheses/concept-cards/random-variable.md
- Updated: wiki/syntheses/concept-card-model.md, wiki/syntheses/concept-cards/linear-equation.md, wiki/syntheses/concept-cards/linear-function.md, wiki/syntheses/concept-cards/probability-distribution.md, wiki/source-notes/math-curriculum-research-korea.md, wiki/source-notes/math-curriculum-research-japan.md, wiki/source-notes/math-curriculum-research-china.md, wiki/source-notes/math-curriculum-research-us.md, wiki/source-notes/math-concept-encyclopedia-korea.md, wiki/source-notes/math-concept-encyclopedia-japan.md, wiki/source-notes/math-concept-encyclopedia-china.md, wiki/source-notes/math-concept-encyclopedia-us.md, wiki/source-notes/math-concept-encyclopedia-multilingual-glossary.md, wiki/source-notes/math-concept-encyclopedia-formula-examples.md, wiki/source-notes/math-concept-encyclopedia-unit-practice-book.md, wiki/source-notes/math-concept-encyclopedia-comparative-problem-book.md, wiki/queries/2026-04-06-research-folder-coverage-audit.md, wiki/index.md, wiki/log.md
- Notes: 연립방정식, 이차함수, 경우의 수, 확률변수 카드를 추가하고 기존 개념 카드의 연결선과 source-note provenance를 보강했다.

## [2026-04-09] ingest | concept cards expansion batch 3
- Sources: docs/math-concept-encyclopedia/korea.md, docs/math-concept-encyclopedia/japan.md, docs/math-concept-encyclopedia/china.md, docs/math-concept-encyclopedia/us.md, docs/math-curriculum-research/korea.md, docs/math-curriculum-research/japan.md, docs/math-curriculum-research/china.md, docs/math-curriculum-research/us.md, docs/math-concept-encyclopedia/multilingual-glossary.md, docs/math-concept-encyclopedia/formula-examples.md, docs/math-concept-encyclopedia/unit-practice-book.md, docs/math-concept-encyclopedia/comparative-problem-book.md
- Added: wiki/syntheses/concept-cards/quadratic-equation.md, wiki/syntheses/concept-cards/trigonometric-function.md, wiki/syntheses/concept-cards/conditional-probability.md, wiki/syntheses/concept-cards/statistical-inference.md
- Updated: wiki/syntheses/concept-card-model.md, wiki/syntheses/concept-cards/quadratic-function.md, wiki/syntheses/concept-cards/trigonometric-ratio.md, wiki/syntheses/concept-cards/counting-principles.md, wiki/syntheses/concept-cards/random-variable.md, wiki/syntheses/concept-cards/probability-distribution.md, wiki/source-notes/math-curriculum-research-korea.md, wiki/source-notes/math-curriculum-research-japan.md, wiki/source-notes/math-curriculum-research-china.md, wiki/source-notes/math-curriculum-research-us.md, wiki/source-notes/math-concept-encyclopedia-korea.md, wiki/source-notes/math-concept-encyclopedia-japan.md, wiki/source-notes/math-concept-encyclopedia-china.md, wiki/source-notes/math-concept-encyclopedia-us.md, wiki/source-notes/math-concept-encyclopedia-multilingual-glossary.md, wiki/source-notes/math-concept-encyclopedia-formula-examples.md, wiki/source-notes/math-concept-encyclopedia-unit-practice-book.md, wiki/source-notes/math-concept-encyclopedia-comparative-problem-book.md, wiki/queries/2026-04-06-research-folder-coverage-audit.md, wiki/index.md, wiki/log.md
- Notes: 이차방정식, 삼각함수, 조건부확률, 통계적 추정 카드를 추가하고 기존 카드의 선행·후속 연결과 provenance를 함께 정리했다.

## [2026-04-09] ingest | concept cards expansion batch 4
- Sources: docs/math-concept-encyclopedia/korea.md, docs/math-concept-encyclopedia/japan.md, docs/math-concept-encyclopedia/china.md, docs/math-concept-encyclopedia/us.md, docs/math-curriculum-research/korea.md, docs/math-curriculum-research/japan.md, docs/math-curriculum-research/china.md, docs/math-curriculum-research/us.md, docs/math-concept-encyclopedia/multilingual-glossary.md, docs/math-concept-encyclopedia/formula-examples.md, docs/math-concept-encyclopedia/unit-practice-book.md, docs/math-concept-encyclopedia/comparative-problem-book.md
- Added: wiki/syntheses/concept-cards/equations-of-geometric-figures.md, wiki/syntheses/concept-cards/regression.md
- Updated: wiki/syntheses/concept-card-model.md, wiki/syntheses/concept-cards/quadratic-equation.md, wiki/syntheses/concept-cards/quadratic-function.md, wiki/syntheses/concept-cards/statistical-inference.md, wiki/source-notes/math-curriculum-research-korea.md, wiki/source-notes/math-curriculum-research-japan.md, wiki/source-notes/math-curriculum-research-china.md, wiki/source-notes/math-curriculum-research-us.md, wiki/source-notes/math-concept-encyclopedia-korea.md, wiki/source-notes/math-concept-encyclopedia-japan.md, wiki/source-notes/math-concept-encyclopedia-china.md, wiki/source-notes/math-concept-encyclopedia-us.md, wiki/source-notes/math-concept-encyclopedia-multilingual-glossary.md, wiki/source-notes/math-concept-encyclopedia-formula-examples.md, wiki/source-notes/math-concept-encyclopedia-unit-practice-book.md, wiki/source-notes/math-concept-encyclopedia-comparative-problem-book.md, wiki/queries/2026-04-06-research-folder-coverage-audit.md, wiki/index.md, wiki/log.md
- Notes: 도형의 방정식과 회귀 카드를 추가하고, 현재까지 반영된 개념 카드 기준으로 남은 교과 개념 backlog를 분야별로 다시 정리했다.

## [2026-04-09] ingest | concept cards expansion batch 5
- Sources: docs/math-concept-encyclopedia/korea.md, docs/math-concept-encyclopedia/japan.md, docs/math-concept-encyclopedia/china.md, docs/math-concept-encyclopedia/us.md, docs/math-concept-encyclopedia/multilingual-glossary.md, docs/math-concept-encyclopedia/formula-examples.md, docs/math-curriculum-research/us.md
- Added: wiki/syntheses/concept-cards/prime-factorization.md, wiki/syntheses/concept-cards/integers-and-rational-numbers.md, wiki/syntheses/concept-cards/variables-and-expressions.md, wiki/syntheses/concept-cards/algebraic-manipulation.md, wiki/syntheses/concept-cards/linear-inequality.md, wiki/syntheses/concept-cards/square-roots-and-real-numbers.md, wiki/syntheses/concept-cards/polynomials-and-factorization.md, wiki/syntheses/concept-cards/basic-geometry-and-construction.md, wiki/syntheses/concept-cards/geometric-proof.md, wiki/syntheses/concept-cards/congruence.md, wiki/syntheses/concept-cards/similarity.md, wiki/syntheses/concept-cards/circle.md, wiki/syntheses/concept-cards/pythagorean-theorem.md, wiki/syntheses/concept-cards/inverse-proportion.md, wiki/syntheses/concept-cards/function-foundations.md, wiki/syntheses/concept-cards/exponential-and-logarithmic-functions.md, wiki/syntheses/concept-cards/sequences.md, wiki/syntheses/concept-cards/limits.md, wiki/syntheses/concept-cards/differentiation.md, wiki/syntheses/concept-cards/integration.md, wiki/syntheses/concept-cards/sets-and-propositions.md, wiki/syntheses/concept-cards/conic-sections.md, wiki/syntheses/concept-cards/vectors.md, wiki/syntheses/concept-cards/spatial-coordinates.md, wiki/syntheses/concept-cards/data-organization.md, wiki/syntheses/concept-cards/correlation-and-scatter-plots.md, wiki/syntheses/concept-cards/normal-distribution.md, wiki/syntheses/concept-cards/sampling.md, wiki/syntheses/concept-cards/practical-statistics.md, wiki/syntheses/concept-cards/economics-math.md, wiki/syntheses/concept-cards/mathematics-for-ai.md, wiki/syntheses/concept-cards/workplace-math.md, wiki/syntheses/concept-cards/mathematics-and-culture.md, wiki/syntheses/concept-cards/integrated-mathematics.md
- Updated: wiki/syntheses/concept-card-model.md, wiki/index.md, wiki/queries/2026-04-06-research-folder-coverage-audit.md, wiki/source-notes/math-concept-encyclopedia-korea.md, wiki/source-notes/math-concept-encyclopedia-japan.md, wiki/source-notes/math-concept-encyclopedia-china.md, wiki/source-notes/math-concept-encyclopedia-us.md, wiki/source-notes/math-concept-encyclopedia-multilingual-glossary.md, wiki/source-notes/math-concept-encyclopedia-formula-examples.md, wiki/source-notes/math-curriculum-research-us.md, wiki/log.md
- Notes: 이전 감사 노트에 남아 있던 대표 개념 backlog를 한 번에 concept card layer로 올리고, 인덱스와 provenance를 새 규모에 맞게 재정리했다.

## [2026-04-09] ingest | concept cards expansion batch 6
- Sources: docs/math-concept-encyclopedia/korea.md, docs/math-concept-encyclopedia/japan.md, docs/math-concept-encyclopedia/us.md, docs/math-concept-encyclopedia/multilingual-glossary.md, docs/math-concept-encyclopedia/unit-practice-book.md, docs/math-curriculum-research/korea.md, docs/math-curriculum-research/us.md, docs/math-curriculum-research/japan.md
- Added: wiki/syntheses/concept-cards/matrix.md, wiki/syntheses/concept-cards/mathematical-induction.md, wiki/syntheses/concept-cards/continuity.md, wiki/syntheses/concept-cards/correlation-coefficient.md, wiki/syntheses/concept-cards/geometric-transformations.md, wiki/syntheses/concept-cards/common-math-1.md, wiki/syntheses/concept-cards/common-math-2.md, wiki/syntheses/concept-cards/korean-algebra-course.md, wiki/syntheses/concept-cards/calculus-1.md, wiki/syntheses/concept-cards/probability-and-statistics-course.md, wiki/syntheses/concept-cards/calculus-2.md, wiki/syntheses/concept-cards/korean-geometry-course.md, wiki/syntheses/concept-cards/basic-math-1-and-2.md, wiki/syntheses/concept-cards/algebra-1.md, wiki/syntheses/concept-cards/geometry-course.md, wiki/syntheses/concept-cards/algebra-2.md, wiki/syntheses/concept-cards/precalculus.md, wiki/syntheses/concept-cards/calculus-course.md, wiki/syntheses/concept-cards/statistics-and-data-analysis.md, wiki/syntheses/concept-cards/advanced-tracks.md
- Updated: wiki/components/internal-knowledge-base.md, wiki/syntheses/concept-card-model.md, wiki/syntheses/concept-cards/basic-geometry-and-construction.md, wiki/syntheses/concept-cards/simultaneous-equations.md, wiki/syntheses/concept-cards/limits.md, wiki/syntheses/concept-cards/correlation-and-scatter-plots.md, wiki/syntheses/concept-cards/regression.md, wiki/syntheses/concept-cards/integrated-mathematics.md, wiki/index.md, wiki/queries/2026-04-06-research-folder-coverage-audit.md, wiki/source-notes/math-curriculum-research-korea.md, wiki/source-notes/math-curriculum-research-us.md, wiki/source-notes/math-curriculum-research-japan.md, wiki/source-notes/math-concept-encyclopedia-korea.md, wiki/source-notes/math-concept-encyclopedia-japan.md, wiki/source-notes/math-concept-encyclopedia-us.md, wiki/source-notes/math-concept-encyclopedia-multilingual-glossary.md, wiki/source-notes/math-concept-encyclopedia-unit-practice-book.md, wiki/log.md
- Notes: 세부 하위 개념 카드와 한국·미국 과목 wrapper 카드를 추가해 concept layer를 과목·경로 단위까지 확장했다.

## [2026-04-09] query | remaining curriculum backlog refresh
- Sources: docs/math-curriculum-research/comparison.md, docs/math-curriculum-research/korea.md, docs/math-curriculum-research/japan.md, docs/math-curriculum-research/china.md, docs/math-curriculum-research/us.md, docs/math-concept-encyclopedia/korea.md, docs/math-concept-encyclopedia/japan.md, docs/math-concept-encyclopedia/china.md, docs/math-concept-encyclopedia/us.md
- Added: wiki/queries/2026-04-09-remaining-curriculum-backlog.md
- Updated: wiki/index.md, wiki/log.md
- Notes: 예전 대표 개념 backlog는 완료로 보고, 현재 실제 남은 일감을 나라별 wrapper·비교 synthesis·세부 하위 개념 기준으로 다시 정리했다.

## [2026-04-09] ingest | concept cards expansion batch 7
- Sources: docs/math-concept-encyclopedia/japan.md, docs/math-concept-encyclopedia/china.md, docs/math-concept-encyclopedia/us.md, docs/math-curriculum-research/japan.md, docs/math-curriculum-research/china.md, docs/math-curriculum-research/us.md, docs/math-concept-encyclopedia/unit-practice-book.md, docs/math-concept-encyclopedia/comparative-problem-book.md
- Added: wiki/syntheses/concept-cards/complex-numbers.md, wiki/syntheses/concept-cards/complex-plane.md, wiki/syntheses/concept-cards/mathematical-modeling-and-inquiry.md, wiki/syntheses/concept-cards/polar-ideas.md, wiki/syntheses/concept-cards/sequences-and-series.md, wiki/syntheses/concept-cards/ap-statistics.md
- Updated: wiki/syntheses/concept-card-model.md, wiki/syntheses/concept-cards/algebra-2.md, wiki/syntheses/concept-cards/precalculus.md, wiki/syntheses/concept-cards/statistics-and-data-analysis.md, wiki/syntheses/concept-cards/advanced-tracks.md, wiki/syntheses/concept-cards/vectors.md, wiki/syntheses/concept-cards/quadratic-equation.md, wiki/syntheses/concept-cards/korean-algebra-course.md, wiki/syntheses/concept-cards/complex-plane.md, wiki/index.md, wiki/queries/2026-04-06-research-folder-coverage-audit.md, wiki/queries/2026-04-09-remaining-curriculum-backlog.md, wiki/source-notes/math-concept-encyclopedia-us.md, wiki/source-notes/math-concept-encyclopedia-japan.md, wiki/source-notes/math-concept-encyclopedia-china.md, wiki/source-notes/math-curriculum-research-us.md, wiki/source-notes/math-curriculum-research-japan.md, wiki/source-notes/math-curriculum-research-china.md, wiki/source-notes/math-concept-encyclopedia-unit-practice-book.md, wiki/source-notes/math-concept-encyclopedia-comparative-problem-book.md, wiki/log.md
- Notes: 복소수 계열, 모델링, 미국 Precalculus·Advanced Statistics 보조 카드를 추가해 concept layer를 75개까지 확장하고 backlog를 세부 하위 카드 기준으로 재정렬했다.

## [2026-04-09] refine | proof-rich anchor concept cards
- Sources: docs/math-concept-encyclopedia/korea.md, docs/math-concept-encyclopedia/japan.md, docs/math-concept-encyclopedia/china.md, docs/math-concept-encyclopedia/us.md, docs/math-concept-encyclopedia/formula-examples.md, docs/math-concept-encyclopedia/unit-practice-book.md, docs/math-concept-encyclopedia/comparative-problem-book.md
- Updated: wiki/syntheses/concept-card-model.md, wiki/syntheses/concept-cards/linear-function.md, wiki/syntheses/concept-cards/trigonometric-ratio.md, wiki/syntheses/concept-cards/quadratic-equation.md, wiki/index.md, wiki/log.md
- Notes: 개념 카드를 빠른 인덱스형에서 `정의-정리-증명-예제-오개념-그래프 링크`가 보이는 정석형 카드로 보강하기 시작했다.

## [2026-04-09] ingest | fine-grained concept cards batch 8
- Sources: docs/math-concept-encyclopedia/korea.md, docs/math-concept-encyclopedia/us.md, docs/math-concept-encyclopedia/japan.md, docs/math-concept-encyclopedia/china.md, docs/math-concept-encyclopedia/multilingual-glossary.md, docs/math-concept-encyclopedia/formula-examples.md, docs/math-concept-encyclopedia/unit-practice-book.md, docs/math-curriculum-research/korea.md
- Added: wiki/syntheses/concept-cards/radians.md, wiki/syntheses/concept-cards/trigonometric-identities.md, wiki/syntheses/concept-cards/complex-number-operations.md, wiki/syntheses/concept-cards/matrix-operations.md, wiki/syntheses/concept-cards/derivative.md
- Updated: wiki/syntheses/concept-card-model.md, wiki/syntheses/concept-cards/trigonometric-function.md, wiki/syntheses/concept-cards/complex-numbers.md, wiki/syntheses/concept-cards/matrix.md, wiki/syntheses/concept-cards/differentiation.md, wiki/syntheses/concept-cards/korean-algebra-course.md, wiki/syntheses/concept-cards/precalculus.md, wiki/syntheses/concept-cards/algebra-2.md, wiki/syntheses/concept-cards/calculus-course.md, wiki/syntheses/concept-cards/common-math-1.md, wiki/index.md, wiki/queries/2026-04-09-remaining-curriculum-backlog.md, wiki/queries/2026-04-06-research-folder-coverage-audit.md, wiki/source-notes/math-concept-encyclopedia-us.md, wiki/source-notes/math-concept-encyclopedia-korea.md, wiki/source-notes/math-concept-encyclopedia-japan.md, wiki/source-notes/math-concept-encyclopedia-china.md, wiki/source-notes/math-concept-encyclopedia-multilingual-glossary.md, wiki/source-notes/math-concept-encyclopedia-formula-examples.md, wiki/source-notes/math-concept-encyclopedia-unit-practice-book.md, wiki/source-notes/math-curriculum-research-korea.md, wiki/log.md
- Notes: 세부 하위 개념을 정석형 카드로 추가하고, 관련 개념 카드와 과목 wrapper에 왕복 그래프 링크를 보강했다.

## [2026-04-09] ingest | fine-grained concept cards batch 9
- Sources: docs/math-concept-encyclopedia/korea.md, docs/math-concept-encyclopedia/japan.md, docs/math-concept-encyclopedia/china.md, docs/math-concept-encyclopedia/us.md, docs/math-concept-encyclopedia/multilingual-glossary.md, docs/math-concept-encyclopedia/formula-examples.md, docs/math-concept-encyclopedia/unit-practice-book.md, docs/math-concept-encyclopedia/comparative-problem-book.md, docs/math-curriculum-research/japan.md, docs/math-curriculum-research/china.md, docs/math-curriculum-research/us.md
- Added: wiki/syntheses/concept-cards/continuity-properties.md, wiki/syntheses/concept-cards/trigonometric-graphs.md, wiki/syntheses/concept-cards/complex-division.md, wiki/syntheses/concept-cards/derivative-graph-interpretation.md, wiki/syntheses/concept-cards/ap-calculus.md, wiki/syntheses/concept-cards/ib-mathematics.md, wiki/syntheses/concept-cards/japan-middle-1.md, wiki/syntheses/concept-cards/japan-middle-2.md, wiki/syntheses/concept-cards/japan-middle-3.md, wiki/syntheses/concept-cards/japan-math-1.md, wiki/syntheses/concept-cards/japan-math-a.md, wiki/syntheses/concept-cards/japan-math-2.md, wiki/syntheses/concept-cards/japan-math-b.md, wiki/syntheses/concept-cards/japan-math-3.md, wiki/syntheses/concept-cards/japan-math-c.md, wiki/syntheses/concept-cards/china-grade-7-semester-1.md, wiki/syntheses/concept-cards/china-grade-7-semester-2.md, wiki/syntheses/concept-cards/china-grade-8-semester-1.md, wiki/syntheses/concept-cards/china-grade-8-semester-2.md, wiki/syntheses/concept-cards/china-grade-9-semester-1.md, wiki/syntheses/concept-cards/china-grade-9-semester-2.md, wiki/syntheses/concept-cards/china-required-curriculum.md, wiki/syntheses/concept-cards/china-elective-required-curriculum.md, wiki/syntheses/concept-cards/china-elective-concept-clusters.md, wiki/syntheses/concept-cards/us-grade-6.md, wiki/syntheses/concept-cards/us-grade-7.md, wiki/syntheses/concept-cards/us-grade-8.md
- Updated: wiki/syntheses/concept-card-model.md, wiki/syntheses/concept-cards/continuity.md, wiki/syntheses/concept-cards/trigonometric-function.md, wiki/syntheses/concept-cards/complex-number-operations.md, wiki/syntheses/concept-cards/derivative.md, wiki/syntheses/concept-cards/advanced-tracks.md, wiki/syntheses/concept-cards/calculus-course.md, wiki/syntheses/concept-cards/complex-numbers.md, wiki/syntheses/concept-cards/matrix.md, wiki/syntheses/concept-cards/differentiation.md, wiki/syntheses/concept-cards/korean-algebra-course.md, wiki/syntheses/concept-cards/precalculus.md, wiki/syntheses/concept-cards/algebra-2.md, wiki/syntheses/concept-cards/common-math-1.md, wiki/index.md, wiki/queries/2026-04-09-remaining-curriculum-backlog.md, wiki/queries/2026-04-06-research-folder-coverage-audit.md, wiki/source-notes/math-concept-encyclopedia-us.md, wiki/source-notes/math-concept-encyclopedia-korea.md, wiki/source-notes/math-concept-encyclopedia-japan.md, wiki/source-notes/math-concept-encyclopedia-china.md, wiki/source-notes/math-concept-encyclopedia-multilingual-glossary.md, wiki/source-notes/math-concept-encyclopedia-formula-examples.md, wiki/source-notes/math-concept-encyclopedia-unit-practice-book.md, wiki/source-notes/math-concept-encyclopedia-comparative-problem-book.md, wiki/source-notes/math-curriculum-research-korea.md, wiki/source-notes/math-curriculum-research-japan.md, wiki/source-notes/math-curriculum-research-china.md, wiki/source-notes/math-curriculum-research-us.md, wiki/log.md
- Notes: 세부 하위 카드와 일본·중국·미국 wrapper 카드를 추가해 docs 기반 남은 교과 개념 및 wrapper backlog를 concept layer에서 사실상 닫았다.

## [2026-04-09] ingest | curriculum hubs and comparison syntheses
- Sources: docs/math-curriculum-research/comparison.md, docs/math-curriculum-research/korea.md, docs/math-curriculum-research/japan.md, docs/math-curriculum-research/china.md, docs/math-curriculum-research/us.md, docs/math-concept-encyclopedia/korea.md, docs/math-concept-encyclopedia/japan.md, docs/math-concept-encyclopedia/china.md, docs/math-concept-encyclopedia/us.md, docs/math-concept-encyclopedia/formula-examples.md, docs/math-concept-encyclopedia/unit-practice-book.md, docs/math-concept-encyclopedia/comparative-problem-book.md
- Added: wiki/syntheses/korea-curriculum-hub.md, wiki/syntheses/japan-curriculum-hub.md, wiki/syntheses/china-curriculum-hub.md, wiki/syntheses/us-curriculum-hub.md, wiki/syntheses/course-track-hub.md, wiki/syntheses/cross-country-function-progression.md, wiki/syntheses/cross-country-geometry-and-proof.md, wiki/syntheses/cross-country-statistics-and-data.md, wiki/syntheses/assessment-systems-comparison.md, wiki/syntheses/example-pattern-library.md, wiki/syntheses/misconception-pattern-map.md, wiki/syntheses/assessment-patterns.md
- Updated: wiki/index.md, wiki/syntheses/concept-card-model.md, wiki/queries/2026-04-09-remaining-curriculum-backlog.md, wiki/queries/2026-04-06-research-folder-coverage-audit.md, wiki/source-notes/math-curriculum-research-comparison.md, wiki/source-notes/math-curriculum-research-korea.md, wiki/source-notes/math-curriculum-research-japan.md, wiki/source-notes/math-curriculum-research-china.md, wiki/source-notes/math-curriculum-research-us.md, wiki/source-notes/math-concept-encyclopedia-korea.md, wiki/source-notes/math-concept-encyclopedia-japan.md, wiki/source-notes/math-concept-encyclopedia-china.md, wiki/source-notes/math-concept-encyclopedia-us.md, wiki/source-notes/math-concept-encyclopedia-multilingual-glossary.md, wiki/source-notes/math-concept-encyclopedia-formula-examples.md, wiki/source-notes/math-concept-encyclopedia-unit-practice-book.md, wiki/source-notes/math-concept-encyclopedia-comparative-problem-book.md, wiki/log.md
- Notes: 나라별 허브, 4개국 비교 synthesis, 예제·오개념·평가 패턴 synthesis를 추가해 research docs 기반 persistent synthesis layer를 완성했다.

## [2026-04-09] refine | persistent hubs and textbook-grade anchor cards
- Sources: docs/math-curriculum-research/korea.md, docs/math-curriculum-research/comparison.md, docs/math-agent-knowledge-plan.md, docs/math-concept-encyclopedia/korea.md, docs/math-concept-encyclopedia/japan.md, docs/math-concept-encyclopedia/china.md, docs/math-concept-encyclopedia/us.md, docs/math-concept-encyclopedia/multilingual-glossary.md, docs/math-concept-encyclopedia/formula-examples.md, docs/math-concept-encyclopedia/unit-practice-book.md, docs/math-concept-encyclopedia/comparative-problem-book.md
- Added: wiki/syntheses/functions-strand.md, wiki/syntheses/geometry-strand.md, wiki/syntheses/statistics-and-probability-strand.md, wiki/syntheses/middle-1-hub.md, wiki/syntheses/middle-2-hub.md, wiki/syntheses/middle-3-hub.md, wiki/syntheses/high-1-hub.md, wiki/syntheses/high-2-hub.md, wiki/syntheses/high-3-hub.md, wiki/syntheses/csat-2028.md
- Updated: wiki/syntheses/concept-card-model.md, wiki/syntheses/korea-curriculum-hub.md, wiki/syntheses/assessment-systems-comparison.md, wiki/syntheses/concept-cards/linear-equation.md, wiki/syntheses/concept-cards/algebraic-manipulation.md, wiki/syntheses/concept-cards/linear-function.md, wiki/syntheses/concept-cards/quadratic-function.md, wiki/syntheses/concept-cards/similarity.md, wiki/syntheses/concept-cards/vectors.md, wiki/syntheses/concept-cards/conditional-probability.md, wiki/syntheses/concept-cards/random-variable.md, wiki/syntheses/concept-cards/probability-distribution.md, wiki/syntheses/concept-cards/probability-and-statistics-course.md, wiki/index.md, wiki/log.md
- Notes: query sample 의존성을 persistent hub로 치환하고 concept/wrapper를 `card_role`과 callout으로 분리했으며, 핵심 개념 카드와 확통 축의 수학적 정밀도와 예제 밀도를 보강했다.

## [2026-04-10] refine | textbook coverage expansion wave
- Sources: docs/math-concept-encyclopedia/korea.md, docs/math-concept-encyclopedia/japan.md, docs/math-concept-encyclopedia/china.md, docs/math-concept-encyclopedia/us.md, docs/math-curriculum-research/korea.md, docs/math-curriculum-research/japan.md, docs/math-curriculum-research/china.md, docs/math-curriculum-research/us.md, docs/math-concept-encyclopedia/multilingual-glossary.md, docs/math-concept-encyclopedia/formula-examples.md, docs/math-concept-encyclopedia/unit-practice-book.md, docs/math-concept-encyclopedia/comparative-problem-book.md
- Updated: wiki/syntheses/concept-cards/proportion.md, wiki/syntheses/concept-cards/function-foundations.md, wiki/syntheses/concept-cards/polynomials-and-factorization.md, wiki/syntheses/concept-cards/exponential-and-logarithmic-functions.md, wiki/syntheses/concept-cards/counting-principles.md, wiki/syntheses/concept-cards/correlation-and-scatter-plots.md, wiki/syntheses/concept-cards/statistical-inference.md, wiki/syntheses/concept-cards/equations-of-geometric-figures.md, wiki/syntheses/concept-cards/trigonometric-function.md, wiki/syntheses/concept-cards/limits.md, wiki/syntheses/concept-cards/continuity.md, wiki/syntheses/concept-cards/integration.md, wiki/syntheses/concept-cards/prime-factorization.md, wiki/syntheses/concept-cards/integers-and-rational-numbers.md, wiki/syntheses/concept-cards/variables-and-expressions.md, wiki/syntheses/concept-cards/linear-inequality.md, wiki/syntheses/concept-cards/basic-geometry-and-construction.md, wiki/syntheses/concept-cards/congruence.md, wiki/syntheses/concept-cards/circle.md, wiki/syntheses/concept-cards/pythagorean-theorem.md, wiki/syntheses/concept-cards/sampling.md, wiki/syntheses/concept-cards/correlation-coefficient.md, wiki/syntheses/concept-cards/inverse-proportion.md, wiki/syntheses/concept-cards/sequences.md, wiki/syntheses/concept-cards/sequences-and-series.md, wiki/syntheses/concept-cards/mathematical-induction.md, wiki/syntheses/concept-cards/complex-numbers.md, wiki/syntheses/concept-cards/complex-plane.md, wiki/syntheses/concept-cards/conic-sections.md, wiki/syntheses/concept-cards/spatial-coordinates.md, wiki/index.md, wiki/log.md
- Notes: 정석형 2차 보강을 함수·확통·기하·미적분·기초 수와 식·고등 심화 카드까지 넓혔고, Worked Examples, Common Pitfalls, 증명 스케치, persistent graph links를 대폭 확장했다.

## [2026-04-10] refine | remaining concept cards normalization
- Sources: docs/math-concept-encyclopedia/korea.md, docs/math-concept-encyclopedia/japan.md, docs/math-concept-encyclopedia/china.md, docs/math-concept-encyclopedia/us.md, docs/math-curriculum-research/korea.md, docs/math-curriculum-research/japan.md, docs/math-curriculum-research/china.md, docs/math-curriculum-research/us.md, docs/math-concept-encyclopedia/multilingual-glossary.md, docs/math-concept-encyclopedia/formula-examples.md, docs/math-concept-encyclopedia/unit-practice-book.md, docs/math-concept-encyclopedia/comparative-problem-book.md
- Updated: wiki/syntheses/concept-cards/data-organization.md, wiki/syntheses/concept-cards/normal-distribution.md, wiki/syntheses/concept-cards/regression.md, wiki/syntheses/concept-cards/practical-statistics.md, wiki/syntheses/concept-cards/economics-math.md, wiki/syntheses/concept-cards/mathematics-for-ai.md, wiki/syntheses/concept-cards/mathematical-modeling-and-inquiry.md, wiki/syntheses/concept-cards/mathematics-and-culture.md, wiki/syntheses/concept-cards/workplace-math.md, wiki/syntheses/concept-cards/prime-factorization.md, wiki/syntheses/concept-cards/integers-and-rational-numbers.md, wiki/syntheses/concept-cards/variables-and-expressions.md, wiki/syntheses/concept-cards/linear-inequality.md, wiki/syntheses/concept-cards/simultaneous-equations.md, wiki/syntheses/concept-cards/square-roots-and-real-numbers.md, wiki/syntheses/concept-cards/matrix.md, wiki/syntheses/concept-cards/mathematical-induction.md, wiki/syntheses/concept-cards/polynomials-and-factorization.md, wiki/syntheses/concept-cards/proportion.md, wiki/syntheses/concept-cards/function-foundations.md, wiki/syntheses/concept-cards/inverse-proportion.md, wiki/syntheses/concept-cards/sequences.md, wiki/syntheses/concept-cards/sequences-and-series.md, wiki/syntheses/concept-cards/geometric-proof.md, wiki/syntheses/concept-cards/geometric-transformations.md, wiki/syntheses/concept-cards/sets-and-propositions.md, wiki/syntheses/concept-cards/polar-ideas.md, wiki/syntheses/concept-cards/differentiation.md, wiki/syntheses/concept-cards/exponential-and-logarithmic-functions.md, wiki/index.md, wiki/log.md
- Notes: 남아 있던 concept 카드의 교재형 섹션 구조를 정리해 모든 concept 카드가 최소 두 가지 이상의 정석형 서술 요소를 갖도록 맞췄고, 확통·응용 카드에는 실제 해석형 예제와 함정까지 보강했다.

## [2026-04-10] ingest | fine-grained curriculum gap cards
- Sources: docs/math-concept-encyclopedia/korea.md, docs/math-concept-encyclopedia/japan.md, docs/math-concept-encyclopedia/china.md, docs/math-concept-encyclopedia/us.md, docs/math-curriculum-research/china.md, docs/math-curriculum-research/comparison.md, docs/math-concept-encyclopedia/multilingual-glossary.md, docs/math-concept-encyclopedia/unit-practice-book.md, docs/math-concept-encyclopedia/comparative-problem-book.md
- Added: wiki/syntheses/concept-cards/fractional-expressions.md, wiki/syntheses/concept-cards/parallelogram.md, wiki/syntheses/concept-cards/projection-and-orthographic-views.md, wiki/syntheses/concept-cards/variance.md, wiki/syntheses/concept-cards/standard-deviation.md, wiki/syntheses/concept-cards/confidence-interval.md, wiki/syntheses/concept-cards/hypothesis-testing.md, wiki/syntheses/concept-cards/equation-of-a-line.md, wiki/syntheses/concept-cards/equation-of-a-circle.md, wiki/syntheses/concept-cards/parabola.md, wiki/syntheses/concept-cards/ellipse.md, wiki/syntheses/concept-cards/hyperbola.md
- Updated: wiki/syntheses/concept-cards/china-grade-8-semester-1.md, wiki/syntheses/concept-cards/china-grade-8-semester-2.md, wiki/syntheses/concept-cards/china-grade-9-semester-2.md, wiki/syntheses/concept-cards/data-organization.md, wiki/syntheses/concept-cards/random-variable.md, wiki/syntheses/concept-cards/probability-distribution.md, wiki/syntheses/concept-cards/statistical-inference.md, wiki/syntheses/concept-cards/normal-distribution.md, wiki/syntheses/concept-cards/probability-and-statistics-course.md, wiki/syntheses/concept-cards/common-math-2.md, wiki/syntheses/concept-cards/equations-of-geometric-figures.md, wiki/syntheses/concept-cards/conic-sections.md, wiki/syntheses/concept-cards/circle.md, wiki/syntheses/geometry-strand.md, wiki/syntheses/statistics-and-probability-strand.md, wiki/syntheses/concept-card-model.md, wiki/queries/2026-04-09-remaining-curriculum-backlog.md, wiki/queries/2026-04-06-research-folder-coverage-audit.md, wiki/source-notes/math-concept-encyclopedia-china.md, wiki/source-notes/math-concept-encyclopedia-korea.md, wiki/source-notes/math-concept-encyclopedia-japan.md, wiki/source-notes/math-concept-encyclopedia-us.md, wiki/source-notes/math-curriculum-research-china.md, wiki/source-notes/math-curriculum-research-korea.md, wiki/source-notes/math-curriculum-research-comparison.md, wiki/source-notes/math-concept-encyclopedia-multilingual-glossary.md, wiki/source-notes/math-concept-encyclopedia-unit-practice-book.md, wiki/source-notes/math-concept-encyclopedia-comparative-problem-book.md, wiki/index.md, wiki/log.md
- Notes: 세분화가 덜 되어 있던 중국 중학·해석기하·확통 노드를 독립 concept card로 올리고, 상위 허브·과목 wrapper·source-note·backlog 메모를 새 그래프 기준으로 다시 연결했다.
