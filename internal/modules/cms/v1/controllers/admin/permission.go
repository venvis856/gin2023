package admin

import (
	"fmt"
	"gin/app/library/jwt"
	"gin/app/library/permission"
	"gin/app/models"
	"gin/global"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gconv"
	"github.com/golang-module/carbon"
)

type PermissionController struct{}

// 获取所有权限列表
func (*PermissionController) Items(c *gin.Context) {
	var param struct {
		Limit       int    `form:"limit" json:"limit"`
		PageIndex   int    `form:"pageIndex" json:"pageIndex"`
		OrderBy     string `form:"orderBy" json:"orderBy"`
		OrderByType string `form:"orderByType" json:"orderByType"`
		Search      string `form:"search" json:"search"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
		return
	}
	fmt.Println(param.Search, "sera")
	model := global.DB.Model(&models.Permission{})
	model = WhereBySearch(model, param.Search)
	model.Where("status != ?", 9)

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
	}
	var result []map[string]interface{}
	model.Find(&result)
	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", map[string]interface{}{"items": result, "total": count})
}

// 新增权限
func (*PermissionController) Create(c *gin.Context) {
	var param struct {
		PermissionName     string `form:"permission_name" json:"permission_name" binding:"required"`
		PermissionCode     string `form:"permission_code" json:"permission_code" binding:"required"`
		SiteId     int64 `form:"site_id" json:"site_id" binding:"required"`
		Type               int64  `form:"type" json:"type" binding:"required"`
		FatherPermissionCode string `form:"father_permission_code" json:"father_permission_code" binding:"required"`
		Status int64 `form:"status" json:"status" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
		return
	}
	rs := global.DB.Model(&models.Permission{}).Create(map[string]interface{}{
		"permission_name":      param.PermissionName,
		"permission_code":      param.PermissionCode,
		"type":                 param.Type,
		"father_permission_code": param.FatherPermissionCode,
		"site_id" :param.SiteId,
		"status":               param.Status,
	})
	if rs.Error != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, rs.Error.Error(), "")
		return
	}
	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", "添加成功")
}

func (*PermissionController) Update(c *gin.Context) {
	var param struct {
		Id int64 `form:"id" json:"id" binding:"required"`
		PermissionName     string `form:"permission_name" json:"permission_name" binding:"required"`
		PermissionCode     string `form:"permission_code" json:"permission_code" binding:"required"`
		SiteId     int64 `form:"site_id" json:"site_id" binding:"required"`
		Type               int64  `form:"type" json:"type" binding:"required"`
		FatherPermissionCode string `form:"father_permission_code" json:"father_permission_code" binding:"required"`
		Status int64 `form:"status" json:"status" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
		return
	}
	data := map[string]interface{}{
		"permission_name":      param.PermissionName,
		"permission_code":      param.PermissionCode,
		"type":                 param.Type,
		"father_permission_code": param.FatherPermissionCode,
		"status":               param.Status,
		"site_id" :param.SiteId,
		"update_time":   carbon.Now().Timestamp(),
	}
	result := global.DB.Model(&models.Permission{}).Where("id = ? and status !=?", param.Id, 9).Updates(data)
	if result.Error != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, result.Error.Error(), "")
		return
	}
	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", result.RowsAffected)
}

// 删除权限
func (*PermissionController) Delete(c *gin.Context) {
	var param struct {
		Id string `form:"id" json:"id" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
		return
	}
	permissionInfo := map[string]interface{}{}
	global.DB.Model(&models.Permission{}).Where("status != ?", 9).First(&permissionInfo, param.Id)
	if permissionInfo == nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, "无该权限", "")
		return
	}
	// 删除用户权限
	rs := global.DB.Unscoped().Where("permission_code = ?", permissionInfo["permission_code"]).Delete(&models.UserPermission{})
	if rs.Error != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, rs.Error.Error(), "")
		return
	}
	// 删除角色权限
	rs = global.DB.Unscoped().Where("permission_code = ?", permissionInfo["permission_code"]).Delete(&models.RolePermission{})
	if rs.Error != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, rs.Error.Error(), "")
		return
	}
	// 删除权限
	rs = global.DB.Unscoped().Where("id = ?", param.Id).Delete(&models.Permission{})
	if rs.Error != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, rs.Error.Error(), "")
		return
	}
	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", "删除成功")
}

func (*PermissionController) GetMenuByUser(c *gin.Context) {
	var param struct {
		SiteId int64 `form:"site_id" json:"site_id" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
		return
	}
	userId := 0
	token := c.Request.Header.Get("token")
	tokenInfo, err := jwt.ParseJwtGoToken(token)
	if err == nil {
		userId = gconv.Int(tokenInfo.Id)
	}
	permissionList:=permission.GetAllPermissionByUser(int64(userId),param.SiteId)
	result:=make([]map[string]interface{},0)
	for _,v:=range permissionList{
		if gconv.Int64(v["type"])==1{
			result=append(result,v)
		}
	}
	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", result)
}