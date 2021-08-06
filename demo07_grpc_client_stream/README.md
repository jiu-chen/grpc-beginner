客户端分批次发送，服务端一次性处理并返回
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
client stream resp: 
users:{user_id:1}  users:{user_id:2  user_score:1}  users:{user_id:3  user_score:2}  users:{user_id:4  user_score:3}  users:{user_id:5  user_score:4}  users:{user_id:1  user_score:5}  users:{user_id:2  user_score:6}  users:{user_id:3  user_score:7}  users:{user_id:4  user_score:8}  users:{user_id:5  user_score:9}  users:{user_id:1  user_score:10}  users:{user_id:2  user_score:11}  users:{user_id:3  user_score:12}  users:{user_id:4  user_score:13}  users:{user_id:5  user_score:14}
```