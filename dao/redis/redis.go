package redis

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var rdb *redis.Client

func Init() {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Printf("%s:%f",
			viper.GetString("redis.host"),
			viper.GetInt("redis.port"),
		),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.password"),
		PoolSize: viper.GetInt("redis.pool_size")
	})
	_, err = rdb.Ping().Result()
	return err
}

func Close(){
	_ = rdb.Close()
}