package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/abowloflrf/shadowsky-qiandao/notification"
	"github.com/abowloflrf/shadowsky-qiandao/notification/discord"
	"github.com/abowloflrf/shadowsky-qiandao/notification/telegram"
)

// checkin login to shadowsky do checkin-job and send notification
func checkin() error {
	now := time.Now()
	cr := &CheckinResult{}
	if !notifyOnly {
		ss, err := NewShadowsky(shadowskyEmail, shadowskyPassword,
			&ShadowskyConfig{
				URL:       shadowskyURL,
				UserAgent: userAgent,
			})
		if err != nil {
			return fmt.Errorf("login to shadowsky %v", err)
		}
		cr, err = ss.Checkin()
		if err != nil {
			return err
		}
	} else {
		// mock check result
		cr.Msg = "Hey There..."
	}
	log.Printf("checkin complete with message: %s, duration: %v", cr.Msg, time.Since(now))
	msgToSend := notification.Message{
		Body: "Shadowsky 签到结果：" + cr.Msg,
	}
	var channels []notification.Channel
	if telegramKey != "" {
		channels = append(channels, &telegram.Channel{
			Key:    telegramKey,
			ChatID: telegramChatID,
		})
	}
	if discordWebhook != "" {
		channels = append(channels, &discord.Channel{
			Webhook: discordWebhook,
		})
	}
	wg := sync.WaitGroup{}
	wg.Add(len(channels))
	for _, c := range channels {
		ch := c
		go func() {
			defer wg.Done()
			err := ch.Send(msgToSend)
			if err != nil {
				log.Printf("send notification error [%s]: %v", ch.Name(), err)
			}
			log.Printf("send notification [%s]", ch.Name())
		}()
	}
	wg.Wait()
	log.Printf("send notification complete, total channels %d, duration: %v", len(channels), time.Since(now))
	return nil
}
