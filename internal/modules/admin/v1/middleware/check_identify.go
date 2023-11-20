
package middleware

import (
"gin/internal/global"
"gin/internal/global/errcode"
"github.com/gin-gonic/gin"
)

func CheckIdentify() gin.HandlerFunc {
	return func(c *gin.Context) {
		identifyId := c.Request.Header.Get("X-Identify-Id")
		if identifyId == "" {
			//c.String(401, "无效的请求")
			//global.Response.Json(c, global.HTTP_SUCCESS, global.TOKEN_FAIL, "请求无效", "")
			global.Response.Error(c, errcode.ERROR_SERVER, "Identify err")
			c.Abort()
			return
		}
		// 赋值
		c.Set("identify_id", identifyId)
		// 继续往下处理
		c.Next()
	}
}
