package handlers

import (
	"github.com/Leon1235532/Bubble_Demo/common"
	"github.com/Leon1235532/Bubble_Demo/models"
	"github.com/Leon1235532/Bubble_Demo/schemas"
	"github.com/Leon1235532/Bubble_Demo/service"
	"github.com/gin-gonic/gin"
)

func UserRegisHandler(c *gin.Context) {
	var userinfo schemas.RegisterRequest
	if err := c.ShouldBindJSON(&userinfo); err != nil {
		common.ErrorResponse(c, common.ParaErrMsg, err.Error())
		return
	}

	if err := service.UserRegister(&userinfo); err != nil {
		common.ErrorResponse(c, "", err.Error())
		return
	}

	respond := schemas.RegisterRespond{
		Username: userinfo.Username,
	}
	common.SucessResponse(c, "", respond)
}

func UserLoginHandler(c *gin.Context) {
	var userinfo schemas.LoginRequest
	if err := c.ShouldBindJSON(&userinfo); err != nil {
		common.ErrorResponse(c, common.ParaErrMsg, err.Error())
		return
	}
	tokenString, err := service.UserLogin(&userinfo)

	if err != nil {
		common.ErrorResponse(c, "", err.Error())
		return
	}

	respond := schemas.LoginRespond{
		Username: userinfo.Username,
		Token:    tokenString,
	}
	common.SucessResponse(c, "", respond)
}

func UpdatePwdHandler(c *gin.Context) {
	var pwdinfo schemas.PwdChange
	if err := c.ShouldBindJSON(&pwdinfo); err != nil {
		common.ErrorResponse(c, common.ParaErrMsg, err.Error())
		return
	}
	userid := c.GetUint("userid")

	if err := service.ModifyPwd(userid, pwdinfo); err != nil {
		common.ErrorResponse(c, "", err.Error())
		return
	}
	message := "password changed successfully! please log in again!"
	common.SucessResponse(c, message, nil)
}

func DeleteUserHandler(c *gin.Context) {
	var (
		inputpwd schemas.LogOutRequest
		user     *models.User
		err      error
	)
	if err = c.ShouldBindJSON(&inputpwd); err != nil {
		common.ErrorResponse(c, common.ParaErrMsg, err.Error())
		return
	}
	userid := c.GetUint("userid")
	if user, err = service.LogOutUser(userid, inputpwd); err != nil {
		common.ErrorResponse(c, "", err.Error())
		return
	}
	message := "Logged out successfully! Please log in again."
	common.SucessResponse(c, message, user)
}
