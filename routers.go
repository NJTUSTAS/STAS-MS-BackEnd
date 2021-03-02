package main

import (
	"DemoProjectGO/controller"
	"github.com/gin-gonic/gin"
)

func CollectRoute(router *gin.Engine) *gin.Engine {
	router.POST("/api/auth/register", controller.Register)
	router.POST("/api/auth/login", controller.Login)
	router.POST("/enroll/receive", controller.EnrollReceive) //招新接收信息
	return router
}
