package Jwt

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func DBInit() {
	username := "root"
	password := "root"
	host := "localhost"
	port := "3306"
	DBName := "blog_db"
	//[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	//  charset=utf8 客户端字符集为utf8
	//  parseTime=true 支持把数据库datetime和date类型转换为golang的time.Time类型
	//  loc=Local 使用系统本地时区
	mysqlDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, DBName)
	db, err := gorm.Open(mysql.Open(mysqlDSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}

	DB = db
}
