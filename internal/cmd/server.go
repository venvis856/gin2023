package cmd

import (
	"gin/internal/common_config"
	"gin/internal/global"
	"github.com/zeebo/errs"
)

func InitServer(conf *common_config.Config) {
	errs := errs.Group{}
	errs.Add(
		global.InitConfig(conf),
		global.InitLogger(&conf.Log),
		global.InitFilesystem(&conf.Filesystem),
		global.InitGorm(&conf.DB),
		//global.InitRpc(),
		//initRedis(&conf.Redis),
	)
	if errs.Err() != nil {
		panic(errs.Err())
	}
}
