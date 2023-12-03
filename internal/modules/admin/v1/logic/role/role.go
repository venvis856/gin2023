package role

import (
	"gin/api/admin/role/v1"
	"gin/internal/global"
	"gin/internal/modules/admin/v1/logic/common"
	"gin/internal/modules/admin/v1/models"
	"gin/internal/modules/admin/v1/service"
	"github.com/golang-module/carbon"
)

type (
	roleLogic struct{}
)

func init() {
	service.RegisterRole(New())
}

func New() service.RoleInterface {
	return &roleLogic{}
}

func (a *roleLogic) Items(param v1.ItemReq) (map[string]interface{}, error) {
	model := global.DB.Model(&models.Role{})
	model = common.WhereBySearch(model, param.Search)
	model.Select("id,vid,role_name,status,type,create_time,update_time")
	model.Where("status != ? and identify_id=?", 9, param.IdentifyId)
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

func (a *roleLogic) Info(param v1.InfoReq) map[string]interface{} {
	result := map[string]interface{}{}
	global.DB.Model(&models.Role{}).Where("status != ?", 9).First(&result, param.Id)
	return result
}

func (a *roleLogic) Create(param v1.CreateReq) (int64, error) {
	uid := service.TableIds().GetAddId("role_service", param.IdentifyId)

	data := models.Role{
		RoleName:   param.RoleName,
		Status:     param.Status,
		Type:       param.Type,
		IdentifyId: param.IdentifyId,
		CreateTime: carbon.Now().Timestamp(),
		Uid:        uid,
	}

	result := global.DB.Model(&models.Role{}).Create(&data)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func (a *roleLogic) Update(param v1.UpdateReq) (int64, error) {
	data := models.Role{
		RoleName:   param.RoleName,
		Status:     param.Status,
		Type:       param.Type,
		IdentifyId: param.IdentifyId,
		UpdateTime: carbon.Now().Timestamp(),
	}
	result := global.DB.Model(&models.Role{}).Where("id = ? and status !=?", param.Id, 9).Updates(&data)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func (a *roleLogic) Delete(param v1.DeleteReq) (int64, error) {
	result := global.DB.Model(&models.Role{}).Where("id = ?", param.Id).Updates(map[string]interface{}{
		"status":      9,
		"delete_time": carbon.Now().Timestamp(),
	})
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
