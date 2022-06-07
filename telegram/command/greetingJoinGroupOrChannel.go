package command

import (
	"fmt"

	"github.com/spf13/viper"
)

func (c *command) GreetingJoinGroupOrChannel() {
	text := c.buildMessage()
	c.sendMessage(text)
}

func (c *command) buildMessage() (text string) {
	name := fmt.Sprintf("%s %s", c.msg.From.FirstName, c.msg.From.LastName)
	text = fmt.Sprintf(viper.GetString("common.join"), name)
	return
}
