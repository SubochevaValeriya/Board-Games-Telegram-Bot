package resources

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os"
	"pwd/clients/telegram"
	"pwd/pkg/consumer"
	"pwd/pkg/service/anticafe"
	"pwd/pkg/service/dice"
	"pwd/pkg/service/game"
	"pwd/pkg/service/links"
)

func InitBot() (*tgbotapi.BotAPI, *telegram.Service) {
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

	return bot, telegramService
}
