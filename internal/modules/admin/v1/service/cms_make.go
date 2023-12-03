package service

import (
	"gin/internal/library/handlePanic"
)

type MakeInterface interface {
	Make(makeType string, pageIds []int, siteId int, isPreview bool) ([]map[string]interface{}, string, error)
	MakePage(pageIds []int, siteId int, isPreview bool) ([]map[string]interface{}, string, error)
}

var makeObj MakeInterface

func CmsMake() MakeInterface {
	if makeObj == nil {
		handlePanic.Panic("audit service panic")
	}
	return makeObj
}

func RegisterCmsMake(i MakeInterface) {
	makeObj = i
}
