---
title: 원
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
  - geometry
  - circle
---

# 원

## Summary

원은 한 점에서 같은 거리에 있는 점들의 집합이라는 기하 개념이다. 거리, 대칭, 각의 관계를 한 번에 묶는 도형이라서 원주각, 접선, 좌표기하, 삼각비로 자연스럽게 이어진다.

## Key Points

- 정의
  - 한 점에서 같은 거리에 있는 점들의 모임을 원이라 한다.
- 핵심 개념
  - 중심
  - 반지름
  - 지름
  - 현
  - 중심각
  - 원주각
- 대표 수식
  - 원의 방정식은 $(x-a)^2+(y-b)^2=r^2$다.
- 성질 1
  - 같은 원에서 같은 중심각은 같은 호와 같은 현을 만든다.
- 성질 2
  - 같은 호를 보는 원주각은 서로 같고, 그 크기는 중심각의 절반이다.
- 성질 3
  - 접선은 접점에서 반지름과 수직이다.
- 증명 스케치
  - `증명 스케치 (추론)`:
  - 같은 호를 보는 원주각은 같은 중심각의 절반이 된다.
  - 이는 중심과 호를 잇는 삼각형들의 이등변 구조를 이용해 각의 합을 비교하면 보일 수 있다.
  - 접선과 반지름의 수직성은 접선이 원과 한 점에서만 만난다는 사실을 자취 관점으로 읽으면 자연스럽다.
- 대표 예제
  - 중심각이 $80^\circ$이면 같은 호를 보는 원주각은 $40^\circ$다.
  - 반지름이 `5`인 원에서 중심에서의 거리가 `5`인 점은 모두 원 위에 있다.
  - 접점에서의 접선은 그 점으로 그은 반지름과 직각을 이룬다.
- Deep Dive
  - 원은 대칭성이 큰 도형이라서 각, 길이, 접선 성질이 서로 강하게 연결된다.
  - 원의 성질은 `호-현-각`의 관계로 정리하면 기억하기 쉽다.
  - 좌표기하에서는 원의 방정식이 `중심에서의 거리 일정`이라는 정의를 식으로 바꾼 것이다.
- Worked Examples

### 예제 1: 원주각

- 중심각이 $120^\circ$이면 같은 호를 보는 원주각은
  $$
  60^\circ
  $$
  이다.

### 예제 2: 원의 방정식

- 중심이 $(2,-1)$이고 반지름이 `3`인 원은
  $$
  (x-2)^2+(y+1)^2=9
  $$
  이다.

## Common Pitfalls

- 원주각은 항상 중심각의 절반이지만, 같은 호를 보고 있어야 한다.
- 접선과 현을 혼동하면 원의 성질을 잘못 적용하기 쉽다.
- 원의 방정식에서 중심의 부호를 반대로 읽기 쉽다.
- 중심과 반지름을 전개식에서 다시 찾지 못하면 좌표기하 문제가 꼬인다.
## Curriculum Context

- 교육과정 배치
  - 한국 대표 배치에서는 `중3`의 닮음·원·삼각비 축에서 다뤄진다.
  - 이후 원주각, 접선, 좌표기하로 이어진다.
- 국가별 배치 스냅샷
  - 한국: `중3`에서 원주각과 중심각, 삼각비 전단계로 이어진다.
  - 일본: `중3`의 `원주각`과 고등학교 `수학A`의 정다각형·원 성질로 이어진다.
  - 중국: `9학년 1학기`의 `원`이 명시적 단원이다.
  - 미국: Grade 7 Geometry와 고등학교 Geometry에서 circles가 반복 등장한다.
- 표현과 문제 감각
  - 다국어 용어: `circle`, `円`, `圆`
  - `집합으로서의 원`과 `원의 성질 문제`를 함께 읽는 것이 중요하다.

## Connections

- 선수 개념은 [basic-geometry-and-construction.md](./basic-geometry-and-construction.md), [congruence.md](./congruence.md)다.
- 다음 개념으로는 [similarity.md](./similarity.md), [trigonometric-ratio.md](./trigonometric-ratio.md), [equation-of-a-circle.md](./equation-of-a-circle.md), [equations-of-geometric-figures.md](./equations-of-geometric-figures.md), [trigonometric-function.md](./trigonometric-function.md)가 이어진다.
- 계통 허브에서는 [geometry-strand.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/syntheses/geometry-strand.md)와 [korea-curriculum-hub.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/syntheses/korea-curriculum-hub.md)를 함께 본다.

## Open Questions

- `원주각`과 `접선`을 각각 독립 카드로 나눌지 후속 판단이 필요하다.

## Sources

- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-concept-encyclopedia/japan.md`
- `docs/math-concept-encyclopedia/china.md`
- `docs/math-concept-encyclopedia/us.md`
- `docs/math-concept-encyclopedia/multilingual-glossary.md`
