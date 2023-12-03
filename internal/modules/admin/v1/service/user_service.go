package service

import (
	"gin/api/admin/user/v1"
	"gin/internal/library/handlePanic"
	"gin/internal/modules/admin/v1/models"
	"github.com/gin-gonic/gin"
)

type UserInterface interface {
	Items(param v1.ItemReq) (map[string]interface{}, error)
	GetUserInfo(c *gin.Context) *v1.UserInfo
	Info(param v1.InfoReq) (map[string]interface{}, error)
	Create(param v1.CreateReq) (int64, error)
	Update(param v1.UpdateReq) (int64, error)
	Delete(param v1.DeleteReq) (int64, error)
	GetSecret(pwd string) string
	GetUserIdentify(c *gin.Context, userId int64) []models.Identify
}

var userObj UserInterface

func User() UserInterface {
	if userObj == nil {
		handlePanic.Panic("user_service service panic")
	}
	return userObj
}

func RegisterUser(i UserInterface) {
	userObj = i
}
