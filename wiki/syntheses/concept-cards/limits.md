---
title: 극한
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
  - calculus
  - limit
---

# 극한

## Summary

극한은 값이 어떤 대상에 점점 가까워질 때 그 끝 behavior를 읽는 개념이다. 무한 과정을 다루는 언어라서 미분과 적분, 연속성 전체의 출발점이 된다.

## Key Points

- 정의
  - 어떤 변수나 수열이 특정 값에 한없이 가까워질 때 함수값이나 항의 경향을 읽는 개념을 극한이라 한다.
- 핵심 개념
  - 한없이 가까워짐
  - 함수의 극한
  - 수열의 극한
  - 연속성
  - 무한 과정
- 대표 수식
  - $\lim_{x\to a}f(x)$
  - $\lim_{n\to\infty}a_n$
- 증명 스케치
  - `증명 스케치 (추론)`:
  - $\frac1n$은 $n$이 커질수록 양수이면서 계속 작아진다.
  - 원하는 만큼 작은 양수보다 더 작게 만들 수 있으므로 $\lim_{n\to\infty}\frac1n=0$이라 읽는다.
- 대표 예제
  - $\lim_{n\to\infty}\frac1n=0$이다.
- 교육과정 배치
  - 한국 대표 배치에서는 `미적분I`와 `미적분II`의 핵심 시작점이다.
- 국가별 배치 스냅샷
  - 한국: `함수의 극한과 연속`, `수열의 극한`으로 나누어 다룬다.
  - 일본: `수학III`의 극한이 대표 단원이다.
  - 중국: `수열과 도함수` 및 고등 미적분 전개에서 극한 감각이 사용된다.
  - 미국: Calculus의 맨 앞에서 limits and continuity로 다룬다.
- 표현과 문제 감각
  - 다국어 용어: `limit`, `極限`, `极限`
  - 극한은 `도착했는가`보다 `얼마나 가까워질 수 있는가`를 묻는 언어다.

## Connections

- 선수 개념은 [sequences.md](./sequences.md), [square-roots-and-real-numbers.md](./square-roots-and-real-numbers.md), [function-foundations.md](./function-foundations.md)다.
- 다음 개념으로는 [continuity.md](./continuity.md), [differentiation.md](./differentiation.md), [integration.md](./integration.md)이 이어진다.

## Open Questions

- `수열의 극한`과 `함수의 극한`을 별도 심화 카드로 더 분리할지 기준이 필요하다.

## Sources

- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-concept-encyclopedia/japan.md`
- `docs/math-concept-encyclopedia/china.md`
- `docs/math-concept-encyclopedia/us.md`
- `docs/math-concept-encyclopedia/multilingual-glossary.md`
