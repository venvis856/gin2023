package models

// CmsTemplate undefined
type CmsTemplate struct {
	ID           int64  `json:"id" gorm:"id"`
	SiteId       int64  `json:"site_id" gorm:"site_id"`
	Status       int64  `json:"status" gorm:"status"`
	TemplateName string `json:"template_name" gorm:"template_name"`
	Type         int64  `json:"type" gorm:"type"` // 1首页  2文章  3分类  4产品  5review  6guide 8 video
	Content      string `json:"content" gorm:"content"`
	ModuleIds    string `json:"module_ids" gorm:"module_ids"`
	CreateTime   int64  `json:"create_time" gorm:"create_time"`
	UpdateTime   int64  `json:"update_time" gorm:"update_time"`
	DeleteTime   int64  `json:"delete_time" gorm:"delete_time"`
}
