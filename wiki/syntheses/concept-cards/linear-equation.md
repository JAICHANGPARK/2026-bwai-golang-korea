---
title: 일차방정식
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
  - docs/math-concept-encyclopedia/unit-practice-book.md
  - docs/math-concept-encyclopedia/comparative-problem-book.md
tags:
  - concept
  - algebra
  - equation
---

# 일차방정식

## Summary

일차방정식은 미지수가 1차로만 나타나는 방정식이며, 등식의 성질을 이용해 해를 구하는 가장 기본적인 대수 개념이다. 문자와 식에서 시작해 연립방정식, 함수, 비례 문제로 이어지는 관문 역할을 한다.

## Definition

- 일차방정식은 미지수의 차수가 1인 방정식이다.
- 가장 일반적인 한 변수 일차방정식은 $ax+b=c\;(a\neq 0)$로 정리할 수 있다.
- 표준형으로 쓰면 $ax+b=0$이며, 이때 해는 $x=-\frac{b}{a}$이다.

## Key Ideas

- 등식의 성질
  - 양변에 같은 수를 더하거나 빼도, 같은 수를 곱하거나 0이 아닌 같은 수로 나눠도 해집합은 변하지 않는다.
- 이항
  - 항을 반대편으로 옮기는 것은 실제로는 양변에 같은 수를 더하거나 빼는 것과 같다.
- 해와 검산
  - 구한 해를 원식에 대입해 참인지 확인한다.
- 항등식과 조건식 구분
  - 어떤 값에서만 참이 되는지, 모든 값에서 참인지 구분해야 한다.
- 일반형 정리
  - 일차식은 전개와 정리를 거치면 언제나 $ax+b=c$ 꼴로 모일 수 있다.

## Theorems and Properties

- 해의 유일성
  - $a\neq 0$이면 $ax+b=c$는 해가 정확히 하나이다.
- 특수한 경우
  - $a=0$이면 식은 상수식으로 바뀐다.
  - $b=c$이면 모든 수가 해가 되고, $b\neq c$이면 해가 없다.
- 해집합 보존
  - 등식의 성질로 만든 변형식은 원식과 동치이며 해집합이 같다.

## Proof Sketch

`ax+b=c`에서 시작하자. 양변에 $-b$를 더하면 $ax=c-b$가 된다. 이제 $a\neq 0$이므로 양변을 $a$로 나누어

$$
x=\frac{c-b}{a}
$$

를 얻는다. 표준형 $ax+b=0$은 $c=0$인 특수한 경우이므로 $x=-\frac{b}{a}$가 된다. 이 증명은 등식의 성질이 해집합을 바꾸지 않는다는 사실에 기대고 있다.

## Worked Examples

- 예제 1
  - $3x-5=16$
  - $3x=21$
  - $x=7$
  - 검산: $3\cdot 7-5=16$
- 예제 2
  - $2(3x-1)-x=11$
  - $6x-2-x=11$
  - $5x=13$
  - $x=\frac{13}{5}$
  - 이 예제는 전개, 동류항 정리, 이항이 한 번에 연결된다는 점을 보여 준다.

## Common Pitfalls

- 부호 실수
  - $-(x-3)$를 $-x-3$로 쓰면 안 되고 $-x+3$이어야 한다.
- 0으로 나누기
  - 변수 쪽 계수를 나눌 때 그 값이 0인지 먼저 확인해야 한다.
- 이항을 기계적으로 외우기
  - 실제로는 양변에 같은 수를 더하거나 빼는 과정이다.
- 검산 생략
  - 계산은 맞아도 원식에 대입하면 틀릴 수 있다.

## Connections

- 선수 개념은 [algebraic-manipulation.md](./algebraic-manipulation.md), [proportion.md](./proportion.md), [variables-and-expressions.md](./variables-and-expressions.md)다.
- 다음 개념으로는 [simultaneous-equations.md](./simultaneous-equations.md), [linear-function.md](./linear-function.md), [linear-inequality.md](./linear-inequality.md)가 이어진다.
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
- `docs/math-concept-encyclopedia/unit-practice-book.md`
- `docs/math-concept-encyclopedia/comparative-problem-book.md`
