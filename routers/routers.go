package routers

import (
	"Checklist/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	//告诉gin框架模板文件引用的静态文件去哪找
	r.Static("/static", "static")
	//告诉gin框架去哪找模板文件
	r.LoadHTMLGlob("templates/*")

	r.GET("/", controller.IndexHandler)

	v1group := r.Group("/v1")
	{
		/****待办事项*****/
		//添加
		v1group.POST("/todo", controller.CreateTask)
		//查看
		//查看所有待办任务
		v1group.GET("/todo", controller.GetAllTask)
		//修改
		v1group.PUT("todo/:id", controller.UpdateTask)
		//删除
		v1group.DELETE("todo/:id", controller.DeleteTask)
	}

	return r
}
