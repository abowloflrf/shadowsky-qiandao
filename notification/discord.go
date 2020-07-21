package notification

import (
	"os"

	"github.com/go-resty/resty/v2"
)

type DiscordChannel struct {
}

func (dc *DiscordChannel) Send(message Message) error {
	client := resty.New()
	_, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]string{"content": message.Body}).
		Post(os.Getenv("DISCORD_WEBHOOK"))
	return err
}

func (dc *DiscordChannel) Name() string {
	return "discord"
}
