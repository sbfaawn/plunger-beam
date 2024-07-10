package cache_test

import (
	"plunger-beam/internal/cache"
	"testing"
)

func TestConnect(t *testing.T) {
	redisOption := cache.RedisOption{
		Address:  "localhost",
		Port:     "6379",
		DbNum:    0,
		Password: "pass123",
	}

	conn := cache.NewRedisConnection(redisOption)

	if err := conn.ConnectToRedis(); err != nil {
		t.Error("Expected successful connection to redis but error : ", err)
	}
}
