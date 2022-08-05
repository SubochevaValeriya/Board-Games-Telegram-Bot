package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"pwd/clients/telegram"
)

func main() {
	//token := mustToken()

	//bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	bot, err := tgbotapi.NewBotAPI("5501151328:AAFVVneFN6O4SLihdwM3qdOTxHmY8mtvNtQ")
	if err != nil {
		panic(err)
	}

	telegram.SendMsg(telegram.ReceiveRequest(bot), bot)
}

//
//func mustToken()
//string{
//token := flag.String("botToken", "", "token for access telegram Board Game bot")
//
//flag.Parse()
//
//if *token == ""{
//log.Fatal("token is not specified")
//}
//
//return *token
//}
