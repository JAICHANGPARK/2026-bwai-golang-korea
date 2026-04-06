# 01. 사용자 흐름 정의 프롬프트

## 목적

사용자가 문제 이미지를 올렸을 때 시스템이 어떤 순서로 동작해야 하는지 정의한다.

## 언제 사용하나

핸즈온 시작 직후, 전체 시스템 그림을 먼저 잡을 때 사용한다.

## 프롬프트

```text
We are designing a math explanation system for middle and high school students.

The user uploads a math problem image and wants:
- the correct answer
- a step-by-step explanation
- key concepts
- a verified final response

Please design the end-to-end user flow in Korean.
Keep it simple and suitable for a 1-hour workshop.
Assume we are not using RAG.

The flow should clearly separate:
- user input
- solving
- verification
- final response
```

## 기대 산출물

- 사용자 입력
- Solver 호출
- Verifier 호출
- 최종 승인 응답
