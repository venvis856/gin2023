package handler

import (
	"gin/api/double_admin/v1/role/v1"
	"gin/internal/global/errcode"
	"gin/internal/global/response"
	"gin/internal/modules/double_admin/v1/service"
	"github.com/gin-gonic/gin"
)

type RoleHandler struct{}

func (*RoleHandler) Items(c *gin.Context) {
	var param v1.ItemReq
	if err := c.ShouldBind(&param); err != nil {
		//global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "param err")
		response.ErrorWithData(c, errcode.ERROR_PARAMS.Wrap(err), "param err")
		return
	}

	rs, err := service.Role().Items(param)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c, rs)
}

func (*RoleHandler) Info(c *gin.Context) {
	var param v1.InfoReq
	if err := c.ShouldBind(&param); err != nil {
		//global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "param err")
		response.ErrorWithData(c, errcode.ERROR_PARAMS.Wrap(err), "param err")
		return
	}
	result := service.Role().Info(param)
	response.Success(c, result)
}

func (*RoleHandler) Create(c *gin.Context) {
	var param v1.CreateReq
	if err := c.ShouldBind(&param); err != nil {
		//global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "param err")
		response.ErrorWithData(c, errcode.ERROR_PARAMS.Wrap(err), "param err")
		return
	}
	rs, err := service.Role().Create(param)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c, rs)
}

func (*RoleHandler) Update(c *gin.Context) {
	var param v1.UpdateReq
	if err := c.ShouldBind(&param); err != nil {
		//global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "param err")
		response.ErrorWithData(c, errcode.ERROR_PARAMS.Wrap(err), "param err")
		return
	}
	rs, err := service.Role().Update(param)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c, rs)
}

// 软删除
func (*RoleHandler) Delete(c *gin.Context) {
	var param v1.DeleteReq
	if err := c.ShouldBind(&param); err != nil {
		//global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "param err")
		response.ErrorWithData(c, errcode.ERROR_PARAMS.Wrap(err), "param err")
		return
	}
	rs, err := service.Role().Delete(param)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c, rs)
}
