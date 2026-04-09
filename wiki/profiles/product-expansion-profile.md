---
title: Product Expansion Profile
type: profile
status: active
updated: 2026-04-06
source_docs:
  - docs/math-agent-system-spec.md
  - docs/math-agent-knowledge-plan.md
  - docs/internal-knowledge-base-build.md
  - docs/math-curriculum-research/README.md
  - docs/math-concept-encyclopedia/README.md
tags:
  - product
  - architecture
  - profile
---

# Product Expansion Profile

## Summary

`product expansion profile`은 실제 서비스 수준의 수학 해설 시스템을 목표로 하며, 멀티 에이전트 워크플로와 내부 지식 계층, 교육과정 연결, 복습 추천까지 포함하는 확장 범위다.

## Key Points

- 사용자 입력은 이미지와 텍스트를 모두 포함한다.
- 시스템은 `solve first, verify second, retrieve for explanation` 원칙을 따른다.
- 설명 품질을 높이기 위해 지식 인프라가 필요하다.
  - 교육과정 원문
  - 성취기준
  - 개념 카드
  - 학년 허브와 개념 그래프
  - 오답 패턴
  - 복습 추천
- 데이터 계층은 세 층으로 나뉜다.
  - 원천 데이터
  - 정규화 데이터
  - 검색용 청크

## Connections

- 시스템 구조는 [../source-notes/math-agent-system-spec.md](../source-notes/math-agent-system-spec.md)에서 본다.
- 데이터 수집 방향은 [../source-notes/math-agent-knowledge-plan.md](../source-notes/math-agent-knowledge-plan.md)에서 본다.
- 내부 지식 설계는 [../components/internal-knowledge-base.md](../components/internal-knowledge-base.md)에 정리되어 있다.
- 구축 순서는 [../syntheses/knowledge-roadmap.md](../syntheses/knowledge-roadmap.md)로 연결된다.
- 개념 문서 구조는 [../syntheses/concept-card-model.md](../syntheses/concept-card-model.md)에서 본다.

## Open Questions

- 어떤 검색 스택을 먼저 붙일지 아직 미정이다.
- 내부 라벨링 체계와 검수 워크플로는 추가 설계가 필요하다.
- 수능 데이터, 중학교 데이터, 사진 업로드 데이터의 균형을 어떻게 맞출지 후속 결정이 필요하다.

## Sources

- `docs/math-agent-system-spec.md`
- `docs/math-agent-knowledge-plan.md`
- `docs/internal-knowledge-base-build.md`
- `docs/math-curriculum-research/README.md`
- `docs/math-concept-encyclopedia/README.md`
