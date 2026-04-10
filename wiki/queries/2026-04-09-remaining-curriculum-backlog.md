---
title: Remaining Curriculum Backlog Refresh
type: query-note
status: active
updated: 2026-04-10
source_docs:
  - docs/math-curriculum-research/comparison.md
  - docs/math-curriculum-research/korea.md
  - docs/math-curriculum-research/japan.md
  - docs/math-curriculum-research/china.md
  - docs/math-curriculum-research/us.md
  - docs/math-concept-encyclopedia/korea.md
  - docs/math-concept-encyclopedia/japan.md
  - docs/math-concept-encyclopedia/china.md
  - docs/math-concept-encyclopedia/us.md
tags:
  - query
  - backlog
  - curriculum
---

# Remaining Curriculum Backlog Refresh

## Summary

예전에 남아 있던 `대표 개념 카드`, `나라별 학년·과목 wrapper`, `지속형 curriculum hub`, `4개국 비교 synthesis`, `예제·오개념·평가 패턴 synthesis` 중심 backlog는 이미 사실상 완료 상태였고, 이번 배치에서 마지막으로 눈에 띄던 `세분화 누락 노드`까지 persistent layer에 편입했다. 이제 `docs/` 기반으로 직접 보이던 필수 확장 항목은 모두 올라왔고, 남은 일은 `더 잘게 쪼갤 세부 카드`, `증명·예제 밀도 보강`, `메타데이터 정규화` 같은 선택형 심화 작업이다.

## Key Points

- 이전 대표 backlog의 상태
  - `소인수분해`, `정수와 유리수`, `문자와 식`, `식의 계산`, `다항식과 인수분해`, `제곱근과 실수`, `일차부등식`
  - `기본 도형과 작도`, `합동`, `닮음`, `원`, `피타고라스 정리`, `평행선·다각형·합동의 논증`
  - `반비례`, `지수함수와 로그함수`, `수열`, `함수 일반론`, `극한`, `미분`, `적분`
  - `집합과 명제`, `이차곡선`, `벡터`, `공간좌표`
  - `자료 정리`, `상관과 산점도`, `정규분포`, `표본조사`, `실용 통계`
  - `경제 수학`, `인공지능 수학`, `직무 수학`, `수학과 문화`, `Integrated Mathematics`
  - 위 목록은 현재 persistent concept layer에 모두 존재한다.
  - 추가로 `복소수`, `복소수평면`, `수학적 모델링·수학 탐구`, `Polar Ideas`, `Sequences and Series`, `AP Statistics`도 이번 배치로 반영되었다.
- 이번 배치에서 닫힌 항목
  - `일본 wrapper`: `중1`, `중2`, `중3`, `수학I`, `수학A`, `수학II`, `수학B`, `수학III`, `수학C`
  - `중국 wrapper`: `7학년 1학기`, `7학년 2학기`, `8학년 1학기`, `8학년 2학기`, `9학년 1학기`, `9학년 2학기`, `필수 과정`, `선택필수 과정`, `선택 과정의 개념군`
  - `미국 wrapper`: `Grade 6`, `Grade 7`, `Grade 8`
  - `나라별 curriculum hub`: 한국, 일본, 중국, 미국
  - `course-track hub`
  - `cross-country comparison synthesis`: 함수 위계, 기하 논증, 통계·데이터
  - `assessment synthesis`
  - `support synthesis`: 예제 패턴, 오개념 패턴, 평가 패턴
  - `세부 하위 카드`: `연속함수의 성질`, `AP Calculus`, `IB Mathematics`, `삼각함수의 그래프`, `복소수의 나눗셈`, `도함수의 그래프 해석`
- 이번 배치에서 추가로 닫힌 `세분화 누락 노드`
  - `분식`
  - `평행사변형`
  - `투영과 투상도`
  - `분산`
  - `표준편차`
  - `신뢰구간`
  - `가설검정`
  - `직선의 방정식`
  - `원의 방정식`
  - `포물선`
  - `타원`
  - `쌍곡선`
- 지금 실제로 남아 있는 선택형 심화 후보
  - `학년 허브`를 query sample에서 persistent synthesis로 승격
  - `정의역·치역`, `좌표평면`, `상자그림`, `사분위범위`, `독립사건`, `연속확률변수`, `베이즈 정리` 같은 더 세분된 하위 카드
  - 앵커 카드의 `정석형 서술 밀도`와 증명 밀도 추가 보강
  - concept card frontmatter와 alias 메타데이터 표준화

### Recommended Next Expansion Order

- 1차
  - `정석형 앵커 카드` 밀도 보강
  - `학년 허브`의 persistent 승격 여부 결정
- 2차
  - `정의역·치역`, `좌표평면`, `상자그림`, `사분위범위` 같은 하위 개념 분화
  - `proof density`가 약한 카드의 증명 보강
- 3차
  - `학생용/교사용` 난이도 뷰 분리
  - 그래프 연결 품질을 위한 alias/frontmatter 표준화

## Connections

- 현재 전체 반영 상태 감사는 [2026-04-06-research-folder-coverage-audit.md](./2026-04-06-research-folder-coverage-audit.md)에 있다.
- 개념 카드 구조는 [../syntheses/concept-card-model.md](../syntheses/concept-card-model.md)에 있다.
- 한국 샘플 그래프는 [2026-04-06-math-curriculum-graph.md](./2026-04-06-math-curriculum-graph.md)와 하위 샘플 노트들을 본다.
- 나라별 직접 요약은 [../source-notes/math-curriculum-research-japan.md](../source-notes/math-curriculum-research-japan.md), [../source-notes/math-curriculum-research-china.md](../source-notes/math-curriculum-research-china.md), [../source-notes/math-curriculum-research-us.md](../source-notes/math-curriculum-research-us.md), [../source-notes/math-concept-encyclopedia-japan.md](../source-notes/math-concept-encyclopedia-japan.md), [../source-notes/math-concept-encyclopedia-china.md](../source-notes/math-concept-encyclopedia-china.md), [../source-notes/math-concept-encyclopedia-us.md](../source-notes/math-concept-encyclopedia-us.md)에 있다.

## Open Questions

- `학년 허브`를 query sample로 유지할지, persistent synthesis로 승격할지 기준이 필요하다.
- `과목 wrapper 카드`와 `나라별 curriculum hub`를 같은 레이어에 둘지 한 단계 더 나눌지 정리가 필요하다.
- 남은 작업의 중심을 `새 카드 수 늘리기`보다 `기존 카드의 증명·예제 밀도 보강`으로 전환할지 판단이 필요하다.

## Sources

- `docs/math-curriculum-research/comparison.md`
- `docs/math-curriculum-research/korea.md`
- `docs/math-curriculum-research/japan.md`
- `docs/math-curriculum-research/china.md`
- `docs/math-curriculum-research/us.md`
- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-concept-encyclopedia/japan.md`
- `docs/math-concept-encyclopedia/china.md`
- `docs/math-concept-encyclopedia/us.md`
