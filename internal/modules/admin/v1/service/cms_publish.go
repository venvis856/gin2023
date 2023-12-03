package service

import (
	"gin/internal/library/handlePanic"
	"github.com/gin-gonic/gin"
)

type PublishParam struct {
	Id     int `form:"id" json:"id" binding:"required"`
	SiteId int `form:"site_id" json:"site_id" binding:"required"`
}

type PublishInterface interface {
	Publish(c *gin.Context, param PublishParam) error
}

var publishObj PublishInterface

func CmsPublish() PublishInterface {
	if publishObj == nil {
		handlePanic.Panic("publish service panic")
	}
	return publishObj
}

func RegisterCmsPublish(i PublishInterface) {
	publishObj = i
}
