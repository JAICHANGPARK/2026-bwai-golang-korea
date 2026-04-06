---
title: LLM Wiki README
type: overview
status: active
updated: 2026-04-06
source_docs:
  - docs/hands-on-math-agent-session-guide.md
  - docs/math-agent-system-spec.md
  - docs/math-agent-knowledge-plan.md
  - docs/internal-knowledge-base-build.md
tags:
  - wiki
  - operations
---

# LLM Wiki

이 폴더는 원천 문서인 `docs/`를 읽고 LLM이 유지하는 구조화된 지식 계층이다.

## 빠른 사용법

- 전체 구조 확인: [`index.md`](./index.md)
- 최근 작업 확인: [`log.md`](./log.md)
- 프로젝트 지도 확인: [`overview/project-map.md`](./overview/project-map.md)

## 현재 범위

- `hands-on profile`
  - 1시간 세션 운영, 프롬프트 팩, Build/Antigravity 재사용 흐름
- `product expansion profile`
  - 멀티 에이전트 시스템, 내부 지식베이스, RAG/검색 확장

## 운영 규칙

- 위키 운영 규칙은 루트의 [`../AGENTS.md`](../AGENTS.md)에 정리되어 있다.
- `docs/`는 기본적으로 원천 문서로 보고, 위키에서만 구조화와 합성을 수행한다.
- ingest를 할 때는 source note, 관련 주제 페이지, index, log를 함께 갱신한다.

## CLI Helpers

- `scripts/wiki-search.sh "query"`
- `scripts/wiki-recent.sh 5`
- `scripts/wiki-lint.sh`
