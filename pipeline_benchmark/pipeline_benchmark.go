package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "10.160.69.177:6379",
		Password: "a2b1c4d3",
		DB:       0,
	})
	ctx := context.Background()
	statuCmd := rdb.Ping(ctx)
	statu, err := statuCmd.Result()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(statu)

	t0 := time.Now()
	rdb.Set(ctx, "benchmark", "one", time.Minute)
	rdb.Set(ctx, "benchmark", "two", time.Minute)
	rdb.Set(ctx, "benchmark", "three", time.Minute)
	rdb.Set(ctx, "benchmark", "four", time.Minute)
	rdb.Set(ctx, "benchmark", "five", time.Minute)
	fmt.Println(time.Since(t0)) //33.262022ms

	pipe := rdb.Pipeline()
	t0 = time.Now()
	pipe.Set(ctx, "benchmark1", "one", time.Minute)
	pipe.Set(ctx, "benchmark2", "two", time.Minute)
	pipe.Set(ctx, "benchmark3", "three", time.Minute)
	pipe.Set(ctx, "benchmark4", "four", time.Minute)
	statuCmd = pipe.Set(ctx, "benchmark5", "five", time.Minute)
	err = statuCmd.Err()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(time.Since(t0)) //21.445µs

}
