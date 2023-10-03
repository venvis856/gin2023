package helper

import (
	"bytes"
	"gin/internal/cmd/system/constraints"
	"math"
	"math/rand"
	"reflect"
	"sort"
	"time"

	"github.com/gogf/gf/util/gconv"
)

// InArray 判断value是否在array中
//
//	a 值   v 数组
func InArray(value interface{}, arr interface{}) bool {
	switch reflect.TypeOf(arr).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(arr)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(value, s.Index(i).Interface()) {
				return true
			}
		}
	}
	return false
}

// Contain 判断searchkey是否在target中，target支持的类型arrary,slice,map
func Contain(searchkey interface{}, target interface{}) bool {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == searchkey {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(searchkey)).IsValid() {
			return true
		}
	}
	return false
}

// 判断字符串是否在字符串数组里面 二分查找法
func InStringArray(target string, str_array []string) bool {
	sort.Strings(str_array)
	index := sort.SearchStrings(str_array, target)
	if index < len(str_array) && str_array[index] == target {
		return true
	}
	return false
}

// SortMap 对slice map排序
func SortMapSlice(data []map[string]interface{}, index string) []map[string]interface{} {
	flag := 0
	sort.Slice(data, func(i, j int) bool {
		if flag == 0 {
			t := reflect.ValueOf(data[i][index]).Kind()
			if t == reflect.String {
				flag = 1
			}
		}
		if flag == 1 {
			return gconv.String(data[i][index]) < gconv.String(data[j][index])
		}
		return gconv.Int(data[i][index]) < gconv.Int(data[j][index])
	})
	return data
}

// Sort 对slice map或者slice struct 排序（统一转成slice map去排序）
func SortSlice(data []interface{}, index string, sortType ...string) []interface{} {
	if len(sortType) == 0 {
		sortType = []string{"asc"}
	}
	flag := 0
	sort.Slice(data, func(i, j int) bool {
		if flag == 0 {
			t := reflect.ValueOf(gconv.Map(data[i])[index]).Kind()
			if t == reflect.String {
				flag = 1
			}
		}
		if flag == 1 {
			if sortType[0] == "desc" {
				return gconv.String(gconv.Map(data[i])[index]) > gconv.String(gconv.Map(data[j])[index])
			}
			return gconv.String(gconv.Map(data[i])[index]) < gconv.String(gconv.Map(data[j])[index])
		}
		if sortType[0] == "desc" {
			return gconv.Int(gconv.Map(data[i])[index]) > gconv.Int(gconv.Map(data[j])[index])
		}
		return gconv.Int(gconv.Map(data[i])[index]) < gconv.Int(gconv.Map(data[j])[index])
	})
	return data
}

// 将一维数组 转化 为 二维数组
func ArrayChunk(ss interface{}, size int) [][]interface{} {
	s := gconv.SliceAny(ss)
	if size < 1 {
		panic("size: cannot be less than 1")
	}
	length := len(s)
	chunks := int(math.Ceil(float64(length) / float64(size)))
	var n [][]interface{}
	for i, end := 0, 0; chunks > 0; chunks-- {
		end = (i + 1) * size
		if end > length {
			end = length
		}
		n = append(n, s[i*size:end])
		i++
	}
	return n
}

// RandSlice 随机取一个slice值
func RandSlice(dataAny interface{}) interface{} {
	data := gconv.SliceAny(dataAny)
	// n := rand.Int() % len(data)
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	n := rand.Intn(len(data))
	return data[n]
}

// Implode map转字符串
func MapStrStrImplode(glue string, pieces map[string]string) string {
	var buf bytes.Buffer
	l := len(pieces)
	for key, str := range pieces {
		buf.WriteString(key)
		buf.WriteString(":")
		buf.WriteString(str)
		if l--; l > 0 {
			buf.WriteString(glue)
		}
	}
	return buf.String()
}

// MapStrStrImplodeSort map转字符串并根据slice排序
func MapStrStrImplodeSort(glue string, pieces map[string]string, slice []string) string {
	var buf bytes.Buffer
	l := len(pieces)
	for _, Key := range slice {
		/* 按顺序从MAP中取值输出 */
		if str, ok := pieces[Key]; ok {
			buf.WriteString(Key)
			buf.WriteString(":")
			buf.WriteString(str)
			if l--; l > 0 {
				buf.WriteString(glue)
			}
		}
	}
	return buf.String()
}

// 多个map合并
func MergeMap(data map[string]interface{}, d ...map[string]interface{}) map[string]interface{} {
	if len(d) > 0 {
		for _, val := range d {
			for k, v := range val {
				data[k] = v
			}
		}
	}
	return data
}

// 多个  []map[string]interface{} 合并
func MergeSliceMap(data []map[string]interface{}, isOnly bool, d ...[]map[string]interface{}) []map[string]interface{} {
	if len(d) > 0 {
		for _, val := range d {
			if !isOnly {
				data = append(data, val...)
			} else {
				for _, valItem := range val {
					if !InArray(valItem, data) {
						data = append(data, valItem)
					}
				}
			}
		}
	}
	return data
}

// []map 转一维数组
func SliceMapToSliceString(data []map[string]interface{}, key string) []string {
	newData := make([]string, 0)
	if len(data) == 0 {
		return newData
	}
	for _, v := range data {
		newKey := ""
		if lKey, ok := v[key].(string); ok {
			newKey = lKey
		}
		newData = append(newData, newKey)
	}
	return newData
}

func Map[T any, K any](arr []T, fn func(item T, k int) K) []K {
	res := make([]K, len(arr))
	for k, v := range arr {
		res[k] = fn(v, k)
	}
	return res
}

func Foreach[T any](arr []T, fn func(item T, k int) error) error {
	for k, v := range arr {
		if err := fn(v, k); err != nil {
			return err
		}
	}
	return nil
}

func Reduce[T any, K any](arr []T, fn func(res *K, item T, k int)) K {
	var res K
	for k, v := range arr {
		fn(&res, v, k)
	}
	return res
}

func SliceMapToSliceInt64(data []map[string]interface{}, key string) []int64 {
	newData := make([]int64, 0)
	if len(data) == 0 {
		return newData
	}
	for _, v := range data {
		if lKey, ok := v[key].(int64); ok {
			newData = append(newData, lKey)
		}
	}
	return newData
}

// Intersect 查交集
func Intersect[T constraints.Ordered](arr1, arr2 []T) []T {
	ret := make([]T, 0)
	l1 := len(arr1)
	l2 := len(arr2)
	if l1 == 0 || l2 == 0 {
		return ret
	}
	if l1 > l2 {
		arr1, arr2 = arr2, arr1
		l1, l2 = l2, l1
	}
	mp := make(map[T]bool, l1)
	for _, s := range arr1 {
		if _, ok := mp[s]; !ok {
			mp[s] = true
		}
	}
	for _, s := range arr2 {
		if _, ok := mp[s]; ok {
			ret = append(ret, s)
		}
	}
	return ret
}

// Split 将arr 按每num个切割
func Split[T any](arr []T, num int64) [][]T {
	max := int64(len(arr))
	res := make([][]T, 0)
	if max == 0 {
		return res
	}
	if num == 0 || max <= num {
		res = append(res, arr)
		return res
	}
	var step = max / num
	if max%num > 0 {
		step += 1
	}
	var beg int64
	var end int64
	for i := int64(0); i < step; i++ {
		beg = i * num
		end = beg + num
		if end > max {
			end = max
		}
		res = append(res, arr[beg:end])
	}
	return res
}
