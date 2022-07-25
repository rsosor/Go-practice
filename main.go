package main

import (
	"go-hello/go/dao"
	"go-hello/go/entity"
	"go-hello/go/routes"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// @title Gin Swagger Demo
// @version 1.0
// @description Swagger API.
// @host localhost:8080
func main() {
	// 連線資料庫
	err := dao.InitPostgres()
	if err != nil {
		panic(err)
	}
	// 程式退出關閉資料庫連線
	defer dao.Close()
	// 繫結模型
	dao.SqlSession.AutoMigrate(&entity.User{})
	// 註冊路由
	r := routes.SetRouter()
	// 啟動 port
	r.Run(":8080")
}
