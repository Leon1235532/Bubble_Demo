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

func ReviewTodo() (todolist *[]Todo, err error) {
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

// Restore Recycle Bin
func RestoreATodo(id string) (err error) {
	err = dao.DB.Unscoped().Model(&Todo{}).
		Where("id =? AND deleted_at IS NOT NULL", id).
		Updates(map[string]any{
			"deleted_at": nil,
		}).Error
	return
}

func ReviewRecycle() (count int64, todolist []Todo, err error) {
	result := dao.DB.Unscoped().
		Where("deleted_at IS NOT NULL").
		Order("ID asc").
		Find(&todolist)
	if result.Error != nil {
		return 0, []Todo{}, err
	}
	count = result.RowsAffected
	return
}

func RestoreAllTodo() (count int64, err error) {
	result := dao.DB.Unscoped().Model(&Todo{}).
		Where("deleted_at IS NOT NULL").
		Updates(map[string]any{
			"deleted_at": nil,
		})
	if err = result.Error; err != nil {
		return 0, err
	}
	count = result.RowsAffected
	return
}

func EmptyAllRecycle() (count int64, err error) {
	result := dao.DB.Unscoped().
		Where("deleted_at IS NOT NULL").
		Delete(&Todo{})
	if result.Error != nil {
		return 0, err
	}
	count = result.RowsAffected
	return
}

func EmptyARecycle(id string) (err error) {
	err = dao.DB.Unscoped().
		Where("id =? AND deleted_at IS NOT NULL", id).
		Delete(&Todo{}).Error
	return
}
