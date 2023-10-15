package service

import (
	"gin/api/double_admin/v1/login/v1"
	"gin/internal/library/handlePanic"
	"github.com/gin-gonic/gin"
)

type LoginInterface interface {
	Login(c *gin.Context, param v1.LoginReq) (map[string]interface{}, error)
	UserInfo(c *gin.Context) (v1.LoginInfo, error)
}

var loginObj LoginInterface

func Login() LoginInterface {
	if loginObj == nil {
		handlePanic.Panic("login service panic")
	}
	return loginObj
}

func RegisterLogin(i LoginInterface) {
	loginObj = i
}
