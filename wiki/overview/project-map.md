---
title: Project Map
type: overview
status: active
updated: 2026-04-15
source_docs:
  - docs/hands-on-math-agent-session-guide.md
  - docs/prompt-pack/README.md
  - docs/math-agent-system-spec.md
  - docs/math-agent-knowledge-plan.md
  - docs/internal-knowledge-base-build.md
  - docs/math-curriculum-research/README.md
  - docs/math-concept-encyclopedia/README.md
tags:
  - overview
  - project-map
---

# Project Map

## Summary

이 프로젝트는 `Build with AI Golang Korea 2026` 이벤트의 `중고등학생 대상 수학 문제 해설 시스템` 핸즈온 workspace를 중심으로 두 개의 운영 축을 다룬다. 하나는 1시간 내에 설계와 앱 생성을 경험하게 하는 `hands-on profile`, 다른 하나는 실제 서비스 확장을 염두에 둔 `product expansion profile`이다.

## Key Points

- 이 workspace의 1차 목적은 행사 핸즈온 자료를 제공하는 것이다.
  - prompt pack
  - step-by-step guide
  - shared assets
  - starter kits
- 공통 코어는 같다.
  - `Solver/Explainer Agent`
  - `Expert Verifier Agent`
  - 학생 친화적 해설과 엄격한 검증의 분리
- 핸즈온 축은 deliberately simple 하다.
  - `RAG 없음`
  - `두 에이전트만 사용`
  - `AI Studio -> Build -> Antigravity` 재사용 흐름 강조
- 제품 확장 축은 지식 인프라를 포함한다.
  - OCR 및 구조화
  - 내부 지식베이스
  - 교육과정 매핑
  - 개념 그래프와 개념 카드
  - 복습 추천
  - File Search 또는 외부 검색 계층

## Connections

- 핸즈온 정리는 [../profiles/hands-on-profile.md](../profiles/hands-on-profile.md)에서 본다.
- 제품형 범위는 [../profiles/product-expansion-profile.md](../profiles/product-expansion-profile.md)에서 본다.
- 에이전트 역할은 [../components/solver-explainer-agent.md](../components/solver-explainer-agent.md), [../components/expert-verifier-agent.md](../components/expert-verifier-agent.md)에 정리되어 있다.
- 지식 인프라는 [../components/internal-knowledge-base.md](../components/internal-knowledge-base.md)와 [../syntheses/knowledge-roadmap.md](../syntheses/knowledge-roadmap.md)에 연결된다.
- 개념 카드 구조는 [../syntheses/concept-card-model.md](../syntheses/concept-card-model.md)에서 본다.
- 프롬프트 운영은 [../components/prompt-pack.md](../components/prompt-pack.md)에서 본다.

## Open Questions

- 실제 raw source 수집 루트는 `docs/` 외에 어떤 폴더 구조로 확장할지 아직 정해지지 않았다.
- `wiki/queries/`와 `wiki/lint/`의 실제 운영 빈도는 아직 비어 있다.
- 향후 제품형 구현에서 어떤 검색 계층을 먼저 붙일지 결정이 필요하다.

## Sources

- `docs/hands-on-math-agent-session-guide.md`
- `docs/prompt-pack/README.md`
- `docs/math-agent-system-spec.md`
- `docs/math-agent-knowledge-plan.md`
- `docs/internal-knowledge-base-build.md`
- `docs/math-curriculum-research/README.md`
- `docs/math-concept-encyclopedia/README.md`
