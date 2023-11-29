package ctrl_admin

import (
	"fmt"
	"gin/api/admin/permissionOperate/v1"
	"gin/internal/global"
	"gin/internal/global/errcode"
	"gin/internal/modules/admin/v1/service"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gconv"
)

// 用户添加权限
func (*PermissionCtrl) UserAddPermission(c *gin.Context) {
	var param v1.UserAddPermissionReq
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS, fmt.Sprintf("param err: %v", err))
		return
	}
	if err := service.PermissionOperate().UserAddPermission(param.UserId, param.PermissionCode, param.IdentifyId); err != nil {
		global.Response.Error(c, errcode.ERROR_SERVER, err.Error())
		return
	}
	global.Response.Success(c, "添加成功")
}

// 角色添加权限
func (*PermissionCtrl) RoleAddPermission(c *gin.Context) {
	var param v1.RoleAddPermissionReq
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS, fmt.Sprintf("param err: %v", err))
		return
	}
	if err := service.PermissionOperate().RoleAddPermission(param.RoleId, param.PermissionCode, param.IdentifyId); err != nil {
		global.Response.Error(c, errcode.ERROR_SERVER, err.Error())
		return
	}
	//global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", "添加成功")
	global.Response.Success(c, "添加成功")
}

// 获取用户的直接权限
func (*PermissionCtrl) GetPermissionByUser(c *gin.Context) {
	var param v1.GetPermissionByUserReq
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS, fmt.Sprintf("param err: %v", err))
		return
	}
	result := service.PermissionOperate().GetPermissionByUser(param.UserId, param.IdentifyId)
	global.Response.Success(c, result)
}

// 获取角色的所有权限
func (*PermissionCtrl) GetAllPermissionByRole(c *gin.Context) {
	var param v1.GetAllPermissionByRoleReq
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS, fmt.Sprintf("param err: %v", err))
		return
	}
	result := service.PermissionOperate().GetPermissionByRole(param.RoleId, param.IdentifyId)
	global.Response.Success(c, result)
}

// 获取用户的所有权限
func (*PermissionCtrl) GetAllPermissionByUser(c *gin.Context) {
	var param v1.GetAllPermissionByUserReq
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS, fmt.Sprintf("param err: %v", err))
		return
	}

	userInfo := service.User().GetUserInfo(c)

	result := service.PermissionOperate().GetAllPermissionByUser(gconv.Int64(userInfo.UserId), gconv.Int64(param.IdentifyId))
	global.Response.Success(c, result)
}

// 获取所有权限列表
//func (*PermissionHandler) GetAllPermission(c *gin.Context) {
//	result := permission.GetAllPermission()
//	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", result)
//}

// 获取用户权限 只限菜单
func (*PermissionCtrl) GetMenuByUser(c *gin.Context) {
	var param v1.GetMenuByUserReq
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS, fmt.Sprintf("param err: %v", err))
		return
	}

	userInfo := service.User().GetUserInfo(c)
	permissionList := service.PermissionOperate().GetAllPermissionByUser(int64(userInfo.UserId), param.IdentifyId)
	// 限制菜单
	var result []map[string]interface{}
	for _, v := range permissionList {
		if gconv.Int64(v["type"]) == 1 {
			result = append(result, v)
		}
	}
	global.Response.Success(c, result)
}
