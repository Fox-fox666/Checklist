package models

import (
	"Checklist/dao"
	"fmt"
)

type Task struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// 关于Task的增删改查
func (t *Task) DeleteTaskById(id string) error {
	return dao.DB.Where("id=?", id).Delete(*t).Error
}

func (t *Task) GetATaskById(id string) error {
	return dao.DB.First(t, id).Error
}

func GetAllTask(tasks *[]Task) error {
	err := dao.DB.Find(tasks).Error

	fmt.Println(&tasks)
	return err
}

func (t *Task) CreateTask() error {
	return dao.DB.Create(t).Error
}

func (t *Task) UpdateTask() error {
	return dao.DB.Model(t).Update("status", !t.Status).Error
}
