package middleware

import (
	"DemoProjectGO/common"
	"DemoProjectGO/util"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

//AuthMiddleware did sth i know less.
func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		sTokenWithPrefix := context.GetHeader("Autherization")
		// 没有token
		// 这个开头检查我没理解，为啥一定要开头有这个字符串
		// 是库规定的还是用户规定的？
		if sTokenWithPrefix == "" || !strings.HasPrefix(sTokenWithPrefix, "Bearer ") {
			context.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			// 中断请求
			context.Abort()
		}
		sToken := sTokenWithPrefix[7:]
		token, claims := common.ParseToken(sToken)
		if !token.Valid {
			context.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			context.Abort()
		}

		user := util.GetUserFormID(claims.UserID)

	}
}
