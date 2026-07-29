package routers

import (
	"github.com/Leon1235532/GoTask/auth"
	"github.com/Leon1235532/GoTask/handlers"
	"github.com/gin-gonic/gin"
)

func RegisAuthRouter(r *gin.Engine) {
	v1 := r.Group("/auth")
	{
		v1.POST("/register", handlers.UserRegisHandler)
		v1.POST("/login", handlers.UserLoginHandler)
		v1.PUT("modifypwd", auth.ExtractUserId(), handlers.UpdatePwdHandler)
		v1.DELETE("logout", auth.ExtractUserId(), handlers.DeleteUserHandler)
	}
}
