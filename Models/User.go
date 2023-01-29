package User

import (
	Init "blog-go/Config"
	"log"
)

type User struct {
	ID       uint   `gorm:"primary_key"`
	Username string `gorm:"username" `
	Password string `gorm:"password" `
}

func LoginUser(username, password string) User {
	var user User
	err := Init.DB.Where("username=?", username).Where("password=?", password).First(&user)
	if err != nil {
		log.Println("get user by id fail : ", err)
	}
	return user
}
