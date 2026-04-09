---
title: 비례
type: synthesis
status: active
updated: 2026-04-10
card_role: concept
source_docs:
  - docs/math-concept-encyclopedia/korea.md
  - docs/math-concept-encyclopedia/japan.md
  - docs/math-concept-encyclopedia/china.md
  - docs/math-concept-encyclopedia/us.md
  - docs/math-curriculum-research/korea.md
  - docs/math-curriculum-research/japan.md
  - docs/math-curriculum-research/china.md
  - docs/math-curriculum-research/us.md
  - docs/math-concept-encyclopedia/multilingual-glossary.md
  - docs/math-concept-encyclopedia/unit-practice-book.md
tags:
  - concept
  - function
  - proportion
---

# 비례

## Summary

비례는 두 양이 함께 변하는 규칙을 식, 표, 그래프로 읽는 출발 개념이다. 정비례와 반비례를 함께 잡아 두면 `좌표와 비례 -> 일차함수 -> 함수 일반론`으로 이어지는 함수 축이 훨씬 선명해진다.

## Key Points

- 정의
  - 정비례는 두 양의 비 $\frac{y}{x}$가 일정한 관계다. 이때 $x \neq 0$이어야 한다.
  - 반비례는 두 양의 곱 $xy$가 일정한 관계다. 이때도 $x \neq 0$이어야 한다.
  - 학교 수학에서는 비례를 보통 `정비례`와 `반비례`를 묶어 읽는다.
- 핵심 개념
  - 비례상수
  - 좌표평면
  - 순서쌍
  - 단위비율
  - 원점 통과 직선
  - 쌍곡선 형태의 그래프
  - 식·표·그래프의 대응
  - 스케일링과 축척
- 대표 성질
  - 정비례 정리: $y=ax$이면 $x\neq 0$에서 $\frac{y}{x}=a$이고, 그래프는 원점을 지나는 직선이다.
  - 반비례 정리: $y=\frac{a}{x}$이면 항상 $xy=a$이고, 그래프는 두 갈래의 곡선이 된다.
  - 확대·축소 성질: 정비례에서는 $x$가 $k$배가 되면 $y$도 $k$배가 되고, 반비례에서는 $x$가 $k$배가 되면 $y$는 $\frac{1}{k}$배가 된다.
- 대표 수식
  - 정비례: $y=ax$
  - 반비례: $y=\frac{a}{x}$
  - 비례식: $\frac{y_1}{x_1}=\frac{y_2}{x_2}$
  - 반비례식: $x_1y_1=x_2y_2$
- 증명 스케치
  - `증명 스케치 (추론)`:
  - 정비례는 `비가 일정하다`를 식으로 옮기면 바로 $y=ax$가 된다.
  - 반대로 $y=ax$에서 $x\neq 0$이면 $\frac{y}{x}=a$이므로 비가 일정하다.
  - 반비례는 `곱이 일정하다`를 식으로 옮기면 바로 $y=\frac{a}{x}$가 된다.
  - 다만 $x=0$은 정의에서 제외되므로, 반비례 그래프를 그릴 때 0을 대입해 값을 찾으려 하면 안 된다.
## Worked Examples
  - 예제 1: $y=3x$에서 $x=4$이면 $y=12$이다. 또한 $x=10$이면 $y=30$이므로 비 $\frac{y}{x}=3$이 일정하다.
  - 예제 2: $y=\frac{12}{x}$에서 $x=3$이면 $y=4$이고, $x=6$이면 $y=2$이다. 두 경우 모두 곱은 $12$로 같다.
  - 예제 3: 한 물건의 가격이 개당 2500원일 때 총가격은 수량에 정비례한다. 수량이 5개면 총가격은 $12500$원이다.
  - 예제 4: 작업자 수가 늘수록 같은 일을 끝내는 시간은 반비례로 줄어드는 모형으로 볼 수 있다.
## Common Pitfalls
  - `비율이 같다`와 `비례한다`를 같은 말처럼 쓰면 안 된다. 비례는 두 양 사이의 함수적 관계다.
  - 반비례에서 $x=0$을 넣어 보려 하면 식이 성립하지 않는다.
  - 원점을 지나지 않는 직선을 모두 정비례로 착각하면 안 된다.
  - 정비례와 반비례의 그래프를 같은 종류의 직선으로 오해하면 안 된다.
- 교육과정 배치
  - 한국 대표 배치에서는 `중1`의 `좌표와 비례` 단원에 놓인다.
  - 이후 [linear-function.md](./linear-function.md)로 자연스럽게 확장된다.
- 국가별 배치 스냅샷
  - 한국: `중1`에서 좌표와 함께 `정비례·반비례`를 배운다.
  - 일본: `중1`의 `비례·반비례와 좌표`가 함수 학습의 출발점이다.
  - 중국: 문서상 `비례`가 독립 대단원으로 먼저 나오기보다 `좌표`, `일차함수`, `닮음`, `반비례함수`에 분산되어 배치된다.
  - 미국: `Grade 6`의 `Ratios and Rates`와 `Grade 7`의 `Proportional Relationships`가 별도 축으로 강하게 조직된다.
- 표현 스냅샷
  - 정비례: `direct proportion`, `比例`, `正比例`
  - 반비례: `inverse proportion`, `反比例`, `反比例`
  - 핵심 포인트는 표·식·그래프가 같은 관계를 다른 언어로 나타낸다는 점이다.

## Connections

- 선수 개념은 [integers-and-rational-numbers.md](./integers-and-rational-numbers.md), [variables-and-expressions.md](./variables-and-expressions.md)다.
- 같은 축의 인접 개념으로는 [inverse-proportion.md](./inverse-proportion.md), [linear-function.md](./linear-function.md), [similarity.md](./similarity.md)가 있다.
- 함수 계통은 [functions-strand.md](../functions-strand.md)에서 다시 묶인다.
- 학년 허브는 [middle-1-hub.md](../middle-1-hub.md)와 [korea-curriculum-hub.md](../korea-curriculum-hub.md)로 이어진다.
- 과목형 관점은 [course-track-hub.md](../course-track-hub.md)에서 함께 본다.

## Open Questions

- `정비례`와 `반비례`를 한 카드에 유지할지, 두 개념 카드로 쪼갤지 세분화 기준이 필요하다.
- `비례식`과 `비례관계`를 같은 카드에서 다룰지 용어 차이를 더 분리할지 검토가 필요하다.

## Sources

- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-concept-encyclopedia/japan.md`
- `docs/math-concept-encyclopedia/china.md`
- `docs/math-concept-encyclopedia/us.md`
- `docs/math-curriculum-research/korea.md`
- `docs/math-curriculum-research/japan.md`
- `docs/math-curriculum-research/china.md`
- `docs/math-curriculum-research/us.md`
- `docs/math-concept-encyclopedia/multilingual-glossary.md`
- `docs/math-concept-encyclopedia/unit-practice-book.md`
