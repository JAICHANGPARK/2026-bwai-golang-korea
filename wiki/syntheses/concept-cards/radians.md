---
title: 라디안
type: synthesis
status: active
updated: 2026-04-09
source_docs:
  - docs/math-concept-encyclopedia/us.md
  - docs/math-concept-encyclopedia/unit-practice-book.md
tags:
  - concept
  - trigonometry
  - angle
---

# 라디안

## Summary

라디안은 각의 크기를 `호의 길이`로 재는 방법이다. 도 단위가 `360등분`의 약속이라면, 라디안은 원 위 길이와 직접 연결되므로 [trigonometric-function.md](./trigonometric-function.md), [derivative.md](./derivative.md), [polar-ideas.md](./polar-ideas.md)로 이어지는 훨씬 수학적인 각도 언어다.

## Key Points

- 정의
  - 반지름이 `r`인 원에서 중심각 `\theta`가 잘라내는 호의 길이를 `s`라 할 때
    $$
    \theta=\frac{s}{r}
    $$
    로 정의한 각의 단위를 라디안이라 한다.
- 핵심 개념
  - 호의 길이
  - 반지름
  - 단위원
  - 각의 함수적 표현
  - 도와 라디안의 변환
- 대표 수식
  - $\theta=\frac{s}{r}$
  - $180^\circ=\pi\text{ rad}$
  - $360^\circ=2\pi\text{ rad}$
- 핵심 해석
  - 라디안은 `각을 길이의 비로 표현한 무차원 수`다.
  - 그래서 삼각함수를 함수처럼 다루거나 미분할 때 자연스럽다.
- 교육과정 배치
  - 미국 대표 경로에서는 `Precalculus`에서 삼각함수와 함께 중심적으로 다룬다.
  - 현재 위키 구조에서는 한국 `대수`의 삼각함수 카드와 미국 `Precalculus` 카드 사이의 세부 하위 개념으로 읽는 것이 적절하다.
- 국가별 배치 스냅샷
  - 한국: 현재 원천 문서에서는 삼각함수 하위 설명으로 간접 연결된다.
  - 일본: `수학II`의 삼각함수 확장에서 자연스럽게 전제되는 개념이다.
  - 중국: 중학의 `예각 삼각함수`보다 고등학교 일반 삼각함수 쪽에서 더 중요한 배경이 된다.
  - 미국: `Precalculus`의 핵심 개념으로 직접 등장한다.
- 표현과 문제 감각
  - `radian`은 각을 재는 새 자일 뿐 아니라, 각과 길이를 연결하는 좌표적 언어라는 감각이 중요하다.

## Deep Dive

### 왜 $\theta=\frac{s}{r}$ 인가

- 원이 클수록 같은 각이라도 호의 길이는 길어진다.
- 따라서 `호의 길이`만으로 각을 재면 원의 크기에 따라 값이 달라진다.
- 반지름 `r`로 한 번 나누면 원의 크기 효과가 사라지고, 같은 모양의 각은 언제나 같은 값으로 표현된다.

### 핵심 정리 1: 반지름과 호의 길이가 같으면 1라디안이다

- 정리
  - 반지름이 `r`이고 호의 길이도 `r`이면 그 중심각의 크기는 `1 rad`다.
- 증명
  - 정의에 의해
    $$
    \theta=\frac{s}{r}
    $$
    이다.
  - 여기서 $s=r$이면
    $$
    \theta=\frac{r}{r}=1
    $$
    이다.
  - 따라서 `1 rad`는 `반지름 길이만큼의 호를 잘라내는 각`이다.

### 핵심 정리 2: $180^\circ=\pi$ 라디안

- `증명 스케치 (추론)`:
  - 반지름이 `r`인 원의 둘레는
    $$
    2\pi r
    $$
    이다.
  - 원 전체에 해당하는 중심각은 한 바퀴이므로
    $$
    \theta=\frac{2\pi r}{r}=2\pi
    $$
    라디안이다.
  - 한 바퀴가 $360^\circ$이므로
    $$
    360^\circ=2\pi\text{ rad}
    $$
    이고, 양변을 2로 나누면
    $$
    180^\circ=\pi\text{ rad}
    $$
    를 얻는다.

## Worked Examples

### 예제 1: 도를 라디안으로 바꾸기

- $60^\circ$를 라디안으로 바꾸면
  $$
  60^\circ=\frac{60}{180}\pi=\frac{\pi}{3}
  $$
  이다.

### 예제 2: 호의 길이 구하기

- 반지름이 `6`이고 중심각이 $\frac{\pi}{2}$ 라디안이면
  $$
  s=r\theta=6\cdot \frac{\pi}{2}=3\pi
  $$
  이다.

## Common Pitfalls

- $\pi$ 라디안을 `180라디안`으로 착각하면 안 된다. $\pi$ 라디안은 `180도`와 같은 크기다.
- 도 단위와 라디안을 한 식 안에서 섞어 쓰면 삼각함수 계산이 틀어지기 쉽다.
- 라디안을 단순 변환 공식으로만 외우면, 왜 미적분에서 중요한지 잘 보이지 않는다.
- 라디안은 길이 그 자체가 아니라 `호의 길이 / 반지름`이라는 비라는 점을 놓치기 쉽다.

## Connections

- 선수 개념은 [circle.md](./circle.md), [trigonometric-ratio.md](./trigonometric-ratio.md), [function-foundations.md](./function-foundations.md)다.
- 같은 축의 인접 개념으로는 [trigonometric-function.md](./trigonometric-function.md), [polar-ideas.md](./polar-ideas.md)가 있다.
- 다음 개념으로는 [trigonometric-identities.md](./trigonometric-identities.md), [derivative.md](./derivative.md), [precalculus.md](./precalculus.md)가 이어진다.
- 학년 허브에서는 [high-2-hub.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/queries/math-curriculum-graph/high-2-hub.md), [high-3-hub.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/queries/math-curriculum-graph/high-3-hub.md)와 연결된다.
- 계통 허브에서는 [functions-strand.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/queries/math-curriculum-graph/functions-strand.md), [geometry-strand.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/queries/math-curriculum-graph/geometry-strand.md)를 함께 본다.

## Open Questions

- `단위원`을 별도 카드로 분리할지, 현재처럼 [trigonometric-function.md](./trigonometric-function.md)와 라디안 카드 사이에서 설명할지 기준이 더 필요하다.
- `호의 길이`와 `부채꼴 넓이`를 라디안 카드 안에 더 깊게 넣을지 후속 설계가 필요하다.

## Sources

- `docs/math-concept-encyclopedia/us.md`
- `docs/math-concept-encyclopedia/unit-practice-book.md`
