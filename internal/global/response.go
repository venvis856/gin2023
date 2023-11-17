package global

import (
	"gin/internal/global/errcode"
	"github.com/gin-gonic/gin"
)

type ResponseStruct struct{}

var Response ResponseStruct

// http 状态码
const (
	// http 返回结果
	HTTP_SUCCESS = 200
	HTTP_FAILURE = 500
)

type ResponseJson struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func (*ResponseStruct) Success(c *gin.Context, data any) {
	c.JSON(HTTP_SUCCESS, ResponseJson{
		Code: int(errcode.SUCCESS),
		Msg:  "",
		Data: data,
	})
}

func (*ResponseStruct) Error(c *gin.Context, code errcode.ErrCode, msg string) {
	c.JSON(HTTP_SUCCESS, ResponseJson{
		Code: int(code),
		Msg:  msg,
		Data: nil,
	})
}
