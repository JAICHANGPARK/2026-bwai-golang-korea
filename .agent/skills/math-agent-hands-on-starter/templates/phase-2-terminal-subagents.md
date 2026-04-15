# Phase 2 Prompt

```text
Use the math-agent-hands-on-starter skill.

Build Phase 2 only.
Take the existing terminal math explainer and refactor it into subagents.

Required subagents:
- problem interpretation agent
- problem solving agent
- solution review agent

Requirements:
- Keep the app runnable from the terminal.
- Add an orchestrator that calls the subagents in sequence.
- The review agent must return an explicit status such as approved, revise, or fail.
- Do not finalize the user-facing explanation unless the review status is approved.
- Show enough terminal output to make the orchestration visible during the demo.
- Keep the code minimal and demo-friendly.

Stop when the orchestration demo works for at least 3 text problems.
```
