package chapter1

import (
	"context"
	"fmt"
	"strings"
	"time"
)

const INCR_SCORE = 432

func (rds *RClient) ArticalVote(articalID, userID string) error {
	//查有序集合得到文章的发布时间
	floatResult := rds.ZScore(context.Background(), "time", articalID)
	t, err := floatResult.Result()
	if err != nil {
		return err
	}
	now := time.Now().Unix()
	if now-int64(t) > 7*24*60*60 {
		return fmt.Errorf("%s is published 7 days ago", articalID)
	}

	//查集合得到是否userid已经给articalid投过票了
	id := articalID[strings.IndexByte(articalID, ':'):]
	intResult := rds.SAdd(context.Background(), "voted"+id, userID)
	count, err := intResult.Result()
	if err != nil {
		return err
	}

	//首次投票，更新score,更新hash
	if count == 1 {
		floatResult = rds.ZIncrBy(context.Background(), "score", INCR_SCORE, articalID)
		_, err = floatResult.Result()
		if err != nil {
			return err
		}

		intResult = rds.HIncrBy(context.Background(), articalID, "votes", 1)
		_, err = intResult.Result()
		return err
	} else {
		return fmt.Errorf("%s has already voted for %s", userID, articalID)
	}
}
