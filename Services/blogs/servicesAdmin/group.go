package servicesAdmin

import (
	Init "blog-go/Config"
	"blog-go/Models/modelsBlogs"
	"log"
)

func CreateGroup(group string, isShow bool) int64 {
	var newGroup = modelsBlogs.BlogGroups{
		Group:  group,
		IsShow: isShow,
	}
	result := Init.DB.Create(&newGroup)
	println(result)
	if result.Error != nil {
		log.Println("Create group fail : ", result)
	}
	return result.RowsAffected
}
