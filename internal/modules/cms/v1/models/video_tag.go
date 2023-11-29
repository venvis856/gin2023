package models

// VideoTag undefined
type VideoTag struct {
	ID          int64  `json:"id" gorm:"id"`
	Name        string `json:"name" gorm:"name"`
	Title       string `json:"title" gorm:"title"`
	Keywords    string `json:"keywords" gorm:"keywords"`
	Description string `json:"description" gorm:"description"`
	SiteId      int64  `json:"site_id" gorm:"site_id"`
	Status      int8   `json:"status" gorm:"status"` // 1 正常 5禁用  9删除
	CreateTime  int64  `json:"create_time" gorm:"create_time"`
	UpdateTime  int64  `json:"update_time" gorm:"update_time"`
	DeleteTime  int64  `json:"delete_time" gorm:"delete_time"`
}
