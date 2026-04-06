# Fallback Prompt 템플릿

생성 결과가 흔들릴 때 바로 넣을 수 있는 짧은 보정 프롬프트 모음이다.

## 1. 범위 축소

```text
Simplify the app.
This is a 1-hour workshop demo.
Keep only:
- image upload
- solve button
- solver result card
- verifier result card
- final approved explanation card
Remove any extra complexity.
```

## 2. JSON 강제

```text
Make both agents return strict JSON only.
Do not include prose outside the JSON object.
```

## 3. 최종 설명 계약 재고정

```text
Do not show a final explanation before verifier approval.
If verdict is revised, show verifier comments and a review-in-progress state.
Stream only the final approved explanation.
```

## 4. UI 단순화

```text
Reduce the UI to a single clean page with:
- upload area
- question input
- solve button
- solver card
- verifier card
- final explanation card
Keep it simple for beginners.
```

## 5. Streaming 및 Markdown 재고정

```text
Use streaming only for the final approved explanation.
Render that explanation with a Markdown component.
Do not stream raw JSON.
```
