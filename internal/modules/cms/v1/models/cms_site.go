package models

type CmsSite struct {
	Id                uint   `gorm:"column:id" db:"id" json:"id" form:"id"`
	SiteName          string `gorm:"column:site_name" db:"site_name" json:"site_name" form:"site_name"`
	Root              string `gorm:"column:root" db:"root" json:"root" form:"root"`
	OnlineUrl         string `gorm:"column:online_url" db:"online_url" json:"online_url" form:"online_url"` //现网url
	OnlineImageUrl    string `gorm:"column:online_image_url" db:"online_image_url" json:"online_image_url" form:"online_image_url"`
	PreviewUrl        string `gorm:"column:preview_url" db:"preview_url" json:"preview_url" form:"preview_url"`
	RsyncPasswordPath string `gorm:"column:rsync_password_path" db:"rsync_password_path" json:"rsync_password_path" form:"rsync_password_path"`
	RsyncAddress      string `gorm:"column:rsync_address" db:"rsync_address" json:"rsync_address" form:"rsync_address"`
	RsyncImageAddress string `gorm:"column:rsync_image_address" db:"rsync_image_address" json:"rsync_image_address" form:"rsync_image_address"`
	Status            int    `gorm:"column:status" db:"status" json:"status" form:"status"` //1正常 5禁用 9删除
	CreateTime        int    `gorm:"column:create_time" db:"create_time" json:"create_time" form:"create_time"`
	UpdateTime        int    `gorm:"column:update_time" db:"update_time" json:"update_time" form:"update_time"`
	DeleteTime        int    `gorm:"column:delete_time" db:"delete_time" json:"delete_time" form:"delete_time"`
}
