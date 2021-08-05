
> demo02
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
GODEBUG=x509ignoreCN=0 go run main.go
```
refer: [GODEBUG=x509ignoreCN=0](https://www.cnblogs.com/jackluo/p/13841286.html)