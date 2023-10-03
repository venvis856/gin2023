package handler

import (
	v1 "gin/api/select/v1"
	"gin/internal/global"
	"gin/internal/global/errcode"
	"gin/internal/global/response"
	"gin/internal/modules/admin/v1/models"
	"github.com/gin-gonic/gin"
)

type SelectHandler struct{}

// 标识列表
func (*SelectHandler) GetIdentifySelectList(c *gin.Context) {
	model := global.DB.Model(&models.Identify{})
	model.Select("id,identify_name")
	model.Where("status = 1")
	var result []map[string]interface{}
	model.Find(&result)
	//global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", result)
	response.Success(c, result)
}

// 角色列表
func (*SelectHandler) GetRoleSelectList(c *gin.Context) {
	var param v1.GetRoleSelectListReq
	if err := c.ShouldBind(&param); err != nil {
		//global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "param err")
		response.ErrorWithData(c, errcode.ERROR_PARAMS.Wrap(err), "param err")
		return
	}
	model := global.DB.Model(&models.Role{})
	model.Select("id,vid,role_name,type")
	model.Where("status =1  and identify_id=?", param.IdentifyId)
	var result []map[string]interface{}
	model.Find(&result)
	//global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", result)
	response.Success(c, result)
}

// 派出所列表
func (*SelectHandler) GetPoliceIdentifySelectList(c *gin.Context) {
	var param struct {
		//IdentifyId int64 `form:"identify_id" json:"identify_id" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		//global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "param err")
		response.ErrorWithData(c, errcode.ERROR_PARAMS.Wrap(err), "param err")
		return
	}
	model := global.DB.Model(&models.Identify{})
	model.Select("id,identify_name,root,type,father_identify_id")
	model.Where("status = 1 and type=2")
	var result []map[string]interface{}
	model.Find(&result)
	//global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", result)
	response.Success(c, result)
}

func (*SelectHandler) GetUserSelectByIdentify(c *gin.Context) {
	var param v1.GetUserSelectByIdentifyReq
	if err := c.ShouldBind(&param); err != nil {
		response.ErrorWithData(c, errcode.ERROR_PARAMS.Wrap(err), "param err")
		return
	}
	model := global.DB.Model(&models.User{})
	model.Select("id,username,phone,realname")
	model.Where("status = 1 and identify_id=?", param.IdentifyId)
	var result []map[string]interface{}
	model.Find(&result)
	//global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", result)
	response.Success(c, result)
}
