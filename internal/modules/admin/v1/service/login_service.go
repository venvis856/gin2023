package service

import (
	"gin/api/admin/login/v1"
	"gin/internal/library/handlePanic"
	"github.com/gin-gonic/gin"
)

type LoginInterface interface {
	Login(c *gin.Context, param v1.LoginReq) (map[string]interface{}, error)
}

var loginObj LoginInterface

func Login() LoginInterface {
	if loginObj == nil {
		handlePanic.Panic("login_service service panic")
	}
	return loginObj
}

func RegisterLogin(i LoginInterface) {
	loginObj = i
}
