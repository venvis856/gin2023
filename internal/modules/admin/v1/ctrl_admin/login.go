package ctrl_admin

import (
	"fmt"
	"gin/api/admin/login/v1"
	"gin/internal/global"
	"gin/internal/global/errcode"
	"gin/internal/modules/admin/v1/service"
	"github.com/gin-gonic/gin"
)

type LoginCtrl struct{}

func (a *LoginCtrl) Login(c *gin.Context) {
	var param v1.LoginReq
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS, fmt.Sprintf("param err: %v", err))
		return
	}
	rs, err := service.Login().Login(c, param)
	if err != nil {
		global.Response.Error(c, errcode.ERROR_SERVER, err.Error())
		return
	}
	global.Response.Success(c, rs)
}

func (a *LoginCtrl) Logout(c *gin.Context) {
	global.Response.Success(c, "")
}
