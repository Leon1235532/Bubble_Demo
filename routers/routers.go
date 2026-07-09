package routers

import (
	"github.com/Leon1235532/Bubble_Demo/handler"
	"github.com/Leon1235532/Bubble_Demo/setting"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	if setting.Conf.Release {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	v1group := r.Group("v1")
	{
		v1group.POST("/todo", handler.CreateHandler)
		v1group.PUT("/todo/:id", handler.UpdateHandler)
		v1group.GET("/todo", handler.ReviewHandler)
		v1group.DELETE("/todo/delete", handler.DeleteHandler)
	}
	v2group := r.Group("v2")
	{
		v2group.GET("/todo", handler.ReviewRecycleHandler)
		v2group.PUT("/todo/:id", handler.RestoreAHandler)
		v2group.PUT("/todo", handler.RestoreAllHandler)
		v2group.PUT("/todo/empty", handler.EmptyAllRecycleHandler)
		v2group.PUT("/todo/empty/delete", handler.EmptyARecycleHandler)
	}
	return r
}
