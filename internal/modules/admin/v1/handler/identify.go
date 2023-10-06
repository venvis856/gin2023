package handler

import (
	"errors"
	"gin/api/admin/identify/v1"
	"gin/internal/global"
	"gin/internal/global/errcode"
	"gin/internal/global/response"
	"gin/internal/modules/admin/v1/models"
	"gin/internal/modules/admin/v1/service"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gconv"
)

type IdentifyHandler struct{}

func (*IdentifyHandler) Items(c *gin.Context) {
	var param v1.ItemReq
	if err := c.ShouldBind(&param); err != nil {
		response.ErrorWithData(c, errcode.ERROR_PARAMS.Wrap(err), "param err")
		return
	}
	result := service.Identify().Items(param)
	response.Success(c, result)
}

func (*IdentifyHandler) Info(c *gin.Context) {
	var param v1.InfoReq
	if err := c.ShouldBind(&param); err != nil {
		response.ErrorWithData(c, errcode.ERROR_PARAMS.Wrap(err), "param err")
		return
	}
	rs, err := service.Identify().Info(param)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c, rs)
}

// "father_identify_id": 0,
func (*IdentifyHandler) Create(c *gin.Context) {
	var param v1.CreateReq
	if err := c.ShouldBind(&param); err != nil {
		response.ErrorWithData(c, errcode.ERROR_PARAMS.Wrap(err), "param err")
		return
	}

	info := make(map[string]interface{})
	global.DB.Model(&models.Identify{}).Where("status != 9 and root = ?", param.Root).First(&info)
	if len(info) != 0 {
		response.Error(c, errors.New("标识符已经存在"))
		return
	}

	if param.Type == gconv.Int8(models.IDENTIFY_TYPE_HOTEL) && param.FatherIdentifyId == 0 {
		response.Error(c, errors.New("请选择对应的派出所"))
		return
	}

	affected, err := service.Identify().Create(param)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c, affected)
}

func (*IdentifyHandler) Update(c *gin.Context) {
	var param v1.UpdateReq

	if err := c.ShouldBind(&param); err != nil {
		//global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "param err")
		response.ErrorWithData(c, errcode.ERROR_PARAMS.Wrap(err), "param err")
		return
	}

	info := make(map[string]interface{})
	global.DB.Model(&models.Identify{}).Where("status != 9 and root = ? and id !=?", param.Root, param.Id).First(&info)
	if len(info) != 0 {
		response.Error(c, errors.New("标识符已经存在"))
		return
	}

	if param.Type == gconv.Int8(models.IDENTIFY_TYPE_HOTEL) && param.FatherIdentifyId == 0 {
		response.Error(c, errors.New("请选择对应的派出所"))
		return
	}

	affected, err := service.Identify().Update(param)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c, affected)
}

// 软删除
func (*IdentifyHandler) Delete(c *gin.Context) {
	var param v1.DeleteReq
	if err := c.ShouldBind(&param); err != nil {
		response.ErrorWithData(c, errcode.ERROR_PARAMS.Wrap(err), "param err")
		return
	}
	affected, err := service.Identify().Delete(param)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c, affected)
}
