package modelAdmin

import "time"

type BlogBlogs struct {
	ID       uint      `gorm:"primary_key"`
	GroupId  int       `gorm:"group_id"`
	Title    string    `gorm:"title"`
	Content  string    `gorm:"content"`
	Visitors int       `gorm:"visitors"`
	IsShow   bool      `gorm:"is_show"`
	CreateAt time.Time `gorm:"create_at"`
	UpdateAt time.Time `gorm:"update_at"`
}

type BlogGroups struct {
	ID       uint      `gorm:"primary_key"`
	Group    string    `gorm:"group"`
	IsShow   bool      `gorm:"is_show"`
	CreateAt time.Time `gorm:"column:create_at"`
	UpdateAt time.Time `gorm:"column:update_at"`
}

type User struct {
	ID       uint   `gorm:"primary_key"`
	Username string `gorm:"user_name" `
	Password string `gorm:"pass_word" `
}
