package discord

import (
	"github.com/abowloflrf/shadowsky-qiandao/notification"
	"github.com/go-resty/resty/v2"
)

// Channel discord
type Channel struct {
	Webhook string
}

var _ notification.Channel = &Channel{}

// Send notification message to discord chat
func (dc *Channel) Send(message notification.Message) error {
	client := resty.New()
	_, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]string{"content": message.Body}).
		Post(dc.Webhook)
	return err
}

// Name of channel
func (dc *Channel) Name() string {
	return "discord"
}
