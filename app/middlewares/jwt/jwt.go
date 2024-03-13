package jwt

import (
	"classroom-system/utils/jwt"
	"github.com/gin-gonic/gin"
)

// 基于jwt认证中间件
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = 200
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 2001 //无token，无权访问
		} else {
			//解析token
			claims, err := jwt.ParseToken(token)
			if err != nil {
				code = 2002 //token不合法
			} else {
				id := claims.Id
				name := claims.Name
				password := claims.Password

				// 将当前请求的user信息保存到请求的上下文c上
				c.Set("id", id)
				c.Set("name", name)
				c.Set("password", password)
			}
		}
		if code != 200 {
			c.JSON(200, gin.H{
				"code": code,
				"msg":  "无权限访问",
				"data": data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
