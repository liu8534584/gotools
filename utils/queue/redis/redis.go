package redisQueue

import (
	"fmt"
	"github.com/go-redis/redis"
	"test.liuda.com/gotest/utils/logging"
)

type RedisInstance struct {
	RedisClient *redis.Client
}

//放入队列
func (redisInstance RedisInstance) PutData(name string, dataString []byte) int64 {
	result, err := redisInstance.RedisClient.LPush(name, dataString).Result()
	if err != nil {
		logging.Error(fmt.Sprintf("放入队列失败,queueName:%s,err:%v", name, err))
		return 0
	}
	//defer redisInstance.RedisClient.Close()
	return result
}

func (redisInstance RedisInstance) MPutData(name string, dataList []string) bool {
	for _, v := range dataList {
		redisInstance.PutData(name, []byte(v))
	}
	return true
}

//获取队列数据
func (redisInstance RedisInstance) GetData(name string, size int) []string {
	var arr []string
	for i := 0; i < size; i++ {
		result, err := redisInstance.RedisClient.RPop(name).Result()
		if err != nil {
			continue
		}
		arr = append(arr, result)
	}
	//defer redisInstance.RedisClient.Close()
	return arr
}

//队列长度
func (redisInstance RedisInstance) Length(name string) int64 {
	result, err := redisInstance.RedisClient.LLen(name).Result()
	//defer redisInstance.RedisClient.Close()
	if err != nil {
		return 0
	}
	return result
}

// 删除
func (redisInstance RedisInstance) Del(name string) int64 {
	result, err := redisInstance.RedisClient.Del(name).Result()
	//defer redisInstance.RedisClient.Close()
	if err != nil {
		return 1
	}
	return result
}
