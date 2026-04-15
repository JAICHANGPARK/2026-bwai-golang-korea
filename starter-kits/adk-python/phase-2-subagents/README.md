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
│       └── reviewer.py
└── tests/
    └── smoke_subagents.md
```

## 완료 기준

- 문제 해석, 풀이, 검토 단계가 분리된다.
- 검토 실패 시 재시도 또는 fallback 메시지를 보여준다.
- 터미널에서 단계별 intermediate result를 확인할 수 있다.
