package command

import "github.com/spf13/viper"

func (c *command) UnknownCommand() {
	c.sendMessage(viper.GetString("common.unknown"))
}
