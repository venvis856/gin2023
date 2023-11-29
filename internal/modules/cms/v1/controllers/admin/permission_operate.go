package admin

import (
	"gin/app/library/jwt"
	"gin/app/library/permission"
	"gin/global"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gconv"
)

// 用户添加权限
func (*PermissionController) UserAddPermission(c *gin.Context) {
	var param struct {
		UserId         int64  `form:"user_id" json:"user_id" binding:"required"`
		PermissionCode string `form:"permission_code" json:"permission_code" binding:"required"`
		SiteId         int64  `form:"site_id" json:"site_id" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
		return
	}
	if err := permission.UserAddPermission(param.UserId, param.PermissionCode, param.SiteId); err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
		return
	}
	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", "添加成功")
}

// 用户删除权限
func (*PermissionController) UserDeletePermission(c *gin.Context) {
	var param struct {
		UserId         int64  `form:"user_id" json:"user_id" binding:"required"`
		PermissionCode string `form:"permission_code" json:"permission_code" binding:"required"`
		SiteId         int64  `form:"site_id" json:"site_id" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
		return
	}
	if err := permission.UserDeletePermission(param.UserId, param.PermissionCode, param.SiteId); err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
		return
	}
	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", "删除成功")
}

// 角色添加权限
func (*PermissionController) RoleAddPermission(c *gin.Context) {
	var param struct {
		RoleId         int64  `form:"role_id" json:"role_id" binding:"required"`
		PermissionCode []string `form:"permission_code" json:"permission_code" binding:"required"`
		SiteId         int64  `form:"site_id" json:"site_id" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "param err")
		return
	}
	if err := permission.RoleAddPermission(param.RoleId, param.PermissionCode, param.SiteId); err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
		return
	}
	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", "添加成功")
}

// 角色删除权限
//func (*PermissionController) RoleDeletePermission(c *gin.Context) {
//	var param struct {
//		RoleId         int64  `form:"role_id" binding:"required"`
//		PermissionCode string `form:"permission_code" binding:"required"`
//		SiteId         int64  `form:"site_id" binding:"required"`
//	}
//	if err := c.ShouldBind(&param); err != nil {
//		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
//		return
//	}
//	if err := permission.RoleDeletePermission(param.RoleId, param.PermissionCode, param.SiteId); err != nil {
//		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
//		return
//	}
//	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", "删除成功")
//}

// 获取用户的直接权限
func (*PermissionController) GetPermissionByUser(c *gin.Context) {
	var param struct {
		UserId int64 `form:"user_id" json:"user_id" binding:"required"`
		SiteId int64 `form:"site_id" json:"site_id" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
		return
	}
	result := permission.GetPermissionByUser(param.UserId, param.SiteId)
	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", result)
}

// 获取角色的所有权限
func (*PermissionController) GetAllPermissionByRole(c *gin.Context) {
	var param struct {
		RoleId int64 `form:"role_id" json:"role_id" binding:"required"`
		SiteId int64 `form:"site_id" json:"site_id" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
		return
	}
	result := permission.GetPermissionByRole(param.RoleId, param.SiteId)
	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", result)
}

// 获取用户的所有权限
func (*PermissionController) GetAllPermissionByUser(c *gin.Context) {
	var param struct {
		Token  string `form:"token" json:"token" binding:"required"`
		SiteId *int64 `form:"site_id" json:"site_id" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
		return
	}

	tokenInfo, err := jwt.ParseJwtGoToken(param.Token)
	if err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
		return
	}

	result := permission.GetAllPermissionByUser(gconv.Int64(tokenInfo.Id), gconv.Int64(param.SiteId))
	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", result)
}

// 获取所有权限列表
//func (*PermissionController) GetAllPermission(c *gin.Context) {
//	result := permission.GetAllPermission()
//	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", result)
//}
