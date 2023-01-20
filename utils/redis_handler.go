package utils

import (
	"github.com/go-redis/redis/v9"
)

var rdb *redis.Client

func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "172.17.0.1:6379",
		Password: "password",
		DB:       0, // use default DB
	})
}

func GetRedis() *redis.Client {
	return rdb
}
