---
title: 조건부확률
type: synthesis
status: active
updated: 2026-04-09
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
  - docs/math-concept-encyclopedia/comparative-problem-book.md
tags:
  - concept
  - probability
  - statistics
---

# 조건부확률

## Summary

조건부확률은 어떤 사건 `B`가 이미 일어났다는 조건 아래에서 사건 `A`가 일어날 확률을 다시 계산하는 개념이다. 표본공간이 바뀐다는 감각을 익히게 해 주며, 독립성, 확률변수, 베이즈적 사고, 통계적 추정의 출발점이 된다.

## Key Points

- 정의
  - $B$가 일어났다고 알고 있을 때 사건 $A$가 일어날 확률을 $P(A\mid B)$로 쓴다.
- 핵심 개념
  - 조건
  - 표본공간의 축소
  - 사건의 교집합
  - 독립성
  - 확률 갱신
- 대표 수식
  - $P(A\mid B)=\frac{P(A\cap B)}{P(B)}\;(P(B)>0)$
  - $P(A\cap B)=P(B)P(A\mid B)$
- 증명 스케치
  - `증명 스케치 (추론)`:
  - $B$가 이미 일어났다고 알면, 원래 전체 표본공간이 아니라 `B 안`만 새로운 표본공간이 된다.
  - 이 안에서 $A$도 함께 일어나는 경우는 $A\cap B$에 해당하므로, 새로운 확률은
    $$
    \frac{\text{B 안에서 A도 일어나는 경우}}{\text{B의 전체 경우}}
    $$
    로 계산된다.
  - 이를 확률 기호로 쓰면
    $$
    P(A\mid B)=\frac{P(A\cap B)}{P(B)}
    $$
    이다.
- 대표 예제
  - 공정한 주사위를 던질 때 $A=\{\text{짝수}\}$, $B=\{\text{4보다 큼}\}$이라 하자.
  - 그러면 $B=\{5,6\}$이고 그 안에서 $A$도 만족하는 경우는 `{6}`뿐이므로
    $$
    P(A\mid B)=\frac12
    $$
    이다.
- 교육과정 배치
  - 한국 대표 배치에서는 고등학교 `확률과 통계`의 기초 개념으로 읽는 것이 자연스럽다.
  - 이후 [random-variable.md](./random-variable.md), [probability-distribution.md](./probability-distribution.md), [statistical-inference.md](./statistical-inference.md)로 이어진다.
- 국가별 배치 스냅샷
  - 한국: `확률과 통계`에서 독립사건, 확률분포, 추정과 함께 읽는 기초 개념이다.
  - 일본: 현재 문서군에서는 독립 표제로 강하게 전면화되기보다 `수학A`의 확률과 `수학B`의 통계적 추측을 잇는 중간 개념으로 보는 것이 자연스럽다.
  - 중국: 고등학교 선택필수의 `조건부확률, 이산확률변수, 정규분포` 축에서 명시적으로 등장한다.
  - 미국: `Statistics and Data Analysis`, `AP Statistics`, 고급 확률 경로에서 의미 해석과 함께 다뤄진다.
- 표현과 문제 감각
  - 다국어 용어: `conditional probability`, `条件付き確率`, `条件概率`
  - 비교 문제집 기준으로 중국은 공식 서술이 직접적이고, 미국은 의미 설명, 한국은 독립성과 함께, 일본은 현재 자료에서 간접적 연결로 읽힌다.

## Connections

- 선수 개념은 [counting-principles.md](./counting-principles.md)와 `확률`이다.
- 다음 개념으로는 [random-variable.md](./random-variable.md), [probability-distribution.md](./probability-distribution.md), [statistical-inference.md](./statistical-inference.md)가 이어진다.
- 계통 허브에서는 [statistics-and-probability-strand.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/queries/math-curriculum-graph/statistics-and-probability-strand.md)를 본다.

## Open Questions

- `독립사건`을 조건부확률 카드의 하위 개념으로 둘지 독립 카드로 뺄지 기준이 필요하다.
- 베이즈 정리를 이 카드에 포함할지 후속 심화 카드로 둘지 결정이 필요하다.

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
- `docs/math-concept-encyclopedia/comparative-problem-book.md`
