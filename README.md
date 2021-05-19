# readme

## 注意

由于不熟悉gin的api，加之对于性能不敏感，本仓库即日起将基于flask开发。
除此之外将会重写测试库，不再检测输入和预期输出是否相符，而是实现一个函数。
函数将会判断对于给定输入输出是否合法。


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

- CI的操作非常简单，整体流程是 drone 打完包，然后推到私有镜像 harbor，最后调高可用 ~~k3s~~(刚刚炸了先用swarm) swarm 集群重新编排容器

- 每次发生操作，都会自动部署到一堆服务器上，不过我用nginx负载均衡到202.119.245.31的80端口了，前端访问202.119.245.31:80即可

- drone的后台是<https://drone.njtechstas.top>
