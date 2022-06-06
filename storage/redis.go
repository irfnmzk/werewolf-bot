package storage

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/irfnmzk/werewolf-arena/werewolf"
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

func (ri *RedisInterface) GetGameState(chatId int64) *werewolf.GameState {
	key := fmt.Sprintf("game:%d", chatId)

	jsonStr, err := ri.client.Get(ctx, key).Result()
	switch {
	case errors.Is(err, redis.Nil):
		ri.logger.Info("game for chat id %d is nil", chatId)
		return nil
	case err != nil:
		ri.logger.Error(err)
		return nil
	default:
		dgs := werewolf.GameState{}
		err := json.Unmarshal([]byte(jsonStr), &dgs)

		if err != nil {
			ri.logger.Error(err)
			return nil
		}
		return &dgs
	}
}

func (ri *RedisInterface) SetGameState(data *werewolf.GameState) {
	key := fmt.Sprintf("game:%d", data.ChatId)

	jBytes, err := json.Marshal(data)
	if err != nil {
		ri.logger.Error(err)
		return
	}

	err = ri.client.Set(ctx, key, jBytes, 0).Err()
	if err != nil {
		ri.logger.Error(err)
		return
	}
}
