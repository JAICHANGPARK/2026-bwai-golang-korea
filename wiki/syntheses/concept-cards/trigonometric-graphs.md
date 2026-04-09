---
title: 삼각함수의 그래프
type: synthesis
status: active
updated: 2026-04-09
source_docs:
  - docs/math-concept-encyclopedia/us.md
  - docs/math-concept-encyclopedia/unit-practice-book.md
  - docs/math-concept-encyclopedia/formula-examples.md
tags:
  - concept
  - trigonometry
  - graph
---

# 삼각함수의 그래프

## Summary

삼각함수의 그래프는 [trigonometric-function.md](./trigonometric-function.md)의 주기성과 [radians.md](./radians.md)의 각도 언어를 눈으로 읽게 해 주는 카드다. 함수값이 반복된다는 사실을 그래프 차원에서 확인하면, 이후 파동 해석과 미분 연결이 훨씬 자연스러워진다.

## Key Points

- 정의
  - $\sin x$, $\cos x$, $\tan x$의 값을 좌표평면에 나타낸 그래프를 삼각함수의 그래프라 한다.
- 핵심 개념
  - 주기
  - 최대값과 최소값
  - 진폭
  - 영점
  - 점근선
- 대표 수식
  - $\sin(x+2\pi)=\sin x$
  - $\cos(x+2\pi)=\cos x$
  - $\tan(x+\pi)=\tan x$
- 핵심 해석
  - 사인과 코사인은 부드럽게 반복되고, 탄젠트는 주기적으로 점근선을 가진다.
- 교육과정 배치
  - 미국 `Precalculus`, 한국 `대수`, 일본 `수학II`의 삼각함수 확장 맥락에서 읽는 것이 적절하다.

## Deep Dive

### 핵심 정리 1: 사인과 코사인의 주기는 $2\pi$다

- `증명 스케치 (추론)`:
  - 단위원에서 각을 `2\pi`만큼 더하는 것은 한 바퀴를 더 도는 것과 같다.
  - 따라서 단위원 위의 점이 원래 위치로 돌아오므로
    $$
    \sin(x+2\pi)=\sin x,\qquad \cos(x+2\pi)=\cos x
    $$
    가 된다.

### 핵심 정리 2: 탄젠트의 주기는 $\pi$다

- `증명 스케치 (추론)`:
  - 탄젠트는
    $$
    \tan x=\frac{\sin x}{\cos x}
    $$
    이다.
  - 각을 $\pi$만큼 더하면 사인과 코사인이 둘 다 부호가 바뀌므로 비의 값은 그대로 남는다.
  - 따라서
    $$
    \tan(x+\pi)=\tan x
    $$
    이다.

## Worked Examples

### 예제 1: 사인 그래프 읽기

- $y=\sin x$는
  - $x=0$에서 0
  - $x=\frac{\pi}{2}$에서 1
  - $x=\pi$에서 0
  - $x=\frac{3\pi}{2}$에서 -1
  - $x=2\pi$에서 다시 0
  를 지난다.

### 예제 2: 탄젠트의 점근선

- $y=\tan x$는 $\cos x=0$인 곳에서 정의되지 않는다.
- 따라서
  $$
  x=\frac{\pi}{2}+k\pi
  $$
  에 수직 점근선이 생긴다.

## Common Pitfalls

- 사인과 코사인의 주기가 $\pi$라고 착각하기 쉽다.
- 탄젠트는 모든 실수에서 정의되지 않는다는 점을 놓치기 쉽다.
- 그래프를 그릴 때 `도`와 `라디안` 축을 섞으면 형태는 맞아도 눈금이 틀어진다.

## Connections

- 선수 개념은 [trigonometric-function.md](./trigonometric-function.md), [radians.md](./radians.md), [trigonometric-identities.md](./trigonometric-identities.md)다.
- 같은 축의 인접 개념으로는 [precalculus.md](./precalculus.md), [polar-ideas.md](./polar-ideas.md)가 있다.
- 다음 개념으로는 [derivative.md](./derivative.md), [calculus-course.md](./calculus-course.md), [vectors.md](./vectors.md)가 이어진다.

## Open Questions

- 진폭과 위상 이동을 별도 그래프 변환 카드로 더 뺄지 기준이 필요하다.

## Sources

- `docs/math-concept-encyclopedia/us.md`
- `docs/math-concept-encyclopedia/unit-practice-book.md`
- `docs/math-concept-encyclopedia/formula-examples.md`
