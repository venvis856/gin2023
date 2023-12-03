package service

import (
	"gin/internal/library/handlePanic"
)

type AuditInterface interface {
	MakeAudit(auditDetailList []map[string]interface{}, makeType int, makeUserId int, siteId int) (int64, error)
	UpdatePageFirstMakeTime(pageIds []int) error
}

var auditObj AuditInterface

func Audit() AuditInterface {
	if auditObj == nil {
		handlePanic.Panic("audit service panic")
	}
	return auditObj
}

func RegisterAudit(i AuditInterface) {
	auditObj = i
}
