package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/irfnmzk/werewolf-arena/telegram"
	"github.com/sirupsen/logrus"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	log := logrus.New()
	log.Formatter = &logrus.TextFormatter{ForceColors: true, FullTimestamp: true}
	log.Level = logrus.DebugLevel

	log.Info("Starting application")

	token := os.Getenv("TELEGRAM_TOKEN")
	if token == "" {
		log.Fatal("no TELEGRAM_TOKEN provided")
	}

	tc := telegram.New(&telegram.ClientConfig{Token: token, Webhook: false, WebhookUrl: ""}, log)

	tc.Start()

	// close signal
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	<-sc
	log.Info("Stopping service")
	time.Sleep(time.Second)
}
