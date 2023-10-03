package middleware

import (
	"gin/internal/global"
	"github.com/gin-gonic/gin"
	"strconv"
)

func RequestLog(logPath string) gin.HandlerFunc {
	return func(c *gin.Context) {
		msg := ""
		clientIP := c.ClientIP()
		msg += "ip: (" + clientIP + ") | " + "route: " + c.Request.RequestURI + " | " + c.Request.Method + " | " + strconv.Itoa(c.Writer.Status())
		global.Logger.Daily(logPath, "info", msg)
		c.Next()
	}
}
