# 07. Build 및 Antigravity 실행 프롬프트

## 목적

통합 명세 프롬프트를 각각 `AI Studio Build`와 `Antigravity`에 넣어 앱을 생성해본다.

## 언제 사용하나

통합 구현 명세 프롬프트가 준비된 뒤 사용한다.

## 7-1. AI Studio Build용 프롬프트

아래 프롬프트에서 `{{INTEGRATION_SPEC}}` 부분에 앞 단계 결과를 넣는다.

```text
Use the following implementation specification to build a simple workshop app:

{{INTEGRATION_SPEC}}

Requirements:
- generate a clean Korean web UI
- keep the project simple and beginner-friendly
- show separate sections for solver result, verifier result, and final approved response
- do not show a final explanation until the verifier verdict is approved
- if verdict is revised, show verifier comments and a retry or review-in-progress state instead of a final answer
- support streaming for the final approved explanation
- render the final approved explanation with a Markdown view
- show clear states for loading, streaming, verifying, and clarification needed
- use TypeScript
- do not add any features outside the specification
```

## 7-2. Antigravity용 프롬프트

아래 프롬프트에서 `{{INTEGRATION_SPEC}}` 부분에 앞 단계 결과를 넣는다.

```text
Create an app based on the following implementation specification:

{{INTEGRATION_SPEC}}

Also apply these development constraints:
- keep the workflow simple for a 1-hour workshop
- maintain two agents only
- structured JSON outputs are required
- no RAG, database, or authentication
- UI text should be in Korean
- stream only the final approved explanation, not raw agent JSON
- if the verifier returns revised, do one bounded retry or keep the app in a review state
- final explanation should stream progressively
- final explanation should be shown in a Markdown-rendered card
```

## 7-3. Antigravity Rules 보강 프롬프트

```text
Please define project rules for this workshop app.

The rules should enforce:
- two-agent architecture only
- no RAG
- no database
- structured JSON outputs
- Korean UI
- beginner-friendly code
- streamed final explanation UX
- Markdown rendering for the final answer card
```

## 7-4. Antigravity Workflow 보강 프롬프트

```text
Please define reusable workflow prompts for this project:
- scaffold app
- add solver
- add verifier
- polish UI
- add streaming and markdown rendering

Keep each workflow short and reusable.
```

## 기대 산출물

- AI Studio Build 버전 앱
- Antigravity 버전 앱
- 선택적으로 Rules / Workflows 초안
