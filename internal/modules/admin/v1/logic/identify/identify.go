package identify

import (
	"encoding/json"
	"errors"
	"gin/api/admin/identify/v1"
	"gin/internal/config"
	"gin/internal/global"
	"gin/internal/library/vcrypto"
	"gin/internal/modules/admin/v1/logic/common"
	"gin/internal/modules/admin/v1/models"
	"gin/internal/modules/admin/v1/service"
	"github.com/gogf/gf/util/gconv"
	"github.com/golang-module/carbon"
	"gorm.io/gorm"
)

type (
	identifyLogic struct {
		handlers map[string]handleObjInterface
	}
)

func init() {
	service.RegisterIdentify(New())
}

type handleObjInterface interface {
	InitPermission(identifyId int64, roleId int64, tx *gorm.DB) error
}

const (
	NORMAL = "normal"
	SYSTEM = "system"
)

func New() service.IdentifyInterface {
	return &identifyLogic{
		handlers: map[string]handleObjInterface{
			NORMAL: &Normal{},
			SYSTEM: &System{},
		},
	}
}

func (a *identifyLogic) Items(param v1.ItemReq) map[string]interface{} {
	model := global.DB.Model(&models.Identify{})
	model = common.WhereBySearch(model, param.Search)
	model.Where("status != ?", 9)
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
	return map[string]interface{}{"items": result, "total": count}
}

func (a *identifyLogic) Info(param v1.InfoReq) (map[string]interface{}, error) {
	identifyInfo := map[string]interface{}{}
	global.DB.Model(&models.Identify{}).Where("status != ?", 9).First(&identifyInfo, param.Id)
	// 先找到系统管理员角色
	roleInfo := make(map[string]interface{})
	global.DB.Model(&models.Role{}).Where("status != 9 and type=1 and identify_id=?", param.Id).First(&roleInfo)
	if len(roleInfo) == 0 {
		return map[string]interface{}{}, errors.New("该公司系统管理员异常")
	}
	userInfo := make(map[string]interface{})
	global.DB.Model(&models.User{}).Where("status != 9").
		Where("role_ids like ? or role_ids like ? or role_ids like ?  or role_ids like ?", "["+gconv.String(roleInfo["id"])+"]", "["+gconv.String(roleInfo["id"])+",", ","+gconv.String(roleInfo["id"])+",", ","+gconv.String(roleInfo["id"])+"]").
		First(&userInfo)
	result := make(map[string]interface{})
	result["identify"] = identifyInfo
	result["userinfo"] = userInfo
	return result, nil
}

func (a *identifyLogic) Create(param v1.CreateReq) (int64, error) {
	tx := global.DB.Begin()
	model := tx.Model(&models.Identify{})

	identifyData := models.Identify{
		IdentifyName:     param.IdentifyName,
		Status:           int64(param.Status),
		Root:             param.Root,
		Type:             param.Type,
		FatherIdentifyId: param.FatherIdentifyId,
		CreateTime:       carbon.Now().Timestamp(),
		ProvinceCode:     param.ProvinceCode,
		ProvinceName:     param.ProvinceName,
		CityCode:         param.CityCode,
		CityName:         param.CityName,
		AreaCode:         param.AreaCode,
		AreaName:         param.AreaName,
		Location:         param.Location,
		LocationX:        param.LocationX,
		LocationY:        param.LocationY,
	}
	result := model.Create(&identifyData)
	if result.Error != nil {
		tx.Rollback()
		return 0, result.Error
	}

	// 新增角色
	roleVid := service.TableIds().GetAddId("role", identifyData.ID)
	roleModel := tx.Model(&models.Role{})
	roleData := models.Role{
		RoleName:   "系统管理员",
		Status:     1,
		Type:       1,
		IdentifyId: identifyData.ID,
		CreateTime: carbon.Now().Timestamp(),
		Vid:        roleVid,
	}
	result = roleModel.Create(&roleData)
	if result.Error != nil {
		tx.Rollback()
		return 0, result.Error
	}

	// 新增管理员
	userModel := tx.Model(&models.User{})
	vid := service.TableIds().GetAddId("user", identifyData.ID)
	//密码加密
	key := gconv.String(config.Cfg.Login.Key)
	pwd := vcrypto.HexEnCrypt(param.PassWord, key, vcrypto.DesCBCEncrypt)
	roleIdsJson, _ := json.Marshal([]int64{roleData.ID})
	roleIds := gconv.String(roleIdsJson)

	userData := models.User{
		Phone:      param.Phone,
		Username:   param.UserName,
		Email:      param.Email,
		Realname:   param.RealName,
		Password:   pwd,
		Status:     param.UserStatus,
		RoleIds:    roleIds,
		CreateTime: carbon.Now().Timestamp(),
		IdentifyId: identifyData.ID,
		Vid:        vid,
	}
	result = userModel.Create(&userData)
	if result.Error != nil {
		tx.Rollback()
		return 0, result.Error
	}

	err := service.Identify().InitIdentifyPermission(param.Type, identifyData.ID, roleData.ID, tx)
	if err != nil {
		return 0, err
	}

	tx.Commit()

	return result.RowsAffected, nil
}

func (a *identifyLogic) Update(param v1.UpdateReq) (int64, error) {
	tx := global.DB.Begin()
	identifyModel := tx.Model(&models.Identify{})

	//data := models.Identify{
	//	IdentifyName:     param.IdentifyName,
	//	Status:           int64(param.Status),
	//	Root:             param.Root,
	//	Type:             param.Type,
	//	FatherIdentifyId: param.FatherIdentifyId,
	//	UpdateTime:       carbon.Now().Timestamp(),
	//	ProvinceCode:     param.ProvinceCode,
	//	ProvinceName:     param.ProvinceName,
	//	CityCode:         param.CityCode,
	//	CityName:         param.CityName,
	//	AreaCode:         param.AreaCode,
	//	AreaName:         param.AreaName,
	//	Location:         param.Location,
	//	LocationX:        param.LocationX,
	//	LocationY:        param.LocationY,
	//}

	data := map[string]interface{}{
		"identify_name":      param.IdentifyName,
		"status":             int64(param.Status),
		"root":               param.Root,
		"type":               param.Type,
		"father_identify_id": param.FatherIdentifyId,
		"update_time":        carbon.Now().Timestamp(),
		"province_code":      param.ProvinceCode,
		"province_name":      param.ProvinceName,
		"city_code":          param.CityCode,
		"city_name":          param.CityName,
		"area_code":          param.AreaCode,
		"area_name":          param.AreaName,
	}

	result := identifyModel.Where("id = ? and status !=?", param.Id, 9).Updates(&data)
	if result.Error != nil {
		tx.Rollback()
		return 0, result.Error
	}

	// 修改管理员用户信息
	//密码加密
	key := gconv.String(config.Cfg.Login.Key)
	pwd := vcrypto.HexEnCrypt(param.PassWord, key, vcrypto.DesCBCEncrypt)
	//identifyIds, _ := json.Marshal([]int64{param.IdentifyId})
	userData := models.User{
		Phone:      param.Phone,
		Username:   param.UserName,
		Email:      param.Email,
		Realname:   param.RealName,
		Password:   pwd,
		Status:     param.UserStatus,
		UpdateTime: carbon.Now().Timestamp(),
		//IdentifyId: param.IdentifyId, // 不可更改归属
	}
	userModel := tx.Model(&models.User{})
	result = userModel.Where("id = ? and status !=?", param.UserId, 9).Updates(&userData)
	if result.Error != nil {
		tx.Rollback()
		return 0, result.Error
	}

	tx.Commit()

	return result.RowsAffected, nil
}

func (a *identifyLogic) Delete(param v1.DeleteReq) (int64, error) {
	result := global.DB.Model(&models.Identify{}).Where("id = ?", param.Id).Updates(map[string]interface{}{
		"status":      9,
		"delete_time": carbon.Now().Timestamp(),
	})
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func (a *identifyLogic) InitIdentifyPermission(identifyType int8, identifyId int64, roleId int64, tx *gorm.DB) error {
	var Itype string
	switch int(identifyType) {
	case models.IDENTIFY_TYPE_SYSTEM:
		Itype = SYSTEM
	case models.IDENTIFY_TYPE_HOTEL:
		Itype = NORMAL
	}
	if Itype == "" {
		return errors.New("初始化权限类型错误")
	}
	hander := a.handlers[Itype]
	err := hander.InitPermission(identifyId, roleId, tx)
	if err != nil {
		return err
	}
	return nil
}

func (a *identifyLogic) GetNoPoliceIdentify(identifyId int64) []models.Identify {
	var thisIden models.Identify
	result := make([]models.Identify, 0)
	global.DB.Model(&models.Identify{}).Where("status =1 and id=?", identifyId).First(&thisIden)
	if thisIden.ID == 0 {
		return result
	}
	if thisIden.Type != gconv.Int8(models.IDENTIFY_TYPE_POLICE) {
		result = append(result, thisIden)
		return result
	}
	result = CollectNonPoliceIdentifies(thisIden.ID)
	return result
}

func CollectNonPoliceIdentifies(identifyId int64) []models.Identify {
	var result []models.Identify
	global.DB.Model(&models.Identify{}).Where("status = 1 and father_identify_id = ?", identifyId).Find(&result)
	collectedIdentifies := make([]models.Identify, 0)
	for _, identify := range result {
		if identify.Type != gconv.Int8(models.IDENTIFY_TYPE_POLICE) {
			collectedIdentifies = append(collectedIdentifies, identify)
		} else {
			collectedIdentifies = append(collectedIdentifies, CollectNonPoliceIdentifies(identify.ID)...)
		}
	}
	return collectedIdentifies
}
