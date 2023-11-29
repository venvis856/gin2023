package identify

import (
	"errors"
	"fmt"
	"gin/api/admin/identify/v1"
	"gin/internal/global"
	"gin/internal/library/vcrypto"
	"gin/internal/modules/admin/v1/config"
	"gin/internal/modules/admin/v1/logic/common"
	"gin/internal/modules/admin/v1/models"
	"gin/internal/modules/admin/v1/service"
	"github.com/gogf/gf/util/gconv"
	"github.com/golang-module/carbon"
)

type identifyLogic struct{}

func init() {
	service.RegisterIdentify(New())
}

func New() service.IdentifyInterface {
	return &identifyLogic{}
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
	result := make(map[string]interface{})
	identifyInfo := make(map[string]interface{})
	global.DB.Model(&models.Identify{}).Where("status != ?", 9).First(&identifyInfo, param.Id)
	// 先找到系统管理员角色 , 身份详情会把系统管理员的信息也带过去
	var roleInfo models.Role
	global.DB.Model(&models.Role{}).Where("status != 9 and type=1 and identify_id=?", param.Id).First(&roleInfo)
	if roleInfo.ID == 0 {
		return result, errors.New("该系统管理员异常")
	}
	// 根据用户角色表，找到管理员的user id
	var userIdsSlice []int64
	err := global.DB.Model(&models.UserRole{}).Where("is_effective == 1 and identify_id = ? and role_id= ?", param.Id, roleInfo.ID).Pluck("id", &userIdsSlice).Error
	if err != nil {
		return result, errors.New(fmt.Sprintf("role get user id err :%v", err))
	}
	var adminUserList []models.User
	global.DB.Model(&models.User{}).Where("status != 9 and id in  ?", userIdsSlice).
		First(&adminUserList)
	result["identify"] = identifyInfo
	result["admin_user_list"] = adminUserList
	return result, nil
}

func (a *identifyLogic) Create(param v1.CreateReq) (int64, error) {
	// 校验如果新增的是系统，只能初始化一次
	if int(param.Type) == models.IDENTIFY_TYPE_SYSTEM {
		var systemIdentifyInfo models.Identify
		global.DB.Model(&models.Identify{}).Where("type = ?", models.IDENTIFY_TYPE_SYSTEM).First(&systemIdentifyInfo)
		if systemIdentifyInfo.ID != 0 {
			return 0, errors.New("系统身份已经初始化")
		}
	}

	// 开始新增
	tx := global.DB.Begin()
	model := tx.Model(&models.Identify{})

	// add identify
	identifyData := models.Identify{
		IdentifyName:     param.IdentifyName,
		IdentifyCode:     param.IdentifyCode,
		Status:           param.Status,
		Type:             param.Type,
		FatherIdentifyId: param.FatherIdentifyId,
		CreateTime:       carbon.Now().Timestamp(),
	}
	result := model.Create(&identifyData)
	if result.Error != nil {
		tx.Rollback()
		return 0, errors.New(fmt.Sprintf("Identify create err :%v", result.Error))
	}

	// 新增角色
	roleModel := tx.Model(&models.Role{})
	roleUid := service.TableIds().GetAddId("role", identifyData.ID)
	roleData := models.Role{
		RoleName:   "系统管理员",
		Status:     1,
		Type:       1,
		IdentifyId: identifyData.ID,
		Uid:        roleUid,
		CreateTime: carbon.Now().Timestamp(),
	}
	result = roleModel.Create(&roleData)
	if result.Error != nil {
		tx.Rollback()
		return 0, errors.New(fmt.Sprintf("role create err :%v", result.Error))
	}

	// 新增管理员
	userModel := tx.Model(&models.User{})
	userUid := service.TableIds().GetAddId("user", identifyData.ID)
	//密码加密
	key := gconv.String(global.Cfg.Login.Key)
	pwd := vcrypto.HexEnCrypt(param.PassWord, key, vcrypto.DesCBCEncrypt)
	userData := models.User{
		Phone:      param.Phone,
		Username:   param.UserName,
		Email:      param.Email,
		Realname:   param.RealName,
		Password:   pwd,
		Status:     param.UserStatus,
		CreateTime: carbon.Now().Timestamp(),
		Uid:        userUid,
	}
	result = userModel.Create(&userData)
	if result.Error != nil {
		tx.Rollback()
		return 0,  errors.New(fmt.Sprintf("user create err :%v", result.Error))
	}

	// 新增 用户角色表
	userRoleData := models.UserRole{
		IdentifyId:  identifyData.ID,
		UserId:      userData.ID,
		RoleId:      roleData.ID,
		IsEffective: config.EFFECTIVE_YES,
	}
	result = tx.Model(&models.UserRole{}).Create(&userRoleData)
	if result.Error != nil {
		tx.Rollback()
		return 0,  errors.New(fmt.Sprintf("user role create err :%v", result.Error))
	}

	// 如新增系统，还需要新增系统的权限
	if int(param.Type) == models.IDENTIFY_TYPE_SYSTEM {
		allPermission := service.PermissionOperate().GetAllPermissionCodes()
		err := service.PermissionOperate().IdentifyAddPermission(identifyData.ID, allPermission)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	tx.Commit()
	return result.RowsAffected, nil
}

func (a *identifyLogic) Update(param v1.UpdateReq) (int64, error) {
	// 如果系统类型存在，则其他类型不能修改为系统类型
	if int(param.Type) == models.IDENTIFY_TYPE_SYSTEM {
		var systemIdentifyInfo models.Identify
		global.DB.Model(&models.Identify{}).Where("type = ? and id !=? ", models.IDENTIFY_TYPE_SYSTEM, param.Id).First(&systemIdentifyInfo)
		if systemIdentifyInfo.ID != 0 {
			return 0, errors.New("无法修改为系统类型")
		}
	}

	tx := global.DB.Begin()
	identifyModel := tx.Model(&models.Identify{})

	data := map[string]interface{}{
		"identify_name":      param.IdentifyName,
		"identify_code":      param.IdentifyCode,
		"status":             param.Status,
		"type":               param.Type,
		"father_identify_id": param.FatherIdentifyId,
		"update_time":        carbon.Now().Timestamp(),
	}

	result := identifyModel.Where("id = ? and status !=?", param.Id, 9).Updates(&data)
	if result.Error != nil {
		tx.Rollback()
		return 0, result.Error
	}

	// 修改管理员用户信息
	//密码加密
	key := gconv.String(global.Cfg.Login.Key)
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
	var info models.Identify
	global.DB.Model(&models.Identify{}).Where("id = ?", param.Id).First(&info)
	if info.ID != 0 {
		return 0, errors.New("不存在的标识")
	}

	if int(info.Type) == models.IDENTIFY_TYPE_SYSTEM {
		return 0, errors.New("系统类型不可删除")
	}

	result := global.DB.Model(&models.Identify{}).Where("id = ?", param.Id).Updates(map[string]interface{}{
		"status":      9,
		"delete_time": carbon.Now().Timestamp(),
	})
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
