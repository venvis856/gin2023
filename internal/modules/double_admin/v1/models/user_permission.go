package models

type UserPermission struct {
	ID             int64  `json:"id" gorm:"id"`                           // ID
	UserId         int64  `json:"user_id" gorm:"user_id"`                 // 用户id
	PermissionCode string `json:"permission_code" gorm:"permission_code"` // 权限code
	IdentifyId     int64  `json:"identify_id" gorm:"identify_id"`         // 标识id
	IsEffective    int8   `json:"is_effective" gorm:"is_effective"`       // 是否有效
}
