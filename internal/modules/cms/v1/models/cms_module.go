package models

type CmsModule struct {
	Id         uint   `gorm:"column:id" db:"id" json:"id" form:"id"`
	ModuleName string `gorm:"column:module_name" db:"module_name" json:"module_name" form:"module_name"`
	SiteId     uint   `gorm:"column:site_id" db:"site_id" json:"site_id" form:"site_id"`
	Status     int    `gorm:"column:status" db:"status" json:"status" form:"status"` //1正常 5禁用 9删除
	Content    string `gorm:"column:content" db:"content" json:"content" form:"content"`
	CreateTime int    `gorm:"column:create_time" db:"create_time" json:"create_time" form:"create_time"`
	UpdateTime int    `gorm:"column:update_time" db:"update_time" json:"update_time" form:"update_time"`
	DeleteTime int    `gorm:"column:delete_time" db:"delete_time" json:"delete_time" form:"delete_time"`
}
