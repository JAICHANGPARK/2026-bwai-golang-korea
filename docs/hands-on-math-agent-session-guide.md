# 1시간 핸즈온 세션 가이드: AI Studio에서 수학 해설 시스템 설계하기

작성일: 2026-04-04  
버전: 0.1  
상태: Draft

## 1. 문서 목적

이 문서는 `Google AI Studio`, `AI Studio Build`, `Google Antigravity`를 사용해 참석자들이 약 1시간 안에 `수학 문제 해설 시스템`을 설계하고, 그 설계를 바탕으로 실제 앱 생성까지 해보는 핸즈온 세션을 설계하기 위한 운영 가이드다.

### 1.1 문서 적용 범위

이 문서는 `핸즈온 프로파일` 전용 문서다. 같은 워크스페이스에 있는 제품형 문서와 범위를 구분해서 읽어야 한다.

| 구분 | 핸즈온 프로파일 | 제품 확장 프로파일 |
| --- | --- | --- |
| 기준 문서 | 이 문서, `docs/prompt-pack/` | `math-agent-system-spec.md`, `math-agent-knowledge-plan.md`, `internal-knowledge-base-build.md` |
| 에이전트 수 | Solver + Verifier | 멀티 에이전트 확장 가능 |
| 지식 검색 | 사용하지 않음 | RAG / File Search 고려 |
| 목표 | 1시간 안에 설계와 앱 생성 | 제품 품질과 확장성 확보 |
| 최종 표시 규칙 | 검증 승인 후에만 최종 설명 표시 | 제품 정책에 따라 확장 가능 |

핸즈온 준비와 진행에서는 반드시 이 문서와 `docs/prompt-pack/`를 우선 기준으로 삼는다.

핸즈온의 조건은 다음과 같다.

- 참석자는 AI Studio 안에서 시스템 설계를 같이 해보는 경험을 해야 한다.
- 참석자는 프론트엔드 감각으로 결과를 이해할 수 있어야 한다.
- `RAG`는 사용하지 않는다.
- 메인 모델은 `gemini-3-flash-preview`를 사용한다.
- 멀티 에이전트 구조를 경험하게 한다.
- 최소한 다음 두 역할을 구현하게 한다.
  - `수학 문제 풀이·해설 에이전트`
  - `전문가 검증 에이전트`

---

## 2. 세션 목표

### 2.1 참석자가 세션 끝에 가져가야 하는 것

세션이 끝났을 때 참석자는 아래를 이해하고 있어야 한다.

- 이미지 문제를 받아 Gemini로 풀이·해설하는 기본 흐름
- 단일 에이전트가 아니라 `풀이`와 `검증`을 분리하는 이유
- AI Studio에서 시스템 요구사항을 프롬프트와 구조화 출력으로 설계하는 방식
- 설계한 명세와 프롬프트를 `AI Studio Build`와 `Antigravity`에 재사용하는 방식
- 룰, 워크플로우, 스킬을 통해 반복 가능한 에이전트 개발 흐름을 만드는 방법

### 2.2 세션 산출물

참석자는 아래 설계 산출물을 완성하는 것을 목표로 한다.

- 사용자 문제 처리 흐름
- Solver / Verifier 에이전트 역할 정의
- Solver 시스템 프롬프트 초안
- Verifier 시스템 프롬프트 초안
- 화면 설계 초안
- 구조화 출력 JSON 초안
- 실패 처리 기준
- 스트리밍 응답 및 Markdown 렌더링 요구사항
- 앱 생성용 베이스 명세 프롬프트

선택 산출물:

- 간단한 AI Studio Build 프로토타입
- Antigravity 버전 프로토타입
- 결과 UI 초안
- 스트리밍과 Markdown view가 반영된 응답 화면

---

## 2.3 세션의 핵심 운영 시나리오

이번 세션은 아래 순서로 진행하는 것이 가장 좋다.

1. `AI Studio`에서 시스템을 설계한다.
2. 설계 내용을 바탕으로 `명세서/프롬프트 패키지`를 만든다.
3. 그 명세를 `AI Studio Build`에 넣어 앱을 만든다.
4. 같은 명세를 `Antigravity`에 넣어 또 다른 방식으로 앱을 만든다.

핵심 메시지는 아래와 같다.

- 좋은 시스템 설계는 재사용 가능하다.
- 좋은 시스템 프롬프트는 앱 생성 도구를 바꿔도 살아남는다.
- 명세를 먼저 만들면 Build와 Antigravity 모두 훨씬 안정적으로 동작한다.

---

## 3. 어떤 툴을 메인으로 쓸까

### 3.1 권장 결론

1시간 세션의 메인 트랙은 `AI Studio에서 시스템 설계와 명세 생성`으로 가져가는 것을 권장한다.  
이유는 가장 중요한 학습 포인트가 `앱 구현`보다 `에이전트 구조 설계`, `시스템 프롬프트 설계`, `출력 계약 정의`, `명세 재사용`이기 때문이다.

`Antigravity`는 확장 트랙 또는 진행자 데모용으로 매우 좋다.  
특히 아래 기능을 강조하고 싶다면 Antigravity가 잘 맞는다.

- Rules
- Workflows
- Skills
- 반복 가능한 에이전트 개발 프로세스

### 3.2 권장 운영 방식

- 1단계: `AI Studio 설계 핸즈온`
- 2단계: `AI Studio Build로 앱 생성`
- 3단계: `Antigravity로 동일 명세 기반 앱 생성`

이렇게 하면 참석자는 1시간 내에 설계부터 구현까지 연결된 흐름을 경험하고, 같은 명세가 서로 다른 생성 도구에서 재사용되는 것도 볼 수 있다.

---

## 4. 공식 기능 관점에서의 판단

핸즈온 설계 시 참고할 공식 정보는 아래와 같다.

- `AI Studio Build`는 프롬프트 기반으로 웹앱 프로토타입을 생성하고, 이후 코드 수정과 annotation 방식으로 점진적으로 다듬을 수 있다.
- AI Studio Build는 기본적으로 `React` 웹앱을 생성한다.
- AI Studio Build는 npm 패키지 사용과 코드 export를 지원한다.
- `Antigravity`는 자연어 기반 앱 생성 경험 위에, Rules, Workflows, Skills 같은 반복 가능한 개발 단위를 둘 수 있다.
- Antigravity는 내부적으로 코어 에이전트와 도구 실행 구조를 갖고 있어, prompt-only 실험보다 팀형 개발에 유리하다.

핸즈온 관점에서는 아래처럼 해석하면 된다.

- 참석자는 "프론트엔드 중심"으로 느끼게 한다.
- 하지만 보안과 실행 안정성을 위해 툴이 생성하는 얇은 서버 레이어는 허용한다.
- 세션의 학습 초점은 백엔드가 아니라 `에이전트 설계`와 `프롬프트 설계`에 둔다.

### 4.1 운영상 주의

AI Studio Build 공식 문서 기준으로, 공유된 앱의 코드는 보는 사람에게 노출될 수 있으므로 실제 API 키를 코드에 직접 넣으면 안 된다.  
세션에서는 아래 원칙을 지키는 것이 좋다.

- `AI Studio 내부 실행/공유`를 기본 시연 경로로 사용한다.
- 외부 공개 배포는 세션 범위에서 제외한다.
- "오늘은 기능 구현과 에이전트 설계에 집중하고, 배포 보안은 다음 단계"라고 명확히 안내한다.

---

## 5. 핸즈온 범위

### 5.1 이번 세션에서 포함할 것

- 사용자 요청 흐름 설계
- Solver/Explainer Agent 설계
- Expert Verifier Agent 설계
- 시스템 프롬프트 설계
- 화면 설계 프롬프트 설계
- 구조화 출력 설계
- 실패 처리 설계
- 스트리밍 및 Markdown 응답 UX 설계
- 앱 생성용 명세 프롬프트 정리
- AI Studio Build 앱 생성
- Antigravity 앱 생성

### 5.2 이번 세션에서 제외할 것

- RAG
- 벡터 DB
- 사용자 계정
- 저장소/히스토리
- 교육과정 데이터베이스
- 개인화 추천
- GraphRAG

### 5.3 왜 이렇게 제한하나

1시간 안에 가장 중요한 학습 포인트는 `에이전트 사고방식`, `프롬프트 분리`, `출력 구조 설계`다.  
RAG나 데이터 인프라까지 넣으면 참가자는 시스템 개념보다 설정과 디버깅에 시간을 다 쓰게 된다.

---

## 6. 권장 최종 결과물

세션 종료 시 아래 수준의 설계 결과가 나오면 충분하다.

1. 사용자가 문제를 올리면 어떤 흐름으로 처리할지 설명할 수 있다.
2. Solver/Explainer Agent의 역할을 정의할 수 있다.
3. Expert Verifier Agent의 역할을 정의할 수 있다.
4. 두 에이전트의 시스템 프롬프트 초안을 만들 수 있다.
5. 화면 설계와 결과 카드 구성을 정의할 수 있다.
6. 두 에이전트의 JSON 출력 구조를 정의할 수 있다.
7. 스트리밍 응답과 Markdown 렌더링 요구를 명세에 반영할 수 있다.
8. 앱 생성에 사용할 통합 명세 프롬프트를 만들 수 있다.
9. 같은 명세로 AI Studio Build와 Antigravity에서 각각 앱을 생성해볼 수 있다.

---

## 7. 1시간 핸즈온 시간표

## 7.1 추천 진행안

| 시간 | 단계 | 목표 |
| --- | --- | --- |
| 0~5분 | 오프닝 | 오늘 설계할 시스템 소개 |
| 5~12분 | 사용자 흐름 정의 | 사용자가 문제를 올리면 어떤 일이 일어나는지 설계 |
| 12~22분 | 에이전트 분리 | Solver / Verifier 역할 정의 |
| 22~30분 | 시스템 프롬프트 설계 | 두 에이전트의 prompt 초안 완성 |
| 30~36분 | 화면 설계 | 업로드, 카드, 상태 영역, Markdown 결과 영역 설계 |
| 36~43분 | 출력 구조 설계 | JSON 구조와 실패 상태 정의 |
| 43~48분 | 통합 명세 프롬프트 만들기 | Build/Antigravity 공용 명세로 정리 |
| 48~54분 | AI Studio Build 생성 | 명세 프롬프트로 앱 생성 및 streaming/Markdown 보강 |
| 54~58분 | Antigravity 생성 | 같은 명세로 앱 생성 및 UX 보강 |
| 58~60분 | 회고 및 확장 아이디어 | RAG, 교육과정, 개인화는 다음 단계로 안내 |

---

## 8. 추천 세션 진행 방식

### 8.1 기본 흐름

- 진행자가 단계별 설계 프롬프트를 먼저 보여준다.
- 참석자는 AI Studio에 프롬프트를 넣고 설계를 같이 정리한다.
- 생성된 결과를 기준으로 작은 수정 프롬프트를 반복한다.
- 세션 후반에는 각자 Solver / Verifier 프롬프트를 조금 바꿔본다.
- Build 단계에서는 streaming과 Markdown view를 마지막 보강 포인트로 넣는다.

### 8.2 중요한 운영 포인트

- 첫 프롬프트는 구체적으로 준다.
- 이후 단계별로 수정 프롬프트를 쪼개서 준다.
- 한 번에 너무 많은 요구사항을 넣지 않는다.
- "왜 이 프롬프트를 넣는지"를 설명한다.
- UI 설계와 스트리밍 UX는 통합 명세 직전에 반드시 고정한다.

---

## 9. 추천 설계 구조

핸즈온의 중심은 UI보다 시스템 설계다.  
따라서 참가자에게는 먼저 아래 구조를 정리하게 하는 것이 좋다.

### 9.1 시스템 구성

- 사용자 입력
- Solver/Explainer Agent
- Expert Verifier Agent
- 최종 승인 응답

### 9.2 선택적 프로토타입 UI

원하면 단일 페이지 UI로 아래 정도만 구현한다.

- 문제 이미지 업로드 영역
- 사용자 질문 입력창
- Solver 결과 카드
- Verifier 결과 카드
- 최종 승인 해설 카드

### 9.3 내부 로직

- `handleSolve()`
  - 이미지와 질문을 Solver에게 보냄
- `handleVerify()`
  - Solver 결과를 Verifier에게 보냄
- `handleRevisionRetry()`
  - verdict가 `revised`면 검토 코멘트를 반영해 최대 1회만 재시도
- `prepareFinalExplanationInput()`
  - verifier가 승인한 구조화 결과만 최종 설명 입력으로 변환
- `streamFinalExplanation()`
  - 승인된 구조화 결과를 바탕으로 최종 Markdown 설명만 스트리밍
- `renderFinalResult()`
  - `approved`일 때만 최종 설명을 표시

### 9.4 데이터 구조

세션에서는 아래 정도의 간단한 상태 구조면 충분하다.

```ts
type SolverResult = {
  problemSummary: string;
  finalAnswer: string;
  steps: { title: string; detail: string }[];
  keyConcepts: string[];
  confidence: number;
};

type VerifiedSolution = {
  finalAnswer: string;
  steps: { title: string; detail: string }[];
  keyConcepts: string[];
};

type VerificationResult = {
  verdict: "approved" | "revised" | "needs_clarification";
  comments: string[];
  approvedSolution?: VerifiedSolution;
  revisionRequests?: string[];
  nextAction:
    | "stream_final_explanation"
    | "retry_solver_once"
    | "ask_for_clearer_image";
};

type FinalExplanationInput = {
  problemSummary: string;
  finalAnswer: string;
  steps: { title: string; detail: string }[];
  keyConcepts: string[];
  verifierComments: string[];
};
```

### 9.5 최종 설명 생성 계약

- Solver와 Verifier는 `JSON only`로 동작한다.
- 앱은 Solver 또는 Verifier의 raw JSON을 사용자용 최종 설명으로 직접 스트리밍하지 않는다.
- `approved`일 때만 `FinalExplanationInput`을 만들고, 그 입력으로 최종 설명 Markdown을 생성해 스트리밍한다.
- `revised`는 최종 표시 상태가 아니라 `재검토 중` 상태다.
- `revised`이면 검토 코멘트를 바탕으로 최대 1회 재시도하고, 재승인되기 전에는 최종 설명 카드를 보여주지 않는다.
- `needs_clarification`이면 더 선명한 이미지나 추가 입력을 요청한다.

---

## 10. AI Studio 메인 트랙: 시스템 설계

### 10.1 목표

참석자가 AI Studio 안에서 수학 해설 시스템을 직접 설계하고, 핵심 prompt와 출력 구조를 완성한다.

### 10.2 핵심 산출물

- 사용자 흐름 초안
- Solver / Verifier 역할 정의
- 시스템 프롬프트 초안
- JSON 출력 구조
- 실패 처리 규칙
- 최종 설명 생성 계약
- 앱 생성용 통합 명세 프롬프트

### 10.3 단계별 설계 프롬프트

#### 1단계: 사용자 흐름 정의

```text
We are designing a math explanation system for middle and high school students.

The user uploads a math problem image and wants:
- the correct answer
- a step-by-step explanation
- key concepts
- a verified final response

Please design the end-to-end user flow in Korean.
Keep it simple and suitable for a 1-hour workshop.
Assume we are not using RAG.
```

#### 2단계: 에이전트 분리

```text
We want to design this as a multi-agent system, not a single agent.

Please define two agents:
1. a Solver/Explainer Agent
2. an Expert Verifier Agent

For each agent, describe:
- role
- responsibilities
- what it should not do
- input
- output

Write the result in Korean and keep it practical for a workshop.
```

#### 3단계: Solver 시스템 프롬프트 초안

```text
Please write a high-quality system prompt for a Solver/Explainer Agent for Korean middle and high school math problems.

The agent should:
- read a math problem image carefully
- solve it step by step
- explain it for students
- produce structured output

Important rules:
- do not guess unreadable text
- do not skip logical steps
- keep explanations clear and educational
- do not act like the final verifier

Write the system prompt in English.
```

#### 4단계: Verifier 시스템 프롬프트 초안

```text
Please write a high-quality system prompt for an Expert Verifier Agent for Korean math problem solving.

The verifier should:
- inspect the solver output
- check correctness
- detect missing conditions
- detect logical inconsistencies
- decide whether the result is approved, revised, or needs clarification

Important rules:
- be strict
- do not approve just because the answer sounds fluent
- focus on correctness before style

Write the system prompt in English.
```

#### 5단계: 구조화 출력 설계

```text
Please design a JSON output schema for:
1. Solver/Explainer Agent
2. Expert Verifier Agent

The schema should be simple enough for a workshop demo and should include:
- answer
- steps
- key concepts
- confidence
- verifier verdict
- comments

Return the schema examples in JSON format.
```

#### 6단계: 실패 처리 설계

```text
Please list the main failure cases for this math explanation system.

Examples:
- unreadable image
- incomplete problem
- uncertain answer
- inconsistent solution

For each failure case, suggest:
- what the system should detect
- what the system should say to the user
- whether the verifier should reject or ask for clarification

Write the result in Korean.
```

#### 7단계: 샘플 테스트 설계

```text
We now have a draft design for the system.

Please help us evaluate it with a sample math problem image.
For the sample, show:
- what the Solver should output
- what the Verifier should check
- what the final approved response should look like

Keep the response structured and easy to review in a workshop.
```

#### 8단계: 앱 생성용 통합 명세 프롬프트 만들기

```text
Based on the design above, create a single implementation specification prompt that can be reused in app builders.

The prompt should include:
- product goal
- target users
- user flow
- two-agent architecture
- Solver role
- Verifier role
- structured JSON outputs
- simple Korean UI requirements
- constraints: no RAG, no database, no auth

Write the result in English so it can be pasted into AI Studio Build or Antigravity.
Keep it concise but implementation-ready.
```

---

## 11. 선택 트랙: AI Studio Build로 간단한 프로토타입 만들기

### 11.1 목표

참석자가 방금 만든 `통합 명세 프롬프트`를 바탕으로 최소 수정으로 `실행되는 수학 해설 앱`을 만든다.

### 11.2 메인 아이디어

AI Studio Build에서는 "처음부터 즉흥적으로 만들기"보다 "설계된 명세를 붙여 넣고 생성하기"를 경험하게 하는 것이 핵심이다.

권장 흐름:

1. 통합 명세 프롬프트 붙여넣기
2. 앱 뼈대 생성
3. Solver / Verifier 연결 확인
4. UI 다듬기

### 11.3 왜 AI Studio Build가 메인 트랙에 적합한가

- 기본적으로 React 웹앱을 빠르게 생성한다.
- 코드 탭에서 바로 수정할 수 있다.
- annotation 방식으로 UI 수정을 쉽게 시연할 수 있다.
- 세션 후 ZIP으로 내보내거나 GitHub로 보낼 수 있다.
- 같은 명세 프롬프트를 바로 앱으로 연결해보기에 좋다.

---

## 12. AI Studio Build 단계별 프롬프트

## 12.1 1단계: 베이스 앱 생성 프롬프트

진행자가 제공할 프롬프트:

```text
Build a Korean-language web app called "Math Explainer Lab".

The app is a simple single-page app for students.
Users can upload one math problem image and optionally type a short question.

The app should have:
- a clean and modern Korean UI
- an image upload area
- a text input for the user's question
- a "문제 풀이 시작" button
- a section for solver output
- a section for verifier output
- a final approved explanation section that appears only after verification approval

Keep the architecture simple and easy for a 1-hour workshop.
Do not add authentication, database, history, RAG, or external search.
Use TypeScript and keep the code easy to read for frontend learners.
```

### 진행자 설명 포인트

- 첫 단계에서는 기능보다 UI 골격이 중요하다.
- 참가자가 "앱이 만들어졌다"는 느낌을 빨리 받아야 한다.

---

## 12.2 2단계: Solver/Explainer Agent 추가 프롬프트

```text
Add a Solver/Explainer agent powered by gemini-3-flash-preview.

When the user uploads a math problem image and clicks the solve button, the app should:
1. send the image and question to the Solver agent
2. ask the agent to read the problem carefully
3. generate:
   - a short problem summary
   - the final answer
   - step-by-step solution
   - key concepts
   - a confidence score

The Solver agent must return structured JSON.

Use this output schema:
{
  "problemSummary": "string",
  "finalAnswer": "string",
  "steps": [{"title": "string", "detail": "string"}],
  "keyConcepts": ["string"],
  "confidence": 0.0
}

Render the result in the solver result section.
Keep the code simple and workshop-friendly.
```

### 진행자 설명 포인트

- "문제를 잘 푸는 프롬프트"보다 "구조화된 출력"이 더 중요하다.
- 구조화 출력이 있어야 뒤 단계에서 검증 에이전트를 붙이기 쉽다.

---

## 12.3 3단계: Solver 시스템 프롬프트 제공

진행자가 참가자에게 붙여주면 좋은 시스템 프롬프트:

```text
You are a math problem solving and explanation agent for middle school and high school students.

Your job is to:
- carefully read the uploaded math problem image
- understand what the problem is asking
- solve the problem step by step
- produce a student-friendly explanation
- return a structured JSON response

Rules:
- do not guess unreadable text or numbers
- if the image is unclear, say so in the summary
- do not skip logical steps
- prefer correct and simple explanations over fancy wording
- keep the answer mathematically consistent

Output only the JSON object requested by the app.
```

### 진행자 설명 포인트

- 초보자에게는 system prompt를 처음부터 직접 쓰게 하지 말고 제공하는 편이 좋다.
- 이후 "tone"이나 "detail level"만 직접 바꾸게 한다.

---

## 12.4 4단계: Verifier Agent 추가 프롬프트

```text
Add a second agent called Expert Verifier using gemini-3-flash-preview.

The Verifier should receive:
- the original user question
- the uploaded image
- the Solver JSON result

The Verifier must check:
- whether the final answer is likely correct
- whether the solution steps are logically consistent
- whether any conditions were missed

The Verifier should return structured JSON:
{
  "verdict": "approved" | "revised" | "needs_clarification",
  "comments": ["string"],
  "approvedSolution": {
    "finalAnswer": "string",
    "steps": [{"title": "string", "detail": "string"}],
    "keyConcepts": ["string"]
  },
  "revisionRequests": ["string"],
  "nextAction": "stream_final_explanation" | "retry_solver_once" | "ask_for_clearer_image"
}

If approved, return the validated solution in approvedSolution.
If revised, do not treat corrected content as final user output. Return revisionRequests and set nextAction to retry_solver_once.
Render the verifier result in a separate UI section.
```

### 진행자 설명 포인트

- 여기서 멀티 에이전트의 핵심 가치를 보여준다.
- "한 번 더 생각하는 시스템"을 만들었다는 감각을 주는 것이 중요하다.

---

## 12.5 5단계: Verifier 시스템 프롬프트 제공

```text
You are an expert math verification agent.

Your job is not to explain first.
Your job is to verify whether the Solver's result is correct and safe to show to students.

Check:
- answer correctness
- logical consistency
- missing conditions
- risky or misleading explanation

Rules:
- be strict
- do not approve just because the explanation sounds fluent
- if the image is unclear, choose needs_clarification
- if the answer is mostly right but steps need correction, choose revised

Return only the JSON object requested by the app.
```

---

## 12.6 6단계: 결과 화면 정리 프롬프트

```text
Improve the app UX for the workshop demo.

Please:
- show the solver result as a card
- show the verifier result as a separate card
- if verdict is approved, show a final approved explanation section
- stream the final approved explanation and render it as Markdown
- if verdict is revised, show verifier comments and a short "재검토 중" state instead of a final answer
- if verdict is needs_clarification, show a friendly Korean message asking for a clearer image

Keep the UI visually clean and beginner-friendly.
Use simple Korean labels.
```

---

## 12.7 7단계: 마무리 프롬프트

```text
Polish the app for a live demo.

Please improve:
- loading states
- error states
- empty states
- mobile layout

Do not add new major features.
Focus on making the current workshop app stable and clear.
```

---

## 13. Antigravity 확장 트랙

Antigravity 트랙은 `룰`, `워크플로우`, `스킬`을 보여주고 싶을 때 좋다.  
다만 1시간 세션 전체를 Antigravity만으로 끌고 가기보다, 아래 둘 중 하나를 권장한다.

- 진행자 데모 트랙
- 빠른 참가자를 위한 확장 트랙

핵심은 `AI Studio에서 만든 통합 명세 프롬프트`를 Antigravity에도 다시 넣어보는 것이다.  
이 경험을 통해 참석자는 "좋은 설계는 빌더가 달라도 재사용된다"는 점을 체감하게 된다.

---

## 14. Antigravity에서 보여주면 좋은 포인트

### 13.1 Rules

팀이 반복해서 지키고 싶은 개발 원칙을 강제한다.

예:

- RAG 금지
- 프론트엔드 중심 유지
- TypeScript 우선
- JSON 출력 강제
- 두 에이전트 구조 고정

### 13.2 Workflows

자주 쓰는 프롬프트 시퀀스를 slash-command처럼 재사용하게 한다.

예:

- `/scaffold-math-app`
- `/add-solver-agent`
- `/add-verifier-agent`
- `/polish-demo-ui`

### 13.3 Skills

특정 작업을 안정적으로 수행하는 재사용 지식 단위로 보여줄 수 있다.

예:

- `math-agent-ui-skill`
- `solver-prompt-skill`
- `verifier-prompt-skill`

---

## 15. Antigravity Rules 예시

진행자가 미리 준비해둘 수 있는 규칙 예시:

```text
Project Rules

1. This project is a 1-hour workshop app.
2. Keep the architecture simple and frontend-focused.
3. Do not add RAG, vector databases, authentication, or history.
4. Use gemini-3-flash-preview for both agents.
5. Keep two agents only:
   - Solver/Explainer
   - Expert Verifier
6. All agent outputs must be structured JSON.
7. Prefer readable TypeScript and simple React components.
8. Do not over-engineer file structure.
9. UI text should be in Korean.
10. Optimize for workshop clarity over production completeness.
```

---

## 16. Antigravity Workflows 예시

## 16.1 `/scaffold-math-app`

```text
Create a simple Korean web app for a 1-hour workshop.
The app should allow a student to upload one math problem image and ask a short question.
Add sections for:
- solver result
- verifier result
- final approved explanation

Keep the app simple, clean, and easy to understand.
Do not add RAG or database features.
```

## 16.2 `/add-solver-agent`

```text
Add a Solver/Explainer agent using gemini-3-flash-preview.
The agent must read the uploaded image, solve the problem, and return structured JSON with:
- problemSummary
- finalAnswer
- steps
- keyConcepts
- confidence
```

## 16.3 `/add-verifier-agent`

```text
Add an Expert Verifier agent using gemini-3-flash-preview.
The verifier must inspect the solver result and return structured JSON with:
- verdict
- comments
- approvedSolution
- revisionRequests
- nextAction
```

## 16.4 `/polish-demo-ui`

```text
Improve the app for a live workshop demo.
Add better card layout, empty states, loading states, streaming states, and clear Korean labels.
Show the final explanation card only after verifier approval.
Do not add new major logic.
```

---

## 17. Antigravity Skills 예시

### 17.1 `solver-prompt-skill`

설명:

- Solver 시스템 프롬프트를 안정적으로 재사용
- structured output 패턴을 유지

핵심 내용:

```text
When building a solver agent for this project:
- read the image carefully
- do not hallucinate unreadable numbers
- solve step by step
- explain for students
- return JSON only
```

### 17.2 `verifier-prompt-skill`

설명:

- 검증 에이전트가 설명자가 아니라 검토자 역할을 유지

핵심 내용:

```text
When building a verifier agent for this project:
- check correctness before style
- identify missing conditions
- reject or revise when uncertain
- return JSON only
```

### 17.3 `workshop-ui-skill`

설명:

- 카드 기반 UI를 일관되게 생성

핵심 내용:

```text
For this workshop app:
- keep UI as a single page
- show solver, verifier, and final answer as separate cards
- use clear Korean labels
- keep code beginner-friendly
```

---

## 18. 단계별 진행자 멘트 가이드

### 18.1 오프닝

- "오늘은 수학 문제 해설 앱을 완성하는 시간이기보다, 그 앱의 핵심 시스템을 AI Studio에서 같이 설계해보는 시간입니다."

### 18.2 Solver 단계

- "첫 번째 에이전트는 문제를 풉니다. 중요한 건 정답보다 구조화된 출력을 만드는 겁니다."

### 18.3 Verifier 단계

- "두 번째 에이전트는 친절한 설명자가 아니라 엄격한 검토자입니다. 이 분리가 오늘의 핵심입니다."

### 18.4 마무리

- "오늘은 RAG 없이도 멀티 에이전트 경험을 만들 수 있다는 걸 본 거고, 다음 단계에서 교육과정 연결과 지식 검색을 넣으면 제품이 더 강해집니다."

---

## 19. 참가자에게 제공하면 좋은 패키지

핸즈온 전에 아래를 미리 제공하면 진행이 빨라진다.

- 샘플 수학 문제 이미지 2~3장
- 설계 워크시트
- Solver 시스템 프롬프트
- Verifier 시스템 프롬프트
- JSON 스키마 예시
- 통합 명세 프롬프트 템플릿
- 단계별 프롬프트 문서
- 실패했을 때 넣을 fallback 프롬프트

권장 위치:

- 자산 안내 문서: [hands-on-assets/README.md](./hands-on-assets/README.md)
- 샘플 문제 이미지 1: [problem-01-linear-equation.svg](./hands-on-assets/sample-problems/problem-01-linear-equation.svg)
- 샘플 문제 이미지 2: [problem-02-function-value.svg](./hands-on-assets/sample-problems/problem-02-function-value.svg)
- 샘플 문제 이미지 3: [problem-03-quadratic-choice.svg](./hands-on-assets/sample-problems/problem-03-quadratic-choice.svg)
- 설계 워크시트: [design-worksheet.md](./hands-on-assets/templates/design-worksheet.md)
- 통합 명세 템플릿: [integration-spec-template.md](./hands-on-assets/templates/integration-spec-template.md)
- fallback prompt 템플릿: [fallback-prompts.md](./hands-on-assets/templates/fallback-prompts.md)

---

## 20. 실패 대비 프롬프트

생성 결과가 너무 복잡해지거나 이상해질 때 진행자가 바로 주면 좋은 프롬프트:

### 20.1 복잡성 축소 프롬프트

```text
Simplify the app.
This is a 1-hour workshop demo.
Keep only:
- image upload
- solve button
- solver result
- verifier result
- final approved result

Remove any extra complexity.
```

### 20.2 JSON 강제 프롬프트

```text
Make both agents return strict JSON only.
Do not include markdown or prose outside the JSON object.
```

### 20.3 UI 축소 프롬프트

```text
Reduce the UI to a single clean page with three result cards.
Keep it simple for beginners.
```

---

## 21. 세션 후 확장 아이디어

핸즈온이 끝난 뒤 다음 단계로 연결하면 좋다.

- 교육과정/성취기준 연결
- 내부 지식베이스 연결
- File Search
- 사용자 학년별 설명
- 풀이 히스토리
- 복습 문제 추천

이때 "오늘 만든 앱은 멀티 에이전트의 최소형이고, 앞으로 지식 검색과 학습 데이터가 붙으면 제품형으로 진화한다"는 메시지를 주는 것이 좋다.

---

## 22. 최종 권장안

1시간 핸즈온의 최적 전략은 아래와 같다.

1. `AI Studio에서 시스템 설계`를 메인 트랙으로 사용한다.
2. Solver/Verifier 두 에이전트 구조만 설계한다.
3. 그 결과를 `통합 명세 프롬프트`로 정리한다.
4. RAG와 데이터 인프라는 제외한다.
5. 단계별 설계 프롬프트를 진행자가 제공한다.
6. 같은 명세를 `AI Studio Build`와 `Antigravity`에 각각 넣어 앱을 만들어본다.
7. `Antigravity`에서는 Rules/Workflows/Skills 개념까지 연결해본다.

핵심은 참가자가 세션 끝에 아래를 말할 수 있게 만드는 것이다.

- "나는 문제를 푸는 에이전트와 검증하는 에이전트를 분리해서 설계해봤다."
- "나는 AI Studio에서 시스템 프롬프트와 출력 구조를 직접 설계해봤다."
- "나는 같은 명세를 Build와 Antigravity에 재사용해봤다."
- "나는 다음 단계로 RAG 연결까지 확장할 수 있겠다."

---

## 23. 참고 링크

- Build mode in Google AI Studio  
  https://ai.google.dev/gemini-api/docs/aistudio-build-mode

- Antigravity 소개  
  https://developers.googleblog.com/en/introducing-antigravity-build-and-deploy-full-stack-ai-apps-in-your-browser/

- Google AI Studio  
  https://ai.google.dev/aistudio

- Gemini API models  
  https://ai.google.dev/gemini-api/docs/models
