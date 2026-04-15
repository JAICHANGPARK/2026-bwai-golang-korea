# Antigravity Workflow Usage Guide

이 문서는 이 저장소에서 `Antigravity`의 workspace-level `Workflows`를 어떻게 활용하면 편한지 정리한다.

## 1. 이 저장소에서 Workflows가 필요한 이유

이 저장소에는 자주 반복되는 작업이 크게 두 가지 있다.

- `LLM Wiki` 운영
- `수학 해설 프로토타입` 구현

특히 수학 해설 프로토타입은 순서를 잘 잡아두는 게 중요하다.

1. terminal single agent
2. terminal subagent orchestration
3. backend plus UI

이 순서를 매번 긴 프롬프트로 다시 설명하지 않고 workflow로 묶어두면, 범위가 흔들릴 일이 훨씬 줄어든다.

## 2. Rule / Workflow / Skill 역할 분리

처음에는 셋이 비슷해 보일 수 있는데, 아래처럼 나눠 생각하면 꽤 단순하다.

- `Rules`
  - 항상 적용되는 제약
  - workspace의 기본 해석
- `Workflows`
  - 반복되는 작업 순서
  - phase별 실행 레시피
- `Skills`
  - 재사용할 도메인 지식과 설계 패턴

이 저장소에서는 보통 아래처럼 나누면 가장 다루기 편하다.

- `Rule`
  - `docs/`와 `wiki/` 역할
  - 기본 구현 제약
- `Workflow`
  - `Phase 1`, `Phase 2`, `Phase 3`, `smoke test`, `polish`
- `Skill`
  - `llm-wiki-workspace`
  - `math-agent-hands-on-starter`

## 3. 포함된 Workflows

이 저장소에는 아래 workflow 파일이 이미 들어 있다.

- `.agent/workflows/phase_1_terminal_single_agent.md`
- `.agent/workflows/phase_2_terminal_subagents.md`
- `.agent/workflows/phase_3_fullstack_ui.md`
- `.agent/workflows/smoke_test_math_agent.md`
- `.agent/workflows/polish_math_agent_demo.md`

## 4. 권장 실행 순서

처음에는 아래 순서대로만 가도 충분하다.

### 첫 시작

1. workspace rules를 먼저 켠다.
2. 필요한 skill이 workspace에 보이는지 확인한다.
3. `phase_1_terminal_single_agent` workflow를 실행한다.

### Phase 1이 끝난 뒤

1. `smoke_test_math_agent` workflow를 실행한다.
2. 정상 경로가 보이면 `phase_2_terminal_subagents` workflow로 넘어간다.

### Phase 2가 끝난 뒤

1. 다시 `smoke_test_math_agent` workflow를 실행한다.
2. review gate가 보이면 `phase_3_fullstack_ui` workflow로 넘어간다.

### UI가 붙은 뒤

1. `smoke_test_math_agent` workflow를 한 번 더 실행한다.
2. 그 다음 `polish_math_agent_demo` workflow를 실행한다.

## 5. 언제 Workflow를 쓰고, 언제 직접 프롬프트를 쓰나

workflow를 쓰는 편이 좋은 경우:

- 작업 순서가 반복될 때
- 같은 phase를 여러 번 시작할 때
- 구현 범위를 자꾸 넓히는 경향이 있을 때
- 팀이나 여러 사용자가 같은 흐름을 따라야 할 때

직접 프롬프트가 더 나은 경우:

- 현재 phase 안에서 작은 수정만 할 때
- 아주 구체적인 버그 수정만 할 때
- 현재 상태 요약, 설명, 비교 같은 짧은 요청을 할 때

## 6. 가장 실용적인 패턴

실제로는 아래 패턴으로 쓰면 가장 덜 헷갈린다.

1. workspace rules는 항상 켠다.
2. 시작 phase는 workflow로 연다.
3. 세부 수정은 직접 프롬프트로 한다.
4. phase 종료 전에는 smoke test workflow를 다시 돌린다.

요약하면:

- 큰 흐름은 `Workflow`
- 작은 수정은 `Prompt`
- 항상 지켜야 하는 건 `Rule`

## 7. 권장 조합

### LLM Wiki 작업 중심

- Rules:
  - `LLM Wiki Workspace Conventions`
- Skill:
  - `llm-wiki-workspace`

### 수학 해설 프로토타입 중심

- Rules:
  - `Math Agent Prototype Scope`
- Skills:
  - `llm-wiki-workspace`
  - `math-agent-hands-on-starter`
- Workflows:
  - `phase_1_terminal_single_agent`
  - `phase_2_terminal_subagents`
  - `phase_3_fullstack_ui`
  - `smoke_test_math_agent`
  - `polish_math_agent_demo`

## 8. 운영 팁

- workflow 이름은 조금 길어도 의미가 바로 읽히는 편이 낫다.
- rules에는 절차보다 기본 제약을 넣고, workflow에는 실제 순서를 넣는다.
- skill에는 현재 작업 순서보다 도메인 구조와 설계 패턴을 넣는 편이 좋다.
- 한 workflow가 너무 많은 책임을 가지기 시작하면 phase를 다시 나누는 편이 더 편하다.
