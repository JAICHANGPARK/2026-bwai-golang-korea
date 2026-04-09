---
title: 연속함수의 성질
type: synthesis
status: active
updated: 2026-04-10
card_role: concept
source_docs:
  - docs/math-concept-encyclopedia/korea.md
  - docs/math-concept-encyclopedia/japan.md
  - docs/math-concept-encyclopedia/us.md
  - docs/math-concept-encyclopedia/unit-practice-book.md
tags:
  - concept
  - calculus
  - continuity
---

# 연속함수의 성질

## Summary

연속함수의 성질은 [continuity.md](./continuity.md)의 정의를 바탕으로, 연속인 함수가 합성이나 사칙연산 아래에서 어떻게 유지되는지를 정리한 카드다. `연속이다`에서 멈추지 않고, 그 연속성이 계산과 그래프 해석에서 어떻게 작동하는지 보여 준다.

## Key Points

- 정의
  - 연속인 함수는 극한값과 함수값이 일치하는 함수이며, 여러 연산을 해도 그 성질이 보존되는 경우가 많다.
- 핵심 개념
  - 합과 차의 연속성
  - 곱의 연속성
  - 몫의 연속성
  - 합성함수의 연속성
  - 다항함수와 유리함수의 연속성
- 대표 수식
  - $\lim_{x\to a}f(x)=f(a)$
  - $(f+g)(a)=f(a)+g(a)$
  - $(fg)(a)=f(a)g(a)$
  - $\left(\frac{f}{g}\right)(a)=\frac{f(a)}{g(a)}\quad (g(a)\neq 0)$
- 핵심 해석
  - 연속성은 한 번 확인하고 끝나는 성질이 아니라, 여러 함수 조합을 안전하게 계산하게 해 주는 규칙이다.
- 교육과정 배치
  - 한국 `미적분I`, 일본 `수학III`, 미국 `Calculus`의 `limits and continuity` 확장에서 읽는 것이 자연스럽다.
- 국가별 배치 스냅샷
  - 한국: 극한과 연속 뒤의 기본 성질로 묶어 읽기 좋다.
  - 일본: `수학III`의 미분 진입 전 정리용 성질로 기능한다.
  - 중국: 도함수로 넘어가기 전 연속적 변화 감각의 배경이다.
  - 미국: Calculus에서 continuity rules로 정리된다.

## Deep Dive

### 핵심 정리 1: 연속함수의 합과 곱은 연속이다

- `증명 스케치 (추론)`:
  - $f$, $g$가 $x=a$에서 연속이면
    $$
    \lim_{x\to a}f(x)=f(a),\qquad \lim_{x\to a}g(x)=g(a)
    $$
    이다.
  - 극한의 연산 법칙에 의해
    $$
    \lim_{x\to a}(f(x)+g(x))=f(a)+g(a)
    $$
    이고,
    $$
    \lim_{x\to a}(f(x)g(x))=f(a)g(a)
    $$
    이다.
  - 따라서 합과 곱도 $x=a$에서 연속이다.

### 핵심 정리 2: 연속함수의 몫은 분모가 0이 아니면 연속이다

- `증명 스케치 (추론)`:
  - $g(a)\neq 0$이고 $g$가 연속이면, $x=a$ 근처에서 분모가 0으로 갑자기 끊어지지 않는다.
  - 극한의 몫 법칙을 적용하면
    $$
    \lim_{x\to a}\frac{f(x)}{g(x)}=\frac{f(a)}{g(a)}
    $$
    이다.
  - 따라서 분모가 0이 아닌 점에서는 몫도 연속이다.

## Worked Examples

### 예제 1: 다항함수

- 함수 $f(x)=x^2+3x+1$은 다항함수이므로 모든 실수에서 연속이다.

### 예제 2: 유리함수

- 함수
  $$
  g(x)=\frac{x+1}{x-2}
  $$
  는 $x=2$를 제외한 모든 실수에서 연속이다.
- $x=2$에서는 분모가 0이 되므로 끊김이 생긴다.

## Common Pitfalls

- `연속`과 `미분 가능`을 같은 것으로 보면 안 된다.
- 유리함수는 항상 연속이 아니라 `분모가 0이 아닌 곳에서만` 연속이다.
- 합과 곱의 연속성을 바로 쓰면서, 함수 자체가 연속인지 확인하지 않는 실수가 많다.

## Connections

- 선수 개념은 [continuity.md](./continuity.md), [limits.md](./limits.md), [function-foundations.md](./function-foundations.md)다.
- 같은 축의 인접 개념으로는 [differentiation.md](./differentiation.md), [derivative.md](./derivative.md)가 있다.
- 다음 개념으로는 [differentiation.md](./differentiation.md), [calculus-1.md](./calculus-1.md), [calculus-course.md](./calculus-course.md)가 이어진다.

## Open Questions

- 중간값정리까지 이 카드에 포함할지, 별도 심화 정리 카드로 둘지 기준이 더 필요하다.

## Sources

- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-concept-encyclopedia/japan.md`
- `docs/math-concept-encyclopedia/us.md`
- `docs/math-concept-encyclopedia/unit-practice-book.md`
