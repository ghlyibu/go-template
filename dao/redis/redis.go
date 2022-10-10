package redis

import (
	"fmt"
	"go-template/setting"

	"github.com/go-redis/redis"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

// Init 初始化连接
func Init(cfg *setting.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password, // no password set
		DB:       cfg.DB,       // use default DB
		PoolSize: cfg.MinIdleConns,
	})

	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func Close() {
	_ = rdb.Close()
}
