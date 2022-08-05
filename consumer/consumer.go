package consumer

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"net/url"
	"pwd/internal"
)

type Consumer interface {
}

func Consume(message *tgbotapi.Message) (tgbotapi.Chattable, error) {
	switch message.Text {
	case "/hello":
		return tgbotapi.NewMessage(message.Chat.ID, internal.MsgHello), nil
	//case "/dice":
	//	return tgbotapi.NewAnimation(message.Chat.ID, internal.GifDiceRolling())
	default:
		return answerWithGameInfo(message), nil
	}

	err := errors.New("Command doesn't meet the conditions")

	return tgbotapi.NewMessage(message.Chat.ID, ""), err
}

func answerWithGameInfo(message *tgbotapi.Message) tgbotapi.Chattable {
	gameInfo := internal.GameInfo{
		Name:        message.Text,
		TeseraLink:  url.URL{},
		VkLinkBNI:   url.URL{},
		AvitoLink:   url.URL{},
		YoutubeLink: "",
		GoogleLink:  url.URL{},
		InfoFromTesera: struct {
			Name                       string
			Description                string
			ImageUrl                   string
			RecommendedAge             string
			NumberOfPlayers            string
			RecommendedNumberOfPlayers string
			GameTime                   string
		}{Name: "", Description: "", ImageUrl: "", RecommendedAge: "", NumberOfPlayers: "", RecommendedNumberOfPlayers: "", GameTime: ""},
	}
	gameInfo.AllLinks(gameInfo.Name)
	var msg tgbotapi.MessageConfig
	if gameInfo.InfoFromTesera.RecommendedAge != "" {
		msg = tgbotapi.NewMessage(message.Chat.ID, internal.MsgGameInfo(gameInfo))
	} else {
		msg = tgbotapi.NewMessage(message.Chat.ID, internal.MsgGameNotFound+internal.MsgShortInfo(gameInfo))
	}
	msg.ParseMode = "HTML"
	msg.ReplyToMessageID = message.MessageID
	return msg
}

//func button() {
//	var button = tgbotapi.NewInlineKeyboardMarkup(
//		tgbotapi.NewInlineKeyboardRow(
//			tgbotapi.NewInlineKeyboardButtonData("Кинуть кубик", "Кинуть кубик")),
//	)
//
//	var button = tgbotapi.NewInlineKeyboardMarkup(
//		tgbotapi.NewInlineKeyboardRow(
//			tgbotapi.NewInlineKeyboardButtonData("Кинуть кубик", "Кинуть кубик")),
//	)
//
//	if update.Message.Text == "open" {
//		msg.ReplyMarkup = button
//	}
//
//	// Send the message.
//	if _, err = bot.Send(msg); err != nil {
//		panic(err)
//	}
//
//	if update.CallbackQuery != nil {
//		// Respond to the callback query, telling Telegram to show the user
//		// a message with the data received.
//		callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
//		if _, err := bot.Request(callback); err != nil {
//			panic(err)
//		}
//
//		// And finally, send a message containing the data received.
//		msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
//		if _, err := bot.Send(msg); err != nil {
//			panic(err)
//		}
//
//		if update.CallbackQuery != nil {
//			// Respond to the callback query, telling Telegram to show the user
//			// a message with the data received.
//			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
//			if _, err := bot.Request(callback); err != nil {
//				panic(err)
//			}
//
//			// And finally, send a message containing the data received.
//			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
//			if _, err := bot.Send(msg); err != nil {
//				panic(err)
//			}
//		}
//	}
//}
