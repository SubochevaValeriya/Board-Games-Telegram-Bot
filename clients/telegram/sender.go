package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"pwd/consumer"
	"pwd/internal"
)

type sender interface {
	SendMsg(updates tgbotapi.UpdatesChannel, bot *tgbotapi.BotAPI)
}

func SendMsg(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if update.Message.Text == "/dice" {
		msgAnim := tgbotapi.NewAnimation(update.Message.Chat.ID, internal.GifDiceRolling())
		if _, err := bot.Send(msgAnim); err != nil {
			panic(err)
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, internal.MsgRollingDice())
		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}
	} else {
		if msg, err := consumer.Consume(update.Message); err == nil {
			if _, err := bot.Send(msg); err != nil {
				// Note that panics are a bad way to handle errors. Telegram can
				// have service outages or network errors, you should retry sending
				// messages or more gracefully handle failures.
				panic(err)
			}
		}
	}
}
