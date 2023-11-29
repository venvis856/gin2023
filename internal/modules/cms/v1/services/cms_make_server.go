package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gin/app/library/helper"
	"gin/app/models"
	"gin/global"
	"github.com/gogf/gf/util/gconv"
	"github.com/golang-module/carbon"
	"math"
	"os"
	"strings"
	"sync"
	"text/template"
)

type CmsMakeServer struct{}

// 公共服务
func (a *CmsMakeServer) ReplaceContent(page map[string]interface{}, param map[string]interface{}) string {
	option := map[string]interface{}{
		"subject":              page["subject"],
		"title":                page["title"],
		"keywords":             page["keywords"],
		"description":          page["description"],
		"image_url":            page["image_url"],
		"first_make_time":      page["first_make_time"],
		"update_time":          page["update_time"],
		"create_time":          page["create_time"],
		"star_number":          page["star_number"],
		"pageList":             param["pageList"],
		"classifyList":         param["classifyList"],
		"classifyInfo":         param["classifyInfo"],
		"htmlUrl":              param["htmlUrl"],
		"videoList":            param["videoList"],
		"preUrl":               param["preUrl"],
		"nextUrl":              param["nextUrl"],
		"product_name":         param["product_name"],
		"product_download_url": param["product_download_url"],
		"product_buy_url":      param["product_buy_url"],
	}
	// 模块处理
	var moduleList []map[string]interface{}
	var pageModuleIds []string
	_ = json.Unmarshal([]byte(gconv.String(page["module_ids"])), &pageModuleIds)
	global.DB.Model(&models.CmsModule{}).Where("status not in ?", []int64{5, 9}).Find(&moduleList, pageModuleIds)
	for _, v := range moduleList {
		option[gconv.String(v["module_name"])] = a.GinReplace(gconv.String(v["content"]), option)
	}
	//fmt.Println(option, "option")
	// 页面内容
	option["content"] = a.GinReplace(gconv.String(page["content"]), option)
	// 模板内容输出
	templateContent := a.GinReplace(gconv.String(page["template_content"]), option)
	return templateContent
}

func (*CmsMakeServer) GinReplace(content string, option map[string]interface{}) string {
	t, _ := template.New("test").Parse(content)
	writer := new(bytes.Buffer)
	// 渲染字段
	_ = t.Execute(writer, option)
	output := fmt.Sprint(writer)
	return output
}

// 1 2 3 4 5 6 7 8 9 10 11 12 13
func (a *CmsMakeServer) GetPagination(thisSize int, sumSize int, url string) []map[string]interface{} {
	var result []map[string]interface{}
	if thisSize <= 6 {
		var sum int
		if sumSize > 10 {
			sum = 10
		} else {
			sum = sumSize
		}
		for i := 1; i <= sum; i++ {
			result = append(result, map[string]interface{}{
				"url":    a.ReplaceUrl(url, i),
				"active": i == thisSize,
				"size":   i,
			})
		}
	}
	if thisSize > 6 {
		var sum int
		if sumSize > thisSize+4 {
			sum = thisSize + 4
		} else {
			sum = sumSize
		}
		for i := thisSize - 5; i <= sum; i++ {
			result = append(result, map[string]interface{}{
				"url":    a.ReplaceUrl(url, i),
				"active": i == thisSize,
				"size":   i,
			})
		}
	}
	return result
}

func (*CmsMakeServer) ReplaceUrl(url string, index int) string {
	newUrl := url
	d := strings.LastIndex(url, ".")
	if index != 1 {
		newUrl = url[:d] + "/" + gconv.String(index) + url[d:]
	}
	return newUrl
}

var templateType = map[string]int64{
	"home":         1,
	"page":         2,
	"classify":     3,
	"product":      4,
	"review":       5,
	"guide":        6,
	"other":        7,
	"video":        8,
	"video_detail": 9,
	"classify_all": 10,
	"video_tag":    11,
}

type VideoLimitSt struct {
	OffSet int64
	Limit  int64
}

const (
	MakeTypePage     = "page"
	MakeTypeVideoTag = "video_tag"
)

func (a *CmsMakeServer) Make(makeType string, pageIds []int, siteId int, isPreview bool, videoLimit VideoLimitSt) ([]map[string]interface{}, string, error) {
	switch makeType {
	case MakeTypePage:
		return a.MakePage(pageIds, siteId, isPreview, videoLimit)
	case MakeTypeVideoTag: // 这里pageids 是video tag id
		return a.MakeVideoTag(pageIds, siteId, isPreview, videoLimit)
	}

	return []map[string]interface{}{}, "", nil
}

// https://github.com/beauty-code-world/cms/blob/master/application/services/Cms/PageMakeServer.php
func (a *CmsMakeServer) MakePage(pageIds []int, siteId int, isPreview bool, videoLimit VideoLimitSt) ([]map[string]interface{}, string, error) {
	var pageInfo []map[string]interface{}
	var newPageArr []map[string]interface{}
	global.DB.Model(&models.CmsPage{}).
		Select("cms_page.*,cms_template.type as template_type,cms_template.content as template_content,"+
			"cms_template.module_ids as module_ids ").
		Where("cms_page.status !=? and cms_page.status!=? and cms_page.site_id=?", 5, 9, siteId).
		Where("cms_page.id in ?", pageIds).
		Joins("left join cms_template on cms_page.template_id = cms_template.id").
		Scan(&pageInfo)
	if len(pageInfo) == 0 {
		return newPageArr, "", errors.New("没有需要生成的页面")
	}
	// 生成页面
	siteInfo := map[string]interface{}{}
	global.DB.Model(&models.CmsSite{}).Where("status != ?", 9).First(&siteInfo, siteId)
	if siteInfo == nil {
		return newPageArr, "", errors.New("站点信息不存在")
	}
	thisPath, _ := os.Getwd()
	dir := thisPath + "/resource/file/" + gconv.String(siteInfo["root"])
	//onlineDir := gconv.String(siteInfo["online_url"])
	var wg sync.WaitGroup
	wg.Add(len(pageInfo))
	for _, page := range pageInfo {
		page["first_make_time"] = carbon.CreateFromTimestamp(gconv.Int64(page["first_make_time"])).ToDateTimeString()
		page["create_time"] = carbon.CreateFromTimestamp(gconv.Int64(page["create_time"])).ToDateTimeString()
		page["update_time"] = carbon.CreateFromTimestamp(gconv.Int64(page["update_time"])).ToDateTimeString()
		if gconv.Int64(page["template_type"]) == gconv.Int64(templateType["home"]) { // 首页
			content := a.ReplaceContent(page, map[string]interface{}{})
			if isPreview { //预览
				return newPageArr, content, nil
			}
			//写入
			filePath := dir + "/" + gconv.String(page["url"])
			if err := helper.CreateFile(filePath); err != nil {
				return newPageArr, "", errors.New("创建文件失败：" + filePath)
			}
			if err := helper.WriteFile(filePath, content); err != nil {
				return newPageArr, "", errors.New("文件写入失败：" + filePath)
			}
			newPageArr = append(newPageArr, map[string]interface{}{
				"id":  page["id"],
				"url": page["url"],
			})
		}

		if gconv.Int64(page["template_type"]) == gconv.Int64(templateType["page"]) { // 文章
			productInfo := make(map[string]interface{})
			global.DB.Model(&models.CmsProduct{}).Where("status not in (5,9) and id=?", page["product_id"]).First(&productInfo)
			param := make(map[string]interface{})
			param["product_name"] = productInfo["product_name"]
			param["product_download_url"] = productInfo["download_url"]
			param["product_buy_url"] = productInfo["buy_url"]
			content := a.ReplaceContent(page, param)
			if isPreview { //预览
				return newPageArr, content, nil
			}
			filePath := dir + "/" + gconv.String(page["url"])
			if err := helper.CreateFile(filePath); err != nil {
				return newPageArr, "", errors.New("创建文件失败：" + filePath)
			}
			if err := helper.WriteFile(filePath, content); err != nil {
				return newPageArr, "", errors.New("文件写入失败：" + filePath)
			}
			newPageArr = append(newPageArr, map[string]interface{}{
				"id":  page["id"],
				"url": page["url"],
			})
		}

		if gconv.Int64(page["template_type"]) == gconv.Int64(templateType["classify"]) { // 分类
			fmt.Println("进来了分类页")
			classifyInfo := map[string]interface{}{}
			global.DB.Model(&models.CmsClassify{}).Where("status != ?", 9).First(&classifyInfo, page["classify_id"])
			if classifyInfo == nil {
				return newPageArr, "", errors.New("分类信息不存在")
			}
			fmt.Println(classifyInfo, "classifyInfoclassifyInfoclassifyInfo")
			// 左边分类列表
			var classifyList []map[string]interface{}
			global.DB.Model(&models.CmsPage{}).Joins(" left join cms_classify on cms_page.classify_id=cms_classify.id").
				Joins("left join cms_template on  cms_page.template_id=cms_template.id").
				Where("cms_page.site_id=? and cms_page.status != ? and cms_page.is_publish = ?", siteId, 9, 1).
				Where("cms_template.type=?", templateType["classify"]).
				Where("cms_classify.is_howto != ?", 1).
				Select("cms_classify.classify_name,cms_page.url").Limit(18).Scan(&classifyList)
			var newClassifyList []map[string]interface{}
			for _, v := range classifyList {
				v["url"] = gconv.String(siteInfo["online_url"]) + "/" + gconv.String(v["url"])
				newClassifyList = append(newClassifyList, v)
			}
			fmt.Println(newClassifyList, "===newClassifyList=====")
			param := make(map[string]interface{})
			param["classifyList"] = newClassifyList
			param["classifyInfo"] = classifyInfo
			pageModel := global.DB.Model(&models.CmsPage{}).Joins("left join cms_template on cms_template.id=cms_page.template_id")
			if gconv.Int(classifyInfo["is_howto"]) == 1 { // 聚合页
				pageModel.Where("cms_page.site_id=? and cms_page.status not in (5,9) and cms_page.is_publish=1 and cms_template.type=2", siteId)
			} else { // 分类页
				pageModel.Where("cms_page.site_id=? and cms_page.status not in (5,9) and cms_page.is_publish=1 and cms_template.type=2 and cms_page.classify_id=? ", siteId, gconv.Int64(page["classify_id"]))
			}
			var pageSumCount int64
			pageModel.Count(&pageSumCount)

			var pageSize float64
			limit := 10
			if pageSumCount == 0 {
				pageSize = 1
			} else {
				pageSize = math.Ceil(gconv.Float64(pageSumCount) / gconv.Float64(limit))
			}
			for i := 1; i <= gconv.Int(pageSize); i++ {
				var pageList []map[string]interface{}
				pageModel.Offset((i - 1) * limit).Limit(limit).Find(&pageList)
				fmt.Println(pageList, "===pageList")
				var newPageList []map[string]interface{}
				// url修改
				for _, v := range pageList {
					v["url"] = gconv.String(siteInfo["online_url"]) + "/" + gconv.String(v["url"])
					v["first_make_time"] = carbon.CreateFromTimestamp(gconv.Int64(v["first_make_time"])).ToDateTimeString()
					newPageList = append(newPageList, v)
				}
				//fmt.Println(newPageList,"newPageList=================1===")
				param["pageList"] = newPageList
				// 翻页
				newPageUrl := gconv.String(siteInfo["online_url"]) + "/" + gconv.String(page["url"])
				param["htmlUrl"] = a.GetPagination(i, gconv.Int(pageSize), newPageUrl)
				// 上页
				preUrl := a.ReplaceUrl(newPageUrl, i-1)
				if i == 1 {
					preUrl = a.ReplaceUrl(newPageUrl, i)
				}
				// 下页
				nextUrl := a.ReplaceUrl(newPageUrl, i+1)
				if i == gconv.Int(pageSize) {
					nextUrl = a.ReplaceUrl(newPageUrl, i)
				}
				param["preUrl"] = preUrl
				param["nextUrl"] = nextUrl
				content := a.ReplaceContent(page, param)
				//fmt.Println(content,"============content")
				if isPreview { //预览
					return newPageArr, content, nil
				}

				url := a.ReplaceUrl(gconv.String(page["url"]), i)
				filePath := dir + "/" + gconv.String(url)
				if err := helper.CreateFile(filePath); err != nil {
					return newPageArr, "", errors.New("创建文件失败：" + filePath)
				}
				if err := helper.WriteFile(filePath, content); err != nil {
					return newPageArr, "", errors.New("文件写入失败：" + filePath)
				}
				newPageArr = append(newPageArr, map[string]interface{}{
					"id":  page["id"],
					"url": url,
				})
			}
		}

		if gconv.Int64(page["template_type"]) == gconv.Int64(templateType["product"]) { // 产品
			content := a.ReplaceContent(page, map[string]interface{}{})
			if isPreview { //预览
				return newPageArr, content, nil
			}
			filePath := dir + "/" + gconv.String(page["url"])
			if err := helper.CreateFile(filePath); err != nil {
				return newPageArr, "", errors.New("创建文件失败：" + filePath)
			}
			if err := helper.WriteFile(filePath, content); err != nil {
				return newPageArr, "", errors.New("文件写入失败：" + filePath)
			}
			newPageArr = append(newPageArr, map[string]interface{}{
				"id":  page["id"],
				"url": page["url"],
			})
		}

		if gconv.Int64(page["template_type"]) == gconv.Int64(templateType["review"]) { // review

		}
		if gconv.Int64(page["template_type"]) == gconv.Int64(templateType["guide"]) { // guide
			content := a.ReplaceContent(page, map[string]interface{}{})
			if isPreview { //预览
				return newPageArr, content, nil
			}
			//写入
			filePath := dir + "/" + gconv.String(page["url"])
			if err := helper.CreateFile(filePath); err != nil {
				return newPageArr, "", errors.New("创建文件失败：" + filePath)
			}
			if err := helper.WriteFile(filePath, content); err != nil {
				return newPageArr, "", errors.New("文件写入失败：" + filePath)
			}
			newPageArr = append(newPageArr, map[string]interface{}{
				"id":  page["id"],
				"url": page["url"],
			})
		}

		if gconv.Int64(page["template_type"]) == gconv.Int64(templateType["video"]) { // 视频
			param := make(map[string]interface{})
			videoModel := global.DB.Model(&models.Video{}).Where("site_id = ? and status =1", siteId)
			var videoSumCount int64
			videoModel.Count(&videoSumCount)

			var pageSize float64
			limit := 9
			if videoSumCount == 0 {
				pageSize = 1
			} else {
				pageSize = math.Ceil(gconv.Float64(videoSumCount) / gconv.Float64(limit))
			}
			for i := 1; i <= gconv.Int(pageSize); i++ {
				var videoList []map[string]interface{}
				videoModel.Offset((i - 1) * limit).Limit(limit).Find(&videoList)
				for k, v := range videoList {
					videoList[k]["url"] = gconv.String(siteInfo["online_url"]) + "/" + helper.SubjectToUrl(gconv.String(v["subject"]), "video")
				}

				//var newPageList []map[string]interface{}
				// url修改
				//for _, v := range pageList {
				//	v["url"] = gconv.String(siteInfo["online_url"]) + "/" + gconv.String(v["url"])
				//	v["first_make_time"] = carbon.CreateFromTimestamp(gconv.Int64(v["first_make_time"])).ToDateTimeString()
				//	v["create_time"] = carbon.CreateFromTimestamp(gconv.Int64(v["create_time"])).ToDateTimeString()
				//	newPageList = append(newPageList, v)
				//}
				// 翻页
				newPageUrl := gconv.String(siteInfo["online_url"]) + "/" + gconv.String(page["url"])
				fmt.Println(pageSize, newPageUrl, "======video==")

				param["htmlUrl"] = a.GetPagination(i, gconv.Int(pageSize), newPageUrl)
				global.Logger.Write("test", "info", i, pageSize, param["htmlUrl"])
				// 上页
				preUrl := a.ReplaceUrl(newPageUrl, i-1)
				if i == 1 {
					preUrl = a.ReplaceUrl(newPageUrl, i)
				}
				// 下页
				nextUrl := a.ReplaceUrl(newPageUrl, i+1)
				if i == gconv.Int(pageSize) {
					nextUrl = a.ReplaceUrl(newPageUrl, i)
				}
				param["preUrl"] = preUrl
				param["nextUrl"] = nextUrl
				param["videoList"] = videoList
				content := a.ReplaceContent(page, param)
				//fmt.Println(content,"============content")
				if isPreview { //预览
					return newPageArr, content, nil
				}

				url := a.ReplaceUrl(gconv.String(page["url"]), i)
				filePath := dir + "/" + gconv.String(url)
				if err := helper.CreateFile(filePath); err != nil {
					return newPageArr, "", errors.New("创建文件失败：" + filePath)
				}
				if err := helper.WriteFile(filePath, content); err != nil {
					return newPageArr, "", errors.New("文件写入失败：" + filePath)
				}
				newPageArr = append(newPageArr, map[string]interface{}{
					"id":  page["id"],
					"url": url,
				})
			}

		}

		if gconv.Int64(page["template_type"]) == gconv.Int64(templateType["video_detail"]) { // 视频详情页

			// 更新搜索页
			content2 := a.ReplaceContent(page, map[string]interface{}{})
			if isPreview { //预览
				return newPageArr, content2, nil
			}
			//写入
			filePath := dir + "/" + gconv.String(page["url"])
			if err := helper.CreateFile(filePath); err != nil {
				return newPageArr, "", errors.New("创建文件失败：" + filePath)
			}
			if err := helper.WriteFile(filePath, content2); err != nil {
				return newPageArr, "", errors.New("文件写入失败：" + filePath)
			}
			newPageArr = append(newPageArr, map[string]interface{}{
				"id":  page["id"],
				"url": page["url"],
			})

			param := make(map[string]interface{})
			var videoList []models.Video
			videoModel := global.DB.Model(&models.Video{}).Where("site_id = ? and status =1", siteId)
			if videoLimit.Limit != 0 {
				videoModel.Offset(int(videoLimit.OffSet)).Limit(int(videoLimit.Limit))
			}
			videoModel.Find(&videoList)
			for _, v := range videoList {
				page["subject"] = v.Subject
				page["title"] = v.Title
				page["description"] = "Download Pornhub Video: " + v.Subject + ", Using our downloader to download this video work is completely free."
				page["image_url"] = v.Thumbnail
				content := a.ReplaceContent(page, param)
				if isPreview { //预览
					return newPageArr, content, nil
				}
				NewUrl := helper.SubjectToUrl(v.Subject, "video")
				//写入
				filePath := dir + "/" + NewUrl
				if err := helper.CreateFile(filePath); err != nil {
					return newPageArr, "", errors.New("创建文件失败：" + filePath)
				}
				if err := helper.WriteFile(filePath, content); err != nil {
					return newPageArr, "", errors.New("文件写入失败：" + filePath)
				}
				newPageArr = append(newPageArr, map[string]interface{}{
					"id":  page["id"],
					"url": NewUrl,
				})
			}

		}

		if gconv.Int64(page["template_type"]) == gconv.Int64(templateType["classify_all"]) { // 统一分类
			param := make(map[string]interface{})
			pageModel := global.DB.Model(&models.CmsPage{}).Joins("left join cms_template on cms_template.id=cms_page.template_id")
			pageModel.Where("cms_page.site_id=? and cms_page.status not in (5,9) and cms_page.is_publish=1 and cms_template.type=2", siteId)
			var pageSumCount int64
			pageModel.Count(&pageSumCount)
			var pageSize float64
			limit := 9
			if pageSumCount == 0 {
				pageSize = 1
			} else {
				pageSize = math.Ceil(gconv.Float64(pageSumCount) / gconv.Float64(limit))
			}
			for i := 1; i <= gconv.Int(pageSize); i++ {
				var pageList []map[string]interface{}
				pageModel.Offset((i - 1) * limit).Limit(limit).Find(&pageList)
				var newPageList []map[string]interface{}
				// url修改
				for _, v := range pageList {
					v["url"] = gconv.String(siteInfo["online_url"]) + "/" + gconv.String(v["url"])
					v["first_make_time"] = carbon.CreateFromTimestamp(gconv.Int64(v["first_make_time"])).ToDateTimeString()
					v["create_time"] = carbon.CreateFromTimestamp(gconv.Int64(v["create_time"])).ToDateTimeString()
					v["subject"] = gconv.String(v["subject"])
					v["image_url"] = gconv.String(v["image_url"])
					newPageList = append(newPageList, v)
				}
				//fmt.Println(newPageList,"newPageList=================1===")
				param["pageList"] = newPageList
				// 翻页
				newPageUrl := gconv.String(siteInfo["online_url"]) + "/" + gconv.String(page["url"])
				param["htmlUrl"] = a.GetPagination(i, gconv.Int(pageSize), newPageUrl)
				// 上页
				preUrl := a.ReplaceUrl(newPageUrl, i-1)
				if i == 1 {
					preUrl = a.ReplaceUrl(newPageUrl, i)
				}
				// 下页
				nextUrl := a.ReplaceUrl(newPageUrl, i+1)
				if i == gconv.Int(pageSize) {
					nextUrl = a.ReplaceUrl(newPageUrl, i)
				}
				param["preUrl"] = preUrl
				param["nextUrl"] = nextUrl
				content := a.ReplaceContent(page, param)
				//fmt.Println(content,"============content")
				if isPreview { //预览
					return newPageArr, content, nil
				}

				url := a.ReplaceUrl(gconv.String(page["url"]), i)
				filePath := dir + "/" + gconv.String(url)
				if err := helper.CreateFile(filePath); err != nil {
					return newPageArr, "", errors.New("创建文件失败：" + filePath)
				}
				if err := helper.WriteFile(filePath, content); err != nil {
					return newPageArr, "", errors.New("文件写入失败：" + filePath)
				}
				newPageArr = append(newPageArr, map[string]interface{}{
					"id":  page["id"],
					"url": url,
				})
			}
		}

		if gconv.Int64(page["template_type"]) == gconv.Int64(templateType["other"]) { // 其他
			content := a.ReplaceContent(page, map[string]interface{}{})
			if isPreview { //预览
				return newPageArr, content, nil
			}
			//写入
			filePath := dir + "/" + gconv.String(page["url"])
			if err := helper.CreateFile(filePath); err != nil {
				return newPageArr, "", errors.New("创建文件失败：" + filePath)
			}
			if err := helper.WriteFile(filePath, content); err != nil {
				return newPageArr, "", errors.New("文件写入失败：" + filePath)
			}
			newPageArr = append(newPageArr, map[string]interface{}{
				"id":  page["id"],
				"url": page["url"],
			})
		}

		wg.Done()
	}
	// 生成记录
	wg.Wait()
	return newPageArr, "", nil
}

func (a *CmsMakeServer) MakeVideoTag(videoTagIds []int, siteId int, isPreview bool, videoLimit VideoLimitSt) ([]map[string]interface{}, string, error) {
	var templateInfo models.CmsTemplate
	global.DB.Model(&models.CmsTemplate{}).Where("site_id = ? and type =?", siteId, templateType["video_tag"]).First(&templateInfo)

	var newPageArr []map[string]interface{}

	// 站点id
	siteInfo := map[string]interface{}{}
	global.DB.Model(&models.CmsSite{}).Where("status != ?", 9).First(&siteInfo, siteId)
	if siteInfo == nil {
		return newPageArr, "", errors.New("站点信息不存在")
	}

	var videoTagList []map[string]interface{}
	global.DB.Model(&models.VideoTag{}).Where("id in ?", videoTagIds).Find(&videoTagList)
	if len(videoTagList) == 0 {
		return newPageArr, "", errors.New("没有需要生成的页面")
	}

	thisPath, _ := os.Getwd()
	dir := thisPath + "/resource/file/" + gconv.String(siteInfo["root"])
	//onlineDir := gconv.String(siteInfo["online_url"])
	var wg sync.WaitGroup
	wg.Add(len(videoTagList))
	//fmt.Println(videoTagList, "===videoTagList")
	for _, page := range videoTagList {
		page["template_content"] = templateInfo.Content
		page["module_ids"] = templateInfo.ModuleIds
		page["first_make_time"] = carbon.CreateFromTimestamp(gconv.Int64(page["create_time"])).ToDateTimeString()
		page["create_time"] = carbon.CreateFromTimestamp(gconv.Int64(page["create_time"])).ToDateTimeString()
		page["update_time"] = carbon.CreateFromTimestamp(gconv.Int64(page["update_time"])).ToDateTimeString()
		page["template_type"] = templateInfo.Type
		fmt.Println(gconv.Int64(page["template_type"]) == gconv.Int64(templateType["video_tag"]), "====================")
		if gconv.Int64(page["template_type"]) == gconv.Int64(templateType["video_tag"]) { // 视频 tag
			videoModel := global.DB.Model(&models.Video{}).Where("status = 1 and site_id = ?", siteId).
				Where("tag_ids like ?", fmt.Sprintf("%%,%v,%%", page["id"])).
				Or("tag_ids like  ?", fmt.Sprintf("[1,%%")).
				Or("tag_ids like  ?", fmt.Sprintf("%%,2]")).
				Or("tag_ids like  ?", fmt.Sprintf("[2]"))

			param := make(map[string]interface{})
			var videoSumCount int64
			videoModel.Count(&videoSumCount)

			var pageSize float64
			limit := 9
			if videoSumCount == 0 {
				pageSize = 1
			} else {
				pageSize = math.Ceil(gconv.Float64(videoSumCount) / gconv.Float64(limit))
			}
			// 补充
			page["subject"] = page["name"]
			page["url"] = helper.SubjectToUrl(gconv.String(page["name"]), "video_tag")
			//fmt.Println(page, "=======page")
			for i := 1; i <= gconv.Int(pageSize); i++ {
				var videoList []map[string]interface{}
				videoModel.Offset((i - 1) * limit).Limit(limit).Find(&videoList)
				fmt.Println(videoList, "=====videoList")
				for k, v := range videoList {
					videoList[k]["url"] = gconv.String(siteInfo["online_url"]) + "/" + helper.SubjectToUrl(gconv.String(v["subject"]), "video")
				}

				// 翻页
				newPageUrl := gconv.String(siteInfo["online_url"]) + "/" + gconv.String(page["url"])
				//fmt.Println(pageSize, newPageUrl, "======video==")

				param["htmlUrl"] = a.GetPagination(i, gconv.Int(pageSize), newPageUrl)
				global.Logger.Write("test", "info", i, pageSize, param["htmlUrl"])
				// 上页
				preUrl := a.ReplaceUrl(newPageUrl, i-1)
				if i == 1 {
					preUrl = a.ReplaceUrl(newPageUrl, i)
				}
				// 下页
				nextUrl := a.ReplaceUrl(newPageUrl, i+1)
				if i == gconv.Int(pageSize) {
					nextUrl = a.ReplaceUrl(newPageUrl, i)
				}
				param["preUrl"] = preUrl
				param["nextUrl"] = nextUrl
				param["videoList"] = videoList
				content := a.ReplaceContent(page, param)
				//fmt.Println(content,"============content")
				if isPreview { //预览
					return newPageArr, content, nil
				}

				url := a.ReplaceUrl(gconv.String(page["url"]), i)
				fmt.Println(url, "=====url", page)
				filePath := dir + "/" + gconv.String(url)
				if err := helper.CreateFile(filePath); err != nil {
					return newPageArr, "", errors.New("创建文件失败：" + filePath)
				}
				if err := helper.WriteFile(filePath, content); err != nil {
					return newPageArr, "", errors.New("文件写入失败：" + filePath)
				}
				newPageArr = append(newPageArr, map[string]interface{}{
					"id":  page["id"],
					"url": url,
				})
			}

		}

		if gconv.Int64(page["template_type"]) == gconv.Int64(templateType["other"]) { // 其他
			content := a.ReplaceContent(page, map[string]interface{}{})
			if isPreview { //预览
				return newPageArr, content, nil
			}
			//写入
			filePath := dir + "/" + gconv.String(page["url"])
			if err := helper.CreateFile(filePath); err != nil {
				return newPageArr, "", errors.New("创建文件失败：" + filePath)
			}
			if err := helper.WriteFile(filePath, content); err != nil {
				return newPageArr, "", errors.New("文件写入失败：" + filePath)
			}
			newPageArr = append(newPageArr, map[string]interface{}{
				"id":  page["id"],
				"url": page["url"],
			})
		}

		wg.Done()
	}
	// 生成记录
	wg.Wait()
	return newPageArr, "", nil
}
