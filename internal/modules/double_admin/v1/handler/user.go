package handler

import (
	"gin/api/double_admin/v1/user/v1"
	"gin/internal/global/errcode"
	"gin/internal/global/response"
	"gin/internal/modules/double_admin/v1/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct{}

func (a *UserHandler) Items(c *gin.Context) {
	var param v1.ItemReq

	if err := c.ShouldBind(&param); err != nil {
		//global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "param err")
		response.ErrorWithData(c, errcode.ERROR_PARAMS.Wrap(err), "param err")
		return
	}
	rs, err := service.User().Items(param)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c, rs)
}

func (*UserHandler) Info(c *gin.Context) {
	var param v1.InfoReq
	if err := c.ShouldBind(&param); err != nil {
		//global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "param err")
		response.ErrorWithData(c, errcode.ERROR_PARAMS.Wrap(err), "param err")
		return
	}
	rs, err := service.User().Info(param)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c, rs)
}

func (*UserHandler) Create(c *gin.Context) {
	var param v1.CreateReq
	if err := c.ShouldBind(&param); err != nil {
		//global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "param err")
		response.ErrorWithData(c, errcode.ERROR_PARAMS.Wrap(err), "param err")
		return
	}

	rs, err := service.User().Create(param)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c, rs)
}

func (*UserHandler) Update(c *gin.Context) {
	var param v1.UpdateReq
	if err := c.ShouldBind(&param); err != nil {
		//global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "param err")
		response.ErrorWithData(c, errcode.ERROR_PARAMS.Wrap(err), "param err")
		return
	}
	rs, err := service.User().Update(param)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c, rs)
}

// 软删除
func (*UserHandler) Delete(c *gin.Context) {
	var param v1.DeleteReq
	if err := c.ShouldBind(&param); err != nil {
		//global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "param err")
		response.ErrorWithData(c, errcode.ERROR_PARAMS.Wrap(err), "param err")
		return
	}

	rs, err := service.User().Delete(param)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c, rs)
}

func (a *UserHandler) GetSecret(c *gin.Context) {
	var param v1.SecretReq
	if err := c.ShouldBind(&param); err != nil {
		response.ErrorWithData(c, errcode.ERROR_PARAMS.Wrap(err), "param err")
		return
	}
	rs := service.User().GetSecret(param.Pwd)
	response.Success(c, rs)
}

// 硬删除
//func (*UserHandler) Delete(c *gin.Context) {
//	var param struct {
//		Id int `form:"id" binding:"required"`
//	}
//	if err := c.ShouldBind(&param); err != nil {
//		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
//		return
//	}
//	result:=global.DB.Unscoped().Where("ID = ?", param.Id).Delete(&models.User{})
//	if result.Error != nil{
//		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, result.Error.Error(), "")
//		return
//	}
//	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", "success")
//}
