package main

import (
	"log"
	"os"
	"regexp"
	"shadowsky-qiandao/notification"
	"strconv"
	"sync"
	"time"
)

// 成功： {Msg:获得了 259 MB流量. Ret:1}
// 已签到： {Msg:您似乎已经签到过了... Ret:1}

type CheckinResult struct {
	Msg string `json:"msg"`
	Ret int    `json:"ret"`
}

// DataN parse the checkin response message and get the number of free data in MB just got
func (cr *CheckinResult) DataN() int {
	compRegex := regexp.MustCompile(`^获得了 (\d+) MB流量$`)
	res := compRegex.FindStringSubmatch(cr.Msg)
	if len(res) != 2 {
		return 0
	}
	n, err := strconv.Atoi(res[1])
	if err != nil {
		return 0
	}
	return n
}

// checkin login to shadowsky do checkin-job and send notification
func checkin() {
	now := time.Now()
	cr := &CheckinResult{}
	if !notifyOnly {
		ss, err := NewShadowsky()
		if err != nil {
			log.Printf("login to shadowsky %v", err)
		}
		cr, err = ss.Checkin()
		if err != nil {
			log.Printf("check in %v", err)
		}
	} else {
		// mock check result
		cr.Msg = "Hey There..."
	}
	log.Println("checkin message:", cr.Msg)

	msgToSend := notification.Message{
		Body: "Shadowsky 签到结果：" + cr.Msg,
	}
	var channels []notification.Channel
	if os.Getenv("TELEGRAM_KEY") != "" {
		channels = append(channels, &notification.TelegramChannel{})
	}
	if os.Getenv("DISCORD_WEBHOOK") != "" {
		channels = append(channels, &notification.DiscordChannel{})
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
	log.Printf("send notification complete, total %d, task duration: %v", len(channels), time.Since(now))
}
