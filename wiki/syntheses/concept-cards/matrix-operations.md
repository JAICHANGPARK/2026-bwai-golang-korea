---
title: 행렬의 연산
type: synthesis
status: active
updated: 2026-04-10
card_role: concept
source_docs:
  - docs/math-concept-encyclopedia/korea.md
  - docs/math-curriculum-research/korea.md
  - docs/math-concept-encyclopedia/multilingual-glossary.md
  - docs/math-concept-encyclopedia/unit-practice-book.md
tags:
  - concept
  - algebra
  - matrix
---

# 행렬의 연산

## Summary

행렬의 연산은 배열로 표현된 정보를 `같은 자리끼리 더하고`, `행과 열을 맞물려 곱하는` 규칙으로 다루는 방법이다. [matrix.md](./matrix.md)가 행렬의 모양과 표현을 소개한다면, 이 카드는 그 표현이 실제 계산 대상으로 어떻게 움직이는지 보여 준다.

## Key Points

- 정의
  - 같은 크기의 행렬은 대응하는 성분끼리 더하고, 행렬곱은 왼쪽 행렬의 행과 오른쪽 행렬의 열을 맞춰 계산한다.
- 핵심 개념
  - 같은 크기
  - 대응 성분
  - 스칼라배
  - 행렬곱
  - 연립관계의 압축 표현
- 대표 수식
  - $A+B=(a_{ij}+b_{ij})$
  - $kA=(ka_{ij})$
  - $(AB)_{ij}=\sum_k a_{ik}b_{kj}$
- 핵심 해석
  - 덧셈은 `같은 위치의 정보 결합`이다.
  - 곱셈은 `행의 정보`와 `열의 정보`를 한 번에 합성하는 연산이다.
- 교육과정 배치
  - 한국 대표 배치에서는 `공통수학1`의 행렬 표현을 바탕으로 이후 심화 대수 감각으로 확장될 수 있다.
  - 현재 문서군에서는 행렬의 독립 계산 단원보다 `표현 -> 연산으로의 확장 가능성`을 보여 주는 세부 하위 카드로 두는 것이 적절하다.
- 국가별 배치 스냅샷
  - 한국: `공통수학1`의 행렬 표현에서 출발해 다음 대수 감각으로 이어질 수 있다.
  - 일본/중국/미국: 현재 문서군에서는 독립 기초 단원보다 상위 대수 또는 선형대수 직관으로 간접 연결된다.
- 표현과 문제 감각
  - 다국어 용어: `matrix`, `行列`, `矩阵`
  - 행렬곱은 보통 곱셈처럼 `순서 바꿔도 같을 것`이라고 생각하기 쉽지만, 실제로는 그렇지 않다.

## Deep Dive

### 왜 덧셈은 같은 크기에서만 가능한가

- 행렬의 각 성분은 `몇 행 몇 열`이라는 위치를 가진다.
- 따라서 같은 위치끼리 더하려면 행과 열의 개수가 같아야 한다.
- 크기가 다르면 대응되는 자리가 없으므로 덧셈을 정의할 수 없다.

### 핵심 정리 1: 행렬 덧셈은 성분별 덧셈이다

- 정리
  - 같은 크기의 행렬 $A=(a_{ij})$, $B=(b_{ij})$에 대해
    $$
    A+B=(a_{ij}+b_{ij})
    $$
    이다.
- 증명 스케치
  - `증명 스케치 (추론)`:
  - 행렬은 각 위치의 값을 담는 표이므로, 두 행렬을 합친다는 것은 같은 위치의 값을 합치는 것으로 보는 것이 가장 자연스럽다.
  - 그래서 $(i,j)$ 위치의 새 값은 $a_{ij}+b_{ij}$가 된다.
  - 이 규칙은 실수의 덧셈 성질을 성분별로 옮긴 것이다.

### 핵심 정리 2: 행렬곱은 일반적으로 교환법칙이 성립하지 않는다

- 정리
  - 보통 $AB\neq BA$ 이다.
- 증명
  - 다음 두 행렬을 보자.
    $$
    A=\begin{bmatrix}1&1\\0&1\end{bmatrix},\qquad
    B=\begin{bmatrix}1&0\\1&1\end{bmatrix}
    $$
  - 먼저
    $$
    AB=\begin{bmatrix}2&1\\1&1\end{bmatrix}
    $$
    이다.
  - 반면
    $$
    BA=\begin{bmatrix}1&1\\1&2\end{bmatrix}
    $$
    이다.
  - 두 결과가 다르므로
    $$
    AB\neq BA
    $$
    이다.

## Worked Examples

### 예제 1: 행렬의 덧셈

- 다음 계산을 보자.
  $$
  \begin{bmatrix}1&2\\3&4\end{bmatrix}
  +
  \begin{bmatrix}5&6\\7&8\end{bmatrix}
  =
  \begin{bmatrix}6&8\\10&12\end{bmatrix}
  $$
  이다.

### 예제 2: 행렬곱

- 다음 두 행렬을 보자.
  $$
  A=\begin{bmatrix}1&2\\3&4\end{bmatrix},\qquad
  B=\begin{bmatrix}2&0\\1&5\end{bmatrix}
  $$
  라 하자.
- 그러면
  $$
  AB=
  \begin{bmatrix}
  1\cdot 2+2\cdot 1 & 1\cdot 0+2\cdot 5\\
  3\cdot 2+4\cdot 1 & 3\cdot 0+4\cdot 5
  \end{bmatrix}
  =
  \begin{bmatrix}
  4&10\\
  10&20
  \end{bmatrix}
  $$
  이다.

## Common Pitfalls

- 크기가 다른 행렬도 그냥 더하려는 실수가 많다.
- 행렬곱에서 `행 곱하기 열`이 아니라 같은 위치끼리 곱해 버리는 오류가 자주 나온다.
- $AB$와 $BA$를 같은 것으로 생각하면 안 된다.
- 행렬곱이 가능하려면 `왼쪽 열 개수 = 오른쪽 행 개수`여야 한다는 조건을 놓치기 쉽다.

## Connections

- 선수 개념은 [matrix.md](./matrix.md), [algebraic-manipulation.md](./algebraic-manipulation.md), [simultaneous-equations.md](./simultaneous-equations.md)다.
- 같은 축의 인접 개념으로는 [common-math-1.md](./common-math-1.md), [vectors.md](./vectors.md)가 있다.
- 다음 개념으로는 [vectors.md](./vectors.md), [mathematics-for-ai.md](./mathematics-for-ai.md), [common-math-1.md](./common-math-1.md)가 이어진다.
- 학년 허브에서는 [high-1-hub.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/syntheses/high-1-hub.md), [high-2-hub.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/syntheses/high-2-hub.md)와 연결된다.

## Open Questions

- 역행렬과 행렬식까지 이 카드에 포함할지, 현재 수준에서 연산 규칙까지만 둘지 기준이 필요하다.
- `행렬 -> 벡터 -> 인공지능 수학` 연결을 별도 synthesis로 더 두껍게 만들지 검토가 필요하다.

## Sources

- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-curriculum-research/korea.md`
- `docs/math-concept-encyclopedia/multilingual-glossary.md`
- `docs/math-concept-encyclopedia/unit-practice-book.md`
