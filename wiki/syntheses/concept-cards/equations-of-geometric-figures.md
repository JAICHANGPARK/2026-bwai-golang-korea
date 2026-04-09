---
title: 도형의 방정식
type: synthesis
status: active
updated: 2026-04-09
source_docs:
  - docs/math-concept-encyclopedia/korea.md
  - docs/math-concept-encyclopedia/japan.md
  - docs/math-concept-encyclopedia/china.md
  - docs/math-concept-encyclopedia/us.md
  - docs/math-curriculum-research/korea.md
  - docs/math-curriculum-research/japan.md
  - docs/math-curriculum-research/china.md
  - docs/math-curriculum-research/us.md
  - docs/math-concept-encyclopedia/multilingual-glossary.md
  - docs/math-concept-encyclopedia/unit-practice-book.md
tags:
  - concept
  - geometry
  - coordinate
---

# 도형의 방정식

## Summary

도형의 방정식은 좌표평면 위의 도형을 식으로 나타내고, 식을 다시 도형으로 읽는 개념이다. 기하와 대수를 직접 연결하는 언어라서, `그림을 계산 가능한 구조로 바꾸는 힘`을 길러 주는 핵심 브리지다.

## Key Points

- 정의
  - 점, 직선, 원, 포물선 같은 도형의 성질을 좌표와 방정식으로 표현하는 단원을 도형의 방정식이라 한다.
- 핵심 개념
  - 좌표평면
  - 거리 공식
  - 중점
  - 직선의 방정식
  - 원의 방정식
  - 자취
- 대표 수식
  - $d=\sqrt{(x_2-x_1)^2+(y_2-y_1)^2}$
  - $M\left(\frac{x_1+x_2}{2},\frac{y_1+y_2}{2}\right)$
  - $(x-a)^2+(y-b)^2=r^2$
- 증명 스케치
  - `증명 스케치 (추론)`:
  - 중심이 $(a,b)$이고 반지름이 $r$인 원 위의 점 $(x,y)$는 중심에서의 거리가 항상 $r$이다.
  - 거리 공식을 쓰면
    $$
    \sqrt{(x-a)^2+(y-b)^2}=r
    $$
    이고, 양변을 제곱하면
    $$
    (x-a)^2+(y-b)^2=r^2
    $$
    를 얻는다.
  - 이 과정은 `거리 조건을 방정식으로 바꾼다`는 좌표기하의 핵심 발상을 보여 준다.
- 대표 예제
  - 중심이 `(2,-1)`이고 반지름이 `3`인 원의 방정식은
    $$
    (x-2)^2+(y+1)^2=9
    $$
    이다.
- 교육과정 배치
  - 한국 대표 배치에서는 고등학교 `공통수학2`의 핵심 구성요소다.
  - 이후 `기하`, `벡터`, `이차곡선`, `해석기하`로 이어진다.
- 국가별 배치 스냅샷
  - 한국: `공통수학2`에서 거리 공식, 중점, 원의 방정식과 함께 좌표기하의 기초로 다룬다.
  - 일본: 고등학교 `수학II`의 `도형と方程式`에서 명시적으로 다룬다.
  - 중국: 중학교의 좌표와 원 개념 위에 고등학교 `평면해석기하`, `원뿔곡선과 방정식` 축으로 확장된다.
  - 미국: 보통 `Geometry` 과목의 `coordinate geometry`에서 증명과 계산을 함께 배운다.
- 표현과 문제 감각
  - 다국어 용어: `equations of geometric figures`, `coordinate geometry`, `図形と方程式`, `解析几何`
  - 한국과 일본은 거리·중점·원의 방정식이 선명하고, 중국은 원뿔곡선 확장이 빠르며, 미국은 `coordinate proof`와 함께 다루는 경우가 많다.

## Connections

- 선수 개념은 `좌표평면`, [linear-function.md](./linear-function.md), [quadratic-function.md](./quadratic-function.md)이다.
- 다음 개념으로는 `기하`, `벡터`, `이차곡선`, `해석기하`가 이어진다.
- 학년 허브에서는 [high-1-hub.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/queries/math-curriculum-graph/high-1-hub.md), [high-2-hub.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/queries/math-curriculum-graph/high-2-hub.md)와 연결된다.
- 계통 허브에서는 [geometry-strand.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/queries/math-curriculum-graph/geometry-strand.md)와 [functions-strand.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/queries/math-curriculum-graph/functions-strand.md)를 함께 본다.

## Open Questions

- `직선의 방정식`, `원의 방정식`, `이차곡선`을 한 카드에 유지할지 더 잘게 분리할지 기준이 필요하다.
- `도형의 방정식` 카드와 `기하` 과목 카드를 어디서 나눌지 후속 설계가 필요하다.

## Sources

- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-concept-encyclopedia/japan.md`
- `docs/math-concept-encyclopedia/china.md`
- `docs/math-concept-encyclopedia/us.md`
- `docs/math-curriculum-research/korea.md`
- `docs/math-curriculum-research/japan.md`
- `docs/math-curriculum-research/china.md`
- `docs/math-curriculum-research/us.md`
- `docs/math-concept-encyclopedia/multilingual-glossary.md`
- `docs/math-concept-encyclopedia/unit-practice-book.md`
