---
title: 타원
type: synthesis
status: active
updated: 2026-04-10
card_role: concept
source_docs:
  - docs/math-concept-encyclopedia/korea.md
  - docs/math-concept-encyclopedia/china.md
  - docs/math-concept-encyclopedia/unit-practice-book.md
tags:
  - concept
  - geometry
  - conic
---

# 타원

## Summary

타원은 두 초점까지의 거리의 합이 일정한 점들의 자취다. 이차곡선 가운데 닫힌 곡선이라는 점이 중요하고, 궤도와 반사 성질 같은 응용으로도 자주 연결된다.

## Key Points

- 정의
  - 두 초점까지의 거리의 합이 일정한 점들의 집합을 타원이라 한다.
- 핵심 개념
  - 두 초점
  - 장축
  - 단축
  - 중심
  - 닫힌 곡선
- 대표 수식
  - $\frac{x^2}{a^2}+\frac{y^2}{b^2}=1\;(a>b>0)$
- 성질 1
  - 중심에 대하여 대칭이다.
- 성질 2
  - 장축 방향으로 더 길게 늘어난 닫힌 곡선이다.
- 증명 스케치
  - `증명 스케치 (추론)`:
  - 두 초점까지의 거리의 합이 일정하다는 조건을 좌표로 두면, 제곱과 정리를 거쳐 표준형으로 옮길 수 있다.
  - 학교 수준에서는 계산 자체보다 `거리합 조건`이 타원의 본질이라는 점이 더 중요하다.

## Deep Dive

타원은 [parabola.md](./parabola.md)와 달리 열린 곡선이 아니라 닫힌 곡선이다. 그래서 같은 이차곡선이어도 운동과 경계의 해석이 달라진다.

또한 타원은 초점 두 개가 있다는 점에서 원과도 다르다. 원은 중심 하나에서의 거리 일정, 타원은 초점 두 개까지의 거리합 일정이라는 차이를 잡아 두면 해석기하 구조가 또렷해진다.

## Worked Examples

- 예제 1:
  $$
  \frac{x^2}{9}+\frac{y^2}{4}=1
  $$
  은 중심이 원점이고 장반경이 `3`, 단반경이 `2`인 타원이다.
- 예제 2: 장반경이 x축 방향에 있으면 더 길게 누운 모양의 타원이 되고, y축 방향에 있으면 세로로 선 타원이 된다.

## Common Pitfalls

- 타원의 초점과 꼭짓점을 같은 것으로 보면 안 된다.
- 분모가 큰 쪽이 더 긴 축이라는 점을 자주 반대로 읽는다.
- 타원을 원이 눌린 그림 정도로만 보면 초점 구조를 놓치게 된다.
- 표준형을 외우더라도 거리합 조건을 잊으면 응용 문제에서 흔들린다.

## Curriculum Context

- 교육과정 배치
  - 한국 대표 배치에서는 `기하`의 이차곡선 세부 개념이다.
  - 중국 해석기하와 미국 Precalculus/coordinate geometry에서도 자연스럽게 이어진다.
- 국가별 배치 스냅샷
  - 한국: `기하`에서 포물선·타원·쌍곡선을 함께 묶는다.
  - 일본: 고교 곡선 해석과 연결되는 심화 좌표기하 개념이다.
  - 중국: `공간벡터와 해석기하` 축의 일부로 읽힌다.
  - 미국: Algebra II, Precalculus에서 conic sections의 일부로 분산 배치된다.

## Connections

- 선수 개념은 [conic-sections.md](./conic-sections.md), [equations-of-geometric-figures.md](./equations-of-geometric-figures.md)다.
- 다음 개념으로는 [hyperbola.md](./hyperbola.md), [vectors.md](./vectors.md), [spatial-coordinates.md](./spatial-coordinates.md)가 이어진다.
- 계통 허브에서는 [geometry-strand.md](../geometry-strand.md), [high-2-hub.md](../high-2-hub.md)와 연결된다.

## Open Questions

- 초점 거리 계산 예제를 별도 카드로 더 확장할지 기준이 필요하다.

## Sources

- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-concept-encyclopedia/china.md`
- `docs/math-concept-encyclopedia/unit-practice-book.md`
