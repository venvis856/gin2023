package models

type Rrweb struct {
	Id        int    `gorm:"column:id" db:"id" json:"id" form:"id"`                                 //数据库ID
	Uid       string `gorm:"column:uid" db:"uid" json:"uid" form:"uid"`                             //uid
	Msg       string `gorm:"column:msg" db:"msg" json:"msg" form:"msg"`                             //msg
	StartTime int    `gorm:"column:start_time" db:"start_time" json:"start_time" form:"start_time"` //开始时间
	EndTime   int    `gorm:"column:end_time" db:"end_time" json:"end_time" form:"end_time"`         //结束时间
}
