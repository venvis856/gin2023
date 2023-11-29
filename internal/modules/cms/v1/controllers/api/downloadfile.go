package api

import (
	"fmt"
	"gin/global"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

type DownloadController struct{}

func (*DownloadController) Download(c *gin.Context) {

	//file:="log/gin"+gtime.Now().Format("Ymd")+".log"
	//fmt.Println(file,1111)
	//f, _ := os.Create(file)
	//gin.DefaultWriter = io.MultiWriter(f)
	var param struct {
		Name string `form:"name" json:"name" `
		Url  string `form:"url" json:"url" `
	}
	if err := c.ShouldBind(&param); err != nil {
		global.Response.Json(c, global.HTTP_SUCCESS, global.ERROR, err.Error(), "")
		return
	}
	url := "https://unsplash.com/photos/cMvn8bXwYcU/download?ixid=MnwxMjA3fDB8MXxhbGx8ODN8fHx8fHwyfHwxNjU0MTM3MDE3&force=true"

	c.Redirect(301,url)
	return




	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprint(os.Stderr, "get url error", err)
	}
	//
	//
	//defer resp.Body.Close()
	//file:="file/"+gconv.String(carbon.Now().Timestamp())+"/"+param.Name
	//out, err := os.Create(file)
	//wt :=bufio.NewWriter(out)
	//
	//defer out.Close()
	// 缓冲区32k循环操作
	//c.Header("content-disposition", "attachment;filename=aa.mp4")
	n, err := io.Copy(c.Writer, resp.Body)
	fmt.Println("write", n)
	if err != nil {
		panic(err)
	}

	//wt.Flush()
	//
	//
	//c.File(file)

	//ghttp.Response.ServeFileDownload(url,"21212.mp4")
}
