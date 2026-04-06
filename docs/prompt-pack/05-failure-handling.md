# 05. 실패 처리 설계 프롬프트

## 목적

이미지 불량, 풀이 불확실, 검증 실패 같은 예외 흐름을 설계한다.

## 언제 사용하나

출력 구조를 정의한 뒤 사용한다.

## 프롬프트

```text
Please list the main failure cases for this math explanation system.

Examples:
- unreadable image
- incomplete problem
- uncertain answer
- inconsistent solution

For each failure case, suggest:
- what the system should detect
- what the system should say to the user
- whether the verifier should reject or ask for clarification

Write the result in Korean.
Keep it practical for a real product demo.
```

## 기대 산출물

- 실패 케이스 목록
- 사용자 메시지 예시
- Verifier 판정 규칙
