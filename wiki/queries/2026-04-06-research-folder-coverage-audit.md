---
title: Research Folder Coverage Audit
type: query-note
status: active
updated: 2026-04-10
source_docs:
  - docs/math-curriculum-research/README.md
  - docs/math-curriculum-research/korea.md
  - docs/math-curriculum-research/comparison.md
  - docs/math-curriculum-research/japan.md
  - docs/math-curriculum-research/china.md
  - docs/math-curriculum-research/us.md
  - docs/math-concept-encyclopedia/README.md
  - docs/math-concept-encyclopedia/korea.md
  - docs/math-concept-encyclopedia/japan.md
  - docs/math-concept-encyclopedia/china.md
  - docs/math-concept-encyclopedia/us.md
  - docs/math-concept-encyclopedia/multilingual-glossary.md
  - docs/math-concept-encyclopedia/formula-examples.md
  - docs/math-concept-encyclopedia/unit-practice-book.md
  - docs/math-concept-encyclopedia/comparative-problem-book.md
tags:
  - query
  - audit
  - coverage
---

# Research Folder Coverage Audit

## Summary

이번 턴 기준으로 `docs/math-curriculum-research/`와 `docs/math-concept-encyclopedia/`의 문서군은 모두 `source-note` 계층에 반영되었고, concept card layer도 `119개`까지 확장되었다. 이전 감사 노트에서 남겨 두었던 `세부 하위 개념 카드`, `나라별 wrapper 카드`, `나라별 curriculum hub`, `course-track hub`, `나라 비교 synthesis`, `예제·오개념·평가 패턴 synthesis`는 물론, 마지막으로 남아 있던 `세분화 누락 노드 12개`도 이번 배치에서 persistent layer로 승격되었다. 이제 남은 과제는 `새 docs 기반 필수 반영`이 아니라 카드 밀도와 메타데이터를 다듬는 선택형 심화 작업이다.

## Key Points

- 전체 상태
  - `math-curriculum-research/`: source-note 계층 반영 완료
  - `math-concept-encyclopedia/`: source-note 계층 반영 완료
  - `persistent concept layer`: 119개 카드에 국가별 배치 또는 course-track 정보 반영
  - `persistent synthesis layer`: 국가 허브, 비교 synthesis, support synthesis 반영 완료
- 현재 직접 반영된 문서
  - `docs/math-curriculum-research/README.md`
  - `docs/math-curriculum-research/korea.md`
  - `docs/math-curriculum-research/comparison.md`
  - `docs/math-curriculum-research/japan.md`
  - `docs/math-curriculum-research/china.md`
  - `docs/math-curriculum-research/us.md`
  - `docs/math-concept-encyclopedia/README.md`
  - `docs/math-concept-encyclopedia/korea.md`
  - `docs/math-concept-encyclopedia/japan.md`
  - `docs/math-concept-encyclopedia/china.md`
  - `docs/math-concept-encyclopedia/us.md`
  - `docs/math-concept-encyclopedia/multilingual-glossary.md`
  - `docs/math-concept-encyclopedia/formula-examples.md`
  - `docs/math-concept-encyclopedia/unit-practice-book.md`
  - `docs/math-concept-encyclopedia/comparative-problem-book.md`
- 이번 배치에서 추가된 카드군
  - 세부 하위 개념: `행렬`, `수학적 귀납법`, `연속성`, `상관계수`, `도형의 이동`
  - 한국 과목 wrapper: `공통수학1`, `공통수학2`, `대수`, `미적분I`, `확률과 통계`, `미적분II`, `기하`, `기본수학1·2`
  - 미국 과목·경로 wrapper: `Algebra I`, `Geometry`, `Algebra II`, `Precalculus`, `Calculus`, `Statistics and Data Analysis`, `Advanced Tracks`
  - 최근 추가 카드: `복소수`, `복소수평면`, `수학적 모델링·수학 탐구`, `Polar Ideas`, `Sequences and Series`, `AP Statistics`, `복소수의 연산`, `행렬의 연산`, `라디안`, `삼각함수 항등식`, `도함수`
- 이번 배치에서 닫힌 확장 후보
  - synthesis 확장: `나라별 curriculum hub`, `course-track hub`, `cross-country comparison synthesis`
  - support synthesis: `example pattern`, `misconception map`, `assessment pattern`
  - 추가 하위 개념: `연속함수의 성질`, `삼각함수의 그래프`, `복소수의 나눗셈`, `도함수의 그래프 해석`, `AP Calculus`, `IB Mathematics`
  - 미국 중학교 wrapper: `Grade 6`, `Grade 7`, `Grade 8`
  - finer gap cards: `분식`, `평행사변형`, `투영과 투상도`, `분산`, `표준편차`, `신뢰구간`, `가설검정`, `직선의 방정식`, `원의 방정식`, `포물선`, `타원`, `쌍곡선`

### Folder Status Details

- `docs/math-curriculum-research/`
  - `README.md`, `korea.md`, `comparison.md`, `japan.md`, `china.md`, `us.md` 모두 source note가 있다.
  - `korea/japan/china/us` 문서는 concept card의 `국가별 배치 스냅샷`과 `course-track note`에 직접 반영된다.
- `docs/math-concept-encyclopedia/`
  - `README.md`, `korea.md`, `japan.md`, `china.md`, `us.md` 모두 source note가 있다.
  - `multilingual-glossary.md`, `formula-examples.md`, `unit-practice-book.md`, `comparative-problem-book.md`도 source note로 편입되어 alias, 수식, 예제, 문제 감각 레이어를 지원한다.

### Recommended Next Expansion Order

- 1차
  - `proof-rich anchor cards` 보강
  - `persistent grade hub` 승격 판단
- 2차
  - alias/frontmatter 표준화
  - finer theorem subcards
- 3차
  - 학습자 레벨별 뷰 분기
  - 검색·추천용 메타데이터 추가

## Connections

- 폴더 설계의 원본 메모는 [../source-notes/math-curriculum-research-readme.md](../source-notes/math-curriculum-research-readme.md), [../source-notes/math-concept-encyclopedia-readme.md](../source-notes/math-concept-encyclopedia-readme.md)에 있다.
- 현재 한국 기준 그래프 샘플은 [2026-04-06-math-curriculum-graph.md](./2026-04-06-math-curriculum-graph.md)와 하위 샘플 노트들에 있다.
- 나라별 직접 요약은 [../source-notes/math-curriculum-research-japan.md](../source-notes/math-curriculum-research-japan.md), [../source-notes/math-curriculum-research-china.md](../source-notes/math-curriculum-research-china.md), [../source-notes/math-curriculum-research-us.md](../source-notes/math-curriculum-research-us.md), [../source-notes/math-concept-encyclopedia-japan.md](../source-notes/math-concept-encyclopedia-japan.md), [../source-notes/math-concept-encyclopedia-china.md](../source-notes/math-concept-encyclopedia-china.md), [../source-notes/math-concept-encyclopedia-us.md](../source-notes/math-concept-encyclopedia-us.md)에 있다.
- 지속 개념 카드 구조는 [../syntheses/concept-card-model.md](../syntheses/concept-card-model.md), [../syntheses/concept-cards/prime-factorization.md](../syntheses/concept-cards/prime-factorization.md), [../syntheses/concept-cards/linear-function.md](../syntheses/concept-cards/linear-function.md), [../syntheses/concept-cards/equations-of-geometric-figures.md](../syntheses/concept-cards/equations-of-geometric-figures.md), [../syntheses/concept-cards/differentiation.md](../syntheses/concept-cards/differentiation.md), [../syntheses/concept-cards/statistical-inference.md](../syntheses/concept-cards/statistical-inference.md), [../syntheses/concept-cards/integrated-mathematics.md](../syntheses/concept-cards/integrated-mathematics.md)에서 본다.

## Open Questions

- 다음 우선순위를 `proof density 보강`과 `새 세부 카드 추가` 중 어디에 둘지 결정이 필요하다.
- 미국처럼 `경로형` 교육과정을 별도 허브로 뺀 현재 구조를 더 세분할지 기준이 필요하다.
- `formula-examples.md`, `unit-practice-book.md`, `comparative-problem-book.md`를 검색 카드 단위로 더 쪼갤지 문서형 자산으로 유지할지 후속 설계가 필요하다.

## Sources

- `docs/math-curriculum-research/README.md`
- `docs/math-curriculum-research/korea.md`
- `docs/math-curriculum-research/comparison.md`
- `docs/math-curriculum-research/japan.md`
- `docs/math-curriculum-research/china.md`
- `docs/math-curriculum-research/us.md`
- `docs/math-concept-encyclopedia/README.md`
- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-concept-encyclopedia/japan.md`
- `docs/math-concept-encyclopedia/china.md`
- `docs/math-concept-encyclopedia/us.md`
- `docs/math-concept-encyclopedia/multilingual-glossary.md`
- `docs/math-concept-encyclopedia/formula-examples.md`
- `docs/math-concept-encyclopedia/unit-practice-book.md`
- `docs/math-concept-encyclopedia/comparative-problem-book.md`
