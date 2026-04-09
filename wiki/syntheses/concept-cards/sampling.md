---
title: 표본조사
type: synthesis
status: active
updated: 2026-04-10
card_role: concept
source_docs:
  - docs/math-concept-encyclopedia/japan.md
  - docs/math-concept-encyclopedia/korea.md
  - docs/math-concept-encyclopedia/us.md
  - docs/math-concept-encyclopedia/china.md
  - docs/math-concept-encyclopedia/multilingual-glossary.md
tags:
  - concept
  - statistics
  - sampling
---

# 표본조사

## Summary

표본조사는 모집단 전체를 모두 조사하지 않고 일부를 뽑아 전체를 추정하는 개념이다. 통계적 추정의 바로 앞 단계이며, `어떻게 뽑았는가`가 계산만큼 중요하다는 점을 보여 준다.

## Definition

- 모집단 전체 대신 일부 표본을 조사해 전체의 성질을 추정하는 방법을 표본조사라 한다.
- 표본조사의 핵심은 `계산식`보다 `표본을 뽑는 설계`에 있다.

## Key Ideas

- 모집단과 표본
  - 조사 대상 전체가 모집단이고, 실제로 조사한 일부가 표본이다.
- 대표성
  - 표본이 모집단의 특성을 고르게 반영해야 한다.
- 무작위성
  - 특정 집단이 과하게 들어가거나 빠지지 않도록 공정하게 뽑아야 한다.
- 편향
  - 표본이 한쪽으로 치우치면 표본 수가 커도 좋은 조사가 되지 않는다.
- 표본오차
  - 잘 뽑은 표본이라도 모집단과 완전히 똑같을 수는 없으므로 흔들림이 생긴다.

## Deep Dive

- 표본조사는 `일부만 보고 전체를 말하는 기술`이지만, 아무 일부나 보면 안 된다.
- 그래서 통계에서는 계산 이전에 `누구를`, `어떤 방식으로`, `몇 명이나` 뽑았는지를 먼저 본다.

## Theorems and Properties

- 대표성의 원칙
  - 모집단의 여러 특징이 표본에 고르게 반영될수록 좋은 추정으로 이어진다.
- 무작위 추출의 장점
  - 무작위 추출은 특정 방향의 편향을 줄이는 기본 전략이다.
- 표본 크기의 효과
  - 일반적으로 표본이 커질수록 우연한 흔들림은 줄어들지만, 편향된 표본이면 크기만 커져도 문제가 해결되지 않는다.

## Proof Sketch

- `증명 스케치 (추론)`:
- 표본이 한쪽으로 치우치면 모집단 전체의 성질을 제대로 반영하지 못한다.
- 예를 들어 학교 전체 만족도를 알고 싶은데 한 반만 조사하면, 그 반의 분위기가 전체 평균처럼 보일 위험이 크다.
- 그래서 좋은 표본조사는 계산 이전에 `대표성`과 `무작위성`을 확보하는 것이 핵심이다.

## Worked Examples

- 예제 1: 편향된 표본
  - 학교 전체 만족도를 알고 싶을 때 우수반 한 반만 조사하면 편향 가능성이 크다.
  - 여러 학년과 여러 반에서 무작위로 표본을 뽑는 편이 더 낫다.
- 예제 2: 표본 크기 비교
  - 같은 학교에서 10명만 조사한 결과와 200명을 무작위로 조사한 결과가 있다면, 보통 200명 표본 쪽이 더 안정적이다.
  - 다만 200명이 모두 한 동아리 학생이라면 여전히 좋은 표본이라고 할 수 없다.

## Common Pitfalls

- 표본이 크기만 크면 무조건 좋은 조사라고 생각하면 안 된다.
- 자발적으로 응답한 사람들만 모은 표본은 대표성이 약할 수 있다.
- 모집단과 표본을 구분하지 않고 표본값을 바로 전체의 참값처럼 말하면 안 된다.
- `무작위`를 `아무나`와 같은 뜻으로 오해하면 조사 설계가 무너진다.

## Curriculum Placement

- 한국 대표 배치에서는 고등학교 통계 추정의 전단계, 일본에서는 중3의 대표 단원이다.
- 국가별 배치 스냅샷
  - 한국: `확률과 통계`와 `실용 통계`의 전단계 개념으로 읽는 것이 자연스럽다.
  - 일본: `중3`의 `표본조사`가 명시적 단원이다.
  - 중국: 중학교 데이터 수집과 고교 통계 추정 사이의 연결 개념으로 배치된다.
  - 미국: Grade 7의 `Probability and Sampling`과 high school Statistics로 이어진다.
- 표현 메모
  - 다국어 용어: `sample survey`, `sampling`, `標本調査`, `抽样调查`

## Connections

- 선수 개념은 [data-organization.md](./data-organization.md)다.
- 다음 개념으로는 [statistical-inference.md](./statistical-inference.md), [practical-statistics.md](./practical-statistics.md), [regression.md](./regression.md)가 이어진다.
- 계통 허브로는 [../statistics-and-probability-strand.md](../statistics-and-probability-strand.md), [../high-2-hub.md](../high-2-hub.md)를 함께 본다.

## Open Questions

- `편향`을 별도 오개념 카드로 분리할지 검토가 필요하다.

## Sources

- `docs/math-concept-encyclopedia/japan.md`
- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-concept-encyclopedia/us.md`
- `docs/math-concept-encyclopedia/china.md`
- `docs/math-concept-encyclopedia/multilingual-glossary.md`
