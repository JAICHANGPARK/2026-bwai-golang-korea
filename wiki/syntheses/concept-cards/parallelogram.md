---
title: 평행사변형
type: synthesis
status: active
updated: 2026-04-10
card_role: concept
source_docs:
  - docs/math-concept-encyclopedia/china.md
  - docs/math-curriculum-research/china.md
  - docs/math-concept-encyclopedia/multilingual-glossary.md
  - docs/math-concept-encyclopedia/unit-practice-book.md
tags:
  - concept
  - geometry
  - quadrilateral
---

# 평행사변형

## Summary

평행사변형은 서로 마주 보는 두 쌍의 변이 각각 평행한 사각형이다. 사각형의 성질을 논증으로 읽는 첫 대표 카드라서, 합동과 벡터, 특수사각형 이해의 중요한 중간 다리다.

## Key Points

- 정의
  - 서로 마주 보는 두 쌍의 변이 각각 평행한 사각형을 평행사변형이라 한다.
- 핵심 개념
  - 대변
  - 대각
  - 대각선
  - 성질과 판정
  - 특수사각형
- 대표 성질
  - 대변의 길이는 각각 같다.
  - 대각의 크기는 각각 같다.
  - 두 대각선은 서로를 이등분한다.
  - 한 쌍의 대변이 평행하고 길이도 같으면 평행사변형으로 판정할 수 있다.
- 대표 수식
  - $\angle A+\angle B=180^\circ$
  - 대각선의 교점을 $M$이라 하면 $AM=CM$, $BM=DM$
- 증명 스케치
  - `증명 스케치 (추론)`:
  - 평행사변형 $ABCD$에서 대각선 $AC$를 그으면, 평행선에서 생기는 엇각이 같아 두 삼각형 $\triangle ABC$, $\triangle CDA$의 각 대응을 잡을 수 있다.
  - 또 대변 $AB,CD$는 평행사변형의 구조상 서로 평행하므로 합동 또는 닮음 구조를 세울 수 있다.
  - 그 결과 대변의 길이와 대각의 크기가 각각 같다는 성질을 얻는다.

## Deep Dive

평행사변형은 단순한 도형 하나가 아니라 `직사각형`, `마름모`, `정사각형`을 모두 품는 상위 구조다. 그래서 한 도형이 평행사변형임을 먼저 보이면, 그 위에 직각 조건이나 등변 조건을 더해 특수사각형으로 세분화할 수 있다.

또한 대각선이 서로를 이등분한다는 성질은 좌표와 벡터에서도 매우 자주 쓰인다. 사각형 문제를 풀다가 중점이 반복해서 보이면, 평행사변형 가능성을 먼저 의심하는 것이 자연스럽다.

## Worked Examples

- 예제 1: 평행사변형에서 한 각이 $70^\circ$이면 이웃각은
  $$
  180^\circ-70^\circ=110^\circ
  $$
  이다.
- 예제 2: 평행사변형의 두 대각선이 만나는 점을 $M$이라 하면, 대각선이 서로를 이등분하므로 $AM=CM$, $BM=DM$이다.
- 예제 3: 한 쌍의 대변이 평행하고 길이도 같다면, 그 사각형은 평행사변형으로 판정할 수 있다.

## Common Pitfalls

- 대변이 평행하다고 해서 모든 사다리꼴이 평행사변형은 아니다.
- 대각선이 서로 같다는 성질은 모든 평행사변형에 성립하지 않는다. 그것은 직사각형 쪽 성질이다.
- 한 각이 작으면 맞은편 각이 작고, 이웃각은 보각이라는 구조를 자주 놓친다.
- 그림이 비뚤어져 보여도 성질은 평행 관계와 길이 관계로 판단해야 한다.

## Curriculum Context

- 교육과정 배치
  - 중국 대표 배치에서는 `8학년 2학기`의 명시적 기하 단원이다.
  - 이후 직사각형, 마름모, 벡터 도형, 좌표기하로 이어진다.
- 국가별 배치 스냅샷
  - 중국: `평행사변형`이 독립 소단원으로 분명하게 나타난다.
  - 한국: 현재 문서군에서는 `평행선·다각형·합동의 논증`과 특수사각형 성질 안에 흡수되어 읽힌다.
  - 일본: 중학교 기하 논증과 `도형의 성질` 축 안에서 간접 연결된다.
  - 미국: Grade 7~8 Geometry와 high school Geometry에서 quadrilateral reasoning으로 다뤄진다.

## Connections

- 선수 개념은 [congruence.md](./congruence.md), [geometric-proof.md](./geometric-proof.md), [basic-geometry-and-construction.md](./basic-geometry-and-construction.md)다.
- 다음 개념으로는 [similarity.md](./similarity.md), [vectors.md](./vectors.md), `직사각형`, `마름모`가 이어진다.
- 중국 학기 허브에서는 [china-grade-8-semester-2.md](./china-grade-8-semester-2.md), [geometry-strand.md](../geometry-strand.md)와 연결된다.

## Open Questions

- `직사각형`, `마름모`, `정사각형`을 평행사변형 카드에서 파생 카드로 더 정리할지 기준이 필요하다.

## Sources

- `docs/math-concept-encyclopedia/china.md`
- `docs/math-curriculum-research/china.md`
- `docs/math-concept-encyclopedia/multilingual-glossary.md`
- `docs/math-concept-encyclopedia/unit-practice-book.md`
