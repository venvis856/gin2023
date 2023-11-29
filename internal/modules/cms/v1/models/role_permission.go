package models

type RolePermission struct {
	Id int `gorm:"column:id" db:"id" json:"id" form:"id"` //ID
	RoleId int `gorm:"column:role_id" db:"role_id" json:"role_id" form:"role_id"` //角色id
	PermissionCode string `gorm:"column:permission_code" db:"permission_code" json:"permission_code" form:"permission_code"` //权限code
	SiteId uint `gorm:"column:site_id" db:"site_id" json:"site_id" form:"site_id"`
	IsEffective uint `gorm:"column:is_effective" db:"is_effective" json:"is_effective" form:"is_effective"` //是否有效
}