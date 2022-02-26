package discord

import (
	"github.com/go-resty/resty/v2"

	"github.com/abowloflrf/shadowsky-qiandao/notification"
)

// channel discord
type channel struct {
	Webhook string
}

var _ notification.Channel = &channel{}

func NewChannel(webhook string) notification.Channel {
	return &channel{
		Webhook: webhook,
	}
}

// Send notification message to discord chat
func (dc *channel) Send(message notification.Message) error {
	client := resty.New()
	_, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]string{"content": message.Body}).
		Post(dc.Webhook)
	return err
}

// Name of channel
func (dc *channel) Name() string {
	return "discord"
}
