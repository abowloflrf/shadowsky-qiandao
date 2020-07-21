package notification

import (
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type TelegramChannel struct {
}

// SendMsg send notification message to specified telegram chat
func (tc *TelegramChannel) Send(message Message) error {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_KEY"))
	if err != nil {
		return err
	}
	chatID, err := strconv.ParseInt(os.Getenv("TELEGRAM_CHATID"), 10, 64)
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

func (tc *TelegramChannel) Name() string {
	return "telegram"
}
