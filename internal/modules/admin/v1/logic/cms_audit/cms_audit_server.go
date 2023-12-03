package services

import (
	"fmt"
	"gin/internal/global"
	"gin/internal/library/helper"
	"gin/internal/modules/admin/v1/models"
	"gin/internal/modules/admin/v1/service"
	"github.com/gogf/gf/util/gconv"
	"github.com/golang-module/carbon"
	"os"
)

type CmsAuditLogic struct{}

func init()  {
	service.RegisterAudit(New())
}

func New() service.AuditInterface {
	return &CmsAuditLogic{}
}

// 生成记录
func (a *CmsAuditLogic) MakeAudit(auditDetailList []map[string]interface{}, makeType int, makeUserId int, siteId int) (int64, error) {
	count := len(auditDetailList)
	if count == 0 {
		return 0, nil
	}
	//fmt.Println(auditDetailList, "auditList")
	auditInfo := map[string]interface{}{}
	global.DB.Model(&models.CmsAudit{}).
		Where("make_user_id =? and type=? and status=? and site_id=?", makeUserId, makeType, 1, siteId).
		First(&auditInfo)
	//fmt.Println(auditInfo, "auditInfo")
	// 站点
	siteInfo := map[string]interface{}{}
	global.DB.Model(&models.CmsSite{}).Where("status != ?", 9).First(&siteInfo, siteId)
	thisPath, _ := os.Getwd()
	dir := thisPath + "/resource/file/" + gconv.String(siteInfo["root"])
	previewDir := gconv.String(siteInfo["preview_url"]) + "/resource/file/" + gconv.String(siteInfo["root"])
	onlineDir := gconv.String(siteInfo["online_url"])
	if makeType == 2 {
		dir = thisPath + "/resource/image/" + gconv.String(siteInfo["root"])
		previewDir = gconv.String(siteInfo["preview_url"]) + "/resource/image/" + gconv.String(siteInfo["root"])
		onlineDir = gconv.String(siteInfo["online_image_url"])
	}
	var auditId int64

	//todu
	if len(auditInfo) == 0 { // 没有找到，全部新增
		auditData := models.CmsAudit{
			FirstUrl:   gconv.String(auditDetailList[0]["url"]),
			Count:      count,
			Type:       makeType,
			MakeUserId: makeUserId,
			SiteId:     gconv.Uint(siteId),
			Status:     1,
			MakeTime:   gconv.Uint(carbon.Now().Timestamp()),
		}
		result := global.DB.Model(&models.CmsAudit{}).Create(&auditData)
		if result.Error != nil {
			return 0, result.Error
		}

		// 详情
		var detailData []map[string]interface{}
		var pageIds []int
		for _, v := range auditDetailList {
			detailData = append(detailData, map[string]interface{}{
				"local_url":   dir + "/" + gconv.String(v["url"]),
				"online_url":  onlineDir + "/" + gconv.String(v["url"]),
				"preview_url": previewDir + "/" + gconv.String(v["url"]),
				"file_url":    v["url"],
				"type":        makeType,
				"status":      1,
				"audit_id":    auditData.Id,
				"page_id":     v["id"], // 非页面是0
				"make_time":   carbon.Now().Timestamp(),
			})
			pageIds = append(pageIds, gconv.Int(v["id"]))
		}
		rs := global.DB.Model(&models.CmsAuditDetail{}).Create(detailData)
		if rs.Error != nil {
			return 0, rs.Error
		}
		// 如果是页面尝试更新第一次时间
		if makeType == 1 {
			err := a.UpdatePageFirstMakeTime(pageIds)
			if err != nil {
				fmt.Println(err, "更新页面第一次时间失败")
			}
		}
		auditId = gconv.Int64(auditData.Id)
	} else { // 已经存在
		// 找出需要新增的记录
		var alreadyData []map[string]interface{}
		var alreadyIds []string
		global.DB.Model(&models.CmsAuditDetail{}).Where("status != ?  and  audit_id=?", 9, auditInfo["id"]).Find(&alreadyData)
		for _, v := range alreadyData {
			alreadyIds = append(alreadyIds, gconv.String(v["file_url"]))
		}
		var addData []map[string]interface{}
		for _, v := range auditDetailList {
			if !helper.InArray(v["url"], alreadyIds) {
				addData = append(addData, v)
			}
		}
		// 更新总数
		count := len(addData)
		rs := global.DB.Model(&models.CmsAudit{}).Where("id=?", auditInfo["id"]).Updates(map[string]interface{}{
			"count": gconv.Int(auditInfo["count"]) + count,
		})
		if rs.Error != nil {
			return 0, rs.Error
		}
		// 详情
		var detailData []map[string]interface{}
		var pageIds []int
		//fmt.Println(addData, "addData")
		if len(addData) != 0 {
			for _, v := range addData {
				detailData = append(detailData, map[string]interface{}{
					"local_url":   dir + "/" + gconv.String(v["url"]),
					"online_url":  onlineDir + "/" + gconv.String(v["url"]),
					"preview_url": previewDir + "/" + gconv.String(v["url"]),
					"file_url":    v["url"],
					"type":        makeType,
					"status":      1,
					"audit_id":    auditInfo["id"],
					"page_id":     v["id"], // 非页面是0
					"make_time":   carbon.Now().Timestamp(),
				})
				pageIds = append(pageIds, gconv.Int(v["id"]))
			}
			rs = global.DB.Model(&models.CmsAuditDetail{}).Create(detailData)
			if rs.Error != nil {
				return 0, rs.Error
			}
		}
		// 如果是页面尝试更新第一次时间
		if makeType == 1 {
			err := a.UpdatePageFirstMakeTime(pageIds)
			if err != nil {
				fmt.Println(err, "更新页面第一次时间失败")
			}
		}
		auditId = gconv.Int64(auditInfo["id"])
	}
	return auditId, nil
}

func (*CmsAuditLogic) UpdatePageFirstMakeTime(pageIds []int) error {
	data := map[string]interface{}{
		"first_make_time": carbon.Now().Timestamp(),
	}
	result := global.DB.Model(&models.CmsPage{}).Where("id in ? and first_make_time=0", pageIds).Updates(data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
