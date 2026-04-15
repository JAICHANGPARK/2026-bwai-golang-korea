---
name: "Math Agent Hands-On Starter"
description: "Guide participants through a 3-phase math explanation agent build: terminal single-agent, terminal subagent orchestration, then backend plus UI prototype with text and image input."
---

# Math Agent Hands-On Starter

이 스킬은 참석자가 `수학 해설 에이전트`를 실제 프로토타입으로 만들도록 돕는 `starter skill`이다.

목표는 프롬프트만 길게 쓰게 하는 것이 아니라, 아래 3단계를 차례대로 완성하게 만드는 것이다.

1. 터미널에서 실행 가능한 단일 수학 해설 에이전트
2. 터미널에서 실행 가능한 subagent 오케스트레이션
3. 텍스트와 이미지를 입력받아 화면에서 결과를 보여주는 웹 프로토타입

## Use When

- 참석자가 핸즈온에서 `수학 해설 에이전트`를 처음 만들 때
- `Antigravity`에서 단계별로 프로토타입을 확장할 때
- `CLI -> orchestration -> UI` 흐름을 잃지 않게 하고 싶을 때

## Companion Skill

이 저장소를 `LLM Wiki workspace`로 읽어야 하는 작업이면 아래 스킬과 함께 쓰는 것이 좋다.

- `llm-wiki-workspace`

이 경우 먼저 `README.md`, `AGENTS.md`, `wiki/README.md`, `wiki/index.md`를 읽고, 설계 지식은 wiki-first로 가져온다.

## Defaults

참석자가 별도 선택을 주지 않으면 아래를 기본값으로 둔다.

- 구현 순서: `Phase 1 -> Phase 2 -> Phase 3`
- 기본 런타임: `adk-python`
- 대안 런타임: `adk-go`
- 프로토타입 범위:
  - `RAG 없음`
  - `DB 없음`
  - `로그인 없음`
  - `개인화 없음`
  - `배포 최적화 없음`

## Required Build Order

이 스킬을 사용할 때는 아래 순서를 건너뛰지 않는다.

### Phase 1

터미널에서 실행 가능한 `단일 수학 해설 에이전트`를 만든다.

필수 요구사항:

- 텍스트 수학문제를 CLI 인자 또는 stdin으로 받는다.
- 문제 해석, 정답, 단계별 풀이를 출력한다.
- 로컬에서 바로 실행 가능해야 한다.
- 코드보다 `실행 가능성`을 우선한다.

### Phase 2

Phase 1 결과를 `subagents`로 분리한다.

권장 기본 구성:

- `문제 해석 에이전트`
- `문제 풀이 에이전트`
- `문제 풀이 검토 에이전트`

필수 요구사항:

- 터미널에서 계속 실행 가능해야 한다.
- 오케스트레이터가 subagent를 순서대로 호출한다.
- `검토 에이전트`가 승인하기 전에는 최종 해설을 확정하지 않는다.
- 실패 또는 수정 필요 상태를 명시적으로 출력한다.

### Optional Extension

기본 two-agent 흐름이 안정적으로 동작한 뒤에는 `wiki knowledge enricher`를 확장 에이전트로 추가할 수 있다.

권장 역할:

- `문제의 난이도` 추정
- `핵심 개념` 정리
- `관련 교육과정` 후보 제안
- `연관 학습 주제` 정리
- `참고한 wiki page` 기록

필수 원칙:

- `검토 에이전트 approved 이후`에만 실행한다.
- 정답을 맞히기 위해 wiki를 참고하지 않는다.
- `learning_context`처럼 짧은 구조화 결과만 final explainer에 넘긴다.
- wiki 근거가 약하면 추정이라고 표시한다.

### Phase 3

Phase 2 결과를 `백엔드 + 화면` 프로토타입으로 확장한다.

필수 요구사항:

- 텍스트 입력
- 이미지 업로드
- 에이전트 실행 요청
- 상태 표시
  - analyzing
  - solving
  - reviewing
- 최종 결과 카드
- 실패 카드

필수 원칙:

- 화면 개발은 Phase 2가 동작한 뒤에 시작한다.
- 프론트는 얇게, 백엔드는 orchestrator 재사용 중심으로 만든다.
- raw intermediate output을 그대로 최종 사용자 화면에 노출하지 않는다.

## Quality Bar

이 스킬로 만든 결과물은 최소한 아래를 만족해야 한다.

- 정상 텍스트 문제 2개 이상 테스트
- 함정 또는 수정 유도 문제 1개 이상 테스트
- 이미지 입력 실패 또는 판독 애매 상황 1개 이상 처리
- Phase 2 이후에는 `review status`가 있어야 함
- Phase 3 이후에는 `UI state`가 있어야 함

## Recommended Runtime Choices

선택 기준은 아래처럼 단순하게 가져간다.

- `adk-python`
  - 참석자가 가장 빨리 결과를 보고 싶을 때
  - orchestration과 backend 확장을 빠르게 이어가고 싶을 때
- `adk-go`
  - 참석자가 Go 기반 CLI 흐름을 선호할 때
  - 터미널 프로토타입을 Go 중심으로 잡고 싶을 때

런타임은 초반에 한 번 고르면 핸즈온 중간에 바꾸지 않는 것이 좋다.

## What To Avoid

아래는 핸즈온 범위 밖이다.

- RAG 추가
- 교육과정 검색 인프라 추가
- 로그인/회원 기능
- 저장 히스토리
- 운영용 모니터링
- 복잡한 디자인 시스템
- 모바일 최적화에 과도한 시간 사용

## Prompt Templates

아래 템플릿을 그대로 시작점으로 써도 된다.

- [templates/phase-1-terminal-single-agent.md](./templates/phase-1-terminal-single-agent.md)
- [templates/phase-2-terminal-subagents.md](./templates/phase-2-terminal-subagents.md)
- [templates/phase-3-fullstack-ui.md](./templates/phase-3-fullstack-ui.md)
- [templates/demo-checklist.md](./templates/demo-checklist.md)
- [templates/sample-problems.md](./templates/sample-problems.md)

## Working Style

이 스킬을 사용하는 에이전트는 아래 방식으로 일한다.

1. 현재 참가자가 어느 phase인지 확인한다.
2. 지정이 없으면 `Phase 1`부터 시작한다.
3. 지정이 없으면 `adk-python`을 기본으로 잡는다.
4. 해당 phase의 최소 동작을 먼저 만든다.
5. 먼저 데모가 되게 하고, 그 다음 구조를 다듬는다.
6. 다음 phase로 넘어가기 전 실행 경로를 확인한다.

## Deliverables

핸즈온 종료 시 산출물은 아래 3개가 모두 있어야 한다.

1. 터미널에서 돌릴 수 있는 수학 해설 프로토타입
2. subagent orchestration이 드러나는 실행 흐름
3. 텍스트와 이미지 입력을 받을 수 있는 화면 프로토타입
