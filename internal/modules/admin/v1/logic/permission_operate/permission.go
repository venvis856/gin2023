package permission_operate

import (
	"encoding/json"
	"errors"
	"fmt"
	"gin/internal/global"
	"gin/internal/library/helper"
	"gin/internal/modules/admin/v1/models"
	"github.com/gogf/gf/util/gconv"
)

/*****************************************权限*****************************************/
// 获取所有权限
func GetAllPermission() []map[string]interface{} {
	var result []map[string]interface{}
	global.DB.Model(&models.Permission{}).Find(&result)
	return result
}

/*****************************************用户-权限*****************************************/
// 用户添加权限
func UserAddPermission(userId int64, permissionCodes []string, identifyId int64) error {
	// 开始事务
	tx := global.DB.Begin()
	model := tx.Model(&models.UserPermission{})
	// 先全部取消列表下对应的权限
	rs := model.Where("user_id = ? and identify_id=?", userId, identifyId).Update("is_effective", 0)
	if rs.Error != nil {
		return errors.New("权限重置失败")
	}
	// 新增权限，存在先更新
	updateDate := map[string]interface{}{
		"is_effective": 1,
	}
	rs = model.Where("user_id = ? and permission_code in ? and identify_id=?", userId, permissionCodes, identifyId).Updates(updateDate)
	if rs.Error != nil {
		tx.Rollback()
		return errors.New("权限更新失败")
	}
	// 找出第一次新增的权限，插入
	var hasPermission []string
	model.Select("permission_code").Where("user_id = ? and identify_id=? ", userId, identifyId).Find(&hasPermission)
	fmt.Println(hasPermission, "=========")
	//return nil
	newPermission := make([]map[string]interface{}, 0)
	for _, v := range permissionCodes {
		if !helper.InArray(v, hasPermission) {
			newPermission = append(newPermission, map[string]interface{}{
				"role_id":         userId,
				"permission_code": v,
				"identify_id":     identifyId,
				"is_effective":    1,
			})
		}
	}
	fmt.Println(newPermission, "newPermission=====")
	model2 := tx.Model(&models.UserPermission{})
	result := model2.Create(newPermission)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	tx.Commit()
	return nil
}

// 用户删除权限
//func UserDeletePermission(userId int64, permissionCode string, identifyId int64) error {
//	result := global.DB.Unscoped().Where("user_id = ? and permission_code=? and identify_id=?", userId, permissionCode, identifyId).Delete(&models.UserPermission{})
//	return result.Error
//}

// 获取用户下直接的权限
func GetPermissionByUser(userId int64, identifyId int64) []map[string]interface{} {
	var result []map[string]interface{}
	global.DB.Model(&models.UserPermission{}).
		Select("permission.id,permission.permission_name,permission.permission_code,permission.type,permission.father_permission_code,permission.identify_id,permission.scene").
		Joins("left join permission on user_permission.permission_code=permission.permission_code and user_permission.identify_id = permission.identify_id").
		Where("user_permission.user_id = ? and permission.status=1 and user_permission.identify_id=? and user_permission.is_effective=1", userId, identifyId).Find(&result)
	return result
}

// 获取用户下所有的权限，包括用户归属角色的权限
func GetAllPermissionByUser(userId int64, identifyId int64) []map[string]interface{} {
	var result []map[string]interface{}
	// 获取用户所在角色
	//var userInfo map[string]interface{}
	userInfo := map[string]interface{}{}
	global.DB.Model(&models.User{}).Where("id = ?", userId).First(&userInfo)
	if len(userInfo) == 0 {
		return nil
	}
	var roleIds []int64
	err := json.Unmarshal(gconv.Bytes(userInfo["role_ids"]), &roleIds)
	if err != nil {
		fmt.Println("json数据转换失败：", err.Error())
		return nil
	}
	// 获取所属角色拥有的权限
	var rolePermissionResult []map[string]interface{}
	isGlobal := false
	if len(roleIds) != 0 {
		var roleRow []map[string]interface{}
		// 判断是否全局权限
		global.DB.Model(&models.Role{}).Where("id in ? and type=1 and status=1 and identify_id=?", roleIds, identifyId).Find(&roleRow)
		if len(roleRow) != 0 {
			isGlobal = true
		}
		if isGlobal {
			global.DB.Model(&models.Permission{}).Select("id,permission_name,permission_code,type,father_permission_code,identify_id,scene").
				Where("status!=9 and identify_id=?", identifyId).Find(&result)
			return result
		}
		for _, roleId := range roleIds {
			rolePermissions := GetPermissionByRole(roleId, identifyId)
			rolePermissionResult = helper.MergeSliceMap(rolePermissionResult, true, rolePermissions)
		}
	}
	// 个人权限
	userPermissionResult := GetPermissionByUser(userId, identifyId)
	result = helper.MergeSliceMap(result, true, rolePermissionResult, userPermissionResult)
	return result
}

// 判断用户是否有权限
func CheckUserHasPermission(userId int64, permissionCode string, identifyId int64) bool {
	userInfo := map[string]interface{}{}
	global.DB.Model(&models.User{}).Where("id = ? and status != 9", userId).First(&userInfo)
	if len(userInfo) == 0 {
		return false
	}
	// 判断用户是否直接有权限
	userPermissionResult := map[string]interface{}{}
	global.DB.Model(&models.UserPermission{}).Where("user_id=? and permission_code=? and identify_id=? and is_effective=1", userId, permissionCode, identifyId).First(&userPermissionResult)
	if userPermissionResult != nil && len(userPermissionResult) != 0 {
		return true
	}
	// 获取用户所在角色
	var roleIds []int
	err := json.Unmarshal(gconv.Bytes(userInfo["role_ids"]), &roleIds)
	if err != nil {
		fmt.Println("json数据转换失败：", err.Error())
		return false
	}
	if len(roleIds) == 0 {
		return false
	}
	// 如果是系统管理员,全部通过
	var globalRoles []map[string]interface{}
	global.DB.Model(&models.Role{}).Where("status =1 and type=1 and identify_id=? and id in ?", identifyId, roleIds).Find(&globalRoles)
	if len(globalRoles) != 0 {
		return true
	}
	fmt.Println(globalRoles, "====globalRoles")
	for _, roleId := range roleIds {
		fmt.Println(CheckRoleHasPermission(int64(roleId), permissionCode, identifyId), "=====11111", roleId, permissionCode)
		if CheckRoleHasPermission(int64(roleId), permissionCode, identifyId) {
			return true
		}
	}
	return false
}

/*****************************************角色-权限*****************************************/

// 判断角色是否有权限
func CheckRoleHasPermission(roleId int64, permissionCode string, identifyId int64) bool {
	//全局角色直接返回true
	roleRow := make(map[string]interface{})
	global.DB.Model(&models.Role{}).Where("id=? and type=1 and status =1 and identify_id=?", roleId, identifyId).First(&roleRow)
	if len(roleRow) != 0 {
		return true
	}
	row := make(map[string]interface{})
	global.DB.Model(&models.RolePermission{}).Where("role_id=? and permission_code=? and is_effective=1 and identify_id=?", gconv.Int64(roleId), permissionCode, identifyId).First(&row)
	fmt.Println(row, "==row===")
	if len(row) != 0 {
		return true
	}
	return false
}

// 角色添加权限,如果已经存在，直接更新 is_effective
func RoleAddPermission(roleId int64, permissionCodes []string, identifyId int64) error {
	// 开始事务
	tx := global.DB.Begin()
	model := tx.Model(&models.RolePermission{})
	// 先全部取消列表下对应的权限
	rs := model.Where("role_id = ? and identify_id=?", roleId, identifyId).Update("is_effective", 0)
	if rs.Error != nil {
		return errors.New("权限重置失败")
	}
	// 新增权限，存在先更新
	updateDate := map[string]interface{}{
		"is_effective": 1,
	}
	rs = model.Where("role_id = ? and permission_code in ? and identify_id=?", roleId, permissionCodes, identifyId).Updates(updateDate)
	if rs.Error != nil {
		tx.Rollback()
		return errors.New("权限更新失败")
	}
	// 找出第一次新增的权限，插入
	var hasPermission []string
	model.Select("permission_code").Where("role_id = ? and identify_id=?", roleId, identifyId).Find(&hasPermission)
	fmt.Println(hasPermission, "=========")
	//return nil
	newPermission := make([]map[string]interface{}, 0)
	for _, v := range permissionCodes {
		if !helper.InArray(v, hasPermission) {
			newPermission = append(newPermission, map[string]interface{}{
				"role_id":         roleId,
				"permission_code": v,
				"identify_id":     identifyId,
				"is_effective":    1,
			})
		}
	}
	fmt.Println(newPermission, "newPermission=====")
	if len(newPermission) > 0 {
		model2 := tx.Model(&models.RolePermission{})
		result := model2.Create(newPermission)
		if result.Error != nil {
			tx.Rollback()
			return result.Error
		}
	}

	tx.Commit()
	return nil
}

// 角色删除权限
//func RoleDeletePermission(roleId int64, permissionCode string, SiteId int64) error {
//  model := global.DB.Model(&models.RolePermission{})
//  updateDate := map[string]interface{}{
//    "is_effective": 0,
//  }
//  result := model.Where("role_id = ? and permission_code=? and site_id=?", roleId, permissionCode, SiteId).Updates(updateDate)
//  return result.Error
//}

// 获取角色下的权限
func GetPermissionByRole(roleId int64, identifyId int64) []map[string]interface{} {
	var result []map[string]interface{}
	global.DB.Model(&models.RolePermission{}).
		Select("permission.id,permission.permission_name,permission.permission_code,permission.type,permission.father_permission_code,permission.identify_id,permission.scene").
		Joins("left join permission on role_permission.permission_code=permission.permission_code and role_permission.identify_id=permission.identify_id ").
		Where("role_permission.role_id = ? and permission.status=1 and permission.identify_id=? and role_permission.is_effective=1", roleId, identifyId).Scan(&result)
	return result
}
