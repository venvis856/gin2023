package ctrl_admin

import (
	"fmt"
	"gin/api/admin/user/v1"
	"gin/internal/global"
	"gin/internal/global/errcode"
	"gin/internal/modules/admin/v1/service"
	"github.com/gin-gonic/gin"
)

type UserCtrl struct{}

func (a *UserCtrl) Items(c *gin.Context) {
	var param v1.ItemReq

	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS, fmt.Sprintf("param err: %v", err))
		return
	}
	rs, err := service.User().Items(param)
	if err != nil {
		global.Response.Error(c, errcode.ERROR_SERVER, err.Error())
		return
	}
	global.Response.Success(c, rs)
}

func (*UserCtrl) Info(c *gin.Context) {
	var param v1.InfoReq
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS, fmt.Sprintf("param err: %v", err))
		return
	}
	rs, err := service.User().Info(param)
	if err != nil {
		global.Response.Error(c, errcode.ERROR_SERVER, err.Error())
		return
	}
	global.Response.Success(c, rs)
}

func (*UserCtrl) Create(c *gin.Context) {
	var param v1.CreateReq
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS, fmt.Sprintf("param err: %v", err))
		return
	}

	rs, err := service.User().Create(param)
	if err != nil {
		global.Response.Error(c, errcode.ERROR_SERVER, err.Error())
		return
	}
	global.Response.Success(c, rs)
}

func (*UserCtrl) Update(c *gin.Context) {
	var param v1.UpdateReq
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS, fmt.Sprintf("param err: %v", err))
		return
	}
	rs, err := service.User().Update(param)
	if err != nil {
		global.Response.Error(c, errcode.ERROR_SERVER, err.Error())
		return
	}
	global.Response.Success(c, rs)
}

// 软删除
func (*UserCtrl) Delete(c *gin.Context) {
	var param v1.DeleteReq
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS, fmt.Sprintf("param err: %v", err))
		return
	}

	rs, err := service.User().Delete(param)
	if err != nil {
		global.Response.Error(c, errcode.ERROR_SERVER, err.Error())
		return
	}
	global.Response.Success(c, rs)
}

func (a *UserCtrl) GetSecret(c *gin.Context) {
	var param v1.SecretReq
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS, fmt.Sprintf("param err: %v", err))
		return
	}
	rs := service.User().GetSecret(param.Pwd)
	global.Response.Success(c, rs)
}

// 硬删除
//func (*UserHandler) Delete(c *gin.Context) {
//	var param struct {
//		Id int `form:"id" binding:"required"`
//	}
//	if err := c.ShouldBind(&param); err != nil {
//		global.Response.Error(c,errcode.ERROR_PARAMS,fmt.Sprintf("param err: %v",err))
//		return
//	}
//	result:=global.DB.Unscoped().Where("ID = ?", param.Id).Delete(&models.User{})
//	if result.Error != nil{
//		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, result.Error.Error(), "")
//		return
//	}
//	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", "success")
//}
