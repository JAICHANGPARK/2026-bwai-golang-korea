---
title: 가설검정
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

# 가설검정

## Summary

가설검정은 표본에서 얻은 결과가 어떤 주장과 얼마나 어긋나는지를 판단하는 개념이다. 통계적 추정이 `어느 범위쯤일까`를 묻는다면, 가설검정은 `이 주장까지 받아들여도 될까`를 묻는 절차다.

## Key Points

- 정의
  - 모집단에 대한 가설을 세우고, 표본 결과가 그 가설과 얼마나 모순되는지 판단하는 절차를 가설검정이라 한다.
- 핵심 개념
  - 귀무가설 $H_0$
  - 대립가설 $H_1$
  - 유의수준 $\alpha$
  - 검정통계량
  - p값
- 대표 수식
  - 귀무가설과 대립가설의 기본 꼴:
    $$
    H_0:\text{현재 주장},\qquad H_1:\text{대안 주장}
    $$
  - 대표 판단 규칙:
    $$
    p\text{-value}<\alpha \Rightarrow H_0 \text{ 기각}
    $$
- 증명 스케치
  - `증명 스케치 (추론)`:
  - 가설검정의 발상은 `만약 현재 주장이 참이라면, 지금 본 표본 결과가 얼마나 드문가`를 보는 것이다.
  - 표본 결과가 너무 드물면, 현재 주장을 계속 유지하기 어렵다고 본다.
  - 학교 수준에서는 이 절차를 `모순의 크기를 재는 판단 규칙`으로 이해하는 것이 자연스럽다.

## Deep Dive

가설검정은 어떤 가설을 `증명`하는 절차가 아니다. 오히려 현재 주장을 잠정적으로 세워 두고, 표본이 그것과 너무 충돌하는지를 점검하는 과정에 가깝다.

그래서 `귀무가설을 기각하지 못했다`는 말은 `귀무가설이 참으로 증명되었다`는 뜻이 아니다. 단지 이번 표본만으로는 충분한 반례가 나오지 않았다는 뜻이다.

## Worked Examples

- 예제 1: 동전이 공정한지 의심한다고 하자. 귀무가설을 `앞면 확률이 0.5다`로 두고, 100번 던져 앞면이 70번 나왔다면 이 결과가 얼마나 드문지 따져 보게 된다.
- 예제 2: 새 학습법이 평균 점수를 올리는지 보려면, 먼저 `차이가 없다`를 귀무가설로 두고 표본 결과가 그 가설과 얼마나 어긋나는지 본다.
- 예제 3: 유의수준을 더 엄격하게 잡으면 같은 자료에서도 귀무가설을 기각하기가 더 어려워진다.

## Common Pitfalls

- `기각하지 못했다`를 `참으로 증명했다`로 읽으면 안 된다.
- 통계적으로 유의하다는 말과 실제 효과가 크다는 말을 같은 뜻으로 쓰면 안 된다.
- p값을 귀무가설이 참일 확률로 해석하면 안 된다.
- 표본 설계가 나쁘면 검정 절차가 정교해도 결론이 흔들린다.

## Curriculum Context

- 교육과정 배치
  - 현재 문서군에서는 독립 소단원보다 `통계적 추정`의 심화 후속 절차로 보는 것이 자연스럽다.
  - 미국 통계 과목과 일본의 `통계적 추측` 연장선에서 특히 잘 읽힌다.
- 국가별 배치 스냅샷
  - 한국: 현재 문서군에서는 `오차와 신뢰` 감각의 후속 심화 개념으로 간접 연결된다.
  - 일본: `표본조사 -> 통계적 추측 -> 신뢰구간/검정` 흐름의 마지막 판단 절차로 읽힌다.
  - 중국: 데이터 분석과 회귀·정규분포의 응용 심화로 이어질 수 있다.
  - 미국: high school Statistics와 AP Statistics에서 대표적인 inference 절차로 다뤄진다.

## Connections

- 선수 개념은 [statistical-inference.md](./statistical-inference.md), [confidence-interval.md](./confidence-interval.md), [conditional-probability.md](./conditional-probability.md)다.
- 다음 개념으로는 [practical-statistics.md](./practical-statistics.md), [mathematics-for-ai.md](./mathematics-for-ai.md)가 이어진다.
- 계통 허브에서는 [statistics-and-probability-strand.md](../statistics-and-probability-strand.md)를 함께 본다.

## Open Questions

- 학교 수준 위키에서 `p값`과 `검정통계량`의 계산을 어디까지 상세히 적을지 기준이 필요하다.

## Sources

- `docs/math-concept-encyclopedia/japan.md`
- `docs/math-concept-encyclopedia/us.md`
- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-curriculum-research/comparison.md`
