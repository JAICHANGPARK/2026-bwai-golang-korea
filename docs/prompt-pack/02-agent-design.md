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
