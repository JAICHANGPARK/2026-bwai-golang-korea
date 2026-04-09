---
title: 이차방정식
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
  - docs/math-concept-encyclopedia/formula-examples.md
  - docs/math-concept-encyclopedia/unit-practice-book.md
  - docs/math-concept-encyclopedia/comparative-problem-book.md
tags:
  - concept
  - algebra
  - equation
---

# 이차방정식

## Summary

이차방정식은 미지수가 2차로 나타나는 방정식의 해를 구하는 개념이다. 인수분해, 완전제곱식, 근의 공식을 묶어 식의 구조를 읽게 해 주며, [quadratic-function.md](./quadratic-function.md)의 x절편 해석과 직접 연결되는 핵심 브리지다.

## Key Points

- 정의
  - 가장 대표적인 형태는 $ax^2+bx+c=0\;(a\neq 0)$이다.
- 핵심 개념
  - 근
  - 인수분해 풀이
  - 완전제곱식
  - 근의 공식
  - 판별식
- 대표 수식
  - $ax^2+bx+c=0$
  - $x=\frac{-b\pm\sqrt{b^2-4ac}}{2a}$
- 증명 스케치
  - `증명 스케치 (추론)`:
  - $ax^2+bx+c=0$의 양변을 $a$로 나누고 완전제곱을 만들면
    $$
    \left(x+\frac{b}{2a}\right)^2=\frac{b^2-4ac}{4a^2}
    $$
    가 된다.
  - 양변에 제곱근을 취해 정리하면
    $$
    x=\frac{-b\pm\sqrt{b^2-4ac}}{2a}
    $$
    를 얻는다.
  - 이 과정은 `제곱식의 구조를 만들어 푼다`는 이차방정식의 핵심 발상을 보여 준다.
- 대표 예제
  - $x^2-4x-5=0$은 인수분해하면
    $$
    (x-5)(x+1)=0
    $$
    이므로 해는 $x=5,-1$이다.
- 교육과정 배치
  - 한국 대표 배치에서는 `중3` 핵심 단원이다.
  - 이후 [quadratic-function.md](./quadratic-function.md), [equations-of-geometric-figures.md](./equations-of-geometric-figures.md), [complex-numbers.md](./complex-numbers.md)의 배경으로 이어진다.
- 국가별 배치 스냅샷
  - 한국: `중3`에서 인수분해와 함께 배우고, 함수의 x절편 해석으로 이어진다.
  - 일본: `중3`에서 배우며, 근의 공식의 절차 설명과 판별식 해석이 중요하다.
  - 중국: `9학년 1학기`에 도입되고, 계산 구조와 함수 연결이 빠르게 이어진다.
  - 미국: 보통 `Algebra I` 또는 `Integrated Mathematics` 경로에서 그래프와 x절편을 함께 읽으며 배운다.
- 표현과 문제 감각
  - 다국어 용어: `quadratic equation`, `二次方程式`, `一元二次方程`
  - 비교 문제집 기준으로 미국은 그래프와 함께 읽는 경우가 많고, 일본은 절차 설명, 한국·중국은 계산 정확도와 속도가 더 강하게 요구된다.

## Connections

- 선수 개념은 `인수분해`, `완전제곱`, [linear-equation.md](./linear-equation.md)이다.
- 다음 개념으로는 [quadratic-function.md](./quadratic-function.md), [equations-of-geometric-figures.md](./equations-of-geometric-figures.md), [complex-numbers.md](./complex-numbers.md)가 이어진다.
- 학년 허브에서는 [middle-3-hub.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/queries/math-curriculum-graph/middle-3-hub.md), [high-1-hub.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/queries/math-curriculum-graph/high-1-hub.md)와 연결된다.

## Open Questions

- `근의 공식`과 `판별식`을 이 카드 안에 유지할지 독립 카드로 분리할지 기준이 필요하다.
- 미국식 그래프 중심 풀이와 한국식 계산 중심 풀이를 별도 비교 synthesis로 뺄지 검토가 필요하다.

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
- `docs/math-concept-encyclopedia/formula-examples.md`
- `docs/math-concept-encyclopedia/unit-practice-book.md`
- `docs/math-concept-encyclopedia/comparative-problem-book.md`
