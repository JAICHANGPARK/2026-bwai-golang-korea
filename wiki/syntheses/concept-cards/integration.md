---
title: 적분
type: synthesis
status: active
updated: 2026-04-09
source_docs:
  - docs/math-concept-encyclopedia/korea.md
  - docs/math-concept-encyclopedia/japan.md
  - docs/math-concept-encyclopedia/china.md
  - docs/math-concept-encyclopedia/us.md
  - docs/math-concept-encyclopedia/multilingual-glossary.md
tags:
  - concept
  - calculus
  - integral
---

# 적분

## Summary

적분은 넓이와 누적량을 읽는 개념이다. 잘게 나눈 양을 모아 전체를 계산하는 언어라서, 미분과 짝을 이루며 변화와 누적을 함께 설명한다.

## Key Points

- 정의
  - 함수값을 작은 조각들로 나누어 누적한 전체량을 구하는 과정을 적분이라 한다.
- 핵심 개념
  - 누적량
  - 넓이
  - 정적분
  - 부정적분
  - 합의 극한
- 대표 수식
  - $\int_a^b f(x)\,dx$
  - $\int x^n\,dx=\frac{x^{n+1}}{n+1}+C\;(n\neq -1)$
- 증명 스케치
  - `증명 스케치 (추론)`:
  - 구간을 아주 작은 폭으로 나누고 각 조각의 넓이를 직사각형으로 근사하면 전체 넓이는 작은 넓이들의 합으로 계산할 수 있다.
  - 이 분할을 무한히 잘게 하면 합이 정확한 넓이에 가까워지고, 이것이 적분의 기본 생각이다.
- 대표 예제
  - $\int_0^2 x\,dx=2$다.
- 교육과정 배치
  - 한국 대표 배치에서는 `미적분I`의 핵심이며 `미적분II`에서 더 고급화된다.
- 국가별 배치 스냅샷
  - 한국: `미적분I`와 `미적분II`로 확장된다.
  - 일본: `수학II`에서 적분의 생각을, `수학III`에서 본격 적분법을 다룬다.
  - 중국: 도함수와 함께 고등 미적분 축에서 다뤄진다.
  - 미국: Calculus의 accumulation and area unit이 대표적이다.
- 표현과 문제 감각
  - 다국어 용어: `integral`, `integration`, `積分`, `积分`
  - 적분은 `넓이`뿐 아니라 `누적 변화량`이라는 해석이 중요하다.

## Connections

- 선수 개념은 [limits.md](./limits.md), [differentiation.md](./differentiation.md)다.
- 다음 개념으로는 `물리의 운동`, `경제 누적량`, `확률밀도`, `대학 미적분`이 이어진다.

## Open Questions

- 정적분과 부정적분을 각각 분리할지 후속 설계가 필요하다.

## Sources

- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-concept-encyclopedia/japan.md`
- `docs/math-concept-encyclopedia/china.md`
- `docs/math-concept-encyclopedia/us.md`
- `docs/math-concept-encyclopedia/multilingual-glossary.md`
