package api

import (
	"gin/app/models"
	"gin/global"
	"github.com/gin-gonic/gin"
)

type SkySearchController struct{}

func (*SkySearchController) Search(c *gin.Context) {
	var param struct {
		Name string `form:"name" json:"name"`
		Limit       int    `form:"limit" json:"limit"`
		PageIndex   int    `form:"pageIndex" json:"pageIndex"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
		return
	}
	model:=global.DB.Model(&models.CmsPage{})
	// subject，时间，赞数,tag
	model.Select("subject,first_make_time,star_number,tag_ids")
	// where  is_publish site_id
	model.Where("is_publish=1 and status!=9 and status!=5 and site_id=2")
	model.Where("subject like '%?%'",param.Name)

	var count int64
	model.Count(&count)

	if param.Limit != 0 {
		if param.PageIndex == 0 {
			param.PageIndex = 1
		}
		model.Offset((param.PageIndex - 1) * param.Limit).Limit(param.Limit)
	}

	var result []map[string]interface{}
	model.Scan(&result)
	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", map[string]interface{}{"items": result, "total": count})
}
