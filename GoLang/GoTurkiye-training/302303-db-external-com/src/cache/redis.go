package cache

import (
	"fmt"
	"github.com/go-redis/redis"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedis() *RedisCache {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost",
		Password: "",
		DB:       0,
	})

	if client == nil {
		fmt.Println("error couldn't connect to redis")
	}

	return &RedisCache{client: client}
}

func (r RedisCache) Get() []byte {
	strCmd := r.client.Get("cities")
	cacheBytes, err := strCmd.Bytes()
	if err != nil {
		return nil
	}
	return cacheBytes

}
func (r RedisCache) Put(value []byte) {
	r.client.Set("cities", value, 0)
}
