package utils

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	ctx       = context.Background()
	redisCli  *redis.Client
	redisPath = flag.String("redisPath", "localhost:6379", "redis连接地址，比如: localhost:6379")
	redisPwd  = flag.String("redisPwd", "", "redis密码")
)

func NewRedisCli() *redis.Client {
	if redisCli == nil {
		// 监控缓存个数
		go func() {
			for {
				cacheCount.WithLabelValues("redis").Set(float64(NewRedisCli().DBSize(ctx).Val()))
				time.Sleep(3 * time.Second)
			}
		}()
		redisCli = redis.NewClient(&redis.Options{
			Addr:     *redisPath,
			Password: *redisPwd, // no password set
			DB:       0,         // use default DB
		})
	}
	return redisCli
}

func ExampleClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	rdb.Exists(ctx, "123").Result()
	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}
