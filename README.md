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

### GO Command

go get -u "repository  path" : go dependency 를 가져 온다

go run ./*.go : go 를 실행시킨다

govendor{https://github.com/kardianos/govendor} fetch "something" : Heroku관련 dependency정보를 fetch한다


[TESTING]
go test : 해당 package의 {Something}_test.go에 Test로 시작하는 메소드를 테스트한다
go test ./... : Sub Package까지 Test실행