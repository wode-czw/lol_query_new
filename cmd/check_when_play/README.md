在游戏开始后联盟的客户端会被关闭，任务管理器无法找到lol的进程，所以其他需要即时访问token的脚本无法执行,这个脚本只要添加好了port_token就没问题
```
go run .\cmd\check_when_play\get_recent_kda_by_name_need_token.go
```