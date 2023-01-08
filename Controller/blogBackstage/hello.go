package blogBackstage

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Hello(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Hello q1mi!",
	})
}
