package api

import (
	"gin/app/models"
	"gin/global"
	"github.com/gin-gonic/gin"
)

//CREATE TABLE `rrweb` (   `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '数据库ID',    `uid` varchar(100) NOT NULL COMMENT 'uid',   `msg` mediumtext  NOT NULL COMMENT 'msg',   `create_time` int(11) DEFAULT NULL COMMENT '用户创建时间',    PRIMARY KEY (`id`) USING BTREE ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

type RrwebController struct {
	Uid       string `form:"uid" json:"uid"`
	Msg       string `form:"msg" json:"msg"`
	StartTime int    `form:"start_time" json:"start_time"`
	EndTime   int    `form:"end_time" json:"end_time"`
}

func (*RrwebController) Add(c *gin.Context) {
	var data RrwebController
	if err := c.ShouldBind(&data); err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
		return
	}
	if data.Uid == "" || data.Msg == "" || data.StartTime == 0 || data.EndTime == 0 {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, "参数错误", "")
		return
	}
	model := global.DB.Model(&models.Rrweb{})
	rrweb := models.Rrweb{
		Uid:       data.Uid,
		Msg:       data.Msg,
		StartTime: data.StartTime,
		EndTime:   data.EndTime,
	}

	result := model.Create(&rrweb) // 将数据指针传递给 Create
	if result.Error != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, "创建失败", "")
		return
	}

	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", "新增成功")
}

func (a *RrwebController) Get(c *gin.Context) {
	uid := c.DefaultPostForm("uid", "")
	time := c.PostForm("time")
	if uid == "" || time == "" {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, "参数错误", "")
	}
	model := global.DB.Model(&models.Rrweb{})
	result := map[string]interface{}{}
	model.Where("uid = ?", uid).Where("start_time <= ? AND end_time >= ?", time, time).First(&result)
	global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", result)
}
