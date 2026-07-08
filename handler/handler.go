package handler

import (
	"net/http"

	"github.com/Leon1235532/Go_backend/models"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := models.CreateTodo(todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  500,
			"msg":   "failed",
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": todo,
	})

}

func ReviewHandler(c *gin.Context) {
	todo_list, err := models.ReviewTodo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  500,
			"msg":   "failed",
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, todo_list)
}

func DeleteHandler(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的id!"})
		return
	}
	if err := models.DeleteTodo(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "id:" + id + " deleted",
	})

}

func UpdateHandler(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的id!"})
		return
	}
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.UpdateTodo(id, todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": todo})
}
