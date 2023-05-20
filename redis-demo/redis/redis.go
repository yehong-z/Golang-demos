package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type Client struct {
	RedisDB *redis.Client
}

func NewClient() (rdb *Client) {
	return &Client{RedisDB: redis.NewClient(&redis.Options{
		Addr:     "10.20.121.247:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})}
}

func (c *Client) Set(key string, value interface{}, expired time.Duration) (err error) {
	err = c.RedisDB.Set(ctx, key, value, expired).Err()
	return
}

func (c *Client) Get(key string) (val string, err error) {
	val, err = c.RedisDB.Get(ctx, key).Result()
	return
}

func Example() {
	redisCli := NewClient()
	_ = redisCli.Set("key", "value", 0)
	val, err := redisCli.Get("key")
	if err == redis.Nil {
		fmt.Println("key does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key", val)
	}
}
