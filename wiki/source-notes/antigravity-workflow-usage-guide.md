---
title: Source Note - Antigravity Workflow Usage Guide
type: source-note
status: active
updated: 2026-04-15
source_docs:
  - docs/antigravity-workflow-usage-guide.md
tags:
  - source-note
  - antigravity
  - workflows
---

# Source Note - Antigravity Workflow Usage Guide

## Source Summary

이 문서는 이 저장소에서 Antigravity의 workspace-level workflows를 어떤 순서와 역할 분리 위에서 사용하는 것이 적절한지 정리한다.

## Notable Claims

- 큰 흐름은 `Workflow`, 항상 적용되는 제약은 `Rule`, 재사용 지식은 `Skill`로 나누는 것이 적절하다.
- 이 저장소에는 `phase_1`, `phase_2`, `phase_3`, `smoke_test`, `polish` 다섯 가지 workflow가 있으면 충분하다.
- phase 시작은 workflow로 열고, 세부 수정은 직접 프롬프트로 하는 패턴이 실용적이다.
- rules에는 절차를 넣지 말고, workflows에 절차를 넣는 편이 낫다.

## Pages Updated

- [../components/antigravity-workspace-automation.md](../components/antigravity-workspace-automation.md)
- [../index.md](../index.md)

## Open Questions

- Antigravity command surface에서 workflow 노출 방식이 이름 기반인지 설명 기반인지에 따라 네이밍 규칙을 더 다듬을 여지가 있다.

## Raw Source

- `docs/antigravity-workflow-usage-guide.md`
