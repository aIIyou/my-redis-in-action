package chapter1

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type Article struct {
	Title  string
	Link   string
	Poster string
	Time   int64
	Votes  int
}

func (rds *RClient) ArticlePost(articel Article, userID string) error {
	intCmd := rds.Incr(context.Background(), "article")
	articleID, err := intCmd.Result()
	if err != nil {
		return err
	}

	//创建出一个set，存储已经投过票的用户id，这里需要直接把作者添加进来
	voteID := "voted:" + fmt.Sprintf("%08d", articleID)
	intCmd = rds.SAdd(context.Background(), voteID, userID)
	err = intCmd.Err()
	if err != nil {
		return err
	}
	boolCmd := rds.Expire(context.Background(), voteID, 7*24*time.Hour)
	err = boolCmd.Err()
	if err != nil {
		return err
	}

	t0 := time.Now().Unix()
	articel.Time = t0
	articel.Votes = 0
	//将新的文章存入hash
	articelKey := "article:" + fmt.Sprintf("%08d", articleID)
	boolCmd = rds.HMSet(context.Background(), articelKey, "title", articel.Title, "link", articel.Link, "time", t0, "poster", userID, "votes", 0)
	err = boolCmd.Err()
	if err != nil {
		return err
	}

	rds.ZAdd(context.Background(), "time", &redis.Z{Score: float64(t0), Member: articelKey})
	rds.ZAdd(context.Background(), "score", &redis.Z{Score: float64(t0), Member: articelKey})
	return nil
}
