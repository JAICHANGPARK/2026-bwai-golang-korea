---
title: 표준편차
type: synthesis
status: active
updated: 2026-04-10
card_role: concept
source_docs:
  - docs/math-concept-encyclopedia/korea.md
  - docs/math-curriculum-research/korea.md
  - docs/math-concept-encyclopedia/multilingual-glossary.md
tags:
  - concept
  - statistics
  - spread
---

# 표준편차

## Summary

표준편차는 분산에 제곱근을 취해 자료의 퍼짐을 원래 단위로 되돌린 값이다. 분산보다 해석이 직관적이어서, `평균에서 대략 얼마나 떨어져 있는가`를 읽는 데 자주 쓰인다.

## Key Points

- 정의
  - 분산의 제곱근을 표준편차라 한다.
- 핵심 개념
  - 분산
  - 제곱근
  - 원래 단위
  - 퍼짐의 크기
- 대표 수식
  - 자료의 표준편차:
    $$
    \sigma=\sqrt{\frac{1}{n}\sum_{i=1}^n (x_i-\bar x)^2}
    $$
  - 확률변수의 표준편차:
    $$
    \sigma=\sqrt{\operatorname{Var}(X)}
    $$
- 성질 1
  - 표준편차는 항상 0 이상이다.
- 성질 2
  - 분산이 0이면 표준편차도 0이다.
- 증명 스케치
  - `증명 스케치 (추론)`:
  - 분산은 퍼짐을 잘 측정하지만 단위가 제곱되어 해석이 불편하다.
  - 그래서 제곱근을 취해 다시 원래 단위로 돌리면, 자료와 같은 단위에서 퍼짐을 읽을 수 있다.

## Deep Dive

분산이 계산의 중심이라면, 표준편차는 해석의 중심이다. 예를 들어 시험 점수의 분산이 `25`라고 하면 직관이 약하지만, 표준편차가 `5점`이라고 하면 평균에서 얼마나 퍼져 있는지 바로 느낌이 온다.

정규분포나 표준점수 해석에서도 표준편차는 핵심 역할을 한다. 평균에서 표준편차 몇 개만큼 떨어졌는지가 상대적 위치를 비교하는 기본 언어가 된다.

## Worked Examples

- 예제 1: 자료 `1,2,3`의 분산이
  $$
  \frac23
  $$
  이므로 표준편차는
  $$
  \sqrt{\frac23}
  $$
  이다.
- 예제 2: 자료 `2,2,2`는 분산이 `0`이므로 표준편차도 `0`이다. 즉 자료가 전혀 퍼져 있지 않다.
- 예제 3: 두 시험의 평균은 같지만 A의 표준편차가 4, B의 표준편차가 10이라면 B가 훨씬 넓게 퍼진 시험으로 해석할 수 있다.

## Common Pitfalls

- 표준편차를 평균과 혼동하면 안 된다.
- 분산과 표준편차를 같은 값으로 기억하면 안 된다.
- 표준편차가 작다고 해서 무조건 좋은 자료라고 단정하면 안 된다. 상황에 따라 다양성이 필요한 경우도 있다.
- 원자료 단위와 표준편차 단위는 같지만, 분산 단위는 다르다는 점을 놓치기 쉽다.

## Curriculum Context

- 교육과정 배치
  - 한국 대표 배치에서는 분산과 함께 산포를 읽는 핵심 세부 개념이다.
  - 이후 [normal-distribution.md](./normal-distribution.md), [statistical-inference.md](./statistical-inference.md) 해석으로 이어진다.
- 국가별 배치 스냅샷
  - 한국: 중등 자료 정리의 후속 산포 개념과 고교 확통 해석을 잇는다.
  - 일본: 통계적 추측과 신뢰 해석의 전단계 variability 감각으로 읽힌다.
  - 중국: 데이터 분석과 정규분포 해석 사이의 브리지로 기능한다.
  - 미국: Statistics와 AP Statistics에서 표준점수, 정규분포 해석과 함께 자주 쓰인다.

## Connections

- 선수 개념은 [variance.md](./variance.md), [data-organization.md](./data-organization.md)다.
- 다음 개념으로는 [normal-distribution.md](./normal-distribution.md), [confidence-interval.md](./confidence-interval.md), [statistical-inference.md](./statistical-inference.md)가 이어진다.
- 계통 허브에서는 [statistics-and-probability-strand.md](../statistics-and-probability-strand.md)를 함께 본다.

## Open Questions

- `표준점수`를 독립 카드로 더 뺄지 기준이 필요하다.

## Sources

- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-curriculum-research/korea.md`
- `docs/math-concept-encyclopedia/multilingual-glossary.md`
