package servicesAdmin

import (
	Init "blog-go/Config"
	"blog-go/Models/modelsBlogs"
	"log"
)

// CreateGroup 创建分组
func CreateGroup(group string) int64 {
	var newGroup = modelsBlogs.BlogGroups{
		Group:  group,
		IsShow: true,
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
	result := Init.DB.Table("blog_groups").Where("id=?", id).Update("is_show", false)
	println(result)
	if result.Error != nil {
		log.Println("DeleteGroup group fail : ", result)
	}
	return result.RowsAffected
}

// UpdateGroup 分组更新
func UpdateGroup(id uint, group string) int64 {
	result := Init.DB.Table("blog_groups").Where("id=?", id).Update("group", group)
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
		log.Println("SelectGroup group fail : ", err)
	}
	return blogGroups
}
