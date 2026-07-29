package routers

import (
	"github.com/Leon1235532/GoTask/setting"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	if setting.Conf.Release {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	RegisTodoRouter(r)
	RegisAuthRouter(r)
	return r
}
