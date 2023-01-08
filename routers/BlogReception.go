package routers

import (
	"blog-go/Controller/blogReception"
	"github.com/gin-gonic/gin"
)

func BlogReception(e *gin.Engine) {
	e.GET("/login", blogReception.Login)
}
