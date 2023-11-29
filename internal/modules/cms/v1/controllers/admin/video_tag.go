package admin

import (
	"gin/app/models"
	"gin/global"
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon"
)

type VideoTagController struct{}

func (*VideoTagController) Items(c *gin.Context) {
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
	model := global.DB.Model(&models.VideoTag{})
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

func (*VideoTagController) Info(c *gin.Context) {
	var param struct {
		Id int `form:"id" json:"id" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
		return
	}
	result := map[string]interface{}{}
	global.DB.Model(&models.VideoTag{}).Where("status != ?", 9).First(&result, param.Id)
	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", result)
}

func (*VideoTagController) Create(c *gin.Context) {
	var param struct {
		SiteId      int    `form:"site_id" json:"site_id" binding:"required"`
		Status      int    `form:"status" json:"status" binding:"required"`
		Name        string `form:"name" json:"name" binding:"required"`
		Title       string `form:"title" json:"title" binding:"required"`
		Keywords    string `form:"keywords" json:"keywords" binding:"required"`
		Description string `form:"description" json:"description" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
		return
	}
	var info models.VideoTag
	global.DB.Model(&models.VideoTag{}).Where("name = ? and site_id=?", param.Name, param.SiteId).First(&info)
	if info.ID != 0 {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, "已经存在的name", "")
		return
	}

	data := map[string]interface{}{
		"site_id":     param.SiteId,
		"status":      param.Status,
		"name":        param.Name,
		"title":       param.Title,
		"keywords":    param.Keywords,
		"description": param.Description,
		"create_time": carbon.Now().Timestamp(),
	}
	result := global.DB.Model(&models.VideoTag{}).Create(data)
	if result.Error != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, result.Error.Error(), "")
		return
	}

	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", result.RowsAffected)
}

func (*VideoTagController) Update(c *gin.Context) {
	var param struct {
		Id          int    `form:"id" json:"id" binding:"required"`
		SiteId      int    `form:"site_id" json:"site_id" binding:"required"`
		Status      int    `form:"status" json:"status" binding:"required"`
		Name        string `form:"name" json:"name" binding:"required"`
		Title       string `form:"title" json:"title" binding:"required"`
		Keywords    string `form:"keywords" json:"keywords" binding:"required"`
		Description string `form:"description" json:"description" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
		return
	}

	var info models.VideoTag
	global.DB.Model(&models.VideoTag{}).Where("name = ? and site_id=? and id != ?", param.Name, param.SiteId, param.Id).First(&info)
	if info.ID != 0 {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, "已经存在的name", "")
		return
	}

	data := map[string]interface{}{
		"site_id":     param.SiteId,
		"status":      param.Status,
		"name":        param.Name,
		"title":       param.Title,
		"keywords":    param.Keywords,
		"description": param.Description,
		"update_time": carbon.Now().Timestamp(),
	}
	result := global.DB.Model(&models.VideoTag{}).Where("id = ? and status !=? ", param.Id, 9, ).Updates(data)
	if result.Error != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, result.Error.Error(), "")
		return
	}

	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", result.RowsAffected)
}

// 软删除
func (*VideoTagController) Delete(c *gin.Context) {
	var param struct {
		Id int `form:"id" json:"id" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
		return
	}
	result := global.DB.Model(&models.VideoTag{}).Where("id = ?", param.Id).Updates(map[string]interface{}{
		"status":      9,
		"delete_time": carbon.Now().Timestamp(),
	})
	if result.Error != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, result.Error.Error(), "")
		return
	}
	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", result.RowsAffected)
}
