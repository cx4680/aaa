package result

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	success = "success"
	failure = "failure"
)

type Response struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(context *gin.Context, data interface{}) {
	context.JSON(http.StatusOK, Response{
		Code: success,
		Msg:  "",
		Data: data,
	})
}

func Failure(context *gin.Context, status int, msg string) {
	context.JSON(status, Response{
		Code: failure,
		Msg:  msg,
	})
}
