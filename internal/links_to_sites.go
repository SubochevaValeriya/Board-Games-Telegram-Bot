package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

type TeseraSearchResponse []struct {
	Type     string `json:"type"`
	Value    string `json:"value"`
	Alias    string `json:"alias"`
	ID       int    `json:"id"`
	TeseraID int    `json:"teseraId"`
	Title    string `json:"title"`
	Title2   string `json:"title2"`
	PhotoURL string `json:"photoUrl"`
}

func (g *GameInfo) TeseraLinkM(s string) {
	//log.Println(s)
	response, err := http.Get("https://api.tesera.ru/search/games?query=%" + s + "&withAdditions=false&WaitHandle.Handle=%7B%7D%22%20")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject TeseraSearchResponse
	json.Unmarshal(responseData, &responseObject)

	//result, err := googlesearch.Search(nil, "https://www.google.com/search?q="+s+" tesera")

	//log.Println(err)
	if len(responseObject) == 0 {
		return
	}
	urlStr := "https://tesera.ru/game/" + responseObject[0].Alias + "/"

	if err == nil {
		url, err := url.Parse(urlStr)
		if err == nil {
			g.TeseraLink = *url
		} else {
			log.Println(err)
		}
	}
}

func (g *GameInfo) GoogleLinkM(s string) {
	url, err := url.Parse("https://www.google.com/search?q=" + s + " настольная игра")
	if err == nil {
		g.GoogleLink = *url
	}
}

func (g *GameInfo) AvitoLinkM(s string) {
	url, err := url.Parse("https://www.avito.ru/rossiya?q=" + s + " настольная игра")
	if err == nil {
		g.AvitoLink = *url
	}
}

func (g *GameInfo) VkLinkBNIM(s string) {
	url, err := url.Parse("https://vk.com/wall-114967596?owners_only=1&q=" + s)
	if err == nil {
		g.VkLinkBNI = *url
	}
}

func (g *GameInfo) YoutubeLinkM(s string) {
	g.YoutubeLink = "https://www.youtube.com/results?search_query=" + s

	//func (g *GameInfo) YoutubeLinkM(s string) {
	//	url, err := url.Parse("https://www.youtube.com/results?search_query=" + s)
	//	if err == nil {
	//		g.VkLinkBNI = *url
	//	}

}

func (g *GameInfo) AllLinks(s string) {
	g.GoogleLinkM(s)
	g.TeseraLinkM(s)
	g.AvitoLinkM(s)
	g.VkLinkBNIM(s)
	g.YoutubeLinkM(s)
	g.InfoFromTesera.TeseraParsing(g.TeseraLink.String())
}
