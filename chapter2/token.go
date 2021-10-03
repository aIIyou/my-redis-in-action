package chapter2

import (
	"context"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

func CheckToken(rdb *redis.Client, token string) (int64, error) {
	key := LOG_KEY + token
	strCmd := rdb.HGet(context.Background(), key, "uid")
	uid, err := strCmd.Result()
	if err != nil {
		return -1, nil
	}

	ret, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		return -1, nil
	}
	return ret, nil
}

func UpdateToken(rdb *redis.Client, token string, user User, item string) error {
	t0 := time.Now().Unix()

	intCmd := rdb.ZAdd(context.Background(), "recent", &redis.Z{Score: float64(t0), Member: token})
	err := intCmd.Err()
	if err != nil {
		return err
	}

	intCmd = rdb.HSet(context.Background(), LOG_KEY+token, "uid", user.UID, "name", user.Name, "last_login_time", t0)
	err = intCmd.Err()
	if err != nil {
		return err
	}
	if item != "" {
		intCmd = rdb.ZAdd(context.Background(), "viewed:"+token, &redis.Z{Score: float64(t0), Member: item})
		err = intCmd.Err()
		if err != nil {
			return err
		}

		intCmd = rdb.ZRemRangeByRank(context.Background(), "viewed:"+token, 0, -26)
		err = intCmd.Err()
		if err != nil {
			return err
		}
	}
	return nil
}
