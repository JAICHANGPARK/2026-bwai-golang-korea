---
title: 조건부확률
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
  - $P(B)>0$일 때, 사건 $B$가 일어났다고 알고 있을 때 사건 $A$가 일어날 확률을
    $$
    P(A\mid B)=\frac{P(A\cap B)}{P(B)}
    $$
    로 정의한다.
- 핵심 개념
  - 조건
  - 표본공간의 축소
  - 사건의 교집합
  - 독립성
  - 확률 갱신
- Deep Dive
  - 조건부확률은 `정보가 추가되었을 때 확률을 다시 정규화하는 규칙`이다.
  - 사건 $B$가 주어지면 전체 표본공간 대신 $B$ 내부에서만 확률을 다시 배분해야 한다.
  - 따라서 조건부확률은 단순한 분수 계산이 아니라, 새로운 정보에 맞게 확률모형을 업데이트하는 과정이다.
- Theorem
  - 곱셈정리:
    $$
    P(A\cap B)=P(B)P(A\mid B)=P(A)P(B\mid A)
    $$
  - 독립성의 동치조건:
    $$
    A,B\text{가 독립} \iff P(A\cap B)=P(A)P(B)
    $$
    이고, $P(B)>0$이면
    $$
    P(A\mid B)=P(A)
    $$
    와도 동치이다.
  - 배반과의 구분:
    - $A\cap B=\varnothing$이면 두 사건은 배반이다.
    - 두 사건이 둘 다 양의 확률을 가지는 경우 배반이면 독립일 수 없다.
- 대표 수식
  - $P(A\mid B)=\frac{P(A\cap B)}{P(B)}\;(P(B)>0)$
  - $P(A\cap B)=P(B)P(A\mid B)$
  - $P(A\mid B)=P(A)$ when $A,B$ are independent and $P(B)>0$
- 증명 스케치
  - `증명 스케치 (추론)`:
  - 아래 설명은 `유한한 등가능 표본공간`에서의 직관이다.
  - $B$가 이미 일어났다고 알면, 원래 전체 표본공간이 아니라 `B 안`만 새로운 표본공간으로 본다.
  - 이 안에서 $A$도 함께 일어나는 경우는 $A\cap B$에 해당하므로, 직관적으로는 새로운 확률이
    $$
    \frac{\text{B 안에서 A도 일어나는 경우}}{\text{B의 전체 경우}}
    $$
    로 계산된다.
  - 확률 이론에서는 이 직관을 일반화해
    $$
    P(A\mid B)=\frac{P(A\cap B)}{P(B)}
    $$
    를 정의로 사용한다.
- 대표 예제
  - 공정한 주사위를 던질 때 $A=\{\text{짝수}\}$, $B=\{\text{4보다 큼}\}$이라 하자.
  - 그러면 $B=\{5,6\}$이고 그 안에서 $A$도 만족하는 경우는 `{6}`뿐이므로
    $$
    P(A\mid B)=\frac12
    $$
    이다.
  - 공정한 동전을 두 번 던질 때 $A=\{\text{첫 번째가 앞면}\}$, $B=\{\text{두 번째가 앞면}\}$이라 하자.
  - 두 시행은 서로 영향을 주지 않으므로
    $$
    P(A\mid B)=P(A)=\frac12
    $$
    이고,
    $$
    P(A\cap B)=\frac14=P(A)P(B)
    $$
    이다.
- Worked Examples
  - 예제 1: 카드 한 장을 무작위로 뽑을 때, 사건 $H$를 `하트`, 사건 $R$을 `빨간 카드`라 하자.
    - $P(H)=\frac{13}{52}=\frac14$
    - $P(H\mid R)=\frac{13/52}{26/52}=\frac12$
    - 빨간 카드라는 조건이 들어오면 표본공간은 26장으로 줄어든다.
  - 예제 2: 빨간 공 3개, 파란 공 2개가 든 상자에서 비복원으로 두 번 뽑는다.
    - $A=\{\text{첫 번째가 빨간 공}\}$, $B=\{\text{두 번째가 빨간 공}\}$
    - $P(A)=\frac35$
    - $P(B\mid A)=\frac24=\frac12$
    - 따라서
      $$
      P(A\cap B)=P(A)P(B\mid A)=\frac35\cdot\frac12=\frac3{10}
      $$
    - 비복원 추출에서는 첫 사건이 둘째 사건에 영향을 주므로 독립이 아니다.
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
  - 조건부확률 문제는 `조건을 정확히 읽고`, `새 표본공간을 제대로 잡는지`가 승부를 가른다.

## Common Pitfalls

- $P(A\mid B)$와 $P(B\mid A)$를 혼동한다.
- $P(B)=0$인데도 조건부확률 공식을 기계적으로 적용한다.
- 독립과 배반을 같은 것으로 착각한다.
- 조건이 붙었는데도 분모를 원래 전체 표본공간 크기로 두는 실수를 한다.

## Connections

- 선수 개념은 [counting-principles.md](./counting-principles.md)와 `확률`이다.
- 다음 개념으로는 [random-variable.md](./random-variable.md), [probability-distribution.md](./probability-distribution.md), [statistical-inference.md](./statistical-inference.md)가 이어진다.
- 과목 허브에서는 [probability-and-statistics-course.md](./probability-and-statistics-course.md)를 본다.
- 계통 허브로는 [../statistics-and-probability-strand.md](../statistics-and-probability-strand.md)를 함께 본다.

## Open Questions

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
