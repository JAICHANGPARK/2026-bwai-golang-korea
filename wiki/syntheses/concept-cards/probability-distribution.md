---
title: 확률분포
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
  - docs/math-concept-encyclopedia/formula-examples.md
  - docs/math-concept-encyclopedia/unit-practice-book.md
  - docs/math-concept-encyclopedia/comparative-problem-book.md
tags:
  - concept
  - statistics
  - probability
---

# 확률분포

## Summary

확률분포는 확률변수가 어떤 값을 얼마나 자주 가질 수 있는지를 수로 정리한 구조다. 중학교의 확률과 자료 해석이 고등학교 `확률과 통계`에서 확률변수, 기대값, 분산으로 확장될 때 핵심이 되는 개념이다.

## Key Points

- 정의
  - 확률변수 `X`가 가질 수 있는 값과 각 값에 대응하는 확률을 정리한 것이 확률분포다.
- 핵심 개념
  - 사건
  - 확률변수
  - 기대값
  - 분산
  - 표본과 모집단
- 대표 수식
  - $P(A\cup B)=P(A)+P(B)-P(A\cap B)$
  - $P(X=x)$
  - $E(X)=\sum xP(X=x)$
  - $\operatorname{Var}(X)=E[(X-E(X))^2]$
- 증명 스케치
  - `증명 스케치 (추론)`:
  - 확률분포의 각 확률은 서로 배반인 경우들을 나눈 값이므로 모두 0 이상이어야 한다.
  - 가능한 모든 경우를 빠짐없이 합치면 전체 사건이 되므로 각 확률의 합은 1이 된다.
  - 기대값은 각 결과값에 그 결과가 나올 가능성을 가중치로 준 평균으로 해석할 수 있다.
- 대표 예제
  - 주사위를 한 번 던질 때 확률변수 $X$를 눈의 수로 두면 $X$는 `1`부터 `6`까지를 각각 확률 $\frac16$으로 갖는다.
  - 확률변수 $X$가 `1,2,3`을 각각 확률 `0.2,0.5,0.3`으로 가지면
    $$
    E(X)=1\cdot 0.2+2\cdot 0.5+3\cdot 0.3=2.1
    $$
    이다.
- 교육과정 배치
  - 한국 대표 배치에서는 `고2`의 `확률과 통계` 핵심 개념으로 읽는 것이 자연스럽다.
  - 선수 개념은 중학교의 `확률`, `산포`, `상관`과 고1의 `경우의 수`다.
- 국가별 배치 스냅샷
  - 한국: `고2`의 `확률과 통계`에서 `확률변수`, `기대값`, `분산`과 함께 읽는 것이 자연스럽다.
  - 일본: 중학교에서는 경험적 확률과 경우의 수를 먼저 다루고, 고등학교 `수학A`의 확률과 `수학B`의 통계적 추측이 분포 개념의 전단계를 이룬다.
  - 중국: `9학년`의 `확률의 기초` 뒤에 고등학교 `확률과 통계`, 그리고 선택필수의 `확률변수·정규분포·회귀`로 확장된다.
  - 미국: `Grade 6~8`에서 표본과 확률모형을 먼저 배우고, 고등학교 `Statistics and Data Analysis` 또는 `AP Statistics`에서 `distribution`, `inference`, `regression`, `probability`가 한 과목으로 묶인다.
- 표현과 문제 감각
  - 다국어 용어: `probability distribution`, `確率分布`, `概率分布`
  - 비교 문제집 기준으로 중국 고교와 미국 Statistics는 분포와 모델링이 더 직접적이고, 일본은 경우의 수 설명, 한국은 기대값 계산 비중이 크다.

## Connections

- 선수 개념은 [counting-principles.md](./counting-principles.md), `확률`, [conditional-probability.md](./conditional-probability.md), [random-variable.md](./random-variable.md)다.
- 다음 개념으로는 [statistical-inference.md](./statistical-inference.md), `실용 통계`, `인공지능 수학`이 이어진다.
- 학년 허브에서는 [high-2-hub.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/queries/math-curriculum-graph/high-2-hub.md), [high-3-hub.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/queries/math-curriculum-graph/high-3-hub.md)와 연결된다.
- 계통 허브에서는 [statistics-and-probability-strand.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/queries/math-curriculum-graph/statistics-and-probability-strand.md)와 [probability-and-statistics-course.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/queries/math-curriculum-graph/probability-and-statistics-course.md)를 본다.

## Open Questions

- `확률분포` 카드 안에 `확률변수`, `기대값`, `분산`을 함께 둘지, 고교 심화에서는 각각 분리할지 기준이 필요하다.
- 이산분포 중심으로 시작한 뒤 연속분포를 후속 카드로 둘지 구조 선택이 필요하다.

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
