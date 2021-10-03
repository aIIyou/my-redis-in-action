package chapter2

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func AddToCart(rdb *redis.Client, user string, item string, count int) error {
	if count <= 0 {
		intCmd := rdb.HDel(context.Background(), "cart:"+user, item)
		err := intCmd.Err()
		if err != nil {
			return err
		}
	} else {
		intCmd := rdb.HSet(context.Background(), "cart:"+user, item, count)
		err := intCmd.Err()
		if err != nil {
			return err
		}
	}
	return nil
}
