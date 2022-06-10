package command

import (
	"fmt"

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

	gameState = werewolf.SetCurrentPlayerState(gameState.ChatId, gameState.CurrentPlayer+1)
	c.redis.SetGameState(gameState)

	c.sendMessage(fmt.Sprintf(viper.GetString("common.player_joined"), c.msg.From.FirstName, c.msg.From.LastName))
	c.maxPlayer(gameState.CurrentPlayer)
}
