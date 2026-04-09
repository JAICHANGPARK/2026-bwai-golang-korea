---
title: 기본 도형과 작도
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
  - construction
---

# 기본 도형과 작도

## Summary

기본 도형과 작도는 점, 선, 각, 원 같은 기하의 기본 대상을 정확히 이해하고 그 조건을 그림으로 구현하는 개념이다. 이후 합동과 닮음, 증명, 좌표기하로 나아가기 위한 기하의 언어를 마련한다.

## Key Points

- 정의
  - 점, 선, 각, 다각형, 원 같은 기본 도형과 이를 정확히 그리는 작도 방법을 다루는 단원이다.
- 핵심 개념
  - 선분
  - 반직선
  - 수직
  - 평행
  - 각의 이등분선
  - 수직이등분선
- 대표 수식
  - 작도는 공식보다 `거리와 각의 조건`으로 설명하는 경우가 많다.
  - 예를 들어 수직이등분선은 `두 점에서 같은 거리인 점들의 집합`이다.
- 성질 1
  - 두 점에서 같은 거리에 있는 점들의 자취는 수직이등분선이다.
- 성질 2
  - 각의 이등분선 위의 점은 그 각의 두 변에서 같은 거리에 있다.
- 증명 스케치
  - `증명 스케치 (추론)`:
  - 두 점 $A,B$에서 같은 거리에 있는 점 $P$를 잡으면 $PA=PB$다.
  - 따라서 $P$는 선분 $AB$를 대칭적으로 보는 위치에 있고, 이런 점들의 자취가 수직이등분선이 된다.
  - 삼각형의 두 변에서 같은 거리에 있는 점들의 자취가 각의 이등분선이 되는 이유도 같은 자취 논리로 설명할 수 있다.
- 대표 예제
  - 선분의 양 끝점에서 같은 반지름의 원을 그려 만나는 두 점을 이으면 수직이등분선을 작도할 수 있다.
  - 각의 이등분선은 각의 두 변에서 같은 거리에 있는 점들의 집합으로 이해하면 작도와 증명이 함께 보인다.
- Deep Dive
  - 작도는 `그릴 수 있다`보다 `무엇이 그려지는가`를 엄밀히 따지는 활동이다.
  - 작도 문제는 결국 거리, 각, 평행, 수직 조건을 기하적으로 번역하는 연습이다.
  - 수직이등분선과 각의 이등분선은 뒤에서 합동, 닮음, 원, 좌표기하의 증명 도구가 된다.
- Worked Examples

### 예제 1: 수직이등분선 작도

- 선분 $AB$가 주어졌을 때, $A$와 $B$를 중심으로 같은 반지름의 원 두 개를 그린다.
- 두 원의 교점을 이으면 $AB$의 수직이등분선을 얻는다.

### 예제 2: 각의 이등분선 작도

- 각의 꼭짓점을 중심으로 호를 그어 두 변과 만나는 점을 잡는다.
- 그 두 점을 중심으로 같은 반지름의 호를 그려 교점을 꼭짓점과 이으면 각의 이등분선이 된다.

## Common Pitfalls

- 작도는 그림을 예쁘게 그리는 일이 아니라 조건을 만족시키는 점들의 집합을 만드는 일이다.
- 수직이등분선과 중선, 각의 이등분선과 높이를 혼동하기 쉽다.
- 원의 교점을 잇는 과정에서 호의 반지름을 다르게 잡으면 대칭성이 깨진다.
- 시각적으로 비슷해 보여도 조건이 다르면 다른 도형이다.
## Curriculum Context

- 교육과정 배치
  - 한국 대표 배치에서는 `중1`의 기본 기하 입문 단원이다.
  - 이후 합동, 원, 닮음, 좌표기하로 이어진다.
- 국가별 배치 스냅샷
  - 한국: `중1`에서 기본 도형과 작도를 함께 다룬다.
  - 일본: `중1`의 `작도와 도형의 이동`에 가깝다.
  - 중국: `7학년`의 `기하도형의 기초`가 대응 단원이다.
  - 미국: Grade 6~7의 geometry and measurement에서 기본 도형 감각이 반복 등장한다.
- 표현과 문제 감각
  - 다국어 용어: `construction`, `作図`, `作图`
  - 기하에서는 계산보다 `조건을 만족하는 점들의 집합`으로 보는 감각이 중요하다.

## Connections

- 선수 개념은 [integers-and-rational-numbers.md](./integers-and-rational-numbers.md)의 좌표와 길이 감각이다.
- 다음 개념으로는 [geometric-transformations.md](./geometric-transformations.md), [congruence.md](./congruence.md), [geometric-proof.md](./geometric-proof.md), [circle.md](./circle.md)가 이어진다.
- 계통 허브에서는 [geometry-strand.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/syntheses/geometry-strand.md)와 [korea-curriculum-hub.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/syntheses/korea-curriculum-hub.md)를 함께 본다.

## Open Questions

- 작도와 자취를 `도형의 이동` 카드와 어느 정도까지 통합할지 기준이 더 필요하다.

## Sources

- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-concept-encyclopedia/japan.md`
- `docs/math-concept-encyclopedia/china.md`
- `docs/math-concept-encyclopedia/us.md`
- `docs/math-concept-encyclopedia/multilingual-glossary.md`
