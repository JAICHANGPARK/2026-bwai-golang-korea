# 03. 시스템 프롬프트 설계 프롬프트

## 목적

두 에이전트의 시스템 프롬프트를 초안 수준이 아니라 실제로 붙여 넣을 수 있는 형태로 만든다.

## 언제 사용하나

에이전트 역할을 정리한 직후 사용한다.

## 3-1. Solver 시스템 프롬프트 생성 프롬프트

```text
Please write a high-quality system prompt for a Solver/Explainer Agent for Korean middle and high school math problems.

The agent should:
- read a math problem image carefully
- solve it step by step
- explain it for students
- produce structured output

Important rules:
- do not guess unreadable text
- do not skip logical steps
- keep explanations clear and educational
- do not act like the final verifier

Write the system prompt in English.
```

## 3-2. Verifier 시스템 프롬프트 생성 프롬프트

```text
Please write a high-quality system prompt for an Expert Verifier Agent for Korean math problem solving.

The verifier should:
- inspect the solver output
- check correctness
- detect missing conditions
- detect logical inconsistencies
- decide whether the result is approved, revised, or needs clarification

Important rules:
- be strict
- do not approve just because the answer sounds fluent
- focus on correctness before style

Write the system prompt in English.
```

## 3-3. Wiki Enricher 시스템 프롬프트 생성 프롬프트 `선택`

기본 two-agent 구조가 안정적으로 동작한 뒤, `난이도`, `핵심 개념`, `관련 교육과정`, `연관 학습 주제`를 `LLM Wiki`에서 보강하고 싶을 때 사용한다.

```text
Please write a high-quality system prompt for a Wiki Knowledge Enricher agent for a Korean math explanation system built with ADK.

The enricher should:
- run only after the solver result is approved by the verifier
- read the minimum relevant wiki pages
- build a structured learning_context
- identify:
  - difficulty
  - core concepts
  - related curriculum placement
  - follow-up topics
  - wiki basis pages

Important rules:
- do not use wiki content to guess the math answer
- do not run before verification approval
- if wiki evidence is weak, mark the mapping as tentative
- keep the result structured and compact

Write the system prompt in English.
```

## 기대 산출물

- Solver 시스템 프롬프트 1개
- Verifier 시스템 프롬프트 1개
- 선택 시 Wiki Enricher 시스템 프롬프트 1개
