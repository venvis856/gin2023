package ctrl_page

import (
	"encoding/json"
	"fmt"
	"gin/internal/global"
	"gin/internal/library/helper"
	"gin/internal/library/jwt"
	"gin/internal/global/errcode"
	services "gin/internal/modules/admin/v1/logic/cms_make"
	"gin/internal/modules/admin/v1/models"
	"gin/internal/modules/admin/v1/service"
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
		global.Response.Error(c, errcode.ERROR_PARAMS,err.Error())
		return
	}
	ids := []int{param.Id}
	// 生成文件
	auditList, _, err := service.CmsMake().Make(services.MakeTypePage, ids, gconv.Int(param.SiteId), false)
	if err != nil {
		global.Response.Error(c, errcode.ERROR_SERVER,  fmt.Sprintf("生成文件失败:%v",err.Error()))
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
	auditId, err := service.Audit().MakeAudit(auditList, 1, userId, gconv.Int(param.SiteId))
	if err != nil {
		global.Response.Error(c, errcode.ERROR_SERVER, fmt.Sprintf("写入记录失败:%v",err.Error()))
		return
	}
	global.Response.Success(c, auditId)
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
		global.Response.Error(c, errcode.ERROR_PARAMS,err.Error())
		return
	}
	var pageList []map[string]interface{}
	global.DB.Model(&models.CmsPage{}).Where("site_id=? and template_id=? and status not in ?", param.SiteId, param.Id, []int{5, 9}).Find(&pageList)
	var ids []int
	for _, v := range pageList {
		ids = append(ids, gconv.Int(v["id"]))
	}

	// 生成文件
	auditList, _, err := service.CmsMake().Make(services.MakeTypePage, ids, gconv.Int(param.SiteId), false)
	if err != nil {
		global.Response.Error(c, errcode.ERROR_SERVER, fmt.Sprintf("生成文件失败:%v",err.Error()))
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
	auditId, err := service.Audit().MakeAudit(auditList, 1, userId, gconv.Int(param.SiteId))
	if err != nil {
		global.Response.Error(c, errcode.ERROR_SERVER, fmt.Sprintf("写入记录失败:%v",err.Error()))
		return
	}
	global.Response.Success(c, auditId)
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
		global.Response.Error(c, errcode.ERROR_PARAMS,err.Error())
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
		global.Response.Error(c, errcode.ERROR_SERVER, "没有需要生成的页面")
		return
	}

	// 生成文件
	auditList, _, err := service.CmsMake().Make(services.MakeTypePage, ids, gconv.Int(param.SiteId), false)
	if err != nil {
		global.Response.Error(c, errcode.ERROR_SERVER, fmt.Sprintf("生成文件失败:%v",err.Error()))
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
	auditId, err := service.Audit().MakeAudit(auditList, 1, userId, gconv.Int(param.SiteId))
	if err != nil {
		global.Response.Error(c, errcode.ERROR_SERVER, fmt.Sprintf("写入记录失败:%v",err.Error()))
		return
	}
	global.Response.Success(c, auditId)
}
