package storage

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/irfnmzk/werewolf-arena/werewolf"
)

func (ri *RedisInterface) GetPlayerState(chatId int64, playerId int64) *werewolf.PlayerState {
	key := fmt.Sprintf("game:%d:%d", chatId, playerId)

	jsonStr, err := ri.client.Get(ctx, key).Result()
	switch {
	case errors.Is(err, redis.Nil):
		ri.logger.Info("player id %d in a game for chat id %d is nil", playerId, chatId)
		return nil
	case err != nil:
		ri.logger.Error(err)
		return nil
	default:
		dgs := werewolf.PlayerState{}
		err := json.Unmarshal([]byte(jsonStr), &dgs)

		if err != nil {
			ri.logger.Error(err)
			return nil
		}
		return &dgs
	}
}

func (ri *RedisInterface) SetPlayerState(data *werewolf.PlayerState) {
	key := fmt.Sprintf("game:%d:%d", data.ChatId, data.UserId)

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
