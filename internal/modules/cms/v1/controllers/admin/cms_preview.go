package admin

import (
	"gin/app/services"
	"gin/global"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gconv"
)

type CmsPreviewController struct{}

func (*CmsPreviewController) Preview(c *gin.Context) {
	var param struct {
		Id     int `form:"id" json:"id" binding:"required"`
		SiteId int `form:"site_id" json:"site_id" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
		return
	}
	ids := []int{param.Id}
	makeService := new(services.CmsMakeServer)
	_, content, err := makeService.Make(services.MakeTypePage, ids, gconv.Int(param.SiteId), true, services.VideoLimitSt{
		OffSet: 0,
		Limit:  0,
	})
	if err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "预览失败")
		return
	}
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(200, content)

}
