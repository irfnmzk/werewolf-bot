package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/irfnmzk/werewolf-arena/werewolf"
)

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

	lobbyTimeOut, _ := time.ParseDuration(os.Getenv("LOBBY_TIMEOUT"))
	err = ri.client.Set(ctx, key, jBytes, lobbyTimeOut).Err()
	if err != nil {
		ri.logger.Error(err)
		return
	}
}
