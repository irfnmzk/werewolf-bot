package storage

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/irfnmzk/werewolf-arena/werewolf"
)

func (ri *RedisInterface) GetRoomState(chatId int64) *werewolf.RoomState {
	key := fmt.Sprintf("room:%d", chatId)

	jsonStr, err := ri.client.Get(ctx, key).Result()
	switch {
	case errors.Is(err, redis.Nil):
		ri.logger.Info(fmt.Sprintf("an user in a game for room %d is nil", chatId))
		return nil
	case err != nil:
		ri.logger.Error(err)
		return nil
	default:
		dgs := werewolf.RoomState{}
		err := json.Unmarshal([]byte(jsonStr), &dgs)

		if err != nil {
			ri.logger.Error(err)
			return nil
		}
		return &dgs
	}
}

func (ri *RedisInterface) SetRoomState(data *werewolf.RoomState) {
	key := fmt.Sprintf("room:%d", data.ChatId)

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
