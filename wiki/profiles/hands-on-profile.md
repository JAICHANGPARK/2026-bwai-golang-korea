---
title: Hands-On Profile
type: profile
status: active
updated: 2026-04-15
source_docs:
  - docs/hands-on-math-agent-session-guide.md
  - docs/prompt-pack/README.md
  - docs/llm-wiki-usage-guide.md
  - docs/antigravity-math-agent-step-by-step.md
  - docs/gemini-cli-math-agent-step-by-step.md
  - docs/antigravity-workspace-rules-draft.md
  - docs/antigravity-workflow-usage-guide.md
  - docs/workspace-distribution-guide.md
tags:
  - hands-on
  - workshop
  - profile
---

# Hands-On Profile

## Summary

`hands-on profile`은 참가자가 약 1시간 안에 수학 문제 해설 시스템의 구조를 설계하고, 그 설계를 `AI Studio Build`와 `Antigravity`에 재사용해보는 체험형 범위다.

## Key Points

- 핵심 학습 포인트는 구현보다 설계다.
  - 사용자 흐름
  - Solver / Verifier 역할 분리
  - 시스템 프롬프트
  - 구조화 출력
  - 스트리밍 및 Markdown UX
- 제약은 명확하다.
  - `RAG 사용 안 함`
  - `두 에이전트만 사용`
  - `메인 모델은 gemini-3-flash-preview`
  - `검증 승인 후에만 최종 설명 표시`
- 세션 산출물은 프롬프트 팩과 통합 명세 프롬프트로 이어진다.
- 세션 이후의 로컬 활용 경로도 만들 수 있다.
  - 저장소 clone
  - `Gemini CLI` 또는 `Antigravity`로 wiki-first 질의
  - ingest와 lint 실습
  - 공유 workspace는 공개용 자산과 runnable starter scaffold를 분리하는 편이 낫다.
- `Antigravity` 전용 진행안도 만들 수 있다.
  - 첫 프롬프트
  - Phase 1
  - Phase 2
  - Phase 3
  - 마지막 polish
  - workspace rules draft
  - workspace workflows
- `Gemini CLI` 전용 진행안도 만들 수 있다.
  - context file 기반 시작
  - Phase 1
  - Phase 2
  - Phase 3
  - 마지막 polish

## Connections

- 운영 가이드는 [../source-notes/hands-on-math-agent-session-guide.md](../source-notes/hands-on-math-agent-session-guide.md)에 요약되어 있다.
- 단계별 프롬프트 묶음은 [../components/prompt-pack.md](../components/prompt-pack.md)에서 본다.
- clone 이후 운영 흐름은 [../components/local-llm-wiki-ops.md](../components/local-llm-wiki-ops.md)에서 본다.
- Antigravity automation 구성은 [../components/antigravity-workspace-automation.md](../components/antigravity-workspace-automation.md)에서 본다.
- 제품형 차이는 [../syntheses/profile-boundary.md](../syntheses/profile-boundary.md)에서 비교한다.

## Open Questions

- 실제 세션 데모에서 어느 시점에 Build와 Antigravity를 나눠 보여줄지 운영안 보강이 가능하다.
- 스트리밍 UI와 Markdown view 예시는 별도 페이지로 확장해도 좋다.
- 참가자용 첫 세션 프롬프트 묶음을 더 짧은 handout으로 분리할지는 아직 정하지 않았다.
- `adk-python`과 `adk-go`를 같은 세션에서 모두 열지, 하나를 기본값으로 고정할지는 운영안 선택이 필요하다.

## Sources

- `docs/hands-on-math-agent-session-guide.md`
- `docs/prompt-pack/README.md`
- `docs/llm-wiki-usage-guide.md`
- `docs/antigravity-math-agent-step-by-step.md`
- `docs/gemini-cli-math-agent-step-by-step.md`
- `docs/antigravity-workspace-rules-draft.md`
- `docs/antigravity-workflow-usage-guide.md`
- `docs/workspace-distribution-guide.md`
