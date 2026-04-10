---
title: 분산
type: synthesis
status: active
updated: 2026-04-10
card_role: concept
source_docs:
  - docs/math-concept-encyclopedia/korea.md
  - docs/math-concept-encyclopedia/china.md
  - docs/math-curriculum-research/korea.md
  - docs/math-curriculum-research/comparison.md
  - docs/math-concept-encyclopedia/multilingual-glossary.md
  - docs/math-concept-encyclopedia/comparative-problem-book.md
tags:
  - concept
  - statistics
  - variance
---

# 분산

## Summary

분산은 자료나 확률변수가 평균 주변에 얼마나 퍼져 있는지를 제곱 기준으로 측정하는 개념이다. `평균만 같아도 분포는 다를 수 있다`는 사실을 수치로 보여 주는 대표적인 산포 지표다.

## Key Points

- 정의
  - 각 값이 평균에서 얼마나 떨어져 있는지를 제곱해 평균낸 양을 분산이라 한다.
- 핵심 개념
  - 평균
  - 편차
  - 편차제곱
  - 산포
  - 퍼짐
- 대표 수식
  - 자료의 분산:
    $$
    \frac{1}{n}\sum_{i=1}^n (x_i-\bar x)^2
    $$
  - 확률변수의 분산:
    $$
    \operatorname{Var}(X)=E[(X-E(X))^2]
    $$
- 성질 1
  - 분산은 항상 0 이상이다.
- 성질 2
  - 모든 값이 같으면 분산은 0이다.
- 증명 스케치
  - `증명 스케치 (추론)`:
  - 평균에서의 차이인 편차를 그냥 더하면 양수와 음수가 서로 상쇄될 수 있다.
  - 그래서 퍼짐을 측정하려면 편차를 제곱해 모두 양수로 만든 뒤 평균을 낸다.
  - 각 항이 제곱이므로 음수가 될 수 없어 분산도 항상 0 이상이다.

## Deep Dive

분산은 `흩어진 정도`를 재는 도구지만, 단위가 원래 자료의 제곱이 된다는 특징이 있다. 그래서 계산 자체는 분산이 맡고, 해석은 [standard-deviation.md](./standard-deviation.md)이 더 직관적인 경우가 많다.

또한 분산은 극단값에 민감하다. 평균에서 멀리 떨어진 값은 제곱 때문에 영향력이 커지므로, 이상치가 있는 자료를 읽을 때는 분산이 크게 튈 수 있다.

## Worked Examples

- 예제 1: 자료 `1,2,3`의 평균은 `2`다. 편차는 `-1,0,1`이고 편차제곱은 `1,0,1`이므로 분산은
  $$
  \frac{1+0+1}{3}=\frac{2}{3}
  $$
  이다.
- 예제 2: 자료 `2,2,2`의 평균은 `2`, 모든 편차는 `0`이므로 분산은 `0`이다.
- 예제 3: 확률변수 $X$가 `0,1,2`를 각각 확률 `1/4,1/2,1/4`로 가질 때 기대값은 `1`이고,
  $$
  \operatorname{Var}(X)=\frac14(0-1)^2+\frac12(1-1)^2+\frac14(2-1)^2=\frac12
  $$
  이다.

## Common Pitfalls

- 분산을 평균과 같은 종류의 대표값으로 오해하면 안 된다. 분산은 `중심`이 아니라 `퍼짐`이다.
- 편차를 제곱하는 이유를 모르고 공식만 외우면 해석이 약해진다.
- 자료의 단위와 분산의 단위를 같다고 생각하면 안 된다.
- 평균이 같다고 해서 분산도 같을 것이라고 생각하면 안 된다.

## Curriculum Context

- 교육과정 배치
  - 한국 대표 배치에서는 중학교 산포 개념과 고등학교 `확률과 통계`의 확률변수 해석을 잇는 세부 개념이다.
  - 중국 문서군에서는 `데이터 분석`, `확률변수·정규분포·회귀`로 이어지는 축에 분명히 놓인다.
- 국가별 배치 스냅샷
  - 한국: `자료와 가능성` 축에서 분산과 표준편차가 분명히 언급된다.
  - 일본: 현재 문서군에서는 `통계적 추측`과 신뢰 개념의 전단계 산포 감각으로 읽힌다.
  - 중국: 중학 데이터 분석과 고교 확률변수 해석 사이의 브리지다.
  - 미국: Statistics와 AP Statistics에서 variability를 정량화하는 핵심 개념이다.

## Connections

- 선수 개념은 [data-organization.md](./data-organization.md), [random-variable.md](./random-variable.md), [probability-distribution.md](./probability-distribution.md)다.
- 다음 개념으로는 [standard-deviation.md](./standard-deviation.md), [normal-distribution.md](./normal-distribution.md), [statistical-inference.md](./statistical-inference.md)가 이어진다.
- 계통 허브에서는 [statistics-and-probability-strand.md](../statistics-and-probability-strand.md)를 함께 본다.

## Open Questions

- 학교 수준 카드에서 `표본분산`과 `모분산`을 분리할지 기준이 더 필요하다.

## Sources

- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-concept-encyclopedia/china.md`
- `docs/math-curriculum-research/korea.md`
- `docs/math-curriculum-research/comparison.md`
- `docs/math-concept-encyclopedia/multilingual-glossary.md`
- `docs/math-concept-encyclopedia/comparative-problem-book.md`
