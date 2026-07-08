package routers

import (
	"github.com/Leon1235532/Go_backend/handler"
	"github.com/Leon1235532/Go_backend/setting"
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
		v1group.DELETE("/todo/:id", handler.DeleteHandler)
	}
	return r
}
