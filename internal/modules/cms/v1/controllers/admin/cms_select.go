package admin

import (
	"encoding/json"
	"fmt"
	"gin/app/library/jwt"
	"gin/app/models"
	"gin/global"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gconv"
)

type CmsSelectController struct{}

func (*CmsSelectController) GetPermissionByUser(c *gin.Context) {
	//user_id
	userId := 0
	token := c.Request.Header.Get("token")
	tokenInfo, err := jwt.ParseJwtGoToken(token)
	if err == nil {
		userId = gconv.Int(tokenInfo.Id)
	}
	userInfo := make(map[string]interface{})
	global.DB.Model(&models.User{}).First(&userInfo, userId)
	// 获取站点信息
	var roleIds []int64
	err = json.Unmarshal(gconv.Bytes(userInfo["role_ids"]), &roleIds)
	if err != nil {
		fmt.Println("json数据转换失败：", err.Error())
		global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", false)
	}
	// 获取所属角色拥有的权限
	isGlobal := false
	if len(roleIds) != 0 {
		for _, roleId := range roleIds {
			// 如果是全局角色直接返回所有权限
			roleRow := map[string]interface{}{}
			global.DB.Model(&models.Role{}).Where("id=? and type =1", roleId).First(&roleRow)
			if len(roleRow) != 0 {
				isGlobal = true
				break
			}
		}
	}
	fmt.Println(isGlobal, "isGlobal")
	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", isGlobal)
}

// 获取站点列表  用户里面配置有
func (*CmsSelectController) GetSiteSelectList(c *gin.Context) {
	//user_id
	userId := 0
	token := c.Request.Header.Get("token")
	tokenInfo, err := jwt.ParseJwtGoToken(token)
	if err == nil {
		userId = gconv.Int(tokenInfo.Id)
	}
	// siteIds
	userInfo := make(map[string]interface{})
	global.DB.Model(&models.User{}).First(&userInfo, userId)
	var siteIds []int64
	_ = json.Unmarshal(gconv.Bytes(userInfo["site_ids"]), &siteIds)
	// 获取站点信息
	model := global.DB.Model(&models.CmsSite{})
	model.Where("status != ? and id in ?", 9, siteIds)
	var result []map[string]interface{}
	model.Find(&result)
	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", result)
}

func (*CmsSelectController) GetTemplateSelectList(c *gin.Context) {
	var param struct {
		SiteId int `form:"site_id" json:"site_id" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
		return
	}
	model := global.DB.Model(&models.CmsTemplate{})
	model.Where("status != ?", 9)
	model.Where("site_id = ?", param.SiteId)
	var result []map[string]interface{}
	model.Select("template_name,id").Find(&result)
	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", result)
}

func (*CmsSelectController) GetModuleSelectList(c *gin.Context) {
	var param struct {
		SiteId int `form:"site_id" json:"site_id" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
		return
	}
	model := global.DB.Model(&models.CmsModule{})
	model.Select("module_name,id")
	model.Where("status != ?", 9)
	model.Where("site_id = ?", param.SiteId)
	var result []map[string]interface{}
	model.Find(&result)
	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", result)
}

func (*CmsSelectController) GetClassifySelectList(c *gin.Context) {
	var param struct {
		SiteId int `form:"site_id" json:"site_id" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
		return
	}
	model := global.DB.Model(&models.CmsClassify{})
	model.Where("status != ?", 9)
	model.Where("site_id = ?", param.SiteId)
	var result []map[string]interface{}
	model.Find(&result)
	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", result)
}

func (*CmsSelectController) GetAuthorSelectList(c *gin.Context) {
	var param struct {
		SiteId int `form:"site_id" json:"site_id" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
		return
	}
	model := global.DB.Model(&models.CmsAuthor{})
	model.Where("status != ?", 9)
	model.Where("site_id = ?", param.SiteId)
	var result []map[string]interface{}
	model.Find(&result)
	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", result)
}

func (*CmsSelectController) GetProductSelectList(c *gin.Context) {
	var param struct {
		SiteId int `form:"site_id" json:"site_id" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
		return
	}
	model := global.DB.Model(&models.CmsProduct{})
	model.Where("status != ?", 9)
	model.Where("site_id = ?", param.SiteId)
	var result []map[string]interface{}
	model.Find(&result)
	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", result)
}

func (*CmsSelectController) GetTagSelectList(c *gin.Context) {
	var param struct {
		SiteId int `form:"site_id" json:"site_id" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
		return
	}
	model := global.DB.Model(&models.CmsTag{})
	model.Where("status != ?", 9)
	model.Where("site_id = ?", param.SiteId)
	var result []map[string]interface{}
	model.Find(&result)
	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", result)
}

func (*CmsSelectController) GetVideoTagSelectList(c *gin.Context) {
	var param struct {
		SiteId int `form:"site_id" json:"site_id" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
		return
	}
	model := global.DB.Model(&models.VideoTag{})
	model.Where("status != ?", 9)
	model.Where("site_id = ?", param.SiteId)
	var result []map[string]interface{}
	model.Find(&result)
	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", result)
}
