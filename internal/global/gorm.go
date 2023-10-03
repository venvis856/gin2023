package global

import (
	"gin/internal/library/db"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitGorm(conf *db.Config) (err error) {
	DB, err = db.NewGormDB(conf)
	return
}
