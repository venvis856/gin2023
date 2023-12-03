package ctrl_page

import (
	"encoding/json"
	"fmt"
	"gin/internal/global"
	"gin/internal/library/vcrypto"
	"gin/internal/modules/admin/v1/models"

	"gin/internal/global/errcode"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gconv"
	"github.com/golang-module/carbon"
)

type UserController struct{}

func (*UserController) Items(c *gin.Context) {
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
	model := global.DB.Model(&models.User{})
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

func (*UserController) Info(c *gin.Context) {
	var param struct {
		Id int `form:"id" json:"id" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS,err.Error())
		return
	}
	result := map[string]interface{}{}
	global.DB.Model(&models.User{}).Where("status != ?", 9).First(&result, param.Id)
	global.Response.Success(c,  result)
}

func (*UserController) Create(c *gin.Context) {
	var param struct {
		UserName string  `form:"username" json:"user_name" binding:"required"`
		PassWord string  `form:"password" json:"pass_word" binding:"required"`
		Status   int     `form:"status" json:"status" binding:"required"`
		RealName string  `form:"realname" json:"real_name"`
		RoleIds  []int64 `form:"role_ids" json:"role_ids"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS,err.Error())
		return
	}
	//密码加密
	key := gconv.String(global.Cfg.Login.Key)
	pwd := vcrypto.HexEnCrypt(param.PassWord, key, vcrypto.DesCBCEncrypt)
	roleIdsJson, _ := json.Marshal(param.RoleIds)
	roleIds := gconv.String(roleIdsJson)
	fmt.Println(param.RoleIds, roleIdsJson, roleIds, "roleIds")
	data := map[string]interface{}{
		"username":    param.UserName,
		"password":    pwd,
		"status":      param.Status,
		"role_ids":    roleIds,
		"create_time": carbon.Now().Timestamp(),
	}
	if param.RealName != "" {
		data["realname"] = param.RealName
	}
	result := global.DB.Model(&models.User{}).Create(data)
	if result.Error != nil {
		global.Response.Error(c, errcode.ERROR_SERVER, result.Error.Error())
		return
	}
	global.Response.Success(c, result.RowsAffected)
}

func (*UserController) Update(c *gin.Context) {
	var param struct {
		Id       int     `form:"id" json:"id" binding:"required"`
		UserName string  `form:"username" json:"user_name" binding:"required"`
		PassWord string  `form:"password" json:"pass_word" binding:"required"`
		Status   int     `form:"status" json:"status" binding:"required"`
		RealName string  `form:"realname" json:"real_name"`
		RoleIds  []int64 `form:"role_ids" json:"role_ids"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS,err.Error())
		return
	}
	//密码加密
	key := gconv.String(global.Cfg.Login.Key)
	pwd := vcrypto.HexEnCrypt(param.PassWord, key, vcrypto.DesCBCEncrypt)
	roleIdsJson, _ := json.Marshal(param.RoleIds)
	roleIds := gconv.String(roleIdsJson)
	data := map[string]interface{}{
		"userName":    param.UserName,
		"passWord":    pwd,
		"status":      param.Status,
		"role_ids":    roleIds,
		"update_time": carbon.Now().Timestamp(),
	}
	if param.RealName != "" {
		data["realname"] = param.RealName
	}
	result := global.DB.Model(&models.User{}).Where("id = ? and status !=?", param.Id, 9).Updates(data)
	if result.Error != nil {
		global.Response.Error(c, errcode.ERROR_SERVER, result.Error.Error())
		return
	}
	global.Response.Success(c,  result.RowsAffected)
}

// 软删除
func (*UserController) Delete(c *gin.Context) {
	var param struct {
		Id int `form:"id" json:"id" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS,err.Error())
		return
	}
	result := global.DB.Model(&models.User{}).Where("id = ?", param.Id).Updates(map[string]interface{}{
		"status":      9,
		"delete_time": carbon.Now().Timestamp(),
	})
	if result.Error != nil {
		global.Response.Error(c,errcode.ERROR_SERVER, result.Error.Error())
		return
	}
	global.Response.Success(c, result.RowsAffected)
}

// 硬删除
//func (*UserController) Delete(c *gin.Context) {
//	var param struct {
//		Id int `form:"id" binding:"required"`
//	}
//	if err := c.ShouldBind(&param); err != nil {
//		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
//		return
//	}
//	result:=global.DB.Unscoped().Where("ID = ?", param.Id).Delete(&models.User{})
//	if result.Error != nil{
//		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, result.Error.Error(), "")
//		return
//	}
//	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", "success")
//}
