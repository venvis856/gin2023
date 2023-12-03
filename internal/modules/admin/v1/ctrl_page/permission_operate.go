package ctrl_page

import (
	"gin/internal/global"
	"gin/internal/library/jwt"
	"gin/internal/modules/admin/v1/service"

	"gin/internal/global/errcode"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gconv"
)

// 用户添加权限
func (*PermissionController) UserAddPermission(c *gin.Context) {
	var param struct {
		UserId         int64  `form:"user_id" json:"user_id" binding:"required"`
		PermissionCode []string `form:"permission_code" json:"permission_code" binding:"required"`
		SiteId         int64  `form:"site_id" json:"site_id" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS,err.Error())
		return
	}
	if err := service.PermissionOperate().UserAddPermission(param.UserId, param.PermissionCode, param.SiteId); err != nil {
		global.Response.Error(c, errcode.ERROR_SERVER, err.Error())
		return
	}
	global.Response.Success(c, "添加成功")
}

// 用户删除权限
//func (*PermissionController) UserDeletePermission(c *gin.Context) {
//	var param struct {
//		UserId         int64  `form:"user_id" json:"user_id" binding:"required"`
//		PermissionCode string `form:"permission_code" json:"permission_code" binding:"required"`
//		SiteId         int64  `form:"site_id" json:"site_id" binding:"required"`
//	}
//	if err := c.ShouldBind(&param); err != nil {
//		global.Response.Error(c, errcode.ERROR_PARAMS,err.Error())
//		return
//	}
//	if err := service.Permission().UserDeletePermission(param.UserId, param.PermissionCode, param.SiteId); err != nil {
//		global.Response.Error(c, errcode.ERROR_SERVER, err.Error())
//		return
//	}
//	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", "删除成功")
//}

// 角色添加权限
func (*PermissionController) RoleAddPermission(c *gin.Context) {
	var param struct {
		RoleId         int64  `form:"role_id" json:"role_id" binding:"required"`
		PermissionCode []string `form:"permission_code" json:"permission_code" binding:"required"`
		SiteId         int64  `form:"site_id" json:"site_id" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS,err.Error())
		return
	}
	if err := service.PermissionOperate().RoleAddPermission(param.RoleId, param.PermissionCode, param.SiteId); err != nil {
		global.Response.Error(c, errcode.ERROR_SERVER, err.Error())
		return
	}
	global.Response.Success(c, "添加成功")
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
//	if err := permission_service.RoleDeletePermission(param.RoleId, param.PermissionCode, param.SiteId); err != nil {
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
		global.Response.Error(c, errcode.ERROR_PARAMS,err.Error())
		return
	}
	result := service.PermissionOperate().GetPermissionByUser(param.UserId, param.SiteId)
	global.Response.Success(c,  result)
}

// 获取角色的所有权限
func (*PermissionController) GetAllPermissionByRole(c *gin.Context) {
	var param struct {
		RoleId int64 `form:"role_id" json:"role_id" binding:"required"`
		SiteId int64 `form:"site_id" json:"site_id" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS,err.Error())
		return
	}
	result := service.PermissionOperate().GetPermissionByRole(param.RoleId, param.SiteId)
	global.Response.Success(c,  result)
}

// 获取用户的所有权限
func (*PermissionController) GetAllPermissionByUser(c *gin.Context) {
	var param struct {
		Token  string `form:"token" json:"token" binding:"required"`
		SiteId *int64 `form:"site_id" json:"site_id" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS,err.Error())
		return
	}

	tokenInfo, err := jwt.ParseJwtGoToken(param.Token)
	if err != nil {
		global.Response.Error(c,errcode.ERROR_SERVER, err.Error())
		return
	}

	result := service.PermissionOperate().GetAllPermissionByUser(gconv.Int64(tokenInfo.Id), gconv.Int64(param.SiteId))
	global.Response.Success(c, result)
}

// 获取所有权限列表
//func (*PermissionController) GetAllPermission(c *gin.Context) {
//	result := permission_service.GetAllPermission()
//	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", result)
//}
