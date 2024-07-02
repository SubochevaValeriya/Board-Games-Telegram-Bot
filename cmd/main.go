package main

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"os"
	"pwd/clients/telegram"
	"pwd/internal/consumer"
	"pwd/internal/service/anticafe"
	"pwd/internal/service/dice"
	"pwd/internal/service/game"
	"pwd/internal/service/links"
)

const envFile = ".env-test"

func main() {
	// The code implements deployment using serverless function and webhooks (pls see api/handler.go)
	// In the main function I left the method using the long polling (all that is needed is to add a loop and change the Update type to UpdateChannels in the SendMsg function (sender.go))

	if err := godotenv.Load(envFile); err != nil {
		panic("can't load env")
	}

	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		panic("can't get bot token")
	}
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(fmt.Errorf("can't start bot: %v", err))
	}
	_, err = bot.Request(tgbotapi.DeleteWebhookConfig{})
	if err != nil {
		panic(fmt.Errorf("can't delete webhook: %v", err))
	}

	diceService := dice.New()
	gameService := game.New()
	linksService := links.New(gameService)
	googleSheetsService := anticafe.New()

	consumerService := consumer.New(gameService, linksService, diceService, googleSheetsService, bot)
	telegramService := telegram.New(consumerService, bot)

	telegramService.SendMsgWithoutWebhook(context.Background(), telegramService.ReceiveRequest())
}
