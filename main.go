package main

import (
	"blog-go/routers"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routers.BlogReception(router)
	routers.BlogBackstage(router)

	if err := router.Run(); err != nil {
		fmt.Printf("startup service failed, err:%v\n\n", err)
	}
}
