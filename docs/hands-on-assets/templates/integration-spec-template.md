# 통합 명세 템플릿

아래 틀을 채운 뒤 AI Studio Build 또는 Antigravity에 넣는다.

```text
Build a simple Korean math explanation workshop app.

Product goal:
- Help middle and high school students upload a math problem image and receive a verified explanation.

Target users:
- Middle school students
- High school students

Architecture:
- Use two agents only
- Solver/Explainer Agent
- Expert Verifier Agent

User flow:
1. User uploads one math problem image and optionally types a short question.
2. Solver returns structured JSON.
3. Verifier reviews the Solver JSON and returns structured JSON.
4. Only if the verifier verdict is approved, create FinalExplanationInput.
5. Stream the final approved explanation in Korean Markdown.

Structured outputs:
- SolverResult
- VerificationResult
- FinalExplanationInput

UI requirements:
- Single page layout
- Korean labels
- Separate cards for Solver, Verifier, and Final Explanation
- Final Explanation card appears only after approval
- Show loading, verifying, review-in-progress, clarification-needed states

Behavior rules:
- Do not use RAG
- Do not use a database
- Do not add authentication
- Do not stream raw agent JSON
- If verdict is revised, do one bounded retry or keep the app in a review state

Implementation notes:
- Use TypeScript
- Keep the code beginner-friendly
- Render the final approved explanation with a Markdown view
```
