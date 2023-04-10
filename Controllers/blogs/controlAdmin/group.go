package controlAdmin

import (
	Group "blog-go/Models/modelsBlogs"
	"blog-go/Services/blogs/servicesAdmin"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddGroup(context *gin.Context) {
	var data Group.BlogGroups
	err := context.ShouldBindJSON(&data)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if !data.IsShow {
		data.IsShow = true
	}
	res := servicesAdmin.CreateGroup(data.Group, data.IsShow)
	if res > 0 {
		context.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "添加成功",
		})
	} else {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "添加失败",
		})
	}

}
