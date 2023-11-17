package models

// RolePermission  角色权限表
type RolePermission struct {
	ID           int64 `gorm:"column:id" json:"id"`                       //  id
	IdentifyId   int64 `gorm:"column:identify_id" json:"identify_id"`     //  标识id
	RoleId       int64 `gorm:"column:role_id" json:"role_id"`             //  角色id
	PermissionId int64 `gorm:"column:permission_id" json:"permission_id"` //  权限id
	IsEffective  int64 `gorm:"column:is_effective" json:"is_effective"`   //  是否有效
}
