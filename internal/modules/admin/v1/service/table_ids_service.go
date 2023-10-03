package service

import "gin/internal/library/handlePanic"

type TableIdsInterface interface {
	GetAddId(tableName string, identifyId int64) int64
}

var tableIds TableIdsInterface

func TableIds() TableIdsInterface {
	if tableIds == nil {
		handlePanic.Panic("table ids service panic")
	}
	return tableIds
}

func RegisterTableIds(i TableIdsInterface) {
	tableIds = i
}
