---
title: 함수 일반론
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
  - function
  - foundations
---

# 함수 일반론

## Summary

함수 일반론은 입력과 출력의 대응 규칙을 하나의 언어로 묶는 개념이다. 일차함수나 이차함수 같은 개별 함수형을 넘어, `함수란 무엇인가`를 정의역·치역·그래프의 언어로 정확히 읽는 카드다.

## Key Points

- 정의
  - 집합 $A$의 각 원소 $x$에 집합 $B$의 원소 하나가 정확히 대응되면 $f: A \to B$를 함수라 한다.
  - 입력 집합을 정의역, 출력이 놓이는 집합을 공역, 실제로 나오는 값들의 집합을 치역이라 한다.
- 핵심 개념
  - 입력과 출력
  - 정의역
  - 치역
  - 공역
  - 그래프
  - 대응 규칙
  - 함수값
  - 합성 함수
  - 역함수의 조건
- 대표 성질
  - 같은 입력이 두 개의 서로 다른 출력에 대응하면 함수가 아니다.
  - 그래프는 임의의 수직선과 많아야 한 점에서만 만나야 함수 그래프다.
  - 서로 다른 두 함수가 같은 정의역에서 모든 입력에 대해 같은 값을 가지면 그 두 함수는 같다.
  - 역함수는 일대일 대응일 때만 정의할 수 있다.
- 대표 수식
  - $f: A \to B$
  - $f(x)$
  - $y=f(x)$
  - $(g \circ f)(x)=g(f(x))$
- 증명 스케치
  - `증명 스케치 (추론)`:
  - 함수의 정의 자체가 `입력 하나에 출력 하나`를 요구하므로, 한 입력에 두 출력이 붙으면 정의를 위반한다.
  - 그래프에서 어떤 수직선이 두 점 이상과 만나면 같은 $x$에 서로 다른 $y$가 대응한 것이므로 함수가 아니다.
  - 역함수는 모든 함수에 존재하지 않는다. 서로 다른 입력이 서로 다른 출력으로 가는 `일대일성`이 있어야 뒤집을 수 있다.
## Deep Dive

- 함수는 `대상`보다 `대응 규칙`을 먼저 보는 언어다.
- 정의역은 실제 입력의 집합이고, 공역은 출력이 놓일 수 있는 집합이며, 치역은 그중 실제로 나온 값들이다.
- 합성함수는 순서가 중요하고, `g \circ f`와 `f \circ g`는 일반적으로 다르다.
- 역함수는 함수의 대응을 되돌리는 역할을 하므로, 일대일 대응이 아니면 존재할 수 없다.

## Worked Examples

### 예제 1: 함수값 계산

- $f(x)=2x+1$이면 $f(3)=7$이고 $f(-2)=-3$이다.

### 예제 2: 함수 여부 판단

- 관계 $\{(1,2),(2,3),(1,4)\}$는 $1$이 두 출력 $2,4$에 대응하므로 함수가 아니다.

### 예제 3: 역함수 존재 조건

- $f(x)=x^2$는 정의역이 $\mathbb{R}$일 때 함수이지만 역함수는 전체 실수에서 바로 존재하지 않는다.
- 정의역을 $x\ge 0$으로 제한해야 역함수가 생긴다.

## Common Pitfalls

- 관계와 함수를 같은 말로 쓰면 안 된다.
- 정의역을 적지 않고 식만 보고 함수를 판단하면 오류가 생긴다.
- 공역과 치역을 같은 말처럼 쓰면 함수의 범위가 흐려진다.
- 모든 함수가 역함수를 가진다고 생각하면 안 된다.
- 그래프가 그려진다고 해서 자동으로 함수는 아니다. 수직선 판정을 먼저 봐야 한다.
- 교육과정 배치
  - 한국 대표 배치에서는 `공통수학2`의 `함수와 그래프` 축에서 일반 함수 해석이 강화된다.
- 국가별 배치 스냅샷
  - 한국: 공통수학의 `함수와 그래프`에서 일반성 있게 읽는다.
  - 일본: 중학교 비례·일차함수·이차함수 위계 뒤에 고교 수학에서 함수 해석이 일반화된다.
  - 중국: 고등학교 필수과정의 `함수`가 명시적 중심 축이다.
  - 미국: Grade 8의 `Functions and Linear Models`와 고교 Algebra/Precalculus 전반의 공통 언어다.
- 표현과 문제 감각
  - 다국어 용어: `function`, `関数`, `函数`
  - 식, 표, 그래프를 같은 관계의 다른 표현으로 읽는 감각이 중요하다.

## Connections

- 선수 개념은 [proportion.md](./proportion.md), [linear-function.md](./linear-function.md), [inverse-proportion.md](./inverse-proportion.md)다.
- 같은 축의 인접 개념으로는 [exponential-and-logarithmic-functions.md](./exponential-and-logarithmic-functions.md), [trigonometric-function.md](./trigonometric-function.md), [limits.md](./limits.md), [sets-and-propositions.md](./sets-and-propositions.md)가 있다.
- 함수 계통은 [functions-strand.md](../functions-strand.md)에서 이어서 본다.
- 고교 공통 축은 [high-1-hub.md](../high-1-hub.md), [korea-curriculum-hub.md](../korea-curriculum-hub.md), [course-track-hub.md](../course-track-hub.md)와 연결해 읽는다.

## Open Questions

- `정의역`과 `치역`을 별도 용어 카드로 더 세분화할지 검토가 필요하다.

## Sources

- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-concept-encyclopedia/china.md`
- `docs/math-concept-encyclopedia/us.md`
- `docs/math-concept-encyclopedia/multilingual-glossary.md`
