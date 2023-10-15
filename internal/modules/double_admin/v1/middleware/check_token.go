package middleware

import (
	"gin/internal/global/errcode"
	"gin/internal/global/response"
	"gin/internal/library/jwt"
	"github.com/gin-gonic/gin"
)

func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			//c.String(401, "无效的请求")
			//global.Response.Json(c, global.HTTP_SUCCESS, global.TOKEN_FAIL, "请求无效", "")
			response.Error(c, errcode.TOKEN_FAIL)
			c.Abort()
			return
		}
		_, err := jwt.ParseJwtGoToken(token)
		if err != nil {
			//global.Response.Json(c, global.HTTP_SUCCESS, global.TOKEN_FAIL, "请求无效："+err.Error(), "")
			response.Error(c, errcode.TOKEN_FAIL.Wrap(err))
			c.Abort()
			return
		}
		// 继续往下处理
		c.Next()
	}
}
