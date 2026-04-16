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

## Optional Wiki Enrichment

`LLM Wiki`를 붙일 경우 backend는 `review approved 이후`에만 `starter-kits/adk-go/shared/wiki_lookup_tool.go`를 호출해 compact context seed를 만든다.

이후 흐름은 아래처럼 잡는 편이 자연스럽다.

1. `problem_interpreter`가 개념 후보를 만든다.
2. `solver`가 풀이 초안을 만든다.
3. `reviewer`가 승인 또는 재시도를 결정한다.
4. 승인되면 `wiki_enricher`가 wiki lookup helper를 호출한다.
5. `final_explainer`가 `approved_solution + learning_context`를 합쳐 응답한다.

UI에는 최종 해설 카드와 별도로 `학습 정보` 카드를 두고, wiki 기반 메타 정보를 분리해서 보여준다.
