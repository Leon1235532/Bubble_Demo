package dao

import "github.com/Leon1235532/Bubble_Demo/models"

// CRUD
func CreateTodo(todo models.Todo) (err error) {
	err = DB.Create(&todo).Error
	return err
}

func ReviewTodo() (todolist *[]models.Todo, err error) {
	if err = DB.Find(&todolist).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateTodo(id string, todo models.Todo) (err error) {
	err = DB.Where("id = ?", id).Updates(&todo).Error
	return
}

func DeleteTodo(ids []uint64) (count int64, err error) {
	res := DB.Where("id IN ?", ids).Delete(&models.Todo{})
	if res.Error != nil {
		return 0, res.Error
	}
	count = res.RowsAffected
	return
}

// Restore Recycle Bin
func RestoreATodo(id string) (err error) {
	err = DB.Unscoped().Model(&models.Todo{}).
		Where("id =? AND deleted_at IS NOT NULL", id).
		Updates(map[string]any{
			"deleted_at": nil,
		}).Error
	return
}

func ReviewRecycle() (count int64, todolist []models.Todo, err error) {
	res := DB.Unscoped().
		Where("deleted_at IS NOT NULL").
		Order("ID asc").
		Find(&todolist)
	if res.Error != nil {
		return 0, []models.Todo{}, res.Error
	}
	count = res.RowsAffected
	return
}

func RestoreAllTodo() (count int64, err error) {
	res := DB.Unscoped().Model(&models.Todo{}).
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
	res := DB.Unscoped().
		Where("deleted_at IS NOT NULL").
		Delete(&models.Todo{})
	if res.Error != nil {
		return 0, res.Error
	}
	count = res.RowsAffected
	return
}

func EmptyARecycle(ids []uint64) (count int64, err error) {
	res := DB.Unscoped().
		Where("id IN ? AND deleted_at IS NOT NULL", ids).
		Delete(&models.Todo{})
	if res.Error != nil {
		return 0, res.Error
	}
	count = res.RowsAffected
	return
}
