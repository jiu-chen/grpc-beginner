
> demo01
```shell
cd proto  
protoc --go_out=plugins=grpc:../services Prod.proto
```

> server
```shell
cd server
go run main.go
```

> client
```shell
cd client
go run main.go
```