package helper

import (
	"encoding/json"
	"sort"
)

/**
map key为string
排序后遍历执行回调函数
*/

func MapSortStringFor(data map[string]interface{}, callback func(string, interface{})) {
	var arr []string
	for k, _ := range data {
		arr = append(arr, k)
	}
	sort.Strings(arr)
	for _, v := range arr {
		callback(v, data[v])
	}
}

/**
map key为int
排序后遍历执行回调函数
*/

func MapSortIntFor(data map[int]interface{}, callback func(int, interface{})) {
	var arr []int
	for k, _ := range data {
		arr = append(arr, k)
	}
	sort.Ints(arr)
	for _, v := range arr {
		callback(v, data[v])
	}
}

/*
*
map 转化为 slice
*/
func MapToSlice(m map[int]string) []string {
	s := make([]string, 0, len(m))
	for _, v := range m {
		s = append(s, v)
	}
	return s
}

/*
*
map 转化为 slice
*/
func MapSliceToSlice(m []map[string]interface{}, key string) []interface{} {
	s := make([]interface{}, 0, len(m))
	for _, v := range m {
		s = append(s, v[key])
	}
	return s
}

// map 转 struct
func MapToStruct(dataMap interface{}, outStruct interface{}) error {
	b, err := json.Marshal(dataMap)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, &outStruct)
	if err != nil {
		return err
	}

	return err
}
