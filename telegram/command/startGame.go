package command

import (
	"github.com/irfnmzk/werewolf-arena/werewolf"
	"github.com/spf13/viper"
)

func (c *command) StartGame() {
	isGroup := c.msg.Chat.Type == "group" || c.msg.Chat.Type == "supergroup"

	if !isGroup {
		c.sendMessage(viper.GetString("common.not_in_group"))
		return
	}

	gameState := c.redis.GetGameState(c.msg.Chat.ID)

	if gameState != nil {
		c.sendMessage(viper.GetString("common.already_created"))
		return
	}

	gameState = werewolf.NewGameState(c.msg.Chat.ID)
	c.redis.SetGameState(gameState)

	c.sendMessage(viper.GetString("common.created"))
}
