# Phase 1 - Terminal Single Agent

목표는 Go 기반으로 텍스트 수학 문제를 입력받아 단일 에이전트 결과를 터미널에서 확인하는 최소형 CLI를 만드는 것이다.

## 권장 구조

```text
phase-1-terminal/
├── README.md
├── cmd/
│   └── terminal/
│       └── main.go
├── internal/
│   └── agent/
│       └── prompts.go
└── testdata/
    └── smoke_text_cases.md
```

## 완료 기준

- 텍스트 문제 입력을 받는다.
- 정답과 단계별 해설을 텍스트로 출력한다.
- 샘플 문제 2개 이상으로 CLI smoke test가 가능하다.
