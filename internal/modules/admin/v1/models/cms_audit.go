package models

type CmsAudit struct {
	Id            uint   `gorm:"column:id" db:"id" json:"id" form:"id"`
	FirstUrl      string `gorm:"column:first_url" db:"first_url" json:"first_url" form:"first_url"`
	Count         int    `gorm:"column:count" db:"count" json:"count" form:"count"`
	Type          int    `gorm:"column:type" db:"type" json:"type" form:"type"`                                             //1 页面 2 图片
	MakeUserId    int    `gorm:"column:make_user_id" db:"make_user_id" json:"make_user_id" form:"make_user_id"`             //生成用户id
	PublushUserId uint   `gorm:"column:publush_user_id" db:"publush_user_id" json:"publush_user_id" form:"publush_user_id"` //发布人id
	Status        uint   `gorm:"column:status" db:"status" json:"status" form:"status"`                                     //1 待发布 2 已发布 9 删除
	SiteId        uint   `gorm:"column:site_id" db:"site_id" json:"site_id" form:"site_id"`
	MakeTime      uint   `gorm:"column:make_time" db:"make_time" json:"make_time" form:"make_time"`
	PublushTime   uint   `gorm:"column:publush_time" db:"publush_time" json:"publush_time" form:"publush_time"`
	DeleteTime    int    `gorm:"column:delete_time" db:"delete_time" json:"delete_time" form:"delete_time"`
}
