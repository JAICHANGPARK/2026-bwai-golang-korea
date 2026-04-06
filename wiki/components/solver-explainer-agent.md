---
title: Solver Explainer Agent
type: component
status: active
updated: 2026-04-06
source_docs:
  - docs/math-agent-system-spec.md
  - docs/hands-on-math-agent-session-guide.md
  - docs/prompt-pack/02-agent-design.md
tags:
  - agent
  - solver
  - explanation
---

# Solver Explainer Agent

## Summary

이 에이전트는 문제를 읽고 풀이 초안과 학생용 설명을 만든다. 핵심은 `정답을 내는 것`만이 아니라 `풀이 구조를 분해해서 교육적으로 설명하는 것`이다.

## Key Points

- 해야 할 일
  - 문제 구조 파악
  - 단계별 풀이 생성
  - 정답 도출
  - 핵심 개념 후보 제안
  - 학생 친화적 설명 작성
- 하지 말아야 할 일
  - 최종 검증자처럼 행동하기
  - 읽히지 않는 문제 조건을 추측하기
  - 논리 단계를 건너뛰기
- 제품형 확장에서는 필요 시 Code Execution과 RAG 보강을 함께 사용한다.

## Connections

- 검증 책임은 [expert-verifier-agent.md](./expert-verifier-agent.md)에 분리되어 있다.
- 핸즈온 범위에서는 이 에이전트의 역할 정의가 프롬프트 설계의 출발점이다.
- 제품형 시스템에서는 워크플로 컨트롤러가 이 에이전트를 먼저 호출한다.

## Open Questions

- 구조화 출력 스키마의 최종 형태는 아직 별도 페이지로 정리되지 않았다.
- 학생 수준별 설명 난이도 조절 규칙은 후속 문서화가 필요하다.

## Sources

- `docs/math-agent-system-spec.md`
- `docs/hands-on-math-agent-session-guide.md`
- `docs/prompt-pack/02-agent-design.md`
