---
description: Extend the orchestrated math explainer into a backend plus UI prototype with text and image input.
---

# Phase 3 Fullstack UI

1. Read `README.md`, `AGENTS.md`, `wiki/README.md`, `wiki/index.md`, and `.agent/skills/math-agent-hands-on-starter/SKILL.md`.
2. Reuse the Phase 2 orchestration as the backend core.
3. Build a thin UI that supports:
   - text input
   - image upload
   - submit action
   - status display
   - final answer and explanation card
   - failure card
4. Show clear states such as `analyzing`, `solving`, and `reviewing`.
5. If image quality is poor or review fails, show a clear fallback instead of pretending success.
6. Do not add `RAG`, authentication, persistence, or unrelated product features.
7. Use sample image files from `docs/hands-on-assets/sample-problems/` when useful.
8. When finished, report:
   - how to run the backend
   - how to run the frontend
   - supported inputs
   - UI states
   - one image-based path
   - one fallback path
