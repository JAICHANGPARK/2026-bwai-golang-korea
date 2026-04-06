# 핸즈온 자산 패키지

이 폴더는 1시간 수학 해설 시스템 핸즈온에서 바로 쓸 수 있는 보조 자산 모음이다.

## 구성

- `sample-problems/`
  - AI Studio Build 또는 Antigravity에서 업로드 테스트에 쓸 샘플 문제 이미지
- `handouts/`
  - 참석자 안내서
  - 진행자 스크립트
  - 실습 순서 카드
- `templates/`
  - 설계 워크시트
  - 통합 명세 템플릿
  - fallback prompt 템플릿

## 사용 순서

1. `handouts/participant-one-pager.md`로 오늘 실습 흐름을 확인한다.
2. `templates/design-worksheet.md`로 설계 결과를 정리한다.
3. `docs/prompt-pack/` 문서를 순서대로 AI Studio에 넣는다.
4. 필요하면 `handouts/session-step-cards.md`를 옆에 띄워 단계별 목표를 확인한다.
5. 샘플 문제 이미지는 `sample-problems/`에서 골라 업로드 테스트를 한다.
6. 생성 결과가 흔들리면 `templates/fallback-prompts.md`의 문구로 범위를 다시 줄인다.
7. 마지막에는 `templates/integration-spec-template.md`에 최종 명세를 정리한다.

## 포함된 샘플 문제

- `problem-01-linear-equation.svg`
- `problem-02-function-value.svg`
- `problem-03-quadratic-choice.svg`

## 메모

- 현재 샘플 이미지는 워크숍용 `clean input` 예시다.
- 실제 촬영 사진형 예시는 필요하면 이후에 별도로 추가한다.

## 빠른 링크

- [participant-one-pager.md](./handouts/participant-one-pager.md)
- [facilitator-one-pager.md](./handouts/facilitator-one-pager.md)
- [session-step-cards.md](./handouts/session-step-cards.md)
- [design-worksheet.md](./templates/design-worksheet.md)
- [integration-spec-template.md](./templates/integration-spec-template.md)
- [fallback-prompts.md](./templates/fallback-prompts.md)
