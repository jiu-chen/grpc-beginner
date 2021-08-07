客户端&&服务端流模式: 客户端分批次发送，服务端分批次处理并返回
```bash
$ chmod u+x gen.sh                     
./gen.sh
```
+ server
```bash
cd server
go run main.go
```
+ client
```bash
cd client
go run main.go
```

+ response
``` bash
client server stream batch resp:
 users:{user_id:1}  users:{user_id:2  user_score:1}  users:{user_id:3  user_score:2}  users:{user_id:4  user_score:3}  users:{user_id:5  user_score:4}
client server stream batch resp:
 users:{user_id:6  user_score:5}  users:{user_id:7  user_score:6}  users:{user_id:8  user_score:7}  users:{user_id:9  user_score:8}  users:{user_id:10  user_score:9}
client server stream batch resp:
 users:{user_id:11  user_score:10}  users:{user_id:12  user_score:11}  users:{user_id:13  user_score:12}  users:{user_id:14  user_score:13}  users:{user_id:15  user_score:14}

```