package controlAdmin

import (
	Jwt "blog-go/Config"
	User "blog-go/Models/modelAdmin"
	"blog-go/Services/blogs/servicesAdmin"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(context *gin.Context) {
	var data User.User
	err := context.ShouldBindJSON(&data)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code":    400,
			"message": err.Error(),
			"token":   nil,
		})
		return
	}
	fmt.Println(data)
	user := servicesAdmin.SelectUser(data.Username, data.Password)
	fmt.Println(user.ID)
	if user.ID < 1 {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "你输入的账号或者密码不正确",
			"token":   nil,
		})
		return
	}
	tokenString, err := Jwt.GenerateJwt(data.Username)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "token错误，请重新登陆",
			"token":   nil,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"code":    200,
		"token":   tokenString,
		"message": "登陆成功",
	})
}

func LoginOut(context *gin.Context) {

	context.JSON(http.StatusOK, gin.H{
		"code":    200,
		"data":    "",
		"message": "退出成功",
	})
}
