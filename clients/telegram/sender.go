package telegram

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type Service struct {
	consumerService consumerService
	bot             *tgbotapi.BotAPI
}

func New(
	consumerService consumerService,
	bot *tgbotapi.BotAPI,
) *Service {
	return &Service{
		consumerService: consumerService,
		bot:             bot,
	}
}

func (s *Service) SendMsg(ctx context.Context, update tgbotapi.Update) {
	var err error
	defer func() {
		if err != nil {
			log.Println(err)
		}
	}()
	var msg tgbotapi.MessageConfig
	msg.ReplyMarkup = keyboard

	if update.CallbackQuery != nil {
		msg, err = s.consumerService.ConsumeCallback(ctx, update.CallbackQuery)
	}

	if update.Message != nil {
		msg, err = s.consumerService.Consume(ctx, update.Message)
	}

	log.Println(msg)
	_, err = s.bot.Send(msg)
}

var keyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Во что поиграть?"),
		tgbotapi.NewKeyboardButton("Бросить кубик!"),
	),
)

func (s *Service) SendMsgWithoutWebhook(ctx context.Context, updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		s.SendMsg(ctx, update)
	}
}
