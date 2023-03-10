package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func (g *GameInfo) TeseraLinkM(s string) error {
	urlSearch := "https://api.tesera.ru/search/games?query=" + url.QueryEscape(s) + "&WaitHandle.Handle=%7B%7D"
	response, err := http.Get(urlSearch)

	if err != nil {
		fmt.Print(err.Error(), "get")
		return err
		//	os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return err
	}

	var responseObject TeseraSearchResponse
	json.Unmarshal(responseData, &responseObject)

	//result, err := googlesearch.Search(nil, "https://www.google.com/search?q="+s+" tesera")

	//log.Println(err)
	if len(responseObject) == 0 {
		log.Println("here")
		url, _ := url.Parse("https://tesera.ru/")
		g.TeseraLink = *url
		return errors.New("not found")
	}
	urlStr := "https://tesera.ru/game/" + responseObject[0].Alias + "/"

	if err == nil {
		url, err := url.Parse(urlStr)
		if err == nil {
			g.TeseraLink = *url
		} else {
			log.Println(err)
			return err
		}
	}
	g.Name = responseObject[0].Value
	log.Println(s, g.Name)
	return nil
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

	id, err := GameSearch(s)
	if err == nil {
		g.BGGLink = fmt.Sprintf("%s%v", BGGLinkToGame, id)
	}
}

func (g *GameInfo) AllLinks(s string) {
	g.GoogleLinkM(s)

	if g.TeseraLinkM(s) == nil {
		g.Info.TeseraParsing(g.TeseraLink.String())
	} else {
		id, err := GameSearch(s)
		if err != nil {
			g.Info.BGGParsing(id)
		}
	}
	g.BGGLinkM(s)
	g.AvitoLinkM(s)
	g.VkLinkBNIM(s)
	g.YoutubeLinkM(s)
}
