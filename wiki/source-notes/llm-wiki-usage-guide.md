---
title: Source Note - LLM Wiki Usage Guide
type: source-note
status: active
updated: 2026-04-15
source_docs:
  - docs/llm-wiki-usage-guide.md
tags:
  - source-note
  - usage-guide
  - llm-wiki
---

# Source Note - LLM Wiki Usage Guide

## Source Summary

이 문서는 저장소를 `git clone`한 뒤 `Gemini CLI` 또는 `Antigravity`로 `wiki/`를 읽고, 질의하고, ingest하고, lint하는 실제 운영 흐름을 안내한다.

## Notable Claims

- 먼저 `README.md`, `AGENTS.md`, `wiki/index.md`를 읽고 workspace 구조를 이해하는 것이 좋다.
- `Gemini CLI`에서는 프로젝트의 `.gemini/settings.json`을 통해 `AGENTS.md`를 context file로 사용할 수 있다.
- `Antigravity`에서는 workspace skill을 통해 이 저장소를 `raw sources + persistent wiki` 패턴으로 다루게 할 수 있다.
- 질문은 `wiki/index.md`에서 시작하고, 필요한 경우에만 `wiki/source-notes/`와 `docs/`로 내려가는 것이 권장된다.
- ingest와 lint도 별도 사용 흐름으로 정리할 수 있다.

## Pages Updated

- [../components/local-llm-wiki-ops.md](../components/local-llm-wiki-ops.md)
- [../profiles/hands-on-profile.md](../profiles/hands-on-profile.md)
- [../index.md](../index.md)

## Open Questions

- Antigravity의 skill auto-activation 동작은 버전에 따라 달라질 수 있어, 환경별 확인이 필요하다.
- `첫 프롬프트 묶음`을 더 축약할지 여부는 아직 열려 있다.

## Raw Source

- `docs/llm-wiki-usage-guide.md`
