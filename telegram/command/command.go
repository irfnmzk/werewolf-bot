package command

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

type command struct {
	log *logrus.Logger
	bot *tgbotapi.BotAPI
	msg *tgbotapi.Message
}

func NewCommand(log *logrus.Logger, bot *tgbotapi.BotAPI, msg *tgbotapi.Message) Command {
	return &command{
		log: log,
		bot: bot,
		msg: msg,
	}
}

type Command interface {
	StartGame()
	UnknownCommand()
	Help()
	About()
	Version()
}
