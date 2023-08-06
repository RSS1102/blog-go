package servicesAdmin

import (
	Init "blog-go/Config"
	"blog-go/Models/modelAdmin"
	Until "blog-go/Until"
	"log"
	"reflect"
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
func UpdateGroup(updates modelAdmin.BlogGroups) int64 {
	updateMap := make(map[string]interface{})
	v := reflect.ValueOf(updates)
	for i := 0; i < v.NumField(); i++ {
		fieldName := v.Type().Field(i).Name
		fieldValue := v.Field(i).Interface()
		lowerFieldName := Until.ConvertToSnakeCase(fieldName)
		if fieldName != "CreateAt" {
			updateMap[lowerFieldName] = fieldValue
		}
	}
	result := Init.DB.Table("blog_groups").Where("id=?", updates.ID).Updates(updateMap)
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
	if Current == -1 || pageSize == -1 {
		errs := Init.DB.Table("blog_groups").Find(&blogGroups)
		if errs.Error != nil {
			return 0, nil
		}
		return int(total), blogGroups
	}
	errs := Init.DB.Table("blog_groups").Limit(pageSize).Offset((Current - 1) * pageSize).Find(&blogGroups)
	if errs.Error != nil {
		log.Println("SelectContent group fail : ", err)
	}
	return int(total), blogGroups
}
