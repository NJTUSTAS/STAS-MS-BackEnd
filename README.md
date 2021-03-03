# readme

## 文档

各api的文档，请参见 doc/api.md.

## 关于每个文件是干啥的

- /common 一些通用文件
    - database.go：数据库io相关
    - jwt.go：token发放

- /controller

  包含各api的业务逻辑

- /model

  数据库中表的格式接口
  
- /util

  存放一些公用函数

- /middleware 认证中间件

- go.mod，go.sum

  存放golang包依赖关系

- main.go

  程序入口

- routers.go

  存放各api
  
- test.py 测试api的py脚本

