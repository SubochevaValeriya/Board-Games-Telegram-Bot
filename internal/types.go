package internal

import (
	"net/url"
)

type GameInfo struct {
	Name           string
	TeseraLink     url.URL
	VkLinkBNI      url.URL
	AvitoLink      url.URL
	YoutubeLink    string
	GoogleLink     url.URL
	Image          url.URL
	InfoFromTesera InfoFromTesera
}

type InfoFromTesera struct {
	Name                       string
	RecommendedAge             string
	NumberOfPlayers            string
	RecommendedNumberOfPlayers string
	GameTime                   string
	Description                string
	ImageUrl                   string
}
