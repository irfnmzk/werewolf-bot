package telegram

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/irfnmzk/werewolf-arena/telegram/command"
	"github.com/spf13/viper"
)

func (tc *Client) handleCommand(msg *tgbotapi.Message, text string) {
	command := command.NewCommand(tc.log, tc.bot, msg, tc.redisClient)
	switch text {
	case "greeting_join_group_or_channel":
		command.GreetingJoinGroupOrChannel()
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

func (tc *Client) handleLeftChat(msg *tgbotapi.Message) {
	// Check if left chat member is bot werewolf
	// and will send a message for the kicker
	if msg.LeftChatMember.IsBot && msg.LeftChatMember.UserName == os.Getenv("USERNAME_BOT") {
		message := tgbotapi.NewMessage(msg.From.ID, viper.GetString("common.left"))
		_, err := tc.bot.Send(message)
		if err != nil {
			log.Panic(err)
		}
	}
}
