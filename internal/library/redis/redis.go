package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/gogf/gf/util/gconv"
	"time"
)

type Config struct {
	Address  string `help:"redis地址" devDefault:"192.168.43.26:6379" default:"127.0.0.1:6379"`
	Password string `help:"redis密码" default:""`
	DB       int    `help:"redis数据库" default:"0"`
}

type Redis struct {
	rdb *redis.Client
}

var ctx = context.Background()

func NewRedis(conf *Config) (*Redis, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.Address,
		Password: conf.Password, // no password set
		DB:       conf.DB,       // use default DB
	})
	err := rdb.Set(ctx, "redis_init", "redis_start", time.Second*60*5).Err()
	if err != nil {
		fmt.Println("redis连接失败：" + err.Error())
		return nil, err
	}
	fmt.Println("redis连接成功", err)
	return &Redis{
		rdb: rdb,
	}, nil
}

func (a *Redis) Set(key string, value interface{}, time time.Duration) error {
	if !a.IsInit() {
		return errors.New("redis未初始化")
	}
	jsonValue, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return a.rdb.Set(ctx, key, jsonValue, time).Err()
}

func (a *Redis) Get(key string) (interface{}, error) {
	if !a.IsInit() {
		return nil, errors.New("redis未初始化")
	}
	byteValue, err := a.rdb.Get(ctx, key).Bytes()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	var value interface{}
	err = json.Unmarshal(byteValue, &value)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (a *Redis) HSet(key string, values interface{}) error {
	if !a.IsInit() {
		return errors.New("redis未初始化")
	}
	return a.rdb.HSet(ctx, key, values).Err()
}

func (a *Redis) HGet(key string, field string) (interface{}, error) {
	if !a.IsInit() {
		return nil, errors.New("redis未初始化")
	}
	val, err := a.rdb.HGet(ctx, key, field).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return val, nil
}

func (a *Redis) HGetAll(key string) (interface{}, error) {
	if !a.IsInit() {
		return nil, errors.New("redis未初始化")
	}
	val, err := a.rdb.HGetAll(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return val, nil
}

func (a *Redis) del(key string) error {
	if !a.IsInit() {
		return errors.New("redis未初始化")
	}
	a.rdb.Del(ctx, key)
	return nil
}

func (a *Redis) LPush(key string, values ...interface{}) error {
	if !a.IsInit() {
		return errors.New("redis未初始化")
	}
	if len(values) > 0 {
		for i := 0; i < len(values); i++ {
			jsonValue, err := json.Marshal(values[i])
			if err != nil {
				return err
			}
			err = a.rdb.LPush(ctx, key, jsonValue).Err()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (a *Redis) LPop(key string) (interface{}, error) {
	if !a.IsInit() {
		return nil, errors.New("redis未初始化")
	}
	str, err := a.rdb.LPop(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	var rs interface{}
	err = json.Unmarshal(gconv.Bytes(str), &rs)
	if err != nil {
		return nil, err
	}
	return rs, nil
}

func (a *Redis) LLen(key string) (int64, error) {
	if !a.IsInit() {
		return 0, errors.New("redis未初始化")
	}
	length, err := a.rdb.LLen(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	return length, nil
}

func (a *Redis) IsInit() bool {
	return a.rdb != nil
}
