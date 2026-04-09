---
title: 일차함수
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
  - linear
---

# 일차함수

## Summary

일차함수는 `y=ax+b` 꼴로 표현되는 가장 기본적인 함수 개념이다. 중학교에서는 일정한 변화율을 식과 그래프로 읽는 첫 본격 함수이고, 고등학교에서는 함수와 그래프 일반론의 출발점이 된다.

## Key Points

- 정의
  - `y=ax+b` 꼴의 함수를 일차함수라 한다. 여기서 `a`는 기울기, `b`는 y절편이다.
- 핵심 개념
  - 함수
  - 독립변수와 종속변수
  - 기울기
  - 절편
  - 변화율
  - 그래프 해석
- 대표 수식
  - $y=ax+b$
  - $a=\frac{y_2-y_1}{x_2-x_1}$
- 핵심 해석
  - `a`는 `x`가 1 증가할 때 `y`가 얼마나 변하는가를 뜻한다.
  - `b`는 그래프가 y축과 만나는 시작 높이이므로 `초깃값`으로 읽을 수 있다.
- 교육과정 배치
  - 한국 대표 배치에서는 `중2` 핵심 단원이고, 고1에서 `함수와 그래프`로 다시 일반화된다.
- 국가별 배치 스냅샷
  - 한국: `중2` 핵심 단원이며, 고1 함수 일반론으로 다시 확장된다.
  - 일본: `중2`의 `일차함수`가 `비례·반비례` 다음 단계로 놓인다.
  - 중국: `8학년 2학기`에 `일차함수`가 놓이고, 이어 `반비례함수`, `이차함수`로 확장된다.
  - 미국: `Grade 8`의 `Functions and Linear Models`와 `Linear Equations and Systems`, 또는 `Algebra I` 경로에서 반복적으로 다뤄진다.
- 표현과 문제 감각
  - 다국어 용어: `linear function`, `一次関数`, `一次函数`
  - 비교 문제집 기준으로 일본은 `변화의 비율` 설명이 자주 붙고, 미국은 `slope-intercept form` 언어가 자주 등장한다.

## Deep Dive

### 왜 그래프가 직선이 되는가

- `x`값이 1만큼 늘어날 때마다 `y`값이 항상 `a`만큼 늘어나면, 그래프 위의 점들은 같은 기울기로 이어진다.
- 기울기가 일정하다는 것은 그래프가 꺾이지 않고 한 방향으로 뻗는다는 뜻이므로, 일차함수의 그래프는 직선이 된다.

### 핵심 정리 1: 일차함수의 변화율은 항상 일정하다

- 정리
  - 함수 $y=ax+b$의 임의의 두 점 $(x_1,y_1)$, $(x_2,y_2)$에 대해 평균변화율은 항상 $a$다.
- 증명
  - $y_1=ax_1+b$, $y_2=ax_2+b$이므로
    $$
    y_2-y_1=(ax_2+b)-(ax_1+b)=a(x_2-x_1)
    $$
    이다.
  - 따라서 $x_1\neq x_2$일 때
    $$
    \frac{y_2-y_1}{x_2-x_1}=a
    $$
    가 된다.
  - 즉, 어느 두 점을 잡아도 변화율은 항상 일정하다.

### 핵심 정리 2: 변화율이 일정하면 식은 $y=ax+b$ 꼴로 쓸 수 있다

- 정리
  - 어떤 함수의 변화율이 항상 일정한 값 `a`라면 그 함수는 $y=ax+b$ 꼴로 표현된다.
- 증명 스케치
  - 그래프 위의 한 점을 $(0,b)$라고 두자.
  - 변화율이 항상 `a`라는 것은 임의의 점 $(x,y)$에 대해
    $$
    \frac{y-b}{x-0}=a
    $$
    로 읽힌다는 뜻이다.
  - 이를 정리하면
    $$
    y-b=ax
    $$
    이고, 결국
    $$
    y=ax+b
    $$
    가 된다.
  - 그래서 `일정한 변화율`과 `일차함수`는 서로 같은 구조를 다른 말로 표현한 것이다.

### 그래프에서 읽어야 할 것

- $a>0$이면 오른쪽으로 갈수록 올라가는 직선이다.
- $a<0$이면 오른쪽으로 갈수록 내려가는 직선이다.
- $a=0$이면 $y=b$인 수평선이 된다.
- $b=0$이면 그래프가 원점을 지나므로 [proportion.md](./proportion.md)의 `정비례`와 연결된다.
- x절편은 $0=ax+b$를 풀어 $x=-\frac{b}{a}$로 구한다.

## Worked Examples

### 예제 1: 두 점을 지나는 일차함수 구하기

- 두 점 `(1,3)`, `(3,7)`을 지나면 기울기는
  $$
  a=\frac{7-3}{3-1}=2
  $$
  이다.
- 따라서 식은 먼저 $y=2x+b$ 꼴이다.
- 점 `(1,3)`을 대입하면
  $$
  3=2\cdot 1+b
  $$
  이므로 $b=1$이다.
- 따라서 일차함수는
  $$
  y=2x+1
  $$
  이다.

### 예제 2: 그래프 해석형 문제

- 함수 $y=-3x+6$에서
  - 기울기는 `-3`
  - y절편은 `6`
  - x절편은 $-3x+6=0$이므로 `2`
  이다.
- 즉 그래프는 `(0,6)`과 `(2,0)`을 지나는 내림차순 직선이다.

## Common Pitfalls

- `정비례`와 `일차함수`를 같은 것으로 보면 안 된다.
  - `정비례`는 $y=ax$이고, `일차함수`는 $y=ax+b$로 더 넓은 개념이다.
- 기울기를 $\frac{x_2-x_1}{y_2-y_1}$로 거꾸로 쓰는 실수가 자주 나온다.
- `b`를 단순 숫자로만 외우지 말고 `그래프의 시작 높이`로 읽어야 문장제 해석이 쉬워진다.
- x절편을 구할 때 $y=0$을 대입해야 하는데, y절편과 혼동하는 경우가 많다.

## Connections

- 선수 개념은 [proportion.md](./proportion.md), [linear-equation.md](./linear-equation.md), [function-foundations.md](./function-foundations.md)이다.
- 같은 축의 인접 개념으로는 [simultaneous-equations.md](./simultaneous-equations.md), [quadratic-function.md](./quadratic-function.md)가 있다.
- 다음 개념으로는 [equations-of-geometric-figures.md](./equations-of-geometric-figures.md), [differentiation.md](./differentiation.md)가 이어진다.
- 학년 허브에서는 [middle-2-hub.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/queries/math-curriculum-graph/middle-2-hub.md), [high-1-hub.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/queries/math-curriculum-graph/high-1-hub.md)와 연결된다.
- 계통 허브에서는 [functions-strand.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/queries/math-curriculum-graph/functions-strand.md)를 본다.

## Open Questions

- `기울기`와 `변화율`을 독립 카드로 분리할지 기준이 더 필요하다.
- `일차함수` 카드 안에 `교점과 연립방정식` 연결을 직접 둘지 별도 브리지 카드로 둘지 검토가 필요하다.

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
