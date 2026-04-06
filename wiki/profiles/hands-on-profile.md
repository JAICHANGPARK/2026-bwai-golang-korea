---
title: Hands-On Profile
type: profile
status: active
updated: 2026-04-06
source_docs:
  - docs/hands-on-math-agent-session-guide.md
  - docs/prompt-pack/README.md
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

## Connections

- 운영 가이드는 [../source-notes/hands-on-math-agent-session-guide.md](../source-notes/hands-on-math-agent-session-guide.md)에 요약되어 있다.
- 단계별 프롬프트 묶음은 [../components/prompt-pack.md](../components/prompt-pack.md)에서 본다.
- 제품형 차이는 [../syntheses/profile-boundary.md](../syntheses/profile-boundary.md)에서 비교한다.

## Open Questions

- 실제 세션 데모에서 어느 시점에 Build와 Antigravity를 나눠 보여줄지 운영안 보강이 가능하다.
- 스트리밍 UI와 Markdown view 예시는 별도 페이지로 확장해도 좋다.

## Sources

- `docs/hands-on-math-agent-session-guide.md`
- `docs/prompt-pack/README.md`
