package v1

import "gin/internal/modules/admin/v1/models"

type LoginReq struct {
	Phone        string `form:"phone" json:"phone" binding:"required_without_all=Email UserName"`
	Email        string `form:"email" json:"email" binding:"required_without_all=Phone UserName""`
	UserName     string `form:"username" json:"username" binding:"required_without_all=Email Phone""`
	PassWord     string `form:"password" json:"password"  binding:"required"`
	IdentifyType int64  `form:"identify_type" json:"identify_type" binding:"required"`
	Scene        uint8  `form:"scene" json:"scene" binding:"required"`
}

type LoginInfo struct {
	*models.User
	RoleId       int64  `json:"roleId"`
	Role         string `json:"role"`
	StreetName   string `json:"street_name"`
	IdentifyName string `json:"identify_name"`
}
