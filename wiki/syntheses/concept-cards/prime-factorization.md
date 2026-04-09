---
title: 소인수분해
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
  - number
  - arithmetic
---

# 소인수분해

## Summary

소인수분해는 자연수를 더 이상 나눌 수 없는 소수들의 곱으로 나타내는 개념이다. 수의 구조를 가장 기본 단위까지 분해해 보는 연산이라서, 약수·배수·최대공약수·최소공배수의 출발점이 된다.

## Key Points

- 정의
  - 1보다 큰 자연수를 소수들의 곱으로 나타내는 과정을 소인수분해라 한다.
- 핵심 개념
  - 소수
  - 합성수
  - 약수
  - 배수
  - 거듭제곱 표기
- 대표 수식
  - $84=2^2\times 3\times 7$
  - $\gcd(a,b)\times \operatorname{lcm}(a,b)=ab$
- 증명 스케치
  - `증명 스케치 (추론)`:
  - 자연수는 합성수이면 더 작은 자연수들의 곱으로 나눌 수 있고, 이 과정을 반복하면 결국 소수들의 곱에 도달한다.
  - 만약 서로 다른 두 소인수분해가 있다면 소수의 나눗셈 성질 때문에 같은 소수들이 같은 횟수로 나타나야 한다.
  - 그래서 순서를 제외하면 소인수분해는 유일하다.
- 대표 예제
  - $60=2^2\times 3\times 5$이므로 $60$의 소인수는 `2,3,5`다.
- 교육과정 배치
  - 한국 대표 배치에서는 `중1` 초반의 핵심 수 개념이다.
- 국가별 배치 스냅샷
  - 한국: `중1`에서 최대공약수·최소공배수와 함께 배운다.
  - 일본: 중학교 초반의 수와 계산 축에서 정수 계산의 기반으로 다뤄진다.
  - 중국: `7학년`의 유리수와 문자식 전단계 감각으로 자연수 구조 학습에 놓인다.
  - 미국: Grade 6 이전 또는 Grade 6의 number sense 맥락에서 divisibility와 함께 반복 등장한다.
- 표현과 문제 감각
  - 다국어 용어: `prime factorization`, `素因数分解`, `素因数分解`
  - 계산 문제뿐 아니라 `왜 같은 수가 같은 소수들의 곱으로 정리되는가`를 이해하는 감각이 중요하다.

## Connections

- 선수 개념은 `소수`, `합성수`, `약수와 배수`다.
- 다음 개념으로는 [integers-and-rational-numbers.md](./integers-and-rational-numbers.md), `최대공약수`, `최소공배수`, [polynomials-and-factorization.md](./polynomials-and-factorization.md)가 이어진다.

## Open Questions

- `최대공약수`와 `최소공배수`를 독립 카드로 분리할지 후속 결정이 필요하다.

## Sources

- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-concept-encyclopedia/japan.md`
- `docs/math-concept-encyclopedia/china.md`
- `docs/math-concept-encyclopedia/us.md`
- `docs/math-concept-encyclopedia/multilingual-glossary.md`
