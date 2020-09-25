package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

var (
	once       bool
	notifyOnly bool
)

func init() {
	flag.BoolVar(&once, "once", false, "run once then exit")
	flag.BoolVar(&notifyOnly, "notify-only", false, "dry-run, send notification only")
}

func main() {
	flag.Parse()

	log.Println("start shadowsky-qiandao, prepare to load config file")
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("load dotenv config file", err.Error())
	}

	if once {
		log.Println("run once...")
		checkin()
		return
	}

	c := cron.New()
	cron := os.Getenv("CRON")
	if len(cron) == 0 {
		cron = "0 8 * * *" // 默认每天上午8点，注意时区
	}
	entryID, err := c.AddFunc(cron, func() {
		checkin()
	})
	if err != nil {
		log.Fatalf("add cronjob %v", err)
	}

	log.Println("start checkin cron job, entry ID", entryID)
	c.Start()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	for range stop {
		log.Println("stop cron job")
		c.Stop()
		return
	}
}
