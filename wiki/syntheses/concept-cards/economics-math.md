---
title: 경제 수학
type: synthesis
status: active
updated: 2026-04-10
card_role: concept
source_docs:
  - docs/math-concept-encyclopedia/korea.md
  - docs/math-concept-encyclopedia/multilingual-glossary.md
tags:
  - concept
  - applied
  - economics
---

# 경제 수학

## Summary

경제 수학은 비율, 이자, 현재가치, 수요·공급 같은 경제 상황을 수학식으로 읽는 과목형 카드다. 함수와 증감률, 최적화가 실제 판단 문제로 연결되는 응용 축이다.

## Key Points

- 정의
  - 경제 상황을 비율, 함수, 수열, 최적화 언어로 해석하는 수학 영역이다.
- 핵심 개념
  - 환율
  - 세금
  - 단리와 복리
  - 현재가치
  - 수요·공급
  - 최적화
- 대표 수식
  - $A=P(1+rn)$
  - $A=P(1+r)^n$
  - $PV=\frac{FV}{(1+r)^n}$
- 대표 예제
  - 원금 `100만 원`, 연이율 `5%`, 2년 복리 후 금액은 $1000000\times 1.05^2$이다.
- 교육과정 배치
  - 한국 대표 배치에서는 선택 과목 `경제 수학`으로 존재한다.
- 국가별 배치 스냅샷
  - 한국: 독립 과목으로 편성된다.
  - 일본: 사회생활과 수학, 금융 수학 맥락에 가깝다.
  - 중국: 경제 모델링과 응용 수학 과목군 안에 가깝다.
  - 미국: financial mathematics, business math, quantitative literacy 과목과 가깝다.

## Deep Dive

경제 수학의 핵심은 현실 문장을 식으로 바꾸는 힘이다. `매달 일정 금액을 저축한다`, `가격이 오르면 수요가 줄어든다`, `비용과 수익이 만나는 지점이 손익분기점이다` 같은 문장을 함수나 수열로 번역해야 한다. 그래서 경제 수학은 계산 과목이면서 동시에 모델링 과목이다.

특히 단리와 복리의 차이는 시간이 길어질수록 크게 벌어진다. 단리는 원금에만 이자가 붙지만, 복리는 이전에 붙은 이자에도 다시 이자가 붙는다. 그래서 같은 이율이라도 장기적으로는 복리 모형이 훨씬 빠르게 증가한다.

## Worked Examples

- 예제 1: 원금이 100만 원, 연이율이 5%인 복리 상품에 2년 동안 예치하면
  $$
  A=1000000(1.05)^2=1102500
  $$
  원이다.
- 예제 2: 상품의 수요 함수가 $q=120-3p$라면 가격 $p$가 10 오를 때 수요량은 30 줄어든다. 기울기 `-3`은 가격과 수요가 반대로 움직인다는 뜻이다.
- 예제 3: 수익 함수가 $R(x)=50x$, 비용 함수가 $C(x)=20x+300$이면 손익분기점은
  $$
  50x=20x+300
  $$
  에서 $x=10$이다.

## Common Pitfalls

- 단리와 복리 공식을 섞어 쓰면 장기 계산에서 큰 차이가 난다.
- 절편이나 기울기를 현실 의미 없이 읽으면 모델 해석이 어긋난다.
- 경제 모델은 현실을 단순화한 것이므로 세금, 할인, 고정비 변화 같은 조건을 빼먹지 않아야 한다.
- 계산 결과가 맞아도 단위가 `원`, `개`, `년` 중 무엇인지 확인하지 않으면 실수가 생긴다.

## Connections

- 선수 개념은 [exponential-and-logarithmic-functions.md](./exponential-and-logarithmic-functions.md), [sequences.md](./sequences.md), [differentiation.md](./differentiation.md)다.
- 다음 개념으로는 `금융 이해력`, `경제 모델링`, `최적화`가 이어진다.
- 응용 허브는 [course-track-hub.md](../course-track-hub.md), [mathematical-modeling-and-inquiry.md](./mathematical-modeling-and-inquiry.md)와 연결된다.

## Open Questions

- `경제 수학`을 과목 카드와 세부 개념 카드로 더 분화할지 검토가 필요하다.
- `단리·복리`, `손익분기점`, `현재가치`를 별도 하위 카드로 승격할지 기준이 더 필요하다.

## Sources

- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-concept-encyclopedia/multilingual-glossary.md`
