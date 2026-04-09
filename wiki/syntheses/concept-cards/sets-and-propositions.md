---
title: 집합과 명제
type: synthesis
status: active
updated: 2026-04-10
card_role: concept
source_docs:
  - docs/math-concept-encyclopedia/korea.md
  - docs/math-concept-encyclopedia/china.md
  - docs/math-concept-encyclopedia/multilingual-glossary.md
tags:
  - concept
  - logic
  - set
---

# 집합과 명제

## Summary

집합과 명제는 수학적 대상을 모임과 조건의 언어로 정리하는 개념이다. 계산 중심 수학에서 한 단계 올라가 `무엇이 성립하는가`, `어떤 조건이 필요한가`를 다루는 논리 카드다.

## Key Points

- 정의
  - 집합은 공통 성질을 가진 대상들의 모임이고, 명제는 참과 거짓을 분명히 가를 수 있는 문장이다.
- 핵심 개념
  - 원소
  - 부분집합
  - 합집합
  - 교집합
  - 필요조건
  - 충분조건
- 대표 수식
  - $A\cup B$
  - $A\cap B$
  - $P\Rightarrow Q$
- 대표 성질
  - $P\Rightarrow Q$가 참이면, $P$는 $Q$의 충분조건이고 $Q$는 $P$의 필요조건이다.
  - 명제의 대우 `\neg Q\Rightarrow \neg P`는 원래 명제와 같은 참거짓 값을 가진다.
  - 집합의 드모르간 법칙은 명제의 부정 법칙과 정확히 대응한다.
- 증명 스케치
  - `증명 스케치 (추론)`:
  - 어떤 원소가 $A\cap B$에 속한다는 것은 `A에도 속하고 B에도 속한다`는 뜻이다.
  - 따라서 교집합은 두 조건을 동시에 만족하는 대상만 모은 집합으로 읽을 수 있다.
  - 명제의 대우는 원래 명제와 논리적으로 동치이므로, 증명에서 자주 사용된다.
## Deep Dive

- 집합은 `무엇이 들어 있는가`를 묻고, 명제는 `어떤 조건이 참인가`를 묻는다.
- 필요조건과 충분조건을 구분하면 증명에서 가정과 결론의 방향이 명확해진다.
- 대우는 원래 명제를 뒤집어 이해하는 강력한 도구이고, 반례는 명제가 거짓임을 보이는 가장 직접적인 방법이다.
- 집합 표기와 논리 표기는 서로 번역 가능하므로 같이 읽어야 한다.

## Worked Examples

### 예제 1: 교집합

- $A=\{1,2,3\}$, $B=\{2,3,4\}$이면
  $$
  A\cap B=\{2,3\}
  $$
  이다.

### 예제 2: 합집합

- 같은 집합에서
  $$
  A\cup B=\{1,2,3,4\}
  $$
  이다.

### 예제 3: 명제의 대우

- `x가 4의 배수이면 x는 짝수이다`의 대우는 `x가 짝수가 아니면 x는 4의 배수가 아니다`이다.

## Common Pitfalls

- 충분조건과 필요조건을 뒤바꾸면 증명 방향이 흔들린다.
- 합집합과 교집합을 `또는`과 `그리고`로 번역할 때 중복 조건을 주의해야 한다.
- 명제의 대우와 역은 다르다.
- 반례는 하나면 충분하지만, 참을 보이려면 일반성을 가져야 한다.
- 교육과정 배치
  - 한국 대표 배치에서는 `공통수학2`에서 함수와 함께 논리 언어로 다뤄진다.
- 국가별 배치 스냅샷
  - 한국: `공통수학2`의 독립 구성요소다.
  - 일본: 현재 문서군에서는 독립 대단원보다 식과 증명, 함수 서술 안에 분산되어 읽힌다.
  - 중국: 고등학교 필수과정의 `예비지식`에서 집합과 명제가 명시적으로 등장한다.
  - 미국: K-12에서 standalone보다는 Algebra, proof, set notation 맥락에 분산되는 편이다.
- 표현과 문제 감각
  - 다국어 용어: `set`, `proposition`, `集合`, `命題`
  - 집합은 `모임`, 명제는 `조건의 방향`으로 읽는 감각이 중요하다.

## Connections

- 선수 개념은 [variables-and-expressions.md](./variables-and-expressions.md), [geometric-proof.md](./geometric-proof.md)다.
- 다음 개념으로는 [equations-of-geometric-figures.md](./equations-of-geometric-figures.md), [function-foundations.md](./function-foundations.md), [conditional-probability.md](./conditional-probability.md)가 이어진다.
- 계통 허브에서는 [geometry-strand.md](../geometry-strand.md), [functions-strand.md](../functions-strand.md), [korea-curriculum-hub.md](../korea-curriculum-hub.md)를 함께 본다.

## Open Questions

- `필요조건과 충분조건`을 별도 논리 카드로 분리할지 검토가 필요하다.

## Sources

- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-concept-encyclopedia/china.md`
- `docs/math-concept-encyclopedia/multilingual-glossary.md`
