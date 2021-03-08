package telegram

import (
	"strconv"

	"github.com/abowloflrf/shadowsky-qiandao/notification"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Channel telegram
type Channel struct {
	Key    string
	ChatID string
}

var _ notification.Channel = &Channel{}

// Send notification message to telegram chat
func (tc *Channel) Send(message notification.Message) error {
	bot, err := tgbotapi.NewBotAPI(tc.Key)
	if err != nil {
		return err
	}
	chatID, err := strconv.ParseInt(tc.ChatID, 10, 64)
	if err != nil {
		return err
	}
	msg := tgbotapi.NewMessage(chatID, message.Body)
	_, err = bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

// Name of channel
func (tc *Channel) Name() string {
	return "telegram"
}
