---
description: Refactor the Phase 1 CLI math explainer into subagents with explicit review gating.
---

# Phase 2 Terminal Subagents

1. Read `README.md`, `AGENTS.md`, `wiki/README.md`, `wiki/index.md`, and `.agent/skills/math-agent-hands-on-starter/SKILL.md`.
2. Keep the existing terminal entry path working.
3. Refactor the current implementation into:
   - problem interpretation agent
   - problem solving agent
   - solution review agent
4. Add an orchestrator that calls the subagents in sequence.
5. Require the review agent to return an explicit status such as `approved`, `revise`, or `fail`.
6. Do not finalize the user-facing explanation unless the review status is `approved`.
7. Keep the output demo-friendly so the orchestration is visible in the terminal.
8. Use sample text problems from `.agent/skills/math-agent-hands-on-starter/templates/sample-problems.md`.
9. When finished, report:
   - orchestrator entry point
   - the three subagents and their responsibilities
   - one approved case
   - one revise or fail case
   - what should change before Phase 3
