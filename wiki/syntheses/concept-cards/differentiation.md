---
title: 미분
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
  - derivative
---

# 미분

## Summary

미분은 함수의 순간변화율을 읽는 개념이다. 직선의 기울기 감각이 곡선의 접선 기울기와 변화율 해석으로 확장되는 지점이라서, 함수 해석의 강력한 도구가 된다.

## Key Points

- 정의
  - 함수의 순간변화율을 구하는 과정을 미분이라 하고, 그 결과를 도함수라 한다.
- 핵심 개념
  - 순간변화율
  - 접선의 기울기
  - 도함수
  - 증가와 감소
  - 극값
- 대표 수식
  - $f'(x)=\lim_{h\to 0}\frac{f(x+h)-f(x)}{h}$
- 증명 스케치
  - `증명 스케치 (추론)`:
  - 평균변화율 $\frac{f(x+h)-f(x)}{h}$는 두 점을 잇는 할선의 기울기다.
  - $h$를 0에 가깝게 보내면 두 점이 하나의 점으로 모이고, 할선은 접선에 가까워진다.
  - 이 극한값이 순간변화율, 즉 미분이다.
- 대표 예제
  - $f(x)=x^2$이면 $f'(x)=2x$다.
- 교육과정 배치
  - 한국 대표 배치에서는 `미적분I`의 중심이며 `미적분II`에서 더 확장된다.
- 국가별 배치 스냅샷
  - 한국: `미적분I`, `미적분II`로 이어진다.
  - 일본: `수학II`에서 미분의 생각을, `수학III`에서 본격 미분법을 다룬다.
  - 중국: `수열과 도함수`에서 변화율과 도함수를 빠르게 연결한다.
  - 미국: Calculus의 핵심 단원이다.
- 표현과 문제 감각
  - 다국어 용어: `differentiation`, `derivative`, `微分`, `導関数`, `导数`
  - 기울기 계산을 넘어 함수의 증감과 최적화를 읽는 것이 중요하다.

## Connections

- 선수 개념은 [limits.md](./limits.md), [linear-function.md](./linear-function.md), [quadratic-function.md](./quadratic-function.md)다.
- 다음 개념으로는 [integration.md](./integration.md), `최적화`, `운동 해석`이 이어진다.

## Open Questions

- `도함수`를 미분 카드 내부 개념으로 둘지 별도 카드로 분리할지 검토가 필요하다.

## Sources

- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-concept-encyclopedia/japan.md`
- `docs/math-concept-encyclopedia/china.md`
- `docs/math-concept-encyclopedia/us.md`
- `docs/math-concept-encyclopedia/multilingual-glossary.md`
