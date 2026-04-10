---
title: 신뢰구간
type: synthesis
status: active
updated: 2026-04-10
card_role: concept
source_docs:
  - docs/math-concept-encyclopedia/japan.md
  - docs/math-concept-encyclopedia/us.md
  - docs/math-concept-encyclopedia/korea.md
  - docs/math-curriculum-research/comparison.md
tags:
  - concept
  - statistics
  - inference
---

# 신뢰구간

## Summary

신뢰구간은 표본에서 얻은 추정값 주위에 `모집단의 참값이 들어 있을 법한 구간`을 제시하는 개념이다. 추정을 한 점으로만 말하지 않고, 불확실성까지 함께 표현하게 해 주는 통계적 추정의 핵심 확장이다.

## Key Points

- 정의
  - 표본으로 계산한 추정값에 오차 범위를 붙여, 모집단의 참값이 들어 있을 것으로 보는 구간을 신뢰구간이라 한다.
- 핵심 개념
  - 점추정
  - 구간추정
  - 신뢰수준
  - 오차한계
  - 표본오차
- 대표 수식
  - 신뢰구간의 기본 형식:
    $$
    \text{추정값}\pm \text{오차한계}
    $$
- 증명 스케치
  - `증명 스케치 (추론)`:
  - 표본은 매번 조금씩 달라질 수 있으므로 추정값도 흔들린다.
  - 그래서 하나의 수만 제시하는 것보다, 그 수 주변의 흔들림 범위를 함께 주는 것이 더 자연스럽다.
  - 학교 수준에서는 엄밀한 유도보다 `표본이 바뀌면 구간도 바뀐다`는 감각을 먼저 익히는 것이 핵심이다.

## Deep Dive

신뢰구간은 `참값이 이미 이 구간 안에 있을 확률이 95%`라는 뜻으로 읽으면 곤란하다. 더 정확히는 `같은 방식으로 표본을 반복해서 뽑아 구간을 만들면, 그중 많은 비율이 참값을 포함하도록 설계되었다`는 의미에 가깝다.

또한 표본 수가 커지면 보통 구간이 더 좁아진다. 같은 신뢰수준이라면 정보가 많아질수록 추정의 흔들림이 줄어드는 것이 자연스럽기 때문이다.

## Worked Examples

- 예제 1: 한 학교에서 만족도 비율을 조사했더니 표본비율이 `0.60`이고 오차한계를 `0.05`로 잡았다고 하자. 그러면 신뢰구간은
  $$
  0.60\pm 0.05
  $$
  이므로 `[0.55,0.65]`로 읽을 수 있다.
- 예제 2: 두 조사 결과의 신뢰구간이 많이 겹친다면, 겉보기 평균 차이가 있어도 성급히 큰 차이라고 단정하기 어렵다.
- 예제 3: 같은 추정값이라도 표본 수가 더 큰 조사는 보통 더 좁은 신뢰구간을 준다.

## Common Pitfalls

- 신뢰구간을 `참값의 확률 구간`으로 기계적으로 말하면 안 된다.
- 구간이 좁다고 해서 무조건 좋은 조사라고 단정하면 안 된다. 편향 표본이면 좁아도 잘못될 수 있다.
- 점추정과 구간추정을 같은 말처럼 쓰면 안 된다.
- 신뢰수준이 높아지면 보통 구간이 넓어진다는 점을 놓치기 쉽다.

## Curriculum Context

- 교육과정 배치
  - 현재 문서군에서는 독립 대단원보다 `통계적 추정`의 심화 후속 개념으로 보이는 것이 자연스럽다.
  - 일본 문서에서는 `통계적 추측`의 연결 개념으로, 미국 문서에서는 Statistics와 Data Analysis의 해석 층으로 읽힌다.
- 국가별 배치 스냅샷
  - 한국: 현재 문서군에서는 `오차와 신뢰` 감각이 통계적 추정 카드 안에 간접적으로 들어 있다.
  - 일본: `표본조사 -> 통계적 추측 -> 신뢰구간`의 흐름이 연결로 나타난다.
  - 중국: 고교 선택필수의 확률변수·회귀 해석과 함께 심화 확통으로 이어질 수 있다.
  - 미국: high school Statistics와 AP Statistics에서 구간추정이 핵심 도구로 자리 잡는다.

## Connections

- 선수 개념은 [sampling.md](./sampling.md), [statistical-inference.md](./statistical-inference.md), [standard-deviation.md](./standard-deviation.md)다.
- 다음 개념으로는 [hypothesis-testing.md](./hypothesis-testing.md), [practical-statistics.md](./practical-statistics.md)가 이어진다.
- 계통 허브에서는 [statistics-and-probability-strand.md](../statistics-and-probability-strand.md)를 함께 본다.

## Open Questions

- 고교 수준 카드에서 `모평균`, `모비율`의 신뢰구간을 별도 예제로 더 분리할지 기준이 필요하다.

## Sources

- `docs/math-concept-encyclopedia/japan.md`
- `docs/math-concept-encyclopedia/us.md`
- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-curriculum-research/comparison.md`
