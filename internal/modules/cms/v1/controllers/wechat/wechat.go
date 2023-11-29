package wechat

import (
  "crypto/sha1"
  "encoding/xml"
  "fmt"
  "gin/global"
  "github.com/gin-gonic/gin"
  "io"
  "io/ioutil"
  "sort"
  "strings"
  "time"
)

type WechatController struct{}

const (
  Token          = "venvis7788"
  AppID          = "AppIDwxeb70a9b974631105"
  AppSecret      = "afde5b5e1810460c56cb82b2b72a1370"
  EncodingAESKey = "aKdJkABQMtkRO8qH4XdmrQcJ94pwBaPLa8hHtAwgms8"
)

// 绑定服务器时微信公众号调用指向该接口
// http://wechat.4itool.com/wechat/check?signature=eaba550f94c94682d9ada82813e1b71b9282fc6f&timestamp=2222&nonce=3333&echostr=success
func (*WechatController) Check(c *gin.Context) {
  var param struct {
    Signature string `form:"signature" json:"signature"`
    Timestamp string `form:"timestamp" json:"timestamp"`
    Nonce     string `form:"nonce" json:"nonce"`
    Echostr   string `form:"echostr" json:"echostr"`
  }
  fmt.Println("===param:", param)
  if err := c.ShouldBind(&param); err != nil {
    global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
    return
  }
  arr := []string{Token, param.Timestamp, param.Nonce}
  sort.Strings(arr)
  str := strings.Join(arr, "")
  h := sha1.New()
  _, _ = io.WriteString(h, str)
  result := fmt.Sprintf("%x", h.Sum(nil))
  rs := ""
  if result == param.Signature {
    rs = param.Echostr
  }
  fmt.Println(param, result, "===========", rs)
  c.String(200, "%v", rs)
}

/*
接收消息来处理函数
<xml>
  <ToUserName><![CDATA[gh_197b808143aa]]></ToUserName>
  <FromUserName><![CDATA[oFxwI6pOlOx238c1SnbsT2C24eF4]]></FromUserName>
  <CreateTime>1660640585</CreateTime>
  <MsgType><![CDATA[text]]></MsgType>
  <Content><![CDATA[你好啊]]></Content>
  <MsgId>23774656060202935</MsgId>
</xml>
*/

type XMLServers struct {
  ToUserName   string `xml:"ToUserName"`
  FromUserName string `xml:"FromUserName"`
  CreateTime   int    `xml:"CreateTime"`
  MsgType      string `xml:"MsgType"`
  Content      string `xml:"Content"`
  MsgId        int    `xml:"MsgId"`
}

type ReceiveXMLServers struct {
  XMLName      xml.Name `xml:"xml"`
  ToUserName   string   `xml:"ToUserName"`
  FromUserName string   `xml:"FromUserName"`
  CreateTime   int64    `xml:"CreateTime"`
  MsgType      string   `xml:"MsgType"`
  Content      string   `xml:"Content"`
}

// http://wechat.4itool.com/wechat/check?
func (*WechatController) Receive(c *gin.Context) {
  data, _ := ioutil.ReadAll(c.Request.Body)
  fmt.Println("===param:", string(data))
  xmlData := XMLServers{}
  err := xml.Unmarshal(data, &xmlData)
  if err != nil {
    fmt.Println(err, "=======xml err")
    return
  }
  receiveData := &ReceiveXMLServers{
    ToUserName:   xmlData.FromUserName,
    FromUserName: xmlData.ToUserName,
    CreateTime:   time.Now().Unix(),
    MsgType:      xmlData.MsgType,
    Content:      "",
  }

  fmt.Println(xmlData, xmlData.Content, "=====xml")
  switch xmlData.Content {
  case "idea":
    receiveData.Content = GetContent()
    c.XML(200, receiveData)
  default:
    c.String(200, "%v", "success")
  }

}

func GetContent() string {
  url, err := ioutil.ReadFile("/home/www/python_selenium/write_data.txt")
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(string(url))
  var content string = "1.无限重置Jetbrains系列产品30天试用期的方法：\n" +
    "https://mp.weixin.qq.com/s?__biz=MzkyNjQwMDAyMw==&mid=2247483684&idx=1&sn=701771125ed160a27f7429b9232ad48a&chksm=c236aa5cf541234a278c1899f92a052aaebe0fc8d1e528524a1ddaf93abb50f1c185c319ac90&token=1081398082&lang=zh_CN#rd\n" +
    "此方法只适用2021.2及以下版本\n\n" +
    "2.支持最新版本Jetbrains系列所有产品\n" +
    "前往以下地址下载 jetbra.zip 并选择对应产品激活码\n" +
    "地址：\n" +
    string(url) + "\n" +
    "使用方式：\n" +
    "1)  mac或linix电脑解压后直接运行cripts/install.sh脚本" +
    "\n    window请运行 scripts\\install-current-user.vbs (只为window当前用户安装) 或 scripts\\install-all-users.vbs (为所有用户安装)\n" +
    "2)  关闭Jetbrains，前往上面地址获取对应产品激活码\n    " +
    "重新打开Jetbrains输入激活码激活\n"
  return content
}
