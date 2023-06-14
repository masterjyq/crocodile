package rediscli

import (
	"github.com/go-redis/redis"
	"github.com/labulaka521/crocodile/core/config"
)

var (
	_redis *redis.Client
)

func GetRedisClient() *redis.Client {
	return _redis
}

func NewRedis(opt *redis.Options) (err error) {
	_redis = redis.NewClient(&redis.Options{
		Addr:     config.CoreConf.Server.Redis.Addr,
		Password: config.CoreConf.Server.Redis.PassWord,
		DB:       config.CoreConf.Server.Redis.DB,
	})
	err = _redis.Ping().Err()
	return
}
