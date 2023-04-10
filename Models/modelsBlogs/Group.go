package modelsBlogs

type BlogGroups struct {
	ID     uint   `gorm:"primary_key"`
	Group  string `gorm:"group" `
	IsShow bool   `gorm:"is_show" `
}
