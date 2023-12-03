package models

type CmsAuditDetail struct {
	Id         uint   `gorm:"column:id" db:"id" json:"id" form:"id"`
	LocalUrl   string `gorm:"column:local_url" db:"local_url" json:"local_url" form:"local_url"`
	OnlineUrl  string `gorm:"column:online_url" db:"online_url" json:"online_url" form:"online_url"`
	PreviewUrl string `gorm:"column:preview_url" db:"preview_url" json:"preview_url" form:"preview_url"`
	FileUrl    string `gorm:"column:file_url" db:"file_url" json:"file_url" form:"file_url"`
	Type       uint8  `gorm:"column:type" db:"type" json:"type" form:"type"`         //1 页面 2图片
	Status     uint8  `gorm:"column:status" db:"status" json:"status" form:"status"` //1 正常 9 删除
	AuditId    uint   `gorm:"column:audit_id" db:"audit_id" json:"audit_id" form:"audit_id"`
	PageId     int    `gorm:"column:page_id" db:"page_id" json:"page_id" form:"page_id"` //页面id
	MakeTime   uint   `gorm:"column:make_time" db:"make_time" json:"make_time" form:"make_time"`
	DeleteTime int    `gorm:"column:delete_time" db:"delete_time" json:"delete_time" form:"delete_time"`
}
