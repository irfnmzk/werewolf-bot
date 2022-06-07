package command

import "github.com/spf13/viper"

func (c *command) About() {
	c.sendMessage(viper.GetString("common.about"))
}
