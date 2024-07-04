package telegram

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type consumerService interface {
	Consume(ctx context.Context, message *tgbotapi.Message) (tgbotapi.MessageConfig, error)
	ConsumeCallback(ctx context.Context, callbackQuery *tgbotapi.CallbackQuery) (tgbotapi.MessageConfig, error)
}
