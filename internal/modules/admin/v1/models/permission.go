package models

const (
	TYPE_MENU   = 1
	TYPE_NORMAL = 2

	SCENE_WEB = 1
	SCENE_APP = 2
	SCENE_QT  = 3
)

// Permission  权限表
type Permission struct {
	ID                   int64  `gorm:"column:id" json:"id"`                                         //  id
	PermissionName       string `gorm:"column:permission_name" json:"permission_name"`               //  权限名称
	PermissionCode       string `gorm:"column:permission_code" json:"permission_code"`               //  权限code
	Type                 int8  `gorm:"column:type" json:"type"`                                     //  类型 1菜单 2普通权限
	Scene                int8  `gorm:"column:scene" json:"scene"`                                   //  场景 1 后台 2 app 3 大屏
	FatherPermissionCode string `gorm:"column:father_permission_code" json:"father_permission_code"` //  父权限code
	Status               int8  `gorm:"column:status" json:"status"`                                 //  1：正常 5 禁用  9 删除
	CreateTime           int64  `gorm:"column:create_time" json:"create_time"`                       //  用户创建时间
	UpdateTime           int64  `gorm:"column:update_time" json:"update_time"`                       //  修改时间
	DeleteTime           int64  `gorm:"column:delete_time" json:"delete_time"`                       //  删除时间
}
