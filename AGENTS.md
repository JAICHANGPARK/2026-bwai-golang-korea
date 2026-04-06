# LLM Wiki Agent Guide

이 작업공간은 `docs/`를 원천 문서 집합으로, `wiki/`를 LLM이 유지하는 파생 지식 계층으로 사용하는 개인용 LLM Wiki다.

## 운영 원칙

- 기본적으로 `docs/`는 `raw sources`처럼 취급한다.
- 원천 문서는 사용자가 명시적으로 수정 요청하지 않는 한 편집하지 않는다.
- 지식의 구조화, 요약, 연결, 모순 표시는 모두 `wiki/`에서 수행한다.
- 새 지식을 반영할 때는 항상 `wiki/index.md`와 `wiki/log.md`를 함께 갱신한다.
- 질의 중 생긴 가치 있는 분석은 필요하면 `wiki/queries/` 또는 기존 주제 페이지에 다시 편입한다.

## 디렉터리 계약

- `docs/`
  - 사람이 작성한 원천 문서.
  - 현재 프로젝트의 기준 자료다.
- `wiki/`
  - LLM이 유지하는 마크다운 위키.
  - Obsidian 친화적으로 유지한다.
- `wiki/source-notes/`
  - 원천 문서별 요약과 핵심 주장, 반영 범위를 정리한다.
- `wiki/overview/`
  - 프로젝트 전체 맥락과 지도 역할을 하는 페이지를 둔다.
- `wiki/profiles/`
  - 핸즈온 프로파일, 제품 확장 프로파일처럼 범위가 다른 운영 모드를 정리한다.
- `wiki/components/`
  - 에이전트, 지식 계층, 프롬프트 팩 같은 구성요소 페이지를 둔다.
- `wiki/syntheses/`
  - 여러 문서를 가로지르는 비교, 해석, 로드맵을 정리한다.
- `wiki/queries/`
  - 사용자 질의로부터 나온 분석 결과를 저장할 때 사용한다.
- `wiki/lint/`
  - 정기 lint 결과나 정리 메모를 저장할 때 사용한다.

## 페이지 프런트매터

`wiki/` 아래 일반 페이지는 아래 YAML frontmatter를 기본으로 사용한다.

```yaml
---
title: 간단한 제목
type: overview | profile | component | synthesis | source-note | query-note | lint-report
status: active
updated: YYYY-MM-DD
source_docs:
  - docs/example.md
tags:
  - tag-name
---
```

권장 규칙:

- `title`은 사람 읽기 기준 제목으로 쓴다.
- `type`은 위 분류 중 하나를 쓴다.
- `updated`는 수정 시마다 오늘 날짜로 바꾼다.
- `source_docs`는 상대 경로로 적고, 실제 근거가 되는 원천 문서만 남긴다.
- `tags`는 검색용 최소 세트만 유지한다.

## 파일 네이밍 규칙

- 파일명은 `kebab-case.md`를 사용한다.
- 원천 문서에 대한 노트는 가능한 한 원문 파일명과 맞춘다.
- 질의 결과 페이지는 `wiki/queries/YYYY-MM-DD-topic.md` 형식을 선호한다.
- lint 결과 페이지는 `wiki/lint/YYYY-MM-DD-summary.md` 형식을 선호한다.

## 페이지 작성 형식

페이지는 장황하게 길어지지 않도록 유지한다. 기본 섹션은 아래를 권장한다.

- `## Summary`
- `## Key Points`
- `## Connections`
- `## Open Questions`
- `## Sources`

소스 노트는 아래 형식을 선호한다.

- `## Source Summary`
- `## Notable Claims`
- `## Pages Updated`
- `## Open Questions`
- `## Raw Source`

## Ingest Workflow

새 문서가 추가되거나 기존 `docs/` 문서를 위키에 반영할 때는 아래 순서를 따른다.

1. 원천 문서를 읽고 핵심 주장, 범위, 대상 독자, 제약을 파악한다.
2. 기존 `wiki/index.md`를 먼저 읽고 이미 있는 관련 페이지를 찾는다.
3. 없으면 새 페이지를 만들고, 있으면 중복을 만들지 말고 기존 페이지를 보강한다.
4. 해당 원천 문서에 대한 `wiki/source-notes/...` 페이지를 만들거나 갱신한다.
5. 관련 `overview`, `profile`, `component`, `synthesis` 페이지를 함께 수정한다.
6. 새 문서가 기존 주장과 충돌하면 숨기지 말고 명시적으로 적는다.
7. `wiki/index.md`에 새 페이지를 등록하거나 설명을 갱신한다.
8. `wiki/log.md`에 아래 형식으로 append-only 엔트리를 추가한다.

```md
## [2026-04-06] ingest | 문서 제목
- Sources: docs/example.md
- Added: wiki/source-notes/example.md
- Updated: wiki/components/example-component.md, wiki/index.md
- Notes: 한 줄 요약
```

## Query Workflow

질문에 답할 때는 바로 원천 문서 전체를 훑지 말고 아래 순서를 우선한다.

1. `wiki/index.md`를 읽어 관련 페이지 후보를 좁힌다.
2. 관련 위키 페이지를 읽어 이미 축적된 합성 지식을 우선 활용한다.
3. 근거가 부족하거나 최신성이 의심되면 해당 `source-note`와 `docs/` 원문을 다시 확인한다.
4. 답변에 가치 있는 새 비교나 해석이 생기면 위키에 다시 반영할지 판단한다.
5. 사용자가 원하면 결과를 `wiki/queries/`에 저장한다.

답변 원칙:

- 핸즈온 범위와 제품 확장 범위를 혼동하지 않는다.
- 교육과정, RAG, 지식 인프라 얘기는 어떤 프로파일 기준인지 명확히 적는다.
- 문서에 없는 내용은 추론이라고 표시하거나 open question으로 남긴다.

## Lint Workflow

정기적으로 아래를 점검한다.

- `index.md`에 등록되지 않은 페이지
- frontmatter 누락 페이지
- 더 이상 최신 주장이 아닌 페이지
- 범위가 겹치지만 통합되지 않은 중복 페이지
- 서로 충돌하는 주장
- 핸즈온 문서와 제품형 문서의 경계가 흐려진 페이지
- 비어 있는 `Open Questions`가 너무 많은데 실제 액션이 없는 페이지

필요하면 `wiki/lint/`에 결과를 저장하고 `wiki/log.md`에 아래 형식으로 남긴다.

```md
## [2026-04-06] lint | health-check
- Checked: wiki/index.md, wiki/components/, wiki/profiles/
- Findings: orphan page 1개, stale claim 1개
- Follow-up: wiki/lint/2026-04-06-health-check.md
```

## 검색 및 운영 도구

가능하면 아래 스크립트를 먼저 사용한다.

- `scripts/wiki-search.sh "query"`
- `scripts/wiki-recent.sh 10`
- `scripts/wiki-lint.sh`

`rg`가 있으면 `rg`를 우선 사용한다.

## 이 작업공간의 현재 해석

- `docs/hands-on-math-agent-session-guide.md`와 `docs/prompt-pack/`은 `핸즈온 프로파일` 축이다.
- `docs/math-agent-system-spec.md`, `docs/math-agent-knowledge-plan.md`, `docs/internal-knowledge-base-build.md`는 `제품 확장 프로파일` 축이다.
- 이 저장소의 첫 번째 위키 목표는 두 축의 경계를 분명히 하면서, 공통 구성요소인 `Solver/Explainer`, `Expert Verifier`, `내부 지식 계층`을 연결하는 것이다.
