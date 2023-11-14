package servicesAdmin

import (
	Init "blog-go/Config"
	"blog-go/Models/modelAdmin"
	"log"
	"time"
)

// TempBlogGroup 覆盖原有的createAt和updateAt的json字段
type TempBlogGroup struct {
	*modelAdmin.BlogGroups        // 嵌入原始结构体
	BlogBlogsCount         int64  `json:"blogBlogsCount"`
	CreatedAt              string `json:"createAt"`
	UpdatedAt              string `json:"updateAt"`
}

// CreateGroup 创建分组
func CreateGroup(group string) int64 {
	var newGroup = modelAdmin.BlogGroups{
		Group:    group,
		IsShow:   true,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}

	result := Init.DB.Create(&newGroup)
	println(result)
	if result.Error != nil {
		log.Println("Create group fail : ", result)
	}
	return result.RowsAffected
}

// UpdateGroup 分组更新
func UpdateGroup(updates modelAdmin.BlogGroups) int64 {
	// 更新记录的不固定字段
	updateData := map[string]interface{}{
		"Group":    updates.Group,
		"IsShow":   updates.IsShow,
		"UpdateAt": time.Now(),
	}
	// 移除未传入的字段
	if updates.Group == "" {
		delete(updateData, "Group")
	}

	result := Init.DB.Table("blog_groups").
		Model(&modelAdmin.BlogGroups{}).
		Where("id=?", updates.ID).
		Updates(updateData)

	if result.Error != nil {
		log.Println("UpdateGroup group fail : ", result)
	}
	return result.RowsAffected
}

// SelectGroupLimit 分组查询
// 第几页
func SelectGroupLimit(Current int, pageSize int) (int, []TempBlogGroup) {

	var blogGroups []TempBlogGroup
	var total int64

	err := Init.DB.Table("blog_groups").Count(&total).Error
	if err != nil {
		log.Println("Failed to get total record count:", err)
		return 0, nil
	}
	if Current == -1 || pageSize == -1 {
		errs := Init.DB.Table("blog_groups").Find(&blogGroups)
		if errs.Error != nil {
			return 0, nil
		}
		for i := range blogGroups {
			blogGroups[i].CreatedAt = blogGroups[i].CreateAt.Format("2006-01-02 15:04:05")
			blogGroups[i].UpdatedAt = blogGroups[i].UpdateAt.Format("2006-01-02 15:04:05")
		}
		return int(total), blogGroups
	}
	errs := Init.DB.Table("blog_groups").
		Joins("JOIN blog_blogs ON blog_blogs.group_id = blog_groups.id").
		Select("blog_groups.*, COUNT(blog_blogs.id) as blog_blogs_count").
		Group("blog_groups.id").
		Find(&blogGroups).
		Limit(pageSize).Offset((Current - 1) * pageSize)
	// 格式化每个 BlogGroup 的 CreatedAt 字段
	for i := range blogGroups {
		blogGroups[i].CreatedAt = blogGroups[i].CreateAt.Format("2006-01-02 15:04:05")
		blogGroups[i].UpdatedAt = blogGroups[i].UpdateAt.Format("2006-01-02 15:04:05")
	}

	if errs.Error != nil {
		log.Println("SelectContent group fail : ", err)
	}
	println(blogGroups)
	return int(total), blogGroups
}
