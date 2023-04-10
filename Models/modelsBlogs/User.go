package modelsBlogs

type User struct {
	ID       uint   `gorm:"primary_key"`
	Username string `gorm:"user_name" `
	Password string `gorm:"pass_word" `
}
