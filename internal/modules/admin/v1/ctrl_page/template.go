package ctrl_page

import (
	"gin/internal/global"
	"gin/internal/global/errcode"
	"gin/internal/modules/admin/v1/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon"
)

type CmsTemplateController struct{}

func (*CmsTemplateController) Items(c *gin.Context) {
	var param struct {
		Limit       int    `form:"limit" json:"limit"`
		PageIndex   int    `form:"pageIndex" json:"pageIndex"`
		OrderBy     string `form:"orderBy" json:"orderBy"`
		OrderByType string `form:"orderByType" json:"orderByType"`
		Search      string `form:"search" json:"search"`
		SiteId      int    `form:"site_id" json:"site_id" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS,err.Error())
		return
	}
	model := global.DB.Model(&models.CmsTemplate{})
	model = WhereBySearch(model, param.Search)
	model.Where("status != ?", 9)
	model.Where("site_id = ?", param.SiteId)
	var count int64
	model.Count(&count)

	if param.Limit != 0 {
		if param.PageIndex == 0 {
			param.PageIndex = 1
		}
		model.Offset((param.PageIndex - 1) * param.Limit).Limit(param.Limit)
	}
	if param.OrderBy != "" && param.OrderByType != "" {
		model.Order(param.OrderBy + " " + param.OrderByType)
	} else {
		model.Order("id desc")
	}
	var result []map[string]interface{}
	model.Find(&result)
	global.Response.Success(c, map[string]interface{}{"items": result, "total": count})
}

func (*CmsTemplateController) Info(c *gin.Context) {
	var param struct {
		Id int `form:"id" json:"id" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS,err.Error())
		return
	}
	result := map[string]interface{}{}
	global.DB.Model(&models.CmsTemplate{}).Where("status != ?", 9).First(&result, param.Id)
	global.Response.Success(c, result)
}

func (*CmsTemplateController) Create(c *gin.Context) {
	var param struct {
		SiteId       int    `form:"site_id" json:"site_id" binding:"required"`
		Status       int    `form:"status" json:"status" binding:"required"`
		TemplateName string `form:"template_name" json:"template_name" binding:"required"`
		Type         string `form:"type" json:"type" binding:"required"`
		ModuleIds    string `form:"module_ids" json:"module_ids"`
		Content      string `form:"content" json:"content" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS,err.Error())
		return
	}
	data := map[string]interface{}{
		"site_id":       param.SiteId,
		"status":        param.Status,
		"template_name": param.TemplateName,
		"type":          param.Type,
		"module_ids":    param.ModuleIds,
		"content":       param.Content,
		"create_time":   carbon.Now().Timestamp(),
	}
	result := global.DB.Model(&models.CmsTemplate{}).Create(data)
	if result.Error != nil {
		global.Response.Error(c, errcode.ERROR_SERVER,result.Error.Error())
		return
	}
	global.Response.Success(c,  result.RowsAffected)
}

func (*CmsTemplateController) Update(c *gin.Context) {
	var param struct {
		Id           int    `form:"id" json:"id" binding:"required"`
		Status       int    `form:"status" json:"status" binding:"required"`
		TemplateName string `form:"template_name" json:"template_name" binding:"required"`
		Type         string `form:"type" json:"type" binding:"required"`
		ModuleIds    string `form:"module_ids" json:"module_ids" binding:"required"`
		Content      string `form:"content" json:"content" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS,err.Error())
		return
	}
	data := map[string]interface{}{
		"status":        param.Status,
		"template_name": param.TemplateName,
		"module_ids":    param.ModuleIds,
		"type":          param.Type,
		"content":       param.Content,
		"update_time":   carbon.Now().Timestamp(),
	}
	result := global.DB.Model(&models.CmsTemplate{}).Where("id = ? and status !=?", param.Id, 9).Updates(data)
	if result.Error != nil {
		global.Response.Error(c, errcode.ERROR_SERVER, result.Error.Error())
		return
	}
	global.Response.Success(c,result.RowsAffected)
}

// 软删除
func (*CmsTemplateController) Delete(c *gin.Context) {
	var param struct {
		Id int `form:"id" json:"id" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS,err.Error())
		return
	}
	result := global.DB.Model(&models.CmsTemplate{}).Where("id = ?", param.Id).Updates(map[string]interface{}{
		"status":      9,
		"delete_time": carbon.Now().Timestamp(),
	})
	if result.Error != nil {
		global.Response.Error(c, errcode.ERROR_SERVER, result.Error.Error())
		return
	}
	global.Response.Success(c,  result.RowsAffected)
}
