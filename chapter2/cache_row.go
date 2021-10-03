package chapter2

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

func Schedule_row_cache(rdb *redis.Client, rowID int, start int64, delay int64) error {
	rdb.ZAdd(context.Background(), "schedule:", &redis.Z{Score: float64(start), Member: rowID})
	rdb.ZAdd(context.Background(), "delay:", &redis.Z{Score: float64(delay), Member: rowID})
	return nil
}

func CacheRow(rdb *redis.Client) {
	for {
		zsliceCmd := rdb.ZRangeWithScores(context.Background(), "schedule:", 0, 0)
		s, err := zsliceCmd.Result()
		if err != nil {
			log.Fatal(err.Error())
		}
		if int64(s[0].Score) > time.Now().Unix() {
			time.Sleep(10 * time.Second)
		}

		floatCmd := rdb.ZScore(context.Background(), "delay:", s[0].Member.(string))
		t, err := floatCmd.Result()
		if err != nil {
			log.Fatal(err.Error())
		}

		Cache(s[0].Member.(string))

		if t <= 0 {
			rdb.ZRem(context.Background(), "schedule:", s[0].Member.(string))
			rdb.ZRem(context.Background(), "delay:", s[0].Member.(string))
		} else {
			rdb.ZAdd(context.Background(), "schedule:", &redis.Z{Score: float64(int64(t) + time.Now().Unix()), Member: s[0].Member})
		}

	}
}

func Cache(rowID string) {

}
