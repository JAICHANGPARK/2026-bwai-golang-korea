---
title: 도형의 이동
type: synthesis
status: active
updated: 2026-04-10
card_role: concept
source_docs:
  - docs/math-concept-encyclopedia/japan.md
  - docs/math-concept-encyclopedia/us.md
  - docs/math-concept-encyclopedia/unit-practice-book.md
tags:
  - concept
  - geometry
  - transformation
---

# 도형의 이동

## Summary

도형의 이동은 도형을 평행이동, 대칭이동, 회전, 확대·축소 같은 변환으로 바라보는 개념이다. 합동과 닮음을 `길이와 각의 조건`만이 아니라 `변환`으로 읽게 해 주는 현대 기하의 핵심 언어다.

## Key Points

- 정의
  - 도형의 모양이나 위치를 일정한 규칙에 따라 옮기는 변환을 도형의 이동이라 한다.
- 핵심 개념
  - 평행이동
  - 대칭이동
  - 회전
  - 확대와 축소
  - 변환 관점
- 대표 수식
  - 평행이동의 예: $(x,y)\mapsto (x+a,y+b)$
- 성질 1
  - 평행이동, 회전, 대칭이동은 길이와 각의 크기를 보존한다.
- 성질 2
  - 확대·축소는 길이를 일정한 비율로 바꾸고 넓이는 그 제곱 비율로 바꾼다.
- 성질 3
  - 합성 변환은 순서에 따라 결과가 달라질 수 있다.
- 증명 스케치
  - `증명 스케치 (추론)`:
  - 평행이동, 회전, 대칭이동은 길이와 각의 크기를 보존한다.
  - 그래서 이런 변환으로 서로 겹쳐지는 도형은 합동으로 읽을 수 있다.
  - 확대·축소는 길이 비를 보존하지 않지만 같은 비율로 늘이므로 닮음을 설명한다.
- 대표 예제
  - 점 $(1,2)$를 $(3,-1)$만큼 평행이동하면 $(4,1)$이 된다.
  - 점 $(2,1)$을 x축 대칭하면 $(2,-1)$이 된다.

## Deep Dive

- 도형의 이동은 합동과 닮음을 설명하는 변환 언어다.
- 평행이동·회전·대칭은 강체이동이고, 확대·축소는 닮음 변환이다.
- 변환을 연속적으로 합성하면 결과가 달라질 수 있으므로 순서가 중요하다.
- 좌표식으로 쓰면 변환의 효과를 계산적으로 확인할 수 있다.

## Worked Examples

### 예제 1: 평행이동

- 점 $(1,2)$를 $(3,-1)$만큼 평행이동하면 $(4,1)$이 된다.

### 예제 2: 대칭이동

- 점 $(2,1)$을 x축 대칭하면 $(2,-1)$이 된다.

### 예제 3: 확대

- 원점 중심, 확대비 $2$의 확대는 점 $(1,-3)$을 $(2,-6)$으로 보낸다.

## Common Pitfalls

- 평행이동의 벡터를 좌우/상하 순서 없이 섞어 쓰면 안 된다.
- 회전과 대칭을 같은 변환으로 취급하면 안 된다.
- 확대·축소는 길이를 보존하지 않으므로 합동 변환과 구분해야 한다.
- 합성 변환에서는 순서가 바뀌면 결과도 달라질 수 있다.

## Curriculum Context

- 교육과정 배치
  - 일본 문서에서는 `작도와 도형의 이동`이 중학교 초반 기하 입문의 한 축이다.
- 국가별 배치 스냅샷
  - 일본: `작도와 도형의 이동`이 명시적으로 등장한다.
  - 미국: Geometry에서 합동·닮음을 변환 관점으로 설명하는 경우가 많다.
  - 한국/중국: 현재 문서군에서는 독립 표제보다 합동·닮음·좌표기하 안에 간접적으로 녹아 있는 편이다.
- 표현과 문제 감각
  - 도형의 이동은 `모양이 같음`을 눈대중이 아니라 변환 규칙으로 설명하게 해 준다.

## Connections

- 선수 개념은 [basic-geometry-and-construction.md](./basic-geometry-and-construction.md)다.
- 다음 개념으로는 [congruence.md](./congruence.md), [similarity.md](./similarity.md), [polar-ideas.md](./polar-ideas.md), [korean-geometry-course.md](./korean-geometry-course.md)가 이어진다.
- 계통 허브에서는 [geometry-strand.md](../geometry-strand.md), [korea-curriculum-hub.md](../korea-curriculum-hub.md)를 함께 본다.

## Open Questions

- `확대와 축소`를 닮음 카드와 더 강하게 결합할지, 변환 카드의 하위 섹션으로 둘지 후속 기준이 필요하다.

## Sources

- `docs/math-concept-encyclopedia/japan.md`
- `docs/math-concept-encyclopedia/us.md`
- `docs/math-concept-encyclopedia/unit-practice-book.md`
