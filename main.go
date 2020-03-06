package main

import (
	"log"
	"os"
	"os/signal"
	"shadowsky-qiandao/notification/tg"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/robfig/cron"
)

// checkin login to shadowsky and checkin
func checkin() {
	ss, err := NewShadowsky()
	if err != nil {
		log.Printf("login to shadowsky %v", err)
	}
	cr, err := ss.Checkin()
	if err != nil {
		log.Printf("check in %v", err)
	}

	log.Println(cr.Msg)
	if os.Getenv("TELEGRAM_KEY") != "" {
		tg.SendMsg("Shadowsky 签到结果：" + cr.Msg)
	}
}

func main() {

	log.Println("start shadowsky-qiandao, prepare to load config file")
	err := godotenv.Load()
	if err != nil {
		log.Println("load dotenv config file", err.Error())
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
