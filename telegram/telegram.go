package telegram

import (
	tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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
}

func New(config *ClientConfig, log *logrus.Logger) *Client {
	bot, err := tgapi.NewBotAPI(config.Token)
	if err != nil {
		log.Fatal("cannot connect to telegram")
	}

	tc := &Client{config, log, bot}
	log.Info("Initializing telegram client")

	return tc
}

func (tc *Client) Start() {
	tc.log.Info("Starting telegram client")

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
	if update.Message.IsCommand() {
		tc.handleCommand(update.Message)
	}
}
