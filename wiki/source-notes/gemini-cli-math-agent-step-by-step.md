---
title: Source Note - Gemini CLI Math Agent Prompt Flow
type: source-note
status: active
updated: 2026-04-15
source_docs:
  - docs/gemini-cli-math-agent-step-by-step.md
tags:
  - source-note
  - gemini-cli
  - prompt-flow
  - prototype
---

# Source Note - Gemini CLI Math Agent Prompt Flow

## Source Summary

이 문서는 저장소 루트에서 `gemini`를 실행하고 `AGENTS.md` 기반 context를 사용해 `Phase 1 -> Phase 2 -> Phase 3` 순서로 수학 해설 에이전트 프로토타입을 만들고, 필요하면 `wiki enricher` 확장을 추가하는 단계별 프롬프트 흐름을 정리한다.

## Notable Claims

- `Gemini CLI`에서는 Antigravity skill 대신 `AGENTS.md`와 `.gemini/settings.json`이 운영 규칙을 고정하는 역할을 한다.
- 첫 프롬프트에서 `README.md`, `AGENTS.md`, `wiki/README.md`, `wiki/index.md`를 먼저 읽게 하는 것이 중요하다.
- 구현 순서는 `터미널 단일 에이전트 -> 터미널 subagent orchestration -> backend + UI`가 적절하다.
- 각 phase마다 바로 복붙 가능한 프롬프트와 종료 확인 프롬프트를 분리해 두면 진행이 안정적이다.
- 필요하면 `review approved 이후`에만 `LLM Wiki`를 읽어 `learning_context`를 만드는 `wiki enricher` 확장을 추가할 수 있다.
- `Gemini CLI`에서는 중간중간 상태 요약 프롬프트를 넣어 컨텍스트를 정리하는 것이 특히 유용하다.

## Pages Updated

- [../components/local-llm-wiki-ops.md](../components/local-llm-wiki-ops.md)
- [../profiles/hands-on-profile.md](../profiles/hands-on-profile.md)
- [../index.md](../index.md)

## Open Questions

- `Gemini CLI`와 `Antigravity` 중 어느 쪽을 기본 예시로 둘지 선택이 필요하다.
- `adk-go` 선택 트랙까지 같은 문서에 직접 포함할지는 아직 열려 있다.

## Raw Source

- `docs/gemini-cli-math-agent-step-by-step.md`
