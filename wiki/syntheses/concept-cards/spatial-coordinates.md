---
title: 공간좌표
type: synthesis
status: active
updated: 2026-04-10
card_role: concept
source_docs:
  - docs/math-concept-encyclopedia/korea.md
  - docs/math-concept-encyclopedia/china.md
  - docs/math-concept-encyclopedia/japan.md
  - docs/math-concept-encyclopedia/multilingual-glossary.md
tags:
  - concept
  - geometry
  - spatial
---

# 공간좌표

## Summary

공간좌표는 3차원 공간의 점과 도형을 좌표로 표현하는 개념이다. 평면좌표의 생각을 입체로 확장해 직선, 평면, 거리, 위치 관계를 계산하게 해 준다.

## Key Points

- 정의
  - 공간 속 한 점의 위치를 세 수 $(x,y,z)$로 나타내는 체계를 공간좌표라 한다.
- 핵심 개념
  - 3차원 좌표축
  - 점의 위치
  - 거리
  - 직선과 평면
  - 투영
- 대표 수식
  - $P=(x,y,z)$
  - 두 점 사이 거리 $d=\sqrt{(x_2-x_1)^2+(y_2-y_1)^2+(z_2-z_1)^2}$
- 성질 1
  - 3차원 거리 공식은 평면의 피타고라스 정리를 두 번 적용해 얻을 수 있다.
- 성질 2
  - 중점은 각 좌표를 각각 평균한 점이다.
- 성질 3
  - 구는 중심에서의 거리가 일정한 점들의 집합이다.
- 증명 스케치
  - `증명 스케치 (추론)`:
  - 공간거리 공식은 평면의 피타고라스 정리를 두 번 적용해 얻을 수 있다.
  - 먼저 바닥면에서의 거리, 그다음 높이 차이를 결합하면 3차원 거리 공식을 얻는다.
  - 중점 공식은 x, y, z좌표 각각에서 `반씩 나눈다`는 사실로 바로 확인된다.
- 대표 예제
  - $(0,0,0)$과 $(1,2,2)$ 사이 거리는 $\sqrt{1+4+4}=3$이다.
  - $(1,2,3)$과 $(5,0,1)$의 중점은 $(3,1,2)$이다.
  - 중심이 $(1,-1,2)$이고 반지름이 `4`인 구는
    $$
    (x-1)^2+(y+1)^2+(z-2)^2=16
    $$
    로 쓸 수 있다.
- Deep Dive
  - 공간좌표는 2차원 좌표의 연장선이지만, 시각화보다 좌표 계산이 훨씬 중요하다.
  - 직선과 평면의 위치 관계, 거리, 투영은 모두 이 좌표 체계 위에서 계산된다.
  - 벡터와 연결하면 공간에서의 방향과 길이를 함께 다루기 쉬워진다.
- Worked Examples

### 예제 1: 거리 구하기

- $(1,1,1)$과 $(4,5,1)$ 사이 거리는
  $$
  \sqrt{(4-1)^2+(5-1)^2+(1-1)^2}=5
  $$
  이다.

### 예제 2: 중점 구하기

- $(2,-2,4)$와 $(6,0,0)$의 중점은
  $$
  (4,-1,2)
  $$
  이다.

## Common Pitfalls

- 3차원 거리 공식에서 세 좌표 차를 빠뜨리면 안 된다.
- 중점은 각 좌표를 평균 내는 점이지, 거리의 중간값이 아니다.
- 공간좌표의 그림은 보조도구이고, 실제 계산은 좌표식으로 해야 한다.
- 구와 원을 혼동하면 차원을 잘못 읽기 쉽다.
- ## Curriculum Context
- 교육과정 배치
  - 한국 대표 배치에서는 `기하`에서 공간도형과 함께 다룬다.
  - 이후 공간벡터, 물리의 운동, 컴퓨터 그래픽스로 이어진다.
- 국가별 배치 스냅샷
  - 한국: `기하`의 `공간도형과 공간좌표` 축에 놓인다.
  - 일본: `수학C`의 벡터·곡선 관련 내용과 연결된다.
  - 중국: `공간벡터와 해석기하`에서 핵심적으로 다뤄진다.
  - 미국: 일반 K-12에서는 제한적이지만 advanced geometry와 calculus 전단계에서 나타난다.
- 표현과 문제 감각
  - 다국어 용어: `three-dimensional coordinates`, `空間座標`, `空间直角坐标系`
  - 2차원 그림으로 3차원 구조를 상상하는 시각화 감각이 중요하다.

## Connections

- 선수 개념은 [vectors.md](./vectors.md), [equations-of-geometric-figures.md](./equations-of-geometric-figures.md), [pythagorean-theorem.md](./pythagorean-theorem.md)다.
- 다음 개념으로는 `공간벡터`, `물리의 운동`, `컴퓨터 그래픽스`, [conic-sections.md](./conic-sections.md)가 이어진다.
- 계통 허브에서는 [geometry-strand.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/syntheses/geometry-strand.md)와 [korea-curriculum-hub.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/syntheses/korea-curriculum-hub.md)를 함께 본다.

## Open Questions

- `직선과 평면의 위치관계`를 별도 카드로 분리할지 검토가 필요하다.

## Sources

- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-concept-encyclopedia/china.md`
- `docs/math-concept-encyclopedia/japan.md`
- `docs/math-concept-encyclopedia/multilingual-glossary.md`
