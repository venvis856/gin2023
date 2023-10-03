package identify

import (
	"errors"
	"gin/internal/library/permission"
	"gin/internal/modules/admin/v1/models"
	"github.com/golang-module/carbon"
	"gorm.io/gorm"
)

type Police struct{}

func (*Police) InitPermission(identifyId int64, roleId int64, tx *gorm.DB) error {
	permissions := []*models.Permission{
		/*************************** web后台 ****************************/
		{PermissionName: "网页后台", PermissionCode: "web_system", FatherPermissionCode: "", IdentifyId: identifyId, Type: models.TYPE_MENU, Scene: models.SCENE_WEB, Status: 1, CreateTime: carbon.Now().Timestamp()},

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

		// 摄像头列表
		{PermissionName: "摄像头列表", PermissionCode: "web_camera_list", FatherPermissionCode: "web_system", IdentifyId: identifyId, Type: models.TYPE_MENU, Scene: models.SCENE_WEB, Status: 1, CreateTime: carbon.Now().Timestamp()},

		// 报警消息
		{PermissionName: "报警消息列表", PermissionCode: "web_alarm_list", FatherPermissionCode: "web_system", IdentifyId: identifyId, Type: models.TYPE_MENU, Scene: models.SCENE_WEB, Status: 1, CreateTime: carbon.Now().Timestamp()},

		// 报警配置设置
		{PermissionName: "报警设置列表", PermissionCode: "web_alarm_config_list", FatherPermissionCode: "web_system", IdentifyId: identifyId, Type: models.TYPE_MENU, Scene: models.SCENE_WEB, Status: 1, CreateTime: carbon.Now().Timestamp()},
		{PermissionName: "报警设置详情", PermissionCode: "web_alarm_config_info", FatherPermissionCode: "web_alarm_config_list", IdentifyId: identifyId, Type: models.TYPE_NORMAL, Scene: models.SCENE_WEB, Status: 1, CreateTime: carbon.Now().Timestamp()},
		{PermissionName: "报警设置新增", PermissionCode: "web_alarm_config_add", FatherPermissionCode: "web_alarm_config_list", IdentifyId: identifyId, Type: models.TYPE_NORMAL, Scene: models.SCENE_WEB, Status: 1, CreateTime: carbon.Now().Timestamp()},
		{PermissionName: "报警设置修改", PermissionCode: "web_alarm_config_update", FatherPermissionCode: "web_alarm_config_list", IdentifyId: identifyId, Type: models.TYPE_NORMAL, Scene: models.SCENE_WEB, Status: 1, CreateTime: carbon.Now().Timestamp()},
		{PermissionName: "报警设置删除", PermissionCode: "web_alarm_config_delete", FatherPermissionCode: "web_alarm_config_list", IdentifyId: identifyId, Type: models.TYPE_NORMAL, Scene: models.SCENE_WEB, Status: 1, CreateTime: carbon.Now().Timestamp()},

		/*************************** app ****************************/
		{PermissionName: "app", PermissionCode: "app_system", FatherPermissionCode: "", IdentifyId: identifyId, Type: models.TYPE_MENU, Scene: models.SCENE_APP, Status: 1, CreateTime: carbon.Now().Timestamp()},

		//// 摄像头
		//{PermissionName: "app摄像头列表", PermissionCode: "app_camera_list", FatherPermissionCode: "app_system", IdentifyId: identifyId, Type:models.TYPE_MENU, Scene: models.SCENE_APP, Status: 1, CreateTime: carbon.Now().Timestamp()},
		//{PermissionName: "app摄像头详情", PermissionCode: "app_camera_info", FatherPermissionCode: "app_camera_list", IdentifyId: identifyId, Type: models.TYPE_NORMAL, Scene: models.SCENE_APP, Status: 1, CreateTime: carbon.Now().Timestamp()},
		//
		//// 盒子节点
		//{PermissionName: "app节点列表", PermissionCode: "app_box_list", FatherPermissionCode: "app_system", IdentifyId: identifyId, Type:models.TYPE_MENU, Scene: models.SCENE_APP, Status: 1, CreateTime: carbon.Now().Timestamp()},
		//{PermissionName: "app节点详情", PermissionCode: "app_box_info", FatherPermissionCode: "app_box_list", IdentifyId: identifyId, Type: models.TYPE_NORMAL, Scene: models.SCENE_APP, Status: 1, CreateTime: carbon.Now().Timestamp()},
		//
		//// app搜图
		//{PermissionName: "app搜图", PermissionCode: "app_picture_list", FatherPermissionCode: "app_system", IdentifyId: identifyId, Type:models.TYPE_MENU, Scene: models.SCENE_APP, Status: 1, CreateTime: carbon.Now().Timestamp()},

		/*************************** 派出所 ****************************/
		{PermissionName: "大屏", PermissionCode: "qt_system", FatherPermissionCode: "", IdentifyId: identifyId, Type: models.TYPE_MENU, Scene: models.SCENE_QT, Status: 1, CreateTime: carbon.Now().Timestamp()},
		//
		//// 摄像头
		//{PermissionName: "大屏摄像头列表", PermissionCode: "qt_camera_list", FatherPermissionCode: "qt_system", IdentifyId: identifyId, Type:models.TYPE_MENU, Scene: models.SCENE_QT, Status: 1, CreateTime: carbon.Now().Timestamp()},
		//{PermissionName: "大屏摄像头详情", PermissionCode: "qt_camera_info", FatherPermissionCode: "qt_camera_list", IdentifyId: identifyId, Type: models.TYPE_NORMAL, Scene: models.SCENE_QT, Status: 1, CreateTime: carbon.Now().Timestamp()},
		//
		//// 盒子节点
		//{PermissionName: "大屏节点列表", PermissionCode: "qt_box_list", FatherPermissionCode: "qt_system", IdentifyId: identifyId, Type:models.TYPE_MENU, Scene: models.SCENE_QT, Status: 1, CreateTime: carbon.Now().Timestamp()},
		//{PermissionName: "大屏节点详情", PermissionCode: "qt_box_info", FatherPermissionCode: "qt_box_list", IdentifyId: identifyId, Type: models.TYPE_NORMAL, Scene: models.SCENE_QT, Status: 1, CreateTime: carbon.Now().Timestamp()},
		//
		////  大屏搜图
		//{PermissionName: "大屏搜图", PermissionCode: "qt_picture_list", FatherPermissionCode: "qt_system", IdentifyId: identifyId, Type:models.TYPE_MENU, Scene: models.SCENE_QT, Status: 1, CreateTime: carbon.Now().Timestamp()},
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
	if err := permission.RoleAddPermission(roleId, codes, identifyId); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
