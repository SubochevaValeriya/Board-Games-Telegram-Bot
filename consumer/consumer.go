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
case "/advise":
		var name string
		var err error
		for {
			name, err = internal.RandomGame(internal.ConnectToBGGClient())
			if err == nil {
				break
			}
		}
		message.Text = name
		fallthrough
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
		BGGLink:     "",
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
