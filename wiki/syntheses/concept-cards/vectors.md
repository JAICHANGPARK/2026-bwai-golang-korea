---
title: 벡터
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
  - vector
---

# 벡터

## Summary

벡터는 크기와 방향을 함께 가진 양을 다루는 개념이다. 도형 정보를 계산 가능한 대수 언어로 바꾸는 강력한 도구라서 기하, 물리, 컴퓨터 그래픽스까지 폭넓게 이어진다. 이 카드에서는 평면 벡터를 중심으로 다룬다.

## Definition

- 벡터는 보통 `크기와 방향을 함께 가지는 양`으로 소개되며, 좌표 계산에서는 성분으로 나타내는 이동량으로 다룬다.
- 평면에서 벡터는 보통 성분쌍 $(a_1,a_2)$로 나타낸다.
- $0$벡터처럼 크기만 있고 방향을 정할 필요가 없는 경우도 함께 포함한다.
- 점은 위치를, 벡터는 이동과 방향을 나타낸다고 구분해서 읽는 것이 좋다.

## Key Ideas

- 성분 표현
  - $\vec a=(a_1,a_2)$
- 벡터의 길이
  - $|\vec a|=\sqrt{a_1^2+a_2^2}$
- 벡터의 덧셈
  - 같은 성분끼리 더한다.
- 스칼라배
  - 각 성분에 같은 수를 곱한다.
- 내적
  - 두 벡터의 방향 관계를 수로 읽는다.

## Theorems and Properties

- 덧셈과 스칼라배의 성분 법칙
  - $(a_1,a_2)+(b_1,b_2)=(a_1+b_1,a_2+b_2)$
  - $k(a_1,a_2)=(ka_1,ka_2)$
- 내적 공식
  - $\vec a\cdot\vec b=a_1b_1+a_2b_2$
  - 두 벡터가 모두 0벡터가 아닐 때는 $\vec a\cdot\vec b=|\vec a||\vec b|\cos\theta$로도 쓸 수 있다.
- 직교 판정
  - 두 벡터가 모두 0벡터가 아닐 때, $\vec a\cdot\vec b=0$이면 서로 수직이고, 서로 수직이면 내적은 0이다.
- 자기 자신과의 내적
  - $\vec a\cdot\vec a=|\vec a|^2$

## Proof Sketch

벡터 내적 공식은 피타고라스 정리와 코사인 법칙에서 자연스럽게 나온다. 두 벡터가 이루는 삼각형에 코사인 법칙을 적용하면

$$
|\vec a-\vec b|^2=|\vec a|^2+|\vec b|^2-2|\vec a||\vec b|\cos\theta
$$

가 된다. 한편 성분으로 전개하면 좌변은 좌표 차이의 제곱합으로도 계산된다. 이 두 표현을 비교하면 $\vec a\cdot\vec b=a_1b_1+a_2b_2$가 되고, 각이 $90^\circ$일 때는 $\cos\theta=0$이므로 내적이 0이 된다.

## Worked Examples

- 예제 1
  - $\vec a=(2,1)$, $\vec b=(1,3)$이면
  - $\vec a+\vec b=(3,4)$
  - $2\vec a=(4,2)$
- 예제 2
  - $\vec a=(2,1)$, $\vec b=(1,3)$이면
  - $\vec a\cdot\vec b=2\cdot 1+1\cdot 3=5$
  - 두 벡터의 길이가 각각 $\sqrt5$, $\sqrt{10}$이므로
    $\cos\theta=\frac{5}{\sqrt{50}}=\frac{1}{\sqrt2}$
  - 따라서 두 벡터의 끼인각은 $45^\circ$다.

## Common Pitfalls

- 점과 벡터 혼동
  - 점은 위치, 벡터는 방향이 있는 변위다.
- 성분 계산 실수
  - 덧셈과 스칼라배는 성분별로 해야 한다.
- 내적을 벡터로 착각
  - 내적의 결과는 스칼라다.
- 수직 판정 오해
  - $\vec a\cdot\vec b=0$은 두 벡터가 모두 0이 아닐 때 직교 판정으로 쓰는 것이 자연스럽다.

## Connections

- 선수 개념은 [equations-of-geometric-figures.md](./equations-of-geometric-figures.md), [pythagorean-theorem.md](./pythagorean-theorem.md), [similarity.md](./similarity.md)다.
- 다음 개념으로는 [complex-plane.md](./complex-plane.md), [spatial-coordinates.md](./spatial-coordinates.md), [conic-sections.md](./conic-sections.md), [mathematics-for-ai.md](./mathematics-for-ai.md)가 이어진다.
- 학년 허브로는 [../high-3-hub.md](../high-3-hub.md)를 함께 본다.
- 계통 허브로는 [../geometry-strand.md](../geometry-strand.md)와 [../functions-strand.md](../functions-strand.md)를 함께 본다.
- 나라·경로 허브로는 [../korea-curriculum-hub.md](../korea-curriculum-hub.md), [../japan-curriculum-hub.md](../japan-curriculum-hub.md), [../china-curriculum-hub.md](../china-curriculum-hub.md), [../us-curriculum-hub.md](../us-curriculum-hub.md), [../course-track-hub.md](../course-track-hub.md)를 본다.

## Sources

- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-concept-encyclopedia/japan.md`
- `docs/math-concept-encyclopedia/china.md`
- `docs/math-concept-encyclopedia/us.md`
- `docs/math-concept-encyclopedia/multilingual-glossary.md`
