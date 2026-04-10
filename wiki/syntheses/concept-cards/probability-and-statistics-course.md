---
title: 확률과 통계
type: synthesis
status: active
updated: 2026-04-10
card_role: wrapper
source_docs:
  - docs/math-concept-encyclopedia/korea.md
  - docs/math-curriculum-research/korea.md
tags:
  - course
  - korea
  - statistics
---

# 확률과 통계

## Summary

> [!info] Course Wrapper
> 이 카드는 개념 자체를 정의하는 페이지가 아니라, 과목이나 경로 안에서 개념들이 어떻게 묶여 배우는지를 안내하는 허브다.

한국의 `확률과 통계` 과목은 경우의 수 심화, 확률분포, 통계적 추정을 하나의 코스로 묶는다. 세기에서 분포로, 분포에서 추정으로 넘어가는 데이터 해석의 중심 축이다.

## Key Points

- 정의
  - 고등학교 선택 과목으로, 조합론과 확률, 분포, 추정을 하나의 연속된 학습 경로로 묶는다.
- Course Map
  - 1단계: [counting-principles.md](./counting-principles.md)와 [conditional-probability.md](./conditional-probability.md)
  - 2단계: [random-variable.md](./random-variable.md)와 [probability-distribution.md](./probability-distribution.md)
  - 3단계: 기대값, [variance.md](./variance.md), [standard-deviation.md](./standard-deviation.md)
  - 4단계: [sampling.md](./sampling.md)과 [statistical-inference.md](./statistical-inference.md), 이후 심화 연결로 [confidence-interval.md](./confidence-interval.md), [hypothesis-testing.md](./hypothesis-testing.md)
- 대표 개념
  - 독립사건
  - [conditional-probability.md](./conditional-probability.md)
  - [random-variable.md](./random-variable.md)
  - [probability-distribution.md](./probability-distribution.md)
  - [variance.md](./variance.md)
  - [standard-deviation.md](./standard-deviation.md)
  - [sampling.md](./sampling.md)
  - [statistical-inference.md](./statistical-inference.md)
  - 심화 연결: [confidence-interval.md](./confidence-interval.md), [hypothesis-testing.md](./hypothesis-testing.md)
- Course Logic
  - 이 과목은 `세기 -> 숫자화 -> 분포화 -> 추정`의 순서로 진행된다.
  - 먼저 경우의 수와 조건부확률로 사건을 세고, 다음에 확률변수로 값을 부여한 뒤, 분포를 만들고, 마지막에 표본을 통해 모집단을 추정한다.
- 대표 수식
  - $P(A\mid B)=\frac{P(A\cap B)}{P(B)}$
  - $E(X)=\sum xP(X=x)$
  - $\operatorname{Var}(X)=E[(X-E(X))^2]$
- 교육과정 배치
  - 한국 대표 배치에서는 `고2` 선택 단계의 핵심 데이터 과목이다.
- 비교 메모
  - 미국 `Statistics and Data Analysis`보다 분포와 수식 비중이 더 높고, `실용 통계`보다 이론성이 강하다.
  - 일본과 중국의 확률·통계 과목을 읽을 때도, 이 과목은 `정리된 학교 문법`에 가장 가깝다.
- Deep Dive
  - 이 과목의 진짜 중심은 `분포`다.
  - 조건부확률은 사건의 구조를 바꾸고, 확률변수는 사건을 숫자로 바꾸고, 확률분포는 그 숫자들의 가능성을 정리한다.
  - 그래서 `확률과 통계`는 사실상 `확률변수와 분포를 이해하기 위한 프레임`이라고 보는 편이 정확하다.
- Worked Examples
  - 예제 1: 주사위를 두 번 던져 합을 다루는 문제에서
    - 경우의 수 단계에서는 가능한 순서쌍을 센다.
    - 확률변수 단계에서는 합을 $X$로 둔다.
    - 분포 단계에서는 $P(X=2),\dots,P(X=12)$를 구한다.
    - 기대값 단계에서는 $E(X)=7$을 얻는다.
  - 예제 2: 빨간 공과 파란 공이 섞인 상자에서 비복원 추출을 할 때
    - 조건부확률로 첫 번째 결과가 둘째 결과에 미치는 영향을 본다.
    - 그다음 확률변수 `빨간 공 개수`를 정의한다.
    - 분포를 만들고 나면 평균적으로 몇 개의 빨간 공이 나오는지를 계산할 수 있다.
- Common Pitfalls
  - 경우의 수와 확률변수를 한 단계로 섞어버린다.
  - 조건부확률을 건너뛰고 바로 분포로 가려 한다.
  - 기대값을 `가장 자주 나오는 값`과 혼동한다.
  - 표본과 모집단을 구분하지 않고 추정 공식을 쓴다.

## Connections

- 핵심 카드로는 [counting-principles.md](./counting-principles.md), [conditional-probability.md](./conditional-probability.md), [random-variable.md](./random-variable.md), [probability-distribution.md](./probability-distribution.md), [variance.md](./variance.md), [standard-deviation.md](./standard-deviation.md), [confidence-interval.md](./confidence-interval.md), [hypothesis-testing.md](./hypothesis-testing.md), [statistical-inference.md](./statistical-inference.md)가 있다.
- 다음 과목 카드로는 [practical-statistics.md](./practical-statistics.md), [statistics-and-data-analysis.md](./statistics-and-data-analysis.md)와 연결된다.
- 학년 허브로는 [../high-2-hub.md](../high-2-hub.md), [../high-3-hub.md](../high-3-hub.md)를 함께 본다.
- 계통 허브로는 [../statistics-and-probability-strand.md](../statistics-and-probability-strand.md)를 함께 본다.

## Open Questions

- `확률과 통계` 카드 안에 `정규분포`를 포함한 심화 분포를 별도 섹션으로 둘지, 현재처럼 개념 카드 쪽으로 분리 유지할지 기준이 필요하다.

## Sources

- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-curriculum-research/korea.md`
