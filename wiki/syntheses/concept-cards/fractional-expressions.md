---
title: 분식
type: synthesis
status: active
updated: 2026-04-10
card_role: concept
source_docs:
  - docs/math-concept-encyclopedia/china.md
  - docs/math-curriculum-research/china.md
  - docs/math-concept-encyclopedia/unit-practice-book.md
tags:
  - concept
  - algebra
  - rational-expression
---

# 분식

## Summary

분식은 문자식이 분자와 분모에 함께 들어 있는 식을 다루는 개념이다. 계산보다 먼저 `어디에서 정의되는가`를 확인해야 한다는 점에서, 다항식보다 한 단계 더 조심스럽게 읽는 대수 카드다.

## Key Points

- 정의
  - 문자식이 분모와 분자로 이루어진 꼴의 식을 분식이라 한다.
  - 분모를 0으로 만드는 값에서는 식이 정의되지 않는다.
- 핵심 개념
  - 분자
  - 분모
  - 정의역 제한
  - 약분
  - 통분
  - 사칙연산
- 대표 수식
  - $\frac{x^2-1}{x-1}$
  - $\frac{a}{b}+\frac{c}{d}=\frac{ad+bc}{bd}\;(bd\neq 0)$
- 성질 1
  - 분모와 분자를 같은 0이 아닌 식으로 나누거나 곱하면 값은 같게 유지된다.
- 성질 2
  - 약분 전과 약분 후의 식이 완전히 같은 것은 아니고, `공통으로 정의되는 범위`에서만 같은 값을 가진다.
- 증명 스케치
  - `증명 스케치 (추론)`:
  - $\frac{ka}{kb}$에서 $k\neq 0$이면 분자와 분모를 모두 $k$로 나눌 수 있으므로 $\frac{a}{b}$와 같은 값을 가진다.
  - 다만 $k=0$이 되는 값이 있으면 애초에 그 식은 정의되지 않으므로, 분식에서는 값뿐 아니라 정의역을 함께 봐야 한다.

## Deep Dive

분식에서 가장 먼저 해야 할 일은 계산이 아니라 `분모가 0이 되는 값을 지우는 것`이다. 이 점을 놓치면 약분 결과만 보고 원래 식의 성질까지 바꾸어 읽게 된다.

예를 들어
$$
\frac{x^2-1}{x-1}=\frac{(x-1)(x+1)}{x-1}=x+1
$$
처럼 보이지만, 원래 식은 $x=1$에서 정의되지 않는다. 그래서 분식은 단순한 식 정리가 아니라 `값의 일치`와 `정의역의 차이`를 함께 보는 훈련이 된다.

## Worked Examples

- 예제 1: 
  $$
  \frac{x^2-1}{x-1}
  $$
  은
  $$
  \frac{(x-1)(x+1)}{x-1}=x+1
  $$
  로 약분할 수 있지만, 원래 식의 정의역은 $x\neq 1$이다.
- 예제 2:
  $$
  \frac{1}{x}+\frac{1}{2x}
  $$
  는 통분하면
  $$
  \frac{2}{2x}+\frac{1}{2x}=\frac{3}{2x}
  $$
  이고, $x\neq 0$이어야 한다.
- 예제 3:
  $$
  \frac{x}{x^2-4}
  $$
  는 분모를 인수분해하면
  $$
  \frac{x}{(x-2)(x+2)}
  $$
  이므로 $x\neq 2,-2$를 먼저 확인해야 한다.

## Common Pitfalls

- 약분을 했다고 해서 정의되지 않던 값까지 허용하면 안 된다.
- 분모가 다른 분식을 더하면서 분자끼리만 바로 더하면 안 된다.
- 분식 문제를 다항식 문제처럼 보고 정의역 제한을 빼먹기 쉽다.
- $x=0$ 같은 금지값을 마지막에 확인하려 하면 중간 계산이 꼬일 수 있다.

## Curriculum Context

- 교육과정 배치
  - 중국 대표 배치에서는 `8학년 1학기`의 명시적 단원이다.
  - 이후 고등학교의 유리식, 함수의 정의역, 연속성 해석으로 이어질 수 있다.
- 국가별 배치 스냅샷
  - 중국: `정식의 곱셈과 인수분해` 다음 단계에서 분식이 분명히 등장한다.
  - 한국: 현재 문서군에서는 독립 표제보다 `식의 계산`, `함수의 정의역`, `유리식` 맥락으로 분산되어 읽힌다.
  - 일본: 현재 문서군에서는 독립 표제보다 식의 변형과 함수 해석의 일부로 간접 연결된다.
  - 미국: Algebra II의 rational expressions와 함수 정의역 감각으로 이어지는 것이 자연스럽다.

## Connections

- 선수 개념은 [algebraic-manipulation.md](./algebraic-manipulation.md), [polynomials-and-factorization.md](./polynomials-and-factorization.md)다.
- 다음 개념으로는 [function-foundations.md](./function-foundations.md), [continuity.md](./continuity.md), `유리함수`가 이어진다.
- 중국 학기 허브에서는 [china-grade-8-semester-1.md](./china-grade-8-semester-1.md), [china-curriculum-hub.md](../china-curriculum-hub.md)와 연결된다.

## Open Questions

- `분식`을 중학 카드로 유지할지, 고등학교 `유리식`과 통합할지 기준이 더 필요하다.

## Sources

- `docs/math-concept-encyclopedia/china.md`
- `docs/math-curriculum-research/china.md`
- `docs/math-concept-encyclopedia/unit-practice-book.md`
