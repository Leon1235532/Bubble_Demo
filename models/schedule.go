package models

import (
	"github.com/Leon1235532/Go_backend/dao"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title  string `gorm:"NOT NULL"`
	Status bool   `gorm:"DEFAULT='false'"`
}

// CRUD
func CreateTodo(todo Todo) (err error) {
	err = dao.DB.Create(&todo).Error
	return err
}

func ReviewTodo() (todolist []Todo, err error) {
	if err = dao.DB.Find(&todolist).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateTodo(id string, todo Todo) (err error) {
	err = dao.DB.Where("id = ?", id).Updates(&todo).Error
	return
}

func DeleteTodo(id string) (err error) {
	err = dao.DB.Where("id = ?", id).Delete(&Todo{}).Error
	return
}
