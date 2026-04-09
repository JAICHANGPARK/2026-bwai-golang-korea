---
title: 도함수
type: synthesis
status: active
updated: 2026-04-10
card_role: concept
source_docs:
  - docs/math-concept-encyclopedia/korea.md
  - docs/math-concept-encyclopedia/us.md
  - docs/math-concept-encyclopedia/japan.md
  - docs/math-concept-encyclopedia/china.md
  - docs/math-concept-encyclopedia/formula-examples.md
  - docs/math-concept-encyclopedia/unit-practice-book.md
  - docs/math-concept-encyclopedia/multilingual-glossary.md
tags:
  - concept
  - calculus
  - derivative
---

# 도함수

## Summary

도함수는 함수의 각 점에서 순간변화율을 새 함수로 모아 놓은 것이다. [differentiation.md](./differentiation.md)가 `미분한다`는 과정을 설명한다면, 이 카드는 그 결과물인 `도함수`가 무엇을 뜻하고 어떻게 계산되는지를 더 또렷하게 보여 준다.

## Key Points

- 정의
  - 함수 $f(x)$에 대해
    $$
    f'(x)=\lim_{h\to 0}\frac{f(x+h)-f(x)}{h}
    $$
    가 존재하면 이를 $f$의 도함수라 한다.
- 핵심 개념
  - 순간변화율
  - 접선의 기울기
  - 미분계수와 도함수
  - 증가와 감소
  - 함수의 변화 해석
- 대표 수식
  - $f'(x)=\lim_{h\to 0}\frac{f(x+h)-f(x)}{h}$
  - $(x^2)'=2x$
  - $(x^3)'=3x^2$
- 핵심 해석
  - 도함수는 원래 함수의 `변화 패턴을 기록한 새 함수`다.
  - 어떤 한 점의 기울기만 말하는 미분계수와 달리, 도함수는 `모든 x에 대한 기울기 규칙`을 준다.
- 교육과정 배치
  - 일본 `수학III`, 중국 `수열과 도함수`, 미국 `Calculus`에서 직접적인 핵심 개념으로 드러난다.
  - 한국 문서군에서는 `미적분I`, `미적분II`의 중심 개념을 세분화한 하위 카드로 읽는 것이 적절하다.
- 국가별 배치 스냅샷
  - 한국: `미적분I/II`의 도함수 활용 문맥에서 핵심 역할을 한다.
  - 일본: `수학III`의 `도함수, 극값, 정적분` 흐름 안에서 분명하게 등장한다.
  - 중국: `수열과 도함수`에서 변화율과 함께 빠르게 연결된다.
  - 미국: `Calculus`의 핵심 언어다.
- 표현과 문제 감각
  - 다국어 용어: `derivative`, `導関数`, `导数`
  - 도함수는 계산 결과표가 아니라 `증가·감소`, `극값`, `최적화`를 읽는 지도라고 보는 것이 중요하다.

## Deep Dive

### 미분계수와 도함수는 어떻게 다른가

- `x=a`에서의 미분계수는 한 점에서의 순간변화율이다.
- 도함수는 그 한 점짜리 정보를 모든 `x`에 대해 모아 만든 함수다.
- 따라서 미분계수는 `값`, 도함수는 `함수`라는 차이가 있다.

### 핵심 정리 1: $f(x)=x^2$의 도함수는 $2x$다

- 증명
  - 정의에 따라
    $$
    f'(x)=\lim_{h\to 0}\frac{(x+h)^2-x^2}{h}
    $$
    이다.
  - 분자를 전개하면
    $$
    (x+h)^2-x^2=x^2+2xh+h^2-x^2=2xh+h^2
    $$
    이다.
  - 따라서
    $$
    f'(x)=\lim_{h\to 0}\frac{2xh+h^2}{h}
    =\lim_{h\to 0}(2x+h)
    $$
    이다.
  - 이제 $h\to 0$을 보내면
    $$
    f'(x)=2x
    $$
    를 얻는다.

### 핵심 정리 2: 도함수의 부호와 증가·감소

- `증명 스케치 (추론)`:
  - 어떤 구간에서 $f'(x)>0$이면, 작은 증가량에 대해 함수값이 대체로 증가하는 방향으로 변한다.
  - 반대로 $f'(x)<0$이면 함수값이 감소하는 방향으로 변한다.
  - 그래서 도함수의 부호를 보면 원래 함수의 증가와 감소를 읽을 수 있다.

## Worked Examples

### 예제 1: $f(x)=x^2+1$의 도함수

- 상수항의 도함수는 0이므로
  $$
  (x^2+1)'=2x
  $$
  이다.

### 예제 2: $f(x)=x^3$의 도함수

- 기본 미분 규칙에 따라
  $$
  (x^3)'=3x^2
  $$
  이다.
- 따라서 $x=2$에서의 순간변화율은
  $$
  f'(2)=3\cdot 2^2=12
  $$
  이다.

## Common Pitfalls

- 미분과 도함수를 완전히 같은 말로만 생각하면 `과정`과 `결과`의 차이가 흐려진다.
- $f'(a)$와 $f'(x)$를 구분하지 않으면 한 점의 기울기와 함수 전체의 규칙이 섞인다.
- 극한 정의에서 분자를 먼저 전개하지 않고 바로 약분하려다 실수하기 쉽다.
- 도함수를 구한 뒤 원래 함수의 증가·감소와 연결하지 않으면 계산만 남고 해석이 사라진다.

## Connections

- 선수 개념은 [limits.md](./limits.md), [continuity.md](./continuity.md), [differentiation.md](./differentiation.md)다.
- 같은 축의 인접 개념으로는 [derivative-graph-interpretation.md](./derivative-graph-interpretation.md), [integration.md](./integration.md), [trigonometric-identities.md](./trigonometric-identities.md)가 있다.
- 다음 개념으로는 [derivative-graph-interpretation.md](./derivative-graph-interpretation.md), [integration.md](./integration.md), [calculus-course.md](./calculus-course.md)가 이어진다.
- 학년 허브에서는 [high-2-hub.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/syntheses/high-2-hub.md), [high-3-hub.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/syntheses/high-3-hub.md)와 연결된다.
- 계통 허브에서는 [functions-strand.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/syntheses/functions-strand.md)를 본다.

## Open Questions

- `합성함수의 미분`, `곱의 미분`, `몫의 미분`을 이 카드에 더 넣을지 별도 규칙 카드로 분리할지 기준이 필요하다.

## Sources

- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-concept-encyclopedia/us.md`
- `docs/math-concept-encyclopedia/japan.md`
- `docs/math-concept-encyclopedia/china.md`
- `docs/math-concept-encyclopedia/formula-examples.md`
- `docs/math-concept-encyclopedia/unit-practice-book.md`
- `docs/math-concept-encyclopedia/multilingual-glossary.md`
