package handler

import (
	"fmt"
	"gin/api/admin/role/v1"
	"gin/internal/global"
	"gin/internal/global/errcode"
	"gin/internal/modules/admin/v1/service"
	"github.com/gin-gonic/gin"
)

type RoleHandler struct{}

func (*RoleHandler) Items(c *gin.Context) {
	var param v1.ItemReq
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS, fmt.Sprintf("param err: %v", err))
		return
	}

	rs, err := service.Role().Items(param)
	if err != nil {
		global.Response.Error(c, errcode.ERROR_SERVER, err.Error())
		return
	}
	global.Response.Success(c, rs)
}

func (*RoleHandler) Info(c *gin.Context) {
	var param v1.InfoReq
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS, fmt.Sprintf("param err: %v", err))
		return
	}
	result := service.Role().Info(param)
	global.Response.Success(c, result)
}

func (*RoleHandler) Create(c *gin.Context) {
	var param v1.CreateReq
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS, fmt.Sprintf("param err: %v", err))
		return
	}
	rs, err := service.Role().Create(param)
	if err != nil {
		global.Response.Error(c, errcode.ERROR_SERVER, err.Error())
		return
	}
	global.Response.Success(c, rs)
}

func (*RoleHandler) Update(c *gin.Context) {
	var param v1.UpdateReq
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS, fmt.Sprintf("param err: %v", err))
		return
	}
	rs, err := service.Role().Update(param)
	if err != nil {
		global.Response.Error(c, errcode.ERROR_SERVER, err.Error())
		return
	}
	global.Response.Success(c, rs)
}

// 软删除
func (*RoleHandler) Delete(c *gin.Context) {
	var param v1.DeleteReq
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS, fmt.Sprintf("param err: %v", err))
		return
	}
	rs, err := service.Role().Delete(param)
	if err != nil {
		global.Response.Error(c, errcode.ERROR_SERVER, err.Error())
		return
	}
	global.Response.Success(c, rs)
}
