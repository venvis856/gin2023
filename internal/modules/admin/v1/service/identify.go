package service

import (
	"gin/api/admin/identify/v1"
	"gin/internal/library/handlePanic"
	"gin/internal/modules/admin/v1/models"
	"gorm.io/gorm"
)

type IdentifyInterface interface {
	Items(param v1.ItemReq) map[string]interface{}
	Create(param v1.CreateReq) (int64, error)
	Update(param v1.UpdateReq) (int64, error)
	Delete(param v1.DeleteReq) (int64, error)
	Info(param v1.InfoReq) (map[string]interface{}, error)
	InitIdentifyPermission(identifyType int8, identifyId int64, roleId int64, tx *gorm.DB) error
	GetNoPoliceIdentify(identifyId int64) []models.Identify
}

var identifyObj IdentifyInterface

func Identify() IdentifyInterface {
	if identifyObj == nil {
		handlePanic.Panic("identify service panic")
	}
	return identifyObj
}

func RegisterIdentify(i IdentifyInterface) {
	identifyObj = i
}
