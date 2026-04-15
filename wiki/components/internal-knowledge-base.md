---
title: Internal Knowledge Base
type: component
status: active
updated: 2026-04-09
source_docs:
  - docs/internal-knowledge-base-build.md
  - docs/math-agent-knowledge-plan.md
  - docs/math-curriculum-research/README.md
  - docs/math-concept-encyclopedia/README.md
tags:
  - knowledge-base
  - rag
  - data
---

# Internal Knowledge Base

## Summary

내부 지식베이스는 정답을 대신 추정하는 시스템이 아니라, 풀이 이후의 개념 설명, 교육과정 연결, 오답 포인트 제안, 복습 추천을 보강하는 구조화 지식 계층이다.

## Key Points

- 역할 범위
  - 개념 설명 보강
  - 성취기준 연결
  - 오답 포인트 제안
  - 복습 주제 추천
  - `learning_context` 같은 구조화 메타데이터 생성
- 금지해야 할 사용
  - 유사 문항 검색으로 정답을 대체
  - 검색 결과를 그대로 정답처럼 복사
- 데이터는 세 층으로 분리한다.
  - 원천 데이터
  - 정규화 데이터
  - 검색용 청크
- 현재의 유력한 raw source 패밀리
  - `docs/math-curriculum-research/`
  - `docs/math-concept-encyclopedia/`
- 위키에서는 `개념 카드`와 `과목·경로 카드`를 사람 읽기용 정규화 레이어로 둘 수 있다.
  - 개념 카드는 학년 허브만으로 표현되지 않는 선수 관계와 재등장 관계를 모은다.
  - 과목·경로 카드는 `공통수학1`, `Algebra I`, `Integrated Mathematics`처럼 실제 수업 단위의 묶음을 보존한다.
  - 필요하면 이후 검색용 카드와 DB 엔터티로 더 분해한다.
- 핵심 지식 타입 예시
  - `CurriculumDocument`
  - `AchievementStandard`
  - `ConceptCard`
  - `CourseCard`
  - `StrategyCard`
  - `ErrorPatternCard`
  - `ProblemMeta`
  - `CurriculumMapping`

## Connections

- 왜 이런 데이터가 필요한지는 [../syntheses/knowledge-roadmap.md](../syntheses/knowledge-roadmap.md)에서 이어진다.
- 제품형 시스템의 상위 맥락은 [../profiles/product-expansion-profile.md](../profiles/product-expansion-profile.md)에 있다.
- 원문 요약은 [../source-notes/internal-knowledge-base-build.md](../source-notes/internal-knowledge-base-build.md)를 본다.
- 개념 문서 구조안은 [../syntheses/concept-card-model.md](../syntheses/concept-card-model.md)에 있다.
- 핸즈온 확장 에이전트 관점에서는 [wiki-enricher-agent.md](./wiki-enricher-agent.md)가 이 지식 계층을 읽어 `learning_context`를 만든다.

## Open Questions

- 정규화 계층을 파일 중심으로 둘지 DB 중심으로 둘지 아직 미정이다.
- 관계 계층을 그래프로 분리할지 룰 엔진으로 둘지 후속 선택이 필요하다.

## Sources

- `docs/internal-knowledge-base-build.md`
- `docs/math-agent-knowledge-plan.md`
- `docs/math-curriculum-research/README.md`
- `docs/math-concept-encyclopedia/README.md`
