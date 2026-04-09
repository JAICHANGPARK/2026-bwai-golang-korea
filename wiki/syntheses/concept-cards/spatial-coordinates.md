---
title: 공간좌표
type: synthesis
status: active
updated: 2026-04-09
source_docs:
  - docs/math-concept-encyclopedia/korea.md
  - docs/math-concept-encyclopedia/china.md
  - docs/math-concept-encyclopedia/japan.md
  - docs/math-concept-encyclopedia/multilingual-glossary.md
tags:
  - concept
  - geometry
  - spatial
---

# 공간좌표

## Summary

공간좌표는 3차원 공간의 점과 도형을 좌표로 표현하는 개념이다. 평면좌표의 생각을 입체로 확장해 직선, 평면, 거리, 위치 관계를 계산하게 해 준다.

## Key Points

- 정의
  - 공간 속 한 점의 위치를 세 수 $(x,y,z)$로 나타내는 체계를 공간좌표라 한다.
- 핵심 개념
  - 3차원 좌표축
  - 점의 위치
  - 거리
  - 직선과 평면
  - 투영
- 대표 수식
  - $P=(x,y,z)$
  - 두 점 사이 거리 $d=\sqrt{(x_2-x_1)^2+(y_2-y_1)^2+(z_2-z_1)^2}$
- 증명 스케치
  - `증명 스케치 (추론)`:
  - 공간거리 공식은 평면의 피타고라스 정리를 두 번 적용해 얻을 수 있다.
  - 먼저 바닥면에서의 거리, 그다음 높이 차이를 결합하면 3차원 거리 공식을 얻는다.
- 대표 예제
  - $(0,0,0)$과 $(1,2,2)$ 사이 거리는 $\sqrt{1+4+4}=3$이다.
- 교육과정 배치
  - 한국 대표 배치에서는 `기하`에서 공간도형과 함께 다룬다.
- 국가별 배치 스냅샷
  - 한국: `기하`의 `공간도형과 공간좌표` 축에 놓인다.
  - 일본: `수학C`의 벡터·곡선 관련 내용과 연결된다.
  - 중국: `공간벡터와 해석기하`에서 핵심적으로 다뤄진다.
  - 미국: 일반 K-12에서는 제한적이지만 advanced geometry와 calculus 전단계에서 나타난다.
- 표현과 문제 감각
  - 다국어 용어: `three-dimensional coordinates`, `空間座標`, `空间直角坐标系`
  - 2차원 그림으로 3차원 구조를 상상하는 시각화 감각이 중요하다.

## Connections

- 선수 개념은 [vectors.md](./vectors.md), [equations-of-geometric-figures.md](./equations-of-geometric-figures.md), [pythagorean-theorem.md](./pythagorean-theorem.md)다.
- 다음 개념으로는 `공간벡터`, `물리의 운동`, `컴퓨터 그래픽스`가 이어진다.

## Open Questions

- `직선과 평면의 위치관계`를 별도 카드로 분리할지 검토가 필요하다.

## Sources

- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-concept-encyclopedia/china.md`
- `docs/math-concept-encyclopedia/japan.md`
- `docs/math-concept-encyclopedia/multilingual-glossary.md`
