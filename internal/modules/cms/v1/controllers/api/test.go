package api

import (
	"fmt"
	"gin/app/models"
	"gin/global"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gconv"
	"time"
)

type TestController struct{}

func (*TestController) Kvrocks(c *gin.Context) {
	fmt.Println(c.Query("key"))
	_ = global.Kvrocks.Set("key", c.Query("key"), time.Second*10000)
	data, _ := global.Kvrocks.Get("key")
	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", data)
}

func (*TestController) Tmp(c *gin.Context) {
	arr := []int{1, 2, 3, 4, 5}
	jsonData, err := global.Json.Marshal(arr)
	if err != nil {
		fmt.Println(err, "====err")
	}
	//rs := helper.InArray(a, arr)
	global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, "", string(jsonData))
}

func (*TestController) TestFirst(c *gin.Context) {
	// 可以
	result := map[string]interface{}{}
	global.DB.Model(&models.User{}).First(&result)
	//global.DB.Table("user").Get(&result)
	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", result)
}

func (*TestController) TestAll(c *gin.Context) {
	// 可以
	//result := map[string]interface{}{}
	//var result []map[interface{}]interface{}
	var result []map[string]interface{}
	global.DB.Model(&models.User{}).Find(&result)

	fmt.Println(result, 11222)
	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", result)
}

func (*TestController) Test1(c *gin.Context) {
	//file:="log/gin"+gtime.Now().Format("Ymd")+".log"
	//fmt.Println(file,1111)
	//f, _ := os.Create(file)
	//gin.DefaultWriter = io.MultiWriter(f)
	a := make(map[string]string)
	a["b"] = "ccccccccc"
	a["b2"] = "33"
	global.Logger.Write("aaaa/bbb", "error", a, "1111", "bbbb")
	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", "hello world")
}

func (*TestController) TestRedis(c *gin.Context) {
	a := make(map[string]string)
	a["bb"] = "bbbbbbb"
	err := global.Redis.Set("cc", a, 10*time.Second)
	if err != nil {
		fmt.Println("设置失败", err)
	}

	rs, err := global.Redis.Get("cc")
	if err != nil {
		fmt.Println("获取失败:", err)
	}
	fmt.Println("获取成功:", rs)
}

type DataParam struct {
	Uid        string `form:"uid" json:"uid"`
	Msg        string `form:"msg" json:"msg"`
	CreateTime string `form:"msg"`
}

func (*TestController) TestWriteLog(c *gin.Context) {
	var data DataParam
	if err := c.ShouldBind(&data); err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
		return
	}
	if data.Uid == "" || gconv.String(data.Msg) == "" {
		return
	}
	global.Logger.Write(data.Uid, "info", data.Msg)
}

func (*TestController) TestGinTemplate(c *gin.Context) {
	//service := new(services.CmsMakeServer)
	//content := service.ReplaceContent("{{.ts}}")
	//global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", content)
}
