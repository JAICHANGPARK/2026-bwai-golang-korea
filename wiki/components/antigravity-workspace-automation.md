---
title: Antigravity Workspace Automation
type: component
status: active
updated: 2026-04-15
source_docs:
  - docs/antigravity-workspace-rules-draft.md
  - docs/antigravity-workflow-usage-guide.md
tags:
  - component
  - antigravity
  - rules
  - workflows
---

# Antigravity Workspace Automation

## Summary

이 컴포넌트는 이 저장소를 Antigravity에서 사용할 때 `Rules`, `Workflows`, `Skills`를 어떤 역할로 나눠 구성하는 것이 적절한지 정리한다.

## Key Points

- `Rules`
  - 항상 적용되는 경계와 제약
  - `docs/`와 `wiki/` 역할
  - 수학 해설 프로토타입 기본 범위
- `Workflows`
  - `Phase 1`, `Phase 2`, `Phase 3`
  - `smoke test`
  - `polish`
- `Skills`
  - `llm-wiki-workspace`
  - `math-agent-hands-on-starter`
- 가장 실용적인 패턴
  - rules는 항상 켠다
  - phase 시작은 workflow로 한다
  - 세부 수정은 직접 프롬프트로 한다

## Connections

- 로컬 사용 흐름은 [../components/local-llm-wiki-ops.md](../components/local-llm-wiki-ops.md)에서 본다.
- 프로토타입 맥락은 [../profiles/hands-on-profile.md](../profiles/hands-on-profile.md)에서 본다.
- rules 초안은 [../source-notes/antigravity-workspace-rules-draft.md](../source-notes/antigravity-workspace-rules-draft.md)에 요약되어 있다.
- workflow 사용 가이드는 [../source-notes/antigravity-workflow-usage-guide.md](../source-notes/antigravity-workflow-usage-guide.md)에 요약되어 있다.

## Open Questions

- workflow 이름을 더 짧게 줄일지, 현재처럼 의미를 명확히 드러낼지는 실제 사용성 확인이 필요하다.
- rules를 둘로 분리할지 하나로 합칠지는 workspace 성격에 따라 달라질 수 있다.

## Sources

- `docs/antigravity-workspace-rules-draft.md`
- `docs/antigravity-workflow-usage-guide.md`
