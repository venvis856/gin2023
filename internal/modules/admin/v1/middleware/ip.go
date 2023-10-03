package middleware

import (
	"github.com/gin-gonic/gin"
)

func IPAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ipList := []string{
			"127.0.0.1",
		}
		flag := false
		clientIP := c.ClientIP() //获取当前ip
		for _, host := range ipList {
			if clientIP == host {
				flag = true
				break
			}
		}
		if !flag {
			c.String(401, "%s ip不允许访问", clientIP)
			c.Abort()
			return
		}
		c.Next()
	}
}
