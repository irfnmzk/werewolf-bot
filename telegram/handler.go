package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (tc *Client) handleCommand(msg *tgbotapi.Message) {
	cmd := msg.Command()

	tc.log.Debug("data")
	tc.log.Info(cmd)
}

func (tc *Client) handleCallback(cb *tgbotapi.CallbackQuery) {
	tc.log.Debug(cb.Message.Chat.ID)
}
