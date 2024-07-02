package api

import (
	"encoding/json"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"pwd/clients/telegram"
)

var bot *tgbotapi.BotAPI

func init() {
	var err error
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		panic("can't get bot token")
	}
	bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(fmt.Errorf("can't start bot: %v", err))
	}

	webhookURL := os.Getenv("WEBHOOK_URL")
	if webhookURL == "" {
		panic("can't get webhook url")
	}

	parsedWebhookURL, err := url.Parse(webhookURL)
	if err != nil {
		panic(fmt.Errorf("can't parse webhook url: %v", err))
	}

	_, err = bot.Request(tgbotapi.WebhookConfig{
		URL: parsedWebhookURL,
	})
	if err != nil {
		panic(fmt.Errorf("can't delete webhook: %v", err))
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Println(err)
	}

	var update tgbotapi.Update

	err = json.Unmarshal(body, &update)
	if err != nil {
		log.Println(err)
		return
	}

	if update.Message.Text != "" {
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		telegram.SendMsg(update, bot)
	}
}
