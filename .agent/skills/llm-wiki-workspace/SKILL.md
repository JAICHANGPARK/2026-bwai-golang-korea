---
name: "LLM Wiki Workspace"
description: "Maintain this repository as a raw-sources plus persistent-wiki workspace. Use when ingesting docs into wiki, answering from wiki, or running wiki lint."
---

# LLM Wiki Workspace

이 workspace는 일반적인 코드 저장소가 아니라 `docs/`를 원천 문서 집합으로, `wiki/`를 LLM이 유지하는 파생 지식 계층으로 쓰는 개인용 LLM Wiki다.

## Core Rules

- `docs/`는 raw sources처럼 취급한다.
- 원천 문서는 사용자가 명시적으로 요청하지 않는 한 편집하지 않는다.
- 구조화, 요약, 연결, 비교, 모순 표시는 `wiki/`에서 수행한다.
- 새 지식을 반영하면 항상 `wiki/index.md`와 `wiki/log.md`를 함께 갱신한다.
- 질의 중 생긴 가치 있는 분석은 `wiki/queries/` 또는 기존 페이지에 다시 편입한다.

## Read Order

이 workspace 관련 요청이 들어오면 가능하면 아래 순서로 읽는다.

1. `README.md`
2. `AGENTS.md`
3. `wiki/README.md`
4. `wiki/index.md`

질문에 답할 때는 바로 `docs/` 전체를 훑지 말고, 먼저 `wiki/index.md`에서 후보 페이지를 좁힌다.

## Query Workflow

1. `wiki/index.md`를 읽는다.
2. 관련 위키 페이지를 먼저 읽는다.
3. 근거가 부족하거나 최신성이 의심될 때만 `wiki/source-notes/`와 `docs/` 원문을 확인한다.
4. 답변에 가치 있는 비교나 해석이 생기면 위키 반영을 제안하거나 수행한다.

## Ingest Workflow

1. 원천 문서를 읽고 핵심 주장, 범위, 대상 독자, 제약을 파악한다.
2. 기존 `wiki/index.md`에서 이미 있는 관련 페이지를 찾는다.
3. 중복 페이지를 만들지 말고 기존 페이지를 보강한다.
4. 해당 문서의 `wiki/source-notes/...` 페이지를 만들거나 갱신한다.
5. 관련 `overview`, `profile`, `component`, `synthesis` 페이지를 함께 수정한다.
6. `wiki/index.md`를 갱신한다.
7. `wiki/log.md`에 append-only 엔트리를 추가한다.

## Lint Workflow

아래 항목을 점검한다.

- `index.md`에 등록되지 않은 페이지
- frontmatter 누락 페이지
- stale claim
- 중복 페이지
- 서로 충돌하는 주장
- `hands-on`과 `product expansion` 경계가 흐려진 페이지

가능하면 아래 스크립트를 먼저 사용한다.

- `scripts/wiki-search.sh "query"`
- `scripts/wiki-recent.sh 10`
- `scripts/wiki-lint.sh`

## Scope Reminder

이 workspace에서 특히 중요한 범위 구분은 아래다.

- `docs/hands-on-math-agent-session-guide.md`와 `docs/prompt-pack/`은 `hands-on profile`
- `docs/math-agent-system-spec.md`, `docs/math-agent-knowledge-plan.md`, `docs/internal-knowledge-base-build.md`는 `product expansion profile`

두 프로파일을 혼동하지 않는다.

## Companion Hands-On Skill

참석자가 실제 `수학 해설 에이전트 프로토타입`을 만드는 핸즈온이면 아래 스킬을 함께 쓰는 것이 좋다.

- `math-agent-hands-on-starter`
