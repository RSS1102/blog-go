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

// SelectBlog blog查询 分页
func SelectBlog(Current int, pageSize int) (int, []modelAdmin.BlogBlogs) {
	var blogBlogs []modelAdmin.BlogBlogs
	var total int64
	err := Init.DB.Table("blog_blogs").Count(&total).Error
	if err != nil {
		log.Println("Failed to get total record count:", err)
		return 0, nil
	}

	errs := Init.DB.Table("blog_blogs").Joins("JOIN blog_groups ON blog_blogs.group_id = blog_groups.ID").
		Limit(pageSize).Offset((Current - 1) * pageSize).Find(&blogBlogs)
	if errs.Error != nil {
		return 0, nil
	}
	return int(total), blogBlogs
}
