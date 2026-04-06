# 04. 구조화 출력 설계 프롬프트

## 목적

두 에이전트가 서로 주고받을 JSON 구조와, 최종 사용자용 설명을 만들기 위한 입력 구조를 설계한다.

## 언제 사용하나

시스템 프롬프트 초안을 만든 뒤 사용한다.

## 프롬프트

```text
Please design a JSON schema for:
1. Solver/Explainer Agent output
2. Expert Verifier Agent output
3. FinalExplanationInput prepared by the app after verification approval

The schema should be simple enough for a workshop demo and should include:
- answer
- steps
- key concepts
- confidence
- verifier verdict
- comments

Important rules:
- the Solver and Verifier return JSON only
- if the verifier verdict is approved, include an approved solution object
- if the verifier verdict is revised, do not treat the corrected content as final user output
- if the verifier verdict is revised, return revision requests and a next_action
- the final user-facing explanation should be generated only from FinalExplanationInput
- the streamed Markdown response must never be raw Solver or Verifier JSON

Return the schema examples in JSON format.
Make them practical for UI rendering.
```

## 기대 산출물

- `SolverResult` JSON 예시
- `VerificationResult` JSON 예시
- `FinalExplanationInput` JSON 예시
