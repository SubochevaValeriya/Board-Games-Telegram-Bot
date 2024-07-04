package api

import (
	"context"
	"encoding/json"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io"
	"log"
	"net/http"
	"pwd/clients/telegram"
	"pwd/pkg/resources"
)

// var bot *tgbotapi.BotAPI
var telegramService *telegram.Service

func init() {
	_, telegramService = resources.InitBot()

	//webhookURL := os.Getenv("WEBHOOK_URL")
	//if webhookURL == "" {
	//	panic("can't get webhook url")
	//}
	//
	//parsedWebhookURL, err := url.Parse(webhookURL)
	//if err != nil {
	//	panic(fmt.Errorf("can't parse webhook url: %v", err))
	//}
	//
	//_, err = bot.Request(tgbotapi.WebhookConfig{
	//	URL: parsedWebhookURL,
	//})
	//if err != nil {
	//	panic(fmt.Errorf("can't set webhook: %v", err))
	//}
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

		telegramService.SendMsg(context.Background(), update)
	}
}
