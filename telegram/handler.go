package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/irfnmzk/werewolf-arena/telegram/command"
)

func (tc *Client) handleCommand(msg *tgbotapi.Message, text string) {
	command := command.NewCommand(tc.log, tc.bot, msg, tc.redisClient)
	switch text {
	case "start", "mulai":
		command.StartGame()
	case "about":
		command.About()
	case "version":
		command.Version()
	case "help":
		command.Help()
	default:
		command.UnknownCommand()
	}
}

func (tc *Client) handleNonCommand(msg *tgbotapi.Message) {
	message := tgbotapi.NewMessage(msg.Chat.ID, "please enter the command!")
	_, err := tc.bot.Send(message)
	if err != nil {
		log.Panic(err)
	}
}

func (tc *Client) handleCallback(cb *tgbotapi.CallbackQuery) {
	// Respond to the callback query, telling Telegram to show the user
	// a message with the data received.
	callback := tgbotapi.NewCallback(cb.ID, cb.Data)
	if _, err := tc.bot.Request(callback); err != nil {
		panic(err)
	}

	tc.handleCommand(cb.Message, cb.Data)
}
