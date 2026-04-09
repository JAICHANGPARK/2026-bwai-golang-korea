---
title: 이차곡선
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
  - conic
---

# 이차곡선

## Summary

이차곡선은 포물선, 타원, 쌍곡선을 하나의 곡선군으로 묶어 읽는 개념이다. 거리 조건과 좌표식이 만나는 대표 영역이라서 해석기하와 물리, 그래픽스의 핵심 언어가 된다.

## Key Points

- 정의
  - 평면을 원뿔로 잘랐을 때 생기는 곡선들을 이차곡선이라 한다.
- 핵심 개념
  - 포물선
  - 타원
  - 쌍곡선
  - 초점
  - 거리 조건
  - 좌표 방정식
- 대표 수식
  - $y^2=4px$
  - $\frac{x^2}{a^2}+\frac{y^2}{b^2}=1$
- 성질 1
  - 포물선은 초점과 준선까지의 거리가 같은 점들의 자취다.
- 성질 2
  - 타원은 두 초점까지의 거리의 합이 일정한 점들의 자취다.
- 성질 3
  - 쌍곡선은 두 초점까지의 거리의 차의 절댓값이 일정한 점들의 자취다.
- 증명 스케치
  - `증명 스케치 (추론)`:
  - 포물선은 한 점(초점)과 한 직선(준선)에서의 거리가 같은 점들의 집합으로 정의할 수 있다.
  - 이 거리 조건을 좌표로 풀어 쓰면 포물선의 방정식을 얻는다.
  - 타원과 쌍곡선도 같은 방식으로 `거리 조건 -> 방정식`으로 옮겨 쓸 수 있다.
- 대표 예제
  - $y^2=4x$는 $p=1$인 포물선이다.
  - 두 초점이 있는 거리합 조건을 좌표화하면 타원의 표준형을 얻는다.
  - 거리차 조건을 좌표화하면 쌍곡선의 표준형을 얻는다.
- Deep Dive
  - 이차곡선은 하나의 원뿔 절단 그림이면서 동시에 거리 조건으로 정의되는 자취다.
  - 학교 수준에서는 각 곡선의 표준형보다 `어떤 거리 조건이 어떤 곡선을 만드는가`가 더 중요하다.
  - 포물선은 광학 반사, 타원은 궤도, 쌍곡선은 위치 판정 문제와 연결된다.
- Worked Examples

### 예제 1: 포물선

- $y^2=8x$는
  $$
  y^2=4px
  $$
  와 비교하면 $p=2$인 포물선이다.

### 예제 2: 타원

- $\frac{x^2}{9}+\frac{y^2}{4}=1$은 중심이 원점이고 장반경이 `3`, 단반경이 `2`인 타원이다.

## Common Pitfalls

- 이차곡선을 모두 같은 식으로 외우면 곡선의 차이를 놓치기 쉽다.
- 포물선의 준선과 초점을 뒤바꾸면 안 된다.
- 타원과 쌍곡선의 표준형에서 분모의 위치를 헷갈리기 쉽다.
- 원뿔 절단 설명만 기억하고 거리 조건을 놓치면 해석기하와의 연결이 약해진다.
- ## Curriculum Context
- 교육과정 배치
  - 한국 대표 배치에서는 `기하`의 핵심 구성요소다.
  - 이후 벡터, 공간좌표, 해석기하로 이어진다.
- 국가별 배치 스냅샷
  - 한국: `기하`에서 포물선·타원·쌍곡선을 함께 다룬다.
  - 일본: `수학C`의 곡선 관련 내용과 연결된다.
  - 중국: `공간벡터와 해석기하` 및 `원뿔곡선` 축에서 다룬다.
  - 미국: Algebra II, Precalculus, coordinate geometry 안에서 분산 배치되는 경우가 많다.
- 표현과 문제 감각
  - 다국어 용어: `conic sections`, `二次曲線`, `圆锥曲线`
  - 식 암기보다 `어떤 거리 조건이 어떤 곡선을 만드는가`가 중요하다.

## Connections

- 선수 개념은 [equations-of-geometric-figures.md](./equations-of-geometric-figures.md), [quadratic-function.md](./quadratic-function.md)다.
- 다음 개념으로는 [vectors.md](./vectors.md), [spatial-coordinates.md](./spatial-coordinates.md), `해석기하`가 이어진다.
- 계통 허브에서는 [geometry-strand.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/syntheses/geometry-strand.md)와 [korea-curriculum-hub.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/syntheses/korea-curriculum-hub.md)를 함께 본다.

## Open Questions

- 포물선·타원·쌍곡선을 각각 독립 카드로 더 세분화할지 검토가 필요하다.

## Sources

- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-concept-encyclopedia/japan.md`
- `docs/math-concept-encyclopedia/china.md`
- `docs/math-concept-encyclopedia/us.md`
- `docs/math-concept-encyclopedia/multilingual-glossary.md`
