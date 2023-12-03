package v2

type LoginReq struct {
	Phone        string `form:"phone" json:"phone" binding:"required_without_all=Email UserName"`
	Email        string `form:"email" json:"email" binding:"required_without_all=Phone UserName"`
	UserName     string `form:"username" json:"username" binding:"required_without_all=Email Phone"`
	PassWord     string `form:"password" json:"password"  binding:"required"`
}