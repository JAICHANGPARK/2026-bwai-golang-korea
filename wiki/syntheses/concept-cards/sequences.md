---
title: 수열
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
  - sequence
  - algebra
---

# 수열

## Summary

수열은 수들이 일정한 규칙으로 늘어선 구조를 다루는 개념이다. 함수의 이산 버전처럼 읽을 수 있어서, 점화식과 일반항, 극한 감각을 연결하는 고등 대수의 핵심 카드다.

## Key Points

- 정의
  - 자연수 $n$에 따라 수 $a_n$을 대응시키는 나열을 수열이라 한다.
  - 수열은 `정의역이 자연수인 함수`로 읽을 수 있다.
- 핵심 개념
  - 항
  - 인덱스
  - 일반항
  - 점화식
  - 등차수열
  - 등비수열
  - 규칙성
- 대표 성질
  - 등차수열은 인접한 두 항의 차가 항상 일정하다.
  - 등비수열은 인접한 두 항의 비가 항상 일정하다.
  - 수열은 항마다 규칙이 정해지면 정의되므로, 같은 규칙을 적는 일반항과 점화식이 서로 왕복된다.
  - 첫째항과 점화식이 주어지면 수열은 보통 유일하게 정해진다.
- 대표 수식
  - $a_n=a_1+(n-1)d$
  - $a_n=a_1r^{n-1}$
  - $a_{n+1}=a_n+d$
  - $a_{n+1}=ra_n$
- 증명 스케치
  - `증명 스케치 (추론)`:
  - 등차수열은 매번 같은 수 $d$를 더하므로 첫째항에서 $(n-1)$번 더한 값이 $n$번째 항이 된다.
  - 그래서 일반항은 $a_n=a_1+(n-1)d$가 된다.
  - 등비수열은 매번 같은 수 $r$을 곱하므로 첫째항에서 $(n-1)$번 곱한 값이 $n$번째 항이 된다.
  - 일반항은 점화식의 반복을 한 번에 적는 표현이고, 점화식은 생성 규칙을 직접 적는 표현이다.
## Deep Dive

- 수열은 이산적 함수이므로 그래프에서 보면 점들의 나열로 읽힌다.
- 등차수열은 선형적 패턴, 등비수열은 지수적 패턴을 대표한다.
- 점화식은 생성 규칙을 드러내고, 일반항은 임의의 항을 직접 계산하게 해 준다.
- 수열의 해석에서는 `몇 번째 항인가`가 식의 일부이므로 인덱스를 반드시 챙겨야 한다.

## Worked Examples

### 예제 1: 등차수열

- 첫째항이 $4$, 공차가 $3$인 등차수열의 8항은
  $$
  a_8=4+7\times 3=25
  $$
  다.

### 예제 2: 등비수열

- 첫째항이 $3$, 공비가 $2$인 등비수열의 5항은
  $$
  a_5=3\cdot 2^4=48
  $$
  이다.

### 예제 3: 점화식

- 점화식 $a_{n+1}=a_n+2$, $a_1=1$이면
  $$
  a_2=3,\ a_3=5,\ a_4=7
  $$
  이다.

## Common Pitfalls

- 공차와 공비를 섞어 쓰면 안 된다.
- $n$번째 항과 $n+1$번째 항을 혼동하면 점화식이 꼬인다.
- 수열은 숫자들의 모음이 아니라 `순서가 있는 규칙`이다.
- 일반항은 맞아도 시작항을 잘못 쓰면 전체 수열이 달라진다.
- 항의 번호를 0부터 시작하는지 1부터 시작하는지 반드시 확인해야 한다.
- 대표 예제
  - 첫째항이 `4`, 공차가 `3`인 등차수열의 8항은 $4+7\times 3=25$다.
- 교육과정 배치
  - 한국 대표 배치에서는 고등학교 `대수`에서 다룬다.
- 국가별 배치 스냅샷
  - 한국: `대수`의 핵심 구성요소다.
  - 일본: `수학B`에서 수열이 명시적으로 배치된다.
  - 중국: 선택필수의 `수열과 도함수`에서 다룬다.
  - 미국: Algebra II와 Precalculus에서 sequences and series의 일부로 배운다.
- 표현과 문제 감각
  - 다국어 용어: `sequence`, `数列`
  - 단순 나열보다 `어떤 규칙이 숨어 있는가`를 함수처럼 읽는 것이 중요하다.

## Connections

- 선수 개념은 [function-foundations.md](./function-foundations.md), [polynomials-and-factorization.md](./polynomials-and-factorization.md)다.
- 같은 축의 인접 개념으로는 [sequences-and-series.md](./sequences-and-series.md), [mathematical-induction.md](./mathematical-induction.md)가 있다.
- 다음 개념으로는 [limits.md](./limits.md), [differentiation.md](./differentiation.md)가 이어진다.
- 함수 계통과 고교 허브는 [functions-strand.md](../functions-strand.md), [high-2-hub.md](../high-2-hub.md), [korean-algebra-course.md](./korean-algebra-course.md)에서 다시 묶인다.

## Open Questions

- `수열의 합`을 별도 카드로 분리할지 추후 결정이 필요하다.

## Sources

- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-concept-encyclopedia/japan.md`
- `docs/math-concept-encyclopedia/china.md`
- `docs/math-concept-encyclopedia/us.md`
- `docs/math-concept-encyclopedia/multilingual-glossary.md`
