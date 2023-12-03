package models

// UserRole  角色表
type UserRole struct {
	ID          int64 `gorm:"column:id" json:"id"`                     //  id
	IdentifyId  int64 `gorm:"column:identify_id" json:"identify_id"`   //  标识id
	UserId      int64 `gorm:"column:user_id" json:"user_id"`           //  user_service id
	RoleId      int64 `gorm:"column:role_id" json:"role_id"`           //  role_service id
	IsEffective int64 `gorm:"column:is_effective" json:"is_effective"` //  是否有效
}
