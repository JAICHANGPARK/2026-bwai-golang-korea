---
title: 삼각함수 항등식
type: synthesis
status: active
updated: 2026-04-10
card_role: concept
source_docs:
  - docs/math-concept-encyclopedia/us.md
  - docs/math-concept-encyclopedia/multilingual-glossary.md
  - docs/math-concept-encyclopedia/formula-examples.md
  - docs/math-concept-encyclopedia/unit-practice-book.md
tags:
  - concept
  - trigonometry
  - identity
---

# 삼각함수 항등식

## Summary

삼각함수 항등식은 특정 각에서만 성립하는 등식이 아니라, 정의역 안의 모든 각에서 항상 성립하는 삼각함수 사이의 관계다. [trigonometric-ratio.md](./trigonometric-ratio.md)의 길이비가 [trigonometric-function.md](./trigonometric-function.md)의 함수 언어로 올라오면서 생기는 핵심 규칙이며, 이후 그래프 해석과 [derivative.md](./derivative.md)의 계산 기반이 된다.

## Key Points

- 정의
  - 삼각함수에 대해 모든 적절한 `x`에서 항상 성립하는 등식을 삼각함수 항등식이라 한다.
- 핵심 개념
  - 항등식
  - 단위원
  - 피타고라스 관계
  - 몫의 관계
  - 함수 사이의 변환
- 대표 수식
  - $\sin^2 x+\cos^2 x=1$
  - $\tan x=\frac{\sin x}{\cos x}$
  - $1+\tan^2 x=\frac{1}{\cos^2 x}$
- 핵심 해석
  - 항등식은 `값을 외우는 공식`이 아니라, 삼각함수끼리 서로 바꿔 쓰는 문법이다.
  - 한 함수만 남겨 식을 단순화하거나, 모르는 값을 다른 함수로 표현할 때 핵심 역할을 한다.
- 교육과정 배치
  - 미국 대표 경로에서는 `Precalculus`의 중심 개념 중 하나다.
  - 한국·일본·중국 문서군에서는 고등학교 삼각함수 확장의 핵심 성질로 읽는 것이 적절하다.
- 국가별 배치 스냅샷
  - 한국: `대수`의 삼각함수 이해를 두껍게 만드는 하위 개념으로 볼 수 있다.
  - 일본: `수학II`의 삼각함수 확장에서 중요한 기본 성질이다.
  - 중국: 중학의 예각 삼각함수에서 출발해 고등학교 일반 삼각함수로 확장될 때 본격화된다.
  - 미국: `Precalculus`에서 identities라는 이름으로 직접 강조된다.
- 표현과 문제 감각
  - 다국어 용어: `identity`, `恒等式`, `恒等式`
  - 식 정리는 계산 속도보다 `어떤 관계를 쓸지` 먼저 보는 습관이 중요하다.

## Deep Dive

### 항등식과 방정식은 어떻게 다른가

- 방정식은 특정 값에서만 참일 수 있다.
- 항등식은 정의역 안의 모든 값에서 참이어야 한다.
- 예를 들어
  $$
  \sin^2 x+\cos^2 x=1
  $$
  은 어느 `x`를 넣어도 성립하므로 항등식이다.

### 핵심 정리 1: $\sin^2 x+\cos^2 x=1$

- 증명
  - 단위원 위의 점을 $(\cos x,\sin x)$라고 두자.
  - 이 점은 항상 원의 방정식
    $$
    X^2+Y^2=1
    $$
    위에 있다.
  - 따라서
    $$
    (\cos x)^2+(\sin x)^2=1
    $$
    이고, 즉
    $$
    \sin^2 x+\cos^2 x=1
    $$
    이다.

### 핵심 정리 2: $\tan x=\frac{\sin x}{\cos x}$

- 증명
  - [trigonometric-ratio.md](./trigonometric-ratio.md)의 정의에 따라
    $$
    \sin x=\frac{\text{높이}}{\text{빗변}},\qquad
    \cos x=\frac{\text{밑변}}{\text{빗변}}
    $$
    이다.
  - 따라서
    $$
    \frac{\sin x}{\cos x}
    =
    \frac{\frac{\text{높이}}{\text{빗변}}}{\frac{\text{밑변}}{\text{빗변}}}
    =
    \frac{\text{높이}}{\text{밑변}}
    =
    \tan x
    $$
    이다.

### 핵심 정리 3: $1+\tan^2 x=\frac{1}{\cos^2 x}$

- 증명
  - $\tan x=\frac{\sin x}{\cos x}$를 이용하면
    $$
    1+\tan^2 x
    =
    1+\frac{\sin^2 x}{\cos^2 x}
    $$
    이다.
  - 분모를 통일하면
    $$
    \frac{\cos^2 x+\sin^2 x}{\cos^2 x}
    $$
    이고, 위의 피타고라스 항등식에 의해 분자는 1이다.
  - 따라서
    $$
    1+\tan^2 x=\frac{1}{\cos^2 x}
    $$
    이다.

## Worked Examples

### 예제 1: 값 구하기

- $\sin x=\frac{3}{5}$이고 `x`가 예각이면
  $$
  \cos^2 x=1-\sin^2 x=1-\frac{9}{25}=\frac{16}{25}
  $$
  이다.
- 따라서
  $$
  \cos x=\frac{4}{5}
  $$
  이다.

### 예제 2: 식 단순화

- 식
  $$
  \sin^2 x+\cos^2 x+2\tan x\cos x
  $$
  를 정리하자.
- 앞의 두 항은 1이고,
  $$
  \tan x\cos x=\frac{\sin x}{\cos x}\cos x=\sin x
  $$
  이다.
- 따라서 전체 식은
  $$
  1+2\sin x
  $$
  로 정리된다.

## Common Pitfalls

- 항등식은 `항상 참`이어야 하는데, 한 각에서 확인하고 끝내는 실수가 많다.
- $\sin^2 x$를 `\sin(x^2)`처럼 읽으면 안 된다. 이것은 `(\sin x)^2`다.
- $\tan x=\frac{\sin x}{\cos x}$는 $\cos x=0$인 곳에서는 쓸 수 없다는 점을 자주 놓친다.
- 식을 정리할 때 무조건 전개만 하면 복잡해질 수 있으므로, 먼저 어떤 항등식을 쓸지 보는 것이 중요하다.

## Connections

- 선수 개념은 [trigonometric-ratio.md](./trigonometric-ratio.md), [trigonometric-function.md](./trigonometric-function.md), [radians.md](./radians.md)다.
- 같은 축의 인접 개념으로는 [polar-ideas.md](./polar-ideas.md), [precalculus.md](./precalculus.md)가 있다.
- 다음 개념으로는 [derivative.md](./derivative.md), [calculus-course.md](./calculus-course.md), [vectors.md](./vectors.md)가 이어진다.
- 학년 허브에서는 [high-2-hub.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/syntheses/high-2-hub.md), [high-3-hub.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/syntheses/high-3-hub.md)와 연결된다.
- 계통 허브에서는 [functions-strand.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/syntheses/functions-strand.md), [geometry-strand.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/syntheses/geometry-strand.md)를 함께 본다.

## Open Questions

- 덧셈정리와 배각공식까지 이 카드에 포함할지, 별도 심화 카드로 뺄지 기준이 더 필요하다.
- `삼각함수 그래프`를 항등식 카드와 분리한 채 유지할지 후속 설계가 필요하다.

## Sources

- `docs/math-concept-encyclopedia/us.md`
- `docs/math-concept-encyclopedia/multilingual-glossary.md`
- `docs/math-concept-encyclopedia/formula-examples.md`
- `docs/math-concept-encyclopedia/unit-practice-book.md`
