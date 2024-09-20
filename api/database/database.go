package database

import (
    "context"
	"github.com/go-redis/redis/v8"
	"os"
}

var ctx = context.Background()
func CreateClient(dbNo int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		addr: os.Getenv("DB_ADDRESS"),
		password: os.Getenv("DB_PASS"),
		DB: dbNo,
	})
	return rdb
}