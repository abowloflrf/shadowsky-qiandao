package telegram

import (
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/abowloflrf/shadowsky-qiandao/notification"
)

// channel telegram
type channel struct {
	Key    string
	ChatID string
}

var _ notification.Channel = &channel{}

func NewChannel(key string, chatID string) notification.Channel {
	return &channel{
		Key:    key,
		ChatID: chatID,
	}
}

// Send notification message to telegram chat
func (tc *channel) Send(message notification.Message) error {
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
func (tc *channel) Name() string {
	return "telegram"
}
