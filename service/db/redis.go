package db


import (
	"cifs/service/config"
	"fmt"
	"github.com/go-redis/redis"
)

var (
	RedisClient *RedisDriver
)

func init() {
	if RedisClient == nil {
		Config := config.NewConfig().LoadConfig("../config/config.json")
		RedisClient = NewRedis(Config)
	}
}

type RedisDriver struct {
	Client *redis.Client
}

func NewRedis(conf *config.Config) *RedisDriver {
	addr := fmt.Sprintf("%s%d", conf.Data.Redis.Host, conf.Data.Redis.Port)
	return &RedisDriver{
		Client:redis.NewClient(&redis.Options{
			Network: "tcp",
			Addr: addr,}),
	}
}
