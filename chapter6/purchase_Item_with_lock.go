package chapter6

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

func PurchaseItemWithLock(rdb *redis.Client, itemID, sellerID, buyerID string) error {
	//提前准备好两个hash，一个有序集合，一个集合的健名
	itemKey := itemID + ":" + sellerID  //查market
	inventory := "inventory:" + buyerID //给买家包裹加商品
	seller := "user:" + sellerID
	buyer := "user:" + buyerID

	//Step1. 获取分布式锁
	uuid, err := AcquireLock(rdb, "market", 5*time.Second)
	if err != nil {
		log.Fatal(err.Error())
	}

	ctx := context.Background()

	//Step2. 查询市场（有序集合）中的商品的价格
	floatCmd := rdb.ZScore(ctx, "market", itemKey)
	price, err := floatCmd.Result()
	if err != nil {
		log.Fatal(err.Error())
	}

	//Step3. 查询买家的账户余额
	stringCmd := rdb.HGet(ctx, buyer, "fund")
	fundstr, err := stringCmd.Result()
	if err != nil {
		log.Fatal(err.Error())
	}
	fund, err := strconv.ParseFloat(fundstr, 64)
	if err != nil {
		log.Fatal(err.Error())
	}

	//Step4. 判断账户余额是否充足
	if fund < price {
		err = ReleaseLock(rdb, "market", uuid)
		if err != nil {
			return err
		}
		return fmt.Errorf("%s's fund is not enough", buyerID)
	}

	//Step5. 给买家账户余额减钱
	floatCmd = rdb.HIncrByFloat(ctx, buyer, "fund", -1*price)
	err = floatCmd.Err()
	if err != nil {
		log.Fatal(err.Error())
	}

	//Step6. 给卖家账户余额加钱
	floatCmd = rdb.HIncrByFloat(ctx, seller, "fund", price)
	err = floatCmd.Err()
	if err != nil {
		log.Fatal(err.Error())
	}

	//Step7. 给买家的包裹家商品
	intCmd := rdb.SAdd(ctx, inventory, itemID)
	err = intCmd.Err()
	if err != nil {
		log.Fatal(err.Error())
	}

	//Step8. 从市场上删除掉商品
	intCmd = rdb.ZRem(ctx, "market", itemKey)
	err = intCmd.Err()
	if err != nil {
		log.Fatal(err.Error())
	}

	//Step9. 释放分布式锁
	err = ReleaseLock(rdb, "market", uuid)
	return err
}
