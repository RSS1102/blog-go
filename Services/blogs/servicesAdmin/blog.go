package servicesAdmin

import (
	Init "blog-go/Config"
	"blog-go/Models/modelsBlogs"
	"log"
	"time"
)

// CreateBlog 创建内容
func CreateBlog(groupId int, title, content string) int64 {
	var newGroup = modelsBlogs.BlogBlogs{
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

// GetBlogIsShow  查询blog是否展示
func GetBlogIsShow(id uint) bool {
	type BlogContent struct {
		IsShow bool
	}
	var blogContent BlogContent
	result := Init.DB.Table("blog_blogs").Select("is_show").Where("id= ?", id).Scan(&blogContent)
	println(result)
	if result.Error != nil {
		log.Println("ContentIsShow group fail : ", result)
	}
	return blogContent.IsShow
}

// UpdateBlogGroupId  blog分组更新
func UpdateBlogGroupId(id uint, groupId int) int64 {
	result := Init.DB.Table("blog_blogs").Where("id=?", id).Updates(&modelsBlogs.BlogBlogs{GroupId: groupId, UpdateAt: time.Now()})
	println(result)
	if result.Error != nil {
		log.Println("group_id update fail : ", result)
	}
	return result.RowsAffected
}

// UpdateBlogTitle  blog标题分组
func UpdateBlogTitle(id uint, title string) int64 {
	result := Init.DB.Table("blog_blogs").Where("id=?", id).Updates(&modelsBlogs.BlogBlogs{Title: title, UpdateAt: time.Now()})
	println(result)
	if result.Error != nil {
		log.Println("group_id update fail : ", result)
	}
	return result.RowsAffected
}

// UpdateBlogData  blog内容更新
func UpdateBlogData(id uint, content string) int64 {
	result := Init.DB.Table("blog_blogs").Where("id=?", id).Updates(&modelsBlogs.BlogBlogs{Content: content, UpdateAt: time.Now()})
	println(result)
	if result.Error != nil {
		log.Println("content update fail : ", result)
	}
	return result.RowsAffected
}

// UpdateBlogIsShow  blog是否展示
func UpdateBlogIsShow(id uint, isShow bool) int64 {
	result := Init.DB.Table("blog_blogs").Where("id=?", id).Updates(&modelsBlogs.BlogBlogs{IsShow: isShow, UpdateAt: time.Now()})
	println(result)
	if result.Error != nil {
		log.Println("is_show update fail : ", result)
	}
	return result.RowsAffected
}

// SelectContent 分组查询
// 第几页
//func SelectContent(paging int, pageSize int) []modelsBlogs.BlogContents {
//	var blogGroups []modelsBlogs.BlogContents
//	err := Init.DB.Table("blog_content").Limit(pageSize).Offset((paging - 1) * pageSize).Find(&blogGroups)
//	if err.Error != nil {
//		log.Println("SelectContent group fail : ", err)
//	}
//	return blogGroups
//}
