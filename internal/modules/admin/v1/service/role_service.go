package service

import (
	"gin/api/admin/role/v1"
	"gin/internal/library/handlePanic"
)

type RoleInterface interface {
	Items(param v1.ItemReq) (map[string]interface{}, error)
	Info(param v1.InfoReq) map[string]interface{}
	Create(param v1.CreateReq) (int64, error)
	Update(param v1.UpdateReq) (int64, error)
	Delete(param v1.DeleteReq) (int64, error)
}

var roleObj RoleInterface

func Role() RoleInterface {
	if roleObj == nil {
		handlePanic.Panic("role service panic")
	}
	return roleObj
}

func RegisterRole(i RoleInterface) {
	roleObj = i
}
