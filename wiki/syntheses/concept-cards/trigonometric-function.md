---
title: 삼각함수
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
  - function
  - trigonometry
---

# 삼각함수

## Summary

삼각함수는 [trigonometric-ratio.md](./trigonometric-ratio.md)를 함수의 언어로 확장한 개념이다. 각의 변화, 주기성, 파동과 회전 같은 현상을 다루는 핵심 도구이며, 고등학교 함수 체계와 미적분으로 들어가는 중요한 문 역할을 한다.

## Key Points

- 정의
  - 각 또는 실수 입력에 대해 $\sin x$, $\cos x$, $\tan x$ 같은 값을 대응시키는 함수를 삼각함수라 한다.
- 핵심 개념
  - 단위원
  - 주기성
  - 라디안
  - 사인, 코사인, 탄젠트
  - 항등식
- 대표 수식
  - $\sin^2 x+\cos^2 x=1$
  - $\sin(x+2\pi)=\sin x$
  - $\cos(x+2\pi)=\cos x$
- 성질 1
  - $\sin x$와 $\cos x$는 각각 주기 $2\pi$를 가진다.
- 성질 2
  - $\sin x$는 홀함수이고 $\cos x$는 짝함수다.
- 증명 스케치
  - `증명 스케치 (추론)`:
  - 단위원 위의 점을 $(\cos x,\sin x)$로 두면, 그 점은 항상 원의 방정식
    $$
    X^2+Y^2=1
    $$
    을 만족한다.
  - 따라서
    $$
    \cos^2 x+\sin^2 x=1
    $$
    이 성립한다.
  - 이 항등식은 삼각함수를 `길이비`를 넘어 `함수와 좌표`의 언어로 읽게 해 준다.
- 대표 예제
  - 삼각함수의 기본 항등식에 따르면
    $$
    \sin^2 x+\cos^2 x=1
    $$
    이다.
  - 예를 들어 $x=\frac{\pi}{3}$이면 $\sin x=\frac{\sqrt3}{2}$, $\cos x=\frac12$이므로 두 제곱의 합은 1이다.
  - $\sin\left(\frac{7\pi}{6}\right)=-\frac12$이고 $\cos\left(\frac{7\pi}{6}\right)=-\frac{\sqrt3}{2}$이다.
  - $\sin(-x)=-\sin x$이므로 $\sin\left(-\frac{\pi}{4}\right)=-\frac{\sqrt2}{2}$이다.

## Deep Dive

  - 삼각함수는 단위원 위의 회전을 실수 전체로 확장한 함수라고 볼 수 있다.
  - $x$가 한 바퀴 $2\pi$만큼 늘어나면 같은 점으로 돌아오므로 주기성이 생긴다.
  - 따라서 삼각함수의 그래프는 반복되는 파동 모양을 갖는다.

## Worked Examples

### 예제 1: 값 계산하기

- $\sin\left(\frac{5\pi}{6}\right)$은 단위원에서 두 번째 사분면의 사인값이므로
  $$
  \sin\left(\frac{5\pi}{6}\right)=\frac12
  $$
  이다.

### 예제 2: 주기성 이용하기

- $\cos\left(\frac{13\pi}{6}\right)$은
  $$
  \cos\left(\frac{13\pi}{6}\right)=\cos\left(\frac{\pi}{6}\right)=\frac{\sqrt3}{2}
  $$
  이다.
- $\tan\left(-\frac{\pi}{3}\right)=-\sqrt3$이다.

## Common Pitfalls

- 도 단위와 라디안을 섞어 쓰면 값이 틀어진다.
- $\sin^2 x+\cos^2 x=1$을 $\sin x+\cos x=1$로 잘못 읽기 쉽다.
- 삼각함수는 예각에서만 정의된다고 생각하면 안 된다.
- 주기성을 쓸 때는 `얼마나` 더했는지 정확히 확인해야 한다.

## Curriculum Context

- 교육과정 배치
  - 한국 대표 배치에서는 고등학교 `대수`에서 배우는 핵심 함수 단원이다.
  - 이후 `미적분`, `기하`, `파동 해석`, `벡터`로 이어진다.
- 국가별 배치 스냅샷
  - 한국: `대수` 과목에서 `지수함수·로그함수·수열`과 함께 묶여 배운다.
  - 일본: `수학II`에서 `삼角関数`로 본격 도입되며, `수학I`의 삼각비를 함수 관점으로 확장한다.
  - 중국: 중학교 `예각 삼각함수` 이후, 고등학교 필수과정의 `함수` 축에서 일반적인 삼각함수로 확장된다.
  - 미국: 보통 `Precalculus`에서 중심적으로 다루고, `Geometry`의 trigonometric ratios가 전단계를 이룬다.
- 표현과 문제 감각
  - 다국어 용어: `trigonometric function`, `三角関数`, `三角函数`
  - 비교 문제집 기준으로 일본은 `삼각비 -> 삼각함수` 위계 설명이 분명하고, 중국은 중학의 예각 삼각함수에서 이어지며, 미국은 `Precalculus`에서 주기성과 identity를 함께 묻는 편이다.

## Connections

- 선수 개념은 [trigonometric-ratio.md](./trigonometric-ratio.md), [radians.md](./radians.md), [function-foundations.md](./function-foundations.md)다.
- 같은 축의 인접 개념으로는 [trigonometric-identities.md](./trigonometric-identities.md), [trigonometric-graphs.md](./trigonometric-graphs.md), [polar-ideas.md](./polar-ideas.md)가 있다.
- 다음 개념으로는 [derivative.md](./derivative.md), [vectors.md](./vectors.md), [precalculus.md](./precalculus.md)가 이어진다.
- 학년 허브에서는 [../high-2-hub.md](../high-2-hub.md), [../high-3-hub.md](../high-3-hub.md)와 연결된다.
- 계통 허브에서는 [../functions-strand.md](../functions-strand.md)와 [../geometry-strand.md](../geometry-strand.md)를 함께 본다.

## Open Questions

- `단위원`을 [radians.md](./radians.md), [trigonometric-function.md](./trigonometric-function.md), [polar-ideas.md](./polar-ideas.md) 중 어디의 중심 브리지로 둘지 기준이 더 필요하다.

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
