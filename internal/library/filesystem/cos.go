package filesystem

import (
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/zeebo/errs"
	"net/http"
	"net/url"
	"os"
)

var ErrFilesystemCos = errs.Class("cos")

type CosConfig struct {
	AccessKeyId     string `help:"AccessKeyId"  default:"AKIDTklkV1WEVbPySgvhnCMIJMeijUBbf93h"`
	AccessKeySecret string `help:"AccessKeySecret"  default:"W8UouVquhDf7KGvJQ9cS2LN8oDct9wSu"`
	Bucket          string `help:"存储桶" default:"nf-1302123357"`
	Region          string `help:"区域" default:"ap-shanghai"`
	BaseUrl         string `help:"cos地址" default:"https://nf-1302123357.cos.ap-shanghai.myqcloud.com"`
	Url             string `help:"自定义地址"  default:"https://img.nextstorage.cn/"`
}

type Cos struct {
	config   *CosConfig
	instance *cos.Client
}

func NewCos(config CosConfig) (*Cos, error) {
	u, err := url.Parse(config.BaseUrl)
	if err != nil {
		return nil, fmt.Errorf("cos init error: %+v", err)
	}
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  config.AccessKeyId,
			SecretKey: config.AccessKeySecret,
		},
	})
	return &Cos{config: &config, instance: client}, nil
}

func (c *Cos) PutFile(ctx context.Context, dist string, src *os.File) error {
	_, _, err := c.instance.Object.Upload(
		ctx, dist, src.Name(), nil,
	)
	return ErrFilesystemCos.Wrap(err)
}

func (c *Cos) Put(ctx context.Context, dist string, src string) error {
	_, _, err := c.instance.Object.Upload(
		ctx, dist, src, nil,
	)
	return ErrFilesystemCos.Wrap(err)
}

func (c *Cos) Url(fileName string) string {
	objectUrl := c.instance.Object.GetObjectURL(fileName)
	return objectUrl.String()
}

func (c *Cos) Exists(ctx context.Context, file string) bool {
	ok, err := c.instance.Object.IsExist(ctx, file)
	if err != nil {
		return false
	}
	return ok
}
