package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrRes struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
	Err  string `json:"err"`
}

type SuccessRes struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

type ParaRes struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
	Err  any    `json:"err"`
}

func ErrorResponse(c *gin.Context, err error) {
	res := ErrRes{
		Code: HTTPErrServer,
		Msg:  FailMsg,
	}
	if err != nil {
		res.Err = err.Error()
	}
	c.JSON(http.StatusInternalServerError, res)
	c.Abort()
}

func SuccessRespData(c *gin.Context, msg string, data any) {
	if msg == "" {
		msg = SuccessMsg
	}
	res := SuccessRes{
		Code: HTTPOk,
		Msg:  msg,
		Data: data,
	}
	c.JSON(http.StatusOK, res)
}

func ParameterError(c *gin.Context, err any) {
	res := ParaRes{
		Code: HTTPErrParam,
		Msg:  MsgParamErr,
	}
	if err != nil {
		res.Err = err
	}
	c.JSON(http.StatusBadRequest, res)
	c.Abort()
}
