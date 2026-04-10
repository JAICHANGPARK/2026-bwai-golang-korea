---
title: 원의 방정식
type: synthesis
status: active
updated: 2026-04-10
card_role: concept
source_docs:
  - docs/math-concept-encyclopedia/korea.md
  - docs/math-concept-encyclopedia/china.md
  - docs/math-concept-encyclopedia/multilingual-glossary.md
  - docs/math-concept-encyclopedia/unit-practice-book.md
tags:
  - concept
  - geometry
  - circle
---

# 원의 방정식

## Summary

원의 방정식은 `중심에서의 거리 일정`이라는 기하 조건을 식으로 바꾼 것이다. 원의 정의가 좌표기하에서 어떻게 계산 가능한 형태가 되는지를 보여 주는 대표 카드다.

## Key Points

- 정의
  - 중심이 $(a,b)$이고 반지름이 $r$인 원 위의 점 $(x,y)$가 만족하는 방정식을 원의 방정식이라 한다.
- 핵심 개념
  - 중심
  - 반지름
  - 거리 공식
  - 표준형
  - 전개형
- 대표 수식
  - $(x-a)^2+(y-b)^2=r^2$
  - $x^2+y^2+Dx+Ey+F=0$
- 성질 1
  - 표준형에서는 중심과 반지름을 바로 읽을 수 있다.
- 성질 2
  - 전개형은 완전제곱으로 묶어야 중심과 반지름이 드러난다.
- 증명 스케치
  - `증명 스케치 (추론)`:
  - 원 위의 모든 점은 중심에서의 거리가 $r$로 일정하다.
  - 거리 공식을 쓰면
    $$
    \sqrt{(x-a)^2+(y-b)^2}=r
    $$
    이고, 양변을 제곱하면
    $$
    (x-a)^2+(y-b)^2=r^2
    $$
    를 얻는다.

## Deep Dive

원의 방정식은 `도형의 정의 -> 거리 공식 -> 방정식`으로 이어지는 좌표기하의 전형적인 흐름을 보여 준다. 그래서 이 카드는 그 자체도 중요하지만, 이후 [parabola.md](./parabola.md), [ellipse.md](./ellipse.md), [hyperbola.md](./hyperbola.md)를 읽는 기준점 역할도 한다.

## Worked Examples

- 예제 1: 중심이 `(2,-1)`이고 반지름이 `3`인 원의 방정식은
  $$
  (x-2)^2+(y+1)^2=9
  $$
  이다.
- 예제 2:
  $$
  x^2+y^2-4x+6y-12=0
  $$
  을 완전제곱으로 정리하면
  $$
  (x-2)^2+(y+3)^2=25
  $$
  이므로 중심은 `(2,-3)`, 반지름은 `5`다.

## Common Pitfalls

- 표준형에서 중심의 부호를 반대로 읽기 쉽다.
- 전개형에서 완전제곱을 잘못 묶으면 반지름이 틀어진다.
- 반지름은 마지막에 제곱근을 취해야 하는데, $r^2$ 값을 그대로 반지름으로 읽으면 안 된다.
- 원의 정의를 잊고 공식만 외우면 응용 문제에서 흔들리기 쉽다.

## Curriculum Context

- 교육과정 배치
  - 한국 대표 배치에서는 `공통수학2`와 좌표기하의 핵심 세부 개념이다.
  - 중국 문서군에서는 중학 원 단원과 거리 개념 위에 고교 해석기하로 이어진다.
- 국가별 배치 스냅샷
  - 한국: `거리 공식`, `중점`, `원의 방정식`이 하나의 좌표기하 축을 이룬다.
  - 일본: `도형과 방정식`에서 자주 등장하는 표준 개념이다.
  - 중국: 중학교의 원 성질과 고교 해석기하 사이를 잇는다.
  - 미국: Geometry와 coordinate proof 맥락에서 반복적으로 다뤄진다.

## Connections

- 선수 개념은 [circle.md](./circle.md), [equations-of-geometric-figures.md](./equations-of-geometric-figures.md), [pythagorean-theorem.md](./pythagorean-theorem.md)다.
- 다음 개념으로는 [conic-sections.md](./conic-sections.md), [equation-of-a-line.md](./equation-of-a-line.md), [vectors.md](./vectors.md)가 이어진다.
- 계통 허브에서는 [geometry-strand.md](../geometry-strand.md), [high-1-hub.md](../high-1-hub.md)와 연결된다.

## Open Questions

- `접선의 방정식`을 원의 방정식 카드에서 더 확장할지 별도 카드로 둘지 기준이 필요하다.

## Sources

- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-concept-encyclopedia/china.md`
- `docs/math-concept-encyclopedia/multilingual-glossary.md`
- `docs/math-concept-encyclopedia/unit-practice-book.md`
