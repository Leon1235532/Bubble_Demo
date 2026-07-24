package routers

import (
	"github.com/Leon1235532/Bubble_Demo/handlers"
	"github.com/Leon1235532/Bubble_Demo/setting"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	if setting.Conf.Release {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	// CRUD router
	v1 := r.Group("/v1")
	{
		v1.POST("/todo/review", handlers.ReviewHandler)
		v1.POST("/todo", handlers.CreateHandler)
		v1.PUT("/todo/:id", handlers.UpdateHandler)
		v1.DELETE("/todo/delete", handlers.DeleteHandler)
	}
	// Restore & Empty
	v2 := r.Group("/v2")
	{
		v2.POST("/todo/review", handlers.ReviewRecyHandler)
		v2.PUT("/todo/recycle", handlers.RtorRecyHandler)
		v2.PUT("/todo/recycle/all", handlers.RtorAllRecHandler)
		v2.DELETE("/todo/empty", handlers.EmptyRecyHandler)
		v2.DELETE("/todo/empty/all", handlers.EmptyAllRecyHandler)
	}
	return r
}
