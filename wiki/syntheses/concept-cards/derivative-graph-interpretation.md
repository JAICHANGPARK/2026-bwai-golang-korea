---
title: 도함수의 그래프 해석
type: synthesis
status: active
updated: 2026-04-10
card_role: concept
source_docs:
  - docs/math-concept-encyclopedia/korea.md
  - docs/math-concept-encyclopedia/us.md
  - docs/math-concept-encyclopedia/japan.md
  - docs/math-concept-encyclopedia/china.md
  - docs/math-concept-encyclopedia/formula-examples.md
  - docs/math-concept-encyclopedia/unit-practice-book.md
tags:
  - concept
  - calculus
  - graph
---

# 도함수의 그래프 해석

## Summary

도함수의 그래프 해석은 [derivative.md](./derivative.md)에서 구한 $f'(x)$를 다시 읽어, 원래 함수 $f(x)$의 증가·감소와 극값을 판단하는 카드다. 계산 결과를 함수 해석으로 되돌리는 지점이어서, 교과과정 안에서도 풀이력 차이가 크게 나는 부분이다.

## Key Points

- 정의
  - 도함수의 부호와 값의 변화를 이용해 원래 함수의 증가·감소, 극대·극소, 접선의 기울기 변화를 읽는 것을 말한다.
- 핵심 개념
  - $f'(x)>0$, $f'(x)<0$
  - 증가와 감소
  - 극값
  - 접선 기울기
  - 부호표
- 대표 수식
  - $f'(x)>0 \Rightarrow f$ 증가
  - $f'(x)<0 \Rightarrow f$ 감소
  - $f'(a)=0$은 극값 후보
- 핵심 해석
  - 도함수는 원래 함수의 `오르내림 지도`다.

## Deep Dive

### 핵심 정리 1: 도함수의 부호와 증가·감소

- `증명 스케치 (추론)`:
  - 어떤 구간에서 $f'(x)>0$이면 접선 기울기가 양수라는 뜻이다.
  - 이는 함수가 오른쪽으로 갈수록 올라가는 경향을 뜻하므로 그 구간에서 함수는 증가한다.
  - 마찬가지로 $f'(x)<0$이면 함수는 감소한다.

### 핵심 정리 2: $f'(a)=0$이면 극값 후보다

- `증명 스케치 (추론)`:
  - 극대나 극소에서는 그래프가 위로 올라가다가 내려오거나, 내려가다가 올라오는 전환이 일어난다.
  - 이때 접선이 수평이 되는 경우가 많아
    $$
    f'(a)=0
    $$
    이 된다.
  - 다만 $f'(a)=0$이라고 해서 항상 극값은 아니므로, 부호 변화를 함께 봐야 한다.

## Worked Examples

### 예제: $f(x)=x^3-3x$의 증감

- 도함수는
  $$
  f'(x)=3x^2-3=3(x-1)(x+1)
  $$
  이다.
- 따라서
  - $x<-1$에서 $f'(x)>0$
  - $-1<x<1$에서 $f'(x)<0$
  - $x>1$에서 $f'(x)>0$
  이다.
- 원래 함수는
  - $(-\infty,-1)$에서 증가
  - $(-1,1)$에서 감소
  - $(1,\infty)$에서 증가
  한다.
- 따라서 $x=-1$에서 극대, $x=1$에서 극소를 가진다.

## Common Pitfalls

- $f'(a)=0$이면 무조건 극값이라고 생각하면 안 된다.
- 도함수의 그래프와 원래 함수의 그래프를 같은 것으로 보면 해석이 뒤섞인다.
- 부호표를 만들 때 인수의 부호를 한꺼번에 잘못 읽는 실수가 많다.

## Connections

- 선수 개념은 [derivative.md](./derivative.md), [differentiation.md](./differentiation.md), [quadratic-function.md](./quadratic-function.md)다.
- 같은 축의 인접 개념으로는 [integration.md](./integration.md), [calculus-course.md](./calculus-course.md)가 있다.
- 다음 개념으로는 [calculus-course.md](./calculus-course.md), [calculus-2.md](./calculus-2.md), [statistics-and-data-analysis.md](./statistics-and-data-analysis.md)가 이어진다.

## Open Questions

- 이계도함수와 오목·볼록 해석까지 별도 카드로 더 뺄지 기준이 필요하다.

## Sources

- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-concept-encyclopedia/us.md`
- `docs/math-concept-encyclopedia/japan.md`
- `docs/math-concept-encyclopedia/china.md`
- `docs/math-concept-encyclopedia/formula-examples.md`
- `docs/math-concept-encyclopedia/unit-practice-book.md`
