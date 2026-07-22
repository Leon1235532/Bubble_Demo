package dao

import (
	"github.com/Leon1235532/test/common"
	"github.com/Leon1235532/test/models"
	"github.com/Leon1235532/test/schemas"
)

// CRUD

func CreateTodo(todo *models.Todo) (err error) {
	err = DB.Create(todo).Error
	return err
}

func UpdateTodo(id string, todo *models.Todo) (err error) {
	err = DB.Model(&models.Todo{}).Where("id = ?", id).Updates(todo).Error
	return err
}

func ReviewTodo(divpage *schemas.Pagination) (todolist []models.Todo, totalpages uint64, err error) {
	if divpage.Page == 0 {
		divpage.Page = 1
	}
	if divpage.PageSize == 0 {
		divpage.PageSize = 3
	}
	var count int64
	if err = DB.Model(&models.Todo{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	totalpages = (uint64(count) + divpage.PageSize - 1) / divpage.PageSize
	offset := (divpage.Page - 1) * divpage.PageSize
	err = DB.Model(&models.Todo{}).Offset(int(offset)).
		Limit(int(divpage.PageSize)).
		Find(&todolist).Error
	if err != nil {
		return nil, 0, err
	}
	return
}

func DeleteTodo(ids *schemas.IDsPara) (count int64, err error) {
	res := DB.Where("id in ?", ids.IDs).
		Delete(&models.Todo{})
	return common.RetCountErr(res)
}

// Restore & Empty

func ReviewRecycle(divpage *schemas.Pagination) (todolist []models.Todo, totalpages uint64, err error) {
	if divpage.Page == 0 {
		divpage.Page = 1
	}
	if divpage.PageSize == 0 {
		divpage.PageSize = 3
	}
	var count int64
	if err = DB.Unscoped().
		Model(&models.Todo{}).
		Where("deleted_at IS NOT NULL").
		Count(&count).Error; err != nil {
		return nil, 0, err
	}
	totalpages = (uint64(count) + divpage.PageSize - 1) / divpage.PageSize
	offset := (divpage.Page - 1) * divpage.PageSize
	if err = DB.Unscoped().
		Model(&models.Todo{}).Where("deleted_at IS NOT NULL").
		Offset(int(offset)).
		Limit(int(divpage.PageSize)).
		Find(&todolist).Error; err != nil {
		return nil, 0, err
	}
	return
}

func RestoreRecycle(ids *schemas.IDsPara) (count int64, err error) {
	res := DB.Unscoped().
		Model(&models.Todo{}).
		Where("id in ? AND deleted_at is NOT NULL", ids.IDs).
		Updates(map[string]any{
			"deleted_at": nil,
		})
	return common.RetCountErr(res)
}

func RestoreAllRecycle() (count int64, err error) {
	res := DB.Unscoped().
		Model(&models.Todo{}).
		Where("deleted_at is NOT NULL").
		Updates(map[string]any{
			"deleted_at": nil,
		})
	return common.RetCountErr(res)
}

func EmptyRecycle(ids *schemas.IDsPara) (count int64, err error) {
	res := DB.Unscoped().
		Where("id in ? AND deleted_at is NOT NULL", ids.IDs).
		Delete(&models.Todo{})
	return common.RetCountErr(res)
}

func EmptyAllRecycle() (count int64, err error) {
	res := DB.Unscoped().
		Where("deleted_at is NOT NULL").
		Delete(&models.Todo{})
	return common.RetCountErr(res)
}
