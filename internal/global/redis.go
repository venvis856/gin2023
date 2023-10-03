package global

import (
	"gin/internal/library/redis"
)

var Redis *redis.Redis

func InitRedis(conf *redis.Config) (err error) {
	Redis, err = redis.NewRedis(conf)
	return
}
