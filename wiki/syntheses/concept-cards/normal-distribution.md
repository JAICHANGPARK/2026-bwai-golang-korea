---
title: 정규분포
type: synthesis
status: active
updated: 2026-04-10
card_role: concept
source_docs:
  - docs/math-concept-encyclopedia/china.md
  - docs/math-concept-encyclopedia/us.md
  - docs/math-concept-encyclopedia/multilingual-glossary.md
  - docs/math-concept-encyclopedia/formula-examples.md
tags:
  - concept
  - statistics
  - distribution
---

# 정규분포

## Summary

정규분포는 평균 근처에 자료가 많이 모이고 양쪽으로 대칭적으로 퍼지는 대표적인 확률분포다. 실제 데이터를 이상화해 읽는 통계 모델의 핵심 카드다.

## Key Points

- 정의
  - 평균을 중심으로 대칭적이고 종 모양으로 퍼지는 대표적인 연속 확률분포를 정규분포라 한다.
- 핵심 개념
  - 평균
  - 분산
  - 대칭성
  - 퍼짐
  - 모델링
- 대표 수식
  - $X\sim N(\mu,\sigma^2)$
  - $Z=\frac{X-\mu}{\sigma}$
- 증명 스케치
  - `증명 스케치 (추론)`:
  - 정규분포는 여러 작은 오차가 합쳐질 때 자주 나타나는 이상화 모델로 이해할 수 있다.
  - 학교 수준에서는 엄밀한 유도보다 `평균 근처 집중 + 양쪽 대칭`이라는 구조를 먼저 읽는 것이 자연스럽다.
- 대표 예제
  - 키 자료가 정규분포를 따른다고 가정하면 평균 $\mu$ 근처에 자료가 가장 많이 모인다.
- 교육과정 배치
  - 현재 제공된 한국 문서에서는 독립 표제로 강하지 않지만, 중국 선택필수와 미국 Statistics 맥락에서는 핵심 개념이다.
- 국가별 배치 스냅샷
  - 한국: 현재 문서군에서는 `확률과 통계`의 심화 분포 개념으로 간접적으로만 읽힌다.
  - 일본: 현재 문서군에서는 `통계적 추측` 중심이라 독립 표제가 약하다.
  - 중국: `확률변수·정규분포·회귀`가 명시적 선택필수 단원이다.
  - 미국: Statistics and Data Analysis에서 distributions의 핵심 예시로 읽힌다.
- 표현과 문제 감각
  - 다국어 용어: `normal distribution`, `正規分布`, `正态分布`
  - 실제 자료를 완벽히 복사한 것이 아니라 `경향을 설명하는 모델`이라는 점이 중요하다.

## Deep Dive

정규분포를 읽을 때 핵심은 `중심`과 `퍼짐`이다. 평균 $\mu$는 그래프의 중심을, 표준편차 $\sigma$는 얼마나 넓게 퍼져 있는지를 결정한다. 평균이 같아도 표준편차가 크면 그래프가 낮고 넓어지고, 표준편차가 작으면 그래프가 높고 좁아진다.

또한 정규분포는 연속확률분포이므로 한 점에서의 확률을 직접 말하기보다 `구간의 넓이`로 확률을 해석해야 한다. 예를 들어 $P(X=a)$보다는 $P(a \le X \le b)$ 같은 형태가 자연스럽다. 실제 수업에서는 표준화 $Z=\frac{X-\mu}{\sigma}$를 통해 서로 다른 정규분포를 공통 기준으로 비교하는 감각이 중요하다.

## Worked Examples

- 예제 1: $X\sim N(100, 15^2)$일 때 점수 $130$의 표준점수는
  $$
  Z=\frac{130-100}{15}=2
  $$
  이다. 즉 평균보다 표준편차 2개만큼 높은 위치다.
- 예제 2: 두 시험 A, B의 평균과 표준편차가 각각 $(70,10)$, $(80,5)$라고 하자. A에서 90점을 받은 학생의 표준점수는 $2$, B에서 90점을 받은 학생의 표준점수는 $2$이므로 상대적 위치는 같다고 해석할 수 있다.
- 예제 3: 키 자료가 정규분포를 따른다고 보면 평균 근처 구간에 학생이 많이 모이고, 평균에서 멀어질수록 학생 수가 줄어드는 경향을 설명할 수 있다.

## Common Pitfalls

- 정규분포를 모든 실제 자료의 `정답 모양`으로 생각하면 안 된다. 현실 데이터는 비대칭이거나 여러 봉우리를 가질 수도 있다.
- 평균과 중앙값, 최빈값이 모두 같다는 성질은 정규분포의 특징이지 모든 분포의 특징이 아니다.
- 연속분포에서 한 점의 확률을 넓이처럼 계산하지 않도록 주의해야 한다.
- 표준화는 값을 바꾸는 조작이 아니라 `비교 기준을 맞추는 과정`이라는 점을 놓치면 안 된다.

## Connections

- 선수 개념은 [probability-distribution.md](./probability-distribution.md), [random-variable.md](./random-variable.md), [statistical-inference.md](./statistical-inference.md)다.
- 다음 개념으로는 [regression.md](./regression.md), [practical-statistics.md](./practical-statistics.md), [mathematics-for-ai.md](./mathematics-for-ai.md)가 이어진다.
- 계통 허브는 [statistics-and-probability-strand.md](../statistics-and-probability-strand.md)에서 다시 볼 수 있다.

## Open Questions

- 한국 교육과정 축에서 `정규분포`를 독립 카드로 얼마나 강하게 둘지 더 검토가 필요하다.

## Sources

- `docs/math-concept-encyclopedia/china.md`
- `docs/math-concept-encyclopedia/us.md`
- `docs/math-concept-encyclopedia/multilingual-glossary.md`
- `docs/math-concept-encyclopedia/formula-examples.md`
