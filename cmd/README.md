1.启动了英雄联盟客户端后执行下列命令可以获得客户端启动时的token和port组成的查询数据的url
```
go run .\cmd\see_port_token\give_url_port.go
```


2.根据召唤师名字可以展示用户的一些简单信息
```
go run .\cmd\show_info\show_my_info.go
```

2.在游戏开始后联盟的客户端会被关闭，任务管理器无法找到lol的进程，所以其他需要即时访问token的脚本无法执行,这个脚本只要添加好了port_token就没问题
```
go run .\cmd\check_when_play\get_recent_kda_by_name_need_token.go
```