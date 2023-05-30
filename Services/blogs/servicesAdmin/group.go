package servicesAdmin

import (
	Init "blog-go/Config"
	"blog-go/Models/modelAdmin"
	"log"
	"time"
)

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
func UpdateGroup(id uint, group string) int64 {
	result := Init.DB.Table("blog_groups").Where("id=?", id).Updates(modelAdmin.BlogGroups{Group: group, UpdateAt: time.Now()})
	println(result)
	if result.Error != nil {
		log.Println("UpdateGroup group fail : ", result)
	}
	return result.RowsAffected
}

// SelectGroup 分组查询
// 第几页
func SelectGroup(Current int, pageSize int) (int, []modelAdmin.BlogGroups) {
	var blogGroups []modelAdmin.BlogGroups
	var total int64
	err := Init.DB.Table("blog_groups").Count(&total).Error
	if err != nil {
		log.Println("Failed to get total record count:", err)
		return 0, nil
	}
	errs := Init.DB.Table("blog_groups").Limit(pageSize).Offset((Current - 1) * pageSize).Find(&blogGroups)
	if errs.Error != nil {
		log.Println("SelectContent group fail : ", err)
	}
	return int(total), blogGroups
}
