package chapter6

import (
	"context"
	"errors"

	"github.com/go-redis/redis/v8"
)

func ReleaseLock(rdb *redis.Client, lockname, uuid string) error {
	lockname = "lock:" + lockname

	for {
		//这里的watch合适吗？是不是应该在加锁的时候就watch
		err := rdb.Watch(context.Background(), func(tx *redis.Tx) error {
			intCmd := tx.Del(context.Background(), lockname) //解锁最核心的操作其实是删除键
			err := intCmd.Err()
			return err
		}, lockname)

		if err == redis.TxFailedErr {
			return errors.New("lock:market has been removed")
		}
		if err == nil {
			break
		}
		if err != nil {
			return err
		}
	}
	return nil
}
