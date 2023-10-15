package router

import (
	adminV1 "gin/internal/modules/double_admin/v1"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	adminV1.InitAdminRoutes(router)
}
