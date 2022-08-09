package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"pwd/internal"
)

func main() {
	//token := mustToken()

	//bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	//bot, err := tgbotapi.NewBotAPI("5501151328:AAFVVneFN6O4SLihdwM3qdOTxHmY8mtvNtQ")
	//if err != nil {
	//	panic(err)
	//}

	//telegram.SendMsg(telegram.ReceiveRequest(bot), bot)
	s := url.QueryEscape("Бэнг")
	//urlSearch := "https://api.tesera.ru/search/games?query=" + s + "&WaitHandle.Handle=%7B%7D"
	//_, err := http.Get("https://api.tesera.ru/search/games?query=" + s + "&WaitHandle.Handle=%7B%7D")
	TeseraLinkMtest(s)
	fmt.Println(TeseraLinkMtest(s))

}

func TeseraLinkMtest(s string) string {
	log.Println(s)
	urlSearch := "https://api.tesera.ru/search/games?query=" + s + "&WaitHandle.Handle=%7B%7D"
	response, err := http.Get(urlSearch)
	//response, err := http.Get("https://api.tesera.ru/search/games?query=&withAdditions=false&WaitHandle.Handle=%7B%7D%22%20/")

	if err != nil {
		fmt.Print(err.Error(), "get")
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject internal.TeseraSearchResponse
	json.Unmarshal(responseData, &responseObject)
	fmt.Println(responseObject)

	//result, err := googlesearch.Search(nil, "https://www.google.com/search?q="+s+" tesera")

	//log.Println(err)
	if len(responseObject) == 0 {
		return ""
	}
	urlStr := "https://tesera.ru/game/" + responseObject[0].Alias + "/"
	return urlStr
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
