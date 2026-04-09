---
title: 다항식과 인수분해
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
  - algebra
  - polynomial
---

# 다항식과 인수분해

## Summary

다항식과 인수분해는 여러 항으로 이루어진 식의 구조를 읽고, 곱의 형태로 다시 분해하는 개념이다. 전개와 인수분해를 왕복할 수 있어야 이차방정식, 이차함수, 나머지와 근의 관계가 한 흐름으로 보인다.

## Key Points

- 정의
  - 한 변수에 대하여 항이 유한 개이고 각 항의 차수가 음이 아닌 정수인 식을 다항식이라 한다.
  - 다항식을 곱의 형태로 바꾸는 과정을 인수분해라 한다.
- 핵심 개념
  - 차수
  - 전개
  - 공통인수
  - 완전제곱식
  - 차의 제곱
  - 근과 인수
  - 다항식의 구조
  - 식의 대칭성
- 대표 성질
  - 전개법칙: 곱셈은 덧셈에 대해 분배된다.
  - 차의 제곱 공식: $a^2-b^2=(a-b)(a+b)$
  - 완전제곱식: $x^2+2ax+a^2=(x+a)^2$
  - 근-인수 관계: $p(a)=0$이면 $x-a$는 $p(x)$의 인수다.
- 대표 수식
  - $a^2-b^2=(a-b)(a+b)$
  - $x^2+2ax+a^2=(x+a)^2$
  - $p(x)=(x-a)q(x)+r$에서 $p(a)=r$
- 증명 스케치
  - `증명 스케치 (추론)`:
  - $(a-b)(a+b)$를 전개하면
    $$
    a^2+ab-ab-b^2=a^2-b^2
    $$
    가 된다.
  - 공통인수 인수분해는 분배법칙을 거꾸로 쓴 것이다. 예를 들어 $ax+ay=a(x+y)$다.
  - $p(x)=(x-a)q(x)+r$이면 $x=a$를 넣었을 때 $p(a)=r$가 되므로, $p(a)=0$이면 나머지 $r$이 0이어야 한다. 따라서 $x-a$가 인수다.
## Worked Examples
  - 예제 1: $x^2-9=x^2-3^2=(x-3)(x+3)$이다.
  - 예제 2: $2x^2+7x+3$은 $2x^2+6x+x+3=2x(x+3)+1(x+3)=(2x+1)(x+3)$로 인수분해된다.
  - 예제 3: $3x^2y-6xy^2=3xy(x-2y)$처럼 공통인수를 먼저 묶어야 한다.
## Deep Dive
  - 다항식의 핵심은 `차수`와 `계수`다. 차수는 얼마나 높은 거듭제곱이 등장하는지를, 계수는 각 항의 크기를 정한다.
  - 인수분해는 전개의 역과정이지만, 단순한 역산이 아니라 식의 구조를 분해하는 작업이다.
  - 나머지정리 관점에서는 $p(x)$를 $(x-a)$로 나눴을 때 나머지가 $p(a)$가 되므로, $p(a)=0$이면 $x-a$가 인수다.
  - 이 원리가 이차방정식의 근과 인수, 그리고 다항식의 해석 전체를 묶어 준다.
## Common Pitfalls
  - 공통인수를 먼저 묶지 않고 공식을 억지로 적용하면 계산이 꼬인다.
  - $a^2-b^2$와 $(a-b)^2$를 혼동하면 안 된다.
  - 인수분해와 방정식의 해 구하기를 같은 작업으로 생각하면 안 된다. 인수분해는 구조를 바꾸는 것이고, 해는 그 구조를 이용해 찾는 것이다.
  - 부호 실수를 많이 내는 곳이므로 마지막에 다시 전개 검산을 해야 한다.
- 교육과정 배치
  - 한국 대표 배치에서는 `중3` 핵심 단원이고, 고등학교 다항식 전반의 기반이 된다.
- 국가별 배치 스냅샷
  - 한국: `중3`에서 다항식과 인수분해가 이차방정식 직전 단계에 놓인다.
  - 일본: `중3`의 `식의 전개와 인수분해`가 대응 단원이다.
  - 중국: `8학년`의 `정식의 곱셈과 인수분해`가 대표 단원이다.
  - 미국: Algebra I와 Algebra II에서 factoring basics와 polynomial functions의 기초로 이어진다.
- 표현과 문제 감각
  - 다국어 용어: `polynomial`, `factorization`, `多項式`, `因数分解`, `因式分解`
  - 인수분해는 계산 기술이면서 동시에 `식의 모양을 알아보는 패턴 인식`이다.

## Connections

- 선수 개념은 [algebraic-manipulation.md](./algebraic-manipulation.md), [variables-and-expressions.md](./variables-and-expressions.md)다.
- 같은 축의 인접 개념으로는 [linear-equation.md](./linear-equation.md), [quadratic-equation.md](./quadratic-equation.md), [quadratic-function.md](./quadratic-function.md)가 있다.
- 다음 개념으로는 [sequences.md](./sequences.md), [complex-number-operations.md](./complex-number-operations.md), [exponential-and-logarithmic-functions.md](./exponential-and-logarithmic-functions.md)가 이어진다.
- 중3 허브와 고1 허브는 [middle-3-hub.md](../middle-3-hub.md), [high-1-hub.md](../high-1-hub.md)에서 다시 볼 수 있다.
- 과목형 연결은 [korean-algebra-course.md](./korean-algebra-course.md), [algebra-1.md](./algebra-1.md), [algebra-2.md](./algebra-2.md), [course-track-hub.md](../course-track-hub.md)로 이어진다.
- 함수 축의 연결은 [functions-strand.md](../functions-strand.md)에서 다시 묶인다.

## Open Questions

- `나머지정리`와 `다항식 나눗셈`을 별도 카드로 분리할지 후속 결정이 필요하다.

## Sources

- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-concept-encyclopedia/japan.md`
- `docs/math-concept-encyclopedia/china.md`
- `docs/math-concept-encyclopedia/us.md`
- `docs/math-concept-encyclopedia/multilingual-glossary.md`
