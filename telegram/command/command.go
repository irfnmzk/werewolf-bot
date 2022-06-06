package command

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/irfnmzk/werewolf-arena/storage"
	"github.com/sirupsen/logrus"
)

type command struct {
	log *logrus.Logger
	bot *tgbotapi.BotAPI
	msg *tgbotapi.Message

	redis *storage.RedisInterface
}

func NewCommand(log *logrus.Logger, bot *tgbotapi.BotAPI, msg *tgbotapi.Message, redis *storage.RedisInterface) Command {
	return &command{
		log:   log,
		bot:   bot,
		msg:   msg,
		redis: redis,
	}
}

type Command interface {
	GreetingJoinGroupOrChannel()
	StartGame()
	UnknownCommand()
	Help()
	About()
	Version()
}
