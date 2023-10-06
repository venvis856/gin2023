package identify

import (
	"errors"
	"gin/internal/modules/admin/v1/logic/permission_operate"
	"gin/internal/modules/admin/v1/models"
	"github.com/golang-module/carbon"
	"gorm.io/gorm"
)

type Normal struct {
}

func (a *Normal) InitPermission(identifyId int64, roleId int64, tx *gorm.DB) error {
	permissions := []*models.Permission{
		{PermissionName: "后台", PermissionCode: "web_system", FatherPermissionCode: "", IdentifyId: identifyId, Type: models.TYPE_MENU, Scene: models.SCENE_WEB, Status: 1, CreateTime: carbon.Now().Timestamp()},

		// 角色列表
		{PermissionName: "角色列表", PermissionCode: "web_role_list", FatherPermissionCode: "web_system", IdentifyId: identifyId, Type: models.TYPE_MENU, Scene: models.SCENE_WEB, Status: 1, CreateTime: carbon.Now().Timestamp()},
		{PermissionName: "角色详情", PermissionCode: "web_role_info", FatherPermissionCode: "web_role_list", IdentifyId: identifyId, Type: models.TYPE_NORMAL, Scene: models.SCENE_WEB, Status: 1, CreateTime: carbon.Now().Timestamp()},
		{PermissionName: "角色新增", PermissionCode: "web_role_add", FatherPermissionCode: "web_role_list", IdentifyId: identifyId, Type: models.TYPE_NORMAL, Scene: models.SCENE_WEB, Status: 1, CreateTime: carbon.Now().Timestamp()},
		{PermissionName: "角色修改", PermissionCode: "web_role_update", FatherPermissionCode: "web_role_list", IdentifyId: identifyId, Type: models.TYPE_NORMAL, Scene: models.SCENE_WEB, Status: 1, CreateTime: carbon.Now().Timestamp()},
		{PermissionName: "角色删除", PermissionCode: "web_role_delete", FatherPermissionCode: "web_role_list", IdentifyId: identifyId, Type: models.TYPE_NORMAL, Scene: models.SCENE_WEB, Status: 1, CreateTime: carbon.Now().Timestamp()},
		{PermissionName: "角色修改权限", PermissionCode: "web_role_change_permission", FatherPermissionCode: "web_role_list", IdentifyId: identifyId, Type: models.TYPE_NORMAL, Scene: models.SCENE_WEB, Status: 1, CreateTime: carbon.Now().Timestamp()},

		// 用户列表
		{PermissionName: "用户列表", PermissionCode: "web_user_list", FatherPermissionCode: "web_system", IdentifyId: identifyId, Type: models.TYPE_MENU, Scene: models.SCENE_WEB, Status: 1, CreateTime: carbon.Now().Timestamp()},
		{PermissionName: "用户详情", PermissionCode: "web_user_info", FatherPermissionCode: "web_user_list", IdentifyId: identifyId, Type: models.TYPE_NORMAL, Scene: models.SCENE_WEB, Status: 1, CreateTime: carbon.Now().Timestamp()},
		{PermissionName: "用户新增", PermissionCode: "web_user_add", FatherPermissionCode: "web_user_list", IdentifyId: identifyId, Type: models.TYPE_NORMAL, Scene: models.SCENE_WEB, Status: 1, CreateTime: carbon.Now().Timestamp()},
		{PermissionName: "用户修改", PermissionCode: "web_user_update", FatherPermissionCode: "web_user_list", IdentifyId: identifyId, Type: models.TYPE_NORMAL, Scene: models.SCENE_WEB, Status: 1, CreateTime: carbon.Now().Timestamp()},
		{PermissionName: "用户删除", PermissionCode: "web_user_delete", FatherPermissionCode: "web_user_list", IdentifyId: identifyId, Type: models.TYPE_NORMAL, Scene: models.SCENE_WEB, Status: 1, CreateTime: carbon.Now().Timestamp()},
	}

	if tx == nil {
		return errors.New("事务失败")
	}

	model := tx.Model(models.Permission{})
	result := model.Create(&permissions)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	codes := make([]string, 0)
	for _, v := range permissions {
		codes = append(codes, v.PermissionCode)
	}
	if err := permission_operate.RoleAddPermission(roleId, codes, identifyId); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}