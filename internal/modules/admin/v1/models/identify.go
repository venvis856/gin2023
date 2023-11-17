package models

// Identify undefined
var (
	IDENTIFY_TYPE_SYSTEM = 1
)

// Identify  身份标识表
type Identify struct {
	ID               int64  `gorm:"column:id" json:"id"`                                 //  id
	IdentifyName     string `gorm:"column:identify_name" json:"identify_name"`           //  身份名
	IdentifyCode     string `gorm:"column:identify_code" json:"identify_code"`           //  身份标识符
	Type             int8   `gorm:"column:type" json:"type"`                             //  1 系统 2 其他
	FatherIdentifyId int64  `gorm:"column:father_identify_id" json:"father_identify_id"` //  父级
	Status           int8   `gorm:"column:status" json:"status"`                         //  1正常 5禁用 9删除
	CreateTime       int64  `gorm:"column:create_time" json:"create_time"`               //  创建时间
	UpdateTime       int64  `gorm:"column:update_time" json:"update_time"`               //  更新时间
	DeleteTime       int64  `gorm:"column:delete_time" json:"delete_time"`               //  删除时间
}
