package user

import (
	"encoding/json"
	"gin/api/admin/user/v1"
	"gin/internal/global"
	"gin/internal/library/jwt"
	"gin/internal/library/vcrypto"
	"gin/internal/modules/admin/v1/logic/common"
	"gin/internal/modules/admin/v1/models"
	"gin/internal/modules/admin/v1/service"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gconv"
	"github.com/golang-module/carbon"
	"sort"
)

type (
	userLogic struct{}
)

func init() {
	service.RegisterUser(New())
}

func New() service.UserInterface {
	return &userLogic{}
}

func (a *userLogic) Items(param v1.ItemReq) (map[string]interface{}, error) {
	model := global.DB.Model(&models.User{})
	model.Select("id,vid,phone,username,realname,email,status,create_time,update_time,role_ids")
	model = common.WhereBySearch(model, param.Search)
	//identifyIdStr := strconv.FormatInt(param.IdentifyId, 10)
	model.Where("status != ? and identify_id=?", 9, param.IdentifyId)
	//model.Where(gorm.Expr("identify_ids like ? or identify_ids like ? or identify_ids like ? or identify_ids = ?", "["+identifyIdStr+",%", "%,"+identifyIdStr+"]", "%,"+identifyIdStr+",%", "["+identifyIdStr+"]"))

	var count int64
	model.Count(&count)
	if param.Limit != 0 {
		if param.PageIndex == 0 {
			param.PageIndex = 1
		}
		model.Offset((param.PageIndex - 1) * param.Limit).Limit(param.Limit)
	}
	if param.OrderBy != "" && param.OrderByType != "" {
		model.Order(param.OrderBy + " " + param.OrderByType)
	} else {
		model.Order("id desc")
	}
	var result []map[string]interface{}
	model.Find(&result)
	return map[string]interface{}{"items": result, "total": count}, nil
}

func (a *userLogic) Info(param v1.InfoReq) (map[string]interface{}, error) {
	result := map[string]interface{}{}
	global.DB.Model(&models.User{}).Where("status != ?", 9).First(&result, param.Id)
	return result, nil
}

func (a *userLogic) Create(param v1.CreateReq) (int64, error) {

	uid := service.TableIds().GetAddId("user_service", param.IdentifyId)

	//密码加密
	key := gconv.String(global.Cfg.Login.Key)
	pwd := vcrypto.HexEnCrypt(param.PassWord, key, vcrypto.DesCBCEncrypt)

	data := models.User{
		Phone:      param.Phone,
		Username:   param.UserName,
		Email:      param.Email,
		Realname:   param.RealName,
		Password:   pwd,
		Status:     param.Status,
		CreateTime: carbon.Now().Timestamp(),
		Uid:        uid,
	}

	result := global.DB.Model(&models.User{}).Create(&data)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func (a *userLogic) Update(param v1.UpdateReq) (int64, error) {
	//密码加密
	key := gconv.String(global.Cfg.Login.Key)
	pwd := vcrypto.HexEnCrypt(param.PassWord, key, vcrypto.DesCBCEncrypt)
	data := models.User{
		Phone:      param.Phone,
		Username:   param.UserName,
		Email:      param.Email,
		Realname:   param.RealName,
		Password:   pwd,
		Status:     param.Status,
		UpdateTime: carbon.Now().Timestamp(),
	}

	result := global.DB.Model(&models.User{}).Where("id = ? and status !=?", param.Id, 9).Updates(&data)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func (a *userLogic) Delete(param v1.DeleteReq) (int64, error) {
	result := global.DB.Model(&models.User{}).Where("id = ?", param.Id).Updates(map[string]interface{}{
		"status":      9,
		"delete_time": carbon.Now().Timestamp(),
	})
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func (a *userLogic) GetUserInfo(c *gin.Context) *v1.UserInfo {
	userInfo := new(v1.UserInfo)
	token := c.Request.Header.Get("token")
	tokenInfo, err := jwt.ParseJwtGoToken(token)
	if err != nil {
		return nil
	}

	err = json.Unmarshal([]byte(tokenInfo.Audience), &userInfo)
	if err != nil {
		return nil
	}
	return userInfo
}

func (a *userLogic) GetUserIdentify(c *gin.Context, userId int64) []models.Identify {
	if userId == 0 {
		userInfo := a.GetUserInfo(c)
		userId = userInfo.ID
	}

	var userRoleIdentifySlice []int64
	global.DB.Model(&models.UserRole{}).Where("user_id = ?", userId).Pluck("identify_id", &userRoleIdentifySlice)

	var userIdentifySlice []int64
	global.DB.Model(&models.UserPermission{}).Where("user_id = ?", userId).Pluck("identify_id", &userIdentifySlice)

	identifySlice := append(userRoleIdentifySlice, userIdentifySlice...)
	sort.Slice(identifySlice, func(i, j int) bool {
		return identifySlice[i] < identifySlice[j]
	})

	var identifyList []models.Identify
	global.DB.Model(&models.Identify{}).Select("id,identify_name,type,father_identify_id,identify_code,status,create_time,update_time").Where("status != 9 and id in ?", identifySlice).Find(&identifyList)
	return identifyList
}

func (a *userLogic) GetSecret(pwd string) string {
	key := gconv.String(global.Cfg.Login.Key)
	secretPwd := vcrypto.HexEnCrypt(pwd, key, vcrypto.DesCBCEncrypt)
	return secretPwd
}
