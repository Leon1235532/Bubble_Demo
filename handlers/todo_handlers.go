package handlers

import (
	"fmt"

	"github.com/Leon1235532/Bubble_Demo/common"
	"github.com/Leon1235532/Bubble_Demo/dao"
	"github.com/Leon1235532/Bubble_Demo/models"
	"github.com/Leon1235532/Bubble_Demo/schemas"
	"github.com/gin-gonic/gin"
)

// CRUD

func CreateHandler(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		common.ErrorResponse(c, common.ParaErrMsg, err.Error())
		return
	}
	todo.UID = c.GetUint("userid")
	if err := dao.CreateTodo(&todo); err != nil {
		common.ErrorResponse(c, "", err.Error())
		return
	}
	common.SucessResponse(c, "", todo)
}

func UpdateHandler(c *gin.Context) {
	var todo models.Todo
	id, ok := c.Params.Get("id")
	if !ok {
		common.ErrorResponse(c, common.IdErrMsg, "")
		return
	}
	if err := c.ShouldBindJSON(&todo); err != nil {
		common.ErrorResponse(c, common.ParaErrMsg, err.Error())
		return
	}
	todo.UID = c.GetUint("userid")
	if err := dao.UpdateTodo(todo.UID, id, &todo); err != nil {
		common.ErrorResponse(c, "", err.Error())
		return
	}
	common.SucessResponse(c, "", todo)
}

func ReviewHandler(c *gin.Context) {
	var (
		uid        uint
		todolist   []models.Todo
		totalpages uint64
		divipage   schemas.Pagination
		err        error
	)
	if err = c.ShouldBindJSON(&divipage); err != nil {
		common.ErrorResponse(c, common.ParaErrMsg, err.Error())
		return
	}
	uid = c.GetUint("userid")
	todolist, totalpages, err = dao.ReviewTodo(uid, &divipage)
	message := fmt.Sprintf("第%d页,共%d页", divipage.Page, totalpages)
	if err != nil {
		common.ErrorResponse(c, "", err.Error())
		return
	}
	common.SucessResponse(c, message, todolist)
}

func DeleteHandler(c *gin.Context) {
	var uid uint
	ids := new(schemas.IDsPara)
	if err := c.ShouldBindJSON(ids); err != nil {
		common.ErrorResponse(c, common.ParaErrMsg, err.Error())
		return
	}
	uid = c.GetUint("userid")
	count, err := dao.DeleteTodo(uid, ids)
	message := fmt.Sprintf("共软删除%d条数据", count)
	if err != nil {
		common.ErrorResponse(c, "", err.Error())
		return
	}
	common.SucessResponse(c, message, nil)
}

// Restore Recycle

func ReviewRecyHandler(c *gin.Context) {
	var (
		uid        uint
		todolist   []models.Todo
		totalpages uint64
		divipage   schemas.Pagination
		err        error
	)
	if err = c.ShouldBindJSON(&divipage); err != nil {
		common.ErrorResponse(c, common.ParaErrMsg, err.Error())
		return
	}
	uid = c.GetUint("userid")
	todolist, totalpages, err = dao.ReviewRecycle(uid, &divipage)
	message := fmt.Sprintf("第%d页,共%d页", divipage.Page, totalpages)
	if err != nil {
		common.ErrorResponse(c, "", err.Error())
		return
	}
	common.SucessResponse(c, message, todolist)
}

func RtorRecyHandler(c *gin.Context) {
	var uid uint
	ids := new(schemas.IDsPara)
	if err := c.ShouldBindJSON(ids); err != nil {
		common.ErrorResponse(c, common.ParaErrMsg, err.Error())
		return
	}
	uid = c.GetUint("userid")
	count, err := dao.RestoreRecycle(uid, ids)
	message := fmt.Sprintf("共恢复%d条数据", count)
	if err != nil {
		common.ErrorResponse(c, "", err.Error())
		return
	}
	common.SucessResponse(c, message, nil)
}

func RtorAllRecHandler(c *gin.Context) {
	var uid uint
	uid = c.GetUint("userid")
	count, err := dao.RestoreAllRecycle(uid)
	message := fmt.Sprintf("共恢复%d条数据", count)
	if err != nil {
		common.ErrorResponse(c, "", err.Error())
		return
	}
	common.SucessResponse(c, message, nil)
}

func EmptyRecyHandler(c *gin.Context) {
	var uid uint
	ids := new(schemas.IDsPara)
	if err := c.ShouldBindJSON(ids); err != nil {
		common.ErrorResponse(c, common.ParaErrMsg, err.Error())
		return
	}
	uid = c.GetUint("userid")
	count, err := dao.EmptyRecycle(uid, ids)
	message := fmt.Sprintf("彻底清空%d条数据", count)
	if err != nil {
		common.ErrorResponse(c, "", err.Error())
		return
	}
	common.SucessResponse(c, message, nil)
}

func EmptyAllRecyHandler(c *gin.Context) {
	var uid uint
	uid = c.GetUint("userid")
	count, err := dao.EmptyAllRecycle(uid)
	message := fmt.Sprintf("彻底清空%d条数据", count)
	if err != nil {
		common.ErrorResponse(c, "", err.Error())
		return
	}
	common.SucessResponse(c, message, nil)
}
