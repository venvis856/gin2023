package v1

type UserAddPermissionReq struct {
	UserId         int64    `form:"user_id" json:"user_id" binding:"required"`
	PermissionCode []string `form:"permission_code" json:"permission_code" binding:"required"`
	IdentifyId     int64    `form:"identify_id" json:"identify_id" binding:"required"`
}

type RoleAddPermissionReq struct {
	RoleId         int64    `form:"role_id" json:"role_id" binding:"required"`
	PermissionCode []string `form:"permission_code" json:"permission_code" binding:"required"`
	IdentifyId     int64    `form:"identify_id" json:"identify_id" binding:"required"`
}

type GetPermissionByUserReq struct {
	UserId     int64 `form:"user_id" json:"user_id" binding:"required"`
	IdentifyId int64 `form:"identify_id" json:"identify_id" binding:"required"`
}

type GetAllPermissionByRoleReq struct {
	RoleId     int64 `form:"role_id" json:"role_id" binding:"required"`
	IdentifyId int64 `form:"identify_id" json:"identify_id" binding:"required"`
}

type GetAllPermissionByUserReq struct {
	IdentifyId int64 `form:"identify_id" json:"identify_id" binding:"required"`
}

type GetMenuByUserReq struct {
	IdentifyId int64 `form:"identify_id" json:"identify_id" binding:"required"`
}
