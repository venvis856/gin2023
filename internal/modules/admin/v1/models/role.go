package models

const (
	ROLE_TYPE_SYSTEM = 1
)

// Role  角色表
type Role struct {
	ID         int64  `gorm:"column:id" json:"id"`                   //  id
	Uid        int64  `gorm:"column:uid" json:"uid"`                 //  序号
	RoleName   string `gorm:"column:role_name" json:"role_name"`     //  角色名称
	Status     int8  `gorm:"column:status" json:"status"`           //  1：正常 5 禁用  9 删除
	Type       int8  `gorm:"column:type" json:"type"`               //  1：系统管理员 2 超级管理员 3普通角色
	CreateTime int64  `gorm:"column:create_time" json:"create_time"` //  用户创建时间
	UpdateTime int64  `gorm:"column:update_time" json:"update_time"` //  修改时间
	DeleteTime int64  `gorm:"column:delete_time" json:"delete_time"` //  删除时间
	IdentifyId int64  `gorm:"column:identify_id" json:"identify_id"` //  标识id
}
