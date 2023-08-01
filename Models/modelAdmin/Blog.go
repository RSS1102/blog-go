package modelAdmin

import "time"

//  前者代表go语言字段 后者代码前端字段

type BlogBlogs struct {
	ID       uint      `gorm:"primary_key" json:"id"`
	GroupId  int       `json:"groupId"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Visitors int       `json:"visitors"`
	IsShow   bool      `json:"isShow"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
}

type BlogGroups struct {
	ID       uint      `gorm:"primary_key" json:"id"`
	Group    string    `json:"group"`
	IsShow   bool      `json:"isShow"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
}

type User struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Username string `json:"username" `
	Password string `json:"password"`
	Account  string `json:"account"`
}

type MergedBlog struct {
	BlogBlogs
	BlogGroup BlogGroups `gorm:"embedded"`
}
