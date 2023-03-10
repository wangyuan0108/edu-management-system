package db

import (
	"github.com/go-redis/redis"
	"log"
	"os"
	"time"
)

var Redis *redis.Client

func RedisDBInit() {
	Redis = redis.NewClient(&redis.Options{
		Addr:        os.Getenv("REDIS_URL"),
		Password:    os.Getenv("REDIS_PASSWORD"),
		DB:          0,
		DialTimeout: 30 * time.Second,
	})
	_, err := Redis.Ping().Result()
	if err != nil {
		log.Fatal("Ping Redis Error:" + err.Error())
	}
}
