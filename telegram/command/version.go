package command

import "os"

func (c *command) Version() {
	c.sendMessage(os.Getenv("version"))
}
