package models

// User  用户表
type User struct {
	ID         int64  `gorm:"column:id" json:"id"`                   //  管理员id
	Uid        int64  `gorm:"column:uid" json:"uid"`                 //  序号
	Username   string `gorm:"column:username" json:"username"`       //  管理员名称
	Phone      string `gorm:"column:phone" json:"phone"`             //  手机号
	Realname   string `gorm:"column:realname" json:"realname"`       //  真实中文名
	Password   string `gorm:"column:password" json:"password"`       //  账号密码
	Email      string `gorm:"column:email" json:"email"`             //  管理员邮箱
	Status     int8   `gorm:"column:status" json:"status"`           //  管理员账户状态，1：正常 5 禁用  9删除
	CreateTime int64  `gorm:"column:create_time" json:"create_time"` //  用户创建时间
	UpdateTime int64  `gorm:"column:update_time" json:"update_time"` //  修改时间
	DeleteTime int64  `gorm:"column:delete_time" json:"delete_time"` //  删除时间
	UserIp     string `gorm:"column:user_ip" json:"user_ip"`         //  用户ip地址
	LoginTime  int64  `gorm:"column:login_time" json:"login_time"`   //  最近一次登录时间
}
