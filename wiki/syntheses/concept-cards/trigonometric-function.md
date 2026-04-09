---
title: 삼각함수
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

- 선수 개념은 [trigonometric-ratio.md](./trigonometric-ratio.md), `단위원`, `좌표평면`이다.
- 다음 개념으로는 `미분`, `파동`, `벡터`, `극좌표`가 이어진다.
- 학년 허브에서는 [high-2-hub.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/queries/math-curriculum-graph/high-2-hub.md), [high-3-hub.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/queries/math-curriculum-graph/high-3-hub.md)와 연결된다.
- 계통 허브에서는 [functions-strand.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/queries/math-curriculum-graph/functions-strand.md)와 [geometry-strand.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/queries/math-curriculum-graph/geometry-strand.md)를 함께 본다.

## Open Questions

- `단위원`과 `라디안`을 삼각함수 카드의 하위 설명으로 둘지 독립 카드로 뺄지 기준이 필요하다.
- `삼각함수의 그래프`를 별도 시각화 카드로 만들지 검토가 필요하다.

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
