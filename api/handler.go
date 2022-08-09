//package api
//
//import (
//	"encoding/json"
//	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
//	"io/ioutil"
//	"log"
//	"net/http"
//	"os"
//)
//
//var bot *tgbotapi.BotAPI
//
//func init() {
//	var err error
//	bot, err = tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
//	if err != nil {
//		log.Panic(err)
//	}
//}
//
//func Repeater(w http.ResponseWriter, r *http.Request) {
//	defer r.Body.Close()
//
//	body, _ := ioutil.ReadAll(r.Body)
//
//	var update tgbotapi.Update
//
//	err := json.Unmarshal(body, &update)
//	if err != nil {
//		log.Println(err)
//		return
//	}
//
//	if update.Message.Text != "" {
//		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
//
//		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
//		_, err := bot.Send(msg)
//		if err != nil {
//			log.Println(err)
//		}
//	}
//}

package api

import (
	"encoding/json"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io"
	"log"
	"net/http"
	"os"
	"pwd/clients/telegram"
)

var bot *tgbotapi.BotAPI

func init() {
	var err error
	bot, err = tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, _ := io.ReadAll(r.Body)

	var update tgbotapi.Update

	err := json.Unmarshal(body, &update)
	if err != nil {
		log.Println(err)
		return
	}

	if update.Message.Text != "" {
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		telegram.SendMsg(telegram.ReceiveRequest(bot), bot)
	}
}
