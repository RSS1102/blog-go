package routers

import (
	"blog-go/Controllers/blogBackstage"
	"blog-go/Controllers/blogReception"
	"fmt"
	"github.com/gin-gonic/gin"
)

func RouterInit() {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.POST("/login", blogReception.Login)
	}
	v2 := router.Group("/v2")
	{
		v2.GET("/hello", blogBackstage.Hello)
	}
	if err := router.Run(); err != nil {
		fmt.Printf("startup service failed, err:%v\n\n", err)
	}
}
