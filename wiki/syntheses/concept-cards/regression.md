---
title: 회귀
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
  - docs/math-concept-encyclopedia/formula-examples.md
  - docs/math-concept-encyclopedia/unit-practice-book.md
  - docs/math-concept-encyclopedia/comparative-problem-book.md
tags:
  - concept
  - statistics
  - modeling
---

# 회귀

## Summary

회귀는 두 변수 사이의 전체 경향을 식으로 요약해 예측에 활용하는 개념이다. 산점도의 점들을 하나하나 보는 수준을 넘어, `대체로 어떤 방향으로 얼마나 변하는가`를 읽게 해 주는 통계 모델링의 핵심 도구다.

## Key Points

- 정의
  - 자료의 경향을 가장 잘 설명하는 식이나 직선을 찾아 변수 사이의 관계를 요약하는 방법을 회귀라 한다.
- 핵심 개념
  - 산점도
  - 추세선
  - 설명변수와 반응변수
  - 예측값
  - 잔차
  - 경향 해석
- 대표 수식
  - $\hat y=ax+b$
  - $e_i=y_i-\hat y_i$
- 증명 스케치
  - `증명 스케치 (추론)`:
  - 각 자료점 $(x_i,y_i)$와 직선 $\hat y=ax+b$ 사이의 차이를 `잔차`라 두면
    $$
    e_i=y_i-(ax_i+b)
    $$
    로 쓸 수 있다.
  - 회귀직선은 보통
    $$
    S(a,b)=\sum_{i=1}^n \left(y_i-(ax_i+b)\right)^2
    $$
    를 가장 작게 하는 선으로 이해한다.
  - 학교 수준에서는 보통 `산점도의 전체 경향을 가장 잘 요약하는 직선`이라는 의미를 먼저 익히는 것이 자연스럽다.
- 대표 예제
  - 회귀직선이
    $$
    \hat y=2x+1
    $$
    이면 $x=5$일 때 예측값은
    $$
    \hat y=11
    $$
    이다.
- 교육과정 배치
  - 한국 대표 배치에서는 `중3`의 `상관과 산점도`가 전단계를 이루고, 고등학교 `실용 통계`에서 `회귀와 상관`이 명시적으로 다뤄진다.
  - 이후 `데이터과학`, `인공지능 수학`, `사회조사`, `비즈니스 분석`으로 이어진다.
- 국가별 배치 스냅샷
  - 한국: 중학교의 상관·산점도 감각 위에 `실용 통계`에서 회귀 해석이 명시적으로 확장된다.
  - 일본: 현재 문서군에서는 독립 대단원으로 강하게 전면화되기보다 `표본조사`와 `통계적 추측`의 응용 개념으로 보는 것이 자연스럽다.
  - 중국: 고등학교 선택필수의 `확률변수·정규분포·회귀`에서 명시적으로 등장한다.
  - 미국: `Statistics and Data Analysis`에서 `line of fit`, `regression`, `inference`가 함께 다뤄진다.
- 표현과 문제 감각
  - 다국어 용어: `regression`, `回帰`, `回归`
  - 비교 문제집 기준으로 미국은 회귀의 의미 해석 비중이 높고, 중국은 회귀직선 계산이 직접적이며, 한국은 상관과 함께 읽는 경향이 강하다.

## Deep Dive

회귀는 `점 하나를 정확히 맞히는 식`을 찾는 작업이 아니다. 자료 전체를 가장 잘 설명하는 경향선을 찾는 작업이다. 그래서 어떤 점은 회귀직선 위에 있고 어떤 점은 위나 아래로 벗어나게 된다. 이 차이를 잔차라고 하고, 잔차의 크기를 통해 모델이 자료를 얼마나 잘 설명하는지 본다.

회귀직선의 기울기 $a$는 `설명변수 $x$가 1 증가할 때 예측값 $\hat y$가 평균적으로 얼마나 변하는가`를 뜻한다. 절편 $b$는 $x=0$일 때의 예측값이지만, 실제 상황에서 $x=0$이 의미 없는 경우도 있으므로 무조건 현실 해석을 붙이면 안 된다.

## Worked Examples

- 예제 1: 회귀직선이
  $$
  \hat y=2x+1
  $$
  이고 실제 자료점이 $(5,12)$이면 예측값은 $11$, 잔차는
  $$
  e=12-11=1
  $$
  이다.
- 예제 2: 공부 시간 $x$와 시험 점수 $y$의 산점도에서 회귀직선의 기울기가 양수라면, 공부 시간이 늘수록 점수가 대체로 증가하는 경향이 있다고 해석할 수 있다.
- 예제 3: 회귀직선이 잘 맞는다고 해도, 자료 범위를 크게 벗어난 $x$값으로 예측하는 것은 조심해야 한다. 관측 구간 밖에서는 같은 경향이 유지된다고 보장할 수 없기 때문이다.

## Common Pitfalls

- 상관이 있다고 해서 인과관계가 있다고 단정하면 안 된다.
- 회귀직선의 절편을 현실 의미 없이 기계적으로 해석하면 오류가 생긴다.
- 산점도 범위를 많이 벗어난 예측값은 신뢰하기 어렵다.
- 회귀직선이 있다고 해서 모든 점이 그 선을 정확히 따라야 한다고 생각하면 안 된다.

## Connections

- 선수 개념은 [linear-function.md](./linear-function.md), [statistical-inference.md](./statistical-inference.md), [correlation-and-scatter-plots.md](./correlation-and-scatter-plots.md), [correlation-coefficient.md](./correlation-coefficient.md)다.
- 다음 개념으로는 [practical-statistics.md](./practical-statistics.md), [statistics-and-data-analysis.md](./statistics-and-data-analysis.md), [mathematics-for-ai.md](./mathematics-for-ai.md), `예측 모델링`이 이어진다.
- 학년 허브에서는 [high-2-hub.md](../high-2-hub.md), [high-3-hub.md](../high-3-hub.md)와 연결된다.
- 계통 허브에서는 [statistics-and-probability-strand.md](../statistics-and-probability-strand.md)와 [probability-and-statistics-course.md](./probability-and-statistics-course.md)를 본다.

## Open Questions

- 단순선형회귀와 다중회귀의 경계를 언제부터 나눌지 후속 설계가 필요하다.

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
- `docs/math-concept-encyclopedia/formula-examples.md`
- `docs/math-concept-encyclopedia/unit-practice-book.md`
- `docs/math-concept-encyclopedia/comparative-problem-book.md`
