package config

import (
	"gin/internal/library/db"
	"gin/internal/library/filesystem"
	"gin/internal/library/logger"
	"gin/internal/library/redis"
)

var Cfg *Config

type Config struct {
	Login struct {
		Key      string `help:"key"  default:"h85HsAaa"`
		UserTime int    `help:"UserTime" devDefault:"432000"  default:"4320000"`
		Secret   string `help:"key" default:"Sadjsadfasdhj"`
	}
	Api struct {
		Server string `help:"监听地址" devEnv:":8101" testEnv:":8401" devDefault:":8101" testDefault:":8101" default:":8181" `
	}
	Filesystem filesystem.Config
	DB         db.Config
	Redis      redis.Config
	Log        logger.Config
	Upload     struct {
		MaxSize int64 `help:"上传文件最大大小" default:"32000000"`
	}
}

func InitConfig(conf *Config) (err error) {
	Cfg = conf
	return
}
