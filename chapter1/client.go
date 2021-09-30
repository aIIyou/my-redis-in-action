package chapter1

import (
	"github.com/go-redis/redis/v8"
)

type RClient struct {
	*redis.Client
}

var Client *redis.Client

func NewClient(host, password string, db int) *RClient {
	Client = redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       db,
	})
	return &RClient{Client: Client}
}
