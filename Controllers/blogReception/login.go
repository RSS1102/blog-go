package blogReception

import (
	User "blog-go/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(context *gin.Context) {
	username, usernameOk := context.GetPostForm("username")
	password, passwordOk := context.GetPostForm("password")
	fmt.Println(username, password)
	if !usernameOk {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "username不可为空",
		})
		return
	}
	if !passwordOk {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "password不可为空",
		})
		return
	}
	user := User.LoginUser(username, password)
	fmt.Println(user.ID)
	if user.ID < 1 {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "你输入的账号或者密码不正确",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "登陆成功",
	})
	//}
}
