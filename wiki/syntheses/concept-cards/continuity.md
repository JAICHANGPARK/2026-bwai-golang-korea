---
title: 연속성
type: synthesis
status: active
updated: 2026-04-09
source_docs:
  - docs/math-concept-encyclopedia/korea.md
  - docs/math-concept-encyclopedia/japan.md
  - docs/math-concept-encyclopedia/us.md
  - docs/math-concept-encyclopedia/multilingual-glossary.md
  - docs/math-concept-encyclopedia/unit-practice-book.md
tags:
  - concept
  - calculus
  - continuity
---

# 연속성

## Summary

연속성은 함수 그래프가 특정 점에서 `끊기지 않고 이어지는가`를 판단하는 개념이다. 극한을 미분으로 넘겨 주는 관문이라서, 미적분에서는 별도 카드로 두는 편이 구조를 더 잘 보여 준다.

## Key Points

- 정의
  - 함수 $f(x)$가 $x=a$에서 연속이라는 것은 $\lim_{x\to a}f(x)=f(a)$가 성립하는 것이다.
- 핵심 개념
  - 함수값
  - 극한값
  - 점에서의 연속
  - 끊김
  - 연결된 그래프
- 대표 수식
  - $\lim_{x\to a}f(x)=f(a)$
- 증명 스케치
  - `증명 스케치 (추론)`:
  - 어떤 점 근처에서 함수값이 한 값으로 가까워지고, 그 값이 실제 함수값과 같으면 그래프를 그 점에서 끊김 없이 이을 수 있다.
  - 그래서 연속성은 `극한값과 실제값이 맞아떨어지는가`를 확인하는 문제로 읽는다.
- 대표 예제
  - $f(x)=2x$는 모든 실수에서 연속이고, 특히 $x=3$에서 $\lim_{x\to 3}f(x)=6=f(3)$이다.
- 교육과정 배치
  - 한국 대표 배치에서는 `미적분I`의 `함수의 극한과 연속`에 놓인다.
- 국가별 배치 스냅샷
  - 한국: `미적분I`의 첫 단원에서 극한과 함께 다룬다.
  - 일본: `수학III`에서 무한 과정과 도함수 사이의 다리로 읽힌다.
  - 미국: Calculus의 `limits and continuity`가 대표 배치다.
  - 중국: 현재 문서군에서는 도함수와 연속적 변화 감각 속에서 간접적으로 읽힌다.
- 표현과 문제 감각
  - 다국어 용어: `continuity`, `連続性`, `连续性`
  - 연속성은 `값이 존재하는가`만이 아니라 `가까워지는 값과 실제값이 일치하는가`를 묻는다.

## Connections

- 선수 개념은 [limits.md](./limits.md), [function-foundations.md](./function-foundations.md)다.
- 같은 축의 인접 개념으로는 [continuity-properties.md](./continuity-properties.md), [differentiation.md](./differentiation.md)가 있다.
- 다음 개념으로는 [continuity-properties.md](./continuity-properties.md), [differentiation.md](./differentiation.md), [calculus-course.md](./calculus-course.md)가 이어진다.

## Open Questions

- `중간값정리`와 `최대최소정리`를 별도 심화 카드로 더 뺄지 후속 설계가 필요하다.

## Sources

- `docs/math-concept-encyclopedia/korea.md`
- `docs/math-concept-encyclopedia/japan.md`
- `docs/math-concept-encyclopedia/us.md`
- `docs/math-concept-encyclopedia/multilingual-glossary.md`
- `docs/math-concept-encyclopedia/unit-practice-book.md`
