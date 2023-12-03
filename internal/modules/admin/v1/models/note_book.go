
package models

type Notebook struct {
	ID int64 `json:"id" gorm:"id"`
	Date string `json:"date" gorm:"date"`
	Content string `json:"content" gorm:"content"`
	UserId int64 `json:"user_id" gorm:"user_id"`
	Status int64 `json:"status" gorm:"status"` // 1 待通知 2 已通知 5 禁用 9 删除
	CreateTime int64 `json:"create_time" gorm:"create_time"`
	UpdateTime int64 `json:"update_time" gorm:"update_time"`
	DeleteTime int64 `json:"delete_time" gorm:"delete_time"`
}
