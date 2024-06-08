# Go-gRPC
studying GoLang with gRPC
---
```
go run server.go
go run client.go
```
### Prg01: hello_grpc

간단하게 숫자를 받아서 제곱 값을 계산하는 예제

1. server.go를 실행하여 50051번 포트를 연다. 
2. client.go를 실행하여 value=4 로 메시지를 만들고 서버로부터 16을 받아온다.

### Prg02: Bidirectional-streaming

입력도 stream, 출력도 stream인 rpc함수

1. server.go를 실행하여 50051번 포트를 연다. 
2. client.go를 실행하여 50051번 채널에 연결하면 client와 server가 5개의 msg를 생성하여 서로에게 streaming한다.

### Prg03: Client-streaming

입력으로 stream을 받고, msg수를 count하여 number로 출력하는 rpc함수

1. server.go를 실행하여 50051번 포트를 연다. 
2. client.go를 실행하여 50051번 채널에 연결하면 client가  msg를 생성하여 server에게 streaming한다. 
3. server는 msg를 loop문으로 받으며 msg수를 count한다.
4. server가 EOF이면 count수를 client에게 응답한다. 

### Prg04: Bidirectional-streaming

입력으로 number를 받고, 출력으로 stream을 내보내는 rpc함수

1. server.go를 실행하여 50051번 포트를 연다. 
2. client.go를 실행하여 value=5로 설정하여 server와 연결한다.
3. server가 client로부터 받은 value=5를 콘솔에 출력한다. 
4. server가 msg 5개를 생성하여 client에게 streaming한다.
