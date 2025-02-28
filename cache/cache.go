package cache

import (
	"defaultproject/serializer"
	"defaultproject/util"
	"encoding/json"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

// RedisClient Redis缓存客户端单例
var RedisClient *redis.Client

// var RedisLogClient *redis.Client

// Redis 在中间件中初始化redis链接
func Redis() {
	db, _ := strconv.ParseUint(os.Getenv("REDIS_DB"), 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr:       os.Getenv("REDIS_ADDR"),
		Password:   os.Getenv("REDIS_PW"),
		DB:         int(db),
		MaxRetries: 1,
	})
	if _, err := client.Ping().Result(); err != nil {
		util.Log().Panic("连接Redis不成功", err)
	}
	RedisClient = client
}

func SetRedis(key string, res interface{}, expiration_input int) (err error) {
	expiry := time.Duration(expiration_input) * time.Second
	redis_data, _ := json.Marshal(res)
	err = RedisClient.Set(key, redis_data, expiry).Err()
	return
}
func GetRedis(key string, r *serializer.Response) (err error) {
	redis_data, err := RedisClient.Get(key).Result()
	if redis_data != "" && err == nil {
		json.Unmarshal([]byte(redis_data), &r)
	}
	return
}
