# LLM Wiki 활용 가이드

이 문서는 이 저장소를 `git clone`한 뒤, `wiki/`를 실제로 `Gemini CLI` 또는 `Antigravity`에서 어떻게 활용하면 좋은지 안내한다.

처음 보면 `docs/`와 `wiki/`가 어떻게 다른지 조금 헷갈릴 수 있는데, 핵심 전제는 아주 단순하다.

- `docs/`는 사람이 만든 원천 문서다.
- `wiki/`는 LLM이 유지하는 파생 지식 계층이다.
- 에이전트는 답변 전에 `wiki/index.md`를 먼저 보고, 필요할 때만 `docs/` 원문으로 내려간다.

이 가이드는 `이미 있는 LLM Wiki를 로컬 에이전트로 탐색하고, 질문하고, 필요할 때 업데이트하는 흐름`을 다룬다.

## 1. 준비물

먼저 아래 정도만 있으면 시작할 수 있다.

필수:

- Git
- 이 저장소를 열 수 있는 개발 환경
- 아래 둘 중 하나
  - `Gemini CLI`
  - `Antigravity`

권장:

- Obsidian
  - `wiki/`를 탐색하고 graph view를 보기 좋다.
- `rg`
  - 빠른 텍스트 검색에 유용하다.

## 2. 먼저 clone 하기

실제 저장소 URL에 맞게 아래 명령만 바꿔서 사용하면 된다.

```bash
git clone https://github.com/<owner>/2026-bwai-golang-korea.git
cd 2026-bwai-golang-korea
```

clone한 직후에는 아래 세 개만 먼저 봐도 감이 꽤 빨리 잡힌다.

1. `README.md`
2. `AGENTS.md`
3. `wiki/index.md`

이 세 파일만 읽어도 이 저장소가 `raw sources + persistent wiki` 구조라는 점을 금방 이해할 수 있다.

## 3. 저장소를 어떻게 읽어야 하나

처음부터 모든 문서를 뒤질 필요는 없다. 아래 규칙만 잡아두면 훨씬 편하다.

- `docs/`를 바로 다 뒤지지 않는다.
- 질문할 때는 먼저 `wiki/index.md`에서 관련 페이지 후보를 좁힌다.
- 관련 위키 페이지를 읽고도 근거가 부족할 때만 `wiki/source-notes/`나 `docs/` 원문으로 내려간다.
- 새로운 분석 가치가 생기면 `wiki/queries/` 또는 기존 페이지에 다시 편입한다.
- 위키를 수정했다면 `wiki/index.md`와 `wiki/log.md`를 같이 갱신한다.

즉, `wiki/`는 "검색 결과를 임시로 보여주는 폴더"가 아니라 "이미 쌓아 둔 지식층"처럼 읽는 편이 좋다.

## 4. Gemini CLI에서 쓰는 방법

### 4.1 설치와 인증

`Gemini CLI`는 아래 셋 중 하나로 금방 시작할 수 있다.

```bash
npx @google/gemini-cli
npm install -g @google/gemini-cli
brew install gemini-cli
```

인증도 보통은 아래 둘 중 하나면 충분하다.

```bash
gemini
```

- 실행 후 Google 로그인 선택

또는

```bash
export GEMINI_API_KEY="YOUR_API_KEY"
gemini
```

이 저장소에는 `.gemini/settings.json`이 들어 있어서, 프로젝트 루트에서 `gemini`를 실행하면 `AGENTS.md`를 project context file로 읽도록 맞춰져 있다.

### 4.2 시작 프롬프트

처음에는 아래 프롬프트를 거의 그대로 써도 충분하다.

```text
This repo is an LLM wiki workspace.
Read README.md, AGENTS.md, wiki/README.md, and wiki/index.md first.
Then summarize:
1. what docs/ is for,
2. what wiki/ is for,
3. the difference between the hands-on profile and the product expansion profile,
4. how you will answer future questions in this repo.
```

### 4.3 Gemini CLI에서 바로 해볼 수 있는 작업

처음에는 아래 세 가지 정도만 해봐도 workspace 감이 꽤 잘 온다.

질의:

```text
Answer from the wiki first.
Read wiki/index.md, then the minimum relevant wiki pages.
Compare the hands-on profile and the product expansion profile.
Only go to docs/ if the wiki is insufficient.
```

ingest:

```text
Read docs/math-agent-system-spec.md and reflect it into the wiki following AGENTS.md.
Update the source note, the relevant component/profile/synthesis pages, wiki/index.md, and wiki/log.md.
Do not edit unrelated docs/.
```

lint:

```text
Run the wiki lint workflow for this repo.
Start with scripts/wiki-lint.sh, then inspect wiki/index.md and wiki/log.md.
If needed, create a report under wiki/lint/ and append a log entry.
```

## 5. Antigravity에서 쓰는 방법

### 5.1 workspace 열기

1. 저장소를 clone 한다.
2. Antigravity에서 이 저장소 루트를 workspace로 연다.
3. 채팅 또는 agent task에서 `wiki` 관련 작업을 요청한다.

이 저장소에는 workspace skill인 `.agent/skills/llm-wiki-workspace/SKILL.md`가 포함되어 있다.  
이 스킬의 역할은 Antigravity가 이 저장소를 일반 코드 저장소가 아니라 `LLM Wiki workspace`처럼 읽게 만드는 것이다.

### 5.2 Antigravity에서 첫 요청 예시

첫 요청은 아래처럼 가볍게 시작하면 된다.

```text
This workspace is a raw-sources plus persistent-wiki repo.
Use the llm-wiki-workspace skill if relevant.
Read README.md, AGENTS.md, wiki/README.md, and wiki/index.md first.
Then give me a short map of the workspace and recommend 3 useful next tasks.
```

### 5.3 Antigravity에서 바로 해볼 수 있는 작업

Antigravity에서도 마찬가지로 `query`, `ingest`, `lint` 정도부터 시작하면 충분하다.

질의:

```text
Use the wiki-first query workflow for this workspace.
Explain how this repository separates hands-on materials from product-expansion materials.
Show the answer with page references.
```

ingest:

```text
Use the llm-wiki-workspace skill.
Ingest docs/llm-wiki-usage-guide.md into the wiki.
Update the appropriate source note, any relevant overview/profile/component pages, wiki/index.md, and wiki/log.md.
```

lint:

```text
Health-check the wiki.
Look for missing index entries, stale claims, weak cross-references, and blurred profile boundaries.
If you find issues, write a markdown report under wiki/lint/ and log it.
```

## 5.5 단계별 문서

도구별 step 문서를 따로 두었으니, 필요할 때 바로 이어서 보면 된다.

- Gemini CLI: [`gemini-cli-math-agent-step-by-step.md`](./gemini-cli-math-agent-step-by-step.md)
- Antigravity: [`antigravity-math-agent-step-by-step.md`](./antigravity-math-agent-step-by-step.md)

Antigravity의 workspace-level automation을 다루는 문서는 아래 두 개다.

- Rules draft: [`antigravity-workspace-rules-draft.md`](./antigravity-workspace-rules-draft.md)
- Workflow usage: [`antigravity-workflow-usage-guide.md`](./antigravity-workflow-usage-guide.md)

## 6. 권장 사용 흐름

처음부터 전부 다 해볼 필요는 없다. 아래 트랙 중 하나만 골라 시작해도 된다.

### 트랙 A. 읽기 전용 시작

가볍게 둘러보는 용도로는 이 트랙이 제일 무난하다.

1. `wiki/index.md`를 읽는다.
2. `hands-on profile`과 `product expansion profile` 차이를 질문한다.
3. `wiki/source-notes/`를 열어 위키가 원천 문서를 어떻게 요약하는지 본다.
4. Obsidian graph view로 `wiki/` 연결 구조를 본다.

### 트랙 B. query 실습

질문형으로 시작하고 싶다면 아래 질문들이 무난하다.

추천 질문:

- `핸즈온 프로파일에서는 왜 RAG를 쓰지 않는가`
- `Solver/Explainer와 Expert Verifier를 왜 분리하는가`
- `이 저장소에서 curriculum research 문서는 어떤 위키 페이지로 흘러갔는가`

### 트랙 C. ingest 실습

문서 하나를 위키에 반영하는 흐름을 보고 싶다면 이 트랙이 좋다.

1. 새 문서를 `docs/`에 넣는다.
2. 에이전트에게 ingest를 요청한다.
3. 아래 파일이 함께 바뀌는지 확인한다.
   - `wiki/source-notes/...`
   - 관련 주제 페이지
   - `wiki/index.md`
   - `wiki/log.md`

### 트랙 D. lint 실습

위키를 계속 쓸 생각이라면 lint 흐름도 한 번 보는 편이 좋다.

1. `scripts/wiki-lint.sh`를 실행한다.
2. agent에게 orphan page, stale claim, 중복 주장 여부를 점검하게 한다.
3. 결과를 `wiki/lint/`에 저장하게 한다.

## 7. 자주 실수하는 지점

처음 쓰면 아래 실수를 자주 하게 된다.

- `docs/`를 위키처럼 직접 수정하려고 한다.
- 질문할 때 `wiki/index.md`를 건너뛰고 원문부터 뒤진다.
- 새 페이지를 만들고 `wiki/index.md`를 갱신하지 않는다.
- 위키를 수정하고 `wiki/log.md`에 기록하지 않는다.
- `hands-on` 범위와 `product expansion` 범위를 섞는다.

## 8. 핵심 한 문장

> 이 저장소의 핵심은 "문서를 매번 다시 검색하는 것"이 아니라, LLM이 `wiki/`를 지속적으로 관리해서 지식을 누적시키는 데 있다.

## 9. 참고 링크

- Gemini CLI README: <https://github.com/google-gemini/gemini-cli>
- Gemini CLI configuration and context files: <https://www.geminicli.com/docs/reference/configuration>
- Antigravity overview: <https://antigravity.im/>
- Antigravity skills guide: <https://antigravity.im/blog/antigravity-skills-guide>
