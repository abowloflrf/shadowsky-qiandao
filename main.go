package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"shadowsky-qiandao/notification"
	"sync"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/robfig/cron"
)

var (
	once       bool
	notifyOnly bool
)

func init() {
	flag.BoolVar(&once, "once", false, "run once then exit")
	flag.BoolVar(&notifyOnly, "notify-only", false, "dry-run, send notification only")
}

// checkin login to shadowsky and checkin
func checkin() {
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

	log.Println(cr.Msg)

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
	log.Printf("send notification complete, total %d", len(channels))
}

func main() {
	flag.Parse()

	log.Println("start shadowsky-qiandao, prepare to load config file")
	err := godotenv.Load()
	if err != nil {
		log.Println("load dotenv config file", err.Error())
	}

	if once {
		log.Println("run once...")
		checkin()
		return
	}

	c := cron.New()
	cron := os.Getenv("CRON")
	if len(cron) == 0 {
		cron = "0 0 8 * * *" // 默认每天上午8点指定
	}
	err = c.AddFunc(cron, func() {
		checkin()
	})
	if err != nil {
		log.Fatalf("add cronjob %v", err)
	}

	log.Println("start checkin cron job")
	c.Start()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	for range stop {
		log.Println("stop cron job")
		c.Stop()
		return
	}
}
