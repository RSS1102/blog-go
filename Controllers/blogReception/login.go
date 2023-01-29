package blogReception

import (
	Jwt "blog-go/Config"
	User "blog-go/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(context *gin.Context) {
	var data User.User
	err := context.ShouldBindJSON(&data)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	user := User.LoginUser(data.Username, data.Password)
	fmt.Println(user.ID)
	if user.ID < 1 {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "你输入的账号或者密码不正确",
		})
		return
	}
	tokenString, err := Jwt.GenerateJwt(data.Username)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"code":    "401",
			"message": "token错误，请重新登陆",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"code":    200,
		"token":   tokenString,
		"message": "登陆成功",
	})
}
