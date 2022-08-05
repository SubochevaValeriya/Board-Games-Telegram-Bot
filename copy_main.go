//package main
//
//import (
//	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
//	"pwd/consumer"
//	"pwd/internal"
//)
//
//func main() {
//	//	token := mustToken()
//
//	//bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
//	bot, err := tgbotapi.NewBotAPI("5501151328:AAFVVneFN6O4SLihdwM3qdOTxHmY8mtvNtQ")
//	if err != nil {
//		panic(err)
//	}
//
//	bot.Debug = true
//
//	// Create a new UpdateConfig struct with an offset of 0. Offsets are used
//	// to make sure Telegram knows we've handled previous values and we don't
//	// need them repeated.
//	updateConfig := tgbotapi.NewUpdate(0)
//
//	// Tell Telegram we should wait up to 30 seconds on each request for an
//	// update. This way we can get information just as quickly as making many
//	// frequent requests without having to send nearly as many.
//	updateConfig.Timeout = 30
//
//	// Start polling Telegram for updates.
//	updates := bot.GetUpdatesChan(updateConfig)
//
//	// Let's go through each update that we're getting from Telegram.
//	for update := range updates {
//		// Telegram can send many types of updates depending on what your Bot
//		// is up to. We only want to look at messages for now, so we can
//		// discard any other updates.
//		if update.Message == nil {
//			continue
//		}
//
//		// Now that we know we've gotten a new message, we can construct a
//		// reply! We'll take the Chat ID and Text from the incoming message
//		// and use it to create a new message.
//		//text :=
//		//"https://www.google.com/search?q=" + update.Message.Text
//
//		//gameInfo := internal.GameInfo{
//		//	Name:        update.Message.Text,
//		//	TeseraLink:  url.URL{},
//		//	VkLinkBNI:   url.URL{},
//		//	AvitoLink:   url.URL{},
//		//	YoutubeLink: "",
//		//	GoogleLink:  url.URL{},
//		//	InfoFromTesera: struct {
//		//		Name                       string
//		//		Description                string
//		//		ImageUrl                   string
//		//		RecommendedAge             string
//		//		NumberOfPlayers            string
//		//		RecommendedNumberOfPlayers string
//		//		GameTime                   string
//		//	}{Name: "", Description: "", ImageUrl: "", RecommendedAge: "", NumberOfPlayers: "", RecommendedNumberOfPlayers: "", GameTime: ""},
//		//}
//
//		//gameInfo.GoogleLinkM(update.Message.Text)
//
//		//gameInfo.AllLinks(gameInfo.Name)
//		//fmt.Println(gameInfo)
//		//gameInfo.TeseraLinkM(gameInfo.Name)
//		//telegram.F()
//
//		//text, _ := telegram.GoogleLink(update.Message.Text)
//		//googleLink := "[Поиск игры в Google]" + "(" + gameInfo.GoogleLink.String() + ")" + "\n"
//		//teseraLink := "[Страница на Тесере]" + "(" + gameInfo.TeseraLink.String() + ")" + "\n"
//		//	fmt.Println(gameInfo.YoutubeLink.String())
//		//youtube := "[Youtube]" + "(" + gameInfo.YoutubeLink + ")" + "\n"
//		//avitoLink := "[Искать игру на Авито]" + "(" + gameInfo.AvitoLink.String() + ")" + "\n" + "\n"
//		//age := "**Рекомендуемый возраст:** " + gameInfo.InfoFromTesera.RecommendedAge + "\n"
//		//numberOfPlayers := "**Количество игроков:** " + gameInfo.InfoFromTesera.NumberOfPlayers + "\n"
//		//recNumberOfPlayers := "Рекомендуемое количество игроков: " + gameInfo.InfoFromTesera.RecommendedNumberOfPlayers + "\n"
//		//time := "Среднее время партии: " + gameInfo.InfoFromTesera.GameTime + "\n"
//		//teseraLink := "[Поиск игры в Tesera]" + "(" + gameInfo.TeseraLink.String() + ")"
//
//		//text := internal.MsgGameInfo(gameInfo) + internal.MsgRollingDice()
//
//		//image := `<a href="` + gameInfo.InfoFromTesera.ImageUrl + `"> </a>` + "\n"
//		//googleLink := `<a href="` + gameInfo.GoogleLink.String() + `">Поиск игры в Google </a>` + "\n"
//		//teseraLink := `<a href="` + gameInfo.TeseraLink.String() + `">Страница на Тесере </a>` + "\n"
//		////	fmt.Println(gameInfo.YoutubeLink.String())
//		//youtube := `<a href="` + gameInfo.YoutubeLink + `">Youtube </a>` + "\n"
//		//avitoLink := `<a href="` + gameInfo.AvitoLink.String() + `">Искать игру на Авито </a>` + "\n" + "\n"
//		//age := "Рекомендуемый возраст: " + gameInfo.InfoFromTesera.RecommendedAge + "\n"
//		//numberOfPlayers := "Количество игроков: " + gameInfo.InfoFromTesera.NumberOfPlayers + "\n"
//		//recNumberOfPlayers := "Рекомендуемое количество игроков: " + gameInfo.InfoFromTesera.RecommendedNumberOfPlayers + "\n"
//		//time := "Среднее время партии: " + gameInfo.InfoFromTesera.GameTime + "\n"
//
//		//fmt.Println(googleLink + teseraLink + youtube + avitoLink + age + numberOfPlayers + recNumberOfPlayers + time)
//		//msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
//		//msg.ParseMode = "markdown"
//
//		//msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
//		if update.Message.Text == "/dice" {
//			msgAnim := tgbotapi.NewAnimation(update.Message.Chat.ID, internal.GifDiceRolling())
//			if _, err := bot.Send(msgAnim); err != nil {
//				panic(err)
//			}
//			msg := tgbotapi.NewMessage(update.Message.Chat.ID, internal.MsgRollingDice())
//			if _, err := bot.Send(msg); err != nil {
//				panic(err)
//			}
//		} else {
//			msg := consumer.Consume(update.Message)
//			if _, err := bot.Send(msg); err != nil {
//				// Note that panics are a bad way to handle errors. Telegram can
//				// have service outages or network errors, you should retry sending
//				// messages or more gracefully handle failures.
//				panic(err)
//			}
//		}
//
//		//msg.ParseMode = "HTML"
//
//		// We'll also say that this message is a reply to the previous message.
//		// For any other specifications than Chat ID or Text, you'll need to
//		// set fields on the `MessageConfig`.
//		//msg.ReplyToMessageID = update.Message.MessageID
//
//		//data, _ := os.ReadFile("dice-roll.gif")
//		//gifka := tgbotapi.FileBytes{
//		//	Name:  "dice-roll.gif",
//		//	Bytes: data,
//		//}
//		//
//		//bot.Send(tgbotapi.NewAnimation(update.Message.Chat.ID, gifka))
//
//		// Okay, we're sending our message off! We don't care about the message
//		// we just sent, so we'll discard it.
//
//	}
//}
//
////
////func mustToken()
////string{
////token := flag.String("botToken", "", "token for access telegram Board Game bot")
////
////flag.Parse()
////
////if *token == ""{
////log.Fatal("token is not specified")
////}
////
////return *token
////}
