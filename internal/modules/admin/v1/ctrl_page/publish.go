package ctrl_page

import (
	"gin/internal/global"
	"gin/internal/global/errcode"
	"gin/internal/modules/admin/v1/service"
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
		global.Response.Error(c, errcode.ERROR_PARAMS,err.Error())
		return
	}
	err := service.CmsPublish().Publish(c, param)
	if err != nil {
		global.Response.Error(c, errcode.ERROR_SERVER, err.Error())
		return
	}
	global.Response.Success(c, "发布成功")
}
