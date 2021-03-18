package middleware

import (
	"DemoProjectGO/common"
	"DemoProjectGO/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

//AuthMiddleware did sth i know less.
func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		Abort401 := func() {
			context.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			// 中断请求
			context.Abort()
		}

		sToken := context.GetHeader("Authorization")

		//rfc6750d的规范要求bearer token以“bearer ”开头
		//参见：https://tools.ietf.org/html/rfc6750
		//这里我们暂时不使用这个特性
		//sPrefixToken := context.GetHeader("Authorization")
		//if !strings.HasPrefix(sPrefixToken, "bearer") {
		//	Abort401()
		//}
		//sToken :=sPrefixToken[7:0]

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
