---
description: Build a minimal terminal-runnable math explanation agent for plain text problems.
---

# Phase 1 Terminal Single Agent

1. Read `README.md`, `AGENTS.md`, `wiki/README.md`, `wiki/index.md`, and `.agent/skills/math-agent-hands-on-starter/SKILL.md`.
2. Default to `adk-python` unless the user explicitly requests `adk-go`.
3. Build only the minimum Phase 1 prototype:
   - terminal runnable
   - plain text problem input from stdin or a CLI argument
   - outputs for problem interpretation, final answer, and step-by-step explanation
4. Do not add `RAG`, database, authentication, image input, or UI.
5. Use sample text problems from `.agent/skills/math-agent-hands-on-starter/templates/sample-problems.md` to smoke-test the CLI flow.
6. When finished, report:
   - entry point file
   - run command
   - sample problems tested
   - what should change before Phase 2
