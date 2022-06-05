package command

func (c *command) UnknownCommand() {
	c.sendMessage("command tidak ditemukan! silahkan ketik /help untuk melihat daftar command.")
}
