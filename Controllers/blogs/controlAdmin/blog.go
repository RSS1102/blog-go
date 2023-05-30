package controlAdmin

import (
	"blog-go/Models/modelAdmin"
	"blog-go/Models/modelPublic"
	"blog-go/Services/blogs/servicesAdmin"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateBlog(context *gin.Context) {
	var data modelAdmin.BlogBlogs
	err := context.ShouldBindJSON(&data)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	res := servicesAdmin.CreateBlog(data.GroupId, data.Content, data.Title)
	if res > 0 {
		context.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "添加内容成功",
		})
	} else {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "添加内容失败",
		})
	}
}

func UpdateBlog(context *gin.Context) {
	var data modelAdmin.BlogBlogs
	err := context.ShouldBindJSON(&data)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}
	println("error,message", data.GroupId)
	res := servicesAdmin.UpdateBlog(data.ID, data.GroupId, data.Title, data.Content, data.IsShow)
	if res > 0 {
		context.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "blog更新成功",
		})
	} else {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "blog更新失败",
		})
	}
}

func SelectBlog(context *gin.Context) {
	var page modelPublic.Page
	err := context.ShouldBindJSON(&page)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	println(page.Current, page.PageSize)
	total, groups := servicesAdmin.SelectBlog(page.Current, page.PageSize)
	if total > 0 {
		context.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "查询成功",
			"total":   total,
			"data":    groups,
		})
	} else {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "查询失败",
			"total":   0,
			"data":    nil,
		})
	}
}
