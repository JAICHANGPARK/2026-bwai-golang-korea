---
title: Profile Boundary
type: synthesis
status: active
updated: 2026-04-06
source_docs:
  - docs/hands-on-math-agent-session-guide.md
  - docs/prompt-pack/README.md
  - docs/math-agent-system-spec.md
  - docs/math-agent-knowledge-plan.md
  - docs/internal-knowledge-base-build.md
tags:
  - synthesis
  - boundary
  - profiles
---

# Profile Boundary

## Summary

이 저장소의 가장 중요한 구분선은 `핸즈온에서 일부러 제외한 것`과 `제품형에서 반드시 다시 도입해야 하는 것`을 명확히 나누는 데 있다.

## Key Points

| 항목 | Hands-On Profile | Product Expansion Profile |
| --- | --- | --- |
| 목표 | 1시간 안에 설계와 앱 생성 경험 | 실제 서비스 품질과 확장성 확보 |
| 에이전트 수 | Solver + Verifier | 멀티 에이전트 확장 가능 |
| 지식 검색 | 사용 안 함 | RAG / File Search 고려 |
| 데이터 인프라 | 제외 | 내부 지식베이스 필요 |
| 결과물 | 프롬프트, 명세, UI 초안 | 워크플로, 데이터, 검증, 검색 계층 |

- 핸즈온은 교육용 압축본이다.
- 제품형은 압축했던 복잡도를 다시 꺼내는 단계다.
- 두 축을 섞으면 세션이 과도하게 무거워지거나, 반대로 제품 설계가 지나치게 단순해진다.

## Connections

- 핸즈온 설명은 [../profiles/hands-on-profile.md](../profiles/hands-on-profile.md)에 있다.
- 제품형 설명은 [../profiles/product-expansion-profile.md](../profiles/product-expansion-profile.md)에 있다.
- 지식 인프라가 왜 제품형에 필요한지는 [knowledge-roadmap.md](./knowledge-roadmap.md)와 [../components/internal-knowledge-base.md](../components/internal-knowledge-base.md)에 연결된다.

## Open Questions

- 어떤 시점에 핸즈온 산출물을 제품형 문서로 승격할지 규칙을 둘 수 있다.
- Build와 Antigravity 산출물을 제품형 구현 명세로 재사용하는 편집 기준이 아직 없다.

## Sources

- `docs/hands-on-math-agent-session-guide.md`
- `docs/prompt-pack/README.md`
- `docs/math-agent-system-spec.md`
- `docs/math-agent-knowledge-plan.md`
- `docs/internal-knowledge-base-build.md`
