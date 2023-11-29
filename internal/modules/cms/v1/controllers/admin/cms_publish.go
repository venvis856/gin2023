package admin

import (
	"gin/app/services"
	"gin/global"
	"github.com/gin-gonic/gin"
)

type CmsPublishController struct {
}

func (*CmsPublishController) Publish(c *gin.Context) {
	var param struct {
		Id     int `form:"id" json:"id" binding:"required"`
		SiteId int `form:"site_id" json:"site_id" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
		return
	}
	publishService := new(services.CmsPublishService)
	err := publishService.Publish(c, param)
	if err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
		return
	}
	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", "发布成功")
}
