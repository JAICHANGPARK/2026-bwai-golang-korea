---
title: 정수와 유리수
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
  - number
  - arithmetic
---

# 정수와 유리수

## Summary

정수와 유리수는 양수와 음수, 0, 분수까지 포함해 수 체계를 확장하는 개념이다. 부호와 절댓값, 역원, 사칙연산 규칙을 익히면서 이후 방정식과 좌표, 함수 입력값 해석의 바탕을 마련한다.

## Key Points

- 정의
  - 정수는 양의 정수, 음의 정수, 0을 포함한 수이고, 유리수는 $\frac{a}{b}$ 꼴로 나타낼 수 있는 수다.
  - 모든 정수는 유리수다. 왜냐하면 $n=\frac{n}{1}$로 쓸 수 있기 때문이다.
- 핵심 개념
  - 부호
  - 절댓값
  - 사칙연산
  - 역수
  - 유한소수와 순환소수
  - 수직선
  - 분수의 동치
- 대표 성질
  - 덧셈 역원: $a+\left(-a\right)=0$
  - 곱셈 역원: $a\neq 0$이면 $\frac{1}{a}$가 존재한다.
  - 부호 규칙: 같은 부호끼리 곱하면 양수, 다른 부호끼리 곱하면 음수다.
  - 유리수 닫힘성: 유리수는 덧셈, 뺄셈, 곱셈, 0이 아닌 수로의 나눗셈에 대해 닫혀 있다.
- 대표 수식
  - $|a|=\begin{cases}a&(a\ge 0)\\-a&(a<0)\end{cases}$
  - $a+\left(-a\right)=0$
  - $\frac{a}{b}=\frac{ka}{kb}$ $(k\neq 0)$
- 증명 스케치
  - `증명 스케치 (추론)`:
  - 절댓값은 부호를 지운 값이 아니라 `0에서 얼마나 떨어져 있는가`를 뜻한다.
  - 그래서 $a$와 $-a$는 방향은 반대지만 절댓값은 같고, 더하면 0이 된다.
  - 분수의 동치 $\frac{a}{b}=\frac{ka}{kb}$는 분자와 분모를 같은 0이 아닌 수로 곱하거나 나누어도 값이 변하지 않는다는 뜻이다.
  - 이 관점이 음수 연산과 수직선 해석의 핵심이다.
## Worked Examples
  - 예제 1: $-3+\frac{5}{2}=-\frac{6}{2}+\frac{5}{2}=-\frac12$이다.
  - 예제 2: $\frac{3}{4}-\frac{5}{6}=\frac{9}{12}-\frac{10}{12}=-\frac{1}{12}$이다.
  - 예제 3: $-2\times(-5)=10$, $(-2)\times 5=-10$이다.
## Common Pitfalls
  - $\frac{a}{b}$에서 $b=0$은 허용되지 않는다.
  - 절댓값은 항상 양수라고 하면 틀린다. $|0|=0$이다.
  - $-(a+b)$를 $-a+b$로 잘못 바꾸면 안 된다.
  - 분수 비교에서 분모가 다르다고 바로 크기를 판단하면 안 된다.
- 교육과정 배치
  - 한국 대표 배치에서는 `중1`의 핵심 수 체계 단원이다.
- 국가별 배치 스냅샷
  - 한국: `중1`에서 정수·유리수와 수직선을 함께 배운다.
  - 일본: `중1`에서 양수·음수와 정수 계산이 분명하게 놓인다.
  - 중국: `7학년 1학기`의 `유리수`가 대표 단원이다.
  - 미국: Grade 6과 Grade 7에서 rational number system extension과 arithmetic로 이어진다.
- 표현과 문제 감각
  - 다국어 용어: `integer`, `rational number`, `整数`, `有理数`
  - 연산 규칙 암기보다 수직선과 방향 감각이 중요하다.

## Connections

- 선수 개념은 [prime-factorization.md](./prime-factorization.md)와 `자연수`다.
- 같은 축의 인접 개념으로는 [variables-and-expressions.md](./variables-and-expressions.md), [linear-equation.md](./linear-equation.md)가 있다.
- 다음 개념으로는 [square-roots-and-real-numbers.md](./square-roots-and-real-numbers.md), [function-foundations.md](./function-foundations.md), [linear-inequality.md](./linear-inequality.md)가 이어진다.
- 중1 허브와 한국 허브는 [middle-1-hub.md](../middle-1-hub.md), [korea-curriculum-hub.md](../korea-curriculum-hub.md)에서 다시 본다.

## Open Questions

- `절댓값`을 독립 카드로 분리할지 추후 결정이 필요하다.

## Sources

- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-concept-encyclopedia/japan.md`
- `docs/math-concept-encyclopedia/china.md`
- `docs/math-concept-encyclopedia/us.md`
- `docs/math-concept-encyclopedia/multilingual-glossary.md`
