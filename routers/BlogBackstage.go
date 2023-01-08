package routers

import (
	"blog-go/Controller/blogBackstage"
	"github.com/gin-gonic/gin"
)

func BlogBackstage(e *gin.Engine) {
	e.GET("/hello", blogBackstage.Hello)
}
