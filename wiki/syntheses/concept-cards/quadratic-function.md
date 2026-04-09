---
title: 이차함수
type: synthesis
status: active
updated: 2026-04-09
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
  - function
  - quadratic
---

# 이차함수

## Summary

이차함수는 선형 관계를 넘어 처음으로 본격적인 비선형 변화를 다루는 함수다. 식의 구조, 그래프의 모양, 꼭짓점과 최대·최소를 함께 읽어야 해서 중학교 후반에서 고등학교 함수 해석으로 넘어가는 핵심 브리지 역할을 한다.

## Definition

- 이차함수는 $f(x)=ax^2+bx+c\;(a\neq 0)$ 꼴의 함수다.
- 그래프는 포물선이며, $a$의 부호에 따라 위로 열리거나 아래로 열린다.

## Key Ideas

- 표준형
  - $y=ax^2+bx+c$
- 꼭짓점형
  - $y=a(x-p)^2+q$
  - 꼭짓점은 $(p,q)$이다.
- 축의 방정식
  - $x=-\frac{b}{2a}$
- 최대·최소
  - $a>0$이면 꼭짓점에서 최소값, $a<0$이면 최대값을 갖는다.
- x절편과 방정식
  - 이차함수의 x절편은 대응하는 이차방정식의 해와 같다.

## Theorems and Properties

- 완전제곱에 의한 꼭짓점형
  - 모든 이차식은 완전제곱을 통해 꼭짓점형으로 바꿀 수 있다.
- 축 대칭성
  - 그래프는 축 $x=-\frac{b}{2a}$를 기준으로 좌우 대칭이다.
- 꼭짓점 극값 정리
  - 제곱항은 항상 0 이상이므로, 꼭짓점에서 최대 또는 최소가 결정된다.
- 판별과 교점
  - 방정식 $ax^2+bx+c=0$의 해의 개수는 그래프가 x축과 만나는 점의 개수와 일치한다.

## Proof Sketch

$y=ax^2+bx+c$를 완전제곱하면

$$
y=a\left(x+\frac{b}{2a}\right)^2+\left(c-\frac{b^2}{4a}\right)
$$

가 된다. 제곱항은 항상 $0$ 이상이므로, $a>0$이면 그 값이 가장 작아지는 지점이 꼭짓점이고, $a<0$이면 가장 커지는 지점이 꼭짓점이다. 따라서 축은 $x=-\frac{b}{2a}$가 되고, 꼭짓점형은 그래프의 모양을 바로 읽게 해 준다.

## Worked Examples

- 예제 1
  - $y=x^2-4x+1$
  - $y=(x-2)^2-3$
  - 꼭짓점은 $(2,-3)$, 축은 $x=2$
  - 최소값은 $-3$
- 예제 2
  - $y=-2x^2+8x-3$
  - $y=-2(x-2)^2+5$
  - 꼭짓점은 $(2,5)$, 축은 $x=2$
  - 최대값은 $5$

## Common Pitfalls

- 축의 공식 착각
  - $-\frac{b}{a}$가 아니라 $-\frac{b}{2a}$이다.
- 꼭짓점과 x절편 혼동
  - 꼭짓점은 그래프의 최고·최저점이고, x절편은 $y=0$일 때의 점이다.
- 부호 실수
  - $a<0$이면 포물선이 아래로 열린다.
- 폭의 변화 오해
  - $|a|$가 클수록 포물선은 더 좁아진다.

## Connections

- 선수 개념은 [linear-function.md](./linear-function.md), [polynomials-and-factorization.md](./polynomials-and-factorization.md), [quadratic-equation.md](./quadratic-equation.md)다.
- 다음 개념으로는 [equations-of-geometric-figures.md](./equations-of-geometric-figures.md), [derivative.md](./derivative.md), [continuity.md](./continuity.md)로 이어진다.
- 계통 허브로는 [../korea-curriculum-hub.md](../korea-curriculum-hub.md), [../japan-curriculum-hub.md](../japan-curriculum-hub.md), [../china-curriculum-hub.md](../china-curriculum-hub.md), [../us-curriculum-hub.md](../us-curriculum-hub.md), [../course-track-hub.md](../course-track-hub.md)를 본다.

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
