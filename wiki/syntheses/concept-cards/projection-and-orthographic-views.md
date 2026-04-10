---
title: 투영과 투상도
type: synthesis
status: active
updated: 2026-04-10
card_role: concept
source_docs:
  - docs/math-concept-encyclopedia/china.md
  - docs/math-curriculum-research/china.md
  - docs/math-concept-encyclopedia/unit-practice-book.md
tags:
  - concept
  - geometry
  - spatial
---

# 투영과 투상도

## Summary

투영과 투상도는 3차원 물체를 2차원 평면에 옮겨 보는 개념이다. 공간 정보를 한 번에 다 보여 줄 수 없다는 사실을 이해하게 해 주며, 공간좌표와 도면 읽기, 시각화 감각의 출발점이 된다.

## Key Points

- 정의
  - 공간 도형을 특정 방향에서 평면으로 비추어 나타낸 그림을 투상도라 하고, 그렇게 옮기는 과정을 투영이라 한다.
- 핵심 개념
  - 정면도
  - 평면도
  - 측면도
  - 시점
  - 정보 손실
  - 3차원과 2차원의 대응
- 대표 성질
  - 하나의 투상도만으로는 물체의 모든 길이와 깊이를 완전히 복원하기 어렵다.
  - 서로 다른 시점의 그림을 함께 보아야 입체 구조를 더 정확히 알 수 있다.
- 대표 수식
  - 이 카드의 핵심은 공식보다 시각화 규칙이다.
- 증명 스케치
  - `증명 스케치 (추론)`:
  - 공간의 점을 평면으로 옮기면 깊이 정보 하나가 사라진다.
  - 그래서 3차원 물체를 한 장의 2차원 그림으로 바꿀 때는 정보가 줄어드는 것이 자연스럽다.
  - 이 때문에 정면도, 평면도, 측면도를 함께 보는 방식이 필요하다.

## Deep Dive

투영은 `그림 그리기 기술`이 아니라 `정보를 어떤 방향에서 읽을 것인가`를 정하는 수학적 선택이다. 앞에서 보면 높이와 너비는 잘 보이지만 깊이는 잘 안 보이고, 위에서 보면 너비와 깊이는 보이지만 높이가 약해진다.

이 개념은 기하 문제뿐 아니라 공학도면, 건축, 제품 설계, 컴퓨터 그래픽스에도 이어진다. 그래서 투상도는 공간감각을 수학 언어로 번역하는 작은 관문이라고 볼 수 있다.

## Worked Examples

- 예제 1: 정육면체를 정면에서 보면 정사각형 하나로 보인다. 하지만 이것만으로는 깊이를 알 수 없으므로 평면도나 측면도를 함께 봐야 한다.
- 예제 2: 계단 모양 블록의 정면도는 직사각형 두 개가 이어진 모양처럼 보여도, 평면도를 보면 뒤쪽 폭 차이가 드러날 수 있다.
- 예제 3: 같은 물체라도 보는 방향을 바꾸면 투상도의 모양이 달라진다. 그래서 시점을 먼저 고정해야 그림을 비교할 수 있다.

## Common Pitfalls

- 한 장의 그림만 보고 입체 전체를 단정하면 안 된다.
- 정면도와 평면도의 방향 기준을 섞으면 물체 모양을 잘못 읽기 쉽다.
- 보이는 길이와 실제 길이를 같은 것으로 오해하면 안 된다.
- 입체의 높이, 너비, 깊이를 어느 그림에서 읽는지 분리하지 않으면 혼란이 커진다.

## Curriculum Context

- 교육과정 배치
  - 중국 대표 배치에서는 `9학년 2학기`의 명시적 공간 시각화 단원이다.
  - 이후 공간좌표, 공학도면, 그래픽스 해석으로 이어진다.
- 국가별 배치 스냅샷
  - 중국: `투영과 투상도`가 명시 소단원으로 나타난다.
  - 한국: 현재 문서군에서는 `공간도형`, `공간좌표`, `직무 수학` 쪽 응용으로 간접 연결된다.
  - 일본: 중학교 공간도형과 고교의 입체 해석 감각으로 분산 연결된다.
  - 미국: middle-school measurement와 Geometry, technical drawing 맥락으로 이어진다.

## Connections

- 선수 개념은 [basic-geometry-and-construction.md](./basic-geometry-and-construction.md), [spatial-coordinates.md](./spatial-coordinates.md)다.
- 다음 개념으로는 [workplace-math.md](./workplace-math.md), `공학도면`, `컴퓨터 그래픽스`가 이어진다.
- 중국 학기 허브에서는 [china-grade-9-semester-2.md](./china-grade-9-semester-2.md), [geometry-strand.md](../geometry-strand.md)와 연결된다.

## Open Questions

- `정투영`과 `원근 표현`을 같은 카드 안에서 다룰지 시각화 하위 카드로 더 쪼갤지 기준이 필요하다.

## Sources

- `docs/math-concept-encyclopedia/china.md`
- `docs/math-curriculum-research/china.md`
- `docs/math-concept-encyclopedia/unit-practice-book.md`
