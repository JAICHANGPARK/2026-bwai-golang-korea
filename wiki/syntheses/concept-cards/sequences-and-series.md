---
title: Sequences and Series
type: synthesis
status: active
updated: 2026-04-10
card_role: concept
source_docs:
  - docs/math-concept-encyclopedia/us.md
  - docs/math-curriculum-research/us.md
tags:
  - concept
  - sequence
  - precalculus
---

# Sequences and Series

## Summary

Sequences and Series는 항들의 규칙과 그 누적 합을 함께 다루는 개념 묶음이다. 미국 Precalculus에서 `수열`을 더 넓은 함수 언어와 연결해 미적분 전 단계로 넘겨 주는 역할을 한다.

## Key Points

- 정의
  - 순서대로 나열된 수의 규칙을 sequence라 하고, 그 항들을 더한 표현을 series라 한다.
  - sequence는 `항의 규칙`, series는 `그 항들의 합`이다.
- 핵심 개념
  - 일반항
  - 점화식
  - 등차수열
  - 등비수열
  - 합의 관점
  - 합 기호
  - 누적
- 대표 성질
  - 등차수열 합은 첫째항과 마지막 항의 평균에 항의 개수를 곱해 구할 수 있다.
  - 등비수열 합은 공비가 1이 아닐 때 표준 공식으로 정리할 수 있다.
  - series는 sequence의 `총합 버전`이 아니라, 수열을 더할 때 생기는 새로운 구조다.
  - 유한급수는 항의 개수가 정해져 있고, 무한급수는 극한과 연결된다.
- 대표 수식
  - $a_n=a_1r^{n-1}$
  - $\sum_{k=1}^n a_k$
  - $S_n=\frac{n(a_1+a_n)}{2}$
  - $S_n=a_1\frac{1-r^n}{1-r}$ $(r\neq 1)$
- 증명 스케치
  - `증명 스케치 (추론)`:
  - 수열이 `항 하나하나의 규칙`이라면, 급수는 그 규칙이 누적될 때 어떤 구조가 생기는지를 보는 관점이다.
  - 등차수열의 합 공식은 첫째항과 마지막 항을 짝지으면 같은 합이 반복된다는 사실에서 나온다.
  - 등비수열 합 공식은 $S_n-rS_n$으로 묶어 인접 항이 소거되는 구조에서 나온다.
  - 따라서 합 공식은 단순 암기가 아니라 `짝짓기`와 `소거`라는 구조적 아이디어로 이해하는 것이 좋다.
## Deep Dive

- sequence는 나열, series는 누적이다. 같은 항들을 보더라도 `한 항씩` 보는지 `전부 더한 값`을 보는지에 따라 관점이 달라진다.
- 등차수열 합은 앞뒤를 짝지으면 같은 값이 반복되고, 등비수열 합은 공비를 곱해 빼면 대부분의 항이 소거된다.
- 무한급수는 `더한 값이 수렴하는가`라는 극한 문제로 넘어가므로, 유한급수와 분리해서 읽어야 한다.

## Worked Examples

### 예제 1: 등차수열의 합

- 등차수열 $3,5,7,9$의 합은
  $$
  3+5+7+9=24
  $$
  이다.

### 예제 2: 등비수열의 합

- $1,2,4,8,\dots$의 처음 5개 항의 합은
  $$
  1+2+4+8+16=31
  $$
  이다.

### 예제 3: 합 공식 확인

- 첫째항이 $2$, 공차가 $3$인 등차수열의 처음 4개 항의 합은
  $$
  2+5+8+11=26
  $$
  이다.

## Common Pitfalls

- sequence와 series를 같은 말로 쓰면 안 된다.
- 합 공식에서 첫째항, 마지막 항, 항의 개수를 혼동하면 안 된다.
- 공비가 1일 때와 아닐 때 공식을 구분해야 한다.
- 무한급수와 유한급수를 구분하지 않으면 적용 범위를 벗어난다.
- 등비수열 합에서 공비가 0이나 1인 특수 경우를 확인해야 한다.
- 교육과정 배치
  - 미국 문서에서는 `Precalculus`의 핵심 개념 묶음으로 직접 언급된다.
- 비교 메모
  - 한국·일본·중국 문서군에서는 `수열`이 더 분명한 독립 표제이고, 미국은 `series`까지 묶어 Precalculus 단계에서 통합적으로 다루는 경향이 보인다.

## Connections

- 선수 개념은 [sequences.md](./sequences.md), [precalculus.md](./precalculus.md)다.
- 같은 축의 인접 개념으로는 [mathematical-induction.md](./mathematical-induction.md), [limits.md](./limits.md)가 있다.
- 다음 개념으로는 [calculus-course.md](./calculus-course.md), [calculus-2.md](./calculus-2.md)로 이어진다.
- 미국 course-track 관점은 [high-2-hub.md](../high-2-hub.md), [course-track-hub.md](../course-track-hub.md)에서 다시 본다.

## Open Questions

- `series`를 실제로 독립 카드로 더 세분화할지, 지금처럼 sequence 확장 카드로 둘지 기준이 필요하다.

## Sources

- `docs/math-concept-encyclopedia/us.md`
- `docs/math-curriculum-research/us.md`
