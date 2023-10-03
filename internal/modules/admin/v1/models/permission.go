package models

const (
	TYPE_MENU   = 1
	TYPE_NORMAL = 2

	SCENE_WEB = 1
	SCENE_APP = 2
	SCENE_QT  = 3
)

type Permission struct {
	ID                   int64  `json:"id" gorm:"id"`                                         // ID
	PermissionName       string `json:"permission_name" gorm:"permission_name"`               // 权限名称
	PermissionCode       string `json:"permission_code" gorm:"permission_code"`               // 权限code
	IdentifyId           int64  `json:"identify_id" gorm:"identify_id"`                       // 标识id
	Type                 int8   `json:"type" gorm:"type"`                                     // 类型 1菜单 2普通权限
	Scene                int8   `json:"scene" gorm:"scene"`                                   // 场景 1 后台 2 app 3 大屏
	FatherPermissionCode string `json:"father_permission_code" gorm:"father_permission_code"` // 父权限code
	Status               int8   `json:"status" gorm:"status"`                                 // 1：正常 5 禁用  9 删除
	CreateTime           int64  `json:"create_time" gorm:"create_time"`                       // 用户创建时间
	UpdateTime           int64  `json:"update_time" gorm:"update_time"`                       // 修改时间
	DeleteTime           int64  `json:"delete_time" gorm:"delete_time"`                       // 删除时间
}
