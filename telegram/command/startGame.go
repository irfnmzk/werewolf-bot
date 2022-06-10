package command

import (
	"os"
	"time"

	"github.com/irfnmzk/werewolf-arena/werewolf"
	"github.com/spf13/viper"
)

func (c *command) StartGame() {
	isGroup := c.msg.Chat.Type == "group" || c.msg.Chat.Type == "supergroup"
	if !isGroup {
		c.sendMessage(viper.GetString("common.not_in_group"))
		return
	}

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
	if gameState != nil {
		c.sendMessage(viper.GetString("common.game_already_created"))
		return
	}

	gameState = werewolf.NewGameState(c.msg.Chat.ID)
	c.redis.SetGameState(gameState)

	lobbyTimeOut, _ := time.ParseDuration(os.Getenv("LOBBY_TIMEOUT"))
	time.AfterFunc(lobbyTimeOut, c.KillGame)

	c.sendMessage(viper.GetString("common.game_created"))
}
