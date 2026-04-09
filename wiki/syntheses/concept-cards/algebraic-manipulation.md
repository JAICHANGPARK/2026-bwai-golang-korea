---
title: 식의 계산
type: synthesis
status: active
updated: 2026-04-09
source_docs:
  - docs/math-concept-encyclopedia/korea.md
  - docs/math-concept-encyclopedia/japan.md
  - docs/math-concept-encyclopedia/china.md
  - docs/math-concept-encyclopedia/us.md
  - docs/math-concept-encyclopedia/multilingual-glossary.md
tags:
  - concept
  - algebra
  - manipulation
---

# 식의 계산

## Summary

식의 계산은 문자식을 전개하고, 동류항을 모으고, 인수로 묶어 구조를 정리하는 대수의 기본 문법이다. 방정식 풀이와 함수식 변형의 바닥에 깔리는 계산 체력이라고 보면 된다.

## Definition

- 식의 계산은 문자를 포함한 식을 전개, 동류항 정리, 인수분해의 방향으로 다루는 과정이다.
- 단항식, 다항식, 동류항, 계수, 차수의 개념을 함께 사용한다.

## Key Ideas

- 전개
  - 괄호를 풀어 곱셈을 덧셈의 형태로 바꾼다.
- 동류항 정리
  - 변수와 차수가 같은 항끼리만 더하거나 뺄 수 있다.
- 분배법칙
  - 곱셈이 덧셈 위로 퍼질 때 식의 구조를 바꿀 수 있다.
- 인수분해와 역전개
  - 계산의 반대 방향으로 공통인자를 묶어 표현을 간결하게 만든다.
- 표준형 정리
  - 다음 단계의 방정식 풀이를 쉽게 하도록 식을 정돈한다.

## Theorems and Properties

- 분배법칙
  - $(a+b)c=ac+bc$
  - $(a-b)c=ac-bc$
- 동류항 결합
  - $ax+bx=(a+b)x$
  - $x^m$과 $x^n$은 $m=n$일 때만 동류항이다.
- 곱셈의 결합과 교환
  - 순서와 묶음은 바뀌어도 값이 변하지 않아 식을 정리할 수 있다.
- 대표 항등식
  - $(a+b)^2=a^2+2ab+b^2$
  - $(a-b)(a+b)=a^2-b^2$

## Proof Sketch

분배법칙은 `a와 b의 합을 c번 반복한다`는 뜻에서 나온다. 예를 들어 $(a+b)c$는 $(a+b)$를 $c$번 더한 것이므로, 결국 $ac$와 $bc$를 각각 $c$번 더한 뒤 합친 것과 같다. 동류항 결합도 같은 원리다. $x+x+x$는 $3x$로 쓸 수 있고, $ax+bx$는 `x`가 반복된 양이므로 $(a+b)x$로 정리된다.

## Worked Examples

- 예제 1
  - $(2x+3)-(-x+1)$
  - $=2x+3+x-1$
  - $=3x+2$
- 예제 2
  - $2(3x-1)-x$
  - $=6x-2-x$
  - $=5x-2$
- 예제 3
  - $6x^2-9x$
  - $=3x(2x-3)$
  - 이 예제는 공통인자를 묶는 것이 전개와 반대 방향의 계산이라는 점을 보여 준다.

## Common Pitfalls

- 부호 오류
  - $-(x-1)$을 $-x-1$로 쓰면 안 된다.
- 동류항 오판
  - $x$와 $x^2$는 동류항이 아니다.
- 전개 누락
  - 곱셈이 걸린 괄호를 한 항씩 분배하지 않으면 식이 틀어진다.
- 인수분해를 너무 빨리 시도
  - 먼저 전개와 정리를 끝내고 공통인자를 보는 편이 안전하다.

## Connections

- 선수 개념은 [variables-and-expressions.md](./variables-and-expressions.md)와 [integers-and-rational-numbers.md](./integers-and-rational-numbers.md)다.
- 다음 개념으로는 [linear-equation.md](./linear-equation.md), [polynomials-and-factorization.md](./polynomials-and-factorization.md), [linear-function.md](./linear-function.md)가 이어진다.
- 계통 허브로는 [../korea-curriculum-hub.md](../korea-curriculum-hub.md), [../japan-curriculum-hub.md](../japan-curriculum-hub.md), [../china-curriculum-hub.md](../china-curriculum-hub.md), [../us-curriculum-hub.md](../us-curriculum-hub.md), [../course-track-hub.md](../course-track-hub.md)를 본다.

## Sources

- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-concept-encyclopedia/japan.md`
- `docs/math-concept-encyclopedia/china.md`
- `docs/math-concept-encyclopedia/us.md`
- `docs/math-concept-encyclopedia/multilingual-glossary.md`
