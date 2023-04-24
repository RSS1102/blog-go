package controlAdmin

import (
	"blog-go/Models/modelsBlogs"
	"blog-go/Services/blogs/servicesAdmin"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateBlog(context *gin.Context) {
	var data modelsBlogs.BlogBlogs
	err := context.ShouldBindJSON(&data)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
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
	var data modelsBlogs.BlogBlogs
	err := context.ShouldBindJSON(&data)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	println("error,message", data.GroupId)
	var message string
	error := 0
	if data.GroupId != 0 {
		res := servicesAdmin.UpdateBlogGroupId(data.ID, data.GroupId)
		if res > 0 {
			message = message + "blog分组更新成功。"
		} else {
			message = message + "blog分组更新失败。"
			error += 1
		}
		println(error, message)
	}
	if data.Title != "" {
		res := servicesAdmin.UpdateBlogTitle(data.ID, data.Title)
		if res > 0 {
			message = message + "blog标题更新成功"
		} else {
			message = message + "blog标题更新失败"
			error += 1
		}
		println(error, message)
	}
	if data.Content != "" {
		res := servicesAdmin.UpdateBlogData(data.ID, data.Content)
		if res > 0 {
			message = message + "blog内容更新成功"
		} else {
			message = message + "blog内容更新失败"
			error += 1
		}
		println(error, message)
	}
	//查询是否相等
	resIsShow := servicesAdmin.GetBlogIsShow(data.ID)
	if data.IsShow != resIsShow {
		if data.IsShow {
			res := servicesAdmin.UpdateBlogIsShow(data.ID, data.IsShow)
			if res > 0 {
				message = message + "blog展示成功"
			} else {
				message = message + "blog展示失败"
				error += 1
			}
		} else {
			res := servicesAdmin.UpdateBlogIsShow(data.ID, data.IsShow)
			if res > 0 {
				message = message + "blog取消展示成功"
			} else {
				message = message + "blog取消展示失败"
				error += 1
			}
		}
	} else {
		message = message + "blog展示状态相等"
	}

	if error > 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": message,
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": message,
		})
	}
}

//func ContentIsShow(context *gin.Context) {
//	var data modelsBlogs.BlogContents
//	err := context.ShouldBindJSON(&data)
//	if err != nil {
//		context.JSON(http.StatusBadRequest, gin.H{
//			"message": err.Error(),
//		})
//		return
//	}
//	res := servicesAdmin.ContentIsShow(data.ID, data.IsShow)
//	if res > 0 {
//		context.JSON(http.StatusOK, gin.H{
//			"code":    200,
//			"message": "内容成功",
//		})
//	} else {
//		context.JSON(http.StatusBadRequest, gin.H{
//			"code":    400,
//			"message": "隐藏内容失败",
//		})
//	}
//}

//type Page struct {
//	Paging   int `json:"paging"`
//	PageSize int `json:"pageSize"`
//}

//func SelectContent(context *gin.Context) {
//	var page Page
//	err := context.ShouldBindJSON(&page)
//	if err != nil {
//		context.JSON(http.StatusBadRequest, gin.H{
//			"message": err.Error(),
//		})
//		return
//	}
//	println(page.Paging, page.PageSize)
//	res := servicesAdmin.SelectGroup(page.Paging, page.PageSize)
//	if len(res) > 0 {
//		context.JSON(http.StatusOK, gin.H{
//			"code":    200,
//			"message": "查询成功",
//			"date":    res,
//		})
//	} else {
//		context.JSON(http.StatusBadRequest, gin.H{
//			"code":    400,
//			"message": "查询失败",
//		})
//	}
//}
