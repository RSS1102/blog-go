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
func UpdateBlog(updates modelAdmin.BlogBlogs) int64 {
	// 更新记录的不固定字段
	updateData := map[string]interface{}{
		"Title":    updates.Title,
		"Content":  updates.Content,
		"Visitors": updates.Visitors,
		"IsShow":   updates.IsShow,
		"GroupId":  updates.GroupId,
	}
	println(updateData)
	// 移除未传入的字段
	if updates.Title == "" {
		delete(updateData, "Title")
	}
	if updates.Content == "" {
		delete(updateData, "Content")
	}
	if updates.Visitors == 0 {
		delete(updateData, "Visitors")
	}
	if updates.GroupId == 0 {
		delete(updateData, "GroupId")
	}

	result := Init.DB.Table("blog_blogs").
		Model(&modelAdmin.BlogBlogs{}).
		Where("id=?", updates.ID).
		Updates(updateData)
	if result.Error != nil {
		log.Println("blog_blogs update fail : ", result)
	}
	return result.RowsAffected
}

// SelectBlog blog查询 分页
func SelectBlog(Current int, pageSize int) (int, []modelAdmin.MergedBlogs) {
	var blogBlogs []modelAdmin.MergedBlogs
	var total int64
	err := Init.DB.Table("blog_blogs").Count(&total).Error
	if err != nil {
		log.Println("Failed to get total record count:", err)
		return 0, nil
	}

	errs := Init.DB.Table("blog_blogs").Joins("JOIN blog_groups ON blog_blogs.group_id = blog_groups.id").
		Select("blog_blogs.*, blog_groups.group, blog_groups.is_show as group_is_show").
		Limit(pageSize).Offset((Current - 1) * pageSize).Find(&blogBlogs)
	if errs.Error != nil {
		return 0, nil
	}
	return int(total), blogBlogs
}
