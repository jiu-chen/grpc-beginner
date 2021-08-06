客户端一次性发送，服务端分批次处理并返回

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
users:{user_id:3 user_score:103} users:{user_id:4 user_score:104}
users:{user_id:5 user_score:105} users:{user_id:6 user_score:106}
users:{user_id:7 user_score:107}
main exit
```