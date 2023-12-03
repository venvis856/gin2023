package ctrl_page

import (
	"gin/internal/global"
	"gin/internal/modules/admin/v1/models"
	"gin/internal/modules/admin/v1/service"
	"github.com/gin-gonic/gin"
)

type SiteCtrl struct{}

func (a *SiteCtrl) SelectList(c *gin.Context) {
	identifyList := service.User().GetUserIdentify(c, 0)
	IdentifyIds := make([]int64, 0)
	for _, v := range identifyList {
		IdentifyIds = append(IdentifyIds, v.ID)
	}
	var siteList []models.CmsSite
	global.DB.Model(&models.CmsSite{}).Where("identify_id in (?)", IdentifyIds).Find(&siteList)
	global.Response.Success(c, siteList)
}