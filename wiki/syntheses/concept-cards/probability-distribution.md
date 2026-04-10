---
title: 확률분포
type: synthesis
status: active
updated: 2026-04-10
card_role: concept
source_docs:
  - docs/math-concept-encyclopedia/korea.md
  - docs/math-concept-encyclopedia/japan.md
  - docs/math-concept-encyclopedia/china.md
  - docs/math-concept-encyclopedia/us.md
  - docs/math-curriculum-research/korea.md
  - docs/math-curriculum-research/japan.md
  - docs/math-curriculum-research/china.md
  - docs/math-curriculum-research/us.md
  - docs/math-concept-encyclopedia/multilingual-glossary.md
  - docs/math-concept-encyclopedia/formula-examples.md
  - docs/math-concept-encyclopedia/unit-practice-book.md
  - docs/math-concept-encyclopedia/comparative-problem-book.md
tags:
  - concept
  - statistics
  - probability
---

# 확률분포

## Summary

확률분포는 확률변수가 어떤 값을 얼마나 자주 가질 수 있는지를 수로 정리한 구조다. 중학교의 확률과 자료 해석이 고등학교 `확률과 통계`에서 확률변수, 기대값, 분산으로 확장될 때 핵심이 되는 개념이다.

## Key Points

- 정의
  - 확률변수 `X`가 취할 수 있는 값들과 그 값에 대응하는 확률 규칙 전체를 확률분포라 한다.
  - 분포는 단순한 표가 아니라 `어떤 값이 가능한가`와 `그 값이 얼마나 가능성이 있는가`를 함께 담는 구조다.
- 핵심 개념
  - 값의 집합
  - 지지도(support)
  - 이산분포
  - 연속분포
  - 확률질량함수/밀도함수
  - 누적분포함수
  - 기대값과 분산
- Deep Dive
  - 이산형 분포에서는 각 값에 확률질량함수 $p(x)=P(X=x)$를 붙이고,
    $$
    p(x)\ge 0,\qquad \sum_x p(x)=1
    $$
    이다.
  - 연속형 분포에서는 밀도함수 $f(x)$를 사용하며,
    $$
    f(x)\ge 0,\qquad \int_{-\infty}^{\infty} f(x)\,dx=1
    $$
    이고,
    $$
    P(a\le X\le b)=\int_a^b f(x)\,dx
    $$
    로 구간 확률을 계산한다.
  - 누적분포함수는
    $$
    F_X(x)=P(X\le x)
    $$
    로 정의되며, 분포를 전체적으로 보는 가장 일반적인 방식이다.
- Theorem
  - 이산형 확률분포의 정규화:
    $$
    \sum_x P(X=x)=1
    $$
  - 기대값과 분산:
    $$
    E(X)=\sum_x xP(X=x),\qquad \operatorname{Var}(X)=E(X^2)-[E(X)]^2
    $$
  - 연속형 확률변수에서는 한 점의 확률이 0이다.
    $$
    P(X=x)=0
    $$
- 대표 수식
  - $P(X=x)$
  - $E(X)=\sum xP(X=x)$
  - $\operatorname{Var}(X)=E[(X-E(X))^2]$
  - $F_X(x)=P(X\le x)$
- 증명 스케치
  - `증명 스케치 (추론)`:
  - 분포는 `가능한 값들의 목록`과 `각 값에 붙는 확률`을 함께 정리한다.
  - 이산형에서는 모든 값의 확률을 합치면 전체가 되어야 하므로 합이 1이어야 한다.
  - 연속형에서는 곡선 아래 넓이가 전체 확률 1을 이룬다.
- Worked Examples
  - 예제 1: 공정한 주사위를 한 번 던질 때 눈의 수를 $X$라 하자.
    - 지지도는 $\{1,2,3,4,5,6\}$
    - 각 값의 확률은 모두 $\frac16$
    - 따라서
      $$
      E(X)=\frac{1+2+3+4+5+6}{6}=3.5
      $$
      이다.
    - 분산은
      $$
      \operatorname{Var}(X)=\frac{35}{12}
      $$
      이다.
  - 예제 2: 동전을 두 번 던져 앞면의 개수를 $Y$라 하자.
    - 지지도는 $\{0,1,2\}$
    - $P(Y=0)=\frac14,\;P(Y=1)=\frac12,\;P(Y=2)=\frac14$
    - 따라서
      $$
      E(Y)=0\cdot\frac14+1\cdot\frac12+2\cdot\frac14=1
      $$
      이고,
      $$
      \operatorname{Var}(Y)=\frac12
      $$
      이다.
  - 예제 3: $X$가 $[0,1]$에서 균등분포를 따른다고 하자.
    - 밀도함수는 $f(x)=1\;(0\le x\le 1)$, 그 밖에서는 $0$이다.
    - 따라서
      $$
      P\left(\frac14\le X\le \frac34\right)=\int_{1/4}^{3/4}1\,dx=\frac12
      $$
    - 한 점의 확률은 $P(X=\tfrac12)=0$이다.
  - 예제 4: 확률변수 $X$가 `1,2,3`을 각각 확률 `0.2,0.5,0.3`으로 가지면
    $$
    E(X)=1\cdot 0.2+2\cdot 0.5+3\cdot 0.3=2.1
    $$
    이다.
- 교육과정 배치
  - 한국 대표 배치에서는 `고2`의 `확률과 통계` 핵심 개념으로 읽는 것이 자연스럽다.
  - 선수 개념은 중학교의 `확률`, `산포`, `상관`과 고1의 `경우의 수`다.
- 국가별 배치 스냅샷
  - 한국: `고2`의 `확률과 통계`에서 `확률변수`, `기대값`, `분산`과 함께 읽는 것이 자연스럽다.
  - 일본: 중학교에서는 경험적 확률과 경우의 수를 먼저 다루고, 고등학교 `수학A`의 확률과 `수학B`의 통계적 추측이 분포 개념의 전단계를 이룬다.
  - 중국: `9학년`의 `확률의 기초` 뒤에 고등학교 `확률과 통계`, 그리고 선택필수의 `확률변수·정규분포·회귀`로 확장된다.
  - 미국: `Grade 6~8`에서 표본과 확률모형을 먼저 배우고, 고등학교 `Statistics and Data Analysis` 또는 `AP Statistics`에서 `distribution`, `inference`, `regression`, `probability`가 한 과목으로 묶인다.
- 표현과 문제 감각
  - 다국어 용어: `probability distribution`, `確率分布`, `概率分布`
  - 비교 문제집 기준으로 중국 고교와 미국 Statistics는 분포와 모델링이 더 직접적이고, 일본은 경우의 수 설명, 한국은 기대값 계산 비중이 크다.
  - 분포는 `확률변수의 성격표`라고 이해하면 가장 빠르다.

## Common Pitfalls

- 확률분포를 단순한 값 목록으로만 보고 확률을 빠뜨린다.
- 이산형과 연속형을 같은 방식으로 계산한다.
- 연속형 분포에서 $P(X=x)$를 양수로 해석한다.
- 분포를 모르면서 기대값이나 분산부터 계산하려 한다.
- 지지도와 값의 범위를 혼동한다.

## Connections

- 선수 개념은 [counting-principles.md](./counting-principles.md), `확률`, [conditional-probability.md](./conditional-probability.md), [random-variable.md](./random-variable.md)다.
- 다음 개념으로는 [variance.md](./variance.md), [standard-deviation.md](./standard-deviation.md), [normal-distribution.md](./normal-distribution.md), [statistical-inference.md](./statistical-inference.md), [practical-statistics.md](./practical-statistics.md)가 이어진다.
- 과목 허브에서는 [probability-and-statistics-course.md](./probability-and-statistics-course.md)를 본다.
- 계통 허브로는 [../statistics-and-probability-strand.md](../statistics-and-probability-strand.md)를 함께 본다.

## Open Questions

- 이산분포 중심으로 시작한 뒤 연속분포를 후속 카드로 둘지 구조 선택이 필요하다.
- 대표적인 이산분포와 연속분포를 각각 독립 카드로 더 쪼갤지 기준이 필요하다.

## Sources

- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-concept-encyclopedia/japan.md`
- `docs/math-concept-encyclopedia/china.md`
- `docs/math-concept-encyclopedia/us.md`
- `docs/math-curriculum-research/korea.md`
- `docs/math-curriculum-research/japan.md`
- `docs/math-curriculum-research/china.md`
- `docs/math-curriculum-research/us.md`
- `docs/math-concept-encyclopedia/multilingual-glossary.md`
- `docs/math-concept-encyclopedia/formula-examples.md`
- `docs/math-concept-encyclopedia/unit-practice-book.md`
- `docs/math-concept-encyclopedia/comparative-problem-book.md`
