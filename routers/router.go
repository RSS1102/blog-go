package routers

import (
	"blog-go/Controllers/blogAdmin"
	"blog-go/Controllers/blogClient"
	"blog-go/Middlewares"
	"fmt"
	"github.com/gin-gonic/gin"
)

func RouterInit() {
	router := gin.Default()
	//client展示
	v1 := router.Group("/client")
	{
		v1.POST("/s", blogClient.Hello)
	}
	//admin登陆
	admin1 := router.Group("/admin")
	{
		admin1.POST("/login", blogAdmin.Login)
	}
	//admin内部操作
	admin2 := router.Group("/admin", Middlewares.JWTAuthMiddleware())
	{
		admin2.POST("/loginOut", blogAdmin.LoginOut)
		admin2.GET("/hello", blogClient.Hello)
	}
	if err := router.Run(); err != nil {
		fmt.Printf("startup service failed, err:%v\n\n", err)
	}
}
