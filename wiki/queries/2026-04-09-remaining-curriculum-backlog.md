---
title: Remaining Curriculum Backlog Refresh
type: query-note
status: active
updated: 2026-04-09
source_docs:
  - docs/math-curriculum-research/comparison.md
  - docs/math-curriculum-research/korea.md
  - docs/math-curriculum-research/japan.md
  - docs/math-curriculum-research/china.md
  - docs/math-curriculum-research/us.md
  - docs/math-concept-encyclopedia/korea.md
  - docs/math-concept-encyclopedia/japan.md
  - docs/math-concept-encyclopedia/china.md
  - docs/math-concept-encyclopedia/us.md
tags:
  - query
  - backlog
  - curriculum
---

# Remaining Curriculum Backlog Refresh

## Summary

예전에 남아 있던 `대표 개념 카드` 중심 backlog는 지금 기준으로 완료 상태다. `복소수`, `복소수평면`, `수학적 모델링·수학 탐구`, `Polar Ideas`, `Sequences and Series`, `AP Statistics`까지 persistent layer에 올라오면서, 현재 남은 일은 개념 자체보다 `나라별 학년·과목 wrapper`, `지속형 curriculum hub`, `4개국 비교 synthesis`, 그리고 더 세밀한 `하위 개념 카드`를 보강하는 쪽으로 이동했다.

## Key Points

- 이전 대표 backlog의 상태
  - `소인수분해`, `정수와 유리수`, `문자와 식`, `식의 계산`, `다항식과 인수분해`, `제곱근과 실수`, `일차부등식`
  - `기본 도형과 작도`, `합동`, `닮음`, `원`, `피타고라스 정리`, `평행선·다각형·합동의 논증`
  - `반비례`, `지수함수와 로그함수`, `수열`, `함수 일반론`, `극한`, `미분`, `적분`
  - `집합과 명제`, `이차곡선`, `벡터`, `공간좌표`
  - `자료 정리`, `상관과 산점도`, `정규분포`, `표본조사`, `실용 통계`
  - `경제 수학`, `인공지능 수학`, `직무 수학`, `수학과 문화`, `Integrated Mathematics`
  - 위 목록은 현재 persistent concept layer에 모두 존재한다.
  - 추가로 `복소수`, `복소수평면`, `수학적 모델링·수학 탐구`, `Polar Ideas`, `Sequences and Series`, `AP Statistics`도 이번 배치로 반영되었다.
- 지금 실제로 남아 있는 1순위
  - `한국 학년 허브`는 아직 query sample 비중이 크다.
  - `일본 wrapper`: `중1`, `중2`, `중3`, `수학I`, `수학A`, `수학II`, `수학B`, `수학III`, `수학C`
  - `중국 wrapper`: `7학년 1학기`, `7학년 2학기`, `8학년 1학기`, `8학년 2학기`, `9학년 1학기`, `9학년 2학기`, `필수 과정`, `선택필수 과정`, `선택 과정의 개념군`
  - `미국 wrapper`: `Grade 6`, `Grade 7`, `Grade 8`
  - 이 항목들은 원천 문서에서 상위 heading으로 직접 보이지만, 아직 persistent synthesis/card로는 대부분 승격되지 않았다.
- 지금 실제로 남아 있는 2순위
  - `나라별 curriculum hub`: 한국, 일본, 중국, 미국별 허브
  - `course-track hub`: 국가 공통과정형, 과목선택형, 경로형을 비교하는 허브
  - `cross-country comparison synthesis`: 함수 위계, 기하 논증, 통계·데이터, 입시 평가 축 비교
  - `assessment synthesis`: 수능, 공통테스트, 가오카오, SAT/미국식 holistic admissions 차이를 수학 학습 경로와 연결하는 허브
- 지금 실제로 남아 있는 3순위
  - `복소수의 연산`
  - `행렬의 연산`
  - `연속함수의 성질`
  - `라디안`
  - `삼각함수 항등식`
  - `AP Calculus`
  - `IB Mathematics`
  - 위 항목들은 기존 원천 문서와 이미 생성된 카드의 open question에서 직접 이어지는 `세부 분화 backlog`다.

### Recommended Next Expansion Order

- 1차
  - `일본 wrapper 카드군`
  - `중국 wrapper 카드군`
  - `미국 Grade 6~8 wrapper 카드군`
- 2차
  - `한국/일본/중국/미국 curriculum hub`
  - `cross-country comparison synthesis`
  - `assessment synthesis`
- 3차
  - `복소수의 연산`
  - `행렬의 연산`
  - `연속함수의 성질`
  - `라디안`
  - `삼각함수 항등식`
  - `AP Calculus`
  - `IB Mathematics`

## Connections

- 현재 전체 반영 상태 감사는 [2026-04-06-research-folder-coverage-audit.md](./2026-04-06-research-folder-coverage-audit.md)에 있다.
- 개념 카드 구조는 [../syntheses/concept-card-model.md](../syntheses/concept-card-model.md)에 있다.
- 한국 샘플 그래프는 [2026-04-06-math-curriculum-graph.md](./2026-04-06-math-curriculum-graph.md)와 하위 샘플 노트들을 본다.
- 나라별 직접 요약은 [../source-notes/math-curriculum-research-japan.md](../source-notes/math-curriculum-research-japan.md), [../source-notes/math-curriculum-research-china.md](../source-notes/math-curriculum-research-china.md), [../source-notes/math-curriculum-research-us.md](../source-notes/math-curriculum-research-us.md), [../source-notes/math-concept-encyclopedia-japan.md](../source-notes/math-concept-encyclopedia-japan.md), [../source-notes/math-concept-encyclopedia-china.md](../source-notes/math-concept-encyclopedia-china.md), [../source-notes/math-concept-encyclopedia-us.md](../source-notes/math-concept-encyclopedia-us.md)에 있다.

## Open Questions

- `학년 허브`를 query sample로 유지할지, persistent synthesis로 승격할지 기준이 필요하다.
- `과목 wrapper 카드`와 `나라별 curriculum hub`를 같은 레이어에 둘지 한 단계 더 나눌지 정리가 필요하다.
- `라디안`, `복소수의 연산`, `연속함수의 성질`처럼 상위 카드 안에 아직 묶여 있는 세부 개념을 언제 독립 카드로 뺄지 기준을 더 선명히 할 필요가 있다.

## Sources

- `docs/math-curriculum-research/comparison.md`
- `docs/math-curriculum-research/korea.md`
- `docs/math-curriculum-research/japan.md`
- `docs/math-curriculum-research/china.md`
- `docs/math-curriculum-research/us.md`
- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-concept-encyclopedia/japan.md`
- `docs/math-concept-encyclopedia/china.md`
- `docs/math-concept-encyclopedia/us.md`
