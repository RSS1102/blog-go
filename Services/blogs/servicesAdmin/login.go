package servicesAdmin

import (
	Init "blog-go/Config"
	User "blog-go/Models/modelAdmin"
	"log"
)

func SelectUser(username, password string) User.User {
	var user User.User
	err := Init.DB.Table("users").Where("username=?", username).Where("password=?", password).First(&user)
	if err != nil {
		log.Println("get user by id fail : ", err)
	}
	return user
}
