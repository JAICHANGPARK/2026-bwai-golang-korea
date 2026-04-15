# 02. 에이전트 역할 정의 프롬프트

## 목적

멀티 에이전트 구조에서 `Solver/Explainer`와 `Expert Verifier`의 역할을 분리한다.

## 언제 사용하나

사용자 흐름을 정의한 다음 바로 사용한다.

## 프롬프트

```text
We want to design this as a multi-agent system, not a single agent.

Please define two agents:
1. a Solver/Explainer Agent
2. an Expert Verifier Agent

For each agent, describe:
- role
- responsibilities
- what it should not do
- input
- output

Write the result in Korean and keep it practical for a workshop.
Make the distinction very clear.
```

## 기대 산출물

- Solver는 문제를 읽고 풀이 초안을 만든다.
- Verifier는 정답과 논리를 재검토한다.
- 두 에이전트가 서로 역할을 침범하지 않는다.

## 선택 확장: Wiki Enricher

기본형이 안정적으로 동작한 뒤에는 `검토 승인 후 학습 정보 보강`용 에이전트를 따로 둘 수 있다. 이 확장 트랙은 [02a-wiki-enricher-extension.md](./02a-wiki-enricher-extension.md)에서 이어서 정리한다.
