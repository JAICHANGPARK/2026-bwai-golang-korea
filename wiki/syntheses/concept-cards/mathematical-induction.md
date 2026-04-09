---
title: 수학적 귀납법
type: synthesis
status: active
updated: 2026-04-10
card_role: concept
source_docs:
  - docs/math-concept-encyclopedia/korea.md
  - docs/math-concept-encyclopedia/multilingual-glossary.md
tags:
  - concept
  - proof
  - induction
---

# 수학적 귀납법

## Summary

수학적 귀납법은 자연수 전체에 대한 명제를 단계적으로 증명하는 방법이다. 대수 과목에서 `패턴을 일반화해 엄밀한 주장으로 바꾸는 법`을 배우게 해 주는 핵심 증명 도구다.

## Key Points

- 정의
  - 어떤 명제가 첫 단계에서 참이고, $k$번째에서 참이면 $k+1$번째에서도 참임을 보이면 모든 자연수에 대해 참이라고 결론내리는 증명 방법이다.
  - 보통 시작 단계, 귀납 가정, 귀납 단계의 3단계로 쓴다.
- 핵심 개념
  - 시작 단계
  - 귀납 가정
  - 귀납 단계
  - 자연수 전체
  - 일반화
  - 명제 $P(n)$
  - 강한 귀납법
- 대표 성질
  - 귀납법은 자연수 전체의 무한한 명제를 유한한 두 단계로 바꾸는 논리다.
  - 시작 단계가 빠지면 첫 번째 도미노가 없어서 전체 결론이 성립하지 않는다.
  - 귀납 가정과 귀납 결론은 정확히 구분해야 하며, 가정에 쓰지 않은 내용을 결론에 새로 넣으면 안 된다.
- 대표 수식
  - $P(1)$이 참
  - $P(k)\Rightarrow P(k+1)$
  - $P(n)\Rightarrow P(n+1)$
- 증명 스케치
  - `증명 스케치 (추론)`:
  - 첫 번째 도미노가 넘어지고, `어느 하나가 넘어지면 다음 것도 넘어진다`가 보이면 모든 도미노가 넘어진다.
  - 수학적 귀납법은 이 발상을 자연수 명제 증명에 옮긴 것이다.
  - 귀납 단계에서는 $P(k)$를 가정한 뒤, 그 가정을 이용해 $P(k+1)$을 도출해야 한다.
## Worked Examples
  - 예제 1: $1+2+\cdots+n=\frac{n(n+1)}2$를 귀납법으로 보일 수 있다.
  - 예제 2: $2^n\ge n+1$ $(n\ge 0)$ 같은 명제도 귀납법으로 증명할 수 있다.
  - 예제 3: $1^3+2^3+\cdots+n^3=\left(\frac{n(n+1)}2\right)^2$도 귀납법의 대표 예다.
## Deep Dive
  - 귀납법의 핵심은 `시작점`과 `전달 규칙`이다.
  - 시작점은 첫 번째 항의 참을 보장하고, 전달 규칙은 하나 참이면 다음도 참임을 보장한다.
  - 강한 귀납법에서는 $k$번째뿐 아니라 그 이전의 여러 참을 함께 가정할 수 있어서, 분할이나 인수 분해가 섞인 명제에 유리하다.
  - 귀납법은 예 몇 개를 보는 실험이 아니라, 자연수 전체를 한 번에 묶는 논리 증명이다.
## Common Pitfalls
  - 시작 단계만 확인하고 귀납 단계를 생략하면 안 된다.
  - $P(k)$를 가정할 때 $P(k+1)$까지 미리 써 버리면 논리 순서가 무너진다.
  - 몇 개의 예를 확인했다고 해서 귀납법이 되는 것은 아니다.
  - 귀납법은 자연수 전체에 대한 방법이지 모든 정수에 대한 방법은 아니다.
- 교육과정 배치
  - 한국 대표 배치에서는 `대수`의 구성 요소다.
- 비교 메모
  - 현재 제공된 일본·중국·미국 문서군에서는 독립 소단원으로 강하게 드러나기보다 상위 대수와 증명 학습의 일부로 읽히는 편이다.
- 표현과 문제 감각
  - 다국어 용어: `mathematical induction`, `数学的帰納法`, `数学归纳法`
  - 핵심은 계산 요령보다 `한 단계에서 다음 단계로 이어지는 구조`를 보여 주는 데 있다.

## Connections

- 선수 개념은 [sequences.md](./sequences.md), [variables-and-expressions.md](./variables-and-expressions.md)다.
- 같은 축의 인접 개념으로는 [sequences-and-series.md](./sequences-and-series.md), [geometric-proof.md](./geometric-proof.md)가 있다.
- 같이 읽을 카드로는 [korean-algebra-course.md](./korean-algebra-course.md), [high-2-hub.md](../high-2-hub.md), [course-track-hub.md](../course-track-hub.md)가 있다.
- 수열과의 연결은 [functions-strand.md](../functions-strand.md)에서도 확인할 수 있다.

## Open Questions

- 귀납법 카드를 `수열`과 함께 더 강하게 묶을지, `증명 방법론` 축으로 독립시킬지 후속 기준이 필요하다.

## Sources

- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-concept-encyclopedia/multilingual-glossary.md`
