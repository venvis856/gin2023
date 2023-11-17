package router

import (
	"gin/internal/common_middleware"
	adminV1 "gin/internal/modules/admin/v1"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	router.Use(common_middleware.Cors())    // 跨域
	router.Use(common_middleware.TraceId()) // 跟踪信息

	adminV1.InitAdminRoutes(router)
}
