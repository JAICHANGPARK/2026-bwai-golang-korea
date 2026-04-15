# Phase 1 - Terminal Single Agent

목표는 터미널에서 텍스트 수학 문제를 넣고, 풀이와 해설을 바로 확인할 수 있는 최소형 에이전트를 만드는 것이다.

## 권장 구조

```text
phase-1-terminal/
├── README.md
├── app/
│   ├── agent.py
│   └── prompts.py
└── tests/
    └── smoke_text_cases.md
```

## 완료 기준

- 텍스트 문제 입력을 받는다.
- 정답과 단계별 해설을 텍스트로 출력한다.
- 샘플 문제 2개 이상으로 터미널 smoke test가 가능하다.
