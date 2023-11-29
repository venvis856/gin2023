package models

type Video struct {
	ID          int64  `json:"id" gorm:"id"`
	Title       string `json:"title" gorm:"title"`
	Subject     string `json:"subject" gorm:"subject"`
	URL         string `json:"url" gorm:"url"`
	Thumbnail   string `json:"thumbnail" gorm:"thumbnail"`
	Description string `json:"description" gorm:"description"`
	SiteId      int64  `json:"site_id" gorm:"site_id"`
	Tags        string `json:"tags" gorm:"tags"`
	TagIds      string `json:"tag_ids" gorm:"tag_ids"`
	CreateTime  int64  `json:"create_time" gorm:"create_time"`
	UpdateTime  int64  `json:"update_time" gorm:"update_time"`
	DeleteTime  int64  `json:"delete_time" gorm:"delete_time"`
	Status      int8   `json:"status" gorm:"status"`
	IsPublic    int8   `json:"is_public" gorm:"is_public"` // 1 正常 2 已发布
}
