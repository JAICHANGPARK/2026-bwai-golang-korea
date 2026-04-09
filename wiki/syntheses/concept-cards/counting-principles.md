---
title: 경우의 수
type: synthesis
status: active
updated: 2026-04-10
card_role: concept
source_docs:
  - docs/math-concept-encyclopedia/korea.md
  - docs/math-concept-encyclopedia/japan.md
  - docs/math-concept-encyclopedia/us.md
  - docs/math-curriculum-research/korea.md
  - docs/math-curriculum-research/japan.md
  - docs/math-curriculum-research/china.md
  - docs/math-curriculum-research/us.md
  - docs/math-concept-encyclopedia/unit-practice-book.md
  - docs/math-concept-encyclopedia/comparative-problem-book.md
tags:
  - concept
  - probability
  - counting
---

# 경우의 수

## Summary

경우의 수는 가능한 결과를 빠짐없이 세는 개념이다. 확률을 정의하기 전의 바닥 공사 역할을 하며, 순열과 조합, 표본공간, 확률모형, 더 나아가 확률변수와 분포를 읽는 출발점이 된다. 학교 수준에서는 `덧셈의 원리`와 `곱셈의 원리`를 정확히 구분하는 것이 가장 중요하다.

## Key Points

- 정의
  - 어떤 조건을 만족하는 경우가 모두 몇 개인지 체계적으로 세는 개념을 경우의 수라 한다.
- 핵심 개념
  - 덧셈의 원리
  - 곱의 원리
  - 순열
  - 조합
  - 표본공간
  - 순서의 고려 여부
- 대표 수식
  - 서로 겹치지 않는 경우의 수는 더한다.
  - 독립적인 단계의 경우의 수는 곱한다.
  - ${}_nP_r=\frac{n!}{(n-r)!}$
  - ${}_nC_r=\frac{n!}{r!(n-r)!}$
- 정리 1: 덧셈의 원리
  - 서로 동시에 일어나지 않는 두 선택이 각각 $m$, $n$가지이면 전체 경우의 수는 $m+n$가지다.
- 정리 2: 곱의 원리
  - 첫 단계에 $m$가지, 둘째 단계에 $n$가지가 있으면 전체 경우의 수는 $mn$가지다.
- 증명 스케치
  - `증명 스케치 (추론)`:
  - 서로 다른 `n`개 중에서 순서를 고려해 `r`개를 뽑으면 첫 자리는 `n`가지, 둘째 자리는 `n-1`가지처럼 곱의 원리로
    $$
    n(n-1)\cdots(n-r+1)=\frac{n!}{(n-r)!}
    $$
    가 된다.
  - 조합은 같은 `r`개를 뽑아 놓고 순서만 다른 경우가 $r!$개씩 중복되므로
    $$
    {}_nC_r=\frac{{}_nP_r}{r!}=\frac{n!}{r!(n-r)!}
    $$
    로 정리된다.
- 대표 예제
  - 5명 중 회장과 부회장을 뽑는 경우의 수는 순서를 고려하므로
    $$
    {}_5P_2=5\cdot 4=20
    $$
    이다.
  - 5명 중 2명을 대표로 뽑는 경우의 수는 순서를 무시하므로
    $$
    {}_5C_2=10
    $$
    이다.
- 교육과정 배치
  - 한국 대표 배치에서는 `중2`의 확률 기초와 `고1 공통수학1`에서의 순열·조합으로 이어지는 축이다.
  - 이후 확률, 조건부확률, 확률변수와 분포의 언어가 된다.
## Deep Dive
  - 덧셈의 원리는 `A 또는 B`처럼 선택이 겹치지 않을 때 쓰고, 곱의 원리는 `한 단계씩 이어지는 선택`에 쓴다.
  - 조합은 `무엇을 골랐는가`만 세고, 순열은 `무엇을 어떤 순서로 놓았는가`까지 센다.
  - 문제를 잘못 읽으면 둘 사이를 뒤섞기 쉬우므로, 먼저 `순서가 중요한가`를 묻는 습관이 핵심이다.
## Worked Examples

### 예제 1: 메뉴 선택

- 메뉴가 `밥/면` 2가지이고, 각 종류마다 3가지 음료를 고르면 전체 조합은
  $$
  2\times 3=6
  $$
  가지다.

### 예제 2: 자리 배치

- 4명 중 2명을 뽑아 한 줄로 세우면
  $$
  {}_4P_2=4\cdot 3=12
  $$
  가지다.
## Common Pitfalls
  - 순서가 중요한데 조합으로 세면 답이 작아진다.
  - 서로 겹치는 경우를 덧셈의 원리로 바로 더하면 중복 계산이 생긴다.
  - `남는 사람`이나 `하지 않는 선택`을 세어야 할 때는 여사건으로 세는 편이 더 쉽다.
## Curriculum Context
- 국가별 배치 스냅샷
  - 한국: `중2`의 확률 기초와 `공통수학1`의 경우의 수 심화로 이어진다.
  - 일본: 중학교에서는 `경우의 수를 바탕으로 한 확률`로 먼저 나타나고, 고등학교 `수학A`에서 `순열·조합` 공식이 분명해진다.
  - 중국: 중학교에서는 확률 기초와 연결되고, 고등학교 선택필수에 `계수원리`가 분명히 등장한다.
  - 미국: 독립 대단원보다 `probability model`, `Statistics`, `Integrated Mathematics` 맥락 속에서 분산 배치되는 편이다.
- 표현과 문제 감각
  - 한국과 일본은 `순열/조합 공식`과 `어떻게 셌는가`를 함께 묻는 편이고, 중국은 계수원리와 확률 연결이 강하며, 미국은 경우의 수를 확률모형 속 설명으로 끌고 가는 문제가 많다.

## Connections

- 선수 개념은 `표`, `식`, `분류` 같은 기초 정리 감각이다.
- 다음 개념으로는 `확률`, [conditional-probability.md](./conditional-probability.md), [random-variable.md](./random-variable.md), [probability-distribution.md](./probability-distribution.md), [statistical-inference.md](./statistical-inference.md)가 이어진다.
- 학년 허브에서는 [middle-2-hub.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/syntheses/middle-2-hub.md), [high-1-hub.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/syntheses/high-1-hub.md), [high-2-hub.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/syntheses/high-2-hub.md)와 연결된다.
- 계통 허브에서는 [statistics-and-probability-strand.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/syntheses/statistics-and-probability-strand.md)와 [probability-and-statistics-course.md](./probability-and-statistics-course.md)를 함께 본다.

## Open Questions

- `경우의 수` 카드 안에 `순열`, `조합`, `곱의 원리`를 함께 둘지 각각 분리할지 기준이 더 필요하다.
- 미국식 `probability model` 표현을 이 카드에 얼마나 직접 반영할지 후속 조정이 필요하다.

## Sources

- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-concept-encyclopedia/japan.md`
- `docs/math-concept-encyclopedia/us.md`
- `docs/math-curriculum-research/korea.md`
- `docs/math-curriculum-research/japan.md`
- `docs/math-curriculum-research/china.md`
- `docs/math-curriculum-research/us.md`
- `docs/math-concept-encyclopedia/unit-practice-book.md`
- `docs/math-concept-encyclopedia/comparative-problem-book.md`
