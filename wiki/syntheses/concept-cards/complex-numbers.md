---
title: 복소수
type: synthesis
status: active
updated: 2026-04-10
card_role: concept
source_docs:
  - docs/math-concept-encyclopedia/japan.md
  - docs/math-concept-encyclopedia/china.md
  - docs/math-concept-encyclopedia/us.md
  - docs/math-curriculum-research/japan.md
  - docs/math-curriculum-research/china.md
  - docs/math-curriculum-research/us.md
  - docs/math-concept-encyclopedia/unit-practice-book.md
tags:
  - concept
  - algebra
  - complex-number
---

# 복소수

## Summary

복소수는 실수 범위를 넘어 $i^2=-1$을 만족하는 허수를 도입해 수 체계를 확장한 개념이다. 이차방정식의 해를 더 넓게 받아들이게 해 주고, 대수와 기하를 연결하는 중요한 관문이 된다.

## Key Points

- 정의
  - $a,b$가 실수일 때 $a+bi$ 꼴로 나타내는 수를 복소수라 한다.
- 핵심 개념
  - 실수부
  - 허수부
  - 허수단위
  - 수 체계 확장
  - 방정식의 해
- 대표 수식
  - $z=a+bi$
  - $i^2=-1$
- 성질 1
  - 두 복소수 $a+bi$, $c+di$가 같으려면 실수부와 허수부가 각각 같아야 한다.
- 성질 2
  - 복소수의 덧셈은 실수부끼리, 허수부끼리 더해 계산한다.
- 성질 3
  - 켤레복소수 $\bar z=a-bi$를 곱하면
    $$
    z\bar z=a^2+b^2
    $$
    가 되어 실수값이 나온다.
- 증명 스케치
  - `증명 스케치 (추론)`:
  - 실수 범위에서는 $x^2+1=0$의 해가 없다.
  - 그래서 $i^2=-1$을 만족하는 새 기호 $i$를 도입하면 $x^2+1=0$의 해를 $\pm i$로 표현할 수 있다.
  - 복소수는 `기존 수 체계로는 풀리지 않던 문제를 표현 가능하게 만드는 확장`으로 이해할 수 있다.
  - $z=a+bi$와 $\bar z=a-bi$를 곱하면 허수부가 서로 지워져 $a^2+b^2$만 남는다.
- 대표 예제
  - 복소수 $3+4i$의 실수부는 `3`, 허수부는 `4`다.
  - $(2+3i)+(1-5i)=3-2i$이다.
  - $(1+i)(1-i)=2$이다.
- Deep Dive
  - 복소수는 `실수로는 닫히지 않는 방정식의 해를 다루기 위해 만든 수 체계`다.
  - 학교 수준에서는 사칙연산과 켤레복소수의 성질을 정확히 익히는 것이 핵심이다.
  - 복소수평면과 연결하면 대수와 기하가 한 번에 보인다.
- Worked Examples

### 예제 1: 덧셈

- $(3-2i)+(4+5i)=7+3i$이다.

### 예제 2: 곱셈

- $(2+i)(3-2i)=6-4i+3i-2i^2=8-i$이다.
- 여기서 $i^2=-1$을 반드시 써야 한다.

## Common Pitfalls

- $i^2=-1$을 $i=-1$로 착각하면 안 된다.
- 실수부와 허수부를 섞어 계산하면 안 된다.
- $\bar z$는 원래 수에서 허수부의 부호만 바꾼 복소수다.
- 복소수의 크기와 실수부를 같은 개념처럼 쓰면 안 된다.
- ## Curriculum Context
- 교육과정 배치
  - 일본 `수학C`, 중국 `기하와 대수`, 미국 `Algebra II`에서 직접적 배경이 드러난다.
- 국가별 배치 스냅샷
  - 한국: 현재 문서군에서는 `이차방정식`의 후속 배경으로 간접 언급된다.
  - 일본: `수학C`의 `복소수평면`으로 연결된다.
  - 중국: `기하와 대수`의 핵심 개념 중 하나다.
  - 미국: `Algebra II`에서 다항식, 지수·로그와 함께 확장된다.
- 표현과 문제 감각
  - 복소수는 `없는 해를 억지로 만든다`기보다 `해의 표현 공간을 넓힌다`는 감각이 중요하다.

## Connections

- 선수 개념은 [quadratic-equation.md](./quadratic-equation.md), [square-roots-and-real-numbers.md](./square-roots-and-real-numbers.md)다.
- 같은 축의 인접 개념으로는 [complex-number-operations.md](./complex-number-operations.md), [algebra-2.md](./algebra-2.md)가 있다.
- 다음 개념으로는 [complex-number-operations.md](./complex-number-operations.md), [complex-plane.md](./complex-plane.md), [vectors.md](./vectors.md), [conic-sections.md](./conic-sections.md)가 이어진다.
- 계통 허브에서는 [geometry-strand.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/syntheses/geometry-strand.md)와 [functions-strand.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/syntheses/functions-strand.md)를 함께 본다.

## Open Questions

- `복소수의 나눗셈`과 `켤레복소수의 활용`을 이 카드 계열 안에서 어디까지 확장할지 기준이 더 필요하다.

## Sources

- `docs/math-concept-encyclopedia/japan.md`
- `docs/math-concept-encyclopedia/china.md`
- `docs/math-concept-encyclopedia/us.md`
- `docs/math-curriculum-research/japan.md`
- `docs/math-curriculum-research/china.md`
- `docs/math-curriculum-research/us.md`
- `docs/math-concept-encyclopedia/unit-practice-book.md`
