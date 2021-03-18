package response

import "github.com/gin-gonic/gin"

func ResponseJson(context *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	context.JSON(httpStatus, gin.H{"code": code, "data": data, "msg": msg})
}

func Success(context *gin.Context, data gin.H, msg string) {
	ResponseJson(context, 200, 200, data, msg)
}

func Fail(context *gin.Context, data gin.H, msg string) {
	ResponseJson(context, 400, 200, data, msg)
}

func Abort(context *gin.Context, data gin.H, msg string) {
	ResponseJson(context, 401, 401, data, msg)
}
