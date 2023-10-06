package handler

import (
	"fmt"
	"gin/api/admin/permissionOperate/v1"
	"gin/internal/global/errcode"
	"gin/internal/global/response"
	"gin/internal/modules/admin/v1/logic/permission_operate"
	"gin/internal/modules/admin/v1/service"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gconv"
)

// 用户添加权限
func (*PermissionHandler) UserAddPermission(c *gin.Context) {
	var param v1.UserAddPermissionReq
	if err := c.ShouldBind(&param); err != nil {
		//global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "param err")
		response.ErrorWithData(c, errcode.ERROR_PARAMS.Wrap(err), "param err")
		return
	}
	if err := permission_operate.UserAddPermission(param.UserId, param.PermissionCode, param.IdentifyId); err != nil {
		//global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
		response.Error(c, err)
		return
	}
	//global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", "添加成功")
	response.Success(c, "添加成功")
}

// 用户删除权限
//func (*PermissionHandler) UserDeletePermission(c *gin.Context) {
//	var param struct {
//		UserId         int64  `form:"user_id" json:"user_id" binding:"required"`
//		PermissionCode string `form:"permission_code" json:"permission_code" binding:"required"`
//		IdentifyId     int64  `form:"identify_id" json:"identify_id" binding:"required"`
//	}
//	if err := c.ShouldBind(&param); err != nil {
//		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
//		return
//	}
//	if err := permission.UserDeletePermission(param.UserId, param.PermissionCode, param.IdentifyId); err != nil {
//		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
//		return
//	}
//	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", "删除成功")
//}

// 角色添加权限
func (*PermissionHandler) RoleAddPermission(c *gin.Context) {
	var param v1.RoleAddPermissionReq
	if err := c.ShouldBind(&param); err != nil {
		//global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "param err")
		response.ErrorWithData(c, errcode.ERROR_PARAMS.Wrap(err), "param err")
		return
	}
	if err := permission_operate.RoleAddPermission(param.RoleId, param.PermissionCode, param.IdentifyId); err != nil {
		//global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
		response.Error(c, err)
		return
	}
	//global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", "添加成功")
	response.Success(c, "添加成功")
}

// 角色删除权限
//func (*PermissionHandler) RoleDeletePermission(c *gin.Context) {
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
func (*PermissionHandler) GetPermissionByUser(c *gin.Context) {
	var param v1.GetPermissionByUserReq
	if err := c.ShouldBind(&param); err != nil {
		//global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "param err")
		response.ErrorWithData(c, errcode.ERROR_PARAMS.Wrap(err), "param err")
		return
	}
	result := permission_operate.GetPermissionByUser(param.UserId, param.IdentifyId)
	//global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", result)
	response.Success(c, result)
}

// 获取角色的所有权限
func (*PermissionHandler) GetAllPermissionByRole(c *gin.Context) {
	var param v1.GetAllPermissionByRoleReq
	if err := c.ShouldBind(&param); err != nil {
		//global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "param err")
		response.ErrorWithData(c, errcode.ERROR_PARAMS.Wrap(err), "param err")
		return
	}
	result := permission_operate.GetPermissionByRole(param.RoleId, param.IdentifyId)
	fmt.Println(result, "=====auto")
	//global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", result)
	response.Success(c, result)
}

// 获取用户的所有权限
func (*PermissionHandler) GetAllPermissionByUser(c *gin.Context) {
	var param v1.GetAllPermissionByUserReq
	if err := c.ShouldBind(&param); err != nil {
		//global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "param err")
		response.ErrorWithData(c, errcode.ERROR_PARAMS.Wrap(err), "param err")
		return
	}

	userInfo := service.User().GetUserInfo(c)

	result := permission_operate.GetAllPermissionByUser(gconv.Int64(userInfo.UserId), gconv.Int64(param.IdentifyId))
	//global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", result)
	response.Success(c, result)
}

// 获取所有权限列表
//func (*PermissionHandler) GetAllPermission(c *gin.Context) {
//	result := permission.GetAllPermission()
//	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", result)
//}

// 获取用户权限 只限菜单
func (*PermissionHandler) GetMenuByUser(c *gin.Context) {
	var param v1.GetMenuByUserReq
	if err := c.ShouldBind(&param); err != nil {
		//global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "param err")
		response.ErrorWithData(c, errcode.ERROR_PARAMS.Wrap(err), "param err")
		return
	}

	userInfo := service.User().GetUserInfo(c)
	permissionList := permission_operate.GetAllPermissionByUser(int64(userInfo.UserId), param.IdentifyId)
	// 限制菜单
	var result []map[string]interface{}
	for _, v := range permissionList {
		if gconv.Int64(v["type"]) == 1 {
			result = append(result, v)
		}
	}
	//global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", result)
	response.Success(c, result)
}
