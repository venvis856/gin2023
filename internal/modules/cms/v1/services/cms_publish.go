package services

import (
	"errors"
	"fmt"
	"gin/app/library/helper"
	"gin/app/library/jwt"
	"gin/app/models"
	"gin/global"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gconv"
	"github.com/golang-module/carbon"
	"os"
	"path"
)

type CmsPublishService struct{}

type PublishParam struct {
	Id     int `form:"id" json:"id" binding:"required"`
	SiteId int `form:"site_id" json:"site_id" binding:"required"`
}

func (receiver *CmsPublishService) Publish(c *gin.Context, param PublishParam) error {
	// 主记录信息
	auditInfo := make(map[string]interface{})
	global.DB.Model(&models.CmsAudit{}).Where("status != ?", 9).First(&auditInfo, param.Id)
	// 获取详情
	var detailList []map[string]interface{}
	global.DB.Model(&models.CmsAuditDetail{}).Where("audit_id = ? and status !=?", param.Id, 9).Find(&detailList)
	fmt.Println(detailList, "detailList")
	if len(detailList) == 0 {
		return errors.New("没有需要发布的文件")
	}

	siteInfo := map[string]interface{}{}
	global.DB.Model(&models.CmsSite{}).Where("status != ?", 9).First(&siteInfo, param.SiteId)
	thisPath, _ := os.Getwd()
	sourceDir := thisPath + "/resource/file/" + gconv.String(siteInfo["root"])
	if gconv.Int(auditInfo["type"]) == 2 {
		sourceDir = thisPath + "/resource/image/" + gconv.String(siteInfo["root"])
	}
	destinationDir := thisPath + "/resource/rsync/" + gconv.String(siteInfo["root"]) + "/" + gconv.String(param.Id)
	// 获取文件移动
	var pageIdsArr []int
	for _, v := range detailList {
		if gconv.Int(v["page_id"]) != 0 {
			pageIdsArr = append(pageIdsArr, gconv.Int(v["page_id"]))
		}
		source := sourceDir + "/" + gconv.String(v["file_url"])
		destination := destinationDir + "/" + gconv.String(v["file_url"])
		err := helper.MkDir(path.Dir(destination))
		global.Logger.Write("rsync", "info", "同步文件，源文件:"+source+"，目标："+destination)
		if err != nil {
			return errors.New("创建目录失败:" + err.Error() + ",路径" + destination)
		}
		_, err = helper.Copy(source, destination)
		if err != nil {
			return errors.New("生成文件没有找到" + err.Error() + ",路径" + source)
		}
	}
	if gin.Mode() == gin.ReleaseMode {
		// 调用命令发布
		//rsync -avz --progress --password-file=/etc/rsyncd/rsyncd.cms /home/www/cms2021/resource/rsync/4itool/7/ cms@157.245.137.85::4itool
		sh := "rsync -avz --progress --password-file=" + gconv.String(siteInfo["rsync_password_path"]) + " " + destinationDir + "/ " +
			gconv.String(siteInfo["rsync_address"])
		if gconv.Int(auditInfo["type"]) == 2 {
			sh = "rsync -avz --progress --password-file=" + gconv.String(siteInfo["rsync_password_path"]) + " " + destinationDir + "/ " +
				gconv.String(siteInfo["rsync_image_address"])
		}
		go SendFile(sh)
	}

	// 更新页面发布了
	global.DB.Model(&models.CmsPage{}).Where("is_publish != 1 and id in ?", pageIdsArr).Updates(map[string]interface{}{
		"is_publish": 1,
	})

	//user_id
	userId := 0
	token := c.Request.Header.Get("token")
	tokenInfo, err := jwt.ParseJwtGoToken(token)
	if err == nil {
		userId = gconv.Int(tokenInfo.Id)
	}

	// 更新状态
	data := map[string]interface{}{
		"status":          2,
		"publush_user_id": userId,
		"publush_time":    carbon.Now().Timestamp(),
	}
	result := global.DB.Model(&models.CmsAudit{}).Where("id = ? and status !=?", param.Id, 9).Updates(data)
	if result.Error != nil {
		return errors.New(result.Error.Error())
	}
	return nil
}

func SendFile(sh string) {
	global.Logger.Write("rsync", "info", "发布命令:"+sh)
	stdout, err := helper.ExecCommand(sh)
	if err != nil {
		global.Logger.Write("rsync", "info", err)
		//global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, "发送失败"+err.Error(), "")
		return
	}
	global.Logger.Write("rsync", "info", stdout)
}
