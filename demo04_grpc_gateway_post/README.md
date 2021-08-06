## 场景练习(1) POST方式提交订单数据
-- gateway实现api/http

```bash
cd server/proto && protoc --go_out=plugins=grpc:../services Prod.proto
protoc --go_out=plugins=grpc:../services Models.proto
protoc --go_out=plugins=grpc:../services Orders.proto

protoc --grpc-gateway_out=logtostderr=true:../services Prod.proto
protoc --grpc-gateway_out=logtostderr=true:../services Orders.proto
cd ../..
```

* 启动服务端
```bash
cd server
go run server.go
# switch another terminal
go run httpserver.go
```
* POST查询（或使用postman请求）
```bash
curl -H "Content-Type: application/json" -X POST -d '{"order_id":34,"order_no":"bj00123456","user_id":8901,"order_money":34.12}' "http://localhost:8080/v1/orders"
```

> client 调用
```bash
cd client
go run main.go
```

> server log
``` go
2021/08/06 13:46:45 grpc server starts...
order_id:34  order_no:"bj00123456"  user_id:8901  order_money:34.12
server收到订单信息:订单号:bj00123456, 订单金额:34.12
```