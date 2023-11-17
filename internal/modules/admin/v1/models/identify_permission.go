package models

// IdentifyPermission  标识权限表
type IdentifyPermission struct {
	ID           int64 `gorm:"column:id" json:"id"`                       //  id
	IdentifyId   int64 `gorm:"column:identify_id" json:"identify_id"`     //  标识id
	PermissionId int64 `gorm:"column:permission_id" json:"permission_id"` //  权限code
	IsEffective  int64 `gorm:"column:is_effective" json:"is_effective"`   //  是否有效
}
