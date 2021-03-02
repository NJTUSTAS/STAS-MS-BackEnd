package main

import (
	"DemoProjectGO/controller"
	"github.com/gin-gonic/gin"
)

func CollectRoute(router *gin.Engine) *gin.Engine {
	//以下是路由列表，每个一行。
	//第一个参数规定相对路径
	//第二个参数决定业务逻辑
	router.POST("/api/auth/register", controller.Register)
	router.POST("/api/auth/login", controller.Login)
	router.POST("/enroll/receive", controller.EnrollReceive) //招新接收信息
	return router
}
