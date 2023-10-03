package helper

import (
	"errors"
	"reflect"
)

// Call 调用方法
func Call(method interface{}, params ...interface{}) ([]reflect.Value, error) {

	if reflect.TypeOf(method).Kind() != reflect.Func {
		return nil, errors.New("the name of input not func!")
	}

	f := reflect.ValueOf(method)
	if len(params) != f.Type().NumIn() {
		return nil, errors.New("the number of input params not match!")
	}
	mp := []reflect.Value{}
	for _, v := range params {
		mp = append(mp, reflect.ValueOf(v))
	}
	return f.Call(mp), nil
}
