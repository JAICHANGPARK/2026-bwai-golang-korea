---
title: 피타고라스 정리
type: synthesis
status: active
updated: 2026-04-10
card_role: concept
source_docs:
  - docs/math-concept-encyclopedia/korea.md
  - docs/math-concept-encyclopedia/japan.md
  - docs/math-concept-encyclopedia/china.md
  - docs/math-concept-encyclopedia/us.md
  - docs/math-concept-encyclopedia/multilingual-glossary.md
tags:
  - concept
  - geometry
  - theorem
---

# 피타고라스 정리

## Summary

피타고라스 정리는 직각삼각형의 세 변 사이의 제곱 관계를 설명하는 정리다. 길이 계산을 넘어 좌표거리, 원의 방정식, 벡터 길이까지 이어지는 기하와 대수의 대표 연결점이다.

## Key Points

- 정의
  - 직각삼각형에서 빗변의 제곱은 나머지 두 변의 제곱의 합과 같다는 정리다.
- 핵심 개념
  - 직각삼각형
  - 빗변
  - 직각변
  - 제곱 관계
  - 거리
- 대표 수식
  - $a^2+b^2=c^2$
- 성질 1
  - 직각삼각형에서만 $a^2+b^2=c^2$가 성립한다.
- 성질 2
  - 역도 성립한다. 세 변의 길이가 $a,b,c$일 때 $a^2+b^2=c^2$이면 그 삼각형은 직각삼각형이다.
- 증명 스케치
  - `증명 스케치 (추론)`:
  - 같은 직각삼각형 네 개를 큰 정사각형 안에 배열하면 남는 넓이를 두 방식으로 계산할 수 있다.
  - 한쪽은 큰 정사각형 넓이, 다른 쪽은 네 삼각형 넓이와 가운데 작은 정사각형 넓이의 합이다.
  - 이를 비교하면 $a^2+b^2=c^2$를 얻는다.
  - 역정리는 피타고라스 정리의 결과를 다시 써서 삼각형이 직각임을 판정하는 방식으로 이해할 수 있다.
- 대표 예제
  - 직각변이 `3,4`이면 빗변은 $5$다.
  - 세 변이 `5,12,13`이면
    $$
    5^2+12^2=13^2
    $$
    이므로 직각삼각형이다.
- Deep Dive
  - 피타고라스 정리는 길이 계산 공식을 넘어 거리의 제곱 구조를 보여 준다.
  - 좌표평면에서 두 점 사이의 거리를 구할 때도 이 정리가 그대로 들어간다.
  - 원의 방정식과 벡터 길이 계산의 바닥에는 이 정리가 있다.
- Worked Examples

### 예제 1: 빗변 구하기

- 직각변이 `6`과 `8`이면 빗변은
  $$
  c=\sqrt{6^2+8^2}=10
  $$
  이다.

### 예제 2: 직각 판정하기

- 세 변이 `7,24,25`이면
  $$
  7^2+24^2=25^2
  $$
  이므로 직각삼각형이다.

## Common Pitfalls

- 직각삼각형이 아닌데도 식을 무조건 적용하면 안 된다.
- 빗변은 항상 가장 긴 변이지만, 가장 긴 변이라고 해서 언제나 빗변인 것은 아니다. 직각삼각형이어야 한다.
- 역정리를 쓸 때는 세 변의 길이를 정확히 비교해야 한다.
- 제곱을 풀 때 `\sqrt{a^2+b^2}`와 `a+b`를 혼동하면 안 된다.
## Curriculum Context
- 교육과정 배치
  - 한국 대표 배치에서는 `중3` 기하 핵심 정리다.
  - 이후 삼각비, 도형의 방정식, 좌표거리로 이어진다.
- 국가별 배치 스냅샷
  - 한국: `중3`의 닮음·원·피타고라스·삼각비 축에서 다룬다.
  - 일본: `중3`의 `닮음·원주각·피타고라스 정리`에 포함된다.
  - 중국: `8학년`의 독립 단원으로 강하게 다룬다.
  - 미국: Grade 8과 Geometry에서 반복 등장한다.
- 표현과 문제 감각
  - 다국어 용어: `Pythagorean theorem`, `三平方の定理`, `勾股定理`
  - 길이 계산 공식으로만 보지 않고 `거리의 제곱 구조`로 읽는 것이 중요하다.

## Connections

- 선수 개념은 [similarity.md](./similarity.md)와 `직각삼각형`이다.
- 다음 개념으로는 [trigonometric-ratio.md](./trigonometric-ratio.md), [equations-of-geometric-figures.md](./equations-of-geometric-figures.md), [vectors.md](./vectors.md), [circle.md](./circle.md)가 이어진다.
- 계통 허브에서는 [geometry-strand.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/syntheses/geometry-strand.md)와 [korea-curriculum-hub.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/syntheses/korea-curriculum-hub.md)를 함께 본다.

## Open Questions

- 넓이 증명 외에 닮음 기반 증명을 별도 카드로 둘지 검토가 필요하다.

## Sources

- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-concept-encyclopedia/japan.md`
- `docs/math-concept-encyclopedia/china.md`
- `docs/math-concept-encyclopedia/us.md`
- `docs/math-concept-encyclopedia/multilingual-glossary.md`
