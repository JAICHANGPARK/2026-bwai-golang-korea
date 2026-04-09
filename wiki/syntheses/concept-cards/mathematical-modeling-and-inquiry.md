---
title: 수학적 모델링·수학 탐구
type: synthesis
status: active
updated: 2026-04-10
card_role: concept
source_docs:
  - docs/math-concept-encyclopedia/china.md
  - docs/math-curriculum-research/china.md
  - docs/math-concept-encyclopedia/unit-practice-book.md
tags:
  - concept
  - modeling
  - inquiry
---

# 수학적 모델링·수학 탐구

## Summary

수학적 모델링·수학 탐구는 현실 문제를 변수, 가정, 식, 자료로 번역해 해석하는 활동 중심 영역이다. 중국 문서군에서 특히 선명하게 드러나지만, 사실상 현대 수학 교육의 공통 확장 축으로 읽을 수 있다.

## Key Points

- 정의
  - 현실 문제를 수학 언어로 바꾸고, 모형을 세우고, 결과를 해석하는 활동을 수학적 모델링·수학 탐구라 한다.
- 핵심 개념
  - 변수 설정
  - 가정
  - 단순화
  - 모델 검증
  - 해석과 보고
- 대표 수식
  - 모델의 예: $\hat y=f(x)$
  - 오차의 예: $e=y-\hat y$
- 증명 스케치
  - `증명 스케치 (추론)`:
  - 좋은 모델은 현실을 완전히 복사하는 것이 아니라, 중요한 구조를 남기고 불필요한 요소를 줄이는 데서 출발한다.
  - 그래서 모델링의 핵심은 `무엇을 남기고 무엇을 버릴지`를 정하는 판단이다.
- 대표 예제
  - 택시 요금을 함수로 모델링하려면 기본요금, 거리, 시간 같은 변수를 먼저 정해야 한다.
- 교육과정 배치
  - 중국 고등학교 문서에서는 `수학적 모델링·수학 탐구`가 독립 활동 영역으로 분명히 보인다.
- 국가별 배치 스냅샷
  - 중국: 필수 과정과 선택 과정 모두에서 모델링 활동이 공식적으로 강조된다.
  - 미국: `Statistics`, `Bivariate Data`, advanced track 전반에 모델링 감각이 강하게 녹아 있다.
  - 한국: `직무 수학`, `경제 수학`, `인공지능 수학`, `실용 통계`에 응용적으로 확장된다.
  - 일본: 실생활 수학과 사회생활 맥락에서 간접적으로 이어진다.
- 표현과 문제 감각
  - 모델링 문제는 답 하나를 맞히는 것보다 `변수와 가정이 왜 타당한가`를 설명하는 힘이 중요하다.

## Deep Dive

모델링은 보통 `현실 문제 파악 -> 변수 설정 -> 가정 세우기 -> 식 또는 자료 구조 만들기 -> 계산 또는 시뮬레이션 -> 결과 해석 -> 모델 수정`의 순환 구조로 움직인다. 한 번에 완벽한 모델이 나오는 것이 아니라, 첫 모델을 세운 뒤 실제 상황과 비교하면서 다듬는 과정이 핵심이다.

좋은 모델은 복잡한 현실을 모두 담는 모델이 아니라, 지금 질문에 필요한 요소를 정확히 남기는 모델이다. 그래서 모델링에서는 `무엇을 버렸는가`를 숨기면 안 된다. 예를 들어 택시 요금 모델에서 교통체증, 할증 시간, 지역 차이를 모두 빼고 선형함수로 시작했다면, 그 단순화가 어디까지 유효한지 함께 말해야 한다.

## Worked Examples

- 예제 1: 택시 요금을 `기본요금 + 거리당 요금`으로 단순화하면
  $$
  y=ax+b
  $$
  꼴의 함수로 시작할 수 있다. 여기서 $b$는 기본요금, $a$는 거리 1단위당 증가하는 요금이다.
- 예제 2: 교실 온도의 변화를 시간에 따라 기록한 뒤 그래프로 그리면, 증가 구간과 감소 구간을 나누어 다른 함수나 추세선으로 설명할 수 있다.
- 예제 3: 매점 판매량을 날씨, 시간대, 가격과 연결해 보면 모든 변수를 한 번에 다루기 어렵다. 이때 핵심 변수 하나를 먼저 고르고 단순 모델을 세운 뒤 점차 보강하는 방식이 자연스럽다.

## Common Pitfalls

- 현실을 단순화했다는 사실을 잊고 모델을 현실 그 자체처럼 다루면 안 된다.
- 변수와 단위를 명확히 정하지 않으면 식은 맞아 보여도 해석이 틀어진다.
- 결과만 제시하고 가정의 타당성을 설명하지 않으면 모델링 문서로서 힘이 약하다.
- 자료 범위를 벗어난 예측을 무비판적으로 일반화하면 안 된다.

## Connections

- 선수 개념은 [function-foundations.md](./function-foundations.md), [data-organization.md](./data-organization.md), [regression.md](./regression.md)다.
- 다음 개념으로는 [economics-math.md](./economics-math.md), [mathematics-for-ai.md](./mathematics-for-ai.md), [workplace-math.md](./workplace-math.md), [practical-statistics.md](./practical-statistics.md)가 이어진다.
- 허브 관점에서는 [course-track-hub.md](../course-track-hub.md), [statistics-and-probability-strand.md](../statistics-and-probability-strand.md)와 함께 읽는다.

## Open Questions

- 모델링을 독립 `concept card`로 유지할지, 이후 `strategy card`나 `project card`로 더 분화할지 기준이 필요하다.

## Sources

- `docs/math-concept-encyclopedia/china.md`
- `docs/math-curriculum-research/china.md`
- `docs/math-concept-encyclopedia/unit-practice-book.md`
