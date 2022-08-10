package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"pwd/consumer"
	"pwd/internal"
)

type sender interface {
	SendMsg(update tgbotapi.Update, bot *tgbotapi.BotAPI)
}

func SendMsg(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if update.Message.Text == "/dice" {
		gif, err := internal.GifDiceRolling()
		if err != nil {
			log.Println(err)
		}
		msgAnim := tgbotapi.NewAnimation(update.Message.Chat.ID, gif)
		if _, err := bot.Send(msgAnim); err != nil {
			log.Println(err)
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, internal.MsgRollingDice())
		if _, err := bot.Send(msg); err != nil {
			log.Println(err)
		}
	} else {
		if msg, err := consumer.Consume(update.Message); err == nil {
			if _, err := bot.Send(msg); err != nil {
				log.Println(err)
			}
		}
	}
}
