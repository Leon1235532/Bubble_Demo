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
	if err := dao.UpdateTodo(id, &todo); err != nil {
		common.ErrorResponse(c, "", err.Error())
		return
	}
	common.SucessResponse(c, "", todo)
}

func ReviewHandler(c *gin.Context) {
	var (
		todolist   []models.Todo
		totalpages uint64
		divipage   schemas.Pagination
		err        error
	)
	if err = c.ShouldBindJSON(&divipage); err != nil {
		common.ErrorResponse(c, common.ParaErrMsg, err.Error())
		return
	}
	todolist, totalpages, err = dao.ReviewTodo(&divipage)
	message := fmt.Sprintf("共%d页,现第%d页", totalpages, divipage.Page)
	common.ResSuccMsgJson(c, message, todolist, err)
}

func DeleteHandler(c *gin.Context) {
	ids := new(schemas.IDsPara)
	if err := c.ShouldBindJSON(ids); err != nil {
		common.ErrorResponse(c, common.ParaErrMsg, err.Error())
		return
	}
	count, err := dao.DeleteTodo(ids)
	message := fmt.Sprintf("共软删除%d条数据", count)
	common.ResSuccMsgJson(c, message, nil, err)
}

// Restore Recycle

func ReviewRecyHandler(c *gin.Context) {
	var (
		todolist   []models.Todo
		totalpages uint64
		divipage   schemas.Pagination
		err        error
	)
	if err = c.ShouldBindJSON(&divipage); err != nil {
		common.ErrorResponse(c, common.ParaErrMsg, err.Error())
		return
	}
	todolist, totalpages, err = dao.ReviewRecycle(&divipage)
	message := fmt.Sprintf("共%d页,现第%d页", totalpages, divipage.Page)
	common.ResSuccMsgJson(c, message, todolist, err)
}

func RtorRecyHandler(c *gin.Context) {
	ids := new(schemas.IDsPara)
	if err := c.ShouldBindJSON(ids); err != nil {
		common.ErrorResponse(c, common.ParaErrMsg, err.Error())
		return
	}
	count, err := dao.RestoreRecycle(ids)
	message := fmt.Sprintf("共恢复%d条数据", count)
	common.ResSuccMsgJson(c, message, nil, err)
}

func RtorAllRecHandler(c *gin.Context) {
	count, err := dao.RestoreAllRecycle()
	message := fmt.Sprintf("共恢复%d条数据", count)
	common.ResSuccMsgJson(c, message, nil, err)
}

func EmptyRecyHandler(c *gin.Context) {
	ids := new(schemas.IDsPara)
	if err := c.ShouldBindJSON(ids); err != nil {
		common.ErrorResponse(c, common.ParaErrMsg, err.Error())
		return
	}
	count, err := dao.EmptyRecycle(ids)
	message := fmt.Sprintf("彻底清空%d条数据", count)
	common.ResSuccMsgJson(c, message, nil, err)
}

func EmptyAllRecyHandler(c *gin.Context) {
	count, err := dao.EmptyAllRecycle()
	message := fmt.Sprintf("彻底清空%d条数据", count)
	common.ResSuccMsgJson(c, message, nil, err)
}
