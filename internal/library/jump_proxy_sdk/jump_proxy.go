package jump_proxy_sdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gin/internal/global"
	"github.com/google/go-querystring/query"
	"github.com/zeebo/errs"
	"go.uber.org/zap"
	"io"
	"net/http"
	"strings"
	"time"
)

var ErrJumpProxy = errs.Class("jumpProxy")

type Config struct {
	BaseUrl string `help:"盒子心跳代理服务地址" devDefault:"http://192.168.43.26:8401" testDefault:"http://192.168.43.26:8401" default:"http://192.168.43.26:8401"`
}

type JumpProxyClient struct {
	baseUrl string
	timeout time.Duration
}

type JumpProxyRequest interface {
	api() string
}

type JumpProxyResponse interface {
	Success() bool
	ErrMsg() string
	Error() error
	setRaw(r []byte)
}

type ResCommon struct {
	raw  []byte `json:"-"`
	Code int    `json:"code"`
	Msg  string `json:"msg,omitempty"`
}

func (receiver *ResCommon) Success() bool {
	return receiver.Code == 0
}

func (receiver *ResCommon) ErrMsg() string {
	return receiver.Msg
}

func (receiver *ResCommon) Error() error {
	if receiver.Success() {
		return nil
	}
	return errors.New(receiver.Msg)
}

func (receiver *ResCommon) String() string {
	return string(receiver.raw)
}

func (receiver *ResCommon) setRaw(r []byte) {
	receiver.raw = r
}

func NewJumpProxyClient(conf *Config) *JumpProxyClient {
	return &JumpProxyClient{
		baseUrl: strings.TrimSuffix(conf.BaseUrl, "/"),
		timeout: time.Second * 60,
	}
}

func (c *JumpProxyClient) Request(req JumpProxyRequest, res JumpProxyResponse) error {
	return c.request(req, res)
}

func (c *JumpProxyClient) RequestRaw(req JumpProxyRequest) (res []byte, err error) {
	res, err = c.post(req)
	if err != nil {
		err = ErrJumpProxy.Wrap(err)
	}
	return
}

func (c *JumpProxyClient) request(req JumpProxyRequest, res JumpProxyResponse) error {
	body, err := c.post(req)
	defer func() {
		_p, _ := json.Marshal(req)
		if err != nil {
			global.Log.Error(req.api(), zap.ByteString("req", _p), zap.ByteString("res", body), zap.Error(err))
		} else {
			global.Log.Info(req.api(), zap.ByteString("req", _p), zap.ByteString("res", body), zap.Error(err))
		}
	}()
	if err != nil {
		return ErrJumpProxy.Wrap(err)
	}
	err = json.Unmarshal(body, &res)
	res.setRaw(body)
	return ErrJumpProxy.Wrap(err)
}

func (c *JumpProxyClient) post(req JumpProxyRequest) (body []byte, err error) {
	httpClient := http.Client{Timeout: c.timeout}
	params, err := json.Marshal(req)
	if err != nil {
		return
	}
	resp, err := httpClient.Post(c.getApiUrl(req.api()), "application/json", bytes.NewReader(params))
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("request(POST) httpCode error:[%d]", resp.StatusCode)
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	return io.ReadAll(resp.Body)
}

func (c *JumpProxyClient) get(req JumpProxyRequest) (body []byte, err error) {
	httpClient := http.Client{Timeout: c.timeout}
	val, err := query.Values(req)
	if err != nil {
		return
	}
	resp, err := httpClient.Get(c.getApiUrl(req.api()) + "?" + val.Encode())
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("request(GET) httpCode error:[%d]", resp.StatusCode)
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	return io.ReadAll(resp.Body)
}

func (c *JumpProxyClient) getApiUrl(api string) string {
	return c.baseUrl + api
}
