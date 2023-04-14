package servicesAdmin

import (
	Init "blog-go/Config"
	"blog-go/Models/modelsBlogs"
	"log"
	"time"
)

// CreateGroup 创建分组
func CreateGroup(group string) int64 {
	now := time.Now()
	var newGroup = modelsBlogs.BlogGroups{
		Group:    group,
		IsShow:   true,
		CreateAt: now,
		UpdateAt: now,
	}

	result := Init.DB.Create(&newGroup)
	println(result)
	if result.Error != nil {
		log.Println("Create group fail : ", result)
	}
	return result.RowsAffected
}

// DeleteGroup 删除分组
func DeleteGroup(id uint) int64 {
	now := time.Now()
	result := Init.DB.Table("blog_groups").Where("id=?", id).Updates(modelsBlogs.BlogGroups{IsShow: false, UpdateAt: now})
	println(result)
	if result.Error != nil {
		log.Println("DeleteGroup group fail : ", result)
	}
	return result.RowsAffected
}

// UpdateGroup 分组更新
func UpdateGroup(id uint, group string) int64 {
	now := time.Now()
	result := Init.DB.Table("blog_groups").Where("id=?", id).Updates(modelsBlogs.BlogGroups{Group: group, UpdateAt: now})
	println(result)
	if result.Error != nil {
		log.Println("UpdateGroup group fail : ", result)
	}
	return result.RowsAffected
}

// SelectGroup 分组查询
// 第几页
func SelectGroup(paging int, pageSize int) []modelsBlogs.BlogGroups {
	var blogGroups []modelsBlogs.BlogGroups
	err := Init.DB.Table("blog_groups").Limit(pageSize).Offset((paging - 1) * pageSize).Find(&blogGroups)
	if err.Error != nil {
		log.Println("SelectContent group fail : ", err)
	}
	return blogGroups
}
