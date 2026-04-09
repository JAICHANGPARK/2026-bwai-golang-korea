---
title: 삼각비
type: synthesis
status: active
updated: 2026-04-10
card_role: concept
source_docs:
  - docs/math-concept-encyclopedia/korea.md
  - docs/math-concept-encyclopedia/japan.md
  - docs/math-concept-encyclopedia/china.md
  - docs/math-concept-encyclopedia/us.md
  - docs/math-curriculum-research/korea.md
  - docs/math-curriculum-research/japan.md
  - docs/math-curriculum-research/china.md
  - docs/math-curriculum-research/us.md
  - docs/math-concept-encyclopedia/multilingual-glossary.md
  - docs/math-concept-encyclopedia/formula-examples.md
  - docs/math-concept-encyclopedia/unit-practice-book.md
  - docs/math-concept-encyclopedia/comparative-problem-book.md
tags:
  - concept
  - geometry
  - trigonometry
---

# 삼각비

## Summary

삼각비는 직각삼각형에서 각과 변의 길이 비를 연결하는 개념이다. 중학교에서는 닮음과 피타고라스 정리 위에서 배우고, 고등학교에서는 삼각함수와 기하, 벡터로 확장되는 핵심 브리지 역할을 한다.

## Key Points

- 정의
  - 직각삼각형에서 예각 `θ`에 대해 변의 길이 비를 이용해 $\sin\theta$, $\cos\theta$, $\tan\theta$를 정의한다.
- 핵심 개념
  - 닮음비
  - 대응변
  - 직각삼각형
  - 피타고라스 정리
  - 사인, 코사인, 탄젠트
- 대표 수식
  - $a^2+b^2=c^2$
  - $\sin\theta=\frac{\text{높이}}{\text{빗변}}$
  - $\cos\theta=\frac{\text{밑변}}{\text{빗변}}$
  - $\tan\theta=\frac{\text{높이}}{\text{밑변}}$
- 핵심 해석
  - 삼각비는 `길이 자체`가 아니라 `길이의 비`를 다루므로, 삼각형이 커지거나 작아져도 값이 변하지 않는다.
  - 그래서 삼각비는 `크기`가 아니라 `각의 모양`을 나타내는 수가 된다.
- 교육과정 배치
  - 한국 대표 배치에서는 `중3`의 `닮음·원·피타고라스·삼각비` 묶음에 놓인다.
  - 이후 고등학교 `삼각함수`, `기하`, `벡터`로 이어진다.
- 국가별 배치 스냅샷
  - 한국: `중3`에서 닮음과 피타고라스 정리 위에 `삼각비`가 놓인다.
  - 일본: 중학교에서는 `닮음·피타고라스`를 준비 개념으로 두고, 고등학교 `수학I`에서 `삼각비`를 본격 도입한다.
  - 중국: `9학년 2학기`에 `예각 삼각함수`로 도입된다.
  - 미국: 보통 `Geometry`에서 `trigonometric ratios`를 다루고, 이후 `Precalculus`에서 삼각함수로 확장한다.
- 표현과 문제 감각
  - 다국어 용어: `trigonometric ratio`, `三角比`, `锐角三角函数`
  - 비교 문제집 기준으로 중국은 중학에서 `锐角三角函数` 표현이 두드러지고, 일본은 `삼각비 -> 삼각함수` 위계가 선명하며, 미국은 Geometry/Precalculus 경로 속에서 연결된다.

## Deep Dive

### 왜 삼각비는 각만으로 정해지는가

- 같은 각 `θ`를 가지는 두 직각삼각형은 [similarity.md](./similarity.md)에 의해 닮음이다.
- 닮음인 두 삼각형에서는 대응변의 비가 모두 같으므로
  $$
  \frac{\text{높이}}{\text{빗변}},\quad \frac{\text{밑변}}{\text{빗변}},\quad \frac{\text{높이}}{\text{밑변}}
  $$
  의 값은 삼각형의 크기와 무관하다.
- 따라서 이 값들은 `각 θ의 성질`이라고 볼 수 있고, 이를 각각 $\sin\theta$, $\cos\theta$, $\tan\theta$로 정의한다.

### 핵심 정리 1: $\tan\theta=\frac{\sin\theta}{\cos\theta}$

- 증명
  - 정의에 의해
    $$
    \sin\theta=\frac{\text{높이}}{\text{빗변}},\qquad \cos\theta=\frac{\text{밑변}}{\text{빗변}}
    $$
    이다.
  - 따라서
    $$
    \frac{\sin\theta}{\cos\theta}
    =
    \frac{\frac{\text{높이}}{\text{빗변}}}{\frac{\text{밑변}}{\text{빗변}}}
    =
    \frac{\text{높이}}{\text{밑변}}
    =
    \tan\theta
    $$
    이다.

### 핵심 정리 2: $\sin^2\theta+\cos^2\theta=1$

- 증명
  - 직각삼각형의 높이, 밑변, 빗변을 각각 $a,b,c$라 하면
    $$
    \sin\theta=\frac{a}{c},\qquad \cos\theta=\frac{b}{c}
    $$
    이다.
  - 따라서
    $$
    \sin^2\theta+\cos^2\theta
    =
    \frac{a^2}{c^2}+\frac{b^2}{c^2}
    =
    \frac{a^2+b^2}{c^2}
    $$
    이다.
  - 이제 [pythagorean-theorem.md](./pythagorean-theorem.md)에 의해 $a^2+b^2=c^2$이므로
    $$
    \sin^2\theta+\cos^2\theta=1
    $$
    을 얻는다.

## Worked Examples

### 예제 1: 삼각비 직접 계산

- 직각삼각형에서 높이 `3`, 밑변 `4`, 빗변 `5`이면
  $$
  \sin\theta=\frac{3}{5},\qquad \cos\theta=\frac{4}{5},\qquad \tan\theta=\frac{3}{4}
  $$
  이다.

### 예제 2: 비를 이용해 길이 구하기

- $\sin\theta=\frac{3}{5}$이고 빗변의 길이가 `10`이면
  $$
  \frac{\text{높이}}{10}=\frac{3}{5}
  $$
  이다.
- 따라서 높이는
  $$
  10\cdot \frac{3}{5}=6
  $$
  이다.

## Common Pitfalls

- `sin`과 `cos`의 분모는 모두 `빗변`이라는 점을 자주 놓친다.
- `tan`은 빗변이 아니라 `밑변`을 분모로 둔다는 점을 혼동하기 쉽다.
- 삼각비는 `길이`가 아니라 `비`라는 사실을 잊으면, 삼각형이 커지면 값도 커진다고 오해하기 쉽다.
- 중학교 삼각비와 고등학교 [trigonometric-function.md](./trigonometric-function.md)를 같은 단계 개념으로 보면 확장 구조가 잘 안 보인다.

## Connections

- 선수 개념은 [similarity.md](./similarity.md), [pythagorean-theorem.md](./pythagorean-theorem.md), [circle.md](./circle.md)다.
- 같은 축의 인접 개념으로는 [trigonometric-function.md](./trigonometric-function.md), [vectors.md](./vectors.md)가 있다.
- 다음 개념으로는 [korean-geometry-course.md](./korean-geometry-course.md), [equations-of-geometric-figures.md](./equations-of-geometric-figures.md), [trigonometric-function.md](./trigonometric-function.md)가 이어진다.
- 학년 허브에서는 [middle-3-hub.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/syntheses/middle-3-hub.md), [high-3-hub.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/syntheses/high-3-hub.md)와 연결된다.
- 계통 허브에서는 [geometry-strand.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/syntheses/geometry-strand.md)와 [functions-strand.md](/Users/jaichang/Documents/GitHub/2026-bwai-golang-korea/wiki/syntheses/functions-strand.md)를 함께 본다.

## Open Questions

- `삼각비`와 [trigonometric-function.md](./trigonometric-function.md)의 경계를 어디까지 분리 유지할지 기준이 더 필요하다.
- `원과 원주각`을 삼각비 카드의 선수 카드로 독립시킬지 후속 설계가 필요하다.

## Sources

- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-concept-encyclopedia/japan.md`
- `docs/math-concept-encyclopedia/china.md`
- `docs/math-concept-encyclopedia/us.md`
- `docs/math-curriculum-research/korea.md`
- `docs/math-curriculum-research/japan.md`
- `docs/math-curriculum-research/china.md`
- `docs/math-curriculum-research/us.md`
- `docs/math-concept-encyclopedia/multilingual-glossary.md`
- `docs/math-concept-encyclopedia/formula-examples.md`
- `docs/math-concept-encyclopedia/unit-practice-book.md`
- `docs/math-concept-encyclopedia/comparative-problem-book.md`
