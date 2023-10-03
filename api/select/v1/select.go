package v1

type GetRoleSelectListReq struct {
	IdentifyId int64 `form:"identify_id" json:"identify_id" binding:"required"`
}

type GetSelectStreetRoadReq struct {
	IdentifyId int64 `form:"identify_id" json:"identify_id" binding:"required"`
}

type GetUserSelectByIdentifyReq struct {
	IdentifyId int64 `form:"identify_id" json:"identify_id" binding:"required"`
}
