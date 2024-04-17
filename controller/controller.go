package controller

import (
	"Checklist/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateTask(c *gin.Context) {
	//前端填写待办任务，点击添加，发送一个请求到这
	//1.从请求中拿出数据
	t := models.Task{}
	c.ShouldBind(&t)
	//2.存入数据库
	if err := t.CreateTask(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	//3.返回响应
	c.JSON(http.StatusOK, t)
}

func GetAllTask(c *gin.Context) {
	tasks := []models.Task{}
	if err := models.GetAllTask(&tasks); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Println(tasks)
	c.JSON(http.StatusOK, tasks)
}

func UpdateTask(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "无效id",
		})
	}

	t := models.Task{}
	if err := t.GetATaskById(id); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := t.UpdateTask(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, nil)
}

func DeleteTask(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "无效id",
		})
	}

	if err := new(models.Task).DeleteTaskById(id); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, nil)
}
