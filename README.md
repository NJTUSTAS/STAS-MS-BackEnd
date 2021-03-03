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

## CI-CD

[![Build Status](https://drone.njtechstas.top/api/badges/NJTUSTAS/STAS-MS-BackEnd/status.svg)](https://drone.njtechstas.top/NJTUSTAS/STAS-MS-BackEnd)

- CI的操作非常简单，整体流程是 drone 打完包，然后推到私有镜像 harbor，最后调用 ~~k3s~~(刚刚炸了先用swarm) swarm 重新编排容器

- 每次发生提交，都会自动部署到202.119.245.31上，后端使用的8080端口无法在公网访问，不过我用nginx代理到80端口了

- drone的后台是https://drone.njtechstas.top
