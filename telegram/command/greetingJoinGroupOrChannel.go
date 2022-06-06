package command

import "fmt"

func (c *command) GreetingJoinGroupOrChannel() {
	text := c.buildMessage()
	c.sendMessage(text)
}

func (c *command) buildMessage() (text string) {
	name := fmt.Sprintf("%s %s", c.msg.From.FirstName, c.msg.From.LastName)
	text = fmt.Sprintf(
		`Hai! kak %s dan kakak yang lain.
Makasih udah add aku ke grup, semoga dengan adanya aku bisa membuat grup ini tambah seru dan menyenangkan
		
Berikut daftar perintah yang kakak bisa gunakan untuk dapat bermain bersama aku:
/help - bantuan
/about - tentang aku
/version - umur aku

Selamat bermain :)
		`, name)
	return
}
