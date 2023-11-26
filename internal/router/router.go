package router

import (
	"gin/internal/middleware"
	adminV1 "gin/internal/modules/admin/v1"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	router.Use(middleware.Cors())    // 跨域
	router.Use(middleware.TraceId()) // 跟踪信息

	adminV1.InitAdminRoutes(router)
}
