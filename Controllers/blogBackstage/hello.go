package blogBackstage

import (
	Jwt "blog-go/Config"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Hello(context *gin.Context) {
	token, err := Jwt.GenerateJwt("name")
	claims, err := Jwt.AnalysisJwt(token)

	fmt.Print(err, token, claims)
	context.JSON(http.StatusOK, gin.H{
		"token":  token,
		"claims": claims,
	})
}
