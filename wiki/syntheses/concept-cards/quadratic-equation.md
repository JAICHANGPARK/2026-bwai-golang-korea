---
title: 이차방정식
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
  - algebra
  - equation
---

# 이차방정식

## Summary

이차방정식은 미지수가 2차로 나타나는 방정식의 해를 구하는 개념이다. 인수분해, 완전제곱식, 근의 공식을 묶어 식의 구조를 읽게 해 주며, [quadratic-function.md](./quadratic-function.md)의 x절편 해석과 직접 연결되는 핵심 브리지다.

## Key Points

- 정의
  - 가장 대표적인 형태는 $ax^2+bx+c=0\;(a\neq 0)$이다.
- 핵심 개념
  - 근
  - 인수분해 풀이
  - 완전제곱식
  - 근의 공식
  - 판별식
- 대표 수식
  - $ax^2+bx+c=0$
  - $x=\frac{-b\pm\sqrt{b^2-4ac}}{2a}$
- 핵심 해석
  - 일차방정식이 직선과의 만남을 다룬다면, 이차방정식은 포물선과 x축의 만남을 읽는 방정식이다.
  - 풀이법은 `인수분해`, `완전제곱`, `근의 공식`으로 나뉘지만 모두 같은 해를 구한다.
- 교육과정 배치
  - 한국 대표 배치에서는 `중3` 핵심 단원이다.
  - 이후 [quadratic-function.md](./quadratic-function.md), [equations-of-geometric-figures.md](./equations-of-geometric-figures.md), [complex-numbers.md](./complex-numbers.md)의 배경으로 이어진다.
- 국가별 배치 스냅샷
  - 한국: `중3`에서 인수분해와 함께 배우고, 함수의 x절편 해석으로 이어진다.
  - 일본: `중3`에서 배우며, 근의 공식의 절차 설명과 판별식 해석이 중요하다.
  - 중국: `9학년 1학기`에 도입되고, 계산 구조와 함수 연결이 빠르게 이어진다.
  - 미국: 보통 `Algebra I` 또는 `Integrated Mathematics` 경로에서 그래프와 x절편을 함께 읽으며 배운다.
- 표현과 문제 감각
  - 다국어 용어: `quadratic equation`, `二次方程式`, `一元二次方程`
  - 비교 문제집 기준으로 미국은 그래프와 함께 읽는 경우가 많고, 일본은 절차 설명, 한국·중국은 계산 정확도와 속도가 더 강하게 요구된다.

## Deep Dive

### 왜 해가 많아야 두 개인가

- 이차방정식은 최고차항이 $x^2$이므로 그래프는 [quadratic-function.md](./quadratic-function.md)의 포물선으로 읽힌다.
- 포물선과 x축은 보통 두 점에서 만나거나, 한 점에서 접하거나, 만나지 않는다.
- 그래서 실수 범위에서 해의 개수는 `2개`, `1개`, `0개`로 나뉜다.

### 핵심 정리 1: 근의 공식

- 정리
  - $ax^2+bx+c=0\;(a\neq 0)$의 해는
    $$
    x=\frac{-b\pm\sqrt{b^2-4ac}}{2a}
    $$
    이다.
- 증명
  - 먼저 양변을 $a$로 나누면
    $$
    x^2+\frac{b}{a}x+\frac{c}{a}=0
    $$
    이다.
  - 상수항을 오른쪽으로 넘기면
    $$
    x^2+\frac{b}{a}x=-\frac{c}{a}
    $$
    이다.
  - 양변에 $\left(\frac{b}{2a}\right)^2$를 더해 완전제곱을 만들면
    $$
    x^2+\frac{b}{a}x+\left(\frac{b}{2a}\right)^2
    =
    -\frac{c}{a}+\left(\frac{b}{2a}\right)^2
    $$
    이다.
  - 왼쪽은
    $$
    \left(x+\frac{b}{2a}\right)^2
    $$
    가 되고, 오른쪽은
    $$
    \frac{b^2-4ac}{4a^2}
    $$
    로 정리된다.
  - 따라서
    $$
    \left(x+\frac{b}{2a}\right)^2=\frac{b^2-4ac}{4a^2}
    $$
    이고, 양변에 제곱근을 취하면
    $$
    x+\frac{b}{2a}=\pm \frac{\sqrt{b^2-4ac}}{2a}
    $$
    이다.
  - 마지막으로 $\frac{b}{2a}$를 이항하면
    $$
    x=\frac{-b\pm\sqrt{b^2-4ac}}{2a}
    $$
    를 얻는다.

### 핵심 정리 2: 판별식의 의미

- $D=b^2-4ac$라고 두면
  - $D>0$: 서로 다른 두 실근
  - $D=0$: 중근 하나
  - $D<0$: 실근 없음
  이다.
- 이는 근의 공식 안에서 제곱근 부분 $\sqrt{D}$가 실제로 존재하는지, 0인지, 양수인지에 따라 바로 결정된다.

## Worked Examples

### 예제 1: 인수분해 풀이

- $x^2-4x-5=0$은
  $$
  x^2-4x-5=(x-5)(x+1)
  $$
  로 인수분해된다.
- 따라서
  $$
  (x-5)(x+1)=0
  $$
  이고, 해는 $x=5,-1$이다.

### 예제 2: 근의 공식 풀이

- $2x^2+x-3=0$에서 $a=2$, $b=1$, $c=-3$이다.
- 근의 공식에 넣으면
  $$
  x=\frac{-1\pm\sqrt{1-4\cdot 2\cdot (-3)}}{4}
  =
  \frac{-1\pm\sqrt{25}}{4}
  $$
  이다.
- 따라서
  $$
  x=\frac{-1\pm 5}{4}
  $$
  이므로 해는 $x=1,\,-\frac{3}{2}$이다.

## Common Pitfalls

- 이차방정식을 표준형 $ax^2+bx+c=0$으로 정리하지 않고 바로 공식을 쓰면 부호 실수가 잦다.
- 인수분해가 되는 식인데도 무조건 근의 공식부터 쓰면 구조를 읽는 힘이 약해진다.
- 판별식을 계산할 때 $b^2-4ac$의 `-4ac` 부호를 특히 많이 틀린다.
- 해를 구한 뒤 [quadratic-function.md](./quadratic-function.md)의 x절편과 연결하지 않으면 그래프 해석이 끊기기 쉽다.

## Connections

- 선수 개념은 [polynomials-and-factorization.md](./polynomials-and-factorization.md), [linear-equation.md](./linear-equation.md), [square-roots-and-real-numbers.md](./square-roots-and-real-numbers.md)이다.
- 같은 축의 인접 개념으로는 [quadratic-function.md](./quadratic-function.md)가 있다.
- 다음 개념으로는 [quadratic-function.md](./quadratic-function.md), [equations-of-geometric-figures.md](./equations-of-geometric-figures.md), [complex-numbers.md](./complex-numbers.md)가 이어진다.
- 학년 허브에서는 [middle-3-hub.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/syntheses/middle-3-hub.md), [high-1-hub.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/syntheses/high-1-hub.md)와 연결된다.

## Open Questions

- `근의 공식`과 `판별식`을 이 카드 안에 유지할지 독립 카드로 분리할지 기준이 필요하다.
- 미국식 그래프 중심 풀이와 한국식 계산 중심 풀이를 별도 비교 synthesis로 뺄지 검토가 필요하다.

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
