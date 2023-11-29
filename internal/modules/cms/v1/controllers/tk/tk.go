package tk

import (
  "gin/global"
  "github.com/gin-gonic/gin"
  "gin/app/library/topsdk"
)

type TkController struct {}

const (
  AppKey="27686241"
  AppSecrect="61992f9b4fb2e158159e82c4de09dc6b"
  AdZoneid="1294800161"
)

func (a TkController) GetItem(c *gin.Context)  {
  topClient := topsdk.NewTopClient(AppKey, AppSecrect)
  param:=map[string]interface{}{
    "Me":"",
    "adzone_id":AdZoneid,
  }
  resp, err := topClient.Request(param)
  if err!=nil{
    global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
    return
  }
  global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", resp)
}