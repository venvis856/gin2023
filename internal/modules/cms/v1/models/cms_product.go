package models

type CmsProduct struct {
	Id uint `gorm:"column:id" db:"id" json:"id" form:"id"`
	ProductName string `gorm:"column:product_name" db:"product_name" json:"product_name" form:"product_name"`
	DownloadUrl string `gorm:"column:download_url" db:"download_url" json:"download_url" form:"download_url"`
	BuyUrl string `gorm:"column:buy_url" db:"buy_url" json:"buy_url" form:"buy_url"`
	SiteId int `gorm:"column:site_id" db:"site_id" json:"site_id" form:"site_id"`
	Status int `gorm:"column:status" db:"status" json:"status" form:"status"` //1正常 5禁用 9删除
	CreateTime int `gorm:"column:create_time" db:"create_time" json:"create_time" form:"create_time"`
	UpdateTime int `gorm:"column:update_time" db:"update_time" json:"update_time" form:"update_time"`
	DeleteTime int `gorm:"column:delete_time" db:"delete_time" json:"delete_time" form:"delete_time"`
}