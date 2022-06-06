package telegram

import (
	"os"

	tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/irfnmzk/werewolf-arena/storage"
	"github.com/irfnmzk/werewolf-arena/werewolf"
	"github.com/sirupsen/logrus"
)

type ClientConfig struct {
	Token      string
	Webhook    bool
	WebhookUrl string
}

type Client struct {
	config *ClientConfig

	log *logrus.Logger
	bot *tgapi.BotAPI

	redisClient *storage.RedisInterface
}

func New(config *ClientConfig, log *logrus.Logger) *Client {
	var redisClient storage.RedisInterface

	redisAddr := os.Getenv("REDIS_ADDR")
	redisPassword := os.Getenv("REDIS_PASS")

	if redisAddr == "" {
		log.Fatal("no REDIS_ADDR is provided")
	}

	err := redisClient.Init(storage.RedisParameters{
		Addr:     redisAddr,
		Username: "",
		Password: redisPassword,
	}, log)

	if err != nil {
		log.Fatal(err)
	}

	bot, err := tgapi.NewBotAPI(config.Token)
	if err != nil {
		log.Fatal("cannot connect to telegram")
	}

	tc := &Client{config, log, bot, &redisClient}
	log.Info("Initializing telegram client")

	return tc
}

func (tc *Client) Start() {
	tc.log.Info("Starting telegram client")

	gameLoop := werewolf.NewGameLoop(tc.redisClient)
	go gameLoop.Execute()

	// handle update
	uc := tgapi.NewUpdate(0)
	uc.Timeout = 30
	updates := tc.bot.GetUpdatesChan(uc)
	for update := range updates {
		tc.handleUpdate(update)
	}
	// TODO start webhook
}

func (tc *Client) handleUpdate(update tgapi.Update) {
	if update.Message != nil {
		if len(update.Message.NewChatMembers) > 0 {
			for _, item := range update.Message.NewChatMembers {
				if item.IsBot && item.UserName == "were_wolf_arena_bot" {
					tc.handleCommand(update.Message, "greeting_join_group_or_channel")
					return
				}
			}
		}

		if update.Message.LeftChatMember != nil {
			tc.handleLeftChat(update.Message)
			return
		}

		if !update.Message.IsCommand() {
			tc.handleNonCommand(update.Message)
			return
		}

		if update.Message.IsCommand() {
			text := update.Message.Command()
			tc.handleCommand(update.Message, text)
			return
		}
	}

	if update.CallbackQuery != nil {
		tc.handleCallback(update.CallbackQuery)
	}
}
