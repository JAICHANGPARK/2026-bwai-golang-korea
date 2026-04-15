---
title: Wiki Index
type: overview
status: active
updated: 2026-04-15
source_docs:
  - docs/hands-on-math-agent-session-guide.md
  - docs/prompt-pack/README.md
  - docs/prompt-pack/02a-wiki-enricher-extension.md
  - docs/llm-wiki-usage-guide.md
  - docs/antigravity-math-agent-step-by-step.md
  - docs/gemini-cli-math-agent-step-by-step.md
  - docs/antigravity-workspace-rules-draft.md
  - docs/antigravity-workflow-usage-guide.md
  - docs/workspace-distribution-guide.md
  - docs/math-agent-system-spec.md
  - docs/math-agent-knowledge-plan.md
  - docs/internal-knowledge-base-build.md
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
  - index
  - navigation
---

# Wiki Index

이 파일은 현재 위키의 콘텐츠 카탈로그다. 질문에 답하거나 새 문서를 ingest할 때 가장 먼저 읽는다.

## Overviews

- [README.md](./README.md) - `Build with AI Golang Korea 2026` 핸즈온 workspace의 목적, 범위, 빠른 진입점을 설명한다.
- [overview/project-map.md](./overview/project-map.md) - 저장소 전체를 `핸즈온`과 `제품 확장` 두 축으로 정리한 지도 페이지다.

## Profiles

- [profiles/hands-on-profile.md](./profiles/hands-on-profile.md) - 1시간 워크숍 프로파일의 목표, 제약, 산출물과 shared workspace 패키징 기준, optional wiki enrichment 확장을 정리한다. Sources: 9
- [profiles/product-expansion-profile.md](./profiles/product-expansion-profile.md) - 제품형 확장 프로파일의 기능 범위와 지식 인프라 요구를 정리한다. Sources: 5

## Components

- [components/solver-explainer-agent.md](./components/solver-explainer-agent.md) - 문제 풀이와 학생용 설명을 담당하는 에이전트의 역할을 정리한다. Sources: 3
- [components/expert-verifier-agent.md](./components/expert-verifier-agent.md) - 엄격한 수학 검증과 승인 결정을 담당하는 에이전트의 역할을 정리한다. Sources: 3
- [components/internal-knowledge-base.md](./components/internal-knowledge-base.md) - 내부 지식 계층의 층위, 타입, 활용 범위를 정리한다. Sources: 4
- [components/prompt-pack.md](./components/prompt-pack.md) - 핸즈온용 프롬프트 팩의 사용 순서와 범위를 정리한다. Sources: 3
- [components/wiki-enricher-agent.md](./components/wiki-enricher-agent.md) - reviewer 승인 후 `LLM Wiki`를 읽어 난이도, 개념, 교육과정, 연관 학습 주제를 보강하는 확장 에이전트 역할을 정리한다. Sources: 2
- [components/local-llm-wiki-ops.md](./components/local-llm-wiki-ops.md) - clone 이후 `Gemini CLI`와 `Antigravity`로 이 저장소를 LLM Wiki 및 hands-on workspace로 운영하고, 공개용 자산과 starter kit을 어떻게 나눌지 정리한다. Sources: 6
- [components/antigravity-workspace-automation.md](./components/antigravity-workspace-automation.md) - Antigravity에서 `Rules`, `Workflows`, `Skills`를 어떻게 나눠 구성할지 정리한다. Sources: 2

## Syntheses

- [syntheses/profile-boundary.md](./syntheses/profile-boundary.md) - 핸즈온 프로파일과 제품 확장 프로파일의 경계를 비교한다. Sources: 5
- [syntheses/knowledge-roadmap.md](./syntheses/knowledge-roadmap.md) - 수집, 정규화, 검색, 설명까지의 지식 구축 로드맵을 요약한다. Sources: 5
- [syntheses/concept-card-model.md](./syntheses/concept-card-model.md) - 학년 허브와 별도로 `개념 카드`를 지속 위키 자산으로 두는 구조, `정석형 서술 + 그래프 연결` 기준, `card_role: concept|wrapper` 분리 규칙을 정리한다.
- [syntheses/korea-curriculum-hub.md](./syntheses/korea-curriculum-hub.md) - 한국의 학년형 중학 구조와 과목형 고교 구조를 묶는 허브다.
- [syntheses/middle-1-hub.md](./syntheses/middle-1-hub.md) - 한국 `중1`의 시작 개념군을 묶는 persistent 학년 허브다.
- [syntheses/middle-2-hub.md](./syntheses/middle-2-hub.md) - 한국 `중2`의 선형 관계·논증·확률 입문을 묶는 persistent 학년 허브다.
- [syntheses/middle-3-hub.md](./syntheses/middle-3-hub.md) - 한국 `중3`의 이차·닮음·삼각비·상관을 묶는 persistent 학년 허브다.
- [syntheses/high-1-hub.md](./syntheses/high-1-hub.md) - 한국 `고1` 공통 기반을 묶는 persistent 학년 허브다.
- [syntheses/high-2-hub.md](./syntheses/high-2-hub.md) - 한국 `고2` 선택 과목 분기를 묶는 persistent 학년 허브다.
- [syntheses/high-3-hub.md](./syntheses/high-3-hub.md) - 한국 `고3` 심화·응용 경로를 묶는 persistent 학년 허브다.
- [syntheses/japan-curriculum-hub.md](./syntheses/japan-curriculum-hub.md) - 일본의 학년형 중학과 `수학I/A/II/B/III/C` 고교 구조를 묶는 허브다.
- [syntheses/china-curriculum-hub.md](./syntheses/china-curriculum-hub.md) - 중국의 학기형 중학과 `필수/선택필수` 고교 구조를 묶는 허브다.
- [syntheses/us-curriculum-hub.md](./syntheses/us-curriculum-hub.md) - 미국의 Grade 6~8과 course-track 고교 구조를 묶는 허브다.
- [syntheses/course-track-hub.md](./syntheses/course-track-hub.md) - 네 나라의 과목·경로 구조를 비교하는 메타 허브다.
- [syntheses/functions-strand.md](./syntheses/functions-strand.md) - 한국 수학의 함수 성장선을 비례에서 미적분까지 이어 보는 persistent 계통 허브다.
- [syntheses/geometry-strand.md](./syntheses/geometry-strand.md) - 작도에서 벡터까지 이어지는 기하 성장선을 묶는 persistent 계통 허브다.
- [syntheses/statistics-and-probability-strand.md](./syntheses/statistics-and-probability-strand.md) - 자료 읽기에서 분포와 추정까지 이어지는 확통 성장선을 묶는 persistent 계통 허브다.
- [syntheses/cross-country-function-progression.md](./syntheses/cross-country-function-progression.md) - 함수 위계의 국가별 전이 방식을 비교한다.
- [syntheses/cross-country-geometry-and-proof.md](./syntheses/cross-country-geometry-and-proof.md) - 기하와 증명 전면화 시점을 국가별로 비교한다.
- [syntheses/cross-country-statistics-and-data.md](./syntheses/cross-country-statistics-and-data.md) - 자료 읽기에서 추정·회귀까지의 국가별 전개를 비교한다.
- [syntheses/csat-2028.md](./syntheses/csat-2028.md) - 한국 `2028학년도 수능`을 교육과정 축과 분리된 평가 맥락 노드로 정리한다.
- [syntheses/assessment-systems-comparison.md](./syntheses/assessment-systems-comparison.md) - 수능, 공통테스트, 가오카오, SAT/holistic admissions의 차이를 비교한다.
- [syntheses/example-pattern-library.md](./syntheses/example-pattern-library.md) - 개념 카드에 재사용할 예제 패턴의 유형을 정리한다.
- [syntheses/misconception-pattern-map.md](./syntheses/misconception-pattern-map.md) - 반복되는 수학 오개념 패턴을 축별로 정리한다.
- [syntheses/assessment-patterns.md](./syntheses/assessment-patterns.md) - 나라별 문항 감각과 평가 패턴을 비교한다.

## Concept Syntheses

현재 `wiki/syntheses/concept-cards/`는 `card_role: concept 79장`, `card_role: wrapper 40장`, 총 `119장`으로 나뉘며, wrapper 카드는 Summary 아래 `Course Wrapper` callout으로 구분한다. 2026-04-10 기준으로 함수·기하·확통·미적분의 핵심 앵커 카드뿐 아니라 남아 있던 concept 카드 전반에도 `Deep Dive`, `Worked Examples`, `Common Pitfalls`, `증명 스케치` 축의 보강이 들어가, 모든 concept 카드가 최소 두 가지 이상의 교재형 서술 요소를 갖도록 정리되어 있다.

### Number And Algebra

- [syntheses/concept-cards/prime-factorization.md](./syntheses/concept-cards/prime-factorization.md) - `소인수분해`를 자연수 구조와 최대공약수·최소공배수의 기반으로 정리한다.
- [syntheses/concept-cards/integers-and-rational-numbers.md](./syntheses/concept-cards/integers-and-rational-numbers.md) - `정수와 유리수`를 수직선, 절댓값, 사칙연산 중심으로 정리한다.
- [syntheses/concept-cards/variables-and-expressions.md](./syntheses/concept-cards/variables-and-expressions.md) - `문자와 식`을 변수, 항, 계수, 일반화 언어 중심으로 정리한다.
- [syntheses/concept-cards/algebraic-manipulation.md](./syntheses/concept-cards/algebraic-manipulation.md) - `식의 계산`을 분배법칙과 동류항 정리 중심으로 정리한다.
- [syntheses/concept-cards/linear-equation.md](./syntheses/concept-cards/linear-equation.md) - `일차방정식`의 정의, 증명 스케치, 예제, 연결 개념을 정리한다.
- [syntheses/concept-cards/linear-inequality.md](./syntheses/concept-cards/linear-inequality.md) - `일차부등식`을 해의 범위와 부등호 반전 규칙 중심으로 정리한다.
- [syntheses/concept-cards/simultaneous-equations.md](./syntheses/concept-cards/simultaneous-equations.md) - `연립방정식`을 소거법, 교점 해석, 선형 관계 브리지 중심으로 정리한다.
- [syntheses/concept-cards/matrix.md](./syntheses/concept-cards/matrix.md) - `행렬`을 배열 표현과 선형 관계 브리지 중심으로 정리한다.
- [syntheses/concept-cards/matrix-operations.md](./syntheses/concept-cards/matrix-operations.md) - `행렬의 연산`을 성분별 덧셈과 행렬곱의 구조 중심으로 정리한다.
- [syntheses/concept-cards/square-roots-and-real-numbers.md](./syntheses/concept-cards/square-roots-and-real-numbers.md) - `제곱근과 실수`를 무리수와 실수 체계 확장 중심으로 정리한다.
- [syntheses/concept-cards/polynomials-and-factorization.md](./syntheses/concept-cards/polynomials-and-factorization.md) - `다항식과 인수분해`를 전개와 구조 인식 중심으로 정리한다.
- [syntheses/concept-cards/fractional-expressions.md](./syntheses/concept-cards/fractional-expressions.md) - `분식`을 약분, 통분, 정의역 감각과 함께 정리한다.
- [syntheses/concept-cards/quadratic-equation.md](./syntheses/concept-cards/quadratic-equation.md) - `이차방정식`을 근의 공식, 판별식, x절편 연결 중심으로 정리한다.
- [syntheses/concept-cards/complex-numbers.md](./syntheses/concept-cards/complex-numbers.md) - `복소수`를 허수단위, 해의 확장, 복소평면 브리지 중심으로 정리한다.
- [syntheses/concept-cards/complex-number-operations.md](./syntheses/concept-cards/complex-number-operations.md) - `복소수의 연산`을 $i^2=-1$ 규칙과 켤레복소수 감각 중심으로 정리한다.
- [syntheses/concept-cards/complex-division.md](./syntheses/concept-cards/complex-division.md) - `복소수의 나눗셈`을 켤레복소수를 이용한 분모 유리화 중심으로 정리한다.
- [syntheses/concept-cards/mathematical-induction.md](./syntheses/concept-cards/mathematical-induction.md) - `수학적 귀납법`을 자연수 명제 증명 방법 중심으로 정리한다.

### Functions And Calculus

- [syntheses/concept-cards/proportion.md](./syntheses/concept-cards/proportion.md) - `비례`를 정비례·반비례, 좌표, 함수 연결 중심으로 정리한다.
- [syntheses/concept-cards/inverse-proportion.md](./syntheses/concept-cards/inverse-proportion.md) - `반비례`를 $y=a/x$ 꼴의 비선형 관계로 정리한다.
- [syntheses/concept-cards/function-foundations.md](./syntheses/concept-cards/function-foundations.md) - `함수 일반론`을 정의역, 치역, 대응 규칙 중심으로 정리한다.
- [syntheses/concept-cards/linear-function.md](./syntheses/concept-cards/linear-function.md) - `일차함수`를 기울기, 그래프, 변화율 중심으로 정리한다.
- [syntheses/concept-cards/quadratic-function.md](./syntheses/concept-cards/quadratic-function.md) - `이차함수`를 꼭짓점, 축, 포물선 해석, 국가별 배치와 함께 정리한다.
- [syntheses/concept-cards/exponential-and-logarithmic-functions.md](./syntheses/concept-cards/exponential-and-logarithmic-functions.md) - `지수함수와 로그함수`를 성장, 역함수성, 로그법칙 중심으로 정리한다.
- [syntheses/concept-cards/sequences.md](./syntheses/concept-cards/sequences.md) - `수열`을 일반항, 점화식, 등차·등비수열 중심으로 정리한다.
- [syntheses/concept-cards/sequences-and-series.md](./syntheses/concept-cards/sequences-and-series.md) - `Sequences and Series`를 수열과 급수의 누적 관점 중심으로 정리한다.
- [syntheses/concept-cards/limits.md](./syntheses/concept-cards/limits.md) - `극한`을 무한 과정과 연속성의 언어로 정리한다.
- [syntheses/concept-cards/continuity.md](./syntheses/concept-cards/continuity.md) - `연속성`을 극한과 함수값의 연결 규칙 중심으로 정리한다.
- [syntheses/concept-cards/continuity-properties.md](./syntheses/concept-cards/continuity-properties.md) - `연속함수의 성질`을 중간값과 최댓값·최솟값 감각 중심으로 정리한다.
- [syntheses/concept-cards/differentiation.md](./syntheses/concept-cards/differentiation.md) - `미분`을 순간변화율과 접선 기울기 중심으로 정리한다.
- [syntheses/concept-cards/derivative.md](./syntheses/concept-cards/derivative.md) - `도함수`를 미분의 결과 함수와 증가·감소 해석 중심으로 정리한다.
- [syntheses/concept-cards/derivative-graph-interpretation.md](./syntheses/concept-cards/derivative-graph-interpretation.md) - `도함수의 그래프 해석`을 증가·감소, 극값, 볼록성 감각 중심으로 정리한다.
- [syntheses/concept-cards/integration.md](./syntheses/concept-cards/integration.md) - `적분`을 넓이와 누적량의 언어로 정리한다.
- [syntheses/concept-cards/trigonometric-ratio.md](./syntheses/concept-cards/trigonometric-ratio.md) - `삼각비`를 닮음 기반 증명 스케치와 함께 정리한다.
- [syntheses/concept-cards/trigonometric-function.md](./syntheses/concept-cards/trigonometric-function.md) - `삼각함수`를 단위원, 주기성, 항등식, 국가별 배치와 함께 정리한다.
- [syntheses/concept-cards/trigonometric-graphs.md](./syntheses/concept-cards/trigonometric-graphs.md) - `삼각함수의 그래프`를 진폭, 주기, 위상이동 해석 중심으로 정리한다.
- [syntheses/concept-cards/radians.md](./syntheses/concept-cards/radians.md) - `라디안`을 호의 길이와 미적분 전단계 브리지 중심으로 정리한다.
- [syntheses/concept-cards/trigonometric-identities.md](./syntheses/concept-cards/trigonometric-identities.md) - `삼각함수 항등식`을 피타고라스 관계와 함수 변환 문법 중심으로 정리한다.
- [syntheses/concept-cards/polar-ideas.md](./syntheses/concept-cards/polar-ideas.md) - `Polar Ideas`를 거리와 각의 표현 전환, Precalculus 브리지 중심으로 정리한다.

### Geometry And Coordinates

- [syntheses/concept-cards/basic-geometry-and-construction.md](./syntheses/concept-cards/basic-geometry-and-construction.md) - `기본 도형과 작도`를 점·선·각·자취 감각 중심으로 정리한다.
- [syntheses/concept-cards/geometric-transformations.md](./syntheses/concept-cards/geometric-transformations.md) - `도형의 이동`을 변환 관점의 기하 언어로 정리한다.
- [syntheses/concept-cards/geometric-proof.md](./syntheses/concept-cards/geometric-proof.md) - `평행선·다각형·합동의 논증`을 증명 구조 중심으로 정리한다.
- [syntheses/concept-cards/congruence.md](./syntheses/concept-cards/congruence.md) - `합동`을 대응과 변환 기하 관점으로 정리한다.
- [syntheses/concept-cards/similarity.md](./syntheses/concept-cards/similarity.md) - `닮음`을 비례식과 확대·축소 관계 중심으로 정리한다.
- [syntheses/concept-cards/parallelogram.md](./syntheses/concept-cards/parallelogram.md) - `평행사변형`을 대변과 대각선의 성질, 증명 연결 중심으로 정리한다.
- [syntheses/concept-cards/circle.md](./syntheses/concept-cards/circle.md) - `원`을 중심각, 원주각, 거리 조건 중심으로 정리한다.
- [syntheses/concept-cards/pythagorean-theorem.md](./syntheses/concept-cards/pythagorean-theorem.md) - `피타고라스 정리`를 거리와 제곱 관계의 브리지로 정리한다.
- [syntheses/concept-cards/sets-and-propositions.md](./syntheses/concept-cards/sets-and-propositions.md) - `집합과 명제`를 논리와 조건 구조의 언어로 정리한다.
- [syntheses/concept-cards/equations-of-geometric-figures.md](./syntheses/concept-cards/equations-of-geometric-figures.md) - `도형의 방정식`을 거리 공식, 원의 방정식, 좌표기하 브리지 중심으로 정리한다.
- [syntheses/concept-cards/equation-of-a-line.md](./syntheses/concept-cards/equation-of-a-line.md) - `직선의 방정식`을 점기울기형, 일반형, 교점 해석 중심으로 정리한다.
- [syntheses/concept-cards/equation-of-a-circle.md](./syntheses/concept-cards/equation-of-a-circle.md) - `원의 방정식`을 중심-반지름 조건과 전개형 해석 중심으로 정리한다.
- [syntheses/concept-cards/conic-sections.md](./syntheses/concept-cards/conic-sections.md) - `이차곡선`을 포물선·타원·쌍곡선의 거리 조건 중심으로 정리한다.
- [syntheses/concept-cards/parabola.md](./syntheses/concept-cards/parabola.md) - `포물선`을 초점, 준선, 표준형 중심으로 정리한다.
- [syntheses/concept-cards/ellipse.md](./syntheses/concept-cards/ellipse.md) - `타원`을 두 초점 거리합과 표준형 중심으로 정리한다.
- [syntheses/concept-cards/hyperbola.md](./syntheses/concept-cards/hyperbola.md) - `쌍곡선`을 거리차 조건과 점근선 감각 중심으로 정리한다.
- [syntheses/concept-cards/vectors.md](./syntheses/concept-cards/vectors.md) - `벡터`를 방향, 성분, 내적 중심으로 정리한다.
- [syntheses/concept-cards/complex-plane.md](./syntheses/concept-cards/complex-plane.md) - `복소수평면`을 복소수와 벡터·좌표의 연결 브리지로 정리한다.
- [syntheses/concept-cards/spatial-coordinates.md](./syntheses/concept-cards/spatial-coordinates.md) - `공간좌표`를 3차원 위치와 거리 계산 중심으로 정리한다.
- [syntheses/concept-cards/projection-and-orthographic-views.md](./syntheses/concept-cards/projection-and-orthographic-views.md) - `투영과 투상도`를 3차원 시각화와 평면 표현의 브리지로 정리한다.

### Data And Probability

- [syntheses/concept-cards/data-organization.md](./syntheses/concept-cards/data-organization.md) - `자료 정리`를 표, 그래프, 대표값 읽기 중심으로 정리한다.
- [syntheses/concept-cards/variance.md](./syntheses/concept-cards/variance.md) - `분산`을 산포의 제곱 평균과 흩어짐 해석 중심으로 정리한다.
- [syntheses/concept-cards/standard-deviation.md](./syntheses/concept-cards/standard-deviation.md) - `표준편차`를 분산의 제곱근과 단위 복원 감각 중심으로 정리한다.
- [syntheses/concept-cards/counting-principles.md](./syntheses/concept-cards/counting-principles.md) - `경우의 수`를 세기 원리, 순열, 조합, 확률 연결 중심으로 정리한다.
- [syntheses/concept-cards/sampling.md](./syntheses/concept-cards/sampling.md) - `표본조사`를 모집단, 표본, 대표성 중심으로 정리한다.
- [syntheses/concept-cards/correlation-and-scatter-plots.md](./syntheses/concept-cards/correlation-and-scatter-plots.md) - `상관과 산점도`를 이변량 데이터 해석 중심으로 정리한다.
- [syntheses/concept-cards/correlation-coefficient.md](./syntheses/concept-cards/correlation-coefficient.md) - `상관계수`를 선형 상관의 수치 요약 중심으로 정리한다.
- [syntheses/concept-cards/conditional-probability.md](./syntheses/concept-cards/conditional-probability.md) - `조건부확률`을 표본공간 축소, 독립성, 확률 갱신 중심으로 정리한다.
- [syntheses/concept-cards/random-variable.md](./syntheses/concept-cards/random-variable.md) - `확률변수`를 기대값, 분산, 분포의 출발점으로 정리한다.
- [syntheses/concept-cards/probability-distribution.md](./syntheses/concept-cards/probability-distribution.md) - `확률분포`를 확률변수, 기대값, 분산과 함께 정리한다.
- [syntheses/concept-cards/statistical-inference.md](./syntheses/concept-cards/statistical-inference.md) - `통계적 추정`을 표본, 대표성, 추정 해석, 국가별 배치와 함께 정리한다.
- [syntheses/concept-cards/confidence-interval.md](./syntheses/concept-cards/confidence-interval.md) - `신뢰구간`을 추정값의 범위 해석과 오차 감각 중심으로 정리한다.
- [syntheses/concept-cards/hypothesis-testing.md](./syntheses/concept-cards/hypothesis-testing.md) - `가설검정`을 귀무가설, 대립가설, 판단 규칙 중심으로 정리한다.
- [syntheses/concept-cards/normal-distribution.md](./syntheses/concept-cards/normal-distribution.md) - `정규분포`를 대표적인 분포 모델과 국가별 배치 중심으로 정리한다.
- [syntheses/concept-cards/regression.md](./syntheses/concept-cards/regression.md) - `회귀`를 산점도, 추세선, 예측, 데이터 모델링 중심으로 정리한다.
- [syntheses/concept-cards/practical-statistics.md](./syntheses/concept-cards/practical-statistics.md) - `실용 통계`를 데이터 설계, 시각화, 해석 실무 중심으로 정리한다.

### Course, Track, And Applied Nodes

- [syntheses/concept-cards/common-math-1.md](./syntheses/concept-cards/common-math-1.md) - 한국 `공통수학1`을 다항식, 경우의 수, 행렬 허브로 정리한다.
- [syntheses/concept-cards/common-math-2.md](./syntheses/concept-cards/common-math-2.md) - 한국 `공통수학2`를 좌표기하, 명제, 함수 허브로 정리한다.
- [syntheses/concept-cards/korean-algebra-course.md](./syntheses/concept-cards/korean-algebra-course.md) - 한국 `대수` 과목을 지수·로그, 삼각함수, 수열, 귀납법 허브로 정리한다.
- [syntheses/concept-cards/calculus-1.md](./syntheses/concept-cards/calculus-1.md) - 한국 `미적분I`를 극한·연속·미분·적분 허브로 정리한다.
- [syntheses/concept-cards/probability-and-statistics-course.md](./syntheses/concept-cards/probability-and-statistics-course.md) - 한국 `확률과 통계` 과목을 분포와 추정 허브로 정리한다.
- [syntheses/concept-cards/calculus-2.md](./syntheses/concept-cards/calculus-2.md) - 한국 `미적분II`를 고급 미적분 허브로 정리한다.
- [syntheses/concept-cards/korean-geometry-course.md](./syntheses/concept-cards/korean-geometry-course.md) - 한국 `기하`를 이차곡선·공간좌표·벡터 허브로 정리한다.
- [syntheses/concept-cards/basic-math-1-and-2.md](./syntheses/concept-cards/basic-math-1-and-2.md) - 한국 `기본수학1·2`를 기초 공통 경로로 정리한다.
- [syntheses/concept-cards/japan-middle-1.md](./syntheses/concept-cards/japan-middle-1.md) - 일본 `중1`을 정수, 비례·반비례, 작도, 자료 읽기 허브로 정리한다.
- [syntheses/concept-cards/japan-middle-2.md](./syntheses/concept-cards/japan-middle-2.md) - 일본 `중2`를 연립방정식, 일차함수, 합동과 증명 허브로 정리한다.
- [syntheses/concept-cards/japan-middle-3.md](./syntheses/concept-cards/japan-middle-3.md) - 일본 `중3`을 이차, 닮음, 표본조사 허브로 정리한다.
- [syntheses/concept-cards/japan-math-1.md](./syntheses/concept-cards/japan-math-1.md) - 일본 `수학I`를 중등-고등 브리지 과목으로 정리한다.
- [syntheses/concept-cards/japan-math-a.md](./syntheses/concept-cards/japan-math-a.md) - 일본 `수학A`를 경우의 수와 확률 허브로 정리한다.
- [syntheses/concept-cards/japan-math-2.md](./syntheses/concept-cards/japan-math-2.md) - 일본 `수학II`를 함수와 미적분 전개 허브로 정리한다.
- [syntheses/concept-cards/japan-math-b.md](./syntheses/concept-cards/japan-math-b.md) - 일본 `수학B`를 수열과 통계적 추측 허브로 정리한다.
- [syntheses/concept-cards/japan-math-3.md](./syntheses/concept-cards/japan-math-3.md) - 일본 `수학III`를 고급 미적분 허브로 정리한다.
- [syntheses/concept-cards/japan-math-c.md](./syntheses/concept-cards/japan-math-c.md) - 일본 `수학C`를 벡터·복소수평면 허브로 정리한다.
- [syntheses/concept-cards/china-grade-7-semester-1.md](./syntheses/concept-cards/china-grade-7-semester-1.md) - 중국 `7학년 1학기`를 유리수·일차방정식 허브로 정리한다.
- [syntheses/concept-cards/china-grade-7-semester-2.md](./syntheses/concept-cards/china-grade-7-semester-2.md) - 중국 `7학년 2학기`를 좌표, 연립방정식, 부등식 허브로 정리한다.
- [syntheses/concept-cards/china-grade-8-semester-1.md](./syntheses/concept-cards/china-grade-8-semester-1.md) - 중국 `8학년 1학기`를 합동, 축대칭, 인수분해 허브로 정리한다.
- [syntheses/concept-cards/china-grade-8-semester-2.md](./syntheses/concept-cards/china-grade-8-semester-2.md) - 중국 `8학년 2학기`를 일차함수와 데이터 분석 허브로 정리한다.
- [syntheses/concept-cards/china-grade-9-semester-1.md](./syntheses/concept-cards/china-grade-9-semester-1.md) - 중국 `9학년 1학기`를 이차와 원, 확률 기초 허브로 정리한다.
- [syntheses/concept-cards/china-grade-9-semester-2.md](./syntheses/concept-cards/china-grade-9-semester-2.md) - 중국 `9학년 2학기`를 닮음과 예각 삼각함수 허브로 정리한다.
- [syntheses/concept-cards/china-required-curriculum.md](./syntheses/concept-cards/china-required-curriculum.md) - 중국 `필수 과정`을 함수·기하와 대수·확통 허브로 정리한다.
- [syntheses/concept-cards/china-elective-required-curriculum.md](./syntheses/concept-cards/china-elective-required-curriculum.md) - 중국 `선택필수 과정`을 도함수·해석기하·확률변수 허브로 정리한다.
- [syntheses/concept-cards/china-elective-concept-clusters.md](./syntheses/concept-cards/china-elective-concept-clusters.md) - 중국 `선택 과정의 개념군`을 메타 허브로 정리한다.
- [syntheses/concept-cards/us-grade-6.md](./syntheses/concept-cards/us-grade-6.md) - 미국 `Grade 6`을 비율, 유리수, 식과 기초 통계 허브로 정리한다.
- [syntheses/concept-cards/us-grade-7.md](./syntheses/concept-cards/us-grade-7.md) - 미국 `Grade 7`을 비례, 일차식, 표본·확률 허브로 정리한다.
- [syntheses/concept-cards/us-grade-8.md](./syntheses/concept-cards/us-grade-8.md) - 미국 `Grade 8`을 함수, 닮음, 이변량 데이터 허브로 정리한다.
- [syntheses/concept-cards/algebra-1.md](./syntheses/concept-cards/algebra-1.md) - 미국 `Algebra I`를 선형·이차 대수 입문 허브로 정리한다.
- [syntheses/concept-cards/geometry-course.md](./syntheses/concept-cards/geometry-course.md) - 미국 `Geometry`를 증명·변환·삼각비 허브로 정리한다.
- [syntheses/concept-cards/algebra-2.md](./syntheses/concept-cards/algebra-2.md) - 미국 `Algebra II`를 고급 함수와 수열 허브로 정리한다.
- [syntheses/concept-cards/precalculus.md](./syntheses/concept-cards/precalculus.md) - 미국 `Precalculus`를 미적분 직전 함수 통합 허브로 정리한다.
- [syntheses/concept-cards/calculus-course.md](./syntheses/concept-cards/calculus-course.md) - 미국 `Calculus`를 상위 미적분 course-track 허브로 정리한다.
- [syntheses/concept-cards/ap-calculus.md](./syntheses/concept-cards/ap-calculus.md) - 미국 `AP Calculus`를 외부평가 연동 미적분 wrapper로 정리한다.
- [syntheses/concept-cards/statistics-and-data-analysis.md](./syntheses/concept-cards/statistics-and-data-analysis.md) - 미국 `Statistics and Data Analysis`를 데이터 해석 허브로 정리한다.
- [syntheses/concept-cards/ap-statistics.md](./syntheses/concept-cards/ap-statistics.md) - 미국 `AP Statistics`를 고급 통계 이수와 외부 평가 wrapper로 정리한다.
- [syntheses/concept-cards/ib-mathematics.md](./syntheses/concept-cards/ib-mathematics.md) - 미국 맥락의 `IB Mathematics`를 국제과정 wrapper로 정리한다.
- [syntheses/concept-cards/advanced-tracks.md](./syntheses/concept-cards/advanced-tracks.md) - 미국 `Advanced Tracks`를 Honors/AP/IB 메타 경로로 정리한다.
- [syntheses/concept-cards/mathematical-modeling-and-inquiry.md](./syntheses/concept-cards/mathematical-modeling-and-inquiry.md) - `수학적 모델링·수학 탐구`를 현실 문제 번역과 해석 활동 중심으로 정리한다.
- [syntheses/concept-cards/economics-math.md](./syntheses/concept-cards/economics-math.md) - `경제 수학`을 이자, 현재가치, 경제 모델링 중심으로 정리한다.
- [syntheses/concept-cards/mathematics-for-ai.md](./syntheses/concept-cards/mathematics-for-ai.md) - `인공지능 수학`을 벡터화, 손실함수, 예측 중심으로 정리한다.
- [syntheses/concept-cards/workplace-math.md](./syntheses/concept-cards/workplace-math.md) - `직무 수학`을 작업량, 단위, 실무형 모델링 중심으로 정리한다.
- [syntheses/concept-cards/mathematics-and-culture.md](./syntheses/concept-cards/mathematics-and-culture.md) - `수학과 문화`를 역사, 예술, 과학기술 연결 중심으로 정리한다.
- [syntheses/concept-cards/integrated-mathematics.md](./syntheses/concept-cards/integrated-mathematics.md) - 미국형 `Integrated Mathematics` course-track을 개념 순환 구조로 정리한다.

## Source Notes

- [source-notes/hands-on-math-agent-session-guide.md](./source-notes/hands-on-math-agent-session-guide.md) - 핸즈온 세션 가이드의 목적과 운영 포인트 요약
- [source-notes/prompt-pack-readme.md](./source-notes/prompt-pack-readme.md) - 프롬프트 팩의 사용 범위와 순서 요약
- [source-notes/prompt-pack-wiki-enricher-extension.md](./source-notes/prompt-pack-wiki-enricher-extension.md) - `wiki enricher` 확장 에이전트와 `learning_context` 흐름 요약
- [source-notes/llm-wiki-usage-guide.md](./source-notes/llm-wiki-usage-guide.md) - clone 이후 `Gemini CLI`와 `Antigravity`에서 위키를 운영하는 흐름 요약
- [source-notes/gemini-cli-math-agent-step-by-step.md](./source-notes/gemini-cli-math-agent-step-by-step.md) - Gemini CLI에서 수학 해설 프로토타입을 단계별로 만드는 프롬프트 흐름 요약
- [source-notes/antigravity-math-agent-step-by-step.md](./source-notes/antigravity-math-agent-step-by-step.md) - Antigravity에서 수학 해설 프로토타입을 단계별로 만드는 프롬프트 흐름 요약
- [source-notes/antigravity-workspace-rules-draft.md](./source-notes/antigravity-workspace-rules-draft.md) - Antigravity workspace rules에 붙여넣을 수 있는 규칙 초안 요약
- [source-notes/antigravity-workflow-usage-guide.md](./source-notes/antigravity-workflow-usage-guide.md) - Antigravity workflow를 어떤 역할 분리와 순서로 쓰는 것이 적절한지 요약
- [source-notes/workspace-distribution-guide.md](./source-notes/workspace-distribution-guide.md) - 공유 workspace에서 공개용 자산, starter kit, facilitator 전용 메모를 어떻게 분리할지 요약
- [source-notes/math-agent-system-spec.md](./source-notes/math-agent-system-spec.md) - 멀티 에이전트 시스템 명세의 핵심 구조 요약
- [source-notes/math-agent-knowledge-plan.md](./source-notes/math-agent-knowledge-plan.md) - 지식 데이터 수집 방향과 교육과정 비교 요약
- [source-notes/internal-knowledge-base-build.md](./source-notes/internal-knowledge-base-build.md) - 내부 지식베이스 구축 규칙과 타입 정의 요약
- [source-notes/math-curriculum-research-readme.md](./source-notes/math-curriculum-research-readme.md) - 수학 교육과정 비교 폴더의 목적과 문서군 구조 요약
- [source-notes/math-curriculum-research-korea.md](./source-notes/math-curriculum-research-korea.md) - 한국 교육과정/입시 구조 문서의 핵심 주장 요약
- [source-notes/math-curriculum-research-comparison.md](./source-notes/math-curriculum-research-comparison.md) - 4개국 비교 문서의 핵심 비교축 요약
- [source-notes/math-curriculum-research-japan.md](./source-notes/math-curriculum-research-japan.md) - 일본 교육과정의 학년 위계와 고교 과목 구조 요약
- [source-notes/math-curriculum-research-china.md](./source-notes/math-curriculum-research-china.md) - 중국 교육과정의 빠른 계통화와 고교 필수/선택필수 구조 요약
- [source-notes/math-curriculum-research-us.md](./source-notes/math-curriculum-research-us.md) - 미국의 course track 중심 수학 경로와 중학교 핵심 개념 배치 요약
- [source-notes/math-concept-encyclopedia-readme.md](./source-notes/math-concept-encyclopedia-readme.md) - 개념 백과사전 폴더의 역할과 산출물 구조 요약
- [source-notes/math-concept-encyclopedia-korea.md](./source-notes/math-concept-encyclopedia-korea.md) - 한국 개념 백과사전의 개념 카드 활용 포인트 요약
- [source-notes/math-concept-encyclopedia-japan.md](./source-notes/math-concept-encyclopedia-japan.md) - 일본 개념 백과사전의 함수 위계와 삼각비/확률 배치 요약
- [source-notes/math-concept-encyclopedia-china.md](./source-notes/math-concept-encyclopedia-china.md) - 중국 개념 백과사전의 방정식/함수/확통 계통 요약
- [source-notes/math-concept-encyclopedia-us.md](./source-notes/math-concept-encyclopedia-us.md) - 미국 개념 백과사전의 경로형 개념 배치와 통계 과목 구조 요약
- [source-notes/math-concept-encyclopedia-multilingual-glossary.md](./source-notes/math-concept-encyclopedia-multilingual-glossary.md) - 다국어 용어집의 alias 및 대표 수식 활용 포인트 요약
- [source-notes/math-concept-encyclopedia-formula-examples.md](./source-notes/math-concept-encyclopedia-formula-examples.md) - 수식과 대표 예제 모음의 활용 포인트 요약
- [source-notes/math-concept-encyclopedia-unit-practice-book.md](./source-notes/math-concept-encyclopedia-unit-practice-book.md) - 단원별 연습문제 자산의 활용 포인트 요약
- [source-notes/math-concept-encyclopedia-comparative-problem-book.md](./source-notes/math-concept-encyclopedia-comparative-problem-book.md) - 나라별 동일 개념 비교 문제집의 활용 포인트 요약

## Query Notes

- [queries/2026-04-06-math-curriculum-graph.md](./queries/2026-04-06-math-curriculum-graph.md) - 중1~고3 수학을 `학년 허브 + 계통 축 + 개념 카드 + 평가 레이어`로 정리하는 그래프 설계안
- [queries/math-curriculum-graph/middle-1-hub.md](./queries/math-curriculum-graph/middle-1-hub.md) - `중1`을 시작 허브로 다루는 샘플 페이지
- [queries/math-curriculum-graph/middle-2-hub.md](./queries/math-curriculum-graph/middle-2-hub.md) - `중2`를 선형 관계와 논증의 확장 구간으로 다루는 샘플 페이지
- [queries/math-curriculum-graph/middle-3-hub.md](./queries/math-curriculum-graph/middle-3-hub.md) - `중3`를 고등수학 전이 허브로 다루는 샘플 페이지
- [queries/math-curriculum-graph/high-1-hub.md](./queries/math-curriculum-graph/high-1-hub.md) - `고1` 공통수학 기반 허브를 다루는 샘플 페이지
- [queries/math-curriculum-graph/high-2-hub.md](./queries/math-curriculum-graph/high-2-hub.md) - `고2` 과목 분기 허브를 다루는 샘플 페이지
- [queries/math-curriculum-graph/high-3-hub.md](./queries/math-curriculum-graph/high-3-hub.md) - `고3` 심화·응용 허브를 다루는 샘플 페이지
- [queries/math-curriculum-graph/functions-strand.md](./queries/math-curriculum-graph/functions-strand.md) - 함수 계통을 학년 횡단 그래프로 묶는 샘플 페이지
- [queries/math-curriculum-graph/geometry-strand.md](./queries/math-curriculum-graph/geometry-strand.md) - 기하 계통을 학년 횡단 그래프로 묶는 샘플 페이지
- [queries/math-curriculum-graph/statistics-and-probability-strand.md](./queries/math-curriculum-graph/statistics-and-probability-strand.md) - 확률·통계 계통을 학년 횡단 그래프로 묶는 샘플 페이지
- [queries/math-curriculum-graph/linear-function.md](./queries/math-curriculum-graph/linear-function.md) - `일차함수` 개념 카드를 중심 노드로 다루는 샘플 페이지
- [queries/math-curriculum-graph/quadratic-function.md](./queries/math-curriculum-graph/quadratic-function.md) - `이차함수` 개념 카드를 브리지 노드로 다루는 샘플 페이지
- [queries/math-curriculum-graph/trigonometric-ratio.md](./queries/math-curriculum-graph/trigonometric-ratio.md) - `삼각비` 개념 카드를 중등-고등 브리지 노드로 다루는 샘플 페이지
- [queries/math-curriculum-graph/probability-and-statistics-course.md](./queries/math-curriculum-graph/probability-and-statistics-course.md) - `확률과 통계` 과목 노드를 다루는 샘플 페이지
- [queries/math-curriculum-graph/csat-2028.md](./queries/math-curriculum-graph/csat-2028.md) - `2028학년도 수능`을 평가 레이어 노드로 다루는 샘플 페이지
- [queries/2026-04-06-research-folder-coverage-audit.md](./queries/2026-04-06-research-folder-coverage-audit.md) - 두 개의 리서치 폴더가 persistent layer에 얼마나 반영되었는지와 남은 선택형 심화 작업을 점검한 감사 노트
- [queries/2026-04-09-remaining-curriculum-backlog.md](./queries/2026-04-09-remaining-curriculum-backlog.md) - docs 기반 대표 backlog가 사실상 닫힌 뒤, 남은 선택형 심화 후보를 다시 정리한 노트

## Query And Lint Destinations

- `wiki/queries/` - 질의에서 나온 분석을 저장할 위치
- `wiki/lint/` - 위키 health-check 결과를 저장할 위치

## Raw Source Roots

- `docs/`
- `docs/prompt-pack/`
