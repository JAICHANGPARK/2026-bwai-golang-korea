---
title: 적분
type: synthesis
status: active
updated: 2026-04-10
card_role: concept
source_docs:
  - docs/math-concept-encyclopedia/korea.md
  - docs/math-concept-encyclopedia/japan.md
  - docs/math-concept-encyclopedia/china.md
  - docs/math-concept-encyclopedia/us.md
  - docs/math-concept-encyclopedia/multilingual-glossary.md
tags:
  - concept
  - calculus
  - integral
---

# 적분

## Summary

적분은 넓이와 누적량을 읽는 개념이다. 잘게 나눈 양을 모아 전체를 계산하는 언어라서, 미분과 짝을 이루며 변화와 누적을 함께 설명한다.

## Definition

- 적분은 작은 변화량을 누적해 전체량을 구하는 개념이다.
- 학교 수학에서는 크게 두 뜻이 함께 쓰인다.
  - 부정적분: 도함수가 주어진 함수가 되는 원시함수의 가족
  - 정적분: 구간 전체의 누적량 또는 부호를 가진 넓이
- 대표 표기는
  $$
  \int_a^b f(x)\,dx,\qquad \int f(x)\,dx
  $$
  이다.

## Key Ideas

- 누적량
  - 적분은 작은 조각들을 더해 전체를 구하는 언어다.
- 정적분과 부정적분
  - 정적분은 `값`, 부정적분은 `함수의 가족`이라는 점을 구분해야 한다.
- 넓이와 부호
  - x축 위에서는 양의 넓이, 아래에서는 음의 넓이로 계산한다.
- 미분과 적분의 연결
  - 적분은 [differentiation.md](./differentiation.md)과 짝을 이루며, [derivative.md](./derivative.md)의 역방향으로도 읽을 수 있다.

## Deep Dive

- 적분의 핵심은 `한 번에 못 세는 전체를 잘게 나눠 더한다`는 생각이다.
- 그래서 넓이 문제뿐 아니라 거리, 질량, 누적 비용, 확률 같은 여러 개념이 적분으로 이어진다.

## Theorems and Properties

- 선형성
  -
    $$
    \int_a^b (af(x)+bg(x))\,dx
    =
    a\int_a^b f(x)\,dx+b\int_a^b g(x)\,dx
    $$
- 미적분의 기본정리
  - $F'(x)=f(x)$이면
    $$
    \int_a^b f(x)\,dx=F(b)-F(a)
    $$
    이다.
- 거듭제곱의 부정적분
  - $n\neq -1$일 때
    $$
    \int x^n\,dx=\frac{x^{n+1}}{n+1}+C
    $$
    이다.

## Proof Sketch

- `증명 스케치 (추론)`:
- 구간 $[a,b]$를 아주 작은 폭 $\Delta x$로 나누고, 각 구간에서 높이를 $f(x)$라고 생각하면 작은 직사각형 넓이는 대략 $f(x)\Delta x$다.
- 이런 작은 넓이들을 모두 더하면 전체 넓이에 가까워진다.
- 분할을 무한히 잘게 하면 이 합의 극한이 정적분
  $$
  \int_a^b f(x)\,dx
  $$
  이 된다.
- 한편 원시함수 $F$가 있으면 이 누적량을 $F(b)-F(a)$로 계산할 수 있는데, 이것이 미적분의 기본정리다.

## Worked Examples

- 예제 1: 정적분 계산
  -
    $$
    \int_0^2 x\,dx
    $$
    를 구하자.
  - 원시함수는 $F(x)=\frac{x^2}{2}$이므로
    $$
    \int_0^2 x\,dx=\left[\frac{x^2}{2}\right]_0^2=\frac{4}{2}-0=2
    $$
    이다.
- 예제 2: 부정적분 계산
  -
    $$
    \int (3x^2-4x+1)\,dx
    $$
    는 항별로 적분하면
    $$
    x^3-2x^2+x+C
    $$
    이다.
- 예제 3: 넓이 해석
  - $y=2$와 x축 사이의 넓이를 $x=1$부터 $x=4$까지 구하면
    $$
    \int_1^4 2\,dx=2(4-1)=6
    $$
    이다.

## Common Pitfalls

- 정적분과 부정적분을 같은 대상으로 착각하기 쉽다.
- 부정적분에서 적분상수 $C$를 빠뜨리면 안 된다.
- x축 아래 영역도 무조건 양수 넓이로 처리하면 정적분 해석이 틀어진다.
- 적분 공식을 외워도 원시함수와 구간 대입 순서를 헷갈리기 쉽다.

## Curriculum Placement

- 한국 대표 배치에서는 `미적분I`의 핵심이며 `미적분II`에서 더 고급화된다.
- 국가별 배치 스냅샷
  - 한국: `미적분I`와 `미적분II`로 확장된다.
  - 일본: `수학II`에서 적분의 생각을, `수학III`에서 본격 적분법을 다룬다.
  - 중국: 도함수와 함께 고등 미적분 축에서 다뤄진다.
  - 미국: Calculus의 accumulation and area unit이 대표적이다.
- 표현 메모
  - 다국어 용어: `integral`, `integration`, `積分`, `积分`

## Connections

- 선수 개념은 [limits.md](./limits.md), [differentiation.md](./differentiation.md)다.
- 같은 축의 인접 개념으로는 [derivative.md](./derivative.md), [calculus-2.md](./calculus-2.md)가 있다.
- 다음 개념으로는 [economics-math.md](./economics-math.md), [normal-distribution.md](./normal-distribution.md), [calculus-course.md](./calculus-course.md)가 이어진다.
- 과목 허브로는 [calculus-1.md](./calculus-1.md), [calculus-2.md](./calculus-2.md), [calculus-course.md](./calculus-course.md)를 본다.
- 계통 허브로는 [../functions-strand.md](../functions-strand.md), [../high-2-hub.md](../high-2-hub.md), [../high-3-hub.md](../high-3-hub.md)를 함께 본다.

## Open Questions

- 정적분과 부정적분을 각각 분리할지 후속 설계가 필요하다.

## Sources

- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-concept-encyclopedia/japan.md`
- `docs/math-concept-encyclopedia/china.md`
- `docs/math-concept-encyclopedia/us.md`
- `docs/math-concept-encyclopedia/multilingual-glossary.md`
