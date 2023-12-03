package ctrl_page

import (
	"encoding/json"
	"fmt"
	"gin/internal/global"
	"gin/internal/library/jwt"
	"gin/internal/global/errcode"
	"gin/internal/modules/admin/v1/models"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gconv"
	"github.com/golang-module/carbon"
)

type CmsPageController struct{}

func (*CmsPageController) Items(c *gin.Context) {
	var param struct {
		Limit       int    `form:"limit" json:"limit"`
		PageIndex   int    `form:"pageIndex" json:"pageIndex"`
		OrderBy     string `form:"orderBy" json:"orderBy"`
		OrderByType string `form:"orderByType" json:"orderByType"`
		Search      string `form:"search" json:"search"`
		SiteId      int    `form:"site_id" json:"site_id" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS,err.Error())
		return
	}
	//fmt.Println(param.Search,"===param")
	model := global.DB.Model(&models.CmsPage{}).Joins("left join cms_template on cms_page.template_id=cms_template.id")
	// 查询字段
	model.Select("cms_page.id,cms_page.status,cms_page.site_id,cms_page.subject,cms_page.title,cms_page.keywords," +
		"cms_page.description,cms_page.content,cms_page.url,cms_page.image_url,cms_page.create_time,cms_page.update_time," +
		"cms_template.template_name,cms_template.type as template_type," +
		"cms_page.template_id,cms_page.classify_id,cms_page.author_id,cms_page.product_id,cms_page.star_number,cms_page.tag_ids," +
		"cms_page.create_user_name,cms_page.create_user_id")
	//搜索添加表名
	// {"url":{"operator":"like","value":"1212","type":"both"}}

	search := gconv.Map(param.Search)
	if _, ok := search["id"]; ok {
		search["cms_page.id"] = search["id"]
		delete(search, "id")
	}
	if _, ok := search["subject"]; ok {
		search["cms_page.subject"] = search["subject"]
		delete(search, "subject")
	}
	if _, ok := search["url"]; ok {
		search["cms_page.url"] = search["url"]
		delete(search, "url")
	}
	if _, ok := search["status"]; ok {
		search["cms_page.status"] = search["status"]
		delete(search, "status")
	}
	if _, ok := search["template_name"]; ok {
		search["cms_template.template_name"] = search["template_name"]
		delete(search, "template_name")
	}
	//
	model = WhereBySearch(model, search)
	model.Where("cms_page.status != ?", 9)
	model.Where("cms_page.site_id = ?", param.SiteId)
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
		model.Order("cms_page.id desc")
	}
	var result []map[string]interface{}
	model.Scan(&result)
	for k, v := range result {
		var tagArr []int
		_ = json.Unmarshal(gconv.Bytes(v["tag_ids"]), &tagArr)
		result[k]["tag_ids"] = tagArr
	}
	global.Response.Success(c,map[string]interface{}{"items": result, "total": count})
}

func (*CmsPageController) Info(c *gin.Context) {
	var param struct {
		Id int `form:"id" json:"id" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS,err.Error())
		return
	}
	result := map[string]interface{}{}
	global.DB.Model(&models.CmsPage{}).Where("status != ?", 9).First(&result, param.Id)
	var tagIds []int
	_ = json.Unmarshal(gconv.Bytes(result["tag_ids"]), &tagIds)
	result["tag_ids"] = tagIds
	global.Response.Success(c,  result)
}

func (*CmsPageController) Create(c *gin.Context) {
	var param struct {
		SiteId      int    `form:"site_id" json:"site_id" binding:"required"`
		Status      int    `form:"status" json:"status" binding:"required"`
		Subject     string `form:"subject" json:"subject" binding:"required"`
		Title       string `form:"title" json:"title" binding:"required"`
		Keywords    string `form:"keywords" json:"keywords" binding:"required"`
		Description string `form:"description" json:"description" binding:"required"`
		Content     string `form:"content" json:"content" binding:"required"`
		Url         string `form:"url" json:"url" binding:"required"`
		ImageUrl    string `form:"image_url" json:"image_url" binding:"required"`
		TemplateId  int    `form:"template_id" json:"template_id" binding:"required"`
		ClassifyId  int    `form:"classify_id" json:"classify_id" binding:"required"`
		AuthorId    int    `form:"author_id" json:"author_id" binding:"required"`
		ProductId   int    `form:"product_id" json:"product_id" binding:"required"`
		StarNumber  string `form:"star_number" json:"star_number" binding:"required"`
		TagIds      []int  `form:"tag_ids" json:"tag_ids"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS,err.Error())
		return
	}
	// todu   last_update_user_id
	lastUpdateUserId := 0
	token := c.Request.Header.Get("token")
	tokenInfo, err := jwt.ParseJwtGoToken(token)
	if err == nil {
		lastUpdateUserId = gconv.Int(tokenInfo.Id)
	}
	tagJson, err := json.Marshal(param.TagIds)
	if err != nil {
		global.Response.Error(c, errcode.ERROR_SERVER, fmt.Sprintf("tag 序列化失败:%v",err.Error()))
		return
	}
	data := map[string]interface{}{
		"site_id":             param.SiteId,
		"status":              param.Status,
		"subject":             param.Subject,
		"title":               param.Title,
		"keywords":            param.Keywords,
		"description":         param.Description,
		"content":             param.Content,
		"url":                 param.Url,
		"image_url":           param.ImageUrl,
		"template_id":         param.TemplateId,
		"classify_id":         param.ClassifyId,
		"author_id":           param.AuthorId,
		"product_id":          param.ProductId,
		"last_update_user_id": lastUpdateUserId,
		"star_number":         param.StarNumber,
		"tag_ids":             tagJson,
		"create_time":         carbon.Now().Timestamp(),
		"create_user_id":      lastUpdateUserId,
		"create_user_name":    tokenInfo.Audience,
	}
	result := global.DB.Model(&models.CmsPage{}).Create(data)
	if result.Error != nil {
		global.Response.Error(c,errcode.ERROR_SERVER, result.Error.Error())
		return
	}
	global.Response.Success(c,result.RowsAffected)
}

func (*CmsPageController) Update(c *gin.Context) {
	var param struct {
		Id          int    `form:"id" json:"id" binding:"required"`
		Status      int    `form:"status" json:"status" binding:"required"`
		Subject     string `form:"subject" json:"subject" binding:"required"`
		Title       string `form:"title" json:"title" binding:"required"`
		Keywords    string `form:"keywords" json:"keywords" binding:"required"`
		Description string `form:"description" json:"description" binding:"required"`
		Content     string `form:"content" json:"content" binding:"required"`
		Url         string `form:"url" json:"url" binding:"required"`
		ImageUrl    string `form:"image_url" json:"image_url" binding:"required"`
		TemplateId  int    `form:"template_id" json:"template_id" binding:"required"`
		ClassifyId  int    `form:"classify_id" json:"classify_id" binding:"required"`
		AuthorId    int    `form:"author_id" json:"author_id" binding:"required"`
		ProductId   int    `form:"product_id" json:"product_id" binding:"required"`
		StarNumber  string `form:"star_number" json:"star_number" binding:"required"`
		TagIds      []int  `form:"tag_ids" json:"tag_ids"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS,err.Error())
		return
	}
	// todu   last_update_user_id
	lastUpdateUserId := 0
	token := c.Request.Header.Get("token")
	tokenInfo, err := jwt.ParseJwtGoToken(token)
	if err == nil {
		lastUpdateUserId = gconv.Int(tokenInfo.Id)
	}
	tagJson, err := json.Marshal(param.TagIds)
	if err != nil {
		global.Response.Error(c,errcode.ERROR_SERVER, fmt.Sprintf("tag 序列化失败:%v",err.Error()))
		return
	}
	data := map[string]interface{}{
		"status":              param.Status,
		"subject":             param.Subject,
		"title":               param.Title,
		"keywords":            param.Keywords,
		"description":         param.Description,
		"content":             param.Content,
		"url":                 param.Url,
		"image_url":           param.ImageUrl,
		"template_id":         param.TemplateId,
		"classify_id":         param.ClassifyId,
		"author_id":           param.AuthorId,
		"product_id":          param.ProductId,
		"last_update_user_id": lastUpdateUserId,
		"star_number":         param.StarNumber,
		"tag_ids":             tagJson,
		"update_time":         carbon.Now().Timestamp(),
	}
	result := global.DB.Model(&models.CmsPage{}).Where("id = ? and status !=?", param.Id, 9).Updates(data)
	if result.Error != nil {
		global.Response.Error(c, errcode.ERROR_SERVER, result.Error.Error())
		return
	}
	global.Response.Success(c,  result.RowsAffected)
}

// 软删除
func (*CmsPageController) Delete(c *gin.Context) {
	var param struct {
		Id int `form:"id" json:"id" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS,err.Error())
		return
	}
	result := global.DB.Model(&models.CmsPage{}).Where("id = ?", param.Id).Updates(map[string]interface{}{
		"status":      9,
		"delete_time": carbon.Now().Timestamp(),
	})
	if result.Error != nil {
		global.Response.Error(c, errcode.ERROR_SERVER, result.Error.Error())
		return
	}
	global.Response.Success(c, result.RowsAffected)
}

func (*CmsPageController) ArticleCreate(c *gin.Context) {
	var param struct {
		SiteId      int    `form:"site_id" json:"site_id" binding:"required"`
		Subject     string `form:"subject" json:"subject" binding:"required"`
		Title       string `form:"title" json:"title" binding:"required"`
		Keywords    string `form:"keywords" json:"keywords" binding:"required"`
		Description string `form:"description" json:"description" binding:"required"`
		Content     string `form:"content" json:"content" binding:"required"`

		Url        string `form:"url" json:"url" binding:"required"`
		TemplateId int64  `form:"template_id" json:"template_id" binding:"required"`
		ImageUrl   string `form:"image_url" json:"image_url" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS,err.Error())
		return
	}
	// todu   last_update_user_id
	lastUpdateUserId := 0
	token := c.Request.Header.Get("token")
	tokenInfo, err := jwt.ParseJwtGoToken(token)
	if err == nil {
		lastUpdateUserId = gconv.Int(tokenInfo.Id)
	}
	tagJson, err := json.Marshal([]int{})
	if err != nil {
		global.Response.Error(c, errcode.ERROR_SERVER, fmt.Sprintf("tag 序列化失败:%v",err.Error()))
		return
	}

	classify_id := 1
	author_id := 1
	product_id := 1
	star_number := 11

	data := map[string]interface{}{
		"site_id":             param.SiteId,
		"status":              1,
		"subject":             param.Subject,
		"title":               param.Title,
		"keywords":            param.Keywords,
		"description":         param.Description,
		"content":             param.Content,
		"url":                 param.Url,
		"image_url":           param.ImageUrl,
		"template_id":         param.TemplateId,
		"classify_id":         classify_id,
		"author_id":           author_id,
		"product_id":          product_id,
		"last_update_user_id": lastUpdateUserId,
		"star_number":         star_number,
		"tag_ids":             tagJson,
		"create_time":         carbon.Now().Timestamp(),
		"create_user_id":      lastUpdateUserId,
		"create_user_name":    tokenInfo.Audience,
	}
	result := global.DB.Model(&models.CmsPage{}).Create(data)
	if result.Error != nil {
		global.Response.Error(c, errcode.ERROR_SERVER, result.Error.Error())
		return
	}

	global.Response.Success(c,  result.RowsAffected)
}

func (*CmsPageController) ArticleUpdate(c *gin.Context) {
	var param struct {
		Id          int    `form:"id" json:"id" binding:"required"`
		Subject     string `form:"subject" json:"subject" binding:"required"`
		Title       string `form:"title" json:"title" binding:"required"`
		Keywords    string `form:"keywords" json:"keywords" binding:"required"`
		Description string `form:"description" json:"description" binding:"required"`
		Content     string `form:"content" json:"content" binding:"required"`

		Url        string `form:"url" json:"url" binding:"required"`
		TemplateId int64  `form:"template_id" json:"template_id" binding:"required"`
		ImageUrl   string `form:"image_url" json:"image_url" binding:"required"`
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Error(c, errcode.ERROR_PARAMS,err.Error())
		return
	}

	classify_id := 1
	author_id := 1
	product_id := 1
	star_number := 11

	// todu   last_update_user_id
	lastUpdateUserId := 0
	token := c.Request.Header.Get("token")
	tokenInfo, err := jwt.ParseJwtGoToken(token)
	if err == nil {
		lastUpdateUserId = gconv.Int(tokenInfo.Id)
	}
	tagJson, err := json.Marshal([]int{})
	if err != nil {
		global.Response.Error(c, errcode.ERROR_SERVER, fmt.Sprintf("tag 序列化失败:%v",err.Error()))
		return
	}
	data := map[string]interface{}{
		"status":              1,
		"subject":             param.Subject,
		"title":               param.Title,
		"keywords":            param.Keywords,
		"description":         param.Description,
		"content":             param.Content,
		"url":                 param.Url,
		"image_url":           param.ImageUrl,
		"template_id":         param.TemplateId,
		"classify_id":         classify_id,
		"author_id":           author_id,
		"product_id":          product_id,
		"last_update_user_id": lastUpdateUserId,
		"star_number":         star_number,
		"tag_ids":             tagJson,
		"update_time":         carbon.Now().Timestamp(),
	}
	result := global.DB.Model(&models.CmsPage{}).Where("id = ? and status !=?", param.Id, 9).Updates(data)
	if result.Error != nil {
		global.Response.Error(c, errcode.ERROR_SERVER, result.Error.Error())
		return
	}
	global.Response.Success(c,result.RowsAffected)
}
