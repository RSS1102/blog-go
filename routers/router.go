package routers

import (
	"blog-go/Controllers/blogBackstage"
	"blog-go/Controllers/blogReception"
	"blog-go/Middlewares"
	"fmt"
	"github.com/gin-gonic/gin"
)

func RouterInit() {
	router := gin.Default()
	v1 := router.Group("/client")
	{
		v1.POST("/s", blogBackstage.Hello)
	}
	v2 := router.Group("/admin")
	{
		v2.GET("/login", blogReception.Login)
	}
	v3 := router.Group("/admin", Middlewares.JWTAuthMiddleware())
	{
		v3.GET("/hello", blogBackstage.Hello)
	}
	if err := router.Run(); err != nil {
		fmt.Printf("startup service failed, err:%v\n\n", err)
	}
}