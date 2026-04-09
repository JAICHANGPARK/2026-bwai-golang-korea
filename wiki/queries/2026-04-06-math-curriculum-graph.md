---
title: Math Curriculum Graph Design
type: query-note
status: active
updated: 2026-04-06
source_docs:
  - docs/math-agent-knowledge-plan.md
  - docs/math-curriculum-research/korea.md
tags:
  - query
  - curriculum
  - graph
---

# Math Curriculum Graph Design

## Summary

중1부터 고3까지의 수학 교과과정은 `학년별 목차`보다 `개념 그래프 + 학년 허브 + 교육과정 버전 + 평가 체제`를 겹쳐서 보는 구조가 더 적합하다. 특히 `2022 개정 교육과정`의 단계적 적용과 `2028학년도 대입 개편`을 구분해야 하므로, 옵시디언에서는 개념 노드를 중심에 두고 학년/과목/평가를 별도 뷰로 연결하는 설계가 안정적이다.

## Key Points

- 기본 단위는 `개념 노드`로 둔다.
  - 예: `일차방정식`, `정비례와 반비례`, `일차함수`, `이차함수`, `삼각비`, `확률분포`
- `학년 페이지`는 개념을 담는 폴더가 아니라 허브 페이지로 둔다.
  - 예: `중1`, `중2`, `중3`, `고1`, `고2`, `고3`
  - 역할: 해당 시기 대표 개념, 선행 개념, 다음 연결 개념을 모아서 보여준다.
- `계통 축 페이지`를 별도로 둔다.
  - 예: `함수 계통`, `방정식 계통`, `기하 계통`, `확률과 통계 계통`
  - 역할: 학년을 가로질러 같은 개념줄기를 추적한다.
- `교육과정 버전`과 `평가 체제`는 개념 자체와 분리한다.
  - `2015 개정`, `2022 개정`은 교육과정 버전 노드 또는 frontmatter 필드로 관리한다.
  - `현행 수능`, `2028 통합형 수능`, `내신 서술형`은 평가 레이어로 연결한다.
- `오개념`과 `복습 경로`를 별도 카드로 둔다.
  - 예: `기울기와 비율 혼동`, `음수로 나눌 때 부호 반전`, `닮음에서 대응변 혼동`

### Recommended Entity Types

- `GradeHub`
  - 중1~고3 시기별 대표 학습 지도를 보여준다.
- `StrandHub`
  - 함수, 방정식, 기하, 확률·통계 같은 계통별 허브다.
- `ConceptCard`
  - 실제 그래프의 중심 노드다.
- `CourseNode`
  - `공통수학1`, `대수`, `미적분I`처럼 고등학교 과목 노드다.
- `AssessmentNode`
  - `내신`, `수능`, `2028 통합형 수능`, `서술형` 같은 평가 맥락이다.
- `MisconceptionCard`
  - 반복 오답과 혼동 포인트를 저장한다.
- `ReviewPath`
  - 특정 개념 부족 시 되돌아가야 할 복습 경로다.

### Recommended Graph Edges

- `prerequisite_of`
  - 먼저 알아야 하는 개념 관계
- `revisited_in`
  - 같은 개념이 다른 학년/과목에서 다시 등장하는 관계
- `belongs_to_strand`
  - 함수, 기하, 확통 같은 계통 연결
- `mapped_to_course`
  - 고등 과목과의 연결
- `assessed_by`
  - 내신/수능/서술형 등의 평가 방식 연결
- `causes_error_in`
  - 특정 오개념이 어떤 문항 유형에서 문제를 만드는지 연결
- `recommends_review_of`
  - 복습 추천 경로 연결

### Recommended Obsidian Layout

- `math/grade-hubs/`
  - `middle-1.md`, `middle-2.md`, `middle-3.md`, `high-1.md`, `high-2.md`, `high-3.md`
- `math/strands/`
  - `equations.md`, `functions.md`, `geometry.md`, `statistics-and-probability.md`
- `math/concepts/`
  - 개념 카드 본체
- `math/courses/`
  - `common-math-1.md`, `common-math-2.md`, `algebra.md`, `calculus-1.md`
- `math/assessments/`
  - `school-exam.md`, `csat-current.md`, `csat-2028.md`
- `math/misconceptions/`
  - 오개념 카드
- `math/review-paths/`
  - 취약 개념별 복습 루트

### Starter Learning Progressions

- 함수 축
  - `정비례와 반비례 -> 일차함수 -> 이차함수 -> 함수와 그래프 -> 함수의 극한 -> 미분`
- 식과 방정식 축
  - `문자와 식 -> 일차방정식 -> 연립방정식 -> 인수분해 -> 이차방정식 -> 방정식과 함수 연결`
- 기하 축
  - `작도와 합동 -> 닮음 -> 피타고라스 정리 -> 삼각비 -> 도형의 방정식 -> 기하와 벡터`
- 확률·통계 축
  - `자료 정리 -> 경우의 수 -> 확률 -> 산포와 상관 -> 확률과 통계 -> 실용 통계`

### Recommended Frontmatter For Concept Cards

```yaml
---
title: 일차함수
type: concept-card
status: active
updated: 2026-04-06
curriculum_versions:
  - 2015
  - 2022
grade_views:
  - middle-2
  - high-1
strands:
  - functions
prerequisites:
  - 정비례와 반비례
  - 좌표평면
leads_to:
  - 기울기
  - 직선의 방정식
  - 함수 해석
assessment_links:
  - school-exam
  - csat-current
misconceptions:
  - 기울기와 비율 혼동
review_paths:
  - 비례에서 함수로
---
```

### Recommended Grade Hubs

- `중1`
  - 소인수분해, 정수와 유리수, 문자와 식, 일차방정식, 좌표와 비례, 기본 도형과 자료 정리 기초
- `중2`
  - 식의 계산, 일차부등식, 연립방정식, 일차함수, 합동과 논증, 확률과 분포 비교
- `중3`
  - 제곱근과 실수, 인수분해, 이차방정식, 이차함수, 닮음, 피타고라스 정리, 삼각비, 상관관계
- `고1`
  - `공통수학1`, `공통수학2`를 기준으로 다항식, 부등식, 경우의 수, 행렬, 도형의 방정식, 집합과 명제, 함수와 그래프
- `고2`
  - 대표적으로 `대수`, `미적분I`, `확률과 통계`
- `고3`
  - 대표적으로 `미적분II`, `기하`, `경제 수학`, `인공지능 수학`, `실용 통계`

### Design Notes About 2028

- `2028`은 새 수학 교육과정 명칭이 아니라 `2028학년도 대입/수능 체제`를 가리킨다.
- 따라서 그래프에서는 `교육과정 버전`과 `평가 체제`를 같은 축에 두지 않는다.
- 정리 기준
  - 교육과정 축: `2015 개정`, `2022 개정`
  - 평가 축: `현행 수능`, `2028 통합형 수능`, `내신`

### Sample Nodes Added In Wiki

- [math-curriculum-graph/middle-1-hub.md](./math-curriculum-graph/middle-1-hub.md)
- [math-curriculum-graph/middle-2-hub.md](./math-curriculum-graph/middle-2-hub.md)
- [math-curriculum-graph/middle-3-hub.md](./math-curriculum-graph/middle-3-hub.md)
- [math-curriculum-graph/high-1-hub.md](./math-curriculum-graph/high-1-hub.md)
- [math-curriculum-graph/high-2-hub.md](./math-curriculum-graph/high-2-hub.md)
- [math-curriculum-graph/high-3-hub.md](./math-curriculum-graph/high-3-hub.md)
- [math-curriculum-graph/functions-strand.md](./math-curriculum-graph/functions-strand.md)
- [math-curriculum-graph/geometry-strand.md](./math-curriculum-graph/geometry-strand.md)
- [math-curriculum-graph/statistics-and-probability-strand.md](./math-curriculum-graph/statistics-and-probability-strand.md)
- [math-curriculum-graph/linear-function.md](./math-curriculum-graph/linear-function.md)
- [math-curriculum-graph/quadratic-function.md](./math-curriculum-graph/quadratic-function.md)
- [math-curriculum-graph/trigonometric-ratio.md](./math-curriculum-graph/trigonometric-ratio.md)
- [math-curriculum-graph/probability-and-statistics-course.md](./math-curriculum-graph/probability-and-statistics-course.md)
- [math-curriculum-graph/csat-2028.md](./math-curriculum-graph/csat-2028.md)

## Connections

- 제품형 지식 인프라의 역할 정의는 [../components/internal-knowledge-base.md](../components/internal-knowledge-base.md)에 있다.
- 지식 정규화 우선순위는 [../syntheses/knowledge-roadmap.md](../syntheses/knowledge-roadmap.md)에서 이어진다.
- 한국 교육과정의 대표 배치는 [../../docs/math-curriculum-research/korea.md](../../docs/math-curriculum-research/korea.md)에 정리되어 있다.
- 교육과정과 개념 관계망이 핵심 자산이라는 점은 [../../docs/math-agent-knowledge-plan.md](../../docs/math-agent-knowledge-plan.md)에 직접 연결된다.

## Open Questions

- 성취기준을 `ConceptCard`에 내장할지 `AchievementStandard`를 별도 노드로 둘지 후속 결정이 필요하다.
- `2015 개정`과 `2022 개정`의 동일 개념을 하나의 노드로 합칠지, 별도 노드와 alias 관계로 둘지 기준이 더 필요하다.
- 고등학교 선택과목의 실제 학교 편차를 어느 수준까지 그래프에 반영할지 정해야 한다.
- 기출 문항 메타데이터를 붙일 때 저작권상 원문 없이 메타데이터만 저장하는 운영 규칙이 필요하다.

## Sources

- `docs/math-agent-knowledge-plan.md`
- `docs/math-curriculum-research/korea.md`
- 교육부 2022 개정 교육과정 고시: <https://www.moe.go.kr/boardCnts/viewRenew.do?boardID=141&boardSeq=93458&lev=0&search=>
- 교육부 2022 개정 수학과 변화 안내: <https://www.moe.go.kr/boardCnts/viewRenew.do?boardID=340&boardSeq=93073&lev=0&m=020201&opType=N&s=moe&statusYN=W>
- 교육부 2028학년도 대입 개편안: <https://www.moe.go.kr/boardCnts/viewRenew.do?boardID=294&boardSeq=103113&lev=0&m=020402>
- 경기도교육청 2022 개정 고등학교 교육과정 Q&A: <https://www.goe.go.kr/resource/old/BBSMSTR_000000030136/BBS_202411010914152012.pdf>
- 경기도교육청 2026 고등학교 학업성적관리 시행지침: <https://www.goe.go.kr/resource/goe/na/bbs_2675/2026/02/410e7664-15c3-434c-82a0-b6a4ee78eec1.pdf>
- 수학 learning progression 참고: <https://www.mathematica.org/publications/mathematical-learning-progressions-a-study-of-teachers-content-adaptations-and-alignment>
