package helper

import (
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/grand"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// CreateId 生成一个24位的id
// 第一个参数是节点0~99
// 第二个参数是生成位数，大于等于20
func CreateId(node ...int) string {
	n := 0
	if len(node) > 0 {
		n = node[0]
	}
	str := gtime.Now().Format("ymdHis") + strconv.Itoa(n)

	clen := len(str)
	max := 24
	if len(node) > 1 && node[1] >= 20 {
		max = node[1]
	}
	maxRand, _ := strconv.Atoi(gstr.Repeat("9", max-clen))
	mixRand, _ := strconv.Atoi("1" + gstr.Repeat("0", max-clen-1))
	randNum := grand.N(mixRand, maxRand)

	str += strconv.Itoa(randNum)
	return str
}

// 获取 min - max 的一个随机数
func GetRangeRandNum(max int, min int) int {
	return rand.Intn(max-min) + min
}

// IsNumeric 判断是否数字
func IsNumeric(val interface{}) bool {
	switch val.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
	case float32, float64, complex64, complex128:
		return true
	case string:
		str := val.(string)
		if str == "" {
			return false
		}
		// Trim any whitespace
		str = strings.Trim(str, " \\t\\n\\r\\v\\f")
		if str[0] == '-' || str[0] == '+' {
			if len(str) == 1 {
				return false
			}
			str = str[1:]
		}
		// hex
		if len(str) > 2 && str[0] == '0' && (str[1] == 'x' || str[1] == 'X') {
			for _, h := range str[2:] {
				if !((h >= '0' && h <= '9') || (h >= 'a' && h <= 'f') || (h >= 'A' && h <= 'F')) {
					return false
				}
			}
			return true
		}
		// 0-9,Point,Scientific
		p, s, l := 0, 0, len(str)
		for i, v := range str {
			if v == '.' { // Point
				if p > 0 || s > 0 || i+1 == l {
					return false
				}
				p = i
			} else if v == 'e' || v == 'E' { // Scientific
				if i == 0 || s > 0 || i+1 == l {
					return false
				}
				s = i
			} else if v < '0' || v > '9' {
				return false
			}
		}
		return true
	}

	return false
}

//获取随机数
func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
