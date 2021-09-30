package chapter1

import (
	"context"
	"time"
)

func (rdb *RClient) GetGroupArticles(group string, order string) ([]string, error) {
	key := order + ":" + group

	intCmd := rdb.SInterStore(context.Background(), key, "group:"+group, order)
	err := intCmd.Err()
	if err != nil {
		return nil, err
	}

	rdb.Expire(context.Background(), key, time.Minute)

	strSliceCmd := rdb.SMembers(context.Background(), key)
	articles, err := strSliceCmd.Result()
	if err != nil {
		return nil, err
	}
	return articles, nil

}
