---
title: Expert Verifier Agent
type: component
status: active
updated: 2026-04-06
source_docs:
  - docs/math-agent-system-spec.md
  - docs/hands-on-math-agent-session-guide.md
  - docs/prompt-pack/02-agent-design.md
tags:
  - agent
  - verifier
  - validation
---

# Expert Verifier Agent

## Summary

이 에이전트는 Solver의 결과를 엄격하게 재검토하고, 승인 여부를 결정한다. 유창한 설명보다 수학적 정확성과 조건 누락 여부를 우선한다.

## Key Points

- 해야 할 일
  - 정답 정확성 점검
  - 계산 오류 점검
  - 논리적 비약 탐지
  - 문제 조건 누락 탐지
  - `approved`, `revised`, `needs_clarification` 판정
- 하지 말아야 할 일
  - 설명이 그럴듯하다는 이유로 승인하기
  - Solver 역할까지 직접 흡수하기
  - 불확실한 상태를 감춘 채 통과시키기
- 핸즈온과 제품형 모두에서 최종 사용자 노출 전 품질 게이트 역할을 맡는다.

## Connections

- 대응되는 생성 책임은 [solver-explainer-agent.md](./solver-explainer-agent.md)에 있다.
- 제품형 시스템에서는 검산 결과와 RAG 근거까지 함께 검토한다.
- 핸즈온에서는 `최종 설명은 검증 승인 후에만 표시` 규칙의 핵심 축이다.

## Open Questions

- 수정 요청 시 재검토 회수와 재시도 정책은 아직 별도 정리가 필요하다.
- 검증 근거를 사용자에게 얼마나 노출할지는 제품 정책으로 더 정교화할 수 있다.

## Sources

- `docs/math-agent-system-spec.md`
- `docs/hands-on-math-agent-session-guide.md`
- `docs/prompt-pack/02-agent-design.md`
