# 03a. 화면 설계 프롬프트

## 목적

핸즈온 앱의 화면 구조를 미리 설계해서, Build나 Antigravity가 너무 엉뚱한 UI를 만들지 않도록 한다.

## 언제 사용하나

에이전트 역할과 시스템 프롬프트를 정리한 뒤, 통합 명세를 만들기 전에 사용한다.

## 프롬프트

```text
Please design the UI structure for a Korean math explanation app for a 1-hour workshop.

The app is for middle school and high school students.
The user uploads a math problem image and gets:
- solver output
- verifier output
- final approved explanation

Please propose a simple single-page screen design in Korean.

The screen should include:
- a top header
- an image upload area
- a short question input
- a primary action button
- a solver result card
- a verifier result card
- a final approved explanation card rendered as Markdown
- a small status area for loading / streaming / error / clarification messages

Important constraints:
- keep the UI clean and beginner-friendly
- do not use too many sections
- prioritize readability on desktop and mobile
- make the three result cards visually distinct
- plan for the final explanation area to show partial streaming text before completion

Write the result in Korean.
```

## 기대 산출물

- 단일 페이지 구조
- 카드 기반 정보 배치
- 업로드 영역, 결과 영역, 상태 영역 분리

## 후속 프롬프트

UI 초안을 더 구체화하고 싶을 때:

```text
Please refine the screen design into a practical implementation outline.

For each section, describe:
- what the user sees
- what data is shown
- what interaction happens
- how streaming status is shown
- how the final explanation switches into a Markdown-rendered view

Keep the app as a simple single-page layout.
Write in Korean.
```
