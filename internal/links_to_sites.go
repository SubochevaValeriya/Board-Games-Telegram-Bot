package internal

import (
	googlesearch "github.com/rocketlaunchr/google-search"
	"net/url"
)

func (g *GameInfo) TeseraLinkM(s string) {
	result, err := googlesearch.Search(nil, "https://www.google.com/search?q="+s+" tesera")
	if err == nil {
		url, err := url.Parse(result[0].URL)
		if err == nil {
			g.TeseraLink = *url
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
