package cmd

import (
	"gin/internal/config"
	"gin/internal/global"
	"github.com/zeebo/errs"
)

func InitServer(conf *config.Config) {
	errs := errs.Group{}
	errs.Add(
		config.InitConfig(conf),
		global.InitLogger(&conf.Log),
		global.InitFilesystem(&conf.Filesystem),
		//global.InitGorm(&conf.DB),
		//global.InitRpc(),
		//initRedis(&conf.Redis),
	)
	if errs.Err() != nil {
		panic(errs.Err())
	}
}
