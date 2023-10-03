package models

type RolePermission struct {
	ID             int64  `json:"id" gorm:"id"`                           // ID
	RoleId         int64  `json:"role_id" gorm:"role_id"`                 // 角色id
	PermissionCode string `json:"permission_code" gorm:"permission_code"` // 权限code
	IdentifyId     int64  `json:"identify_id" gorm:"identify_id"`         // 标识id
	IsEffective    int64  `json:"is_effective" gorm:"is_effective"`       // 是否有效
}
