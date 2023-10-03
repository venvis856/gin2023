package models

type User struct {
	ID         int64  `json:"id" gorm:"id"` // 管理员ID
	Vid        int64  `json:"vid" gorm:"vid"`
	Username   string `json:"username" gorm:"username"`       // 管理员名称
	Phone      string `json:"phone" gorm:"phone"`             // 手机号
	Realname   string `json:"realname" gorm:"realname"`       // 真实中文名
	Password   string `json:"-" gorm:"password"`              // 账号密码
	Email      string `json:"email" gorm:"email"`             // 管理员邮箱
	Status     int8   `json:"status" gorm:"status"`           // 管理员账户状态，0：停用 1：正常
	CreateTime int64  `json:"create_time" gorm:"create_time"` // 用户创建时间
	UpdateTime int64  `json:"update_time" gorm:"update_time"` // 修改时间
	DeleteTime int64  `json:"delete_time" gorm:"delete_time"`
	UserIp     string `json:"user_ip" gorm:"user_ip"`         // 用户IP地址
	LoginTime  int64  `json:"login_time" gorm:"login_time"`   // 最近一次登录时间
	RoleIds    string `json:"role_ids" gorm:"role_ids"`       // 角色ids
	IdentifyId int64  `json:"identify_id" gorm:"identify_id"` // 标识id

	Identify Identify `json:"-" gorm:"foreignKey:identify_id;references:id"`
}
