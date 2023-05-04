package servicesAdmin

import (
	Init "blog-go/Config"
	"blog-go/Models/modelAdmin"
	"log"
	"time"
)

// CreateBlog 创建内容
func CreateBlog(groupId int, title, content string) int64 {
	var newGroup = modelAdmin.BlogBlogs{
		GroupId:  groupId,
		Title:    title,
		Content:  content,
		Visitors: 0,
		IsShow:   true,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}
	result := Init.DB.Table("blog_blogs").Create(&newGroup)
	println(result)
	if result.Error != nil {
		log.Println("Create group fail : ", result)
	}
	return result.RowsAffected
}

//// GetBlogIsShow  查询blog是否展示
//func GetBlogIsShow(id uint) bool {
//	type BlogContent struct {
//		IsShow bool
//	}
//	var blogContent BlogContent
//	result := Init.DB.Table("blog_blogs").Select("is_show").Where("id= ?", id).Scan(&blogContent)
//	println(result)
//	if result.Error != nil {
//		log.Println("ContentIsShow group fail : ", result)
//	}
//	return blogContent.IsShow
//}

// UpdateBlog blog更新
func UpdateBlog(id uint, groupId int, title string, content string, isShow bool) int64 {
	result := Init.DB.Table("blog_blogs").
		Where("id=?", id).
		Updates(&modelAdmin.BlogBlogs{GroupId: groupId, Title: title, Content: content, IsShow: isShow, UpdateAt: time.Now()})
	println(result)
	if result.Error != nil {
		log.Println("blog_blogs update fail : ", result)
	}
	return result.RowsAffected
}

// SelectBlog blog查询
// 第几页
func SelectBlog(paging int, pageSize int) []modelAdmin.BlogGroups {
	var blogGroups []modelAdmin.BlogGroups
	err := Init.DB.Table("blog_blogs").Limit(pageSize).Offset((paging - 1) * pageSize).Find(&blogGroups)
	if err.Error != nil {
		log.Println("SelectContent group fail : ", err)
	}
	return blogGroups
}
