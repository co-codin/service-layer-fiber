package database

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var Cache *redis.Client
var CacheChannel chan string

func SetupRedis() {
	Cache = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
}

func SetupCacheChannel() {
	CacheChannel = make(chan string)

	for {
		go func(ch chan string) {
			time.Sleep(time.Second * 5)

			Cache.Del(context.Background(), <-ch)
		}(CacheChannel)
	}

}

func ClearCache(keys ...string) {
	for _, key := range keys {
		CacheChannel <- key
	}
}
