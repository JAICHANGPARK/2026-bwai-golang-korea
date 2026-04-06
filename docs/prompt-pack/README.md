# 수학 해설 시스템 핸즈온 프롬프트 팩

이 폴더는 `AI Studio -> 명세 생성 -> AI Studio Build -> Antigravity` 흐름으로 진행하는 핸즈온 세션에서 참석자가 바로 복사해 사용할 수 있는 단계별 프롬프트 모음이다.

## 권장 순서

1. [01-user-flow.md](./01-user-flow.md)
2. [02-agent-design.md](./02-agent-design.md)
3. [03-system-prompts.md](./03-system-prompts.md)
4. [03a-ui-screen-design.md](./03a-ui-screen-design.md)
5. [04-output-schema.md](./04-output-schema.md)
6. [05-failure-handling.md](./05-failure-handling.md)
7. [06-integration-spec.md](./06-integration-spec.md)
8. [07-build-and-antigravity.md](./07-build-and-antigravity.md)
9. [07a-streaming-markdown.md](./07a-streaming-markdown.md)

## 사용 원칙

- 메인 모델은 `gemini-3-flash-preview`
- `RAG`는 사용하지 않음
- 두 에이전트만 사용
  - `Solver/Explainer`
  - `Expert Verifier`
- 먼저 설계와 명세를 만들고, 그 결과를 Build와 Antigravity에 재사용함

## 문서 성격

- 이 폴더의 문서는 `참석자용`이다.
- 진행자 운영 메모와 체크 포인트는 [hands-on-facilitator-checkpoints.md](../hands-on-facilitator-checkpoints.md)에서 따로 본다.

## 적용 범위

- 이 프롬프트 팩과 [hands-on-math-agent-session-guide.md](../hands-on-math-agent-session-guide.md)는 `1시간 핸즈온 프로파일` 기준 문서다.
- 같은 워크스페이스의 [math-agent-system-spec.md](../math-agent-system-spec.md), [math-agent-knowledge-plan.md](../math-agent-knowledge-plan.md), [internal-knowledge-base-build.md](../internal-knowledge-base-build.md)는 `제품 확장/다음 단계` 참고 문서다.
- 따라서 핸즈온에서는 `RAG 없음`, `두 에이전트만 사용`, `최종 설명은 검증 승인 후에만 표시` 원칙을 우선한다.

## 권장 산출물

- 사용자 흐름
- 에이전트 역할 정의
- 시스템 프롬프트
- 화면 설계 방향
- JSON 출력 스키마
- 실패 처리 규칙
- 통합 구현 명세 프롬프트
- 스트리밍 및 Markdown 렌더링 요구사항
