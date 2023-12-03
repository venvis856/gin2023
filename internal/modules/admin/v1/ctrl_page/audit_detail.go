package ctrl_page

import (
	"gin/internal/global"
	"gin/internal/global/errcode"
	"gin/internal/modules/admin/v1/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon"
)

type CmsAuditDetailController struct{}

func (*CmsAuditDetailController) Items(c *gin.Context) {
	var param struct {
		Limit       int    `form:"limit" json:"limit"`
		PageIndex   int    `form:"pageIndex" json:"pageIndex"`
		OrderBy     string `form:"orderBy" json:"orderBy"`
		OrderByType string `form:"orderByType" json:"orderByType"`
		Search      string `form:"search" json:"search"`
		AuditId     int    `form:"audit_id" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS,err.Error())
		return
	}
	model := global.DB.Model(&models.CmsAuditDetail{})
	model = WhereBySearch(model, param.Search)
	model.Where("status != ?", 9)
	model.Where("audit_id = ?", param.AuditId)
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
	global.Response.Success(c,  map[string]interface{}{"items": result, "total": count})
}

// 软删除
func (*CmsAuditDetailController) Delete(c *gin.Context) {
	var param struct {
		Id int `form:"id" json:"id" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS,err.Error())
		return
	}
	result := global.DB.Model(&models.CmsAuditDetail{}).Where("id = ?", param.Id).Updates(map[string]interface{}{
		"status":      9,
		"delete_time": carbon.Now().Timestamp(),
	})
	if result.Error != nil {
		global.Response.Error(c, errcode.ERROR_SERVER, result.Error.Error())
		return
	}
	global.Response.Success(c,  result.RowsAffected)
}
