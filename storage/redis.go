package storage

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

var ctx = context.Background()

type RedisParameters struct {
	Addr     string
	Username string
	Password string
}

type RedisInterface struct {
	client *redis.Client
	logger *logrus.Logger
}

func (ri *RedisInterface) Init(params interface{}, logger *logrus.Logger) error {
	redisParams := params.(RedisParameters)
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisParams.Addr,
		Username: redisParams.Username,
		Password: redisParams.Password,
		DB:       0, // use default db
	})

	ri.logger = logger
	ri.client = rdb

	return nil
}

func (ri *RedisInterface) DelState(key string) error {
	ri.client.Del(ctx, key)
	return nil
}
