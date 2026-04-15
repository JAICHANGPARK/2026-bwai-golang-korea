---
title: Source Note - Prompt Pack Wiki Enricher Extension
type: source-note
status: active
updated: 2026-04-15
source_docs:
  - docs/prompt-pack/02a-wiki-enricher-extension.md
tags:
  - source-note
  - prompt-pack
  - wiki-enricher
---

# Source Note - Prompt Pack Wiki Enricher Extension

## Source Summary

이 문서는 기본 `Solver/Explainer + Expert Verifier` 구조가 안정적으로 돌아간 뒤, `LLM Wiki`를 참고해 난이도, 핵심 개념, 교육과정, 연관 학습 주제를 보강하는 `Wiki Knowledge Enricher` 확장 에이전트 흐름을 정리한다.

## Notable Claims

- `wiki enricher`는 정답을 맞히는 용도가 아니라 승인된 풀이를 학습 정보로 보강하는 용도다.
- `reviewer approved 이후`에만 실행하는 것이 핵심 원칙이다.
- `ADK`에서는 `approved_solution`과 `learning_context`를 분리해 final explainer에 주입하는 구조가 자연스럽다.
- 참가자용 prompt instruction에는 `wiki/index.md -> 관련 페이지 최소 조회 -> learning_context 생성` 흐름을 명시하는 편이 안정적이다.

## Pages Updated

- [../components/prompt-pack.md](../components/prompt-pack.md)
- [../components/wiki-enricher-agent.md](../components/wiki-enricher-agent.md)
- [../profiles/hands-on-profile.md](../profiles/hands-on-profile.md)

## Open Questions

- 핸즈온 본 세션에 `wiki enricher`를 기본 트랙으로 넣을지, 확장 트랙으로만 둘지는 아직 선택 여지가 있다.

## Raw Source

- `docs/prompt-pack/02a-wiki-enricher-extension.md`
