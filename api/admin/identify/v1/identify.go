package v1

type ItemReq struct {
	Limit       int    `form:"limit" json:"limit"`
	PageIndex   int    `form:"pageIndex" json:"pageIndex"`
	OrderBy     string `form:"orderBy" json:"orderBy"`
	OrderByType string `form:"orderByType" json:"orderByType"`
	Search      string `form:"search" json:"search"`
}

type InfoReq struct {
	Id int `form:"id" json:"id" binding:"required"`
}

type CreateReq struct {
	IdentifyName     string `form:"identify_name" json:"identify_name" binding:"required"`
	IdentifyCode     string `form:"identify_code" json:"identify_code" binding:"required"`
	Type             int8   `form:"type" json:"type" binding:"required"`
	FatherIdentifyId int64  `form:"father_identify_id" json:"father_identify_id"`
	Status           int8   `form:"status" json:"status" binding:"required"`

	Phone      string `form:"phone" json:"phone" binding:"required_without_all=Email UserName"`
	Email      string `form:"email" json:"email" binding:"required_without_all=Phone UserName"`
	UserName   string `form:"username" json:"username" binding:"required_without_all=Email Phone"`
	PassWord   string `form:"password" json:"password" binding:"required"`
	UserStatus int8   `form:"user_status" json:"user_status" binding:"required"`
	RealName   string `form:"realname" json:"realname"`
}

type UpdateReq struct {
	Id               int    `form:"id" json:"id" binding:"required"`
	IdentifyName     string `form:"identify_name" json:"identify_name" binding:"required"`
	IdentifyCode     string `form:"identify_code" json:"identify_code" binding:"required"`
	Type             int8   `form:"type" json:"type" binding:"required"`
	FatherIdentifyId int64  `form:"father_identify_id" json:"father_identify_id"`
	Status           int8   `form:"status" json:"status" binding:"required"`

	UserId     int64  `form:"user_id" json:"user_id" binding:"required"`
	Phone      string `form:"phone" json:"phone" binding:"required_without_all=Email UserName"`
	Email      string `form:"email" json:"email" binding:"required_without_all=Phone UserName"`
	UserName   string `form:"username" json:"username" binding:"required_without_all=Email Phone"`
	PassWord   string `form:"password" json:"password" binding:"required"`
	UserStatus int8   `form:"user_status" json:"user_status" binding:"required"`
	RealName   string `form:"realname" json:"realname"`
}

type DeleteReq struct {
	Id int `form:"id" json:"id" binding:"required"`
}
