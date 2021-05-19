package response

import (
	"DemoProjectGO/model"
	"github.com/gin-gonic/gin"
)

func ReturnJson(context *gin.Context, httpStatus int, data gin.H, msg string) {
	context.JSON(httpStatus, gin.H{"code": httpStatus, "data": data, "msg": msg})
}

func ReturnArray(context *gin.Context, httpStatus int, data []model.Note) {
	context.JSON(httpStatus, data)
}

func Success(context *gin.Context, data gin.H, msg string) {
	ReturnJson(context, 200,  data, msg)
}

func Fail(context *gin.Context, data gin.H, msg string) {
	ReturnJson(context, 400,  data, msg)
}

func Abort(context *gin.Context, data gin.H, msg string) {
	ReturnJson(context, 401,  data, msg)
}
