package common

import (
	"github.com/gin-gonic/gin"
)

func ResSuccMsgJson(c *gin.Context, message string, data any, err error) {
	if err != nil {
		ErrorResponse(c, "", err.Error())
		return
	}
	SucessResponse(c, message, data)
}
