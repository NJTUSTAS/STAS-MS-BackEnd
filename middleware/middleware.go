package middleware

import (
	"DemoProjectGO/common"
	"DemoProjectGO/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

//AuthMiddleware did sth i know less.
func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {

		sToken := context.GetHeader("Authorization")

		Abort401 := func() {
			context.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			// 中断请求
			context.Abort()
		}
		if sToken == "" {
			Abort401()
			return
		}

		token, claims := common.ParseToken(sToken)
		if !(token.Valid) {
			Abort401()
			return
		}

		user := util.GetUserFormID(claims.UserID)
		if user.ID == 0 {
			Abort401()
			return
		}
		// 用户信息写入上下文
		context.Set("user", user)
		context.Next() //这是干啥的我不是很理解

	}
}
