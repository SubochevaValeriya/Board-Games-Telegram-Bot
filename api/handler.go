package handler

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var bot *tgbotapi.BotAPI

func init() {
	var err error
	bot, err = tgbotapi.NewBotAPI("5501151328:AAFVVneFN6O4SLihdwM3qdOTxHmY8mtvNtQ")
	if err != nil {
		log.Panic(err)
	}
}

type Hito struct {
	URI   string
	Title string
	fecha time.Time
}

type Response struct {
	Msg    string `json:"text"`
	ChatID int64  `json:"chat_id"`
	Method string `json:"method"`
}

var hitos = []Hito{
	Hito{
		URI:   "0.Repositorio",
		Title: "Datos básicos y repo",
		fecha: time.Date(2020, time.September, 29, 11, 30, 0, 0, time.UTC),
	},
	Hito{
		URI:   "1.Infraestructura",
		Title: "HUs y entidad principal",
		fecha: time.Date(2020, time.October, 6, 11, 30, 0, 0, time.UTC),
	},
	Hito{
		URI:   "2.Tests",
		Title: "Tests iniciales",
		fecha: time.Date(2020, time.October, 16, 11, 30, 0, 0, time.UTC),
	},
	Hito{
		URI:   "3.Contenedores",
		Title: "Contenedores",
		fecha: time.Date(2020, time.October, 26, 11, 30, 0, 0, time.UTC),
	},
	Hito{
		URI:   "4.CI",
		Title: "Integración continua",
		fecha: time.Date(2020, time.November, 6, 23, 59, 0, 0, time.UTC),
	},
	Hito{
		URI:   "5.Serverless",
		Title: "Trabajando con funciones serverless",
		fecha: time.Date(2020, time.November, 24, 11, 30, 0, 0, time.UTC),
	},
}

func Handler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	var update tgbotapi.Update
	if err := json.Unmarshal(body, &update); err != nil {
		log.Fatal("Error en el update →", err)
	}
	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	currentTime := time.Now()
	var next int
	var queda time.Duration
	for indice, hito := range hitos {
		if hito.fecha.After(currentTime) {
			next = indice
			queda = hito.fecha.Sub(currentTime)
		}
	}
	if update.Message.IsCommand() {
		text := ""
		if next == 0 {
			text = "Ninguna entrega próxima"
		} else {

			switch update.Message.Command() {
			case "kk":
				text = queda.String()
			case "kekeda":
				text = fmt.Sprintf("→ Próximo hito %s\n🔗 https://jj.github.io/IV/documentos/proyecto/%s\n📅 %s",
					hitos[next].Title,
					hitos[next].URI,
					hitos[next].fecha.String(),
				)
			default:
				text = "Usa /kk para lo que queda para el próximo hito, /kekeda para + detalle"
			}
		}
		data := Response{Msg: text,
			Method: "sendMessage",
			ChatID: update.Message.Chat.ID}

		msg, _ := json.Marshal(data)
		log.Printf("Response %s", string(msg))
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, string(msg))
	}
}

//func t(w http.ResponseWriter, r *http.Request) {
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
//		gameInfo := internal.GameInfo{
//			Name:        update.Message.Text,
//			TeseraLink:  url.URL{},
//			VkLinkBNI:   url.URL{},
//			AvitoLink:   url.URL{},
//			YoutubeLink: "",
//			GoogleLink:  url.URL{},
//			Image:       url.URL{},
//		}
//
//		//gameInfo.GoogleLinkM(update.Message.Text)
//
//		gameInfo.AllLinks(gameInfo.Name)
//		fmt.Println(gameInfo)
//		//gameInfo.TeseraLinkM(gameInfo.Name)
//		//internal.F()
//
//		//text, _ := internal.GoogleLink(update.Message.Text)
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
//		image := `<a href="` + gameInfo.InfoFromTesera.ImageUrl + `"> </a>` + "\n"
//		googleLink := `<a href="` + gameInfo.GoogleLink.String() + `">Поиск игры в Google </a>` + "\n"
//		teseraLink := `<a href="` + gameInfo.TeseraLink.String() + `">Страница на Тесере </a>` + "\n"
//		//	fmt.Println(gameInfo.YoutubeLink.String())
//		youtube := `<a href="` + gameInfo.YoutubeLink + `">Youtube </a>` + "\n"
//		avitoLink := `<a href="` + gameInfo.AvitoLink.String() + `">Искать игру на Авито </a>` + "\n" + "\n"
//		age := "Рекомендуемый возраст: " + gameInfo.InfoFromTesera.RecommendedAge + "\n"
//		numberOfPlayers := "Количество игроков: " + gameInfo.InfoFromTesera.NumberOfPlayers + "\n"
//		recNumberOfPlayers := "Рекомендуемое количество игроков: " + gameInfo.InfoFromTesera.RecommendedNumberOfPlayers + "\n"
//		time := "Среднее время партии: " + gameInfo.InfoFromTesera.GameTime + "\n"
//
//		fmt.Println(googleLink + teseraLink + youtube + avitoLink + age + numberOfPlayers + recNumberOfPlayers + time)
//		msg := tgbotapi.NewMessage(update.Message.Chat.ID, image+googleLink+teseraLink+youtube+avitoLink+age+numberOfPlayers+recNumberOfPlayers+time)
//		//msg.ParseMode = "markdown"
//
//		//msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
//		msg.ParseMode = "HTML"
//
//		// We'll also say that this message is a reply to the previous message.
//		// For any other specifications than Chat ID or Text, you'll need to
//		// set fields on the `MessageConfig`.
//		msg.ReplyToMessageID = update.Message.MessageID
//
//		// Okay, we're sending our message off! We don't care about the message
//		// we just sent, so we'll discard it.
//		if _, err := bot.Send(msg); err != nil {
//			// Note that panics are a bad way to handle errors. Telegram can
//			// have service outages or network errors, you should retry sending
//			// messages or more gracefully handle failures.
//			panic(err)
//		}
//	}
//}
