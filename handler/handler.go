package handler

import (
	"fmt"

	"github.com/Leon1235532/Bubble_Demo/common"
	"github.com/Leon1235532/Bubble_Demo/models"
	"github.com/Leon1235532/Bubble_Demo/schemas"
	"github.com/gin-gonic/gin"
)

/*
url     --> handler  --> logic   -->    model
请求来了  -->  控制器   --> 业务逻辑  --> 模型层的增删改查

控制器 handler（门面 / 中转站）
管请求、参数、响应、路由对接，不碰数据库、不写复杂业务。

业务逻辑 service/logic（大脑）
管核心业务规则（判断、计算、流程、事务），不直接处理 HTTP 请求。

模型 model（数据层）
只管数据库 CRUD，只操作表，不懂业务。
*/

func CreateHandler(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		common.ParameterError(c, err.Error())
		return
	}
	err := models.CreateTodo(todo)
	if err != nil {
		common.ErrorResponse(c, err)
		return
	}
	common.SuccessRespData(c, "", todo)
}

func ReviewHandler(c *gin.Context) {
	todolist, err := models.ReviewTodo()
	if err != nil {
		common.ErrorResponse(c, err)
		return
	}
	common.SuccessRespData(c, "", todolist)
}

func DeleteHandler(c *gin.Context) {
	var ids schemas.IDsPara
	if err := c.ShouldBindJSON(&ids); err != nil {
		common.ParameterError(c, err.Error())
		return
	}
	count, err := models.DeleteTodo(ids.IDs)
	if err != nil {
		common.ErrorResponse(c, err)
		return
	}
	message := fmt.Sprintf("%d条待办事项已删除", count)
	common.SuccessRespData(c, message, models.Todo{})
}

func UpdateHandler(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		common.ParameterError(c, common.IdErrMsg)
		return
	}
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		common.ParameterError(c, err.Error())
		return
	}
	if err := models.UpdateTodo(id, todo); err != nil {
		common.ParameterError(c, err.Error())
		return
	}
	common.SuccessRespData(c, "", todo)
}

func RestoreAHandler(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		common.ParameterError(c, common.IdErrMsg)
		return
	}
	if err := models.RestoreATodo(id); err != nil {
		common.ErrorResponse(c, err)
		return
	}
	message := "id:" + id + " restored"
	common.SuccessRespData(c, message, models.Todo{})
}

func RestoreAllHandler(c *gin.Context) {
	count, err := models.RestoreAllTodo()
	if err != nil {
		common.ParameterError(c, err.Error())
		return
	}
	message := fmt.Sprintf("%d个待办事项已恢复!", count)
	common.SuccessRespData(c, message, models.Todo{})
}

func ReviewRecycleHandler(c *gin.Context) {
	count, list, err := models.ReviewRecycle()
	if err != nil {
		common.ErrorResponse(c, err)
		return
	}
	message := fmt.Sprintf("回收站中共%d条数据", count)
	common.SuccessRespData(c, message, list)
}

func EmptyAllRecycleHandler(c *gin.Context) {
	count, err := models.EmptyAllRecycle()
	if err != nil {
		common.ErrorResponse(c, err)
		return
	}
	message := fmt.Sprintf("已彻底删除%d条数据", count)
	common.SuccessRespData(c, message, models.Todo{})
}

func EmptyARecycleHandler(c *gin.Context) {
	var ids schemas.IDsPara
	if err := c.ShouldBindJSON(&ids); err != nil {
		common.ParameterError(c, err.Error())
	}
	count, err := models.EmptyARecycle(ids.IDs)
	if err != nil {
		common.ErrorResponse(c, err)
		return
	}
	message := fmt.Sprintf("已彻底删除%d条数据", count)
	common.SuccessRespData(c, message, models.Todo{})
}
