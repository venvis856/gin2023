package middleware

import (
	"fmt"
	"gin/internal/global"
	"gin/internal/global/errcode"
	"gin/internal/modules/admin/v1/service"
	"github.com/gin-gonic/gin"
)

func CheckPermission(permissionCodes ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userInfo := service.User().GetUserInfo(c)
		if userInfo == nil {
			global.Response.Error(c, errcode.ERROR_PARAMS, "获取用户信息失败")
			c.Abort()
			return
		}

		if userInfo.ID == 0 {
			global.Response.Error(c, errcode.ERROR_PARAMS, "用户无效")
			c.Abort()
			return
		}

		//method := c.Request.Method
		//fmt.Println(method, "====method")
		if len(permissionCodes) < 0 {
			global.Response.Error(c, errcode.ERROR_PARAMS, "无此权限")
			c.Abort()
			return
		}

		// 获取 identify_service list
		identifyList := service.User().GetUserIdentify(c, userInfo.ID)
		if len(identifyList) == 0 {
			global.Response.Error(c, errcode.ERROR_PARAMS, "无此系统权限")
			c.Abort()
			return
		}
		bol := false
		fmt.Println(userInfo.ID, permissionCodes, identifyList,"=====22222222222222222222222")
		for _, v := range permissionCodes {
			for _, identify := range identifyList {
				fmt.Println(userInfo.ID, v, identify.ID, "=====code")
				if ok := service.PermissionOperate().CheckUserHasPermission(userInfo.ID, v, identify.ID); ok {
					bol = true
				}
			}
		}
		if !bol {
			global.Response.Error(c, errcode.ERROR_PARAMS, "权限不足")
			c.Abort()
			return
		}

		// 继续往下处理
		c.Next()
	}
}