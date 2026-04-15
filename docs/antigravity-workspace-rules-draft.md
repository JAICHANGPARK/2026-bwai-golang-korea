# Antigravity Workspace Rules Draft

이 문서는 `Antigravity`의 workspace-level `Rules`에 붙여넣기 좋은 초안을 정리한다.

목적은 두 가지다.

- 이 저장소를 `raw sources + persistent wiki` 패턴으로 일관되게 읽게 하기
- 수학 해설 프로토타입 구현 범위를 항상 같은 바닥선 위에 고정하기

## 1. 사용 방식

권장 방식:

1. `Antigravity`의 workspace rules 설정 화면을 연다.
2. 아래 초안 중 필요한 rule block을 복사해 붙여넣는다.
3. 이 저장소처럼 역할이 분명한 repo라면 `Always On`으로 두는 편이 좋다.

## 2. Rule Draft A: LLM Wiki Workspace Conventions

```text
This workspace uses a raw-sources plus persistent-wiki pattern.

Repository conventions:
- Treat docs/ as raw sources.
- Do not edit raw source documents unless explicitly requested.
- Treat wiki/ as the maintained derived knowledge layer.
- When adding knowledge, update wiki/index.md and wiki/log.md together.
- For questions, read wiki/index.md first, then the minimum relevant wiki pages.
- Only go back to docs/ source documents when the wiki is insufficient or needs verification.
- If a useful analysis emerges from a query, prefer filing it back into wiki/queries/ or an existing wiki page.
```

## 3. Rule Draft B: Math Agent Prototype Scope

```text
This workspace is also used to build a math explanation agent prototype.

Default implementation constraints:
- Default runtime is adk-python unless explicitly changed.
- Build order is:
  1. terminal single agent
  2. terminal subagent orchestration
  3. backend plus UI
- Do not skip phases.
- Do not add RAG, database, authentication, persistence, or personalization by default.
- Do not finalize the user-facing explanation unless the review status is approved.
- Prefer minimum runnable behavior first, then polish.
- Keep the frontend thin and reuse backend orchestration logic.
```

## 4. Combined Draft

workspace를 단일 규칙으로 관리하고 싶다면 아래처럼 합쳐도 된다.

```text
This workspace uses a raw-sources plus persistent-wiki pattern and is also used to build a math explanation agent prototype.

Repository conventions:
- Treat docs/ as raw sources.
- Do not edit raw source documents unless explicitly requested.
- Treat wiki/ as the maintained derived knowledge layer.
- When adding knowledge, update wiki/index.md and wiki/log.md together.
- For questions, read wiki/index.md first, then the minimum relevant wiki pages.
- Only go back to docs/ source documents when the wiki is insufficient or needs verification.

Prototype constraints:
- Default runtime is adk-python unless explicitly changed.
- Build order is terminal single agent, terminal subagent orchestration, then backend plus UI.
- Do not skip phases.
- Do not add RAG, database, authentication, persistence, or personalization by default.
- Do not finalize the user-facing explanation unless the review status is approved.
- Prefer minimum runnable behavior first, then polish.
```

## 5. Rule Design Notes

workspace rule에는 아래만 넣는 것이 좋다.

- 항상 참이어야 하는 제약
- every task에 공통으로 적용돼야 하는 규칙
- 현재 repo의 경계와 기본 해석

workspace rule에 아래는 넣지 않는 편이 좋다.

- 긴 단계별 절차
- phase별 상세 구현 지시
- 현재 작업에서만 필요한 일회성 프롬프트

그런 내용은 `Workflows`로 분리하는 편이 낫다.
