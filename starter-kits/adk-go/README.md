# ADK Go Starter Kit

`adk-go`를 사용하는 대안 런타임 starter scaffold 초안이다.

## 포함 단계

- `phase-1-terminal/`
  - 텍스트 입력 기반 단일 에이전트
- `phase-2-subagents/`
  - 문제 해석, 풀이, 검토 에이전트 오케스트레이션
- `phase-3-fullstack/`
  - 이미지 입력과 화면까지 포함한 프로토타입
- `shared/`
  - runtime context 빌더로 재사용할 `wiki lookup helper`

## 권장 사용법

1. `phase-1-terminal/README.md`에서 최소 CLI 구조를 만든다.
2. 같은 저장소에서 `phase-2-subagents/README.md` 요구사항을 추가한다.
3. 마지막으로 `phase-3-fullstack/README.md` 구조로 확장한다.

## Wiki Lookup Interface

`LLM Wiki`를 붙이고 싶다면 `starter-kits/adk-go/shared/wiki_lookup_tool.go`를 로컬 helper나 ADK tool wrapper로 사용한다.

고정 인터페이스는 아래처럼 잡으면 된다.

```go
result, err := wikilookup.LookupWikiContext(
    "2x + 3 = 11을 풀어라.",
    "양변에서 3을 빼고 2로 나누면 x=4이다.",
    []string{"일차방정식", "이항"},
    "KR-2022-middle",
    4,
    ".",
)
```

이 함수는 아래 역할만 맡는다.

- 입력: 수학 문제와 개념 후보
- 검색: `wiki/index.md`와 관련 wiki page 최소 조회
- 출력: `matches`와 `index_excerpt`를 포함한 compact evidence bundle

그다음 `Wiki Knowledge Enricher`가 이 결과를 받아 `learning_context`를 만든다.
