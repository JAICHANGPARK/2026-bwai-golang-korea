---
title: 직선의 방정식
type: synthesis
status: active
updated: 2026-04-10
card_role: concept
source_docs:
  - docs/math-concept-encyclopedia/korea.md
  - docs/math-concept-encyclopedia/japan.md
  - docs/math-concept-encyclopedia/us.md
  - docs/math-concept-encyclopedia/unit-practice-book.md
tags:
  - concept
  - geometry
  - coordinate
---

# 직선의 방정식

## Summary

직선의 방정식은 좌표평면의 직선을 식으로 나타내는 개념이다. 기울기와 한 점, 또는 두 점만 알면 직선을 수식으로 바꿀 수 있다는 점에서 좌표기하의 가장 기본적인 언어다.

## Key Points

- 정의
  - 좌표평면 위의 직선을 식으로 나타낸 것을 직선의 방정식이라 한다.
- 핵심 개념
  - 기울기
  - 절편
  - 점기울기형
  - 일반형
  - 수직선
- 대표 수식
  - $y=mx+b$
  - $y-y_1=m(x-x_1)$
  - $ax+by+c=0$
- 성질 1
  - 기울기가 같으면 두 직선은 평행하다.
- 성질 2
  - 수직선은 $x=k$ 꼴로 나타난다.
- 증명 스케치
  - `증명 스케치 (추론)`:
  - 직선 위의 모든 점에서 변화율이 일정하다면
    $$
    \frac{y-y_1}{x-x_1}=m
    $$
    로 볼 수 있다.
  - 이를 정리하면
    $$
    y-y_1=m(x-x_1)
    $$
    을 얻고, 이것이 직선의 점기울기형이다.

## Deep Dive

직선의 방정식은 모양이 하나가 아니다. `기울기-절편형`, `점기울기형`, `일반형`은 모두 같은 직선을 다른 관점으로 표현한 것일 뿐이다. 문제에서 주어진 정보가 무엇이냐에 따라 가장 편한 형태를 고르는 것이 중요하다.

또한 수직선은 기울기를 정의할 수 없으므로 $y=mx+b$ 꼴에 억지로 넣으면 안 된다. 이 점은 좌표기하에서 자주 나오는 예외다.

## Worked Examples

- 예제 1: 기울기가 `2`이고 점 `(1,-3)`을 지나는 직선은
  $$
  y+3=2(x-1)
  $$
  이고, 정리하면
  $$
  y=2x-5
  $$
  이다.
- 예제 2: 점 `(1,2)`와 `(3,6)`을 지나는 직선의 기울기는
  $$
  \frac{6-2}{3-1}=2
  $$
  이므로 방정식은
  $$
  y-2=2(x-1)
  $$
  이다.
- 예제 3: x좌표가 항상 `4`인 수직선의 방정식은
  $$
  x=4
  $$
  이다.

## Common Pitfalls

- 수직선을 기울기형에 넣으려 하면 안 된다.
- 기울기 계산에서 분자와 분모 순서를 뒤섞으면 부호가 틀어진다.
- 두 점을 지나는 직선에서 한 점만 대입하고 다른 점으로 검산하지 않으면 실수가 남기 쉽다.
- $y=mx+b$에서 $b$는 y절편이지 x절편이 아니다.

## Curriculum Context

- 교육과정 배치
  - 한국 대표 배치에서는 `공통수학2`의 도형의 방정식 핵심 세부 개념이다.
  - 일본 `수학II`, 미국 Geometry/Algebra, 중국의 좌표-함수 브리지에서도 자연스럽게 읽힌다.
- 국가별 배치 스냅샷
  - 한국: 거리 공식, 중점, 원의 방정식과 함께 좌표기하의 시작점이다.
  - 일본: `도형과 방정식` 안에서 기울기와 직선 표현이 분명하게 다뤄진다.
  - 중국: 중학교 좌표평면과 일차함수 위에 고등 해석기하로 확장된다.
  - 미국: `Linear Equations and Systems`, `coordinate geometry`에서 반복 등장한다.

## Connections

- 선수 개념은 [linear-function.md](./linear-function.md), [equations-of-geometric-figures.md](./equations-of-geometric-figures.md)다.
- 다음 개념으로는 [equation-of-a-circle.md](./equation-of-a-circle.md), [parabola.md](./parabola.md), [vectors.md](./vectors.md)로 이어진다.
- 계통 허브에서는 [geometry-strand.md](../geometry-strand.md), [high-1-hub.md](../high-1-hub.md)와 연결된다.

## Open Questions

- `두 직선의 위치 관계`를 별도 하위 카드로 더 뺄지 검토가 필요하다.

## Sources

- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-concept-encyclopedia/japan.md`
- `docs/math-concept-encyclopedia/us.md`
- `docs/math-concept-encyclopedia/unit-practice-book.md`
