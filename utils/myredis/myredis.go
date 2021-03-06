package myredis

import (
	"github.com/go-redis/redis"
	"test.liuda.com/gotest/utils/setting"
)

var Client *redis.Client

func Setup() {
	Client = redis.NewClient(&redis.Options{
		Addr:     setting.RedisSetting.Host,
		Password: setting.RedisSetting.Password,
		DB:       setting.RedisSetting.Db,
	})
	//	//return client
}

func Close() {
	//err := Client.Close()
	//if err != nil {
	//	logging.Error(err)
	//}

}
