package command

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/irfnmzk/werewolf-arena/storage"
	"github.com/irfnmzk/werewolf-arena/werewolf"
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
	JoinInGame()
	KillGame()
}

func (c *command) getAdministratorInRoom() (userIds []int64, err error) {
	config := tgbotapi.ChatAdministratorsConfig{
		ChatConfig: tgbotapi.ChatConfig{
			ChatID:             c.msg.Chat.ID,
			SuperGroupUsername: c.msg.Chat.UserName,
		},
	}

	admins, err := c.bot.GetChatAdministrators(config)
	if err != nil {
		c.log.Info("get administrator is error: %s", err)
		return
	}

	for i := range admins {
		if admins[i].IsAdministrator() || admins[i].IsCreator() {
			userIds = append(userIds, admins[i].User.ID)
		}
	}

	return
}

func (c *command) setAdministratorRoom() (isAdmin bool, err error) {
	userIds, err := c.getAdministratorInRoom()
	if err != nil {
		return
	}

	roomState := werewolf.NewRoomState(c.msg.Chat.ID, userIds)
	c.redis.SetRoomState(roomState)

	isAdmin, err = c.checkAdministratorRoom(roomState)

	return
}

func (c *command) checkAdministratorRoom(roomState *werewolf.RoomState) (isAdmin bool, err error) {
	for i := range roomState.UserId {
		if roomState.UserId[i] == c.msg.From.ID {
			isAdmin = true
			return
		}
	}

	err = fmt.Errorf("player is not an administrator")

	return
}
