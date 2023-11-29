package permission_operate

import (
	"errors"
	"fmt"
	"gin/internal/global"
	"gin/internal/library/helper"
	"gin/internal/modules/admin/v1/config"
	"gin/internal/modules/admin/v1/models"
	"gin/internal/modules/admin/v1/service"
	"github.com/gogf/gf/util/gconv"
)

type permissionOperateLogic struct{}

func init() {
	service.RegisterPermissionOperate(New())
}

func New() service.PermissionOperateInterface {
	return &permissionOperateLogic{}
}

/*****************************************权限*****************************************/
// 获取所有权限
func (*permissionOperateLogic) GetAllPermission() []map[string]interface{} {
	var result []map[string]interface{}
	global.DB.Model(&models.Permission{}).Find(&result)
	return result
}

// 获取权限的code
func (*permissionOperateLogic) GetAllPermissionCodes() []string {
	var result []string
	global.DB.Model(&models.Permission{}).Pluck("permission_code", &result)
	return result
}

// 获取某个 identify 下的所有权限
func (*permissionOperateLogic) GetAllPermissionByIdentify(identifyId int64) []map[string]interface{} {
	var result []map[string]interface{}
	global.DB.Model(&models.IdentifyPermission{}).Where("identify_id=? and is_effective = ?", identifyId, config.EFFECTIVE_YES).Find(&result)
	return result
}

// identify 添加权限
func (a *permissionOperateLogic) IdentifyAddPermission(identifyId int64, permissionCodes []string) error {
	// 开始事务
	tx := global.DB.Begin()
	model := tx.Model(&models.IdentifyPermission{})
	// 先全部取消列表下对应的权限
	rs := model.Where("identify_id = ? ", identifyId).Update("is_effective", 0)
	if rs.Error != nil {
		return errors.New("权限重置失败")
	}
	// 新增权限，存在先更新
	updateDate := map[string]interface{}{
		"is_effective": config.EFFECTIVE_YES,
	}
	permissionIds := a.GetPermissionIdsByPermissionCode(permissionCodes)
	rs = model.Where("permission_id in ? and identify_id=?", permissionIds, identifyId).Updates(updateDate)
	if rs.Error != nil {
		tx.Rollback()
		return errors.New("权限更新失败")
	}
	// 找出第一次新增的权限，插入
	var hasPermission []int64
	tx.Model(&models.IdentifyPermission{}).Where("identify_id=?", identifyId).Pluck("permission_id", &hasPermission)
	//fmt.Println(hasPermission, "=========")
	//return nil
	newPermission := make([]map[string]interface{}, 0)
	for _, v := range permissionIds {
		if !helper.InArray(v, hasPermission) {
			newPermission = append(newPermission, map[string]interface{}{
				"permission_id": v,
				"identify_id":   identifyId,
				"is_effective":  config.EFFECTIVE_YES,
			})
		}
	}
	//fmt.Println(newPermission, "newPermission=====")
	if len(newPermission) > 0 {
		model2 := tx.Model(&models.IdentifyPermission{})
		result := model2.Create(newPermission)
		if result.Error != nil {
			tx.Rollback()
			return result.Error
		}
	}

	tx.Commit()
	return nil
}

/*****************************************用户-权限*****************************************/
// 用户添加权限
func (a *permissionOperateLogic) UserAddPermission(userId int64, permissionCodes []string, identifyId int64) error {
	// 开始事务
	tx := global.DB.Begin()
	model := tx.Model(&models.UserPermission{})
	// 先全部取消列表下对应的权限
	rs := model.Where("user_id = ? and identify_id=?", userId, identifyId).Update("is_effective", config.EFFECTIVE_NO)
	if rs.Error != nil {
		return errors.New("权限重置失败")
	}
	// 新增权限，存在先更新
	updateDate := map[string]interface{}{
		"is_effective": config.EFFECTIVE_YES,
	}
	permissionIds := a.GetPermissionIdsByPermissionCode(permissionCodes)
	rs = model.Where("user_id = ? and permission_id in ? and identify_id=?", userId, permissionIds, identifyId).Updates(updateDate)
	if rs.Error != nil {
		tx.Rollback()
		return errors.New("权限更新失败")
	}
	// 找出第一次新增的权限，插入
	var hasPermission []int64
	model.Select("permission_id").Where("user_id = ? and identify_id=? ", userId, identifyId).Find(&hasPermission)
	//fmt.Println(hasPermission, "=========")
	//return nil
	newPermission := make([]map[string]interface{}, 0)
	for _, v := range permissionCodes {
		if !helper.InArray(v, hasPermission) {
			newPermission = append(newPermission, map[string]interface{}{
				"user_id":       userId,
				"permission_id": v,
				"identify_id":   identifyId,
				"is_effective":  config.EFFECTIVE_YES,
			})
		}
	}
	//fmt.Println(newPermission, "newPermission=====")
	model2 := tx.Model(&models.UserPermission{})
	result := model2.Create(newPermission)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	tx.Commit()
	return nil
}

// 获取用户下直接的权限
func (*permissionOperateLogic) GetPermissionByUser(userId int64, identifyId int64) []map[string]interface{} {
	var result []map[string]interface{}
	global.DB.Model(&models.UserPermission{}).
		Select("permission.id,permission.permission_name,permission.permission_code,permission.type,permission.father_permission_code,permission.identify_id").
		Joins("left join permission on user_permission.permission_id=permission.id").
		Where("user_permission.user_id = ? and permission.status=1 and user_permission.identify_id=? and user_permission.is_effective=1", userId, identifyId).Find(&result)
	return result
}

// 获取用户下所有的权限，包括用户归属角色的权限
func (a *permissionOperateLogic) GetAllPermissionByUser(userId int64, identifyId int64) []map[string]interface{} {
	var result []map[string]interface{}
	// 获取用户所在角色
	//var userInfo map[string]interface{}
	userInfo := map[string]interface{}{}
	global.DB.Model(&models.User{}).Where("id = ?", userId).First(&userInfo)
	if len(userInfo) == 0 {
		return nil
	}
	var roleIds []int64
	global.DB.Model(models.UserRole{}).Where("user_id = ? and identify_id=?", userId, identifyId).Pluck("role_id", &roleIds)
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
			global.DB.Model(&models.Permission{}).Select("id,permission_name,permission_code,type,father_permission_code,identify_id").
				Where("status!=9 and identify_id=?", identifyId).Find(&result)
			return result
		}
		for _, roleId := range roleIds {
			rolePermissions := a.GetPermissionByRole(roleId, identifyId)
			rolePermissionResult = helper.MergeSliceMap(rolePermissionResult, true, rolePermissions)
		}
	}
	// 个人权限
	userPermissionResult := a.GetPermissionByUser(userId, identifyId)
	result = helper.MergeSliceMap(result, true, rolePermissionResult, userPermissionResult)
	return result
}

// 判断用户是否有权限
func (a *permissionOperateLogic) CheckUserHasPermission(userId int64, permissionCode string, identifyId int64) bool {
	var userInfo models.User
	global.DB.Model(&models.User{}).Where("id = ? and status != 9", userId).First(&userInfo)
	if userInfo.ID == 0 {
		return false
	}
	// 判断用户是否直接有权限
	permissionIds := a.GetPermissionIdsByPermissionCode([]string{permissionCode})
	if len(permissionIds) == 0 {
		return false
	}
	permissionId := permissionIds[0]

	userPermissionResult := map[string]interface{}{}
	global.DB.Model(&models.UserPermission{}).Where("user_id=? and permission_id=? and identify_id=? and is_effective=1", userId, permissionId, identifyId).First(&userPermissionResult)
	if userPermissionResult != nil && len(userPermissionResult) != 0 {
		return true
	}
	// 获取用户所在角色
	var roleIds []int
	global.DB.Model(models.UserRole{}).Where("user_id = ? and identify_id=?", userId, identifyId).Pluck("role_id", &roleIds)
	if len(roleIds) == 0 {
		return false
	}
	// 如果是系统管理员,全部通过
	var globalRoles []map[string]interface{}
	global.DB.Model(&models.Role{}).Where("status =1 and type=? and identify_id=? and id in ?", models.ROLE_TYPE_SYSTEM, identifyId, roleIds).Find(&globalRoles)
	if len(globalRoles) != 0 {
		return true
	}
	fmt.Println(globalRoles, "====globalRoles")
	for _, roleId := range roleIds {
		//fmt.Println(CheckRoleHasPermission(int64(roleId), permissionCode, identifyId), "=====11111", roleId, permissionCode)
		if a.CheckRoleHasPermission(int64(roleId), permissionId, identifyId) {
			return true
		}
	}
	return false
}

/*****************************************角色-权限*****************************************/

// 判断角色是否有权限
func (*permissionOperateLogic) CheckRoleHasPermission(roleId int64, permissionId int64, identifyId int64) bool {
	//全局角色直接返回true
	var roleInfo models.Role
	global.DB.Model(&models.Role{}).Where("id=? and type=? and status =1 and identify_id=?", roleId, models.ROLE_TYPE_SYSTEM, identifyId).First(&roleInfo)
	if roleInfo.ID != 0 {
		return true
	}
	// 非全局角色
	var row models.RolePermission
	global.DB.Model(&models.RolePermission{}).Where("role_id=? and permission_id=? and is_effective=1 and identify_id=?", gconv.Int64(roleId), permissionId, identifyId).First(&row)
	//fmt.Println(row, "==row===")
	if row.ID != 0 {
		return true
	}
	return false
}

// 角色添加权限,如果已经存在，直接更新 is_effective
func (a *permissionOperateLogic) RoleAddPermission(roleId int64, permissionCodes []string, identifyId int64) error {
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
		"is_effective": config.EFFECTIVE_YES,
	}
	permissionIds := a.GetPermissionIdsByPermissionCode(permissionCodes)
	rs = model.Where("role_id = ? and permission_id in ? and identify_id=?", roleId, permissionIds, identifyId).Updates(updateDate)
	if rs.Error != nil {
		tx.Rollback()
		return errors.New("权限更新失败")
	}
	// 找出第一次新增的权限，插入
	var hasPermission []int64
	tx.Model(&models.RolePermission{}).Where("role_id = ? and identify_id=?", roleId, identifyId).Pluck("permission_id", &hasPermission)
	//fmt.Println(hasPermission, "=========")
	//return nil
	newPermission := make([]map[string]interface{}, 0)
	for _, v := range permissionIds {
		if !helper.InArray(v, hasPermission) {
			newPermission = append(newPermission, map[string]interface{}{
				"role_id":       roleId,
				"permission_id": v,
				"identify_id":   identifyId,
				"is_effective":  config.EFFECTIVE_YES,
			})
		}
	}
	//fmt.Println(newPermission, "newPermission=====")
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

// 获取角色下的权限
func (*permissionOperateLogic) GetPermissionByRole(roleId int64, identifyId int64) []map[string]interface{} {
	var result []map[string]interface{}
	global.DB.Model(&models.RolePermission{}).
		Select("permission.id,permission.permission_name,permission.permission_code,permission.type,permission.father_permission_code,permission.identify_id").
		Joins("left join permission on role_permission.permission_id=permission.id").
		Where("role_permission.role_id = ? and permission.status=1 and role_permission.identify_id=? and role_permission.is_effective=1", roleId, identifyId).Scan(&result)
	return result
}

// permission code  to permission id
func (*permissionOperateLogic) GetPermissionIdsByPermissionCode(permissionCodes []string) []int64 {
	var permissionIds []int64
	global.DB.Model(&models.Permission{}).Where("permission_code in ? ", permissionCodes).Pluck("id", &permissionIds)
	return permissionIds
}