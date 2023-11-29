package api

import (
	"fmt"
	"gin/global"
	"github.com/gin-gonic/gin"
	"io"
	"net"
	"net/http"
	"time"
)

func Hello(c *gin.Context) {
	global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, "", "hello world")
	fmt.Println("cccc")
}

func Hello2(c *gin.Context) {
	//file:="log/gin"+gtime.Now().Format("Ymd")+".log"
	//fmt.Println(file,1111)
	//f, _ := os.Create(file)
	//gin.DefaultWriter = io.MultiWriter(f)
	//a := make(map[string]string)
	//a["b"] = "ccccccccc"
	//a["b2"] = "33"
	//global.Logger.Write("aaaa/bbb", "error", a, "1111", "bbbb")
	//global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", "hello world")
	url := "http://kfpaydata.oa.com/composer.json"
	//global.Response.Json(c, global.HTTP_SUCCESS, global.SUCCESS, "", url)
	timeout := time.Duration(5) * time.Second
	transport := &http.Transport{
		ResponseHeaderTimeout: timeout,
		Dial: func(network, addr string) (net.Conn, error) {
			return net.DialTimeout(network, addr, timeout)
		},
		DisableKeepAlives: true,
	}
	client := &http.Client{
		Transport: transport,
	}

	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	c.Writer.Header().Set("Content-Disposition", "attachment; filename=111.txt")
	c.Writer.Header().Set("Content-Type", c.Request.Header.Get("Content-Type"))
	c.Writer.Header().Set("Content-Length", c.Request.Header.Get("Content-Length"))

	//stream the body to the client without fully loading it into memory
	io.Copy(c.Writer, resp.Body)
}
