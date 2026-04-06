---
title: Source Note - Math Agent System Spec
type: source-note
status: active
updated: 2026-04-06
source_docs:
  - docs/math-agent-system-spec.md
tags:
  - source-note
  - system-spec
---

# Source Note - Math Agent System Spec

## Source Summary

이 문서는 중고등학생 대상 수학문제 해설 서비스의 제품형 멀티 에이전트 시스템 명세를 정의한다.

## Notable Claims

- 시스템은 `Solver/Explainer`와 `Expert Verifier` 두 축을 중심으로 설계된다.
- 제품형 범위에서는 RAG, File Search, 내부 지식 DB를 고려한다.
- 핵심 실행 원칙은 `solve first, verify second, retrieve for explanation`이다.
- 검증 실패 시 사용자에게 무리한 답변을 주지 않는 보수적 정책을 따른다.

## Pages Updated

- [../profiles/product-expansion-profile.md](../profiles/product-expansion-profile.md)
- [../components/solver-explainer-agent.md](../components/solver-explainer-agent.md)
- [../components/expert-verifier-agent.md](../components/expert-verifier-agent.md)
- [../overview/project-map.md](../overview/project-map.md)

## Open Questions

- 워크플로 컨트롤러와 관측성 로깅을 별도 구성요소 페이지로 분리할 수 있다.

## Raw Source

- `docs/math-agent-system-spec.md`
