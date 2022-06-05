package command

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *command) sendMessage(msg string) {
	message := tgbotapi.NewMessage(c.msg.Chat.ID, msg)
	_, err := c.bot.Send(message)
	if err != nil {
		c.log.Panic(err)
	}
}

func (c *command) sendMessageWithMarkup(msg string, numericKeyboard tgbotapi.InlineKeyboardMarkup) {
	message := tgbotapi.NewMessage(c.msg.Chat.ID, msg)
	message.ReplyMarkup = numericKeyboard
	_, err := c.bot.Send(message)
	if err != nil {
		c.log.Panic(err)
	}
}
