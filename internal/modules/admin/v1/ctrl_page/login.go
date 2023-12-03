package ctrl_page

import (
	"fmt"
	"gin/internal/global"
	"gin/internal/library/jwt"
	"gin/internal/library/vcrypto"
	"gin/internal/modules/admin/v1/models"

	"gin/internal/global/errcode"
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
		global.Response.Error(c, errcode.ERROR_PARAMS,err.Error())
		return
	}
	//密码加密
	key := gconv.String(global.Cfg.Login.Key)
	pwd := vcrypto.HexEnCrypt(param.PassWord, key, vcrypto.DesCBCEncrypt)

	result := map[string]interface{}{}
	global.DB.Model(&models.User{}).Where("username = ? and password=? and status=?", param.UserName, pwd, 1).First(&result)

	if len(result) == 0 {
		global.Response.Error(c,errcode.ERROR_SERVER, "用户或密码错误")
		return
	}

	token, err := jwt.CreateJwtGoToken(param.UserName, gconv.String(result["id"]))
	if err != nil {
		global.Response.Error(c, errcode.ERROR_SERVER, fmt.Sprintf("登录失败:%v",err.Error()))
		return
	}
	global.Response.Success(c, token)
}
