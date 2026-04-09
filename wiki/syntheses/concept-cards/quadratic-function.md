---
title: 이차함수
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
  - function
  - quadratic
---

# 이차함수

## Summary

이차함수는 선형 관계를 넘어 처음으로 본격적인 비선형 변화를 다루는 함수 개념이다. 식의 구조, 그래프의 모양, 꼭짓점과 최대·최소를 함께 읽어야 해서 중학교 후반에서 고등학교 함수 해석으로 넘어가는 핵심 브리지 역할을 한다.

## Key Points

- 정의
  - 가장 대표적인 형태는 $y=ax^2+bx+c\;(a\neq 0)$이며, 그래프는 포물선이다.
- 핵심 개념
  - 포물선
  - 꼭짓점
  - 축
  - 최대와 최소
  - x절편
  - 대칭성
- 대표 수식
  - $y=ax^2+bx+c=a(x-p)^2+q$
  - $x=-\frac{b}{2a}$
- 증명 스케치
  - `증명 스케치 (추론)`:
  - $y=ax^2+bx+c$를 완전제곱하면
    $$
    y=a\left(x+\frac{b}{2a}\right)^2+\left(c-\frac{b^2}{4a}\right)
    $$
    꼴이 된다.
  - 제곱항은 항상 0 이상이므로, $a>0$이면 꼭짓점에서 최소값을, $a<0$이면 최대값을 가진다.
  - 따라서 이차함수의 그래프는 꼭짓점을 중심으로 대칭이며, 꼭짓점형이 그래프 해석에 가장 직접적이다.
- 대표 예제
  - $y=x^2-4x+1$을 완전제곱하면
    $$
    y=(x-2)^2-3
    $$
    이므로 꼭짓점은 `(2,-3)`이고 축은 $x=2$다.
- 교육과정 배치
  - 한국 대표 배치에서는 `중3` 핵심 단원이고, 고등학교에서 함수 일반론과 최적화 감각으로 다시 확장된다.
- 국가별 배치 스냅샷
  - 한국: `중3`에서 `이차방정식`과 밀접하게 연결되어 다뤄진다.
  - 일본: `중3`에서 `비례 -> 일차함수 -> 이차함수` 계통의 정점처럼 배치되고, 고등학교 `수학I`에서 다시 중요해진다.
  - 중국: `9학년 1학기`에 도입되며, 구조적 식 변형과 그래프 해석이 함께 강조된다.
  - 미국: 보통 `Algebra I`나 `Integrated Mathematics` 경로에서 처음 배우고, 이후 `Algebra II`와 `Calculus` 직관으로 이어진다.
- 표현과 문제 감각
  - 다국어 용어: `quadratic function`, `二次関数`, `二次函数`
  - 비교 문제집 기준으로 한국은 꼭짓점 계산, 일본은 그래프 서술, 중국은 식 변형, 미국은 `vertex form` 변환이 자주 강조된다.

## Connections

- 선수 개념은 [linear-function.md](./linear-function.md), `인수분해`, [quadratic-equation.md](./quadratic-equation.md)이다.
- 다음 개념으로는 `최적화`, [equations-of-geometric-figures.md](./equations-of-geometric-figures.md), `미분의 극값 해석`이 이어진다.
- 학년 허브에서는 [middle-3-hub.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/queries/math-curriculum-graph/middle-3-hub.md), [high-1-hub.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/queries/math-curriculum-graph/high-1-hub.md)와 연결된다.
- 계통 허브에서는 [functions-strand.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/queries/math-curriculum-graph/functions-strand.md)를 본다.

## Open Questions

- `이차함수` 카드 안에 `이차방정식과 x절편의 관계`를 더 강하게 넣을지 별도 브리지 카드로 뺄지 기준이 필요하다.
- `최대·최소`를 이 카드의 하위 설명으로 둘지 독립 개념 카드로 분리할지 후속 결정이 필요하다.

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
