---
title: 반비례
type: synthesis
status: active
updated: 2026-04-10
card_role: concept
source_docs:
  - docs/math-concept-encyclopedia/korea.md
  - docs/math-concept-encyclopedia/japan.md
  - docs/math-concept-encyclopedia/china.md
  - docs/math-concept-encyclopedia/multilingual-glossary.md
tags:
  - concept
  - function
  - proportion
---

# 반비례

## Summary

반비례는 한 변수가 커질수록 다른 변수가 일정한 곱을 유지하며 작아지는 관계를 다루는 개념이다. 정비례와 함께 중학교 함수 감각을 처음 세우는 비선형 관계 카드다.

## Key Points

- 정의
  - $a\neq 0$일 때 두 변수 $x,y$가 $xy=a$를 만족하면 $y=\frac{a}{x}$ 꼴의 반비례 관계에 있다고 한다.
  - 학교 수학에서는 `반비례`와 `반비례함수`를 거의 같은 출발점으로 읽는다.
- 핵심 개념
  - 비례상수
  - 정의역과 치역
  - 함수값
  - 쌍곡선
  - 쌍곡선 모양
  - 비선형 변화
  - 축에 대한 점근선
- 대표 성질
  - 곱 보존: $xy=a$이면 $x$가 커질수록 $y$는 그에 맞게 작아진다.
  - 함수형: $a\neq 0$일 때 $y=\frac{a}{x}$는 $x\neq 0$에서 정의되는 함수다.
  - 대칭성: $a>0$이면 그래프는 1사분면과 3사분면에, $a<0$이면 2사분면과 4사분면에 놓인다.
  - 점근선: 그래프는 축 $x=0$과 $y=0$에 가까워지지만 만나지는 않는다.
- 대표 수식
  - $y=\frac{a}{x}$
  - $xy=a$
  - $x\neq 0$
- 증명 스케치
  - `증명 스케치 (추론)`:
  - $xy=a$에서 $x\neq 0$이면 양변을 $x$로 나누어 $y=\frac{a}{x}$를 얻는다.
  - 따라서 한 변수가 커질수록 다른 변수는 같은 곱을 유지하도록 조정된다.
  - 그래프가 축을 만나지 않는 이유는 $x=0$이면 식 자체가 정의되지 않고, $y=0$이면 $xy=a\neq 0$와 모순되기 때문이다.
## Deep Dive

- 반비례는 `한쪽이 늘면 다른 쪽이 자동으로 줄어드는` 현상을 곱의 보존으로 설명한다.
- 그래프는 직선이 아니라 쌍곡선 형태이며, 원점에 가까워질수록 값의 변동이 커진다.
- $x$와 $y$가 모두 양수인 경우만 보는지, 음수까지 허용하는지에 따라 그래프의 사분면 해석이 달라진다.
- 실생활에서는 속도-시간, 작업자 수-시간, 넓이 일정한 직사각형의 변 길이 관계처럼 나타난다.

## Worked Examples

### 예제 1: 기본 계산

- $xy=12$일 때 $x=3$이면 $y=4$다.

### 예제 2: 음수 밑의 반비례

- $y=\frac{-8}{x}$에서 $x=-2$이면 $y=4$이고, $x=4$이면 $y=-2$다.

### 예제 3: 작업 모형

- 어떤 일을 6명이 10시간에 끝낸다면, 같은 속도 모형에서는 인원수와 시간의 곱이 일정하다고 볼 수 있다.

## Common Pitfalls

- $x=0$을 대입해 값을 구하려 하면 안 된다.
- 반비례를 단순히 `줄어드는 관계`로만 보면 안 된다. 핵심은 곱이 일정하다는 점이다.
- 그래프가 직선이 아니라고 해서 임의의 곡선과 같다고 보면 안 된다.
- 정비례의 `비율 일정`과 반비례의 `곱 일정`을 뒤섞으면 안 된다.
- 점근선을 실제로 만나는 선으로 착각하면 안 된다.
- 대표 예제
  - $xy=12$일 때 $x=3$이면 $y=4$다.
- 교육과정 배치
  - 한국 대표 배치에서는 `중1`의 좌표와 비례 단원에서 정비례와 함께 배운다.
- 국가별 배치 스냅샷
  - 한국: `중1`에서 정비례와 함께 다룬다.
  - 일본: `중1`의 `비례·반비례와 좌표`에 포함된다.
  - 중국: `9학년 2학기`에 `반비례함수`로 독립 단원이 있다.
  - 미국: 대체로 middle school의 proportional reasoning 안에서 간접적으로 다뤄지고, 함수 단원에서 명시화된다.
- 표현과 문제 감각
  - 다국어 용어: `inverse proportion`, `反比例`
  - 정비례처럼 직선이 아니라는 점과 $x=0$이 정의되지 않는다는 점이 핵심이다.

## Connections

- 선수 개념은 [proportion.md](./proportion.md), `좌표평면`이다.
- 같은 축의 인접 개념으로는 [function-foundations.md](./function-foundations.md), [linear-function.md](./linear-function.md), [exponential-and-logarithmic-functions.md](./exponential-and-logarithmic-functions.md)가 있다.
- 다음 개념으로는 `유리함수`, `모델링`이 이어진다.
- 함수 계통은 [functions-strand.md](../functions-strand.md)에서 다시 묶인다.
- 중1 허브와 한국 허브는 [middle-1-hub.md](../middle-1-hub.md), [korea-curriculum-hub.md](../korea-curriculum-hub.md)에서 함께 본다.

## Open Questions

- `반비례함수`를 중등 카드와 고교 유리함수 카드로 더 나눌지 검토가 필요하다.

## Sources

- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-concept-encyclopedia/japan.md`
- `docs/math-concept-encyclopedia/china.md`
- `docs/math-concept-encyclopedia/multilingual-glossary.md`
