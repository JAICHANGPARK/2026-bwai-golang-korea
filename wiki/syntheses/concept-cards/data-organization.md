---
title: 자료 정리
type: synthesis
status: active
updated: 2026-04-10
card_role: concept
source_docs:
  - docs/math-concept-encyclopedia/korea.md
  - docs/math-concept-encyclopedia/japan.md
  - docs/math-concept-encyclopedia/china.md
  - docs/math-concept-encyclopedia/us.md
  - docs/math-concept-encyclopedia/multilingual-glossary.md
tags:
  - concept
  - statistics
  - data
---

# 자료 정리

## Summary

자료 정리는 데이터를 표와 그래프로 정리하고 대표값을 읽는 개념이다. 통계의 시작점으로, 수를 계산하는 것보다 `자료의 모습이 무엇을 말하는가`를 읽는 훈련에 가깝다.

## Key Points

- 정의
  - 자료를 모아 표와 그래프로 정리하고 평균, 중앙값, 최빈값 같은 대표값으로 특징을 읽는 과정을 자료 정리라 한다.
- 핵심 개념
  - 도수분포표
  - 히스토그램
  - 평균
  - 중앙값
  - 최빈값
  - 상대도수
- 대표 수식
  - $\bar x=\frac{x_1+\cdots+x_n}{n}$
- 증명 스케치
  - `증명 스케치 (추론)`:
  - 평균은 모든 자료값을 똑같이 반영한 대표값이고, 중앙값은 순서상 가운데 값을 나타낸다.
  - 따라서 같은 자료라도 어떤 대표값을 쓰느냐에 따라 해석이 달라질 수 있다.
- 대표 예제
  - 자료 `3,5,7,9`의 평균은 $6$이다.
  - 자료를 크기순으로 늘어놓았을 때 가운데 위치를 찾으면 중앙값을 읽을 수 있다.
- 교육과정 배치
  - 한국 대표 배치에서는 `중1`의 자료 정리 기초에서 시작한다.
- 국가별 배치 스냅샷
  - 한국: `중1`의 자료 정리 기초가 시작점이다.
  - 일본: 중1의 히스토그램·상대도수 감각, 중2의 상자수염그림으로 이어진다.
  - 중국: `7학년` 자료 수집과 `8학년` 데이터 분석이 연결된다.
  - 미국: Grade 6 Statistics가 대표적이다.
- 표현과 문제 감각
  - 다국어 용어: `data`, `mean`, `median`, `histogram`, `データ`, `ヒストグラム`, `数据`
  - 숫자 하나보다 분포의 모양을 함께 읽는 것이 중요하다.

## Deep Dive

대표값은 모두 `자료의 중심`을 말하지만 민감도가 다르다. 평균은 모든 값을 반영하므로 극단값의 영향을 크게 받는다. 반면 중앙값은 순서만 보므로 극단값이 있어도 비교적 안정적이다. 그래서 자료를 읽을 때는 평균 하나만 보는 대신 `평균과 중앙값이 얼마나 떨어져 있는지`, `그래프의 한쪽 꼬리가 길게 늘어졌는지`를 함께 봐야 한다.

도수분포표와 히스토그램은 자료의 개별 값을 모두 읽기 어려울 때 전체 모양을 보여 준다. 이때 계급의 폭을 어떻게 잡느냐에 따라 그래프 인상이 달라질 수 있으므로, 그래프는 `자료를 요약한 그림`이지 자료 자체와 완전히 같지 않다는 점을 기억해야 한다.

## Worked Examples

- 예제 1: 자료가 `2, 4, 4, 10`이면
  $$
  \bar x=\frac{2+4+4+10}{4}=5
  $$
  이다. 크기순으로 놓으면 가운데 두 수가 `4,4`이므로 중앙값은 `4`, 최빈값도 `4`다.
- 예제 2: 자료가 `1, 2, 2, 3, 20`이면 평균은
  $$
  \bar x=\frac{1+2+2+3+20}{5}=5.6
  $$
  이지만 중앙값은 `2`다. 큰 값 `20` 하나 때문에 평균이 중앙값보다 훨씬 커졌음을 읽어야 한다.
- 예제 3: 20명의 시험 점수를 `0~20`, `20~40`, `40~60`, `60~80`, `80~100` 구간으로 나눠 히스토그램을 그리면, 점수 하나하나보다 어느 구간에 학생이 많이 몰렸는지 빠르게 읽을 수 있다.

## Common Pitfalls

- 평균이 대표값이라고 해서 항상 가장 좋은 요약이라고 생각하면 안 된다. 극단값이 있으면 중앙값이 더 적절할 수 있다.
- 히스토그램 막대의 높이만 보고 전체 학생 수를 직접 읽으려 하면 안 된다. 계급과 도수를 함께 봐야 한다.
- 최빈값은 `가장 큰 값`이 아니라 `가장 자주 나타나는 값`이다.
- 그래프 모양만 보고 원자료의 정확한 값까지 안다고 생각하면 안 된다. 자료 정리는 요약이지 복사가 아니다.

## Connections

- 선수 개념은 [integers-and-rational-numbers.md](./integers-and-rational-numbers.md)다.
- 다음 개념으로는 [variance.md](./variance.md), [standard-deviation.md](./standard-deviation.md), [sampling.md](./sampling.md), [correlation-and-scatter-plots.md](./correlation-and-scatter-plots.md), [statistical-inference.md](./statistical-inference.md)가 이어진다.
- 계통 허브는 [statistics-and-probability-strand.md](../statistics-and-probability-strand.md)에서 다시 묶인다.

## Open Questions

- `상자그림`, `사분위범위` 같은 중등 자료 해석 카드를 [data-organization.md](./data-organization.md)에서 더 세분화할지 검토가 남아 있다.

## Sources

- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-concept-encyclopedia/japan.md`
- `docs/math-concept-encyclopedia/china.md`
- `docs/math-concept-encyclopedia/us.md`
- `docs/math-concept-encyclopedia/multilingual-glossary.md`
