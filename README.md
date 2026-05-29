# 2026-bwai-golang-korea

이 저장소는 `Build with AI Golang Korea 2026` 이벤트의 `수학 문제 해설 에이전트` 핸즈온 세션 자료를 중심으로 정리한 workspace다.

행사에서는 `AI Studio -> Build -> Antigravity` 흐름으로 수학 해설 에이전트를 설계하고, 터미널 프로토타입과 화면 프로토타입까지 직접 만들어보는 것을 목표로 한다. 이 저장소에는 그 흐름을 바로 따라갈 수 있는 가이드, 프롬프트 팩, handout, starter kit, workspace automation 초안이 함께 들어 있다.

같은 workspace 안에는 행사 이후를 위한 다음 단계 문서도 같이 정리해 둔다.

- `hands-on profile`: 1시간 안에 `AI Studio -> Build -> Antigravity` 흐름으로 설계와 프로토타이핑을 경험하는 워크숍 트랙
- `product expansion profile`: 멀티 에이전트, 내부 지식베이스, 교육과정 매핑, RAG 확장까지 포함하는 제품형 기획 트랙

이 저장소는 `docs/`를 원천 문서 집합으로, `wiki/`를 LLM이 유지하는 구조화된 파생 지식 계층으로 사용한다.

## 저장소 구성

| 경로 | 역할 |
| --- | --- |
| `docs/` | 사람이 작성한 원천 문서와 연구 자료 |
| `docs/prompt-pack/` | 핸즈온 참가자가 바로 복사해 쓸 수 있는 단계별 프롬프트 |
| `docs/math-curriculum-research/` | 한국, 일본, 중국, 미국 수학 교육과정 비교 문서 |
| `docs/math-concept-encyclopedia/` | 국가별 커리큘럼을 개념 백과사전처럼 읽을 수 있게 정리한 자료 |
| `starter-kits/` | `adk-python`, `adk-go` 기준의 runnable starter scaffold 초안 |
| `wiki/` | LLM이 유지하는 마크다운 위키와 합성 지식 계층 |
| `scripts/` | 위키 탐색과 점검용 헬퍼 스크립트 |

## 시작 경로

### 1. 핸즈온 트랙부터 볼 때

행사 핸즈온 흐름으로 바로 들어가려면 아래 순서가 가장 자연스럽다.

- [`docs/hands-on-math-agent-session-guide.md`](./docs/hands-on-math-agent-session-guide.md)
- [`docs/prompt-pack/README.md`](./docs/prompt-pack/README.md)
- [`docs/workspace-distribution-guide.md`](./docs/workspace-distribution-guide.md)
- [`docs/hands-on-assets/README.md`](./docs/hands-on-assets/README.md)
- [`starter-kits/README.md`](./starter-kits/README.md)
- [`docs/llm-wiki-usage-guide.md`](./docs/llm-wiki-usage-guide.md)
- [`docs/gemini-cli-math-agent-step-by-step.md`](./docs/gemini-cli-math-agent-step-by-step.md)
- [`docs/antigravity-math-agent-step-by-step.md`](./docs/antigravity-math-agent-step-by-step.md)
- [`docs/antigravity-workspace-rules-draft.md`](./docs/antigravity-workspace-rules-draft.md)
- [`docs/antigravity-workflow-usage-guide.md`](./docs/antigravity-workflow-usage-guide.md)

핵심 범위는 아래와 같다.

- `RAG`를 사용하지 않는다.
- `Solver/Explainer`와 `Expert Verifier` 두 에이전트만 사용한다.
- 먼저 설계와 명세를 만들고, 그 결과를 `Build`와 `Antigravity`에 재사용한다.

### 2. 제품 확장 트랙부터 볼 때

- [`docs/math-agent-system-spec.md`](./docs/math-agent-system-spec.md)
- [`docs/math-agent-knowledge-plan.md`](./docs/math-agent-knowledge-plan.md)
- [`docs/internal-knowledge-base-build.md`](./docs/internal-knowledge-base-build.md)

이 트랙은 아래 주제를 다룬다.

- 멀티 에이전트 오케스트레이션
- 내부 지식 계층과 검색용 청크 설계
- 교육과정/성취기준 매핑
- 복습 추천, 오답 포인트, RAG 확장

### 3. 현재 정리된 합성 지식을 바로 볼 때

- [`wiki/index.md`](./wiki/index.md)
- [`wiki/overview/project-map.md`](./wiki/overview/project-map.md)
- [`wiki/log.md`](./wiki/log.md)

`wiki/`는 원천 문서를 그대로 반복하지 않고, 범위 구분, 연결, 비교, 충돌, 확장 메모를 구조화해 둔 계층이다.

## 두 가지 프로파일

| 구분 | Hands-on Profile | Product Expansion Profile |
| --- | --- | --- |
| 목표 | 1시간 안에 설계와 앱 생성 경험 | 제품 품질과 확장성 확보 |
| 기준 문서 | `hands-on-math-agent-session-guide.md`, `docs/prompt-pack/` | `math-agent-system-spec.md`, `math-agent-knowledge-plan.md`, `internal-knowledge-base-build.md` |
| 에이전트 수 | `Solver/Explainer` + `Expert Verifier` | 멀티 에이전트 확장 가능 |
| 지식 검색 | 사용하지 않음 | `RAG`, `File Search`, 내부 지식베이스 고려 |
| 세션 초점 | 프롬프트 설계, 출력 계약, 앱 생성 재사용 | 데이터 설계, 검색 계층, 운영 구조 |

행사에서는 보통 `hands-on profile`을 먼저 사용하고, `product expansion profile`은 확장 아이디어나 후속 작업 문서처럼 읽으면 된다. 두 축을 섞지 않고 보는 편이 이해하기 쉽다.

## 추천 읽기 순서

1. [`wiki/overview/project-map.md`](./wiki/overview/project-map.md)에서 저장소 전체 지도를 본다.
2. 목적이 워크숍이면 `hands-on` 문서군을, 제품 기획이면 `product expansion` 문서군을 읽는다.
3. 교육과정 비교가 필요하면 [`docs/math-curriculum-research/README.md`](./docs/math-curriculum-research/README.md)부터 시작한다.
4. 개념 구조와 카드형 지식이 필요하면 [`docs/math-concept-encyclopedia/README.md`](./docs/math-concept-encyclopedia/README.md)와 `wiki/syntheses/`를 본다.
5. 현재 위키 반영 상태와 최근 변경을 확인하려면 [`wiki/index.md`](./wiki/index.md), [`wiki/log.md`](./wiki/log.md)를 읽는다.

## 위키 운영 방식

이 저장소는 일반적인 문서 폴더가 아니라 `raw sources + persistent wiki` 구조를 전제로 운영한다.

- `docs/`는 기본적으로 원천 문서처럼 취급한다.
- 원천 문서는 사용자가 명시적으로 요청하지 않는 한 수정하지 않는다.
- 구조화, 요약, 연결, 비교, 충돌 표시는 `wiki/`에서 수행한다.
- 새 지식을 반영할 때는 `wiki/index.md`와 `wiki/log.md`를 함께 갱신한다.
- 세부 운영 규칙은 [`AGENTS.md`](./AGENTS.md)에 정리되어 있다.

## CLI 헬퍼

가능하면 아래 스크립트를 먼저 사용한다.

```bash
scripts/wiki-search.sh "query"
scripts/wiki-recent.sh 10
scripts/wiki-lint.sh
```

텍스트 검색은 `rg`를 우선 사용한다.

## 현재 초점

현재 위키의 첫 번째 목표는 아래 공통 축을 분명히 연결하는 것이다.

- `Solver/Explainer Agent`
- `Expert Verifier Agent`
- `내부 지식 계층`
- `핸즈온 프로파일`과 `제품 확장 프로파일`의 경계

즉, 이 저장소는 `Build with AI Golang Korea 2026` 핸즈온 자료를 중심에 두고, 그 위에 제품형 확장 문서와 위키 지식 계층을 함께 쌓아 가는 workspace다.
