package database

import "github.com/redis/go-redis/v9"

var Cache *redis.Client

func SetupRedis() {
	Cache = redis.NewClient(&redis.Options{
		
	})
}