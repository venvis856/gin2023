package db

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OrderBy struct {
	Field     string `json:"order_field" form:"order_field" query:"order_field"`
	OrderDesc string `json:"order_by" form:"order_by" query:"order_by"`
}

func (o OrderBy) OrderByQuery() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if o.Field == "" {
			return db
		}
		return db.Order(clause.OrderByColumn{Column: clause.Column{Name: o.Field}, Desc: o.OrderDesc == "desc"})
	}
}
