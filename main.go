package main

import (
	"mytodolist/db"
	"mytodolist/models"
	"mytodolist/routers"
)

func main() {
	// 连接数据库
	err := db.InitMySql()
	if err != nil {
		panic(err)
	}

	db.MyDB.AutoMigrate(&models.Todo{})
	r := routers.SetUpRouter()
	r.Run(":3000")
}
