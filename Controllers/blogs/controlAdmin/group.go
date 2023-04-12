package controlAdmin

import (
	Group "blog-go/Models/modelsBlogs"
	"blog-go/Services/blogs/servicesAdmin"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateGroup(context *gin.Context) {
	var data Group.BlogGroups
	err := context.ShouldBindJSON(&data)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	res := servicesAdmin.CreateGroup(data.Group)
	if res > 0 {
		context.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "添加分组成功",
		})
	} else {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "添加分组失败",
		})
	}
}
func UpdateGroup(context *gin.Context) {
	var data Group.BlogGroups
	err := context.ShouldBindJSON(&data)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	res := servicesAdmin.UpdateGroup(data.ID, data.Group)
	if res > 0 {
		context.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "更新分组成功",
		})
	} else {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "更新分组失败",
		})
	}
}
func DeleteGroup(context *gin.Context) {
	var data Group.BlogGroups
	err := context.ShouldBindJSON(&data)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	res := servicesAdmin.DeleteGroup(data.ID)
	if res > 0 {
		context.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "删除分组成功",
		})
	} else {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "删除分组失败",
		})
	}
}

type Page struct {
	Paging   int `json:"paging"`
	PageSize int `json:"pageSize"`
}

func SelectGroup(context *gin.Context) {
	var page Page
	err := context.ShouldBindJSON(&page)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	println(page.Paging, page.PageSize)
	res := servicesAdmin.SelectGroup(page.Paging, page.PageSize)
	if len(res) > 0 {
		context.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "查询成功",
			"date":    res,
		})
	} else {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "查询失败",
		})
	}
}
