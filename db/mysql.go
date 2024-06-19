package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	MyDB *gorm.DB //数据库变量
)

const (
	Dsn string = "root:gh20030629@tcp(localhost:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local&timeout=10s" //数据库连接地址
)

// 数据库初始化
func InitMySql() (err error) {
	MyDB, err = gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	if err != nil {
		return
	}
	tmp, _ := MyDB.DB()
	// 测试数据库连通性
	return tmp.Ping()
}
