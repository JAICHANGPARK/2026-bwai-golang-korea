---
title: 쌍곡선
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

# 쌍곡선

## Summary

쌍곡선은 두 초점까지의 거리의 차의 절댓값이 일정한 점들의 자취다. 두 갈래로 벌어지는 곡선이라는 점에서 타원과 대비되고, 점근선 개념까지 함께 읽히는 대표 이차곡선이다.

## Key Points

- 정의
  - 두 초점까지의 거리의 차의 절댓값이 일정한 점들의 집합을 쌍곡선이라 한다.
- 핵심 개념
  - 두 초점
  - 중심
  - 두 갈래 곡선
  - 점근선
  - 거리차 조건
- 대표 수식
  - $\frac{x^2}{a^2}-\frac{y^2}{b^2}=1$
- 성질 1
  - 중심에 대하여 대칭이다.
- 성질 2
  - 그래프는 두 갈래로 나뉘며, 멀리서 보면 점근선에 가까워진다.
- 증명 스케치
  - `증명 스케치 (추론)`:
  - 거리차 조건을 좌표로 옮기면 제곱과 정리를 거쳐 표준형을 얻는다.
  - 학교 수준에서는 계산보다 `타원은 거리합`, `쌍곡선은 거리차`라는 대비를 먼저 잡는 것이 중요하다.

## Deep Dive

쌍곡선은 타원과 초점 구조는 비슷하지만 결과 모양이 완전히 다르다. 타원은 닫혀 있고, 쌍곡선은 두 갈래로 열려 있다. 이 차이는 거리의 `합`을 고정하느냐 `차`를 고정하느냐에서 나온다.

또한 쌍곡선은 점근선 개념을 통해 `곡선이 직선에 점점 가까워지는 방식`을 보여 준다. 이는 이후 함수 그래프 해석과도 이어지는 중요한 감각이다.

## Worked Examples

- 예제 1:
  $$
  \frac{x^2}{9}-\frac{y^2}{4}=1
  $$
  은 중심이 원점인 쌍곡선이고, x축 방향으로 열린다.
- 예제 2: 위 식의 점근선은
  $$
  y=\pm \frac{2}{3}x
  $$
  이다.

## Common Pitfalls

- 타원과 쌍곡선의 표준형 부호 차이를 자주 헷갈린다.
- 두 갈래라는 그림만 기억하고 거리차 조건을 놓치기 쉽다.
- 점근선이 그래프 위의 선이라고 오해하면 안 된다.
- 분모가 큰 쪽이 장축이라고 기계적으로 읽으면 쌍곡선에서는 해석이 어긋난다.

## Curriculum Context

- 교육과정 배치
  - 한국 대표 배치에서는 `기하`의 이차곡선 세부 개념이다.
  - 중국 해석기하와 미국 Precalculus/coordinate geometry의 확장 카드로도 자연스럽다.
- 국가별 배치 스냅샷
  - 한국: `기하`에서 포물선·타원·쌍곡선이 함께 조직된다.
  - 일본: 고교 좌표기하 심화 곡선으로 간접 연결된다.
  - 중국: `공간벡터와 해석기하`에서 원뿔곡선 일부로 읽힌다.
  - 미국: Algebra II, Precalculus에서 conic sections 일부로 분산 배치된다.

## Connections

- 선수 개념은 [conic-sections.md](./conic-sections.md), [ellipse.md](./ellipse.md), [equations-of-geometric-figures.md](./equations-of-geometric-figures.md)다.
- 다음 개념으로는 [vectors.md](./vectors.md), [spatial-coordinates.md](./spatial-coordinates.md), `해석기하`가 이어진다.
- 계통 허브에서는 [geometry-strand.md](../geometry-strand.md), [high-2-hub.md](../high-2-hub.md)와 연결된다.

## Open Questions

- 점근선과 함수 그래프의 점근선을 별도 비교 카드로 둘지 기준이 필요하다.

## Sources

- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-concept-encyclopedia/china.md`
- `docs/math-concept-encyclopedia/unit-practice-book.md`
