---
title: 복소수의 연산
type: synthesis
status: active
updated: 2026-04-09
source_docs:
  - docs/math-concept-encyclopedia/us.md
  - docs/math-concept-encyclopedia/china.md
  - docs/math-concept-encyclopedia/unit-practice-book.md
tags:
  - concept
  - algebra
  - complex-number
---

# 복소수의 연산

## Summary

복소수의 연산은 $a+bi$ 꼴의 수를 실수 계산 규칙 위에서 확장해 다루는 방법이다. [complex-numbers.md](./complex-numbers.md)가 `왜 복소수가 필요한가`를 보여 준다면, 이 카드는 `그 복소수를 실제로 어떻게 계산하는가`를 보여 주는 정석형 하위 카드다.

## Key Points

- 정의
  - 복소수 $z_1=a+bi$, $z_2=c+di$에 대해 덧셈, 뺄셈, 곱셈은 실수의 분배법칙과 $i^2=-1$을 이용해 계산한다.
- 핵심 개념
  - 실수부와 허수부
  - 동류항 정리
  - $i^2=-1$
  - 켤레복소수
  - 사칙연산의 확장
- 대표 수식
  - $(a+bi)+(c+di)=(a+c)+(b+d)i$
  - $(a+bi)(c+di)=(ac-bd)+(ad+bc)i$
  - $\overline{a+bi}=a-bi$
- 핵심 해석
  - 복소수 연산도 `실수부끼리`, `허수부끼리`, 그리고 $i^2=-1$ 규칙을 지키면 된다.
  - 곱셈은 단순 결합이 아니라 `실수부와 허수부가 섞여 새 항이 생긴다`는 점이 중요하다.
- 교육과정 배치
  - 미국 대표 경로에서는 `Algebra II`의 복소수 확장 안에서 읽는 것이 자연스럽다.
  - 중국 문서군에서는 고등학교 `기하와 대수`의 일부 맥락으로 이어지고, 일본에서는 `복소수평면`의 전단계 계산 감각으로 기능한다.
- 국가별 배치 스냅샷
  - 한국: 현재 문서군에서는 `이차방정식`의 실근 한계를 넘어가는 배경 지식으로 읽힌다.
  - 일본: `수학C`의 `복소수평면`으로 들어가기 전 대수적 계산 감각을 이룬다.
  - 중국: `기하와 대수`의 복소수 확장과 연결된다.
  - 미국: `Algebra II`에서 복소수 도입과 함께 계산 규칙을 익힌다.
- 표현과 문제 감각
  - 복소수 연산은 새 규칙을 많이 외우는 단원보다, `분배 -> $i^2=-1$ 정리 -> 실수부/허수부 묶기`의 흐름을 익히는 단원이다.

## Deep Dive

### 왜 덧셈은 성분별로 해도 되는가

- 복소수 $a+bi$는 `실수부 a`와 `허수부 b`를 가진 한 쌍으로 볼 수 있다.
- 따라서 같은 종류의 항끼리 묶으면
  $$
  (a+bi)+(c+di)=a+c+(b+d)i
  $$
  가 된다.
- 이는 [algebraic-manipulation.md](./algebraic-manipulation.md)의 동류항 정리 감각이 복소수로 확장된 것이다.

### 핵심 정리 1: 복소수 곱셈 공식

- 정리
  - 두 복소수의 곱은
    $$
    (a+bi)(c+di)=(ac-bd)+(ad+bc)i
    $$
    이다.
- 증명
  - 분배법칙으로 전개하면
    $$
    (a+bi)(c+di)=ac+adi+bci+bdi^2
    $$
    이다.
  - 여기서 $i^2=-1$이므로
    $$
    bdi^2=-bd
    $$
    이다.
  - 따라서 실수항과 허수항을 묶으면
    $$
    (ac-bd)+(ad+bc)i
    $$
    를 얻는다.

### 핵심 정리 2: 켤레를 곱하면 실수가 된다

- 정리
  - 복소수 $z=a+bi$에 대해
    $$
    z\overline z=(a+bi)(a-bi)=a^2+b^2
    $$
    이다.
- 증명
  - 전개하면
    $$
    (a+bi)(a-bi)=a^2-abi+abi-b^2i^2
    $$
    이다.
  - 가운데 두 항은 소거되고, $i^2=-1$이므로
    $$
    a^2-b^2(-1)=a^2+b^2
    $$
    이다.
  - 따라서 결과는 항상 실수다.

## Worked Examples

### 예제 1: 덧셈

- 복소수 $(4-2i)+(1+3i)$는
  $$
  (4+1)+(-2+3)i=5+i
  $$
  이다.

### 예제 2: 곱셈

- $(2+3i)(1-i)$를 계산하면
  $$
  2-2i+3i-3i^2
  $$
  이다.
- $i^2=-1$이므로
  $$
  2+i+3=5+i
  $$
  가 된다.

## Common Pitfalls

- $i^2=-1$인데 $i^2=1$처럼 계산하는 실수가 가장 흔하다.
- 덧셈에서는 실수부와 허수부를 따로 묶어야 하는데, 항을 뒤섞어 놓치기 쉽다.
- 곱셈에서 `FOIL`처럼 전개한 뒤 $i^2$ 정리를 빠뜨리면 답이 틀어진다.
- 켤레복소수 $a-bi$를 `-a+bi`로 착각하면 이후 계산이 전부 흔들린다.

## Connections

- 선수 개념은 [complex-numbers.md](./complex-numbers.md), [algebraic-manipulation.md](./algebraic-manipulation.md), [quadratic-equation.md](./quadratic-equation.md)다.
- 같은 축의 인접 개념으로는 [complex-division.md](./complex-division.md), [complex-plane.md](./complex-plane.md), [algebra-2.md](./algebra-2.md)가 있다.
- 다음 개념으로는 [complex-division.md](./complex-division.md), [complex-plane.md](./complex-plane.md), [polar-ideas.md](./polar-ideas.md)가 이어진다.
- 학년 허브에서는 [high-1-hub.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/queries/math-curriculum-graph/high-1-hub.md), [high-3-hub.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/queries/math-curriculum-graph/high-3-hub.md)와 연결된다.

## Open Questions

- `복소수 연산 -> 복소수평면의 회전 해석` 브리지 카드를 따로 둘지 검토가 필요하다.

## Sources

- `docs/math-concept-encyclopedia/us.md`
- `docs/math-concept-encyclopedia/china.md`
- `docs/math-concept-encyclopedia/unit-practice-book.md`
