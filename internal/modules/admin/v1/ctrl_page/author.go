package ctrl_page

import (
	"gin/internal/global"
	"gin/internal/global/errcode"
	"gin/internal/modules/admin/v1/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon"
)

type CmsAuthorController struct{}

func (*CmsAuthorController) Items(c *gin.Context) {
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
	model := global.DB.Model(&models.CmsAuthor{})
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

func (*CmsAuthorController) Info(c *gin.Context) {
	var param struct {
		Id int `form:"id" json:"id" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS,err.Error())
		return
	}
	result := map[string]interface{}{}
	global.DB.Model(&models.CmsAuthor{}).Where("status != ?", 9).First(&result, param.Id)
	global.Response.Success(c,  result)
}

func (*CmsAuthorController) Create(c *gin.Context) {
	var param struct {
		SiteId     int    `form:"site_id" json:"site_id" binding:"required"`
		Status     int    `form:"status" json:"status" binding:"required"`
		AuthorName string `form:"author_name" json:"author_name" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS,err.Error())
		return
	}
	data := map[string]interface{}{
		"site_id":     param.SiteId,
		"status":      param.Status,
		"author_name": param.AuthorName,
		"create_time": carbon.Now().Timestamp(),
	}
	result := global.DB.Model(&models.CmsAuthor{}).Create(data)
	if result.Error != nil {
		global.Response.Error(c, errcode.ERROR_SERVER, result.Error.Error())
		return
	}
	global.Response.Success(c, result.RowsAffected)
}

func (*CmsAuthorController) Update(c *gin.Context) {
	var param struct {
		Id         int    `form:"id" json:"id" binding:"required"`
		Status     int    `form:"status" json:"status" binding:"required"`
		AuthorName string `form:"author_name" json:"author_name" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS,err.Error())
		return
	}
	data := map[string]interface{}{
		"status":      param.Status,
		"author_name": param.AuthorName,
		"update_time": carbon.Now().Timestamp(),
	}
	result := global.DB.Model(&models.CmsAuthor{}).Where("id = ? and status !=?", param.Id, 9).Updates(data)
	if result.Error != nil {
		global.Response.Error(c, errcode.ERROR_SERVER, result.Error.Error())
		return
	}
	global.Response.Success(c, result.RowsAffected)
}

// 软删除
func (*CmsAuthorController) Delete(c *gin.Context) {
	var param struct {
		Id int `form:"id" json:"id" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS,err.Error())
		return
	}
	result := global.DB.Model(&models.CmsAuthor{}).Where("id = ?", param.Id).Updates(map[string]interface{}{
		"status":      9,
		"delete_time": carbon.Now().Timestamp(),
	})
	if result.Error != nil {
		global.Response.Error(c, errcode.ERROR_SERVER,result.Error.Error())
		return
	}
	global.Response.Success(c,  result.RowsAffected)
}
