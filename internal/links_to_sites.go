package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func (g *GameInfo) TeseraLinkM(s string) {
	log.Println(s)
	urlSearch := "https://api.tesera.ru/search/games?query=" + url.QueryEscape(s) + "&WaitHandle.Handle=%7B%7D"
	response, err := http.Get(urlSearch)

	if err != nil {
		fmt.Print(err.Error(), "get")
		//	os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}

	var responseObject TeseraSearchResponse
	json.Unmarshal(responseData, &responseObject)

	//result, err := googlesearch.Search(nil, "https://www.google.com/search?q="+s+" tesera")

	//log.Println(err)
	if len(responseObject) == 0 {
		url, _ := url.Parse("https://tesera.ru/")
		g.TeseraLink = *url
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
	g.Name = responseObject[0].Value
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
	url, err := url.Parse("https://vk.com/wall-114967596?owners_only=1&q=" + url.QueryEscape(s))
	if err == nil {
		g.VkLinkBNI = *url
	}
}

func (g *GameInfo) YoutubeLinkM(s string) {
	g.YoutubeLink = "https://www.youtube.com/results?search_query=" + s
}

func (g *GameInfo) BGGLinkM(s string) {

	url, err := FindTheGame(ConnectToBGGClient(), s)
	if err == nil {
		g.BGGLink = url
	}
}

func (g *GameInfo) AllLinks(s string) {
	g.GoogleLinkM(s)
	g.TeseraLinkM(s)
	g.AvitoLinkM(s)
	g.VkLinkBNIM(s)
	g.YoutubeLinkM(s)
	g.BGGLinkM(g.Name)
	g.InfoFromTesera.TeseraParsing(g.TeseraLink.String())
}
