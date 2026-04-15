# Shared Workspace Packaging Guide

이 문서는 수학 해설 에이전트 workspace를 다른 사람과 공유할 때, 어떤 자료를 repo 안에 두면 좋고 어떤 자료는 따로 빼두는 편이 좋은지 정리한다.

공유용 repo를 만들다 보면 문서, 샘플 자산, 실행 코드, 운영 메모가 금방 한 폴더에 섞이기 쉽다. 이 문서는 그 경계를 조금 더 편하게 잡기 위한 기준이라고 보면 된다.

## 목표

- 공유받은 사람이 바로 탐색하고 실행할 수 있는 구조를 유지한다.
- `docs/`에는 설명 문서와 공개 가능한 자산만 둔다.
- runnable code와 운영 메모를 같은 폴더에 섞지 않는다.

## 공유 workspace에 남길 것

공유용 repo에 남기면 좋은 건 아래와 같다.

- `docs/prompt-pack/` 같은 단계별 설계 문서
- `docs/hands-on-assets/` 같은 공개용 샘플 문제, 템플릿, 한 장 요약
- `docs/antigravity-*.md`, `docs/gemini-cli-*.md` 같은 도구별 사용 가이드
- `.agent/skills/`, `.agent/workflows/` 같은 workspace 자동화 자산
- `starter-kits/` 같은 runnable starter scaffold

## 공유 workspace에서 뺄 것

반대로 아래 항목은 shared workspace 밖에서 관리하는 편이 안전하다.

- facilitator용 시간표, 오프닝 멘트, 진행 스크립트
- 내부 운영 체크리스트와 트러블슈팅 메모
- 행사용 명단, 링크 모음, 비공개 안내 문서
- 토큰, API 키, 환경 변수 예시처럼 오해를 부를 수 있는 민감 설정

## 권장 배치

폴더를 나눌 때는 아래 정도로 생각하면 깔끔하다.

- `docs/hands-on-assets/`
  - 공개 가능한 handout, sample input, worksheet만 둔다.
- `starter-kits/`
  - `adk-python`, `adk-go` 같은 런타임별 starter code 뼈대를 둔다.
- `scripts/`
  - repo 운영용 스크립트만 둔다.
- facilitator 전용 자료
  - 별도 private repo, 노션, 로컬 폴더처럼 공유 workspace 밖에서 관리한다.

## 공개 전 체크리스트

공개 직전에 아래 항목만 한 번 훑어봐도 사고를 많이 줄일 수 있다.

1. 역할 중심 내부 라벨이 파일명과 문구에 과도하게 드러나지 않는지 다시 본다.
2. `docs/hands-on-assets/` 안에 실행 코드나 내부 운영 메모가 섞이지 않았는지 확인한다.
3. `starter-kits/` 안에 비밀값, 개인 경로, 임시 스크립트가 없는지 확인한다.
4. root `README.md`에서 공유용 진입 경로가 `docs/`, `starter-kits/`, `wiki/`로 깔끔하게 이어지는지 확인한다.
5. 공개 후 첫 사용자 입장에서 `README.md -> docs/hands-on-assets/README.md -> starter-kits/README.md` 순서가 자연스러운지 확인한다.

## 한 줄 기준

공유 workspace에는 `바로 읽거나 바로 실행할 수 있는 자료`만 남기고, 개인 운영 메모는 분리한다고 생각하면 된다.
