---
title: 평행선·다각형·합동의 논증
type: synthesis
status: active
updated: 2026-04-10
card_role: concept
source_docs:
  - docs/math-concept-encyclopedia/korea.md
  - docs/math-concept-encyclopedia/japan.md
  - docs/math-concept-encyclopedia/china.md
  - docs/math-concept-encyclopedia/us.md
tags:
  - concept
  - geometry
  - proof
---

# 평행선·다각형·합동의 논증

## Summary

평행선·다각형·합동의 논증은 기하 명제가 왜 참인지 말로 설명하고 근거를 연결하는 개념이다. 계산 중심의 수학에서 벗어나 `가정 -> 추론 -> 결론` 구조를 익히는 첫 본격 증명 카드다.

## Key Points

- 정의
  - 평행선, 각의 관계, 다각형의 성질, 합동 조건을 논리적으로 연결해 결론을 보이는 과정을 기하 논증이라 한다.
- 핵심 개념
  - 가정
  - 결론
  - 동위각
  - 엇각
  - 내각과 외각
  - 합동조건
- 대표 수식
  - 삼각형 내각의 합은 $180^\circ$다.
  - 평행선에서는 동위각과 엇각의 관계가 핵심 근거가 된다.
- 성질 1
  - 삼각형의 외각은 이웃하지 않는 두 내각의 합과 같다.
- 성질 2
  - 이등변삼각형의 두 밑각은 같다.
- 성질 3
  - 평행선과 한 횡단선이 만들 각은 대응 관계로 서로 같거나 합이 $180^\circ$가 된다.
- 증명 스케치
  - `증명 스케치 (추론)`:
  - 삼각형의 한 꼭짓점을 지나 밑변과 평행한 직선을 그리면, 평행선의 각 관계로 세 내각이 한 직선 위에 놓인 세 각과 대응된다.
  - 직선 위의 각의 합은 $180^\circ$이므로 삼각형 내각의 합도 $180^\circ$다.
  - 같은 논리로 외각 정리와 이등변삼각형의 밑각 성질도 증명할 수 있다.
## Deep Dive

- 기하 증명은 `같다`, `평행이다`, `합동이다`를 연결하는 언어다.
- 핵심은 이미 아는 성질을 새로운 그림에 옮겨 쓰는 것이다.
- 보조선을 그어 평행선을 만들거나, 합동인 삼각형을 찾아 각과 변을 넘겨 쓰는 방식이 자주 쓰인다.
- 증명은 보통 `가정 -> 근거 -> 결론` 순서로 적는다.

## Worked Examples

### 예제 1: 삼각형 내각의 합

- 한 꼭짓점을 지나 밑변과 평행한 직선을 그으면 세 내각이 한 직선 위의 각들과 대응된다.
- 따라서 삼각형 내각의 합은
  $$
  180^\circ
  $$
  이다.

### 예제 2: 외각 정리

- 삼각형의 한 외각은 그와 이웃하지 않는 두 내각의 합과 같다.

### 예제 3: 이등변삼각형의 밑각

- 두 변이 같은 이등변삼각형에서 꼭짓점을 지나는 선을 그어 합동인 두 삼각형을 만들면 밑각이 같음을 보일 수 있다.

## Common Pitfalls

- 결론만 적고 근거를 생략하면 증명이 아니다.
- 평행선의 각 관계에서 동위각, 엇각, 같은 쪽 내각을 헷갈리기 쉽다.
- 합동 조건을 쓸 때 대응을 잘못 잡으면 증명이 무너진다.
- 보조선을 그었는데 왜 그 보조선이 필요한지 설명하지 않으면 논리 구조가 약해진다.
- 교육과정 배치
  - 한국 대표 배치에서는 `중2`의 논증 중심 단원이다.
- 국가별 배치 스냅샷
  - 한국: `중2`에서 평행선, 다각형, 합동의 논증이 함께 묶인다.
  - 일본: `중2`의 `합동과 증명`이 강한 대응 축이다.
  - 중국: `교차선과 평행선`, `합동삼각형`이 분리되어 있지만 논증 문화는 일찍 등장한다.
  - 미국: Grade 8과 Geometry에서 proof, congruence, similarity가 분산 연결된다.
- 표현과 문제 감각
  - 한국·일본은 서술형 증명 문화가 강하고, 미국은 변환 관점의 설명이 상대적으로 강한 편이다.

## Connections

- 선수 개념은 [basic-geometry-and-construction.md](./basic-geometry-and-construction.md)다.
- 다음 개념으로는 [congruence.md](./congruence.md), [similarity.md](./similarity.md), [trigonometric-ratio.md](./trigonometric-ratio.md), [sets-and-propositions.md](./sets-and-propositions.md)가 이어진다.
- 계통 허브에서는 [geometry-strand.md](../geometry-strand.md), [korea-curriculum-hub.md](../korea-curriculum-hub.md)를 함께 본다.

## Open Questions

- `증명 문장 템플릿`을 별도 학습 카드로 뺄지 검토가 필요하다.

## Sources

- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-concept-encyclopedia/japan.md`
- `docs/math-concept-encyclopedia/china.md`
- `docs/math-concept-encyclopedia/us.md`
