package ctrl_page

import (
	"bufio"
	"fmt"
	"gin/internal/global"
	"gin/internal/global/errcode"
	"gin/internal/library/helper"
	"gin/internal/library/jwt"
	"gin/internal/modules/admin/v1/models"
	"gin/internal/modules/admin/v1/service"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gconv"
	"github.com/golang-module/carbon"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

type CmsPictureController struct{}

func (*CmsPictureController) Upload(c *gin.Context) {
	siteId := c.PostForm("site_id")
	pathUrl := c.DefaultPostForm("path", "")
	//_,file, err := c.Request.FormFile("file")
	if siteId == "" {
		global.Response.Error(c, errcode.ERROR_PARAMS, "参数site_id不对")
		return
	}
	form, _ := c.MultipartForm()
	files := form.File["files"]
	for _, file := range files {
		//log.Println(file.Filename)
		//headers.Size 获取文件大小
		if file.Size > 1024*1024*50 {
			global.Response.Error(c, errcode.ERROR_SERVER, fmt.Sprintf("%v文件超过 1024*1024*50",file.Filename))
			return
		}
	}

	siteInfo := map[string]interface{}{}
	global.DB.Model(&models.CmsSite{}).Where("status != ?", 9).First(&siteInfo, siteId)
	if len(siteInfo) == 0 {
		global.Response.Error(c, errcode.ERROR_SERVER, fmt.Sprintf("站点信息不存在"))
		return
	}
	dir := "./resource/image/" + gconv.String(siteInfo["root"]) + "/"
	if pathUrl != "" {
		dir = dir + pathUrl + "/"
	}
	_, err := os.Stat(dir)
	if err != nil { //不存在创建
		_ = os.MkdirAll(path.Dir(dir), os.ModePerm)
		_, _ = os.Create(dir)
	}
	var newArr []map[string]interface{}
	for _, file := range files {
		filename := dir + file.Filename
		if err := c.SaveUploadedFile(file, filename); err != nil {
			global.Response.Error(c, errcode.ERROR_SERVER, fmt.Sprintf("保存文件失败"))
			return
		}
		newArr = append(newArr, map[string]interface{}{
			"id":  0,
			"url": pathUrl + "/" + file.Filename,
		})
	}
	//user_id
	userId := 0
	token := c.Request.Header.Get("token")
	tokenInfo, err := jwt.ParseJwtGoToken(token)
	if err == nil {
		userId = gconv.Int(tokenInfo.Id)
	}
	// 生成记录
	_, err = service.Audit().MakeAudit(newArr, 2, userId, gconv.Int(siteId))
	if err != nil {
		global.Response.Error(c, errcode.ERROR_SERVER, fmt.Sprintf("写入记录失败"))
		return
	}
	global.Response.Success(c, "上传成功")
}

// 获取目录文件
func (*CmsPictureController) GetFileList(c *gin.Context) {
	siteId := c.PostForm("site_id")
	path := c.DefaultPostForm("path", "")
	if siteId == "" {
		global.Response.Error(c, errcode.ERROR_SERVER, fmt.Sprintf("参数site_id不对"))
		return
	}
	result := map[string]interface{}{}
	global.DB.Model(&models.CmsSite{}).Where("status != ?", 9).First(&result, siteId)
	if len(result) == 0 {
		global.Response.Error(c, errcode.ERROR_SERVER, fmt.Sprintf("站点信息不存在"))
		return
	}
	dir := "./resource/image/" + gconv.String(result["root"])
	previewDir := gconv.String(result["preview_url"]) + "/resource/image/" + gconv.String(result["root"])
	onlineDir := gconv.String(result["online_url"]) + "/img"
	if path != "" {
		dir = dir + path
		previewDir = previewDir + path
		onlineDir = onlineDir + path
	}
	// 获取目录下所有文件
	list, err := helper.ReadDirFiles(dir)
	if err != nil {
		global.Response.Error(c, errcode.ERROR_SERVER, err.Error())
		return
	}
	for _, v := range list {
		v["preview_url"] = previewDir + "/" + gconv.String(v["label"])
		v["online_url"] = onlineDir + "/" + gconv.String(v["label"])
	}
	global.Response.Success(c,  list)
}

// 获取目录结构
func (*CmsPictureController) GetDirList(c *gin.Context) {
	siteId := c.PostForm("site_id")
	path := c.PostForm("path")
	if siteId == "" {
		global.Response.Error(c, errcode.ERROR_SERVER, fmt.Sprintf("参数site_id不对"))
		return
	}
	result := map[string]interface{}{}
	global.DB.Model(&models.CmsSite{}).Where("status != ?", 9).First(&result, siteId)
	if len(result) == 0 {
		global.Response.Error(c,errcode.ERROR_SERVER, fmt.Sprintf("站点信息不存在"))
		return
	}
	dir := "./resource/image/" + gconv.String(result["root"]) + "/" + path
	list, err := helper.ReadDirTree(dir, true, "")
	if err != nil {
		global.Response.Error(c, errcode.ERROR_SERVER, err.Error())
		return
	}
	global.Response.Success(c, list)
}

// 新增目录
func (*CmsPictureController) CreateDir(c *gin.Context) {
	siteId := c.PostForm("site_id")
	path := c.PostForm("path")
	isLevel := c.PostForm("is_level") // 1 同级  2 子集
	name := c.PostForm("name")
	if siteId == "" {
		global.Response.Error(c,  errcode.ERROR_SERVER, fmt.Sprintf("参数site_id不对"))
		return
	}
	result := map[string]interface{}{}
	global.DB.Model(&models.CmsSite{}).Where("status != ?", 9).First(&result, siteId)
	if len(result) == 0 {
		global.Response.Error(c, errcode.ERROR_SERVER, fmt.Sprintf("站点信息不存在"))
		return
	}
	dir := "./resource/image/" + gconv.String(result["root"]) + "/" + path
	if isLevel == "1" {
		lastIndex := strings.LastIndex(dir, "/")
		dir = dir[:lastIndex]
	}
	mkdirDir := dir + "/" + name
	err := helper.IsNotExistMkDir(mkdirDir)
	if err != nil {
		global.Response.Error(c, errcode.ERROR_SERVER, err.Error())
		return
	}
	global.Response.Success(c,  "创建成功")
}

func (*CmsPictureController) ChangeOnlinePic(c *gin.Context) {
	siteId := c.PostForm("site_id")
	url := c.PostForm("url")
	//_,file, err := c.Request.FormFile("file")
	if siteId == "" {
		global.Response.Error(c, errcode.ERROR_SERVER, "参数site_id不对")
		return
	}
	if url == "" {
		global.Response.Error(c, errcode.ERROR_SERVER,  fmt.Sprintf("参数url不正确"))
		return
	}

	siteInfo := map[string]interface{}{}
	global.DB.Model(&models.CmsSite{}).Where("status != ?", 9).First(&siteInfo, siteId)
	if len(siteInfo) == 0 {
		global.Response.Error(c, errcode.ERROR_SERVER, fmt.Sprintf("站点信息不存在"))
		return
	}
	// 	1 下载url到本地
	//imgPath := "/Users/user_service/Documents/home/www/mygin-proj/"
	thisPath, _ := os.Getwd()
	imgPath := thisPath + "/resource/image/" + gconv.String(siteInfo["root"]) + "/change_pic/"
	err := helper.MkDir(path.Dir(imgPath))
	if err != nil {
		global.Response.Error(c, errcode.ERROR_SERVER, fmt.Sprintf("创建路径失败"))
		return
	}

	fileName := path.Base(url)
	newFileName := "change_" + gconv.String(carbon.Now().Timestamp()) + "_" + fileName
	localUrl := imgPath + newFileName

	res, err := http.Get(url)
	if err != nil {
		global.Response.Error(c, errcode.ERROR_SERVER, fmt.Sprintf("获取图片失败"))
		return
	}
	defer res.Body.Close()
	// 获得get请求响应的reader对象
	reader := bufio.NewReaderSize(res.Body, 32*1024)
	file, err := os.Create(localUrl)
	if err != nil {
		panic(err)
	}
	// 获得文件的writer对象
	writer := bufio.NewWriter(file)
	written, _ := io.Copy(writer, reader)
	fmt.Printf("Total length: %d", written)

	// 2 本地url 上传 到 远方服务器
	sh := "rsync -avz --progress --password-file=" + gconv.String(siteInfo["rsync_password_path"]) + " " + localUrl +
		" " + gconv.String(siteInfo["rsync_image_address"]) + "/change_pic/ "
	SendFile(sh)

	// 3.返回远方服务器图片地址
	onlineUrl := gconv.String(siteInfo["online_url"])
	newUrl := onlineUrl + "/img/change_pic/" + newFileName
	global.Response.Success(c, newUrl)

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

type UploadResData struct {
	Url  string `json:"url"`
	Alt  string `json:"alt"`
	Href string `json:"href"`
}

func (*CmsPictureController) UploadAndPublish(c *gin.Context) {
	siteId := c.PostForm("site_id")
	pathUrl := c.DefaultPostForm("path", "")
	if siteId == "" {
		c.JSON(200, map[string]interface{}{
			"errno":   1,
			"message": "参数site_id不对",
		})
		return
	}
	form, _ := c.MultipartForm()
	files := form.File["files"]
	for _, file := range files {
		//log.Println(file.Filename)
		//headers.Size 获取文件大小
		if file.Size > 1024*1024*50 {
			c.JSON(200, map[string]interface{}{
				"errno":   1,
				"message": file.Filename + "文件超过 1024*1024*50",
			})
			return
		}
	}
	siteInfo := map[string]interface{}{}
	global.DB.Model(&models.CmsSite{}).Where("status != ?", 9).First(&siteInfo, siteId)
	if len(siteInfo) == 0 {
		c.JSON(200, map[string]interface{}{
			"errno":   1,
			"message": "站点信息不存在",
		})
		return
	}

	dir := "./resource/image/" + gconv.String(siteInfo["root"]) + "/"
	if pathUrl != "" {
		dir = dir + pathUrl + "/"
	}
	_, err := os.Stat(dir)
	if err != nil { //不存在创建
		_ = os.MkdirAll(path.Dir(dir), os.ModePerm)
		_, _ = os.Create(dir)
	}
	var newArr []map[string]interface{}
	var resFilename string
	for _, file := range files {
		filename := dir + file.Filename
		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.JSON(200, map[string]interface{}{
				"errno":   1,
				"message": "保存文件失败",
			})
			return
		}
		newArr = append(newArr, map[string]interface{}{
			"id":  0,
			"url": file.Filename,
		})
		resFilename = file.Filename
	}

	//user_id
	userId := 0
	token := c.Request.Header.Get("token")
	tokenInfo, err := jwt.ParseJwtGoToken(token)
	if err == nil {
		userId = gconv.Int(tokenInfo.Id)
	}
	// 生成记录
	auditId, err := service.Audit().MakeAudit(newArr, 2, userId, gconv.Int(siteId))
	if err != nil {
		c.JSON(200, map[string]interface{}{
			"errno":   1,
			"message": "写入记录失败" + err.Error(),
		})
		return
	}

	publishParam := service.PublishParam{
		Id:     gconv.Int(auditId),
		SiteId: gconv.Int(siteId),
	}
	err = service.CmsPublish().Publish(c, publishParam)
	if err != nil {
		c.JSON(200, map[string]interface{}{
			"errno":   1,
			"message": err.Error(),
			"err":     "publish",
		})
		return
	}
	onlineDir := gconv.String(siteInfo["online_url"]) + "/img/"
	c.JSON(200, map[string]interface{}{
		"errno":   0,
		"message": "上传成功",
		"data": UploadResData{
			Url:  onlineDir + resFilename,
			Alt:  "",
			Href: onlineDir + resFilename,
		},
	})
}
