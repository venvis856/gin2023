package admin

import (
	"gin/app/models"
	"gin/global"
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon"
)

type CmsTagController struct{}

func (*CmsTagController) Items(c *gin.Context) {
	var param struct {
		Limit       int    `form:"limit" json:"limit"`
		PageIndex   int    `form:"pageIndex" json:"pageIndex"`
		OrderBy     string `form:"orderBy" json:"orderBy"`
		OrderByType string `form:"orderByType" json:"orderByType"`
		Search      string `form:"search" json:"search"`
		SiteId      int    `form:"site_id" json:"site_id" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
		return
	}
	model := global.DB.Model(&models.CmsTag{})
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
	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", map[string]interface{}{"items": result, "total": count})
}

func (*CmsTagController) Info(c *gin.Context) {
	var param struct {
		Id int `form:"id" json:"id" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
		return
	}
	result := map[string]interface{}{}
	global.DB.Model(&models.CmsTag{}).Where("status != ?", 9).First(&result, param.Id)
	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", result)
}

func (*CmsTagController) Create(c *gin.Context) {
	var param struct {
		SiteId      int    `form:"site_id" json:"site_id" binding:"required"`
		Status      int    `form:"status" json:"status" binding:"required"`
		TagName string `form:"tag_name" json:"tag_name" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
		return
	}
	data := map[string]interface{}{
		"site_id":      param.SiteId,
		"status":       param.Status,
		"tag_name": param.TagName,
		"create_time":  carbon.Now().Timestamp(),
	}
	result := global.DB.Model(&models.CmsTag{}).Create(data)
	if result.Error != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, result.Error.Error(), "")
		return
	}
	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", result.RowsAffected)
}

func (*CmsTagController) Update(c *gin.Context) {
	var param struct {
		Id          int    `form:"id" json:"id" binding:"required"`
		SiteId      int    `form:"site_id" json:"site_id" binding:"required"`
		Status      int    `form:"status" json:"status" binding:"required"`
		TagName string `form:"tag_name" json:"tag_name" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
		return
	}
	data := map[string]interface{}{
		"status":       param.Status,
		"tag_name": param.TagName,
		"update_time":  carbon.Now().Timestamp(),
	}
	result := global.DB.Model(&models.CmsTag{}).Where("id = ? and status !=?", param.Id, 9).Updates(data)
	if result.Error != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, result.Error.Error(), "")
		return
	}
	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", result.RowsAffected)
}

// 软删除
func (*CmsTagController) Delete(c *gin.Context) {
	var param struct {
		Id int `form:"id" json:"id" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
		return
	}
	result := global.DB.Model(&models.CmsTag{}).Where("id = ?", param.Id).Updates(map[string]interface{}{
		"status":      9,
		"delete_time": carbon.Now().Timestamp(),
	})
	if result.Error != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, result.Error.Error(), "")
		return
	}
	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", result.RowsAffected)
}
