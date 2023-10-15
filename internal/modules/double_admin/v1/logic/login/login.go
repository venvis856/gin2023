package login

import (
	"encoding/json"
	"errors"
	"gin/api/double_admin/v1/login/v1"
	"gin/internal/config"
	"gin/internal/global"
	"gin/internal/global/errcode"
	"gin/internal/library/jwt"
	"gin/internal/library/vcrypto"
	"gin/internal/modules/double_admin/v1/logic/permission_operate"
	"gin/internal/modules/double_admin/v1/models"
	"gin/internal/modules/double_admin/v1/service"
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
	//密码加密
	key := gconv.String(config.Cfg.Login.Key)
	pwd := vcrypto.HexEnCrypt(param.PassWord, key, vcrypto.DesCBCEncrypt)

	model := global.DB.Model(&models.User{})
	model.Joins("left join identify on user.identify_id =identify.id ")
	model.Select("user.id as user_id, user.vid,user.username,user.phone,user.realname,user.email,user.create_time,user.identify_id,identify.identify_name,identify.type as identify_type")
	model.Where("user.status != 9 and user.password = ?", pwd)
	model.Where("identify.type =?", param.IdentifyType)

	if param.Phone != "" {
		model.Where("user.phone=?", param.Phone)
	} else if param.Email != "" {
		model.Where("user.email=?", param.Email)
	} else if param.UserName != "" {
		model.Where("user.username=?", param.UserName)
	}

	result := service.User().GetUserInfo(c)
	model.Find(&result)
	if result.UserId == 0 {
		return map[string]interface{}{}, errors.New("账号或密码错误")
	}
	result.Scene = param.Scene
	// 校验权限
	permissionSlice := permission_operate.GetAllPermissionByUser(result.UserId, result.IdentifyId)
	authBool := false
	for _, v := range permissionSlice {
		if gconv.Uint8(v["scene"]) == param.Scene {
			authBool = true
			break
		}
	}

	if !authBool {
		return map[string]interface{}{}, errcode.TOKEN_FORBIDDEN
	}

	jsonResult, err := json.Marshal(result)
	if err != nil {
		return map[string]interface{}{}, err
	}
	token, err := jwt.CreateJwtGoToken(string(jsonResult), gconv.String(result.UserId))
	if err != nil {
		return map[string]interface{}{}, err
	}

	rs := map[string]interface{}{
		"userInfo": result,
		"token":    token,
	}

	go a.LoginLog(c, result.UserId)
	return rs, nil
}

func (a *loginLogic) LoginLog(c *gin.Context, userId int64) {
	data := models.User{
		UserIp:    c.ClientIP(),
		LoginTime: carbon.Now().Timestamp(),
	}
	global.DB.Model(&models.User{}).Where("id = ?", userId).Updates(&data)
}

func (a *loginLogic) UserInfo(c *gin.Context) (v1.LoginInfo, error) {
	userInfo := service.User().GetUserInfo(c)
	if userInfo.UserId == 0 {
		return v1.LoginInfo{}, errcode.TOKEN_FAIL
	}
	retUser := struct {
		*models.User
		RoleId       int64  `json:"roleId"`
		Role         string `json:"role"`
		StreetName   string `json:"street_name"`
		IdentifyName string `json:"identify_name"`
	}{}
	err := global.DB.Model(models.User{}).Preload("Identify").Find(&retUser.User, userInfo.UserId).Error
	if err != nil {
		return v1.LoginInfo{}, err
	}
	rols := models.Role{}
	var roles []int64
	err = json.Unmarshal([]byte(retUser.RoleIds), &roles)
	if err != nil {
		return v1.LoginInfo{}, err
	}
	if len(roles) == 0 {
		return retUser, nil
	}
	err = global.DB.Model(models.Role{}).Where("id in ?", roles).Take(&rols).Error
	if err != nil {
		return v1.LoginInfo{}, err
	}
	retUser.RoleId = rols.ID
	retUser.Role = rols.RoleName
	retUser.StreetName = retUser.User.Identify.Location
	retUser.IdentifyName = retUser.User.Identify.IdentifyName
	return retUser, nil
}
