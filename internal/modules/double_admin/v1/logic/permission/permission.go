package permission

import (
	"errors"
	"gin/api/double_admin/v1/permission/v1"
	"gin/internal/global"
	"gin/internal/modules/double_admin/v1/logic/common"
	"gin/internal/modules/double_admin/v1/logic/permission_operate"
	"gin/internal/modules/double_admin/v1/models"
	"gin/internal/modules/double_admin/v1/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon"
)

type (
	permissionLogic struct{}
)

func init() {
	service.RegisterPermission(New())
}

func New() service.PermissionInterface {
	return &permissionLogic{}
}

func (a *permissionLogic) CheckAuth(c *gin.Context, authCode string, IdentifyId int64) bool {
	userInfo := service.User().GetUserInfo(c)
	return permission_operate.CheckUserHasPermission(int64(userInfo.UserId), authCode, IdentifyId)
}

func (a *permissionLogic) Items(param v1.ItemReq) (map[string]interface{}, error) {
	model := global.DB.Model(&models.Permission{})
	model = common.WhereBySearch(model, param.Search)
	model.Select("id,permission_name,permission_code,type,father_permission_code,status,create_time,update_time")
	model.Where("status != ? and identify_id=?", 9, param.IdentifyId)

	var count int64
	model.Count(&count)

	if param.Limit != 0 {
		if param.PageIndex == 0 {
			param.PageIndex = 1
		}
		model.Offset((param.PageIndex - 1) * param.Limit).Limit(param.Limit)
	}
	if param.OrderBy != "" && param.OrderByType != "" {
		model.Order(param.OrderBy + " " + param.OrderByType)
	}
	var result []map[string]interface{}
	model.Find(&result)

	return map[string]interface{}{"items": result, "total": count}, nil
}

func (a *permissionLogic) Create(param v1.CreateReq) (int64, error) {
	data := models.Permission{
		PermissionName:       param.PermissionName,
		PermissionCode:       param.PermissionCode,
		Type:                 param.Type,
		FatherPermissionCode: param.FatherPermissionCode,
		IdentifyId:           param.IdentifyId,
		Status:               param.Status,
		CreateTime:           carbon.Now().Timestamp(),
	}

	rs := global.DB.Model(&models.Permission{}).Create(&data)
	if rs.Error != nil {
		return 0, rs.Error
	}
	return rs.RowsAffected, nil
}

func (a *permissionLogic) Update(param v1.UpdateReq) (int64, error) {
	data := models.Permission{
		PermissionName:       param.PermissionName,
		PermissionCode:       param.PermissionCode,
		Type:                 param.Type,
		FatherPermissionCode: param.FatherPermissionCode,
		IdentifyId:           param.IdentifyId,
		Status:               param.Status,
		UpdateTime:           carbon.Now().Timestamp(),
	}
	result := global.DB.Model(&models.Permission{}).Where("id = ? and status !=?", param.Id, 9).Updates(&data)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func (a *permissionLogic) Delete(c *gin.Context, param v1.DeleteReq) (int64, error) {
	permissionInfo := map[string]interface{}{}
	global.DB.Model(&models.Permission{}).Where("status != ?", 9).First(&permissionInfo, param.Id)
	if len(permissionInfo) == 0 {
		return 0, errors.New("无该权限")
	}
	// 删除用户权限
	rs := global.DB.Unscoped().Where("permission_code = ? and identify_id=?", permissionInfo["permission_code"], permissionInfo["identify_id"]).Delete(&models.UserPermission{})
	if rs.Error != nil {
		return 0, rs.Error
	}
	// 删除角色权限
	rs = global.DB.Unscoped().Where("permission_code = ? and identify_id=? ", permissionInfo["permission_code"], permissionInfo["identify_id"]).Delete(&models.RolePermission{})
	if rs.Error != nil {
		return 0, rs.Error
	}
	// 删除权限
	rs = global.DB.Unscoped().Where("id = ?", param.Id).Delete(&models.Permission{})
	if rs.Error != nil {
		return 0, rs.Error
	}
	// 日志
	go func() {
		userInfo := service.User().GetUserInfo(c)
		global.Logger.Daily("delete_permision", "info", map[string]interface{}{
			"delete_user":     userInfo,
			"permission_info": permissionInfo,
		})
	}()
	return rs.RowsAffected, nil
}
