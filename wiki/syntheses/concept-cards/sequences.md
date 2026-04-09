---
title: 수열
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
  - sequence
  - algebra
---

# 수열

## Summary

수열은 수들이 일정한 규칙으로 늘어선 구조를 다루는 개념이다. 함수의 이산 버전처럼 읽을 수 있어서, 점화식과 일반항, 극한 감각을 연결하는 고등 대수의 핵심 카드다.

## Key Points

- 정의
  - 자연수 순서에 따라 수를 대응시킨 나열을 수열이라 한다.
- 핵심 개념
  - 일반항
  - 점화식
  - 등차수열
  - 등비수열
  - 규칙성
- 대표 수식
  - $a_n=a_1+(n-1)d$
  - $a_n=a_1r^{n-1}$
- 증명 스케치
  - `증명 스케치 (추론)`:
  - 등차수열은 매번 같은 수 $d$를 더하므로 첫째항에서 $(n-1)$번 더한 값이 $n$번째 항이 된다.
  - 그래서 일반항은 $a_n=a_1+(n-1)d$가 된다.
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
- 다음 개념으로는 [limits.md](./limits.md), [differentiation.md](./differentiation.md), `수학적 귀납법`이 이어진다.

## Open Questions

- `수열의 합`을 별도 카드로 분리할지 추후 결정이 필요하다.

## Sources

- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-concept-encyclopedia/japan.md`
- `docs/math-concept-encyclopedia/china.md`
- `docs/math-concept-encyclopedia/us.md`
- `docs/math-concept-encyclopedia/multilingual-glossary.md`
