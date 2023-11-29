package admin

import (
	"encoding/json"
	"fmt"
	"gin/app/library/helper"
	"gin/app/library/jwt"
	"gin/app/models"
	"gin/app/services"
	"gin/global"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gconv"
	"strconv"
)

type CmsMakeController struct{}

// 1.文章生成  2.模块生成 3.模板生成

// 生成记录  生成文件
func (*CmsMakeController) PageMake(c *gin.Context) {
	var param struct {
		Id          int   `form:"id" json:"id" binding:"required"`
		SiteId      int   `form:"site_id" json:"site_id" binding:"required"`
		VideoOffset int64 `form:"video_offset" json:"video_offset"`
		VideoLimit  int64 `form:"video_limit" json:"video_limit"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
		return
	}
	ids := []int{param.Id}
	makeService := new(services.CmsMakeServer)
	// 生成文件
	auditList, _, err := makeService.Make(services.MakeTypePage, ids, gconv.Int(param.SiteId), false, services.VideoLimitSt{
		OffSet: param.VideoOffset,
		Limit:  param.VideoLimit,
	})
	if err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "生成文件失败")
		return
	}
	//user_id
	userId := 0
	token := c.Request.Header.Get("token")
	tokenInfo, err := jwt.ParseJwtGoToken(token)
	if err == nil {
		userId = gconv.Int(tokenInfo.Id)
	}
	// 生成记录
	auditServer := new(services.CmsAuditServer)
	auditId, err := auditServer.MakeAudit(auditList, 1, userId, gconv.Int(param.SiteId))
	if err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "写入记录失败")
		return
	}
	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", auditId)
}

// 模板生成
func (*CmsMakeController) TemplateMake(c *gin.Context) {
	var param struct {
		Id          int   `form:"id" json:"id" binding:"required"`
		SiteId      int   `form:"site_id" json:"site_id" binding:"required"`
		VideoOffset int64 `form:"video_offset" json:"video_offset"`
		VideoLimit  int64 `form:"video_limit" json:"video_limit"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
		return
	}
	var pageList []map[string]interface{}
	global.DB.Model(&models.CmsPage{}).Where("site_id=? and template_id=? and status not in ?", param.SiteId, param.Id, []int{5, 9}).Find(&pageList)
	var ids []int
	for _, v := range pageList {
		ids = append(ids, gconv.Int(v["id"]))
	}
	// 生成
	makeService := new(services.CmsMakeServer)
	// 生成文件
	auditList, _, err := makeService.Make(services.MakeTypePage, ids, gconv.Int(param.SiteId), false, services.VideoLimitSt{
		OffSet: param.VideoOffset,
		Limit:  param.VideoLimit,
	})
	if err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "生成文件失败")
		return
	}
	//user_id
	userId := 0
	token := c.Request.Header.Get("token")
	tokenInfo, err := jwt.ParseJwtGoToken(token)
	if err == nil {
		userId = gconv.Int(tokenInfo.Id)
	}
	// 生成记录
	auditServer := new(services.CmsAuditServer)
	auditId, err := auditServer.MakeAudit(auditList, 1, userId, gconv.Int(param.SiteId))
	if err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "写入记录失败")
		return
	}
	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", auditId)
}

// 模块生成
func (*CmsMakeController) ModuleMake(c *gin.Context) {
	var param struct {
		Id          int   `form:"id" json:"id" binding:"required"`
		SiteId      int   `form:"site_id" json:"site_id" binding:"required"`
		VideoOffset int64 `form:"video_offset" json:"video_offset"`
		VideoLimit  int64 `form:"video_limit" json:"video_limit"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
		return
	}
	// 获取模块对应模板所有的页面id
	var templateList []map[string]interface{}
	global.DB.Model(&models.CmsTemplate{}).Select("id,template_name,site_id,module_ids").Where("status not in ? and site_id=?", []int{5, 9}, param.SiteId).Find(&templateList)
	var ids []int
	for _, v := range templateList {
		var moduleIds []string
		_ = json.Unmarshal(gconv.Bytes(v["module_ids"]), &moduleIds)
		if helper.InArray(strconv.Itoa(param.Id), moduleIds) {
			var pageList []map[string]interface{}
			global.DB.Model(&models.CmsPage{}).Where("site_id=? and template_id=? and status not in ?", param.SiteId, v["id"], []int{5, 9}).Find(&pageList)
			for _, v := range pageList {
				if !helper.InArray(gconv.Int(v["id"]), ids) {
					ids = append(ids, gconv.Int(v["id"]))
				}
			}
		}
	}
	if len(ids) == 0 {
		fmt.Println("ids为空")
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, "没有需要生成的页面", "生成文件失败")
		return
	}

	// 生成
	makeService := new(services.CmsMakeServer)
	// 生成文件
	auditList, _, err := makeService.Make(services.MakeTypePage, ids, gconv.Int(param.SiteId), false, services.VideoLimitSt{
		OffSet: param.VideoOffset,
		Limit:  param.VideoLimit,
	})
	if err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "生成文件失败")
		return
	}
	//user_id
	userId := 0
	token := c.Request.Header.Get("token")
	tokenInfo, err := jwt.ParseJwtGoToken(token)
	if err == nil {
		userId = gconv.Int(tokenInfo.Id)
	}
	// 生成记录
	auditServer := new(services.CmsAuditServer)
	auditId, err := auditServer.MakeAudit(auditList, 1, userId, gconv.Int(param.SiteId))
	if err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "写入记录失败")
		return
	}
	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", auditId)
}

// 视频tag 生成  // todu
func (*CmsMakeController) VideoTagMake(c *gin.Context) {
	var param struct {
		Id     int `form:"id" json:"id" binding:"required"`
		SiteId int `form:"site_id" json:"site_id" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
		return
	}
	ids := []int{param.Id}
	makeService := new(services.CmsMakeServer)
	// 生成文件
	auditList, _, err := makeService.Make(services.MakeTypeVideoTag, ids, gconv.Int(param.SiteId), false, services.VideoLimitSt{})
	if err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "生成文件失败")
		return
	}
	//user_id
	userId := 0
	token := c.Request.Header.Get("token")
	tokenInfo, err := jwt.ParseJwtGoToken(token)
	if err == nil {
		userId = gconv.Int(tokenInfo.Id)
	}
	// 生成记录
	auditServer := new(services.CmsAuditServer)
	auditId, err := auditServer.MakeAudit(auditList, 1, userId, gconv.Int(param.SiteId))
	if err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "写入记录失败")
		return
	}
	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", auditId)
}
