package models

type CmsPage struct {
	ID               int64  `json:"id" gorm:"id"`
	SiteId           int64  `json:"site_id" gorm:"site_id"`
	Status           int64  `json:"status" gorm:"status"` // 1正常 5禁用 9删除
	Subject          string `json:"subject" gorm:"subject"`
	Title            string `json:"title" gorm:"title"`
	Keywords         string `json:"keywords" gorm:"keywords"`
	Description      string `json:"description" gorm:"description"`
	Content          string `json:"content" gorm:"content"`
	URL              string `json:"url" gorm:"url"`             // 页面地址
	ImageUrl         string `json:"image_url" gorm:"image_url"` // 页面封面图片地址
	TemplateId       int64  `json:"template_id" gorm:"template_id"`
	ClassifyId       int64  `json:"classify_id" gorm:"classify_id"`
	AuthorId         int64  `json:"author_id" gorm:"author_id"`
	ProductId        int64  `json:"product_id" gorm:"product_id"`
	TagIds           string `json:"tag_ids" gorm:"tag_ids"` // Tags
	LastUpdateUserId int64  `json:"last_update_user_id" gorm:"last_update_user_id"`
	CreateTime       int64  `json:"create_time" gorm:"create_time"`
	UpdateTime       int64  `json:"update_time" gorm:"update_time"`
	DeleteTime       int64  `json:"delete_time" gorm:"delete_time"`
	FirstMakeTime    int64  `json:"first_make_time" gorm:"first_make_time"`
	IsPublish        int64  `json:"is_publish" gorm:"is_publish"`   // 1 发布过 2没发布过
	StarNumber       string `json:"star_number" gorm:"star_number"` // 点赞数
	CreateUserId     int64  `json:"create_user_id" gorm:"create_user_id"`
	CreateUserName   string `json:"create_user_name" gorm:"create_user_name"`
}
