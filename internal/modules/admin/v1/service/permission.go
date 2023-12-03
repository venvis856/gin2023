package service

import (
	"gin/api/admin/permission/v1"
	"gin/internal/library/handlePanic"
	"github.com/gin-gonic/gin"
)

type PermissionInterface interface {
	Items(param v1.ItemReq) (map[string]interface{}, error)
	Create(param v1.CreateReq) (int64, error)
	Update(param v1.UpdateReq) (int64, error)
	Delete(c *gin.Context, param v1.DeleteReq) (int64, error)
}

var permissionObj PermissionInterface

func Permission() PermissionInterface {
	if permissionObj == nil {
		handlePanic.Panic("permission_service service panic")
	}
	return permissionObj
}

func RegisterPermission(i PermissionInterface) {
	permissionObj = i
}
