package global

import "gin/internal/common_config"

var Cfg *common_config.Config

func InitConfig(conf *common_config.Config) (err error) {
	Cfg = conf
	return
}
