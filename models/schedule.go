package models

import (
	"github.com/Leon1235532/Bubble_Demo/dao"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title  string `gorm:"NOT NULL" json:"title"`
	Status bool   `gorm:"DEFAULT='false'" json:"status"`
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

func DeleteTodo(ids []uint64) (count int64, err error) {
	res := dao.DB.Where("id IN ?", ids).Delete(&Todo{})
	if res.Error != nil {
		return 0, res.Error
	}
	count = res.RowsAffected
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
	res := dao.DB.Unscoped().
		Where("deleted_at IS NOT NULL").
		Order("ID asc").
		Find(&todolist)
	if res.Error != nil {
		return 0, []Todo{}, res.Error
	}
	count = res.RowsAffected
	return
}

func RestoreAllTodo() (count int64, err error) {
	res := dao.DB.Unscoped().Model(&Todo{}).
		Where("deleted_at IS NOT NULL").
		Updates(map[string]any{
			"deleted_at": nil,
		})
	if err = res.Error; err != nil {
		return 0, res.Error
	}
	count = res.RowsAffected
	return
}

func EmptyAllRecycle() (count int64, err error) {
	res := dao.DB.Unscoped().
		Where("deleted_at IS NOT NULL").
		Delete(&Todo{})
	if res.Error != nil {
		return 0, res.Error
	}
	count = res.RowsAffected
	return
}

func EmptyARecycle(ids []uint64) (count int64, err error) {
	res := dao.DB.Unscoped().
		Where("id IN ? AND deleted_at IS NOT NULL", ids).
		Delete(&Todo{})
	if res.Error != nil {
		return 0, res.Error
	}
	count = res.RowsAffected
	return
}
