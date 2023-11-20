package handler

import (
	"fmt"
	"gin/api/admin/identify/v1"
	"gin/internal/global"
	"gin/internal/global/errcode"
	"gin/internal/modules/admin/v1/models"
	"gin/internal/modules/admin/v1/service"
	"github.com/gin-gonic/gin"
)

type IdentifyHandler struct{}

func (*IdentifyHandler) Items(c *gin.Context) {
	var param v1.ItemReq
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS, fmt.Sprintf("param err: %v", err))
		return
	}
	result := service.Identify().Items(param)
	global.Response.Success(c, result)
}

func (*IdentifyHandler) Info(c *gin.Context) {
	var param v1.InfoReq
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS, fmt.Sprintf("param err: %v", err))
		return
	}
	rs, err := service.Identify().Info(param)
	if err != nil {
		global.Response.Error(c, errcode.ERROR_SERVER, err.Error())
		return
	}
	global.Response.Success(c, rs)
}

// "father_identify_id": 0,
func (*IdentifyHandler) Create(c *gin.Context) {
	var param v1.CreateReq
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS, fmt.Sprintf("param err: %v", err))
		return
	}

	affected, err := service.Identify().Create(param)
	if err != nil {
		global.Response.Error(c, errcode.ERROR_SERVER, err.Error())
		return
	}
	global.Response.Success(c, affected)
}

func (*IdentifyHandler) InitCreate(c *gin.Context) {
	var param v1.CreateReq
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS, fmt.Sprintf("param err: %v", err))
		return
	}
	param.Type = int8(models.IDENTIFY_TYPE_SYSTEM)
	affected, err := service.Identify().Create(param)
	if err != nil {
		global.Response.Error(c, errcode.ERROR_SERVER, err.Error())
		return
	}
	global.Response.Success(c, affected)
}

func (*IdentifyHandler) Update(c *gin.Context) {
	var param v1.UpdateReq

	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS, fmt.Sprintf("param err: %v", err))
		return
	}

	affected, err := service.Identify().Update(param)
	if err != nil {
		global.Response.Error(c, errcode.ERROR_SERVER, err.Error())
		return
	}
	global.Response.Success(c, affected)
}

// 软删除
func (*IdentifyHandler) Delete(c *gin.Context) {
	var param v1.DeleteReq
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS, fmt.Sprintf("param err: %v", err))
		return
	}
	affected, err := service.Identify().Delete(param)
	if err != nil {
		global.Response.Error(c, errcode.ERROR_SERVER, err.Error())
		return
	}
	global.Response.Success(c, affected)
}
