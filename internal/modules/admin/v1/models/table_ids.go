package models

type TableIds struct {
	ID         int64  `json:"id" gorm:"id"`
	TableName  string `json:"table_name" gorm:"table_name"`
	IdentifyId int64  `json:"identify_id" gorm:"identify_id"`
	MaxId      int64  `json:"max_id" gorm:"max_id"`
}
