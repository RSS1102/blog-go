package Middlewares

import (
	Jwt "blog-go/Config"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// JWTAuthMiddleware JWT中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    302,
				"message": "请携带Token发起请求。",
			})
			c.Abort()
			return
		}
		parts := strings.Split(authHeader, ".")
		if len(parts) != 3 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 2004,
				"msg":  "请求头中auth格式有误",
			})
			c.Abort()
			return
		}
		claims, err := Jwt.AnalysisJwt(authHeader)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 205,
				"msg":  "无效的Token",
			})
			c.Abort()
			return
		}
		fmt.Print(claims)
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("username", claims.Username)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
