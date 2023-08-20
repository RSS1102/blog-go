package controlAdmin

import (
	"blog-go/Models/modelAdmin"
	"blog-go/Models/modelPublic"
	"blog-go/Services/blogs/servicesAdmin"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateGroup(context *gin.Context) {
	var data modelAdmin.BlogGroups
	err := context.ShouldBindJSON(&data)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	res := servicesAdmin.CreateGroup(data.Group)
	if res > 0 {
		context.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "添加分组成功",
			"data":    res,
		})
	} else {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "添加分组失败",
			"data":    nil,
		})
	}
}
func UpdateGroup(context *gin.Context) {
	var data modelAdmin.BlogGroups
	err := context.ShouldBindJSON(&data)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	res := servicesAdmin.UpdateGroup(data)
	if res > 0 {
		context.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "更新分组成功",
			"data":    res,
		})
	} else {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "更新分组失败",
			"data":    nil,
		})
	}
}

func SelectGroupLimit(context *gin.Context) {
	var page modelPublic.Page
	err := context.ShouldBindJSON(&page)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    200,
			"message": err.Error(),
			"data":    make([]int, 0),
			"total":   0,
		})
		return
	}
	println(page.Current, page.PageSize)
	total, data := servicesAdmin.SelectGroupLimit(page.Current, page.PageSize)
	if total > 0 {
		context.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "查询成功",
			"data":    data,
			"total":   total,
		})
	} else {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "查询失败",
			"data":    nil,
			"total":   0,
		})
	}
}
