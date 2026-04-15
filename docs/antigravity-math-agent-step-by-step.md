# Antigravity 단계별 프롬프트 가이드

이 문서는 `Antigravity`에서 이 저장소를 열고, 단계별 프롬프트로 `수학 해설 에이전트 프로토타입`을 차근차근 만드는 흐름을 안내한다.

처음에는 어디서부터 시작해야 할지 조금 막막할 수 있는데, 이 문서대로 `Phase 1 -> Phase 2 -> Phase 3` 순서로만 가면 된다. 한 번에 완성형 앱을 만들려고 하기보다, 먼저 돌아가는 흐름을 만들고 점점 붙여 나간다고 생각하면 편하다.

핵심은 아래 두 스킬을 같이 쓰는 것이다.

- `llm-wiki-workspace`
- `math-agent-hands-on-starter`

첫 번째 스킬은 이 저장소를 `raw sources + persistent wiki` workspace로 읽게 한다.  
두 번째 스킬은 구현 흐름을 `Phase 1 -> Phase 2 -> Phase 3` 순서로 자연스럽게 잡아준다.

`Rules`와 `Workflows`를 더 적극적으로 쓰고 싶다면 아래 문서를 함께 보면 된다.

- [`antigravity-workspace-rules-draft.md`](./antigravity-workspace-rules-draft.md)
- [`antigravity-workflow-usage-guide.md`](./antigravity-workflow-usage-guide.md)

## 1. 목표

이 문서를 따라가면 아래 3가지를 순서대로 만들게 된다.

1. 터미널에서 실행 가능한 단일 수학 해설 에이전트
2. subagent 오케스트레이션이 보이는 터미널 프로토타입
3. 텍스트와 이미지 입력을 받을 수 있는 화면 프로토타입

## 2. 시작 전에 알아둘 규칙

미리 딱 기억해두면 좋은 건 아래 정도다.

- 기본 런타임은 `adk-python`
- `adk-go`는 선택 트랙
- `RAG`, DB, 로그인, 개인화는 기본 범위에서 제외함
- 먼저 `터미널 실행 경로`를 만든 뒤 화면으로 확장함
- `review`가 승인하기 전에는 최종 해설을 확정하지 않음

## 3. Antigravity에서 처음 보내는 프롬프트

workspace를 연 뒤 첫 메시지는 아래처럼 보내면 무난하다.

```text
Use the llm-wiki-workspace skill and the math-agent-hands-on-starter skill.

Read README.md, AGENTS.md, wiki/README.md, and wiki/index.md first.

We are building a prototype for a math explanation agent.
Use adk-python unless I explicitly ask for adk-go.
Start with Phase 1 only.

Before coding, briefly summarize:
1. what this repo is,
2. what Phase 1 should build,
3. what you will not build yet.
```

이 첫 프롬프트는 아래 세 가지를 한 번에 정리해준다.

- workspace를 제대로 읽게 하기
- 스킬 두 개를 모두 활성화하기
- 한 번에 너무 많이 만들지 않게 `Phase 1 only`를 고정하기

## 4. Phase 1: 단일 수학 해설 에이전트

여기서는 욕심내지 말고, "터미널에서 문제 하나 넣으면 설명이 나온다"까지 만들면 충분하다.

### 권장 프롬프트

```text
Use the math-agent-hands-on-starter skill.

Build Phase 1 only.
Create a minimal terminal-runnable math explanation agent.

Requirements:
- Use adk-python.
- Accept a plain text math problem from stdin or a CLI argument.
- Return:
  1. problem interpretation
  2. final answer
  3. step-by-step explanation
- Keep the implementation minimal and locally runnable.
- Do not add RAG, auth, database, image handling, or UI yet.
- Include a short run command.

Stop when the CLI flow works for at least 3 sample text problems.
```

### Phase 1에서 확인할 것

이 단계에서는 아래만 확인되면 된다.

- 터미널에서 문제를 넣을 수 있는가
- 정답이 나오는가
- 풀이 단계가 나오는가
- 실행 명령이 분명한가

### Phase 1 테스트용 문제

테스트 문제는 아래 정도면 충분하다.

```text
Solve for x: 3x - 7 = 11
```

```text
Given f(x) = 2x + 3, find f(5).
```

```text
Which values of x satisfy x^2 - 5x + 6 = 0?
```

### Phase 1이 끝났을 때 보내는 확인 프롬프트

잘 돌아가기 시작하면 아래 프롬프트로 상태를 정리해보면 편하다.

```text
Summarize the Phase 1 result briefly.
Show me:
- the entry point file
- the run command
- the 3 sample problems you tested
- what must change before moving to Phase 2
```

## 5. Phase 2: subagent 오케스트레이션

Phase 1이 끝났다면 이제 역할을 나눌 차례다. 여기서 중요한 건 "에이전트를 많이 늘리는 것"보다 "각 역할이 왜 분리되어야 하는지 보이게 만드는 것"이다.

### 권장 프롬프트

```text
Use the math-agent-hands-on-starter skill.

Build Phase 2 only.
Take the existing terminal math explainer and refactor it into subagents.

Required subagents:
- problem interpretation agent
- problem solving agent
- solution review agent

Requirements:
- Keep the app runnable from the terminal.
- Add an orchestrator that calls the subagents in sequence.
- The review agent must return an explicit status such as approved, revise, or fail.
- Do not finalize the user-facing explanation unless the review status is approved.
- Show enough terminal output to make the orchestration visible.
- Keep the code minimal and demo-friendly.

Stop when the orchestration flow works for at least 3 text problems.
```

### Phase 2에서 확인할 것

이 단계에서는 구조가 눈에 보이는지가 핵심이다.

- subagent가 분리되어 보이는가
- orchestrator가 있는가
- `approved`, `revise`, `fail` 같은 상태가 있는가
- review가 통과하기 전 최종 해설이 막히는가

### Phase 2가 끝났을 때 보내는 확인 프롬프트

```text
Summarize the Phase 2 result briefly.
Show me:
- the orchestrator entry point
- the three subagents and their responsibilities
- one approved case
- one revise or fail case
- what must change before moving to Phase 3
```

## 6. Phase 3: 화면과 백엔드

이제야 화면을 붙인다. 이미 Phase 2까지 잘 만들어졌다면, 여기서는 새 기능을 크게 더하기보다 기존 흐름을 화면으로 자연스럽게 보여주는 데 집중하면 된다.

### 권장 프롬프트

```text
Use the math-agent-hands-on-starter skill.

Build Phase 3 only.
Extend the existing orchestrated math explainer into a backend plus UI prototype.

Requirements:
- Reuse the Phase 2 orchestration as the backend core.
- Add a frontend that supports:
  - text input
  - image upload
  - submit action
  - status display
  - final answer and explanation card
  - failure card
- The UI must show clear states such as analyzing, solving, and reviewing.
- If image quality is poor or the review fails, show a friendly fallback instead of pretending success.
- Keep the frontend thin and the backend simple.
- Do not add RAG, auth, persistence, or unrelated product features.

Use the sample SVG problem files under docs/hands-on-assets/sample-problems/ if useful.

Stop when a local end-to-end flow works from the screen.
```

### Phase 3에서 확인할 것

UI는 화려함보다 상태가 잘 보이는지가 더 중요하다.

- 텍스트 입력이 되는가
- 이미지 업로드가 되는가
- 실행 상태가 보이는가
- 최종 결과 카드가 있는가
- 실패 카드가 있는가
- review 실패 시 거짓 성공을 보여주지 않는가

### 이미지 테스트용 파일

아래 샘플을 그대로 써도 된다.

- `docs/hands-on-assets/sample-problems/problem-01-linear-equation.svg`
- `docs/hands-on-assets/sample-problems/problem-02-function-value.svg`
- `docs/hands-on-assets/sample-problems/problem-03-quadratic-choice.svg`

### Phase 3이 끝났을 때 보내는 확인 프롬프트

```text
Summarize the Phase 3 result briefly.
Show me:
- how to run the backend
- how to run the frontend
- the supported inputs
- the UI states
- one image-based demo path
- one fallback path
```

## 7. 마지막 다듬기 프롬프트

프로토타입이 이미 동작할 때만 아래 프롬프트를 넣는 편이 좋다. 아직 기본 흐름이 안 잡혔다면 polish보다 먼저 동작을 맞추는 쪽이 낫다.

```text
Improve the current demo without changing the core architecture.

Focus on:
- clearer loading, reviewing, and failure states
- cleaner Korean labels
- simpler card layout
- more stable text and image demo flow

Do not add new major features.
Do not add RAG, auth, persistence, or extra product scope.
```

## 8. 최종 확인 체크리스트

마지막에는 아래 정도만 확인되면 데모하기에 충분하다.

- 정상 텍스트 문제 1개
- 조금 까다로운 텍스트 문제 1개
- 이미지 문제 1개
- 실패 또는 fallback 사례 1개

확인할 화면 상태:

- analyzing
- solving
- reviewing
- approved result
- fallback or failure

## 9. 막혔을 때 쓰는 프롬프트

### 구현 범위가 너무 커졌을 때

```text
Reduce scope.
Keep only the minimum needed for the current phase.
Do not jump to the next phase yet.
```

### UI부터 너무 복잡해졌을 때

```text
Simplify the UI.
Keep only:
- one text input
- one image upload
- one submit button
- one status area
- one final result card
- one failure card
```

### review 게이트가 안 지켜질 때

```text
Fix the review gate.
The app must not show the final explanation unless the review status is approved.
```

### 이미지 처리가 불안정할 때

```text
Add a clearer fallback for image parsing uncertainty.
If the image is unreadable or ambiguous, ask for a clearer upload instead of pretending success.
```

## 10. 가장 짧은 버전

진짜 짧게 시작하려면 아래 3개만 순서대로 넣어도 된다.

1. 첫 시작

```text
Use the llm-wiki-workspace skill and the math-agent-hands-on-starter skill.
Read README.md, AGENTS.md, wiki/README.md, and wiki/index.md first.
Use adk-python and start with Phase 1 only.
```

2. Phase 2 전환

```text
Use the math-agent-hands-on-starter skill.
Phase 1 works now.
Move to Phase 2 and refactor it into interpretation, solving, and review subagents.
Keep it terminal-runnable.
```

3. Phase 3 전환

```text
Use the math-agent-hands-on-starter skill.
Phase 2 works now.
Move to Phase 3 and build a backend plus UI prototype with text input, image upload, status states, result card, and failure card.
```
