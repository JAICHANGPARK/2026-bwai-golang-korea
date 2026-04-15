---
title: Source Note - Antigravity Math Agent Prompt Flow
type: source-note
status: active
updated: 2026-04-15
source_docs:
  - docs/antigravity-math-agent-step-by-step.md
tags:
  - source-note
  - antigravity
  - prompt-flow
  - prototype
---

# Source Note - Antigravity Math Agent Prompt Flow

## Source Summary

이 문서는 Antigravity에서 이 저장소를 열고 `llm-wiki-workspace`와 `math-agent-hands-on-starter` 스킬을 사용해 `Phase 1 -> Phase 2 -> Phase 3` 순서로 수학 해설 에이전트 프로토타입을 만드는 단계별 프롬프트 흐름을 정리한다.

## Notable Claims

- 첫 프롬프트에서 두 스킬을 함께 지정하고 `Phase 1 only`를 고정하는 것이 중요하다.
- 구현 순서는 `터미널 단일 에이전트 -> 터미널 subagent orchestration -> backend + UI`가 적절하다.
- 각 phase마다 바로 복붙 가능한 프롬프트와 종료 확인 프롬프트를 분리해 두면 진행이 안정적이다.
- 가장 중요한 제약은 `RAG 없음`, `Phase 건너뛰지 않음`, `review 승인 전 최종 해설 금지`다.
- 마지막에는 정상 텍스트, 까다로운 텍스트, 이미지, fallback 사례까지 점검하는 흐름이 권장된다.

## Pages Updated

- [../components/local-llm-wiki-ops.md](../components/local-llm-wiki-ops.md)
- [../profiles/hands-on-profile.md](../profiles/hands-on-profile.md)
- [../index.md](../index.md)

## Open Questions

- Antigravity UI에서 스킬 선택 노출 방식이 환경마다 얼마나 일관적인지는 추가 확인이 필요하다.
- `adk-go` 선택 트랙까지 같은 문서에 직접 포함할지는 아직 열려 있다.

## Raw Source

- `docs/antigravity-math-agent-step-by-step.md`
