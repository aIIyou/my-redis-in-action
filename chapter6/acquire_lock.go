package chapter6

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

func AcquireLock(rdb *redis.Client, lockname string, timeout time.Duration) (string, error) {
	ctx := context.Background()
	key := "lock:" + lockname
	rand.Seed(time.Now().UnixNano())
	uuid := rand.Int()
	t0 := time.Now()
	for {
		boolCmd := rdb.SetNX(ctx, key, uuid, 0)
		b, err := boolCmd.Result()
		if err != nil {
			log.Fatal(err.Error())
		}
		if b {
			return strconv.FormatInt(int64(uuid), 10), nil
		} else {
			time.Sleep(time.Second)
		}
		if time.Since(t0) > timeout {
			break
		}
	}
	return "", errors.New("acquire lock timeout")
}
