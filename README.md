# 데봉이치킨 - Employee 컴포넌트


### API
| Method | URL | 설명        |
| :------------ | :----------- | :------------------- |
| GET    | /employees          | 직원들 조회 |
| GET   | /employees/{:id}      | 직원ID로 직원 조회 |
| POST    | /employees  | 직원 생성 (* 필수 데이터 미지정)|
| PUT | /employees/ {:id} | 직원 정보 수정 |
| DELETE | /employees/ {:id} | 직원 정보 삭제 |


### 추가 예정 API
| Method | URL | 설명 |
| :------------ | :---------------- | :----------- |
| ? | /employees/sendAuthentication | 인증번호 발송 |
| ? | /employees/verify | 인증번호 검증 |

