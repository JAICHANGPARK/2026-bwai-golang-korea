# 06. 통합 구현 명세 프롬프트

## 목적

앞 단계에서 만든 설계 내용을 `AI Studio Build`와 `Antigravity`에 재사용할 수 있는 하나의 통합 명세 프롬프트로 합친다.

## 언제 사용하나

사용자 흐름, 에이전트 역할, 시스템 프롬프트, JSON 구조, 실패 처리까지 정리한 뒤 사용한다.

## 프롬프트

```text
Based on the design above, create a single implementation specification prompt that can be reused in app builders.

The prompt should include:
- product goal
- target users
- user flow
- two-agent architecture
- Solver role
- Verifier role
- optional post-verification wiki enrichment step for learning metadata
- if the wiki enrichment step is used, define a local wiki lookup helper or context builder instead of relying on prompt-only file reading
- structured JSON outputs
- a final explanation generation step after verification approval
- if the wiki enrichment step is used, define a learning_context object for:
  - difficulty
  - concepts
  - curriculum mapping
  - related topics
- a rule that the final explanation is streamed only after approved verification
- a rule that revised means retry or hold, not immediate final display
- simple Korean UI requirements
- separate cards for solver, verifier, and final answer
- if the wiki enrichment step is used, add a learning info card
- streaming response requirement for the final user-facing explanation
- Markdown rendering requirement for the final approved explanation
- status handling for loading, streaming, verifying, completed, clarification needed
- constraints: no RAG, no database, no auth

Write the result in English so it can be pasted into AI Studio Build or Antigravity.
Keep it concise but implementation-ready.
```

## 기대 산출물

- 하나의 통합 구현 명세 프롬프트
- 선택 시 `wiki enrichment -> learning_context -> final explanation` 흐름까지 포함된 통합 명세
