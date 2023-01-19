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
	BGGLink        string
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
