package v2

type ItemReq struct {
	Limit       int    `form:"limit" json:"limit"`
	PageIndex   int    `form:"pageIndex" json:"pageIndex"`
	OrderBy     string `form:"orderBy" json:"orderBy"`
	OrderByType string `form:"orderByType" json:"orderByType"`
	Search      string `form:"search" json:"search"`
	IdentifyId  int64  `form:"identify_id" json:"identify_id" binding:"required"`
}

type InfoReq struct {
	Id int `form:"id" json:"id" binding:"required"`
}

type CreateReq struct {
	Phone      string  `form:"phone" json:"phone" binding:"required_without_all=Email UserName"`
	Email      string  `form:"email" json:"email" binding:"required_without_all=Phone UserName""`
	UserName   string  `form:"username" json:"username" binding:"required_without_all=Email Phone""`
	PassWord   string  `form:"password" json:"password" binding:"required"`
	Status     int8    `form:"status" json:"status" binding:"required"`
	RealName   string  `form:"realname" json:"realname"`
	RoleIds    []int64 `form:"role_ids" json:"role_ids" binding:"required"`
	IdentifyId int64   `form:"identify_id" json:"identify_id" binding:"required"`
}

type UpdateReq struct {
	Id         int     `form:"id" json:"id" binding:"required"`
	Phone      string  `form:"phone" json:"phone" binding:"required_without_all=Email UserName"`
	Email      string  `form:"email" json:"email" binding:"required_without_all=Phone UserName""`
	UserName   string  `form:"username" json:"username" binding:"required_without_all=Email Phone""`
	PassWord   string  `form:"password" json:"password" binding:"required"`
	Status     int8    `form:"status" json:"status" binding:"required"`
	RealName   string  `form:"realname" json:"realname"`
	RoleIds    []int64 `form:"role_ids" json:"role_ids" binding:"required"`
	IdentifyId int64   `form:"identify_id" json:"identify_id" binding:"required"`
}

type DeleteReq struct {
	Id int `form:"id" json:"id" binding:"required"`
}

type UserInfo struct {
	CreateTime   int64  `json:"create_time"`
	Email        string `json:"email"`
	ID           int64  `json:"id"`
	IdentifyId   int64  `json:"identify_id"`
	IdentifyName string `json:"identify_name"`
	IdentifyType int    `json:"identify_type"`
	Phone        string `json:"phone"`
	Realname     string `json:"realname"`
	Username     string `json:"username"`
	Threshold    int    `json:"threshold"`
	Vid          int    `json:"vid"`

}

type SecretReq struct {
	Pwd string `form:"pwd" json:"pwd" binding:"required"`
}
