---
title: 수학적 귀납법
type: synthesis
status: active
updated: 2026-04-09
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
- 핵심 개념
  - 시작 단계
  - 귀납 가정
  - 귀납 단계
  - 자연수 전체
  - 일반화
- 대표 수식
  - $P(1)$이 참
  - $P(k)\Rightarrow P(k+1)$
- 증명 스케치
  - `증명 스케치 (추론)`:
  - 첫 번째 도미노가 넘어지고, `어느 하나가 넘어지면 다음 것도 넘어진다`가 보이면 모든 도미노가 넘어진다.
  - 수학적 귀납법은 이 발상을 자연수 명제 증명에 옮긴 것이다.
- 대표 예제
  - $1+2+\cdots+n=\frac{n(n+1)}2$를 귀납법으로 보일 수 있다.
- 교육과정 배치
  - 한국 대표 배치에서는 `대수`의 구성 요소다.
- 비교 메모
  - 현재 제공된 일본·중국·미국 문서군에서는 독립 소단원으로 강하게 드러나기보다 상위 대수와 증명 학습의 일부로 읽히는 편이다.
- 표현과 문제 감각
  - 다국어 용어: `mathematical induction`, `数学的帰納法`, `数学归纳法`
  - 핵심은 계산 요령보다 `한 단계에서 다음 단계로 이어지는 구조`를 보여 주는 데 있다.

## Connections

- 선수 개념은 [sequences.md](./sequences.md), [variables-and-expressions.md](./variables-and-expressions.md)다.
- 같이 읽을 카드로는 [korean-algebra-course.md](./korean-algebra-course.md), [geometric-proof.md](./geometric-proof.md)가 있다.

## Open Questions

- 귀납법 카드를 `수열`과 함께 더 강하게 묶을지, `증명 방법론` 축으로 독립시킬지 후속 기준이 필요하다.

## Sources

- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-concept-encyclopedia/multilingual-glossary.md`
