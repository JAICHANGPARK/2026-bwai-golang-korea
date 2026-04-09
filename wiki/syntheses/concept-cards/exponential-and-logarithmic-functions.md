---
title: 지수함수와 로그함수
type: synthesis
status: active
updated: 2026-04-10
card_role: concept
source_docs:
  - docs/math-concept-encyclopedia/korea.md
  - docs/math-concept-encyclopedia/japan.md
  - docs/math-concept-encyclopedia/china.md
  - docs/math-concept-encyclopedia/us.md
  - docs/math-concept-encyclopedia/multilingual-glossary.md
tags:
  - concept
  - function
  - exponential
---

# 지수함수와 로그함수

## Summary

지수함수와 로그함수는 빠른 성장과 감소를 다루는 함수 개념이다. 지수는 반복 곱셈의 구조를, 로그는 그 구조를 역으로 푸는 언어를 제공하므로 복리, 척도, 지수 방정식, 미적분의 기초가 된다.

## Key Points

- 정의
  - $a>0$, $a\neq 1$일 때 $y=a^x$ 꼴의 함수를 지수함수라 한다.
  - $a>0$, $a\neq 1$, $x>0$일 때 $y=\log_a x$ 꼴의 함수를 로그함수라 한다.
  - 로그함수는 지수함수의 역함수다.
- 핵심 개념
  - 지수 증가
  - 지수 감소
  - 로그의 역함수성
  - 밑
  - 곱셈과 덧셈의 전환
  - 지수법칙
  - 밑변환
  - 단조성
- 대표 성질
  - 지수법칙: $a^{m+n}=a^ma^n$
  - 지수의 0승: $a^0=1$
  - 로그의 곱셈 법칙: $\log_a(xy)=\log_a x+\log_a y$
  - 역함수 법칙: $\log_a(a^x)=x$, $a^{\log_a x}=x$
  - $a>1$이면 지수함수는 증가하고, $0<a<1$이면 감소한다.
  - 로그함수와 지수함수는 서로 $y=x$에 대한 대칭 관계를 이룬다.
- 대표 수식
  - $a^{m+n}=a^ma^n$
  - $\log_a(xy)=\log_a x+\log_a y$
  - $a^{-n}=\frac{1}{a^n}$
  - $\log_a x=\frac{\ln x}{\ln a}$
- 증명 스케치
  - `증명 스케치 (추론)`:
  - $a^m$은 $a$를 $m$번 곱한 값, $a^n$은 $n$번 곱한 값이다.
  - 따라서 둘을 곱하면 총 $m+n$번 곱한 값이 되어 $a^{m+n}$이 된다.
  - 로그 공식은 이 지수법칙을 역으로 읽은 결과다. 즉 $\log_a x$는 `a를 몇 번 곱해야 x가 되는가`를 묻는 질문이다.
  - 역함수 법칙은 함수와 역함수의 정의를 그대로 적용한 것이다.
  - 지수함수와 로그함수는 서로 역함수이므로, 한쪽의 그래프를 뒤집으면 다른 쪽이 된다.

## Deep Dive

- 지수함수는 반복 곱셈을 한 번에 표현하는 함수이고, 로그함수는 그 반복 곱셈의 횟수를 묻는 함수다.
- 지수법칙은 곱셈을 지수의 덧셈으로 바꾸고, 로그법칙은 곱셈을 덧셈으로 바꾼다.
- 지수성장은 일정한 `비율`로 커지므로 초반에는 느려 보여도 나중에 급격히 커진다.
- 로그는 큰 수를 다루는 척도 압축 도구로 자주 쓰인다.

## Worked Examples

### 예제 1: 지수방정식

- $2^x=16$이면 $16=2^4$이므로
  $$
  x=4
  $$
  다.

### 예제 2: 로그 계산

- $\log_3 81-\log_3 3=4-1=3$이다.

### 예제 3: 성장 모델

- 월 복리 1%로 원금이 $P$일 때 12개월 뒤 금액은
  $$
  P(1.01)^{12}
  $$
  로 쓸 수 있다.

## Common Pitfalls

- $\log_a(x+y)=\log_a x+\log_a y$는 일반적으로 성립하지 않는다.
- $\log_a x$는 $x>0$에서만 정의된다.
- 밑이 1인 로그나 음수 밑은 학교 수학의 기본 로그함수 범위가 아니다.
- 지수함수를 `선형 증가`로 착각하면 성장 속도 해석이 틀어진다.
- 역함수 관계를 생각하지 않고 그래프를 외우면 지수와 로그를 쉽게 혼동한다.
- 교육과정 배치
  - 한국 대표 배치에서는 고등학교 `대수`의 핵심 함수 단원이다.
- 국가별 배치 스냅샷
  - 한국: `대수`에서 수열, 삼각함수와 함께 배운다.
  - 일본: `수학II`의 `지수·로그함수`에 놓인다.
  - 중국: 고등학교 필수과정 `함수` 축에서 명시적으로 다뤄진다.
  - 미국: Algebra II와 Precalculus에서 주로 다뤄진다.
- 표현과 문제 감각
  - 다국어 용어: `exponential function`, `logarithmic function`, `指数関数`, `対数関数`, `指数函数`, `对数函数`
  - 성장 문제와 척도 문제를 같은 함수 언어로 읽는 것이 중요하다.

## Connections

- 선수 개념은 [function-foundations.md](./function-foundations.md), [polynomials-and-factorization.md](./polynomials-and-factorization.md)다.
- 같은 축의 인접 개념으로는 [limits.md](./limits.md), [differentiation.md](./differentiation.md), [economics-math.md](./economics-math.md)가 있다.
- 함수 계통은 [functions-strand.md](../functions-strand.md)에서 이어진다.
- 고교 허브는 [high-2-hub.md](../high-2-hub.md), [korea-curriculum-hub.md](../korea-curriculum-hub.md), [course-track-hub.md](../course-track-hub.md)로 묶인다.
- 과목형 연결은 [algebra-2.md](./algebra-2.md), [precalculus.md](./precalculus.md), [calculus-course.md](./calculus-course.md)로 확장된다.

## Open Questions

- `로그의 밑변환`을 독립 카드로 분리할지 후속 검토가 필요하다.

## Sources

- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-concept-encyclopedia/japan.md`
- `docs/math-concept-encyclopedia/china.md`
- `docs/math-concept-encyclopedia/us.md`
- `docs/math-concept-encyclopedia/multilingual-glossary.md`
