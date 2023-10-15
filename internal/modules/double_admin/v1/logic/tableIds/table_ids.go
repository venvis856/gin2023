package tableIds

import (
	"gin/internal/global"
	"gin/internal/modules/double_admin/v1/models"
	"gin/internal/modules/double_admin/v1/service"
	"github.com/gogf/gf/util/gconv"
)

type tableIds struct{}

func init() {
	service.RegisterTableIds(New())
}

func New() service.TableIdsInterface {
	return &tableIds{}
}

func (receiver *tableIds) GetAddId(tableName string, identifyId int64) int64 {
	info := make(map[string]interface{})
	global.DB.Model(&models.TableIds{}).Select("max_id").Where("table_name=? and identify_id=?", tableName, identifyId).First(&info)
	if len(info) == 0 {
		global.DB.Model(&models.TableIds{}).Create(map[string]interface{}{
			"table_name":  tableName,
			"identify_id": identifyId,
			"max_id":      1,
		})
		return 1
	}
	newId := gconv.Int64(info["max_id"]) + 1
	global.DB.Model(&models.TableIds{}).Where("table_name=? and identify_id=?", tableName, identifyId).Updates(map[string]interface{}{
		"max_id": newId,
	})
	return newId
}
