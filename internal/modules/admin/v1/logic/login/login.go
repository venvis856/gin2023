package login

import (
	"encoding/json"
	"errors"
	"gin/api/admin/login/v1"
	"gin/internal/common_config"
	"gin/internal/global"
	"gin/internal/library/jwt"
	"gin/internal/library/vcrypto"
	"gin/internal/modules/admin/v1/logic/permission_operate"
	"gin/internal/modules/admin/v1/models"
	"gin/internal/modules/admin/v1/service"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gconv"
	"github.com/golang-module/carbon"
)

type (
	loginLogic struct{}
)

func init() {
	service.RegisterLogin(New())
}

func New() service.LoginInterface {
	return &loginLogic{}
}

func (a *loginLogic) Login(c *gin.Context, param v1.LoginReq) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	//密码加密
	key := gconv.String(common_config.Cfg.Login.Key)
	pwd := vcrypto.HexEnCrypt(param.PassWord, key, vcrypto.DesCBCEncrypt)

	var userInfo models.User
	model := global.DB.Model(&models.User{})
	model.Select("user.id as user_id, user.uid,user.username,user.phone,user.realname,user.email,user.create_time")
	model.Where("user.status != 9 and user.password = ?", pwd)

	if param.Phone != "" {
		model.Where("user.phone=?", param.Phone)
	} else if param.Email != "" {
		model.Where("user.email=?", param.Email)
	} else if param.UserName != "" {
		model.Where("user.username=?", param.UserName)
	}
	model.First(&userInfo)

	if userInfo.ID == 0 {
		return result, errors.New("账号或密码错误")
	}

	identifyId, _ := c.Get("identify_id")
	// 校验权限
	permissionSlice := permission_operate.GetAllPermissionByUser(userInfo.ID, identifyId.(int64))

	if len(permissionSlice) == 0 {
		return map[string]interface{}{}, errors.New("没有权限无效")
	}

	jsonResult, err := json.Marshal(userInfo)
	if err != nil {
		return map[string]interface{}{}, err
	}
	token, err := jwt.CreateJwtGoToken(string(jsonResult), gconv.String(userInfo.ID))
	if err != nil {
		return map[string]interface{}{}, err
	}

	result = map[string]interface{}{
		"userInfo": result,
		"token":    token,
	}

	go a.LoginLog(c, userInfo.ID)
	return result, nil
}

func (a *loginLogic) LoginLog(c *gin.Context, userId int64) {
	data := models.User{
		UserIp:    c.ClientIP(),
		LoginTime: carbon.Now().Timestamp(),
	}
	global.DB.Model(&models.User{}).Where("id = ?", userId).Updates(&data)
}