package service

import "gin/internal/library/handlePanic"

type PermissionOperateInterface interface {
	GetAllPermission() []map[string]interface{}
	GetAllPermissionCodes() []string
	GetAllPermissionByIdentify(identifyId int64) []map[string]interface{}
	IdentifyAddPermission(identifyId int64, permissionCodes []string) error
	UserAddPermission(userId int64, permissionCodes []string, identifyId int64) error
	GetPermissionByUser(userId int64, identifyId int64) []map[string]interface{}
	GetAllPermissionByUser(userId int64, identifyId int64) []map[string]interface{}
	CheckUserHasPermission(userId int64, permissionCode string, identifyId int64) bool
	CheckRoleHasPermission(roleId int64, permissionId int64, identifyId int64) bool
	RoleAddPermission(roleId int64, permissionCodes []string, identifyId int64) error
	GetPermissionByRole(roleId int64, identifyId int64) []map[string]interface{}
	GetPermissionIdsByPermissionCode(permissionCodes []string) []int64
}

var permissionOperateObj PermissionOperateInterface

func PermissionOperate() PermissionOperateInterface {
	if permissionOperateObj == nil {
		handlePanic.Panic("permissionOperate service panic")
	}
	return permissionOperateObj
}

func RegisterPermissionOperate(i PermissionOperateInterface) {
	permissionOperateObj = i
}