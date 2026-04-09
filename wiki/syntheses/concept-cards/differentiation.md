---
title: 미분
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
- 대표 성질
  - 미분 가능하면 연속이다.
  - 도함수의 값은 해당 점에서의 접선 기울기와 같다.
  - 평균변화율의 극한이 존재해야 순간변화율을 정의할 수 있다.
- 증명 스케치
  - `증명 스케치 (추론)`:
  - 평균변화율 $\frac{f(x+h)-f(x)}{h}$는 두 점을 잇는 할선의 기울기다.
  - $h$를 0에 가깝게 보내면 두 점이 하나의 점으로 모이고, 할선은 접선에 가까워진다.
  - 이 극한값이 순간변화율, 즉 미분이다.
  - 미분 가능하면 함수값이 주변에서 한쪽으로 끊어지지 않으므로 연속이어야 한다.
## Deep Dive

- 미분은 국소 선형근사의 언어다. 함수가 아주 짧은 구간에서는 직선처럼 보인다는 점을 수식으로 잡아낸다.
- 미분 가능한 점에서는 그래프가 접선으로 잘 설명되지만, 뾰족점이나 불연속점에서는 미분이 실패할 수 있다.
- 도함수는 원함수의 증가·감소와 굽음을 판정하는 도구로 이어진다.

## Worked Examples

### 예제 1: 이차함수 미분

- $f(x)=x^2$이면
  $$
  f'(x)=2x
  $$
  다.

### 예제 2: 직선의 미분

- $f(x)=3x-1$이면
  $$
  f'(x)=3
  $$
  이다. 직선의 기울기는 어디서나 일정하다.

### 예제 3: 미분 불가능한 예

- $f(x)=|x|$는 $x=0$에서 뾰족점이 생겨 미분 가능하지 않다.

## Common Pitfalls

- 미분값과 함수값을 같은 것으로 보지 말아야 한다.
- 평균변화율과 순간변화율을 혼동하면 안 된다.
- 미분 가능하면 연속이지만, 연속이라고 해서 항상 미분 가능한 것은 아니다.
- 뾰족점, 수직접선, 불연속점에서는 미분이 실패할 수 있다.
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
- 같은 축의 인접 개념으로는 [derivative.md](./derivative.md), [continuity.md](./continuity.md)가 있다.
- 다음 개념으로는 [derivative.md](./derivative.md), [integration.md](./integration.md), [calculus-course.md](./calculus-course.md)가 이어진다.
- 계통 허브에서는 [functions-strand.md](../functions-strand.md), [korea-curriculum-hub.md](../korea-curriculum-hub.md), [course-track-hub.md](../course-track-hub.md)를 함께 본다.

## Open Questions

- `미분법의 계산 규칙`과 `도함수의 그래프 해석`을 별도 세부 카드로 더 분리할지 검토가 필요하다.

## Sources

- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-concept-encyclopedia/japan.md`
- `docs/math-concept-encyclopedia/china.md`
- `docs/math-concept-encyclopedia/us.md`
- `docs/math-concept-encyclopedia/multilingual-glossary.md`
