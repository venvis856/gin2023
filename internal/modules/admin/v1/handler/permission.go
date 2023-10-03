package handler

import (
	v1 "gin/api/permission/v1"
	"gin/internal/global/errcode"
	"gin/internal/global/response"
	"gin/internal/modules/admin/v1/service"
	"github.com/gin-gonic/gin"
)

type PermissionHandler struct{}

func (*PermissionHandler) Items(c *gin.Context) {
	var param v1.ItemReq
	if err := c.ShouldBind(&param); err != nil {
		response.ErrorWithData(c, errcode.ERROR_PARAMS.Wrap(err), "param err")
		return
	}
	rs, err := service.Permission().Items(param)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c, rs)
}

// 新增权限
func (*PermissionHandler) Create(c *gin.Context) {
	var param v1.CreateReq
	if err := c.ShouldBind(&param); err != nil {
		//global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "param err")
		response.ErrorWithData(c, errcode.ERROR_PARAMS.Wrap(err), "param err")
		return
	}
	rs, err := service.Permission().Create(param)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c, rs)
}

func (*PermissionHandler) Update(c *gin.Context) {
	var param v1.UpdateReq
	if err := c.ShouldBind(&param); err != nil {
		response.ErrorWithData(c, errcode.ERROR_PARAMS.Wrap(err), "param err")
		return
	}
	rs, err := service.Permission().Update(param)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c, rs)
}

// 删除权限
func (*PermissionHandler) Delete(c *gin.Context) {
	var param v1.DeleteReq
	if err := c.ShouldBind(&param); err != nil {
		response.ErrorWithData(c, errcode.ERROR_PARAMS.Wrap(err), "param err")
		return
	}
	rs, err := service.Permission().Delete(c, param)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c, rs)
}
