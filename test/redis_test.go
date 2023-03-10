package test

import (
	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
	"testing"
	"time"
)

func TestRedisDB(t *testing.T) {
	ENV, err := godotenv.Read()
	var Redis *redis.Client
	Redis = redis.NewClient(&redis.Options{
		Addr:        ENV["REDIS_URL"],
		Password:    ENV["REDIS_PASSWORD"],
		DB:          0,
		DialTimeout: 10 * time.Second,
	})
	pong, err := Redis.Ping().Result()
	if err != nil {
		t.Fatal("Ping Redis Error:" + err.Error())
	}
	defer func(client *redis.Client) {
		err := client.Close()
		if err != nil {
			t.Fatal(err)
		}
	}(Redis)
	t.Log("Successfully connected redis and " + pong)
}
