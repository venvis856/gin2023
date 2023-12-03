package ctrl_page

import (
	"encoding/json"
	"fmt"
	"gin/internal/global"
	"gin/internal/modules/admin/v1/models"

	"gin/internal/global/errcode"

	"gin/internal/library/jwt"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gconv"
	"github.com/golang-module/carbon"

)

type NoteBookController struct{}

func (*NoteBookController) Items(c *gin.Context) {
	var param struct {
		Limit       int    `form:"limit" json:"limit"`
		PageIndex   int    `form:"pageIndex" json:"pageIndex"`
		OrderBy     string `form:"orderBy" json:"orderBy"`
		OrderByType string `form:"orderByType" json:"orderByType"`
		Search      string `form:"search" json:"search"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS,err.Error())
		return
	}
	model := global.DB.Model(&models.Notebook{})
	model = WhereBySearch(model, param.Search)
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
	global.Response.Success(c, map[string]interface{}{"items": result, "total": count})
}

func (*NoteBookController) Info(c *gin.Context) {
	var param struct {
		Date string `form:"date" json:"date" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS,err.Error())
		return
	}
	result := map[string]interface{}{}
	global.DB.Model(&models.Notebook{}).Where("status != ? and date=?", 9, param.Date).First(&result)
	var content interface{}
	err := json.Unmarshal(gconv.Bytes(result["content"]), &content)
	if err != nil {
        global.Response.Error(c,errcode.ERROR_SERVER, fmt.Sprintf("json err:%v",err.Error()))
        return
	}
    result["content"]=content
	global.Response.Success(c,  result)
}

func (*NoteBookController) CreateOrUpdate(c *gin.Context) {
	var param struct {
		Date    string   `form:"date" json:"date" binding:"required"`
		Content interface{} `form:"content" json:"content" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS,err.Error())
		return
	}
	// user_id
	userId := 0
	token := c.Request.Header.Get("token")
	tokenInfo, err := jwt.ParseJwtGoToken(token)
	if err == nil {
		userId = gconv.Int(tokenInfo.Id)
	}
	if userId==0 {
		global.Response.Error(c, errcode.ERROR_SERVER,fmt.Sprintf("user_service token无效,userId=%v",userId))
		return
	}
	contentStr, err := json.Marshal(param.Content)
	if err != nil {
		global.Response.Error(c,errcode.ERROR_SERVER, fmt.Sprintf("json失败:%v",err.Error()))
		return
	}
    findInfo:=make(map[string]interface{})
    global.DB.Model(&models.Notebook{}).Where("status != 9 and date=?",param.Date).First(&findInfo)
    if len(findInfo) == 0 { // 新增
        data := map[string]interface{}{
            "date":        param.Date,
            "content":     contentStr,
            "user_id" : userId,
            "status":      1,
            "create_time": carbon.Now().Timestamp(),
            }
            result := global.DB.Model(&models.Notebook{}).Create(data)
            if result.Error != nil {
                global.Response.Error(c, errcode.ERROR_SERVER, result.Error.Error())
                return
            }
    }else{
        data := map[string]interface{}{
            "content": contentStr,
            "update_time": carbon.Now().Timestamp(),
            }
            result := global.DB.Model(&models.Notebook{}).Where("status !=9 and date=?", param.Date).Updates(data)
            if result.Error != nil {
                global.Response.Error(c, errcode.ERROR_SERVER, result.Error.Error())
                return
            }
    }
	global.Response.Success(c,  "update success")
}

// 软删除
func (*NoteBookController) Delete(c *gin.Context) {
	var param struct {
		Id int `form:"id" json:"id" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS,err.Error())
		return
	}
	result := global.DB.Model(&models.Notebook{}).Where("id = ?", param.Id).Updates(map[string]interface{}{
		"status":      9,
		"delete_time": carbon.Now().Timestamp(),
	})
	if result.Error != nil {
		global.Response.Error(c,errcode.ERROR_SERVER, result.Error.Error())
		return
	}
	global.Response.Success(c,  result.RowsAffected)
}
