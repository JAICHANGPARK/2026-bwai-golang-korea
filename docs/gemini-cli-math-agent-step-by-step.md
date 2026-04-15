# Gemini CLI 단계별 프롬프트 가이드

이 문서는 이 저장소 루트에서 `gemini`를 실행하고, 단계별 프롬프트로 `수학 해설 에이전트 프로토타입`을 차근차근 만드는 흐름을 안내한다.

`Antigravity`와 달리 `Gemini CLI`에서는 workspace skill을 직접 쓰지 않는다.  
대신 이 저장소의 `.gemini/settings.json`이 `AGENTS.md`를 project context file로 읽게 하므로, `AGENTS.md + README + wiki/index.md`를 기반으로 진행한다.

처음에는 그냥 긴 세션에서 하나씩 이어 붙인다고 생각하면 된다. 먼저 터미널에서 돌고, 그 다음 subagent로 나누고, 마지막에 화면을 붙이는 흐름이다. 기본형이 안정적으로 돌아가면 마지막에 `wiki enricher` 확장을 선택해서 붙일 수도 있다.

## 1. 목표

이 문서를 따라가면 아래 3가지를 기본 흐름으로 만들게 된다.

1. 터미널에서 실행 가능한 단일 수학 해설 에이전트
2. subagent 오케스트레이션이 보이는 터미널 프로토타입
3. 텍스트와 이미지 입력을 받을 수 있는 화면 프로토타입

선택 확장:

4. `LLM Wiki`를 읽어 난이도, 핵심 개념, 교육과정, 연관 학습 주제를 보강하는 `wiki enricher`

## 2. 시작 전에 알아둘 규칙

먼저 아래 정도만 기억해두면 된다.

- 기본 런타임은 `adk-python`
- `adk-go`는 선택 트랙
- `RAG`, DB, 로그인, 개인화는 기본 범위에서 제외함
- 먼저 `터미널 실행 경로`를 만든 뒤 화면으로 확장함
- `review`가 승인하기 전에는 최종 해설을 확정하지 않음
- 학습 메타데이터가 필요하면 `review approved 이후`에만 `wiki enrichment`를 붙임

## 3. 설치와 실행

공식 `Gemini CLI` 문서 기준으로는 아래 중 하나로 시작하면 된다.

```bash
npm install -g @google/gemini-cli
```

또는

```bash
brew install gemini-cli
```

인증도 어렵지 않다. 아래처럼 시작하면 된다.

```bash
gemini
```

또는 API 키를 직접 쓸 수 있다.

```bash
export GEMINI_API_KEY="YOUR_API_KEY"
gemini
```

이 저장소 루트에서 실행하면 프로젝트 `.gemini/settings.json`이 `AGENTS.md`를 context file로 읽는다. 그래서 첫 프롬프트부터 repo 구조를 좀 더 잘 따라오는 편이다.

## 4. Gemini CLI에서 처음 보내는 프롬프트

터미널에서 `gemini`를 실행한 뒤 첫 메시지는 아래처럼 보내면 무난하다.

```text
Read README.md, AGENTS.md, wiki/README.md, and wiki/index.md first.

We are building a prototype for a math explanation agent in this repository.
Use adk-python unless I explicitly ask for adk-go.
Start with Phase 1 only.

Before coding, briefly summarize:
1. what this repo is,
2. what Phase 1 should build,
3. what you will not build yet.
```

이 첫 프롬프트는 아래 세 가지를 한 번에 정리해준다.

- workspace를 제대로 읽게 하기
- 한 번에 너무 많이 만들지 못하게 `Phase 1 only`를 고정하기
- `docs/`와 `wiki/` 역할을 먼저 이해하게 하기

## 5. Phase 1: 단일 수학 해설 에이전트

여기서는 "일단 돌게 만들기"가 목표다. 코드가 아주 예쁘지 않아도 괜찮고, 터미널에서 문제를 넣었을 때 답과 설명이 나오는 흐름만 먼저 잡으면 된다.

### 권장 프롬프트

```text
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

이 단계에서는 아래만 확인되면 충분하다.

- 터미널에서 문제를 넣을 수 있는가
- 정답이 나오는가
- 풀이 단계가 나오는가
- 실행 명령이 분명한가

### Phase 1 테스트용 문제

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

동작이 잡히면 아래 프롬프트로 상태를 짧게 정리해보면 편하다.

```text
Summarize the Phase 1 result briefly.
Show me:
- the entry point file
- the run command
- the 3 sample problems you tested
- what must change before moving to Phase 2
```

## 6. Phase 2: subagent 오케스트레이션

Phase 2에서는 역할 분리가 핵심이다. 에이전트 수를 늘리는 게 목적이 아니라, 문제 해석, 풀이, 검토가 서로 다른 책임이라는 걸 구조로 보여주면 된다.

### 권장 프롬프트

```text
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

이 단계에서는 역할과 상태가 분명히 드러나는지 보면 된다.

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

## 7. Phase 3: 화면과 백엔드

이제 마지막 단계다. 여기서는 새 구조를 더 만드는 것보다, Phase 2에서 만든 흐름을 사용자가 화면에서 이해할 수 있게 보여주는 데 집중하면 된다.

### 권장 프롬프트

```text
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

UI는 화려함보다도 상태 전달이 잘 되는지가 중요하다.

- 텍스트 입력이 되는가
- 이미지 업로드가 되는가
- 실행 상태가 보이는가
- 최종 결과 카드가 있는가
- 실패 카드가 있는가
- review 실패 시 거짓 성공을 보여주지 않는가

### 이미지 테스트용 파일

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

## 8. 선택 확장: Wiki Enricher

기본 `solver + reviewer + UI` 흐름이 안정적으로 돌아간 뒤에만 이 확장을 붙이는 편이 좋다. 여기서는 정답 풀이를 바꾸지 않고, `LLM Wiki`를 읽어 학습 정보만 보강한다.

### 권장 프롬프트

```text
Keep the current approved solution flow as-is.
Add an optional wiki enricher step after review approval.

Requirements:
- Run the wiki enrichment step only after the review status is approved.
- Read wiki/index.md first.
- Read only the minimum relevant wiki pages.
- Build a compact learning_context with:
  - difficulty
  - core concepts
  - related curriculum placement
  - follow-up topics
  - wiki basis pages
- Do not use wiki content to guess the math answer.
- If wiki evidence is weak, mark the mapping as tentative.
- Show the learning metadata in a separate learning info card in the UI.
```

### 이 확장에서 확인할 것

- `review approved 이후`에만 실행되는가
- `learning_context`가 별도 구조로 보이는가
- 난이도, 개념, 교육과정, 연관 주제가 분리되어 나오는가
- wiki 근거가 약할 때 추정이라고 표시하는가
- 정답 풀이와 학습 정보가 서로 섞이지 않는가

### 확장이 끝났을 때 보내는 확인 프롬프트

```text
Summarize the wiki enricher extension briefly.
Show me:
- where the wiki enrichment step runs
- what learning_context contains
- which wiki pages were used for one example
- how the UI shows the learning info separately from the solution
```

## 9. 마지막 다듬기 프롬프트

기본 기능이 안정적으로 돌아간 뒤에만 이 단계를 넣는 편이 좋다.

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

## 10. 최종 확인 체크리스트

- 정상 텍스트 문제 1개
- 조금 까다로운 텍스트 문제 1개
- 이미지 문제 1개
- 실패 또는 fallback 사례 1개
- 선택 확장을 붙였다면 learning info 사례 1개

확인할 화면 상태:

- analyzing
- solving
- reviewing
- approved result
- fallback or failure

선택 확장을 붙였다면 추가로:

- learning info card
- tentative curriculum mapping state

## 11. 막혔을 때 쓰는 프롬프트

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

### wiki enrichment가 정답 풀이와 섞일 때

```text
Separate the wiki enrichment step from the math solving step.
The app must solve and review first.
Only after approved review, build a separate learning_context from the wiki.
Do not use wiki content to guess the answer.
```

## 12. Gemini CLI에서 추가로 유용한 점

- `Antigravity`의 skill 이름을 직접 쓸 필요는 없다.
- 대신 `AGENTS.md`와 현재 phase를 프롬프트에서 명확히 고정하는 것이 중요하다.
- 긴 세션이 되면 현재 상태를 주기적으로 요약하게 해서 컨텍스트를 정리하는 편이 좋다.

예:

```text
Summarize the current state of this project.
Show:
- completed phase
- remaining work
- run commands
- current blockers
```

## 13. 가장 짧은 버전

1. 첫 시작

```text
Read README.md, AGENTS.md, wiki/README.md, and wiki/index.md first.
Use adk-python and start with Phase 1 only.
```

2. Phase 2 전환

```text
Phase 1 works now.
Move to Phase 2 and refactor it into interpretation, solving, and review subagents.
Keep it terminal-runnable.
```

3. Phase 3 전환

```text
Phase 2 works now.
Move to Phase 3 and build a backend plus UI prototype with text input, image upload, status states, result card, and failure card.
```

4. 선택 확장

```text
Keep the approved solution flow.
Add a wiki enricher step after review approval and show the result in a separate learning info card.
```
