package ctrl_page

import (
	"fmt"
	"gin/internal/global"
	"gin/internal/global/errcode"
	services "gin/internal/modules/admin/v1/logic/cms_make"
	"gin/internal/modules/admin/v1/service"
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
		global.Response.Error(c, errcode.ERROR_PARAMS,err.Error())
		return
	}
	ids := []int{param.Id}
	_, content, err := service.CmsMake().Make(services.MakeTypePage, ids, gconv.Int(param.SiteId), true)
	if err != nil {
		global.Response.Error(c, errcode.ERROR_SERVER, fmt.Sprintf("预览失败: %v",err.Error()))
		return
	}
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(200, content)
}
