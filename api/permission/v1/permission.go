package v1

type ItemReq struct {
	Limit       int    `form:"limit" json:"limit"`
	PageIndex   int    `form:"pageIndex" json:"pageIndex"`
	OrderBy     string `form:"orderBy" json:"orderBy"`
	OrderByType string `form:"orderByType" json:"orderByType"`
	Search      string `form:"search" json:"search"`
	IdentifyId  int64  `form:"identify_id" json:"identify_id" binding:"required"`
}

type CreateReq struct {
	PermissionName       string `form:"permission_name" json:"permission_name" binding:"required"`
	PermissionCode       string `form:"permission_code" json:"permission_code" binding:"required"`
	IdentifyId           int64  `form:"identify_id" json:"identify_id" binding:"required"`
	Type                 int8   `form:"type" json:"type" binding:"required"`
	FatherPermissionCode string `form:"father_permission_code" json:"father_permission_code"`
	Status               int8   `form:"status" json:"status" binding:"required"`
}

type UpdateReq struct {
	Id                   int64  `form:"id" json:"id" binding:"required"`
	PermissionName       string `form:"permission_name" json:"permission_name" binding:"required"`
	PermissionCode       string `form:"permission_code" json:"permission_code" binding:"required"`
	IdentifyId           int64  `form:"identify_id" json:"identify_id" binding:"required"`
	Type                 int8   `form:"type" json:"type" binding:"required"`
	FatherPermissionCode string `form:"father_permission_code" json:"father_permission_code"`
	Status               int8   `form:"status" json:"status" binding:"required"`
}

type DeleteReq struct {
	Id string `form:"id" json:"id" binding:"required"`
}
