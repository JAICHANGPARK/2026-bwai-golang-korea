---
title: 연속성
type: synthesis
status: active
updated: 2026-04-10
card_role: concept
source_docs:
  - docs/math-concept-encyclopedia/korea.md
  - docs/math-concept-encyclopedia/japan.md
  - docs/math-concept-encyclopedia/us.md
  - docs/math-concept-encyclopedia/multilingual-glossary.md
  - docs/math-concept-encyclopedia/unit-practice-book.md
tags:
  - concept
  - calculus
  - continuity
---

# 연속성

## Summary

연속성은 함수 그래프가 특정 점에서 `끊기지 않고 이어지는가`를 판단하는 개념이다. 극한을 미분으로 넘겨 주는 관문이라서, 미적분에서는 별도 카드로 두는 편이 구조를 더 잘 보여 준다.

## Definition

- 함수 $f(x)$가 $x=a$에서 연속이라는 것은 다음 세 조건이 모두 성립하는 것이다.
  - $f(a)$가 정의된다.
  - $\lim_{x\to a}f(x)$가 존재한다.
  - $\lim_{x\to a}f(x)=f(a)$이다.
- 학교 수학에서는 보통 이 세 조건을 한 줄로
  $$
  \lim_{x\to a}f(x)=f(a)
  $$
  로 요약한다.

## Key Ideas

- 함수값과 극한값
  - 연속성은 `실제 함수값`과 `가까워지는 값`을 비교하는 개념이다.
- 점에서의 연속
  - 모든 점에서 연속이면 그 구간에서 그래프가 끊김 없이 이어진다.
- 불연속의 유형
  - `구멍이 있는 경우`, `점프하는 경우`, `수직 점근선처럼 발산하는 경우`를 구분하면 좋다.
- 연속성과 미분
  - 미분가능하면 연속이지만, 연속이라고 해서 항상 미분가능한 것은 아니다.

## Deep Dive

- 연속성은 그림으로 보면 `연필을 떼지 않고 그릴 수 있는가`와 비슷하지만, 계산에서는 반드시 `극한`으로 확인해야 한다.
- 특히 함수값만 보고 판단하면 안 되고, 좌극한과 우극한이 같은지도 같이 봐야 한다.

## Theorems and Properties

- 다항함수의 연속성
  - 다항함수는 모든 실수에서 연속이다.
- 유리함수의 연속성
  - 분모가 0이 아닌 점에서는 유리함수도 연속이다.
- 합성 및 사칙연산의 보존
  - 연속함수들의 합, 차, 곱은 연속이고, 나눗셈도 분모가 0이 아닌 점에서는 연속이다.
- 미분가능하면 연속
  - 한 점에서 미분가능한 함수는 그 점에서 연속이다.

## Proof Sketch

- `증명 스케치 (추론)`:
- 함수 $f(x)=2x$를 보자.
- $x$가 3에 가까워지면 $2x$는 6에 가까워지므로
  $$
  \lim_{x\to 3}2x=6
  $$
  이다.
- 동시에 실제 함수값도
  $$
  f(3)=2\cdot 3=6
  $$
  이다.
- 따라서 극한값과 함수값이 일치하므로 $f(x)=2x$는 $x=3$에서 연속이다.

## Worked Examples

- 예제 1: 연속 확인
  - $f(x)=x^2+1$은 다항함수이므로 모든 실수에서 연속이다.
  - 특히 $x=2$에서
    $$
    \lim_{x\to 2}(x^2+1)=5=f(2)
    $$
    이다.
- 예제 2: 매개변수 결정
  -
    $$
    f(x)=
    \begin{cases}
    x+2 & (x\neq 1)\\
    k & (x=1)
    \end{cases}
    $$
    가 $x=1$에서 연속이 되려면 $k$를 구하자.
  - $x\to 1$일 때 극한은 $1+2=3$이다.
  - 연속이 되려면 $f(1)=k$가 이 값과 같아야 하므로
    $$
    k=3
    $$
    이다.
- 예제 3: 불연속 판별
  - $g(x)=\frac{x^2-1}{x-1}$은 $x\neq 1$에서 $x+1$과 같지만, 원래 식은 $x=1$에서 정의되지 않는다.
  - 따라서 $x=1$에서는 `구멍이 있는 불연속`이다.

## Common Pitfalls

- 함수값이 존재한다고 해서 바로 연속이라고 결론내리면 안 된다.
- 좌극한과 우극한이 다른데도 하나의 극한이 있다고 착각하기 쉽다.
- `미분가능이면 연속`을 `연속이면 미분가능`로 거꾸로 외우면 안 된다.
- 약분해서 얻은 식만 보고 원래 함수의 정의역 문제를 놓치기 쉽다.

## Curriculum Placement

- 한국 대표 배치에서는 `미적분I`의 `함수의 극한과 연속`에 놓인다.
- 국가별 배치 스냅샷
  - 한국: `미적분I`의 첫 단원에서 극한과 함께 다룬다.
  - 일본: `수학III`에서 무한 과정과 도함수 사이의 다리로 읽힌다.
  - 미국: Calculus의 `limits and continuity`가 대표 배치다.
  - 중국: 현재 문서군에서는 도함수와 연속적 변화 감각 속에서 간접적으로 읽힌다.
- 표현 메모
  - 다국어 용어: `continuity`, `連続性`, `连续性`

## Connections

- 선수 개념은 [limits.md](./limits.md), [function-foundations.md](./function-foundations.md)다.
- 같은 축의 인접 개념으로는 [continuity-properties.md](./continuity-properties.md), [differentiation.md](./differentiation.md)가 있다.
- 다음 개념으로는 [continuity-properties.md](./continuity-properties.md), [differentiation.md](./differentiation.md), [calculus-course.md](./calculus-course.md)가 이어진다.
- 과목 허브로는 [calculus-1.md](./calculus-1.md), [calculus-course.md](./calculus-course.md)를 본다.
- 계통 허브로는 [../functions-strand.md](../functions-strand.md), [../high-2-hub.md](../high-2-hub.md)를 함께 본다.

## Open Questions

- `중간값정리`와 `최대최소정리`를 별도 심화 카드로 더 뺄지 후속 설계가 필요하다.

## Sources

- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-concept-encyclopedia/japan.md`
- `docs/math-concept-encyclopedia/us.md`
- `docs/math-concept-encyclopedia/multilingual-glossary.md`
- `docs/math-concept-encyclopedia/unit-practice-book.md`
