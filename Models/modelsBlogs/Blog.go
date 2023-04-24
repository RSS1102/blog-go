package modelsBlogs

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
