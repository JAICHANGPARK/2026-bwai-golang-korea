# 07a. Streaming 및 Markdown 렌더링 프롬프트

## 목적

Gemini 응답을 `streaming` 방식으로 받고, 최종 설명은 `Markdown view`로 렌더링하는 UX를 Build 또는 Antigravity 결과물에 반영한다.

## 왜 필요한가

Gemini 공식 문서 기준으로:

- JavaScript SDK에서 `generateContentStream()` 사용 가능
- REST에서 `streamGenerateContent?alt=sse` 사용 가능
- Gemini 3 계열은 긴 응답과 멀티모달 작업에 적합

반면 `Markdown 렌더링`은 Gemini API 자체 기능이라기보다 앱 UI 설계 선택이다.  
즉, 스트리밍은 모델/API 기능이고, Markdown view는 프론트엔드 렌더링 전략이다.

## 언제 사용하나

- 통합 명세 프롬프트를 만든 뒤
- Build 또는 Antigravity에 구현 요구사항을 보강할 때

## 1. AI Studio Build 보강 프롬프트

```text
Please update the app to support streaming responses from Gemini.

Requirements:
- use streaming only for the final approved user-facing explanation
- show partial text as it arrives
- display a loading state while streaming
- once complete, render the final explanation using a Markdown view
- keep solver result, verifier result, and final explanation visually separated

Important notes:
- internal agent outputs may stay structured
- do not stream raw Solver or Verifier JSON
- create the streamed explanation from a FinalExplanationInput object after verifier approval
- the final approved explanation should be easy to read as Markdown
- the UI should feel responsive during streaming
- do not add new major features outside this requirement
```

## 2. Antigravity 보강 프롬프트

```text
Enhance the app with a better response UX.

Please add:
- streaming output for the final approved explanation only
- Markdown rendering for the final approved explanation
- clear status handling for:
  - generating
  - verifying
  - review in progress
  - completed
  - clarification needed

Constraints:
- keep the app simple for a workshop
- do not add RAG or backend complexity
- keep Korean UI labels
- do not expose revised content as the final answer before approval
```

## 3. 화면 설계 보강 프롬프트

```text
Please refine the UI to support streaming and Markdown viewing.

The final explanation area should:
- show a typing or streaming state while tokens arrive
- switch to a readable Markdown-rendered card when complete
- preserve line breaks, lists, formulas-as-text, and emphasis

The app should:
- keep Solver and Verifier cards structured
- show the final explanation card only after approval
- treat revised as an intermediate review state

The page should still remain simple and suitable for beginners.
```

## 4. 구현 요구사항 보강 프롬프트

```text
Please update the implementation plan with these UX requirements:

- Gemini responses should be received with streaming if supported by the runtime
- the final approved explanation should be rendered through a Markdown component/view
- during streaming, the UI should progressively append text
- the final Markdown card should be styled for readability in Korean
- the app should prepare FinalExplanationInput from approved verified data before requesting the streamed explanation
- the app should not show revised content as a final approved answer

If structured JSON and streaming conflict, prefer:
1. structured data for internal agent steps
2. streamed Markdown for the final user-facing explanation
```

## 기대 산출물

- 스트리밍 응답 UI
- Markdown 렌더링이 되는 최종 설명 영역
- loading / verifying / complete 상태 표시

## 참고 메모

공식 문서 확인 기준:

- `generateContentStream()`은 JavaScript SDK에서 사용 가능
- REST에서는 `streamGenerateContent?alt=sse` 사용 가능
- Markdown 렌더링은 앱 UI에서 처리해야 함
