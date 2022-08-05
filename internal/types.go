package internal

import (
	"net/url"
)

type linkFetcher interface {
	AllLinks(s string)
}

type teseraFetcher interface {
	TeseraParsing(webUrl string)
}

type GameInfo struct {
	Name           string
	TeseraLink     url.URL
	VkLinkBNI      url.URL
	AvitoLink      url.URL
	YoutubeLink    string
	GoogleLink     url.URL
	InfoFromTesera InfoFromTesera
}

type InfoFromTesera struct {
	Name                       string
	Description                string
	ImageUrl                   string
	RecommendedAge             string
	NumberOfPlayers            string
	RecommendedNumberOfPlayers string
	GameTime                   string
}
