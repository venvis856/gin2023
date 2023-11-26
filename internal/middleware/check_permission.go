package middleware

import (
	"gin/internal/global"
	"gin/internal/global/errcode"
	"gin/internal/modules/admin/v1/service"
	"github.com/gin-gonic/gin"
)

func CheckPermission(permissionCodes ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userInfo := service.User().GetUserInfo(c)
		if userInfo == nil {
			global.Response.Error(c,errcode.ERROR_PARAMS,"获取用户信息失败")
			c.Abort()
			return
		}

		if userInfo.UserId == 0 {
			global.Response.Error(c,errcode.ERROR_PARAMS,"用户无效")
			c.Abort()
			return
		}

		//method := c.Request.Method
		//fmt.Println(method, "====method")
		if len(permissionCodes) < 0 {
			global.Response.Error(c,errcode.ERROR_PARAMS,"无此权限")
			c.Abort()
			return
		}

		bol := false
		for _, v := range permissionCodes {
			if ok := service.Permission().CheckAuth(c, v, userInfo.IdentifyId); ok {
				bol = true
			}
		}
		if !bol {
			global.Response.Error(c,errcode.ERROR_PARAMS,"权限不足")
			c.Abort()
			return
		}

		// 继续往下处理
		c.Next()
	}
}
