---
title: Wiki Enricher Agent
type: component
status: active
updated: 2026-04-15
source_docs:
  - docs/prompt-pack/02a-wiki-enricher-extension.md
  - docs/internal-knowledge-base-build.md
tags:
  - agent
  - enrichment
  - wiki
  - curriculum
---

# Wiki Enricher Agent

## Summary

이 에이전트는 `reviewer`가 승인한 풀이를 바탕으로 `LLM Wiki`에서 최소한의 관련 페이지를 읽고, 난이도, 핵심 개념, 교육과정, 연관 학습 주제를 구조화된 `learning_context`로 압축한다.

## Key Points

- 해야 할 일
  - `wiki/index.md`에서 관련 페이지 후보 좁히기
  - 관련 허브와 개념 카드 최소 조회
  - 난이도 추정
  - 핵심 개념 추출
  - 교육과정 후보 연결
  - 연관 학습 주제와 참고 wiki page 기록
- 하지 말아야 할 일
  - 정답을 맞히기 위해 wiki를 사용하기
  - reviewer 승인 전 실행하기
  - wiki 전체를 final context에 그대로 밀어 넣기
  - 근거가 약한 매핑을 확정처럼 말하기
- 권장 출력
  - `difficulty`
  - `concepts`
  - `curriculum_candidates`
  - `related_topics`
  - `wiki_basis_pages`

## Connections

- 정답과 풀이 초안은 [solver-explainer-agent.md](./solver-explainer-agent.md)가 만든다.
- 승인 여부는 [expert-verifier-agent.md](./expert-verifier-agent.md)가 결정한다.
- 이 에이전트가 만든 결과는 `learning_context`로 final explainer에 넘어간다.
- 지식 계층의 역할 정의는 [internal-knowledge-base.md](./internal-knowledge-base.md)에 있다.

## Open Questions

- 난이도를 `easy/medium/hard` 같은 단순 라벨로 둘지, 평가 맥락별 다중 축으로 둘지는 추가 설계가 필요하다.
- 교육과정 매핑 기준을 한국 단일 기준으로 둘지, 국가별 프로파일 입력으로 넓힐지는 아직 열려 있다.

## Sources

- `docs/prompt-pack/02a-wiki-enricher-extension.md`
- `docs/internal-knowledge-base-build.md`
