package dao

import (
	"github.com/Leon1235532/GoTask/common"
	"github.com/Leon1235532/GoTask/models"
	"github.com/Leon1235532/GoTask/schemas"
)

// CRUD

func CreateTodo(todo *models.Todo) (err error) {
	err = DB.Create(todo).Error
	return err
}

func UpdateTodo(uid uint, id string, todo *models.Todo) (err error) {
	if err = DB.Model(&models.Todo{}).
		Where("id = ? AND uid = ?", id, uid).
		Updates(todo).Error; err != nil {
		return
	}
	return DB.First(todo, id).Error
}

func ReviewTodo(uid uint, divpage *schemas.Pagination) (todolist []models.Todo, totalpages uint64, err error) {
	if divpage.Page == 0 {
		divpage.Page = 1
	}

	if divpage.PageSize == 0 {
		divpage.PageSize = 3
	}

	var count int64
	if err = DB.Model(&models.Todo{}).Where("uid = ?", uid).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	totalpages = (uint64(count) + divpage.PageSize - 1) / divpage.PageSize

	if totalpages < divpage.Page {
		divpage.Page = totalpages
	}
	offset := (divpage.Page - 1) * divpage.PageSize
	err = DB.Where("uid = ?", uid).Offset(int(offset)).
		Limit(int(divpage.PageSize)).
		Find(&todolist).Error

	if err != nil {
		return nil, 0, err
	}
	return
}

func DeleteTodo(uid uint, ids *schemas.IDsPara) (count int64, err error) {
	res := DB.Where("id in ? AND uid = ?", ids.IDs, uid).
		Delete(&models.Todo{})
	return common.RetCountErr(res)
}

// Restore & Empty

func ReviewRecycle(uid uint, divpage *schemas.Pagination) (todolist []models.Todo, totalpages uint64, err error) {
	if divpage.Page == 0 {
		divpage.Page = 1
	}

	if divpage.PageSize == 0 {
		divpage.PageSize = 3
	}

	var count int64
	if err = DB.Unscoped().
		Model(&models.Todo{}).
		Where("deleted_at IS NOT NULL AND uid = ?", uid).
		Count(&count).Error; err != nil {
		return nil, 0, err
	}
	totalpages = (uint64(count) + divpage.PageSize - 1) / divpage.PageSize

	if totalpages < divpage.Page {
		divpage.Page = totalpages
	}
	offset := (divpage.Page - 1) * divpage.PageSize

	if err = DB.Unscoped().
		Where("deleted_at IS NOT NULL AND uid = ?", uid).
		Offset(int(offset)).
		Limit(int(divpage.PageSize)).
		Find(&todolist).Error; err != nil {
		return nil, 0, err
	}
	return
}

func RestoreRecycle(uid uint, ids *schemas.IDsPara) (count int64, err error) {
	res := DB.Unscoped().
		Model(&models.Todo{}).
		Where("id in ? AND uid = ?", ids.IDs, uid).
		Where("deleted_at is NOT NULL").
		Updates(map[string]any{
			"deleted_at": nil,
		})
	return common.RetCountErr(res)
}

func RestoreAllRecycle(uid uint) (count int64, err error) {
	res := DB.Unscoped().
		Model(&models.Todo{}).
		Where("deleted_at is NOT NULL AND uid = ?", uid).
		Updates(map[string]any{
			"deleted_at": nil,
		})
	return common.RetCountErr(res)
}

func EmptyRecycle(uid uint, ids *schemas.IDsPara) (count int64, err error) {
	res := DB.Unscoped().
		Where("id in ? AND uid = ?", ids.IDs, uid).
		Where("deleted_at is NOT NULL").
		Delete(&models.Todo{})
	return common.RetCountErr(res)
}

func EmptyAllRecycle(uid uint) (count int64, err error) {
	res := DB.Unscoped().
		Where("deleted_at is NOT NULL AND uid = ?", uid).
		Delete(&models.Todo{})
	return common.RetCountErr(res)
}
