package global

import (
	log2 "gin/internal/library/logger"
)

var Logger *log2.Logger

var Log *log2.Log

func InitLogger(conf *log2.Config) (err error) {
	Log, err = log2.NewLog(conf)
	Logger, err = log2.NewLogger(*conf)
	return
}
