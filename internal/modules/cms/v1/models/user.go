package models

type User struct {
	Id         int    `gorm:"column:id" db:"id" json:"id" form:"id"`                                     //管理员ID
	Username   string `gorm:"column:username" db:"username" json:"username" form:"username"`             //管理员名称
	Realname   string `gorm:"column:realname" db:"realname" json:"realname" form:"realname"`             //真实中文名
	Password   string `gorm:"column:password" db:"password" json:"password" form:"password"`             //账号密码
	Email      string `gorm:"column:email" db:"email" json:"email" form:"email"`                         //管理员邮箱
	Status     int8   `gorm:"column:status" db:"status" json:"status" form:"status"`                     //管理员账户状态，0：停用 1：正常
	CreateTime int    `gorm:"column:create_time" db:"create_time" json:"create_time" form:"create_time"` //用户创建时间
	UpdateTime int    `gorm:"column:update_time" db:"update_time" json:"update_time" form:"update_time"` //修改时间
	DeleteTime int    `gorm:"column:delete_time" db:"delete_time" json:"delete_time" form:"delete_time"`
	UserIp     string `gorm:"column:user_ip" db:"user_ip" json:"user_ip" form:"user_ip"`             //用户IP地址
	LoginTime  int    `gorm:"column:login_time" db:"login_time" json:"login_time" form:"login_time"` //最近一次登录时间
	RoleIds    string `gorm:"column:role_ids" db:"role_ids" json:"role_ids" form:"role_ids"`         //角色ids
	SiteIds    string `gorm:"column:site_ids" db:"site_ids" json:"site_ids" form:"site_ids"`         //站点ids
}
