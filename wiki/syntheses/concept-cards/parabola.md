---
title: 포물선
type: synthesis
status: active
updated: 2026-04-10
card_role: concept
source_docs:
  - docs/math-concept-encyclopedia/korea.md
  - docs/math-concept-encyclopedia/japan.md
  - docs/math-concept-encyclopedia/china.md
  - docs/math-concept-encyclopedia/unit-practice-book.md
tags:
  - concept
  - geometry
  - conic
---

# 포물선

## Summary

포물선은 한 초점과 한 준선까지의 거리가 같은 점들의 자취다. 중학교의 [quadratic-function.md](./quadratic-function.md)와 고등학교 [conic-sections.md](./conic-sections.md)를 가장 직접적으로 이어 주는 대표 곡선이다.

## Key Points

- 정의
  - 한 점(초점)과 한 직선(준선)까지의 거리가 같은 점들의 집합을 포물선이라 한다.
- 핵심 개념
  - 초점
  - 준선
  - 축
  - 꼭짓점
  - 대칭성
- 대표 수식
  - $y=a(x-p)^2+q$
  - $y^2=4px$
- 성질 1
  - 포물선은 축에 대하여 대칭이다.
- 성질 2
  - $a$의 부호는 열린 방향, $|a|$는 벌어짐 정도를 바꾼다.
- 증명 스케치
  - `증명 스케치 (추론)`:
  - 초점이 $(p,0)$, 준선이 $x=-p$인 포물선 위의 점 $(x,y)$는
    $$
    \sqrt{(x-p)^2+y^2}=x+p
    $$
    를 만족한다.
  - 양변을 제곱하고 정리하면
    $$
    y^2=4px
    $$
    를 얻는다.

## Deep Dive

포물선은 같은 곡선을 `함수 그래프`로도, `거리 조건의 자취`로도 읽을 수 있다는 점이 중요하다. 중학교에서는 $y=ax^2+bx+c$의 그래프로 먼저 만나고, 고등학교에서는 초점과 준선을 가진 해석기하의 곡선으로 다시 만난다.

그래프 문제에서는 꼭짓점과 축이 중요하고, 해석기하 문제에서는 초점과 준선이 중요하다. 두 언어가 같은 곡선을 서로 다른 시점에서 설명하고 있다고 이해하면 훨씬 단단해진다.

## Worked Examples

- 예제 1:
  $$
  y=-2(x-1)^2+3
  $$
  의 꼭짓점은 `(1,3)`이고 아래로 열린다.
- 예제 2:
  $$
  y^2=8x
  $$
  는
  $$
  y^2=4px
  $$
  와 비교하면 $p=2$이므로 초점은 `(2,0)`이다.

## Common Pitfalls

- [quadratic-function.md](./quadratic-function.md)의 포물선과 해석기하의 포물선을 전혀 다른 대상으로 보면 안 된다.
- $a$의 부호와 크기가 각각 무엇을 바꾸는지 헷갈리기 쉽다.
- 초점과 꼭짓점을 같은 점으로 오해하면 안 된다.
- 축을 x축이나 y축 하나로만 고정해서 생각하면 다양한 형태를 놓치게 된다.

## Curriculum Context

- 교육과정 배치
  - 한국에서는 `중3` 이차함수와 `기하`의 이차곡선을 잇는 세부 곡선이다.
  - 일본과 중국에서도 이차함수 해석과 곡선 해석 사이의 중요한 브리지다.
- 국가별 배치 스냅샷
  - 한국: `이차함수`의 포물선과 `기하`의 포물선이 같은 곡선으로 이어진다.
  - 일본: `이차함수` 감각 위에 고교 곡선 해석으로 확장된다.
  - 중국: `이차함수`와 `공간벡터와 해석기하` 사이의 연결점이다.
  - 미국: Algebra I/II의 quadratic graph와 Precalculus의 conic view가 연결된다.

## Connections

- 선수 개념은 [quadratic-function.md](./quadratic-function.md), [equations-of-geometric-figures.md](./equations-of-geometric-figures.md)다.
- 다음 개념으로는 [conic-sections.md](./conic-sections.md), [ellipse.md](./ellipse.md), [hyperbola.md](./hyperbola.md)가 이어진다.
- 계통 허브에서는 [geometry-strand.md](../geometry-strand.md), [high-2-hub.md](../high-2-hub.md)와 연결된다.

## Open Questions

- 초점과 준선 문제를 별도 예제 카드로 더 분화할지 기준이 필요하다.

## Sources

- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-concept-encyclopedia/japan.md`
- `docs/math-concept-encyclopedia/china.md`
- `docs/math-concept-encyclopedia/unit-practice-book.md`
