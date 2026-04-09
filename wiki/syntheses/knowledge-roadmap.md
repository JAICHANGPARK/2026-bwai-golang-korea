---
title: Knowledge Roadmap
type: synthesis
status: active
updated: 2026-04-06
source_docs:
  - docs/math-agent-knowledge-plan.md
  - docs/internal-knowledge-base-build.md
  - docs/math-agent-system-spec.md
  - docs/math-curriculum-research/README.md
  - docs/math-concept-encyclopedia/README.md
tags:
  - synthesis
  - roadmap
  - knowledge
---

# Knowledge Roadmap

## Summary

제품형 수학 해설 시스템에서 지식 구축은 `자료를 많이 모으는 것`보다 `무엇을 어떤 층으로 정규화하고, 어떤 검색 단위로 제공할지`를 정하는 일에 가깝다.

## Key Points

- 1차 우선순위
  - 교육과정 원문
  - 성취기준
  - 단원/개념 체계
  - 나라별 커리큘럼 연구 문서
  - 개념 백과사전과 예제 문서
  - 공식 평가 문항 메타데이터
  - 오답 포인트 카드
  - 복습 추천 카드
- 구축 순서
  1. 원천 문서를 수집한다.
  2. 성취기준, 개념, 단원 같은 정규화 엔터티를 만든다.
  3. 사람 읽기용 `개념 카드`와 허브 문서를 위키에 만든다.
  4. 검색에 적합한 카드형 청크를 만든다.
  5. 해설 에이전트와 검증 에이전트가 공통 근거로 참조하게 한다.
- 중요한 데이터 설계 포인트
  - 교육과정 버전 구분
  - 학교급/과목/단원 메타데이터
  - 검수 규칙
  - 저작권과 공개 범위

## Connections

- 내부 지식 계층의 역할 정의는 [../components/internal-knowledge-base.md](../components/internal-knowledge-base.md)에 있다.
- 전체 제품 범위는 [../profiles/product-expansion-profile.md](../profiles/product-expansion-profile.md)에 있다.
- 원천 논의는 [../source-notes/math-agent-knowledge-plan.md](../source-notes/math-agent-knowledge-plan.md), [../source-notes/internal-knowledge-base-build.md](../source-notes/internal-knowledge-base-build.md)에 나뉘어 있다.
- 개념 문서 구조는 [./concept-card-model.md](./concept-card-model.md)에서 이어진다.

## Open Questions

- AI-Hub, KICE, 교육과정 문서를 어떤 ingest cadence로 넣을지 아직 없다.
- 검색용 카드와 서비스 DB의 경계를 어느 도구에서 구현할지 미정이다.

## Sources

- `docs/math-agent-knowledge-plan.md`
- `docs/internal-knowledge-base-build.md`
- `docs/math-agent-system-spec.md`
- `docs/math-curriculum-research/README.md`
- `docs/math-concept-encyclopedia/README.md`
