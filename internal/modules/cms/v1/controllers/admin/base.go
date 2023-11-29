package admin

import (
	"github.com/gogf/gf/util/gconv"
	"gorm.io/gorm"
)

// 前端请求的参数转换
// w参数格式：w["id"] = {"value":"","operator":"=","type":""}
func WhereBySearch(model *gorm.DB, w interface{}) *gorm.DB {
	m := gconv.Map(w)
outLoop:
	for k, v := range m {
		wv := gconv.Map(v)
		operator := gconv.String(wv["operator"])
		val := gconv.String(wv["value"])
		if operator == "like" {
			switch gconv.String(wv["type"]) {
			case "both":
				val = "%" + val + "%"
			case "left":
				val = "%" + val
			case "right":
				val = val + "%"
			default:
				val = "%" + val + "%"
			}
			model.Where(k+" like ?", val)
		} else if operator == "range" {
			// 格式示例：{value:['xx','xx'],operator:'range',type:"time"}
			rangeValues := gconv.SliceStr(wv["value"])
			if len(rangeValues) >= 1 {
				model.Where(k+" >=? ", rangeValues[0])
			}
			if len(rangeValues) == 2 {
				model.Where(k+" <=? ", rangeValues[1])
			}
			continue outLoop
		} else if operator == "in" {
			model.Where(k+" "+operator+" ? ", wv["value"])
			continue outLoop
		} else {
			model.Where(k+" "+operator+" ? ", wv["value"])
		}
	}

	return model
}
