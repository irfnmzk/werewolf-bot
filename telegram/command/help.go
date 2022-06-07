package command

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/viper"
)

func (c *command) Help() {
	helpKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("about", "about"),
			tgbotapi.NewInlineKeyboardButtonData("version", "version"),
		),
	)

	c.sendMessageWithMarkup(viper.GetString("common.help"), helpKeyboard)
}
