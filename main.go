package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	_ "time/tzdata"

	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

var (
	// cmd flag
	once       bool
	notifyOnly bool

	// env config
	cronExp           string // cron schedule expressions
	timezone          string // timezone
	shadowskyURL      string
	shadowskyEmail    string
	shadowskyPassword string
	userAgent         string
	telegramKey       string
	telegramChatID    string
	discordWebhook    string
)

func main() {
	log.Println("start shadowsky-qiandao, prepare to load config")
	// parse cmd flag
	flag.BoolVar(&once, "once", false, "run once then exit")
	flag.BoolVar(&notifyOnly, "notify-only", false, "dry-run, send notification only")
	flag.Parse()
	// parse env config
	if err := godotenv.Load(); err != nil {
		if os.IsNotExist(err) {
			log.Println(".env file not found")
		} else {
			log.Fatalln("load dotenv config file", err.Error())
		}
	}
	cronExp = getEnvOr("CRON", "0 8 * * *")
	// timezone = getEnvOr("TZ", "Asia/Shanghai")
	shadowskyURL = getEnvOr("SHADOWSKY_URL", "https://www.shadowsky.xyz")
	shadowskyEmail = os.Getenv("SHADOWNSKY_EMAIL")
	shadowskyPassword = os.Getenv("SHADOWSKY_PASSWORD")
	userAgent = getEnvOr("SHADOWSKY_UA", "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.72 Safari/537.36")
	telegramChatID = os.Getenv("TELEGRAM_CHATID")
	telegramKey = os.Getenv("TELEGRAM_KEY")
	discordWebhook = os.Getenv("DISCORD_WEBHOOK")

	if once {
		log.Println("run once...")
		if err := checkin(); err != nil {
			log.Printf("checkin with error: %v", err)
			os.Exit(1)
		}
		return
	}

	c := cron.New()
	entryID, err := c.AddFunc(cronExp, func() {
		if err := checkin(); err != nil {
			log.Printf("checkin with error: %v", err)
		}
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

// getEnvOr use default value
func getEnvOr(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
