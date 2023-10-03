package models

// Identify undefined
var (
	IDENTIFY_TYPE_HOTEL  = 1
	IDENTIFY_TYPE_POLICE = 2
	IDENTIFY_TYPE_BUILD  = 3
	IDENTIFY_TYPE_PARK   = 4
	IDENTIFY_TYPE_SYSTEM = 9
)

type Identify struct {
	ID               int64  `json:"id" gorm:"id"`                                 // ID
	IdentifyName     string `json:"identify_name" gorm:"identify_name"`           // 身份名
	Root             string `json:"root" gorm:"root"`                             // 身份标识符
	Type             int8   `json:"type" gorm:"type"`                             // 1 酒店 2 派出所 3 大厦 4 园区  9 系统
	FatherIdentifyId int64  `json:"father_identify_id" gorm:"father_identify_id"` // 父级
	Status           int64  `json:"status" gorm:"status"`                         // 1正常 5禁用 9删除
	CreateTime       int64  `json:"create_time" gorm:"create_time"`               // 创建时间
	UpdateTime       int64  `json:"update_time" gorm:"update_time"`               // 更新时间
	DeleteTime       int64  `json:"delete_time" gorm:"delete_time"`               // 删除时间
	ProvinceCode     string `json:"province_code" gorm:"province_code"`           // 省编码
	ProvinceName     string `json:"province_name" gorm:"province_name"`           // 省
	CityCode         string `json:"city_code" gorm:"city_code"`                   // 市编码
	CityName         string `json:"city_name" gorm:"city_name"`                   // 市
	AreaCode         string `json:"area_code" gorm:"area_code"`                   // 县区编码
	AreaName         string `json:"area_name" gorm:"area_name"`                   // 县区
	Location         string `json:"location" gorm:"location"`                     // 位置
	LocationX        string `json:"location_x" gorm:"location_x"`                 // 经度
	LocationY        string `json:"location_y" gorm:"location_y"`                 // 维度

}
