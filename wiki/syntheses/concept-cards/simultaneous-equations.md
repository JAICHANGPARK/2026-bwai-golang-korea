---
title: 연립방정식
type: synthesis
status: active
updated: 2026-04-10
card_role: concept
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
  - docs/math-concept-encyclopedia/comparative-problem-book.md
tags:
  - concept
  - algebra
  - equation
---

# 연립방정식

## Summary

연립방정식은 여러 방정식을 동시에 만족하는 해를 찾는 개념이다. 식을 푸는 대수 조작과 그래프의 교점을 읽는 함수 해석이 만나는 첫 핵심 브리지라서, 중학교 대수에서 고등학교 행렬과 선형대수로 이어지는 중요한 관문이 된다.

## Key Points

- 정의
  - 두 개 이상의 방정식을 동시에 만족하는 해를 구하는 문제를 연립방정식이라 한다.
- 핵심 개념
  - 소거법
  - 대입법
  - 해의 존재와 개수
  - 교점과 해의 관계
  - 식과 그래프의 대응
- 대표 수식
  - $\begin{cases}a_1x+b_1y=c_1\\a_2x+b_2y=c_2\end{cases}$
- 증명 스케치
  - `증명 스케치 (추론)`:
  - 점 $(x,y)$가 두 직선의 그래프 위에 동시에 있으면 첫째 식과 둘째 식을 모두 만족한다.
  - 반대로 $(x,y)$가 두 식을 동시에 만족하면, 그 점은 각 식이 나타내는 두 직선 위에 모두 놓인다.
  - 따라서 연립일차방정식의 해는 두 직선의 교점으로 읽을 수 있다.
## Worked Examples
  - 예제 1: $\begin{cases}x+y=7\\x-y=1\end{cases}$에서 두 식을 더하면 $2x=8$이므로 $x=4$이다. 이를 첫째 식에 대입하면 $y=3$이므로 해는 $(4,3)$이다.
  - 예제 2: $\begin{cases}2x+3y=12\\x-y=1\end{cases}$는 대입법으로 풀면 $x=3$, $y=2$가 된다.
## Deep Dive
  - 연립방정식은 `두 조건을 동시에 만족하는 공통 해`를 찾는 문제다.
  - 두 일차식은 그래프에서 두 직선을 이루므로, 해의 개수는 보통 교점의 개수와 같다.
  - 교점이 하나면 해가 하나, 평행하면 해가 없고, 같은 직선이면 해가 무한히 많다.
  - 소거법과 대입법은 이 교점 문제를 식의 계산으로 옮겨 푸는 두 가지 방식이다.
## Common Pitfalls
  - 한 식만 만족하는 값을 해로 끝내면 안 된다.
  - 소거 과정에서 부호를 잘못 바꾸면 교점이 달라진다.
  - 두 직선이 평행하거나 일치하는 경우를 반드시 확인해야 한다.
  - 연립방정식의 해는 보통 순서쌍으로 쓴다.
- 교육과정 배치
  - 한국 대표 배치에서는 `중2`의 핵심 선형 관계 단원이다.
  - 이후 `일차함수의 교점`, `행렬`, `선형대수의 기초 발상`으로 이어진다.
- 국가별 배치 스냅샷
  - 한국: `중2`에서 `일차함수`와 나란히 놓이며 선형 관계를 식과 그래프로 함께 읽게 한다.
  - 일본: `중2`의 `연립일차방정식`으로 도입되고, 설명형 해석이 붙는 경우가 잦다.
  - 중국: `7학년 2학기`에 비교적 이른 시기에 도입된다.
  - 미국: 보통 `Grade 8`의 `Linear Equations and Systems` 또는 `Algebra I` 경로에서 다룬다.
- 표현과 문제 감각
  - 다국어 용어: `system of equations`, `連立方程式`, `方程组 / 二元一次方程组`
  - 비교 문제집 기준으로 한국·중국은 계산 절차 훈련 비중이 크고, 일본은 교점 해석 설명이 붙기 쉬우며, 미국은 그래프와 식을 함께 읽는 문제가 자주 나온다.

## Connections

- 선수 개념은 [linear-equation.md](./linear-equation.md)와 `좌표평면`이다.
- 같이 읽을 개념은 [linear-function.md](./linear-function.md)다.
- 다음 개념으로는 [matrix.md](./matrix.md), [equations-of-geometric-figures.md](./equations-of-geometric-figures.md), `선형대수`가 이어진다.
- 학년 허브에서는 [middle-2-hub.md](../middle-2-hub.md)와 연결된다.
- 함수 계통은 [functions-strand.md](../functions-strand.md)에서 함께 본다.

## Open Questions

- `연립방정식` 카드 안에 `해의 개수`와 `평행/일치/교차`를 별도 시각화 카드로 둘지 기준이 필요하다.

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
- `docs/math-concept-encyclopedia/comparative-problem-book.md`
