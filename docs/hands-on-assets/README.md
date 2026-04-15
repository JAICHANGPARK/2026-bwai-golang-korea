# 공유용 수학 해설 자산 패키지

이 폴더는 공유 workspace에 그대로 포함해도 되는 공개용 보조 자산 모음이다.

## 구성

- `sample-problems/`
  - 업로드 테스트에 쓸 샘플 문제 이미지
- `handouts/`
  - 빠르게 흐름을 잡는 원페이지 요약
  - 단계별 목표를 확인하는 카드형 가이드
- `templates/`
  - 설계 워크시트
  - 통합 명세 템플릿
  - fallback prompt 템플릿

실행 가능한 starter code는 이 폴더가 아니라 [`starter-kits/README.md`](../../starter-kits/README.md)에 둔다.

## 사용 순서

1. `handouts/math-agent-one-pager.md`로 전체 흐름을 확인한다.
2. `templates/design-worksheet.md`로 설계 결과를 정리한다.
3. `docs/prompt-pack/` 문서를 순서대로 AI Studio에 넣는다.
4. 필요하면 `handouts/math-agent-step-cards.md`를 옆에 띄워 단계별 목표를 확인한다.
5. 샘플 문제 이미지는 `sample-problems/`에서 골라 업로드 테스트를 한다.
6. 생성 결과가 흔들리면 `templates/fallback-prompts.md`의 문구로 범위를 다시 줄인다.
7. runnable scaffold가 필요하면 `starter-kits/`에서 런타임별 뼈대를 고른다.
8. 마지막에는 `templates/integration-spec-template.md`에 최종 명세를 정리한다.

## 포함된 샘플 문제

- `problem-01-linear-equation.svg`
- `problem-02-function-value.svg`
- `problem-03-quadratic-choice.svg`

## 공개 범위 메모

- 이 폴더에는 공유 workspace에 남겨도 되는 자료만 둔다.
- facilitator용 타이밍 메모, 진행 스크립트, 내부 체크리스트는 이 폴더에 두지 않는다.
- 현재 샘플 이미지는 `clean input` 예시다.
- 실제 촬영 사진형 예시는 필요하면 이후에 별도로 추가한다.

## 빠른 링크

- [math-agent-one-pager.md](./handouts/math-agent-one-pager.md)
- [math-agent-step-cards.md](./handouts/math-agent-step-cards.md)
- [design-worksheet.md](./templates/design-worksheet.md)
- [integration-spec-template.md](./templates/integration-spec-template.md)
- [fallback-prompts.md](./templates/fallback-prompts.md)
- [starter-kits/README.md](../../starter-kits/README.md)
