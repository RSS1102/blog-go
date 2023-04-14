package modelsBlogs

import "time"

type BlogGroups struct {
	ID       uint      `gorm:"primary_key"`
	Group    string    `gorm:"group"`
	IsShow   bool      `gorm:"is_show"`
	CreateAt time.Time `gorm:"column:create_at"`
	UpdateAt time.Time `gorm:"column:update_at"`
}
