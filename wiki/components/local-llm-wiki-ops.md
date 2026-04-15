---
title: Local LLM Wiki Ops
type: component
status: active
updated: 2026-04-15
source_docs:
  - docs/llm-wiki-usage-guide.md
  - docs/antigravity-math-agent-step-by-step.md
  - docs/gemini-cli-math-agent-step-by-step.md
  - docs/antigravity-workspace-rules-draft.md
  - docs/antigravity-workflow-usage-guide.md
  - docs/workspace-distribution-guide.md
tags:
  - component
  - llm-wiki
  - local-ops
  - usage
---

# Local LLM Wiki Ops

## Summary

이 컴포넌트는 저장소를 로컬로 clone한 뒤 `Gemini CLI` 또는 `Antigravity`로 이 workspace를 실제 `LLM Wiki`처럼 다루는 운영 흐름을 정리한다.

## Key Points

- 시작점은 도구가 아니라 workspace 해석이다.
  - `docs/`는 raw sources
  - `wiki/`는 persistent wiki
  - `AGENTS.md`는 운영 규칙
- `Gemini CLI` 경로
  - 프로젝트 `.gemini/settings.json`에서 `AGENTS.md`를 context file로 지정한다.
  - 첫 세션에서 `README.md`, `wiki/README.md`, `wiki/index.md`를 명시적으로 읽게 하는 프롬프트가 유용하다.
  - `Phase 1 -> Phase 2 -> Phase 3`를 프롬프트로 고정하는 편이 안정적이다.
- `Antigravity` 경로
  - workspace skill로 `raw sources + persistent wiki` 규칙을 고정한다.
  - workspace `Rules`로 항상 적용되는 제약을 고정할 수 있다.
  - workspace `Workflows`로 `Phase 1 -> Phase 2 -> Phase 3` 순서를 재사용할 수 있다.
  - `math-agent-hands-on-starter` 스킬로 `Phase 1 -> Phase 2 -> Phase 3` 빌드 순서를 고정한다.
  - 질의, ingest, lint 모두 같은 규칙 위에서 수행한다.
- `query -> ingest -> lint` 순서로 사용하는 것이 자연스럽다.
  - 프로토타입 구현은 `terminal single agent -> subagents -> backend + UI` 순서가 자연스럽다.
- 공유 workspace 패키징도 분리하는 편이 안정적이다.
  - `docs/hands-on-assets/`는 공개용 샘플 문제, 템플릿, quick guide만 둔다.
  - runnable scaffold는 `starter-kits/`로 분리한다.
  - facilitator 전용 운영 메모는 shared workspace 밖에서 관리한다.

## Connections

- 프로파일 맥락은 [../profiles/hands-on-profile.md](../profiles/hands-on-profile.md)에서 본다.
- prompt pack 자체는 [../components/prompt-pack.md](../components/prompt-pack.md)에 정리되어 있다.
- 사용 가이드의 원문 요약은 [../source-notes/llm-wiki-usage-guide.md](../source-notes/llm-wiki-usage-guide.md)에 있다.
- Gemini CLI 단계별 프롬프트 가이드는 [../source-notes/gemini-cli-math-agent-step-by-step.md](../source-notes/gemini-cli-math-agent-step-by-step.md)에 요약되어 있다.
- Antigravity 단계별 프롬프트 가이드는 [../source-notes/antigravity-math-agent-step-by-step.md](../source-notes/antigravity-math-agent-step-by-step.md)에 요약되어 있다.
- Antigravity automation 구성은 [../components/antigravity-workspace-automation.md](../components/antigravity-workspace-automation.md)에서 본다.

## Open Questions

- `Gemini CLI`와 `Antigravity` 중 어느 경로를 기본 예시로 둘지 선택이 필요하다.
- clone 이후 `Obsidian`까지 같이 열지, 아니면 agent UI만으로 충분한지 선택이 필요하다.
- Antigravity 첫 화면에서 스킬 이름을 얼마나 명시적으로 노출할지 정리가 필요하다.
- rules와 workflows를 어느 정도까지 기본 제공할지가 repo 성격마다 달라질 수 있다.

## Sources

- `docs/llm-wiki-usage-guide.md`
- `docs/antigravity-math-agent-step-by-step.md`
- `docs/gemini-cli-math-agent-step-by-step.md`
- `docs/antigravity-workspace-rules-draft.md`
- `docs/antigravity-workflow-usage-guide.md`
- `docs/workspace-distribution-guide.md`
