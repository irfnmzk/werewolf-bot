package command

import (
	"fmt"

	"github.com/spf13/viper"
)

func (c *command) KillGame() {
	roomState := c.redis.GetRoomState(c.msg.Chat.ID)
	if roomState == nil {
		isAdmin, err := c.setAdministratorRoom()
		if err != nil || !isAdmin {
			return
		}
	} else {
		isAdmin, err := c.checkAdministratorRoom(roomState)
		if err != nil || !isAdmin {
			return
		}
	}

	gameState := c.redis.GetGameState(c.msg.Chat.ID)
	if gameState == nil {
		return
	}

	key := fmt.Sprintf("game:%d", c.msg.Chat.ID)
	c.redis.DelState(key)
	c.redis.DelPlayerState(key)
	c.sendMessage(viper.GetString("common.game_is_over"))
}
