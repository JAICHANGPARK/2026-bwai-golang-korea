---
title: Concept Card Model
type: synthesis
status: active
updated: 2026-04-09
source_docs:
  - docs/math-agent-knowledge-plan.md
  - docs/internal-knowledge-base-build.md
  - docs/math-curriculum-research/README.md
  - docs/math-concept-encyclopedia/README.md
  - docs/math-curriculum-research/japan.md
  - docs/math-curriculum-research/china.md
  - docs/math-curriculum-research/us.md
  - docs/math-concept-encyclopedia/japan.md
  - docs/math-concept-encyclopedia/china.md
  - docs/math-concept-encyclopedia/us.md
tags:
  - synthesis
  - concept-card
  - curriculum
---

# Concept Card Model

## Summary

학년 허브만으로는 중1부터 고3까지 이어지는 수학 개념의 재등장과 선행관계를 충분히 표현하기 어렵다. 그래서 지속 위키 계층에서는 `학년 허브`와 별도로 `개념 카드`를 기본 단위로 두는 것이 적합하다.

## Key Points

- 왜 개념 카드가 필요한가
  - `일차방정식`, `비례`, `일차함수`, `삼각비`, `확률분포` 같은 개념은 한 학년에만 머물지 않는다.
  - 학생 설명, 복습 추천, 오개념 진단은 학년보다 개념 단위가 더 잘 맞는다.
  - 제품형 지식베이스에서도 `ConceptCard`는 검색과 설명 보강의 핵심 엔터티다.
- 개념 카드에 넣을 기본 요소
  - 정의
  - 핵심 개념과 대표 수식
  - 증명 또는 증명 스케치
  - 대표 예제
  - 교육과정 배치
  - 국가별 배치 스냅샷
  - 연결 개념
  - 오개념과 복습 경로
- 증명 섹션 운영 원칙
  - 원천 문서에 직접 증명이 있으면 그 구조를 따른다.
  - 원천 문서에 없지만 개념 이해에 꼭 필요하면 `증명 스케치 (추론)`처럼 명시한다.
  - 과도한 형식 논증보다 중고등 학생 눈높이 설명을 우선한다.
- 위키 구조에서의 위치
  - `학년 허브`: 언제 배우는가
  - `계통 허브`: 무엇으로 이어지는가
  - `개념 카드`: 그 개념이 정확히 무엇인가
  - `과목·경로 카드`: 어떤 과목이나 course-track 안에서 그 개념군이 어떻게 묶이는가
  - `평가 노드`: 시험에서 어떻게 읽히는가
- 과목·경로 카드 운영 원칙
  - `공통수학1`, `미적분I`, `Algebra I`, `Precalculus`처럼 실제 수업 단위가 개념 묶음을 설명할 때는 별도 wrapper 카드를 둔다.
  - 개념 카드가 `무엇인가`를 설명한다면, 과목·경로 카드는 `어떻게 묶여 배우는가`를 설명한다.
  - 과목 카드는 가능한 한 원천 문서의 과목명과 구성 순서를 보존한다.
- 국가별 렌즈를 다루는 방법
  - 한국 기준 카드만 만들지 말고, 같은 개념이 `일본/중국/미국`에서 어디에 놓이는지 함께 적으면 비교와 전이 학습에 도움이 된다.
  - 미국처럼 `학년 고정형`이 아닌 경우는 `Grade 8 또는 Algebra I`, `Geometry 또는 Precalculus`처럼 `경로형 배치`로 적는다.
  - 어떤 나라 문서에 개념이 독립 단원으로 없으면 `문서상 분산 배치` 또는 `상위 과목에서 본격화`라고 적어 과잉 일반화를 피한다.
- 현재 규모 메모
  - persistent concept layer는 현재 `75개` 카드로 구성된다.
- 현재 카드 세트
  - 수와 대수 축: [concept-cards/prime-factorization.md](./concept-cards/prime-factorization.md), [concept-cards/integers-and-rational-numbers.md](./concept-cards/integers-and-rational-numbers.md), [concept-cards/variables-and-expressions.md](./concept-cards/variables-and-expressions.md), [concept-cards/algebraic-manipulation.md](./concept-cards/algebraic-manipulation.md), [concept-cards/linear-equation.md](./concept-cards/linear-equation.md), [concept-cards/linear-inequality.md](./concept-cards/linear-inequality.md), [concept-cards/simultaneous-equations.md](./concept-cards/simultaneous-equations.md), [concept-cards/matrix.md](./concept-cards/matrix.md), [concept-cards/square-roots-and-real-numbers.md](./concept-cards/square-roots-and-real-numbers.md), [concept-cards/polynomials-and-factorization.md](./concept-cards/polynomials-and-factorization.md), [concept-cards/quadratic-equation.md](./concept-cards/quadratic-equation.md), [concept-cards/complex-numbers.md](./concept-cards/complex-numbers.md), [concept-cards/mathematical-induction.md](./concept-cards/mathematical-induction.md)
  - 함수와 해석 축: [concept-cards/proportion.md](./concept-cards/proportion.md), [concept-cards/inverse-proportion.md](./concept-cards/inverse-proportion.md), [concept-cards/function-foundations.md](./concept-cards/function-foundations.md), [concept-cards/linear-function.md](./concept-cards/linear-function.md), [concept-cards/quadratic-function.md](./concept-cards/quadratic-function.md), [concept-cards/exponential-and-logarithmic-functions.md](./concept-cards/exponential-and-logarithmic-functions.md), [concept-cards/sequences.md](./concept-cards/sequences.md), [concept-cards/sequences-and-series.md](./concept-cards/sequences-and-series.md), [concept-cards/limits.md](./concept-cards/limits.md), [concept-cards/continuity.md](./concept-cards/continuity.md), [concept-cards/differentiation.md](./concept-cards/differentiation.md), [concept-cards/integration.md](./concept-cards/integration.md), [concept-cards/trigonometric-ratio.md](./concept-cards/trigonometric-ratio.md), [concept-cards/trigonometric-function.md](./concept-cards/trigonometric-function.md), [concept-cards/polar-ideas.md](./concept-cards/polar-ideas.md)
  - 기하와 좌표 축: [concept-cards/basic-geometry-and-construction.md](./concept-cards/basic-geometry-and-construction.md), [concept-cards/geometric-transformations.md](./concept-cards/geometric-transformations.md), [concept-cards/geometric-proof.md](./concept-cards/geometric-proof.md), [concept-cards/congruence.md](./concept-cards/congruence.md), [concept-cards/similarity.md](./concept-cards/similarity.md), [concept-cards/circle.md](./concept-cards/circle.md), [concept-cards/pythagorean-theorem.md](./concept-cards/pythagorean-theorem.md), [concept-cards/sets-and-propositions.md](./concept-cards/sets-and-propositions.md), [concept-cards/equations-of-geometric-figures.md](./concept-cards/equations-of-geometric-figures.md), [concept-cards/conic-sections.md](./concept-cards/conic-sections.md), [concept-cards/vectors.md](./concept-cards/vectors.md), [concept-cards/complex-plane.md](./concept-cards/complex-plane.md), [concept-cards/spatial-coordinates.md](./concept-cards/spatial-coordinates.md)
  - 확률과 통계 축: [concept-cards/data-organization.md](./concept-cards/data-organization.md), [concept-cards/counting-principles.md](./concept-cards/counting-principles.md), [concept-cards/sampling.md](./concept-cards/sampling.md), [concept-cards/correlation-and-scatter-plots.md](./concept-cards/correlation-and-scatter-plots.md), [concept-cards/correlation-coefficient.md](./concept-cards/correlation-coefficient.md), [concept-cards/conditional-probability.md](./concept-cards/conditional-probability.md), [concept-cards/random-variable.md](./concept-cards/random-variable.md), [concept-cards/probability-distribution.md](./concept-cards/probability-distribution.md), [concept-cards/statistical-inference.md](./concept-cards/statistical-inference.md), [concept-cards/normal-distribution.md](./concept-cards/normal-distribution.md), [concept-cards/regression.md](./concept-cards/regression.md), [concept-cards/practical-statistics.md](./concept-cards/practical-statistics.md)
  - 과목·응용·경로 축: [concept-cards/common-math-1.md](./concept-cards/common-math-1.md), [concept-cards/common-math-2.md](./concept-cards/common-math-2.md), [concept-cards/korean-algebra-course.md](./concept-cards/korean-algebra-course.md), [concept-cards/calculus-1.md](./concept-cards/calculus-1.md), [concept-cards/probability-and-statistics-course.md](./concept-cards/probability-and-statistics-course.md), [concept-cards/calculus-2.md](./concept-cards/calculus-2.md), [concept-cards/korean-geometry-course.md](./concept-cards/korean-geometry-course.md), [concept-cards/basic-math-1-and-2.md](./concept-cards/basic-math-1-and-2.md), [concept-cards/algebra-1.md](./concept-cards/algebra-1.md), [concept-cards/geometry-course.md](./concept-cards/geometry-course.md), [concept-cards/algebra-2.md](./concept-cards/algebra-2.md), [concept-cards/precalculus.md](./concept-cards/precalculus.md), [concept-cards/calculus-course.md](./concept-cards/calculus-course.md), [concept-cards/statistics-and-data-analysis.md](./concept-cards/statistics-and-data-analysis.md), [concept-cards/ap-statistics.md](./concept-cards/ap-statistics.md), [concept-cards/advanced-tracks.md](./concept-cards/advanced-tracks.md), [concept-cards/mathematical-modeling-and-inquiry.md](./concept-cards/mathematical-modeling-and-inquiry.md), [concept-cards/economics-math.md](./concept-cards/economics-math.md), [concept-cards/mathematics-for-ai.md](./concept-cards/mathematics-for-ai.md), [concept-cards/workplace-math.md](./concept-cards/workplace-math.md), [concept-cards/mathematics-and-culture.md](./concept-cards/mathematics-and-culture.md), [concept-cards/integrated-mathematics.md](./concept-cards/integrated-mathematics.md)

## Connections

- 지식 엔터티 설계는 [../components/internal-knowledge-base.md](../components/internal-knowledge-base.md)에 있다.
- 지식 구축 순서는 [./knowledge-roadmap.md](./knowledge-roadmap.md)와 연결된다.
- 학년 허브와 계통 허브의 샘플은 [../queries/2026-04-06-math-curriculum-graph.md](../queries/2026-04-06-math-curriculum-graph.md) 및 하위 샘플 페이지에 있다.
- 개념 백과사전 폴더의 성격은 [../source-notes/math-concept-encyclopedia-readme.md](../source-notes/math-concept-encyclopedia-readme.md)를 본다.

## Open Questions

- `개념 카드`를 frontmatter 기준으로 더 엄격히 표준화할지 아직 정하지 않았다.
- `개념 카드`와 `과목·경로 카드`를 같은 폴더와 템플릿 안에 둘지 더 분리할지 기준이 필요하다.
- 하나의 카드에 `중등 설명`과 `고등 설명`을 함께 둘지, 난이도 뷰를 분리할지 기준이 더 필요하다.
- 증명 난이도를 학생용, 교사용, 검증용으로 나눌지 후속 설계가 필요하다.
- 국가별 카드 차이를 `별도 문서`로 분기할지, 하나의 카드 안에 `국가별 배치 스냅샷`으로 유지할지 기준이 더 필요하다.

## Sources

- `docs/math-agent-knowledge-plan.md`
- `docs/internal-knowledge-base-build.md`
- `docs/math-curriculum-research/README.md`
- `docs/math-concept-encyclopedia/README.md`
- `docs/math-curriculum-research/japan.md`
- `docs/math-curriculum-research/china.md`
- `docs/math-curriculum-research/us.md`
- `docs/math-concept-encyclopedia/japan.md`
- `docs/math-concept-encyclopedia/china.md`
- `docs/math-concept-encyclopedia/us.md`
