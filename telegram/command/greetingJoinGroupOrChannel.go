package command

import (
	"fmt"

	"github.com/spf13/viper"
)

func (c *command) GreetingJoinGroupOrChannel() {
	roomState := c.redis.GetRoomState(c.msg.Chat.ID)
	if roomState == nil {
		_, err := c.setAdministratorRoom()
		if err != nil {
			return
		}
	}

	text := buildMessage(c.msg.From.FirstName, c.msg.From.LastName)
	c.sendMessage(text)
}

func buildMessage(firstName string, lastName string) (text string) {
	name := fmt.Sprintf("%s %s", firstName, lastName)
	text = fmt.Sprintf(viper.GetString("common.join"), name)
	return
}
