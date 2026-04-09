---
title: 극한
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
  - calculus
  - limit
---

# 극한

## Summary

극한은 값이 어떤 대상에 점점 가까워질 때 그 끝 behavior를 읽는 개념이다. 무한 과정을 다루는 언어라서 미분과 적분, 연속성 전체의 출발점이 된다.

## Definition

- 변수나 수열의 항이 어떤 값에 한없이 가까워질 때, 그 경향을 읽는 개념을 극한이라 한다.
- 함수의 극한은 `x=a에서의 함수값` 그 자체보다 `x가 a에 가까워질 때의 함수값`을 본다.
- 대표 표기는
  $$
  \lim_{x\to a}f(x),\qquad \lim_{n\to\infty}a_n
  $$
  이다.

## Key Ideas

- 한없이 가까워짐
  - 극한은 실제로 도착했는지보다 `얼마나 가까워질 수 있는가`를 본다.
- 좌극한과 우극한
  - 왼쪽과 오른쪽에서 가까워지는 값이 같아야 함수의 극한이 존재한다.
- 함수의 극한과 수열의 극한
  - 둘 다 `무한 과정의 끝 behavior`를 읽지만, 입력이 연속적으로 움직이느냐 항 번호가 커지느냐가 다르다.
- 극한과 연속성
  - [continuity.md](./continuity.md)은 `극한값`과 `실제 함수값`이 일치하는지 보는 개념이다.
- 극한과 변화율
  - [differentiation.md](./differentiation.md)의 순간변화율도 결국 차분비의 극한으로 정의된다.

## Deep Dive

- 극한은 `값이 없던 곳을 억지로 채우는 계산`이 아니다.
- 예를 들어 함수가 어떤 점에서 정의되지 않아도, 그 점 근처에서 가까워지는 경향은 충분히 읽을 수 있다.
- 그래서 `극한이 존재한다`와 `그 점에서 함수값이 존재한다`는 서로 다른 말이다.

## Theorems and Properties

- 극한의 사칙연산
  - $\lim_{x\to a}f(x)=L$, $\lim_{x\to a}g(x)=M$이면
    $$
    \lim_{x\to a}(f(x)\pm g(x))=L\pm M,\qquad \lim_{x\to a}f(x)g(x)=LM
    $$
    이다.
- 다항함수의 극한
  - 다항함수는 모든 점에서 `대입`으로 극한을 계산할 수 있다.
- 샌드위치 정리
  - 어떤 함수가 두 함수 사이에 끼여 있고, 양옆 함수의 극한이 같으면 가운데 함수의 극한도 같다.
- 수열 극한의 대표 성질
  - $\lim_{n\to\infty}\frac{1}{n}=0$은 많은 극한 직관의 출발점이다.

## Proof Sketch

- `증명 스케치 (추론)`:
- 수열 $\frac1n$을 생각하자.
- $n$이 커질수록 $\frac1n$은 양수이면서 계속 작아진다.
- 원하는 만큼 작은 양수 $\varepsilon$를 정해도, $n>\frac1\varepsilon$가 되면
  $$
  \frac1n<\varepsilon
  $$
  가 된다.
- 즉 $\frac1n$은 0에 한없이 가까워질 수 있으므로
  $$
  \lim_{n\to\infty}\frac1n=0
  $$
  이다.

## Worked Examples

- 예제 1: 다항함수의 극한
  - $\lim_{x\to 2}(x^2-3x+1)$을 구하자.
  - 다항함수이므로 바로 대입하면
    $$
    2^2-3\cdot 2+1=-1
    $$
    이다.
  - 따라서 극한값은 $-1$이다.
- 예제 2: 약분 후 극한
  - $\lim_{x\to 1}\frac{x^2-1}{x-1}$을 구하자.
  - 분자를 인수분해하면
    $$
    \frac{(x-1)(x+1)}{x-1}=x+1\quad(x\neq 1)
    $$
    이다.
  - 따라서
    $$
    \lim_{x\to 1}\frac{x^2-1}{x-1}=\lim_{x\to 1}(x+1)=2
    $$
    이다.
- 예제 3: 수열 극한
  - $\lim_{n\to\infty}\frac{3n+1}{n}$은
    $$
    \frac{3n+1}{n}=3+\frac1n
    $$
    이므로
    $$
    \lim_{n\to\infty}\frac{3n+1}{n}=3
    $$
    이다.

## Common Pitfalls

- `x=a`를 바로 대입했을 때 식이 안 된다고 해서 극한이 없다고 결론내리면 안 된다.
- 좌극한과 우극한이 다른데도 극한이 있다고 착각하기 쉽다.
- `극한값`과 `함수값`을 같은 것으로 보면 [continuity.md](./continuity.md)에서 바로 헷갈린다.
- 분모가 0이 되는 식에서 약분 없이 바로 계산하려 하면 오류가 난다.

## Curriculum Placement

- 한국 대표 배치에서는 `미적분I`와 `미적분II`의 핵심 시작점이다.
- 국가별 배치 스냅샷
  - 한국: `함수의 극한과 연속`, `수열의 극한`으로 나누어 다룬다.
  - 일본: `수학III`의 극한이 대표 단원이다.
  - 중국: `수열과 도함수` 및 고등 미적분 전개에서 극한 감각이 사용된다.
  - 미국: Calculus의 맨 앞에서 limits and continuity로 다룬다.
- 표현 메모
  - 다국어 용어: `limit`, `極限`, `极限`

## Connections

- 선수 개념은 [sequences.md](./sequences.md), [square-roots-and-real-numbers.md](./square-roots-and-real-numbers.md), [function-foundations.md](./function-foundations.md)다.
- 같은 축의 인접 개념으로는 [continuity.md](./continuity.md), [differentiation.md](./differentiation.md), [integration.md](./integration.md)이 있다.
- 과목 허브로는 [calculus-1.md](./calculus-1.md), [calculus-2.md](./calculus-2.md), [calculus-course.md](./calculus-course.md)를 본다.
- 계통 허브로는 [../functions-strand.md](../functions-strand.md), [../high-2-hub.md](../high-2-hub.md), [../high-3-hub.md](../high-3-hub.md)를 함께 본다.

## Open Questions

- `수열의 극한`과 `함수의 극한`을 별도 심화 카드로 더 분리할지 기준이 필요하다.

## Sources

- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-concept-encyclopedia/japan.md`
- `docs/math-concept-encyclopedia/china.md`
- `docs/math-concept-encyclopedia/us.md`
- `docs/math-concept-encyclopedia/multilingual-glossary.md`
