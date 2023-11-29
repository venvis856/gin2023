package admin

import (
	"gin/app/library/jwt"
	"gin/app/library/vcrypto"
	"gin/app/models"
	"gin/global"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gconv"
)

type LoginController struct{}

func (*LoginController) Login(c *gin.Context) {
	var param struct {
		UserName string `form:"username" json:"username" binding:"required"`
		PassWord string `form:"password" json:"password"  binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
		return
	}
	//密码加密
	key := gconv.String(global.Config.Get("login.key"))
	pwd := vcrypto.HexEnCrypt(param.PassWord, key, vcrypto.DesCBCEncrypt)

	result := map[string]interface{}{}
	global.DB.Model(&models.User{}).Where("username = ? and password=? and status=?", param.UserName, pwd, 1).First(&result)

	if len(result) == 0 {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, "用户或密码错误", "")
		return
	}

	token, err := jwt.CreateJwtGoToken(param.UserName, gconv.String(result["id"]))
	if err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "登录失败")
		return
	}
	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", token)
}
