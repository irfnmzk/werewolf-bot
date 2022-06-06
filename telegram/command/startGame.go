package command

import (
	"github.com/irfnmzk/werewolf-arena/state"
)

func (c *command) StartGame() {
	isGroup := c.msg.Chat.Type == "group" || c.msg.Chat.Type == "supergroup"

	if !isGroup {
		c.sendMessage("Perintah hanya bisa di lakukan dalam group")
		return
	}

	gameState := c.redis.GetGameState(c.msg.Chat.ID)

	if gameState != nil {
		c.sendMessage("Game sudah di buat!")
		return
	}

	gameState = state.NewGameState(c.msg.Chat.ID)
	c.redis.SetGameState(gameState)

	c.sendMessage("Game berhasil di buat")
}
