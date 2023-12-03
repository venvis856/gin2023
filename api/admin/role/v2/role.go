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
	RoleName   string `form:"role_name" json:"role_name" binding:"required"`
	Type       int8   `form:"type" json:"type" binding:"required"`
	Status     int8   `form:"status" json:"status" binding:"required"`
	IdentifyId int64  `form:"identify_id" json:"identify_id" binding:"required"`
}

type UpdateReq struct {
	Id         int    `form:"id" json:"id" binding:"required"`
	RoleName   string `form:"role_name" json:"role_name" binding:"required"`
	Type       int8   `form:"type" json:"type" binding:"required"`
	IdentifyId int64  `form:"identify_id" json:"identify_id" binding:"required"`
	Status     int8   `form:"status" json:"status" binding:"required"`
}

type DeleteReq struct {
	Id int `form:"id" json:"id" binding:"required"`
}
