package redis

import (
	"web_app/settings"

	"github.com/go-redis/redis"
	_ "github.com/go-redis/redis"
)

var Rdb *redis.Client

func Init(cfg settings.RedisConfig) (err error) {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     cfg.Host + ":" + cfg.Port,
		Password: cfg.PassWord,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})
	err = Rdb.Ping().Err()
	if err != nil {
		return
	}
	return nil
}
