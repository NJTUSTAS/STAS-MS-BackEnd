package main

import (
	"DemoProjectGO/controller"
	"DemoProjectGO/middleware"

	"github.com/gin-gonic/gin"
)

// CollectRoute _
func CollectRoute(router *gin.Engine) *gin.Engine {
	//以下是路由列表，每个一行。
	//第一个参数规定相对路径
	//第二个参数决定业务逻辑
	//由于反向代理要求，请把所有相对路径放在/api下
	router.POST("/api/auth/register", controller.Register)
	router.POST("/api/auth/login", controller.Login)
	router.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)

	router.POST("/api/enroll/receive", controller.EnrollReceive) //招新接收信息
	return router
}
