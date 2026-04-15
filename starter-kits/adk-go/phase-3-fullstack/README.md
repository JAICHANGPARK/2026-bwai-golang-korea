# Phase 3 - Fullstack UI

목표는 Go backend와 단일 화면 UI를 연결해 텍스트 입력과 이미지 입력을 모두 처리하는 프로토타입을 만드는 것이다.

## 권장 구조

```text
phase-3-fullstack/
├── README.md
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── api/
│   └── agents/
├── web/
│   └── src/
└── testdata/
    └── smoke_ui.md
```

## 완료 기준

- 텍스트 입력과 이미지 입력이 모두 가능하다.
- backend가 solve, review, final explanation 흐름을 노출한다.
- UI가 입력, 처리 상태, 검토 결과, 최종 해설, fallback 상태를 보여준다.
