package handler

import (
	"fmt"
	"gin/api/admin/permission/v1"
	"gin/internal/global"
	"gin/internal/global/errcode"
	"gin/internal/modules/admin/v1/service"
	"github.com/gin-gonic/gin"
)

type PermissionHandler struct{}

func (*PermissionHandler) Items(c *gin.Context) {
	var param v1.ItemReq
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS, fmt.Sprintf("param err: %v", err))
		return
	}
	rs, err := service.Permission().Items(param)
	if err != nil {
		global.Response.Error(c, errcode.ERROR_SERVER, err.Error())
		return
	}
	global.Response.Success(c, rs)
}

// 新增权限
func (*PermissionHandler) Create(c *gin.Context) {
	var param v1.CreateReq
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS, fmt.Sprintf("param err: %v", err))
		return
	}
	rs, err := service.Permission().Create(param)
	if err != nil {
		global.Response.Error(c, errcode.ERROR_SERVER, err.Error())
		return
	}
	global.Response.Success(c, rs)
}

func (*PermissionHandler) Update(c *gin.Context) {
	var param v1.UpdateReq
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS, fmt.Sprintf("param err: %v", err))
		return
	}
	rs, err := service.Permission().Update(param)
	if err != nil {
		global.Response.Error(c, errcode.ERROR_SERVER, err.Error())
		return
	}
	global.Response.Success(c, rs)
}

// 删除权限
func (*PermissionHandler) Delete(c *gin.Context) {
	var param v1.DeleteReq
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS, fmt.Sprintf("param err: %v", err))
		return
	}
	rs, err := service.Permission().Delete(c, param)
	if err != nil {
		global.Response.Error(c, errcode.ERROR_SERVER, err.Error())
		return
	}
	global.Response.Success(c, rs)
}
