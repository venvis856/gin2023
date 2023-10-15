package common_middleware

import (
	"gin/internal/global/response"
	"github.com/gin-gonic/gin"
)

func Tracker() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_ = response.SetTrackId(ctx)
		// 处理请求
		ctx.Next()
	}
}
