package models

type Permission struct {
	Id                   int    `gorm:"column:id" db:"id" json:"id" form:"id"`                                                                                 //ID
	PermissionName       string `gorm:"column:permission_name" db:"permission_name" json:"permission_name" form:"permission_name"`                             //权限名称
	PermissionCode       string `gorm:"column:permission_code" db:"permission_code" json:"permission_code" form:"permission_code"`                             //权限code
	SiteId               int    `gorm:"column:site_id" db:"site_id" json:"site_id" form:"site_id"`                                                             //站点id
	Type                 int8   `gorm:"column:type" db:"type" json:"type" form:"type"`                                                                         //类型 1站点 2菜单 3普通权限
	FatherPermissionCode string `gorm:"column:father_permission_code" db:"father_permission_code" json:"father_permission_code" form:"father_permission_code"` //父权限code
	Status               int8   `gorm:"column:status" db:"status" json:"status" form:"status"`                                                                 //1：正常 5 禁用  9 删除
	CreateTime           int    `gorm:"column:create_time" db:"create_time" json:"create_time" form:"create_time"`                                             //用户创建时间
	UpdateTime           int    `gorm:"column:update_time" db:"update_time" json:"update_time" form:"update_time"`                                             //修改时间
	DeleteTime           int    `gorm:"column:delete_time" db:"delete_time" json:"delete_time" form:"delete_time"`                                             //删除时间
}
