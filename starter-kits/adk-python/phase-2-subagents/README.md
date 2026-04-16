# Phase 2 - Terminal Subagents

목표는 단일 에이전트를 `문제 해석`, `문제 풀이`, `풀이 검토` 역할로 분리해 터미널에서 검증 가능한 오케스트레이션 구조를 만드는 것이다.

## 권장 구조

```text
phase-2-subagents/
├── README.md
├── app/
│   ├── orchestrator.py
│   └── agents/
│       ├── problem_interpreter.py
│       ├── solver.py
│       ├── reviewer.py
│       └── wiki_enricher.py   # optional
└── tests/
    └── smoke_subagents.md
```

## 완료 기준

- 문제 해석, 풀이, 검토 단계가 분리된다.
- 검토 실패 시 재시도 또는 fallback 메시지를 보여준다.
- 터미널에서 단계별 intermediate result를 확인할 수 있다.

## Optional Wiki Lookup

`reviewer`가 승인한 뒤에만 `starter-kits/adk-python/shared/wiki_lookup_tool.py`를 호출해 관련 wiki page 후보를 가져온다.

이 단계의 역할은 간단하다.

- 입력: `problem_text`, `concept_candidates`, `approved_solution_summary`
- 검색: `wiki/index.md`와 관련 wiki page 최소 조회
- 출력: `query_keywords`, `matches`, `index_excerpt`

Phase 2에서는 이 결과를 터미널에 그대로 출력해도 충분하다. 중요한 건 `solver`가 아니라 `wiki_enricher` 또는 `orchestrator`가 이 lookup을 담당한다는 점이다.
