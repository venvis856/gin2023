package helper

import (
	"encoding/base64"
	"errors"
	"regexp"
)

// checkBase64 检测是否base64图片
//不能判断一定是，可以判断一定不是。判断方式，base64只包含特定字符;解码再转码，查验是否相等。目前貌似没有能一定判断是的方法
func CheckBase64(str string) (bool, error) {
	pattern := "^([A-Za-z0-9+/]{4})*([A-Za-z0-9+/]{4}|[A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{2}==)$"
	matched, err := regexp.MatchString(pattern, str)
	if err != nil {
		return false, errors.New("检测base64字符不匹配")
	}
	if !(len(str)%4 == 0 && matched) {
		return false, errors.New("检测base64除4有余数,固格式不正确")
	}
	unCodeStr, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return false, errors.New("检测base64的时候,base64解码失败")
	}
	tranStr := base64.StdEncoding.EncodeToString(unCodeStr)
	//return str==base64.StdEncoding.EncodeToString(unCodeStr)
	if str == tranStr {
		return true, nil
	}
	return false, errors.New("检测base64的时候,发现不匹配")
}
