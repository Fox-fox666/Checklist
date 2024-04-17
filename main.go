package main

import (
	"Checklist/dao"
	"Checklist/models"
	"Checklist/routers"
	_ "github.com/go-sql-driver/mysql"
)

// 请求  --->  controller ---> logic  --->  dao
//
//	控制器         逻辑处理       模型层增删改查

func main() {
	//连接数据库
	err := dao.InitDB()
	if err != nil {
		panic(err)
	}
	//程序退出关闭数据库连接
	defer dao.DB.Close()
	//模型绑定
	dao.DB.AutoMigrate(new(models.Task))

	r := routers.SetupRouter()
	r.Run(":8080")
}
