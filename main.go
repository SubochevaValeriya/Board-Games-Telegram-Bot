package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os"
	"pwd/clients/telegram"
)

func main() {
	// The code implements deployment using serverless function and webhooks (pls see api/handler.go)
	// In the main function I left the method using the long polling (all that is needed is to add a loop and change the Update type to UpdateChannels in the SendMsg function (sender.go))

	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		panic(err)
	}

	telegram.SendMsg(telegram.ReceiveRequest(bot), bot)
}
