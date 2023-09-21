
# lol_query_tools
query_rank_info_by_name is the main file ,modify the name you want ,
later I will give a gin   GUI to get the infomation.
make sure the LOL  client is active when you use this file.
the account you want to query must in the same area with your account who is logined.

查询英雄联盟的最近30把排位的熟悉英雄的召唤师信息--根据名字,id,并且会进行很多公平的计算，
评分数据会比wegame的更加可靠，因为会计算送头的的负面信息,如果可能还会计算无效伤害

即使这个人隐藏了战绩也没用关系，因为wegame不让查的是wegame的数据库，我们使用lcu数据库查的是英雄联盟存的数据

There are several script to get infomation in cmd folder. 

1.启动了英雄联盟客户端后执行下列命令可以获得客户端启动时的token和port组成的查询数据的url
```
go run .\cmd\get_token_port\give_url_port.go
```


2.根据召唤师名字可以展示用户的一些简单信息
```
go run .\cmd\show_my_account\show_my_info.go
```

编译程序就能够得到应用程序，只需要在启动客户端之后点击程序就在终端查询了。（还没实现）
```
go build app.go
```
