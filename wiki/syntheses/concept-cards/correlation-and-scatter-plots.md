---
title: 상관과 산점도
type: synthesis
status: active
updated: 2026-04-10
card_role: concept
source_docs:
  - docs/math-concept-encyclopedia/korea.md
  - docs/math-concept-encyclopedia/china.md
  - docs/math-concept-encyclopedia/us.md
  - docs/math-concept-encyclopedia/multilingual-glossary.md
tags:
  - concept
  - statistics
  - correlation
---

# 상관과 산점도

## Summary

상관과 산점도는 두 변수가 함께 어떻게 변하는지를 시각적으로 읽는 개념이다. `함께 움직인다`는 사실과 `원인이다`는 판단을 구분하게 해 주는 데이터 해석의 핵심 카드다.

## Key Points

- 정의
  - 자료쌍 $(x_i,y_i)$를 점으로 찍어 관계를 보는 그림을 산점도라 하고, 그 경향을 상관이라 한다.
- 핵심 개념
  - 산점도
  - 양의 상관
  - 음의 상관
  - 무상관
  - 이상치
  - 경향선
- 대표 수식
  - 자료쌍 $(x_i,y_i)$
  - 경향선 $\hat y=ax+b$
- 성질 1
  - 점들이 오른쪽 위로 갈수록 함께 커지는 경향이 보이면 양의 상관으로 읽는다.
- 성질 2
  - 점들이 오른쪽 아래로 갈수록 한 변수가 커질 때 다른 변수는 작아지는 음의 상관으로 읽는다.
- 증명 스케치
  - `증명 스케치 (추론)`:
  - 산점도에서 점들이 오른쪽 위로 모이면 한 변수가 커질수록 다른 변수도 커지는 경향을 읽을 수 있다.
  - 하지만 점의 배치만으로 인과관계까지 증명되지는 않는다.
  - 이상치는 전체 경향선을 크게 흔들 수 있으므로, 상관 해석은 점의 모양과 함께 봐야 한다.
- 대표 예제
  - 공부 시간과 점수의 산점도가 오른쪽 위로 모이면 양의 상관 경향을 읽을 수 있다.
  - 키와 신발 사이즈의 산점도가 비교적 오른쪽 위로 늘어서면 양의 상관이 있으나, 한 변수로 다른 변수를 설명할 수 있다는 뜻은 아니다.
  - 기온과 아이스크림 판매량의 산점도가 오른쪽 위로 모일 수 있지만, 두 변수가 함께 변한다고 해서 반드시 한쪽이 다른 쪽의 원인이라고 말할 수는 없다.
## Deep Dive

- 산점도는 `점의 위치`가 곧 정보다.
- 점들이 한 줄에 가깝게 모일수록 선형 경향이 뚜렷하고, 퍼져 있을수록 관계가 약하다.
- 이상치 하나가 전체 경향선을 왜곡할 수 있으므로, 시각적 판단은 항상 자료 맥락과 함께 읽어야 한다.
- 상관은 방향과 경향의 세기이고, 인과는 원인과 결과의 주장이다. 둘은 같은 말이 아니다.

## Worked Examples

### 예제 1: 공부 시간과 성적

- 1일 공부 시간이 늘수록 점수가 대체로 오른다면 산점도는 오른쪽 위로 퍼진다.
- 이때 `공부 시간과 점수는 양의 상관이 있다`고 말할 수 있다.

### 예제 2: 기온과 따뜻한 음료 판매량

- 기온이 낮을수록 따뜻한 음료 판매량이 늘고, 기온이 높을수록 판매량이 줄면 산점도는 오른쪽 아래로 내려가는 모양을 보인다.
- 이때 `음의 상관`으로 읽되, 계절이나 시간대 같은 숨은 변수를 함께 점검해야 한다.

## Common Pitfalls

- 상관을 인과로 착각하면 안 된다.
- 산점도가 곡선 모양인데 직선 하나로만 해석하면 관계를 놓치기 쉽다.
- 이상치를 무시하면 경향을 잘못 읽을 수 있다.
- 상관계수는 산점도의 모양을 보조하는 값이지, 산점도를 대체하는 값이 아니다.
## Curriculum Context
- 교육과정 배치
  - 한국 대표 배치에서는 `중3`의 핵심 데이터 해석 단원이다.
- 국가별 배치 스냅샷
  - 한국: `중3`의 `상관과 산점도`가 명시적 단원이다.
  - 일본: 현재 문서군에서는 상자수염그림과 표본조사 축 안에서 분산된 데이터 해석으로 읽힌다.
  - 중국: `데이터 분석`과 고교 통계·회귀로 이어진다.
  - 미국: Grade 8의 `Bivariate Data`가 대표 단원이다.
- 표현과 문제 감각
  - 다국어 용어: `scatter plot`, `correlation`, `散布図`, `散点图`, `相关关系`
  - `함께 변한다`와 `원인이다`를 구분하는 것이 가장 중요하다.

## Connections

- 선수 개념은 [data-organization.md](./data-organization.md), [linear-function.md](./linear-function.md)다.
- 다음 개념으로는 [correlation-coefficient.md](./correlation-coefficient.md), [regression.md](./regression.md), [statistical-inference.md](./statistical-inference.md), [practical-statistics.md](./practical-statistics.md), [mathematics-for-ai.md](./mathematics-for-ai.md)가 이어진다.
- 계통 허브에서는 [statistics-and-probability-strand.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/syntheses/statistics-and-probability-strand.md)와 [probability-and-statistics-course.md](./probability-and-statistics-course.md)를 함께 본다.
- 학년 허브에서는 [middle-3-hub.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/syntheses/middle-3-hub.md), [high-2-hub.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/syntheses/high-2-hub.md)와 연결된다.

## Open Questions

- 산점도의 시각 해석과 상관계수의 수치 해석을 어느 수준까지 한 카드에서 병행할지 기준이 필요하다.

## Sources

- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-concept-encyclopedia/china.md`
- `docs/math-concept-encyclopedia/us.md`
- `docs/math-concept-encyclopedia/multilingual-glossary.md`
