package command

import (
	"github.com/irfnmzk/werewolf-arena/werewolf"
	"github.com/spf13/viper"
)

func (c *command) JoinInGame() {
	gameState := c.redis.GetGameState(c.msg.Chat.ID)
	if gameState == nil {
		c.sendMessage(viper.GetString("common.game_not_created_yet"))
		return
	}

	playerState := c.redis.GetPlayerState(c.msg.Chat.ID, c.msg.From.ID)
	if playerState != nil {
		c.sendMessage(viper.GetString("common.player_already_joined"))
		return
	}

	playerState = werewolf.NewPlayerState(c.msg.Chat.ID, c.msg.From.ID)
	c.redis.SetPlayerState(playerState)

	c.sendMessage(viper.GetString("common.player_joined"))
}
