---
title: 상관계수
type: synthesis
status: active
updated: 2026-04-10
card_role: concept
source_docs:
  - docs/math-concept-encyclopedia/korea.md
  - docs/math-concept-encyclopedia/us.md
  - docs/math-concept-encyclopedia/multilingual-glossary.md
  - docs/math-concept-encyclopedia/unit-practice-book.md
tags:
  - concept
  - statistics
  - correlation
---

# 상관계수

## Summary

상관계수는 두 변수의 선형 관계가 얼마나 강한지를 하나의 수로 요약하는 개념이다. 산점도에서 눈으로 읽던 경향을 수치로 정리해 회귀와 통계 해석으로 넘겨 주는 브리지 역할을 한다.

## Definition

- 두 변수의 선형 상관 정도를 `-1`과 `1` 사이의 수로 나타낸 값을 상관계수라 한다.
- 학교 통계에서는 보통 `산점도의 경향을 수치로 요약한 값`으로 이해하면 좋다.

## Key Ideas

- 방향
  - 양수면 양의 상관, 음수면 음의 상관으로 읽는다.
- 크기
  - 절댓값이 1에 가까울수록 선형 경향이 강하다.
- 선형 관계
  - 상관계수는 특히 `직선처럼 정렬되는 경향`을 읽는 수치다.
- 산점도와 함께 읽기
  - 상관계수 하나만 보지 말고 [correlation-and-scatter-plots.md](./correlation-and-scatter-plots.md)와 함께 봐야 한다.

## Deep Dive

- 상관계수는 편리하지만, 그림을 숫자 하나로 압축한 값이기 때문에 맹목적으로 믿으면 안 된다.
- 같은 상관계수라도 자료의 모양이 완전히 다를 수 있고, 이상치 하나가 값을 크게 흔들 수도 있다.

## Theorems and Properties

- 범위
  - 상관계수 $r$은 항상
    $$
    -1\le r\le 1
    $$
    을 만족한다.
- 부호의 의미
  - $r>0$이면 대체로 함께 증가하는 경향, $r<0$이면 한쪽이 증가할 때 다른 쪽이 감소하는 경향을 뜻한다.
- 해석의 한계
  - $r$이 0에 가깝다고 해서 관계가 전혀 없다는 뜻은 아니고, `선형 관계가 약하다`는 뜻에 가깝다.

## Proof Sketch

- `증명 스케치 (추론)`:
- 현재 원천 문서군은 `상관관계`를 해석 중심으로 제시하고, 상관계수 공식을 독립적으로 전개하지는 않는다.
- 이 카드는 그 해석을 수치화하는 학교 통계의 대표 관행을 정리한 확장 카드다.
- 핵심은 상관계수가 `산점도의 직선적 정렬 정도를 -1과 1 사이 값으로 압축한다`는 점이다.

## Worked Examples

- 예제 1: 양의 상관
  - 공부 시간과 시험 점수의 산점도가 오른쪽 위로 길게 모여 있고 $r=0.85$라면 강한 양의 선형 상관으로 읽는다.
- 예제 2: 음의 상관
  - 물건 가격과 판매량의 산점도가 오른쪽 아래로 길게 모여 있고 $r=-0.78$이라면 강한 음의 선형 상관으로 읽는다.
- 예제 3: 해석 주의
  - 키와 수학 점수 자료에서 $r=0.05$라면 `선형 상관이 매우 약하다`고 해석하는 편이 맞다.
  - 이것이 곧 `두 변수 사이에 어떤 관계도 없다`는 뜻은 아니다.

## Common Pitfalls

- 상관이 크다고 해서 인과관계가 있다고 결론내리면 안 된다.
- $r=0$에 가깝다고 해서 관계가 전혀 없다고 해석하면 안 된다.
- 상관계수만 보고 이상치의 영향을 무시하면 안 된다.
- 비선형 관계 자료를 상관계수 하나로만 판단하면 정보가 많이 사라진다.

## Curriculum Placement

- 한국 대표 문서에서는 `중3 상관과 산점도`와 `실용 통계` 사이의 확장 개념으로 읽는 것이 자연스럽다.
- 비교 메모
  - 미국 Statistics에서도 산점도와 회귀를 수치 요약과 함께 읽는 경우가 많다.
  - 현재 제공된 한·미 문서군은 해석의 중요성을 더 강조하므로, 이 카드는 `수치화된 해석` 레이어라고 보는 편이 적절하다.
- 표현 메모
  - 상관계수는 `인과관계의 증거`가 아니라 `선형 경향의 강도와 방향`을 요약한 값이다.

## Connections

- 선수 개념은 [correlation-and-scatter-plots.md](./correlation-and-scatter-plots.md), [data-organization.md](./data-organization.md)다.
- 다음 개념으로는 [regression.md](./regression.md), [practical-statistics.md](./practical-statistics.md), [statistics-and-data-analysis.md](./statistics-and-data-analysis.md)가 이어진다.
- 계통 허브로는 [../statistics-and-probability-strand.md](../statistics-and-probability-strand.md), [../high-3-hub.md](../high-3-hub.md)를 함께 본다.

## Open Questions

- 학교 수준에서 상관계수의 정확한 계산 공식을 이 카드에 둘지, 해석 중심 카드로 유지할지 기준이 더 필요하다.

## Sources

- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-concept-encyclopedia/us.md`
- `docs/math-concept-encyclopedia/multilingual-glossary.md`
- `docs/math-concept-encyclopedia/unit-practice-book.md`
