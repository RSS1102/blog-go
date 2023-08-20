package modelAdmin

import "time"

//  前者代表go语言字段 后者代码前端字段

type BlogBlogs struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	GroupId     int       `json:"groupId"`
	Title       string    `json:"title"`
	ContentMd   string    `json:"contentMd"`
	ContentHtml string    `json:"contentHtml"`
	Visitors    int       `gorm:"default:0" json:"visitors"`
	IsShow      bool      `gorm:"default:true" json:"isShow"`
	CreateAt    time.Time `json:"createAt"`
	UpdateAt    time.Time `json:"updateAt"`
}

type BlogGroups struct {
	ID       uint      `gorm:"primary_key" json:"id"`
	Group    string    `json:"group"`
	IsShow   bool      `gorm:"default:true" json:"isShow"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
}

type User struct {
	ID       uint      `gorm:"primary_key" json:"id"`
	Username string    `json:"username" `
	Password string    `json:"password"`
	Account  string    `json:"account"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
}

// MergedBlogs 连表查询 IsShow为原字段,groupIsShow为输出字段
type MergedBlogs struct {
	BlogBlogs
	Group       string `json:"group"`
	GroupIsShow bool   `json:"groupIsShow"`
}
