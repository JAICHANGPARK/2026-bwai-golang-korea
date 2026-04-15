# Phase 2 - Terminal Subagents

목표는 Go 기반으로 `문제 해석`, `문제 풀이`, `풀이 검토` 역할을 분리한 오케스트레이션 CLI를 만드는 것이다.

## 권장 구조

```text
phase-2-subagents/
├── README.md
├── cmd/
│   └── terminal/
│       └── main.go
├── internal/
│   ├── orchestrator/
│   └── agents/
│       ├── problem_interpreter.go
│       ├── solver.go
│       └── reviewer.go
└── testdata/
    └── smoke_subagents.md
```

## 완료 기준

- 문제 해석, 풀이, 검토 단계가 분리된다.
- 검토 실패 시 재시도 또는 fallback 메시지를 보여준다.
- 터미널에서 단계별 intermediate result를 확인할 수 있다.
