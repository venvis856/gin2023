package jwt

import (
	"encoding/json"
	"fmt"
	"gin/internal/common_config"
	"github.com/dgrijalva/jwt-go"
	"github.com/gogf/gf/util/gconv"
	"strings"
	"time"
)

// 声明一个标准的JWT接口
type IJwt interface {
	//设置头部
	SetHeader(string)
	//设置签名算法
	SetSignFunc(SignFunc)
	//设置编码算法
	SetEncodeFunc(EncodeFunc)

	//写入body
	WriteBody(map[string]interface{})

	//生成jwt
	CreateJwtString() (string, error)
	//验证jwt
	CheckJwtString(string) bool
}

// 规范header的格式
type Header struct {
	Type string `json:"type"`
	Alg  string `json:"alg"`
}

// 签名算法
type SignFunc func([]byte) string

// 编码算法
type EncodeFunc func([]byte) string

// 声明一个结构图 去实现 标准的JWT接口
type Jwt struct {
	Header Header
	Body   map[string]interface{}

	signFun   SignFunc
	encodeFun EncodeFunc
}

// 设置头部信息，说明你使用的签名算法
func (j *Jwt) SetHeader(headerType string) {
	j.Header = Header{
		Type: "JWT",
		Alg:  headerType,
	}
}

// 设置签名算法
func (j *Jwt) SetSignFunc(signFunc SignFunc) {
	j.signFun = signFunc
}

// 设置对 header 和 body 的加密算法
func (j *Jwt) SetEncodeFunc(encodeFunc EncodeFunc) {
	j.encodeFun = encodeFunc
}

// 写入要加密的内容
func (j *Jwt) WriteBody(body map[string]interface{}) {
	j.Body = body
}

// 生成token
func (j *Jwt) CreateJwtString() (string, error) {
	//编码header
	headerByte, err := json.Marshal(j.Header)
	if err != nil {
		return "", err
	}
	headerStr := j.encodeFun(headerByte)

	//编码body
	bodyByte, err := json.Marshal(j.Body)
	if err != nil {
		return "", err
	}
	bodyStr := j.encodeFun(bodyByte)

	//签名
	signByte := j.signFun([]byte(string(headerStr) + "." + string(bodyStr)))

	return fmt.Sprintf("%s.%s.%s", headerStr, bodyStr, signByte), nil
}

// 验证 token 是否合规
func (j *Jwt) CheckJwtString(input string) bool {
	arr := strings.Split(input, ".")
	//格式是否正确
	if len(arr) != 3 {
		return false
	}
	//签名
	signByte := j.signFun([]byte(string(arr[0]) + "." + string(arr[1])))
	if string(signByte) != arr[2] {
		return false
	}
	return true
}

// jwt-go
func CreateJwtGoToken(audience string, id string) (string, error) {
	expiresTime := time.Now().Unix() + gconv.Int64(common_config.Cfg.Login.UserTime)
	claims := jwt.StandardClaims{
		Audience:  audience,          // 受众
		ExpiresAt: expiresTime,       // 失效时间
		Id:        id,                // 编号
		IssuedAt:  time.Now().Unix(), // 签发时间
		Issuer:    "system",          // 签发人
		NotBefore: time.Now().Unix(), // 生效时间
		Subject:   "login",           // 主题
	}
	var jwtSecret = []byte(gconv.String(common_config.Cfg.Login.Secret))
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseJwtGoToken(token string) (*jwt.StandardClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(gconv.String(common_config.Cfg.Login.Secret)), nil
	})
	if err == nil && jwtToken != nil {
		if claim, ok := jwtToken.Claims.(*jwt.StandardClaims); ok && jwtToken.Valid {
			return claim, nil
		}
	}
	return nil, err
}
