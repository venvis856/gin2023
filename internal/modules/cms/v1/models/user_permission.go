package models

type UserPermission struct {
	Id int `gorm:"column:id" db:"id" json:"id" form:"id"` //ID
	UserId int `gorm:"column:user_id" db:"user_id" json:"user_id" form:"user_id"` //用户id
	PermissionCode string `gorm:"column:permission_code" db:"permission_code" json:"permission_code" form:"permission_code"` //权限code
	SiteId int `gorm:"column:site_id" db:"site_id" json:"site_id" form:"site_id"`
}