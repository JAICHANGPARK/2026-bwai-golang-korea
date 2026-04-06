# 진행자 스크립트

이 문서는 1시간 핸즈온을 짧고 안정적으로 운영하기 위한 진행자용 한 장 요약이다.

## 0~5분 오프닝

말할 포인트:
- 오늘은 완성 앱보다 `좋은 시스템 설계`를 만드는 시간이란 점
- Solver와 Verifier를 분리하는 것이 핵심이라는 점
- RAG 없이도 멀티 에이전트 경험을 만들 수 있다는 점

확인할 것:
- 참석자가 AI Studio를 열었는지
- [participant-one-pager.md](./participant-one-pager.md)와 [design-worksheet.md](../templates/design-worksheet.md)를 같이 열었는지

## 5~22분 설계 구간

순서:
- [01-user-flow.md](../../prompt-pack/01-user-flow.md)
- [02-agent-design.md](../../prompt-pack/02-agent-design.md)
- [03-system-prompts.md](../../prompt-pack/03-system-prompts.md)

말할 포인트:
- Solver는 푸는 역할
- Verifier는 승인하는 역할
- 두 역할이 섞이면 품질이 흔들린다

## 22~36분 UI와 스키마 구간

순서:
- [03a-ui-screen-design.md](../../prompt-pack/03a-ui-screen-design.md)
- [04-output-schema.md](../../prompt-pack/04-output-schema.md)
- [05-failure-handling.md](../../prompt-pack/05-failure-handling.md)

말할 포인트:
- 최종 설명 카드는 `approved` 이후에만 보인다
- `revised`는 최종 답이 아니라 `재검토 중` 상태다
- raw JSON은 최종 설명으로 직접 노출하지 않는다

## 36~48분 통합 명세 구간

순서:
- [06-integration-spec.md](../../prompt-pack/06-integration-spec.md)

말할 포인트:
- 좋은 명세는 Build와 Antigravity에서 재사용된다
- streaming과 Markdown 요구를 이 단계에서 꼭 넣어야 한다

## 48~58분 앱 생성 구간

순서:
- [07-build-and-antigravity.md](../../prompt-pack/07-build-and-antigravity.md)
- [07a-streaming-markdown.md](../../prompt-pack/07a-streaming-markdown.md)

말할 포인트:
- Build는 빠른 프로토타입
- Antigravity는 Rules/Workflows 감각
- 둘 다 같은 명세에서 출발한다

## 58~60분 마무리

말할 포인트:
- 오늘은 최소형 멀티 에이전트를 설계했다는 점
- 다음 단계에서 교육과정 연결, RAG, 개인화가 붙을 수 있다는 점

## 막힐 때 바로 쓰는 문서

- fallback prompt: [fallback-prompts.md](../templates/fallback-prompts.md)
- 진행자 체크: [hands-on-facilitator-checkpoints.md](../../hands-on-facilitator-checkpoints.md)
- 샘플 문제: [problem-01-linear-equation.svg](../sample-problems/problem-01-linear-equation.svg)
