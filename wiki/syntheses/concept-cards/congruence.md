---
title: 합동
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
  - congruence
---

# 합동

## Summary

합동은 두 도형이 위치만 바꾸면 정확히 겹친다는 개념이다. 모양뿐 아니라 크기까지 완전히 같다는 뜻이라서, 기하 증명과 변환 기하의 기본 언어가 된다.

## Key Points

- 정의
  - 두 도형을 평행이동, 회전, 대칭이동으로 옮겼을 때 정확히 겹치면 합동이라 한다.
- 핵심 개념
  - 대응점
  - 대응변
  - 대응각
  - 합동조건
  - 길이 보존
  - 각 보존
- 대표 수식
  - 삼각형 합동조건으로 `SSS`, `SAS`, `ASA`가 자주 쓰인다.
- 성질 1
  - 합동인 도형은 대응하는 모든 변의 길이와 각의 크기가 같다.
- 성질 2
  - 평행이동, 회전, 대칭이동 같은 강체이동은 길이와 각을 보존한다.
- 증명 스케치
  - `증명 스케치 (추론)`:
  - 삼각형의 세 변의 길이가 각각 같으면, 한 변을 기준으로 다른 꼭짓점의 위치가 하나로 결정된다.
  - 따라서 두 삼각형은 같은 크기와 같은 모양을 가져 서로 겹친다.
  - 이 논리는 `강체이동으로 정확히 포개진다`는 합동의 정의와 맞물린다.
- 대표 예제
  - 두 삼각형에서 세 변의 길이가 각각 같으면 두 삼각형은 합동이다.
  - 두 삼각형에서 두 변과 끼인각이 각각 같으면 역시 합동이다.
  - 합동인 두 도형의 둘레와 넓이는 각각 같다.
- Deep Dive
  - 합동은 `모양이 비슷하다`보다 강한 조건이다.
  - 합동조건은 삼각형의 경우 증명의 출발점이 되며, 다른 도형의 성질을 옮겨 쓰게 해 준다.
  - 중등 기하에서 `무엇이 같아야 도형이 완전히 같은가`를 판단하는 기준이 합동이다.
- Worked Examples

### 예제 1: SSS

- 두 삼각형의 세 변 길이가 각각 `3,4,5`와 `3,4,5`라면 두 삼각형은 합동이다.

### 예제 2: SAS

- 두 삼각형에서 두 변과 그 끼인각이 각각 같으면 두 삼각형은 합동이다.
- 따라서 대응하는 나머지 한 변과 두 각도 같다.

## Common Pitfalls

- 합동과 닮음을 같은 말로 쓰면 안 된다. 합동은 크기까지 같아야 한다.
- 대응 순서를 잘못 잡으면 같은 합동조건도 오독하게 된다.
- `세 변이 같다`는 사실만으로 임의의 두 도형이 합동이라고 할 수는 없다. 삼각형에서는 특별히 합동조건으로 성립한다.
- 강체이동과 임의의 늘이기·줄이기를 섞어 생각하면 안 된다.
## Curriculum Context

- 교육과정 배치
  - 한국 대표 배치에서는 `중2`의 핵심 기하 논증 개념이다.
  - 이후 닮음, 원, 삼각비, 좌표기하로 이어진다.
- 국가별 배치 스냅샷
  - 한국: `중2`의 합동 논증과 함께 배운다.
  - 일본: `중2`의 `합동과 증명`이 대표 단원이다.
  - 중국: `8학년`의 `합동삼각형`이 명시적 단원이다.
  - 미국: Grade 8과 Geometry에서 transformation-based congruence가 강조된다.
- 표현과 문제 감각
  - 다국어 용어: `congruence`, `合同`, `全等`
  - 합동은 단순히 `닮았음`보다 강한 조건이라는 점이 중요하다.

## Connections

- 선수 개념은 [basic-geometry-and-construction.md](./basic-geometry-and-construction.md), [geometric-proof.md](./geometric-proof.md)다.
- 다음 개념으로는 [similarity.md](./similarity.md), [circle.md](./circle.md), [vectors.md](./vectors.md)가 이어진다.
- 계통 허브에서는 [geometry-strand.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/syntheses/geometry-strand.md)와 [korea-curriculum-hub.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/syntheses/korea-curriculum-hub.md)를 함께 본다.

## Open Questions

- 변환 관점의 합동을 별도 시각화 카드로 더 보강할지 검토가 필요하다.

## Sources

- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-concept-encyclopedia/japan.md`
- `docs/math-concept-encyclopedia/china.md`
- `docs/math-concept-encyclopedia/us.md`
- `docs/math-concept-encyclopedia/multilingual-glossary.md`
