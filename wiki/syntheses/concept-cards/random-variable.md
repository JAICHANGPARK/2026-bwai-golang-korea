---
title: 확률변수
type: synthesis
status: active
updated: 2026-04-09
source_docs:
  - docs/math-concept-encyclopedia/korea.md
  - docs/math-concept-encyclopedia/china.md
  - docs/math-concept-encyclopedia/us.md
  - docs/math-concept-encyclopedia/japan.md
  - docs/math-curriculum-research/korea.md
  - docs/math-curriculum-research/china.md
  - docs/math-curriculum-research/us.md
  - docs/math-curriculum-research/japan.md
  - docs/math-concept-encyclopedia/multilingual-glossary.md
  - docs/math-concept-encyclopedia/formula-examples.md
  - docs/math-concept-encyclopedia/unit-practice-book.md
  - docs/math-concept-encyclopedia/comparative-problem-book.md
tags:
  - concept
  - probability
  - statistics
---

# 확률변수

## Summary

확률변수는 우연한 결과를 숫자로 바꾸어 다루는 개념이다. 경우의 수와 확률을 계산 단계에서 끝내지 않고, 기대값·분산·확률분포 같은 통계적 구조로 넘어가게 해 주는 핵심 번역 장치다.

## Key Points

- 정의
  - 표본공간의 각 결과에 실수를 대응시키는 규칙을 확률변수라 한다.
- 핵심 개념
  - 표본공간
  - 사건
  - 값의 대응
  - 기대값
  - 분산
  - 분포
- 대표 수식
  - $X:\Omega\to\mathbb{R}$
  - $P(X=x)$
  - $E(X)=\sum xP(X=x)$
- 증명 스케치
  - `증명 스케치 (추론)`:
  - 확률변수 $X$가 취하는 값마다 사건 $\{X=x\}$를 생각하면, 이 사건들은 서로 겹치지 않고 가능한 모든 결과를 빠짐없이 나눈다.
  - 따라서 각 값에 붙는 확률 $P(X=x)$는 모두 0 이상이고, 가능한 모든 값에 대한 합은 1이 된다.
  - 이 구조 위에서 기대값은 `결과값 x 가중치 확률`의 합으로 정의된다.
- 대표 예제
  - 주사위를 한 번 던질 때 눈의 수를 확률변수 $X$로 두면 $X$는 `1`부터 `6`까지를 각각 확률 $\frac16$으로 가진다.
  - 확률변수 $X$가 `1,2,3`을 각각 확률 `0.2, 0.5, 0.3`으로 가지면
    $$
    E(X)=1\cdot 0.2+2\cdot 0.5+3\cdot 0.3=2.1
    $$
    이다.
- 교육과정 배치
  - 한국 대표 배치에서는 `고2`의 `확률과 통계`에서 본격적으로 등장한다.
  - 이후 [probability-distribution.md](./probability-distribution.md), `통계적 추정`, `회귀`, `실용 통계`로 이어진다.
- 국가별 배치 스냅샷
  - 한국: `확률과 통계`에서 `기대값`, `분산`, `확률분포`와 함께 묶인다.
  - 일본: 문서상 `확률변수`가 독립 표제로 강하게 드러나기보다 `수학A`의 확률과 `수학B`의 통계적 추측 사이의 상위 통계 개념으로 읽힌다.
  - 중국: 고등학교 선택필수의 `확률변수·정규분포·회귀`에서 명시적으로 등장한다.
  - 미국: 보통 고등학교 `Statistics and Data Analysis`, `AP Statistics`, 또는 상위 경로에서 분포와 함께 다뤄진다.
- 표현과 문제 감각
  - 다국어 용어: `random variable`, `確率変数`, `随机变量`
  - 비교 문제집 기준으로 한국은 기대값 계산, 일본은 경우의 수 설명, 중국과 미국은 분포와 모델링 연결이 더 직접적이다.

## Connections

- 선수 개념은 [counting-principles.md](./counting-principles.md), `확률`, [conditional-probability.md](./conditional-probability.md)다.
- 바로 다음 개념은 [probability-distribution.md](./probability-distribution.md), [statistical-inference.md](./statistical-inference.md)다.
- 계통 허브에서는 [statistics-and-probability-strand.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/queries/math-curriculum-graph/statistics-and-probability-strand.md)를 본다.
- 과목 노드로는 [probability-and-statistics-course.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/queries/math-curriculum-graph/probability-and-statistics-course.md)와 연결된다.

## Open Questions

- `확률변수` 카드와 [probability-distribution.md](./probability-distribution.md)를 어느 정도까지 분리 유지할지 기준이 더 필요하다.
- 연속확률변수와 이산확률변수를 같은 카드에서 다룰지 후속 분화가 필요하다.

## Sources

- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-concept-encyclopedia/china.md`
- `docs/math-concept-encyclopedia/us.md`
- `docs/math-concept-encyclopedia/japan.md`
- `docs/math-curriculum-research/korea.md`
- `docs/math-curriculum-research/china.md`
- `docs/math-curriculum-research/us.md`
- `docs/math-curriculum-research/japan.md`
- `docs/math-concept-encyclopedia/multilingual-glossary.md`
- `docs/math-concept-encyclopedia/formula-examples.md`
- `docs/math-concept-encyclopedia/unit-practice-book.md`
- `docs/math-concept-encyclopedia/comparative-problem-book.md`
