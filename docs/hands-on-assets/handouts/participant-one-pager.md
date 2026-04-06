# 참석자 안내서

이 세션의 목표는 `수학 해설 시스템을 직접 설계하고`, 그 설계를 `AI Studio Build`와 `Antigravity`로 앱까지 연결해보는 것이다.

## 오늘 만드는 것

- Solver/Explainer 에이전트
- Expert Verifier 에이전트
- 화면 구조
- JSON 출력 구조
- 최종 설명 streaming + Markdown UX
- 앱 생성용 통합 명세

## 오늘 규칙

- `RAG`는 사용하지 않는다.
- 에이전트는 `2개만` 사용한다.
- 최종 설명은 `검증 승인 후에만` 보여준다.
- raw JSON은 최종 답으로 바로 보여주지 않는다.

## 실습 순서

1. 사용자 흐름 설계: [01-user-flow.md](../../prompt-pack/01-user-flow.md)
2. 에이전트 역할 분리: [02-agent-design.md](../../prompt-pack/02-agent-design.md)
3. 시스템 프롬프트 작성: [03-system-prompts.md](../../prompt-pack/03-system-prompts.md)
4. 화면 설계: [03a-ui-screen-design.md](../../prompt-pack/03a-ui-screen-design.md)
5. 출력 스키마와 실패 처리: [04-output-schema.md](../../prompt-pack/04-output-schema.md), [05-failure-handling.md](../../prompt-pack/05-failure-handling.md)
6. 통합 명세 만들기: [06-integration-spec.md](../../prompt-pack/06-integration-spec.md)
7. 앱 생성과 보강: [07-build-and-antigravity.md](../../prompt-pack/07-build-and-antigravity.md), [07a-streaming-markdown.md](../../prompt-pack/07a-streaming-markdown.md)

## 같이 열어두면 좋은 문서

- 워크시트: [design-worksheet.md](../templates/design-worksheet.md)
- 통합 명세 템플릿: [integration-spec-template.md](../templates/integration-spec-template.md)
- fallback prompt 모음: [fallback-prompts.md](../templates/fallback-prompts.md)
- 샘플 문제 이미지: [problem-01-linear-equation.svg](../sample-problems/problem-01-linear-equation.svg), [problem-02-function-value.svg](../sample-problems/problem-02-function-value.svg), [problem-03-quadratic-choice.svg](../sample-problems/problem-03-quadratic-choice.svg)

## 최종 체크

- 사용자 흐름이 정리되었는가
- Solver와 Verifier 역할이 분리되었는가
- 구조화 출력이 정의되었는가
- 최종 설명이 `approved` 이후에만 보이도록 설계되었는가
- streaming과 Markdown 요구가 명세에 포함되었는가

## 막혔을 때

- 먼저 범위를 줄인다.
- UI보다 구조를 먼저 고정한다.
- `fallback-prompts.md`의 문구를 그대로 사용한다.
