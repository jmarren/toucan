package cache

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client

func InitCache() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "password", // no password set
		DB:       0,          // use default DB
	})
	err := Rdb.Ping(context.Background()).Err()
	if err != nil {
		fmt.Printf("ERROR: ", err)
	}
	fmt.Println("Cache initialized successfully")
}
