package global

import "gin/internal/config"

var Cfg *config.Config

func InitConfig(conf *config.Config) (err error) {
	Cfg = conf
	return
}
