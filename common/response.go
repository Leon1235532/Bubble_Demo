package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SuccessCode struct {
	Code int64  `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

type ErrCode struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
	Err  string `json:"err"`
}

func SucessResponse(c *gin.Context, msg string, data any) {
	if msg == "" {
		msg = SucMsg
	}
	res := SuccessCode{
		Code: HttpOk,
		Data: data,
		Msg:  msg,
	}
	c.JSON(http.StatusOK, res)
}

func ErrorResponse(c *gin.Context, msg string, err string) {
	if msg == "" {
		msg = FailMsg
	}
	res := ErrCode{
		Code: HttpErrPara,
		Msg:  msg,
		Err:  err,
	}
	c.JSON(http.StatusBadRequest, res)
	c.Abort()
}
