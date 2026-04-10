---
title: 확률변수
type: synthesis
status: active
updated: 2026-04-10
card_role: concept
source_docs:
  - docs/math-concept-encyclopedia/korea.md
  - docs/math-concept-encyclopedia/china.md
  - docs/math-concept-encyclopedia/us.md
  - docs/math-concept-encyclopedia/japan.md
  - docs/math-curriculum-research/korea.md
  - docs/math-curriculum-research/china.md
  - docs/math-curriculum-research/us.md
  - docs/math-curriculum-research/japan.md
  - docs/math-concept-encyclopedia/multilingual-glossary.md
  - docs/math-concept-encyclopedia/formula-examples.md
  - docs/math-concept-encyclopedia/unit-practice-book.md
  - docs/math-concept-encyclopedia/comparative-problem-book.md
tags:
  - concept
  - probability
  - statistics
---

# 확률변수

## Summary

확률변수는 우연한 결과를 숫자로 바꾸어 다루는 개념이다. 경우의 수와 확률을 계산 단계에서 끝내지 않고, 기대값·분산·확률분포 같은 통계적 구조로 넘어가게 해 주는 핵심 번역 장치다.

## Key Points

- 정의
  - 확률공간 $(\Omega,\mathcal F,P)$ 위에서 정의된 함수
    $$
    X:\Omega\to\mathbb R
    $$
    를 확률변수라 한다.
  - 각 결과 $\omega\in\Omega$에 실수 $X(\omega)$를 대응시켜 우연한 결과를 수로 번역한다.
- 핵심 개념
  - 표본공간
  - 사건
  - 값의 대응
  - 기대값
  - 분산
  - 분포
- Deep Dive
  - 확률변수는 `사건` 자체가 아니라 `사건을 수치로 읽는 규칙`이다.
  - 그래서 확률변수의 진짜 역할은 결과를 숫자로 바꾸는 것이 아니라, 숫자 위에서 확률분포를 만들게 하는 데 있다.
  - 이 관점에서 사건 $\{X=x\}$, $\{X\le a\}$, $\{a\le X\le b\}$는 모두 다시 확률을 묻는 새로운 사건이 된다.
- Theorem
  - 학교 수준에서 주로 다루는 이산확률변수에 대해 선형성:
    $$
    E(aX+b)=aE(X)+b
    $$
  - 학교 수준에서 주로 다루는 이산확률변수에 대해 가법성:
    $$
    E(X+Y)=E(X)+E(Y)
    $$
  - 독립인 이산확률변수 $X,Y$에 대해:
    $$
    E(XY)=E(X)E(Y)
    $$
  - 분산의 이동·확대:
    $$
    \operatorname{Var}(aX+b)=a^2\operatorname{Var}(X)
    $$
- 대표 수식
  - $X:\Omega\to\mathbb{R}$
  - $F_X(x)=P(X\le x)$
  - 이산확률변수에서는 $P(X=x)$와
    $$
    E(X)=\sum_x xP(X=x)
    $$
    를 쓴다.
  - $\operatorname{Var}(X)=E[(X-E(X))^2]$
- 증명 스케치
  - `증명 스케치 (추론)`:
  - 학교 수준에서 먼저 다루는 `이산확률변수`를 생각하자.
  - 이산확률변수 $X$가 취하는 값마다 사건 $\{X=x\}$를 생각하면, 이 사건들은 서로 겹치지 않고 가능한 모든 결과를 나눈다.
  - 따라서 각 값에 붙는 확률 $P(X=x)$는 모두 0 이상이고, 가능한 모든 값에 대한 합은 1이 된다.
  - 이 구조 위에서 기대값은 `값 x 와 그 값이 나올 확률`의 가중합으로 정의된다.
- Worked Examples
  - 예제 1: 공정한 주사위를 한 번 던지고, 눈의 수를 $X$라 하자.
    - $X(\omega)\in\{1,2,3,4,5,6\}$
    - $P(X=x)=\frac16\quad(x=1,\dots,6)$
    - 기대값은
      $$
      E(X)=\frac{1+2+3+4+5+6}{6}=3.5
      $$
      이다.
  - 예제 2: 동전을 두 번 던지고 앞면의 개수를 $Y$라 하자.
    - $Y\in\{0,1,2\}$
    - $P(Y=0)=\frac14,\;P(Y=1)=\frac12,\;P(Y=2)=\frac14$
    - 따라서
      $$
      E(Y)=0\cdot\frac14+1\cdot\frac12+2\cdot\frac14=1
      $$
    - 같은 방식으로 분산도 계산할 수 있다.
  - 예제 3: $X$가 `1,2,3`을 각각 확률 `0.2, 0.5, 0.3`으로 가질 때
    $$
    E(X)=1\cdot 0.2+2\cdot 0.5+3\cdot 0.3=2.1
    $$
    이다.
- 교육과정 배치
  - 한국 대표 배치에서는 `고2`의 `확률과 통계`에서 본격적으로 등장한다.
  - 이후 [probability-distribution.md](./probability-distribution.md), `통계적 추정`, `회귀`, `실용 통계`로 이어진다.
- 국가별 배치 스냅샷
  - 한국: `확률과 통계`에서 `기대값`, `분산`, `확률분포`와 함께 묶인다.
  - 일본: 문서상 `확률변수`가 독립 표제로 강하게 드러나기보다 `수학A`의 확률과 `수학B`의 통계적 추측 사이의 상위 통계 개념으로 읽힌다.
  - 중국: 고등학교 선택필수의 `확률변수·정규분포·회귀`에서 명시적으로 등장한다.
  - 미국: 보통 고등학교 `Statistics and Data Analysis`, `AP Statistics`, 또는 상위 경로에서 분포와 함께 다뤄진다.
- 표현과 문제 감각
  - 다국어 용어: `random variable`, `確率変数`, `随机变量`
  - 비교 문제집 기준으로 한국은 기대값 계산, 일본은 경우의 수 설명, 중국과 미국은 분포와 모델링 연결이 더 직접적이다.
  - 확률변수는 `결과`와 `숫자` 사이의 다리다.

## Common Pitfalls

- 확률변수와 그 값 하나를 같은 것으로 착각한다.
- $P(X=x)$와 $P(X\le x)$를 혼동한다.
- 연속형 확률변수에서 $P(X=x)$를 양수로 생각한다.
- 기대값을 `가장 자주 나오는 값`과 같은 것으로 오해한다.
- 사건과 확률변수의 기호를 섞어서 쓴다.

## Connections

- 선수 개념은 [counting-principles.md](./counting-principles.md), `확률`, [conditional-probability.md](./conditional-probability.md)다.
- 바로 다음 개념은 [variance.md](./variance.md), [standard-deviation.md](./standard-deviation.md), [probability-distribution.md](./probability-distribution.md), [statistical-inference.md](./statistical-inference.md)다.
- 과목 허브로는 [probability-and-statistics-course.md](./probability-and-statistics-course.md)와 연결된다.
- 계통 허브로는 [../statistics-and-probability-strand.md](../statistics-and-probability-strand.md)를 함께 본다.

## Open Questions

- 연속확률변수와 이산확률변수를 같은 카드에서 다룰지 후속 분화가 필요하다.
- `기대값`을 [variance.md](./variance.md)와 대칭되는 독립 카드로 승격할지 검토가 남아 있다.

## Sources

- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-concept-encyclopedia/china.md`
- `docs/math-concept-encyclopedia/us.md`
- `docs/math-concept-encyclopedia/japan.md`
- `docs/math-curriculum-research/korea.md`
- `docs/math-curriculum-research/china.md`
- `docs/math-curriculum-research/us.md`
- `docs/math-curriculum-research/japan.md`
- `docs/math-concept-encyclopedia/multilingual-glossary.md`
- `docs/math-concept-encyclopedia/formula-examples.md`
- `docs/math-concept-encyclopedia/unit-practice-book.md`
- `docs/math-concept-encyclopedia/comparative-problem-book.md`
