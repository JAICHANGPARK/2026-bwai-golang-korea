# 수학 해설 에이전트 단계 카드

한 단계씩 빠르게 따라가기 위한 카드형 요약이다.

## 카드 1. 사용자 흐름

- 문서: [01-user-flow.md](../../prompt-pack/01-user-flow.md)
- 목표: 입력 -> 풀이 -> 검증 -> 최종 응답 흐름 정리
- 저장할 것: 사용자 흐름 초안

## 카드 2. 에이전트 역할

- 문서: [02-agent-design.md](../../prompt-pack/02-agent-design.md)
- 목표: Solver와 Verifier 책임 분리
- 저장할 것: 각 에이전트 역할 정의

## 카드 3. 시스템 프롬프트

- 문서: [03-system-prompts.md](../../prompt-pack/03-system-prompts.md)
- 목표: 두 에이전트의 시스템 프롬프트 초안 만들기
- 저장할 것: Solver prompt, Verifier prompt

## 카드 4. 화면 설계

- 문서: [03a-ui-screen-design.md](../../prompt-pack/03a-ui-screen-design.md)
- 목표: 업로드 영역, 결과 카드, 상태 영역 설계
- 저장할 것: 단일 페이지 UI 초안

## 카드 5. 출력 구조와 실패 처리

- 문서: [04-output-schema.md](../../prompt-pack/04-output-schema.md), [05-failure-handling.md](../../prompt-pack/05-failure-handling.md)
- 목표: JSON 계약과 예외 흐름 정의
- 저장할 것: SolverResult, VerificationResult, FinalExplanationInput

## 카드 6. 통합 명세

- 문서: [06-integration-spec.md](../../prompt-pack/06-integration-spec.md)
- 목표: Build와 Antigravity에 공용으로 넣을 명세 만들기
- 저장할 것: 통합 구현 명세 프롬프트

## 카드 7. 앱 생성

- 문서: [07-build-and-antigravity.md](../../prompt-pack/07-build-and-antigravity.md)
- 목표: 같은 명세로 Build와 Antigravity 앱 생성
- 저장할 것: 생성된 앱 초안

## 카드 8. Streaming과 Markdown

- 문서: [07a-streaming-markdown.md](../../prompt-pack/07a-streaming-markdown.md)
- 목표: 최종 승인 설명만 streaming + Markdown으로 보이게 만들기
- 저장할 것: UX 보강 결과

## 카드 9. 막힐 때

- 워크시트: [design-worksheet.md](../templates/design-worksheet.md)
- fallback prompt: [fallback-prompts.md](../templates/fallback-prompts.md)
- 샘플 문제: [problem-01-linear-equation.svg](../sample-problems/problem-01-linear-equation.svg), [problem-02-function-value.svg](../sample-problems/problem-02-function-value.svg), [problem-03-quadratic-choice.svg](../sample-problems/problem-03-quadratic-choice.svg)
