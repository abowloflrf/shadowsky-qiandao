package tg

import (
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// SendMsg send notification message to specified telegram chat
func SendMsg(message string) error {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_KEY"))
	if err != nil {
		return err
	}
	chatID, err := strconv.ParseInt(os.Getenv("TELEGRAM_CHATID"), 10, 64)
	if err != nil {
		return err
	}
	msg := tgbotapi.NewMessage(chatID, message)
	_, err = bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}
