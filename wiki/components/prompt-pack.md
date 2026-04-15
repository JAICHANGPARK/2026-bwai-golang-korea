---
title: Prompt Pack
type: component
status: active
updated: 2026-04-15
source_docs:
  - docs/prompt-pack/README.md
  - docs/hands-on-math-agent-session-guide.md
  - docs/prompt-pack/02a-wiki-enricher-extension.md
tags:
  - prompt-pack
  - workshop
  - prompts
---

# Prompt Pack

## Summary

`docs/prompt-pack/`은 핸즈온 참가자가 순서대로 복사해 쓸 수 있는 프롬프트 모음이다. 목적은 설계 산출물을 단계적으로 만들고, 그 결과를 Build와 Antigravity에 재사용하는 것이다. 기본형은 two-agent 구조를 따르고, 필요하면 `wiki enricher`를 확장 트랙으로 붙일 수 있다.

## Key Points

- 권장 순서
  - 사용자 흐름
  - 에이전트 설계
  - 선택적 `wiki enricher` 확장
  - 시스템 프롬프트
  - 화면 설계
  - 출력 스키마
  - 실패 처리
  - 통합 구현 명세
  - Build 및 Antigravity
- 적용 범위
  - shared workspace에서 바로 쓰는 설계 문서
  - 1시간 핸즈온 프로파일 기준
  - 기본 원칙은 `RAG 없음`, `두 에이전트`
  - 필요하면 승인 후 `wiki enrichment` 확장 가능
- 이 폴더는 운영 메모가 아니라 참가자용 artifacts에 가깝다.

## Connections

- 핸즈온 전체 맥락은 [../profiles/hands-on-profile.md](../profiles/hands-on-profile.md)에 있다.
- 프로젝트 지도는 [../overview/project-map.md](../overview/project-map.md)를 본다.
- 세부 소스 노트는 [../source-notes/prompt-pack-readme.md](../source-notes/prompt-pack-readme.md)에 있다.
- 확장 에이전트 정의는 [wiki-enricher-agent.md](./wiki-enricher-agent.md)에 정리되어 있다.

## Open Questions

- prompt pack 각 파일의 결과물을 위키에 어떻게 재편입할지는 아직 정해지지 않았다.

## Sources

- `docs/prompt-pack/README.md`
- `docs/hands-on-math-agent-session-guide.md`
- `docs/prompt-pack/02a-wiki-enricher-extension.md`
