---
title: 복소수의 나눗셈
type: synthesis
status: active
updated: 2026-04-09
source_docs:
  - docs/math-concept-encyclopedia/us.md
  - docs/math-concept-encyclopedia/china.md
  - docs/math-concept-encyclopedia/unit-practice-book.md
tags:
  - concept
  - algebra
  - complex-number
---

# 복소수의 나눗셈

## Summary

복소수의 나눗셈은 분모의 허수를 없애기 위해 켤레복소수를 곱하는 계산이다. [complex-number-operations.md](./complex-number-operations.md)의 켤레 개념이 실제로 왜 중요한지를 가장 잘 보여 주는 하위 카드다.

## Key Points

- 정의
  - 복소수의 나눗셈은 분모와 분자에 분모의 켤레복소수를 곱해 분모를 실수로 바꾸어 계산한다.
- 핵심 개념
  - 켤레복소수
  - 분모 유리화와 유사한 발상
  - $i^2=-1$
  - 실수 분모 만들기
- 대표 수식
  - $\frac{a+bi}{c+di}=\frac{(a+bi)(c-di)}{c^2+d^2}\quad (c+di\neq 0)$
- 핵심 해석
  - 복소수의 나눗셈은 새 규칙이 아니라 `곱셈 + 켤레`의 응용이다.

## Deep Dive

### 핵심 정리: 분모에 켤레를 곱하면 실수 분모가 된다

- 증명
  - 분모가 $c+di$라고 하자.
  - 분모와 분자에 $c-di$를 곱하면
    $$
    \frac{a+bi}{c+di}
    =
    \frac{(a+bi)(c-di)}{(c+di)(c-di)}
    $$
    이다.
  - 분모는
    $$
    (c+di)(c-di)=c^2+d^2
    $$
    로 실수가 된다.
  - 따라서 나눗셈을 실수 분모를 가진 형태로 바꿀 수 있다.

## Worked Examples

### 예제: 복소수 나눗셈


  $$
  \frac{1+2i}{1-i}
  $$
  를 계산하자.
- 분자와 분모에 $1+i$를 곱하면
  $$
  \frac{(1+2i)(1+i)}{(1-i)(1+i)}
  $$
  이다.
- 분자는
  $$
  1+i+2i+2i^2= -1+3i
  $$
  이고, 분모는
  $$
  1-(-1)=2
  $$
  이다.
- 따라서
  $$
  \frac{1+2i}{1-i}=\frac{-1+3i}{2}
  $$
  이다.

## Common Pitfalls

- 분모만 켤레로 바꾸고 분자까지 곱하지 않는 실수가 많다.
- $(c+di)(c-di)=c^2-d^2$로 잘못 쓰면 안 된다.
- $i^2=-1$ 정리를 마지막에 놓치면 분자가 잘못 정리된다.

## Connections

- 선수 개념은 [complex-number-operations.md](./complex-number-operations.md), [complex-numbers.md](./complex-numbers.md), [algebraic-manipulation.md](./algebraic-manipulation.md)다.
- 같은 축의 인접 개념으로는 [complex-plane.md](./complex-plane.md), [algebra-2.md](./algebra-2.md)가 있다.
- 다음 개념으로는 [complex-plane.md](./complex-plane.md), [polar-ideas.md](./polar-ideas.md), [vectors.md](./vectors.md)가 이어진다.

## Open Questions

- 극형식과 드무아브르 정리까지 이어지는 별도 심화 카드가 필요한지 검토가 필요하다.

## Sources

- `docs/math-concept-encyclopedia/us.md`
- `docs/math-concept-encyclopedia/china.md`
- `docs/math-concept-encyclopedia/unit-practice-book.md`
