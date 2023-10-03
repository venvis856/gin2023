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
	Root             string `form:"root" json:"root" binding:"required"`
	Type             int8   `form:"type" json:"type" binding:"required"`
	FatherIdentifyId int64  `form:"father_identify_id" json:"father_identify_id"`
	Status           int    `form:"status" json:"status" binding:"required"`
	Location         string `form:"location" json:"location" binding:"required"`
	LocationX        string `form:"location_x" json:"location_x" binding:"required"`
	LocationY        string `form:"location_y" json:"location_y" binding:"required"`
	ProvinceCode     string `form:"province_code" json:"province_code"`            // 省编码
	ProvinceName     string `form:"province_name" json:"province_name"`            // 省
	CityCode         string `form:"city_code" json:"city_code" binding:"required"` // 市编码
	CityName         string `form:"city_name" json:"city_name" binding:"required"` // 市
	AreaCode         string `form:"area_code" json:"area_code" binding:"required"` // 县区编码
	AreaName         string `form:"area_name" json:"area_name" binding:"required"` // 县区

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
	Root             string `form:"root" json:"root" binding:"required"`
	Type             int8   `form:"type" json:"type" binding:"required"`
	FatherIdentifyId int64  `form:"father_identify_id" json:"father_identify_id"`
	Status           int    `form:"status" json:"status" binding:"required"`
	Location         string `form:"location" json:"location" binding:"required"`
	LocationX        string `form:"location_x" json:"location_x" binding:"required"`
	LocationY        string `form:"location_y" json:"location_y" binding:"required"`
	ProvinceCode     string `form:"province_code" json:"province_code"`            // 省编码
	ProvinceName     string `form:"province_name" json:"province_name"`            // 省
	CityCode         string `form:"city_code" json:"city_code" binding:"required"` // 市编码
	CityName         string `form:"city_name" json:"city_name" binding:"required"` // 市
	AreaCode         string `form:"area_code" json:"area_code" binding:"required"` // 县区编码
	AreaName         string `form:"area_name" json:"area_name" binding:"required"` // 县区

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
