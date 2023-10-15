package handler

import (
	"gin/api/double_admin/v1/login/v1"
	"gin/internal/global/errcode"
	"gin/internal/global/response"
	"gin/internal/modules/double_admin/v1/service"
	"github.com/gin-gonic/gin"
)

type LoginHandler struct{}

func (a *LoginHandler) Login(c *gin.Context) {
	var param v1.LoginReq
	if err := c.ShouldBind(&param); err != nil {
		response.ErrorWithData(c, errcode.ERROR_PARAMS.Wrap(err), "param err")
		return
	}
	rs, err := service.Login().Login(c, param)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c, rs)
}

func (a *LoginHandler) UserInfo(ctx *gin.Context) {
	retUser, err := service.Login().UserInfo(ctx)
	if err != nil {
		response.Error(ctx, err)
		return
	}
	response.Success(ctx, retUser)
}

func (a *LoginHandler) Logout(c *gin.Context) {
	response.Success(c, "")
}
