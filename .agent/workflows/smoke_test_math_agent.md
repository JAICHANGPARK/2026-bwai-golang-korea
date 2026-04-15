---
description: Run a lightweight smoke test against the current math agent phase using sample text and image problems.
---

# Smoke Test Math Agent

1. Inspect the current project state and identify whether it is in Phase 1, Phase 2, or Phase 3.
2. Use sample text problems from `.agent/skills/math-agent-hands-on-starter/templates/sample-problems.md`.
3. If the project already supports image input, also use one sample image under `docs/hands-on-assets/sample-problems/`.
4. Verify at least:
   - one normal success path
   - one slightly tricky path
   - one fallback or failure path
5. Check whether the current implementation respects the review gate.
6. Report:
   - detected phase
   - commands or steps used
   - pass/fail results
   - the biggest remaining gap
