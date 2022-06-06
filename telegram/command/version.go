package command

func (c *command) Version() {
	c.sendMessage("Beta")
}
