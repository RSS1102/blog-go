package modelAdmin

import "time"

type BlogBlogs struct {
	ID       uint      `json:"primary_key"`
	GroupId  int       `json:"group_id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Visitors int       `json:"visitors"`
	IsShow   bool      `json:"is_show"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

type BlogGroups struct {
	ID       uint      `gorm:"primary_key" json:"id"`
	Group    string    `json:"group"`
	IsShow   bool      `json:"is_show"`
	CreateAt time.Time `json:"column:create_at"`
	UpdateAt time.Time `json:"column:update_at"`
}

type User struct {
	ID       uint   `json:"primary_key"`
	Username string `json:"username" `
	Password string `json:"password" `
}
