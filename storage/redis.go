package storage

import "github.com/go-redis/redis/v8"

type RedisParameters struct {
	Addr     string
	Username string
	Password string
}

type RedisInterface struct {
	client *redis.Client
}

func (ri *RedisInterface) Init(params interface{}) error {
	redisParams := params.(RedisParameters)
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisParams.Addr,
		Username: redisParams.Username,
		Password: redisParams.Password,
		DB:       0, // use default db
	})

	ri.client = rdb

	return nil
}
