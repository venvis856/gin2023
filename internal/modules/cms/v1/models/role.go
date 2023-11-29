package models

type Role struct {
	Id         int    `gorm:"column:id" db:"id" json:"id" form:"id"`                                     //ID
	RoleName   string `gorm:"column:role_name" db:"role_name" json:"role_name" form:"role_name"`         //角色名称
	Status     int8   `gorm:"column:status" db:"status" json:"status" form:"status"`                     //1：正常 5 禁用  9 删除
	CreateTime int    `gorm:"column:create_time" db:"create_time" json:"create_time" form:"create_time"` //用户创建时间
	UpdateTime int    `gorm:"column:update_time" db:"update_time" json:"update_time" form:"update_time"` //修改时间
	DeleteTime int    `gorm:"column:delete_time" db:"delete_time" json:"delete_time" form:"delete_time"` //删除时间
}
