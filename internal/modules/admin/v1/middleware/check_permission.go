package middleware

import (
	"errors"
	"gin/internal/global/errcode"
	"gin/internal/global/response"
	"gin/internal/modules/admin/v1/service"
	"github.com/gin-gonic/gin"
)

func CheckPermission(permissionCodes ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userInfo := service.User().GetUserInfo(c)
		if userInfo == nil {
			//global.Response.Json(c, global.HTTP_SUCCESS, global.TOKEN_FAIL, "获取用户信息失败", "")
			response.Error(c, errcode.TOKEN_FAIL.Wrap(errors.New("获取用户信息失败")))
			c.Abort()
			return
		}

		if userInfo.UserId == 0 {
			//c.String(401, "无效的请求")
			//global.Response.Json(c, global.HTTP_SUCCESS, global.TOKEN_FAIL, "用户无效", "")
			response.Error(c, errcode.TOKEN_FAIL.Wrap(errors.New("用户无效")))
			c.Abort()
			return
		}

		//method := c.Request.Method
		//fmt.Println(method, "====method")
		if len(permissionCodes) < 0 {
			response.Error(c, errcode.TOKEN_FAIL.Wrap(errors.New("用户无效")))
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
			response.Error(c, errcode.FORBIDDEN_METHOD.Wrap(errors.New("权限不足")))
			c.Abort()
			return
		}

		// 继续往下处理
		c.Next()
	}
}
