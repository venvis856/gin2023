package models

type CmsClassify struct {
	Id           uint   `gorm:"column:id" db:"id" json:"id" form:"id"`
	ClassifyName string `gorm:"column:classify_name" db:"classify_name" json:"classify_name" form:"classify_name"`
	IsHowto      int8   `gorm:"column:is_howto" db:"is_howto" json:"is_howto" form:"is_howto"` //1 是 2 否
	SiteId       int    `gorm:"column:site_id" db:"site_id" json:"site_id" form:"site_id"`
	Status       int    `gorm:"column:status" db:"status" json:"status" form:"status"` //1正常 5禁用 9删除
	CreateTime   int    `gorm:"column:create_time" db:"create_time" json:"create_time" form:"create_time"`
	UpdateTime   int    `gorm:"column:update_time" db:"update_time" json:"update_time" form:"update_time"`
	DeleteTime   int    `gorm:"column:delete_time" db:"delete_time" json:"delete_time" form:"delete_time"`
}
