# Starter Kits

이 디렉터리는 공유 workspace에서 바로 사용할 수 있는 runnable starter scaffold를 둔다.

## 원칙

- 설명 문서와 실행 코드를 분리한다.
- `docs/hands-on-assets/`는 공개용 자산 패키지로 유지한다.
- 실제 프로젝트 뼈대는 `starter-kits/`에 둔다.

## 런타임 선택

- `adk-python/`
  - 기본 런타임
  - 빠른 프로토타이핑과 핸즈온 진행에 적합하다.
- `adk-go/`
  - 대안 런타임
  - Go 기반 구조를 선호할 때 사용한다.

## 공통 진행 순서

1. `phase-1-terminal/`에서 텍스트 입력 기반 단일 에이전트를 만든다.
2. `phase-2-subagents/`에서 문제 해석, 풀이, 검토 에이전트로 분리한다.
3. `phase-3-fullstack/`에서 backend와 UI를 연결한다.

## 함께 쓰는 자산

- 샘플 이미지 문제는 `docs/hands-on-assets/sample-problems/`를 사용한다.
- 설계 워크시트와 fallback prompt는 `docs/hands-on-assets/templates/`를 사용한다.
- 도구별 사용 흐름은 `docs/antigravity-math-agent-step-by-step.md`, `docs/gemini-cli-math-agent-step-by-step.md`를 본다.
