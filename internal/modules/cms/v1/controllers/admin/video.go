package admin

import (
	"encoding/json"
	"gin/app/models"
	"gin/global"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gconv"
	"github.com/golang-module/carbon"
)

type VideoController struct{}

func (*VideoController) Items(c *gin.Context) {
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
	model := global.DB.Model(&models.Video{})
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
	for k, v := range result {
		var tagArr []int
		_ = json.Unmarshal(gconv.Bytes(v["tag_ids"]), &tagArr)
		result[k]["tag_ids"] = tagArr
	}
	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", map[string]interface{}{"items": result, "total": count})
}

func (*VideoController) Info(c *gin.Context) {
	var param struct {
		Id int `form:"id" json:"id" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
		return
	}
	result := map[string]interface{}{}
	global.DB.Model(&models.Video{}).Where("status != ?", 9).First(&result, param.Id)
	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", result)
}

func (*VideoController) Create(c *gin.Context) {
	var param struct {
		SiteId      int    `form:"site_id" json:"site_id" binding:"required"`
		Status      int    `form:"status" json:"status" binding:"required"`
		Subject     string `form:"subject" json:"subject" binding:"required"`
		Url         string `form:"url" json:"url" binding:"required"`
		Thumbnail   string `form:"thumbnail" json:"thumbnail" binding:"required"`
		Description string `form:"description" json:"description" binding:"required"`
		Tags        string `form:"tags" json:"tags"` // 旧的
		TagIds      []int  `form:"tag_ids" json:"tag_ids" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
		return
	}
	var info models.Video
	global.DB.Model(&models.Video{}).Where("url = ?", param.Url).First(&info)
	if info.ID != 0 {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, "已经存在的url", "")
		return
	}

	tagJson, err := json.Marshal(param.TagIds)
	if err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "tag 序列化失败")
		return
	}

	data := map[string]interface{}{
		"site_id":     param.SiteId,
		"status":      param.Status,
		"subject":     param.Subject,
		"url":         param.Url,
		"thumbnail":   param.Thumbnail,
		"description": param.Description,
		//"tags":        param.Tags,
		"tag_ids":     tagJson,
		"create_time": carbon.Now().Timestamp(),
	}
	result := global.DB.Model(&models.Video{}).Create(data)
	if result.Error != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, result.Error.Error(), "")
		return
	}

	// 判断tag库是否有存在
	//tagSlice := strings.Split(param.Tags, ",")
	//for _, v := range tagSlice {
	//	var tagInfo models.VideoTag
	//	global.DB.Model(&models.VideoTag{}).Where("name = ?", v).Find(&tagInfo)
	//	if tagInfo.ID == 0 {
	//		global.DB.Model(&models.VideoTag{}).Create(&models.VideoTag{Name: v, CreateTime: carbon.Now().Timestamp()})
	//	}
	//}

	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", result.RowsAffected)
}

func (*VideoController) Update(c *gin.Context) {
	var param struct {
		Id          int    `form:"id" json:"id" binding:"required"`
		SiteId      int    `form:"site_id" json:"site_id" binding:"required"`
		Status      int    `form:"status" json:"status" binding:"required"`
		Subject     string `form:"subject" json:"subject" binding:"required"`
		Url         string `form:"url" json:"url" binding:"required"`
		Thumbnail   string `form:"thumbnail" json:"thumbnail" binding:"required"`
		Description string `form:"description" json:"description" binding:"required"`
		Tags        string `form:"tags" json:"tags" `
		TagIds      []int  `form:"tag_ids" json:"tag_ids" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
		return
	}
	tagJson, err := json.Marshal(param.TagIds)
	if err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "tag 序列化失败")
		return
	}
	data := map[string]interface{}{
		"status":      param.Status,
		"subject":     param.Subject,
		"url":         param.Url,
		"thumbnail":   param.Thumbnail,
		"description": param.Description,
		//"tags":        param.Tags,
		"tag_ids":     tagJson,
		"update_time": carbon.Now().Timestamp(),
	}
	result := global.DB.Model(&models.Video{}).Where("id = ? and status !=?", param.Id, 9).Updates(data)
	if result.Error != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, result.Error.Error(), "")
		return
	}
	// 判断tag库是否有存在
	//tagSlice := strings.Split(param.Tags, ",")
	//for _, v := range tagSlice {
	//	var tagInfo models.VideoTag
	//	global.DB.Model(&models.VideoTag{}).Where("name = ?", v).Find(&tagInfo)
	//	if tagInfo.ID == 0 {
	//		global.DB.Model(&models.VideoTag{}).Create(&models.VideoTag{Name: v, CreateTime: carbon.Now().Timestamp()})
	//	}
	//}
	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", result.RowsAffected)
}

// 软删除
func (*VideoController) Delete(c *gin.Context) {
	var param struct {
		Id int `form:"id" json:"id" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
		return
	}
	result := global.DB.Model(&models.Video{}).Where("id = ?", param.Id).Updates(map[string]interface{}{
		"status":      9,
		"delete_time": carbon.Now().Timestamp(),
	})
	if result.Error != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, result.Error.Error(), "")
		return
	}
	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", result.RowsAffected)
}

func (*VideoController) WSearchItems(c *gin.Context) {
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
	model := global.DB.Model(&models.Video{})
	model.Where("status != ?", 9)
	model.Where("site_id = ?", param.SiteId)
	if param.Search != "" {
		model.Where("tags like ? or subject like ?", param.Search, param.Search)
	}
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
