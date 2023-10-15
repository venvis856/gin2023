package models

type Role struct {
	ID         int64  `json:"id" gorm:"id"` // ID
	Vid        int64  `json:"vid" gorm:"vid"`
	RoleName   string `json:"role_name" gorm:"role_name"`     // 角色名称
	Status     int8   `json:"status" gorm:"status"`           // 1：正常 5 禁用  9 删除
	Type       int8   `json:"type" gorm:"type"`               // 1：系统管理员 2 超级管理员 3普通角色
	CreateTime int64  `json:"create_time" gorm:"create_time"` // 用户创建时间
	UpdateTime int64  `json:"update_time" gorm:"update_time"` // 修改时间
	DeleteTime int64  `json:"delete_time" gorm:"delete_time"` // 删除时间
	IdentifyId int64  `json:"identify_id" gorm:"identify_id"` // 标识id
}
