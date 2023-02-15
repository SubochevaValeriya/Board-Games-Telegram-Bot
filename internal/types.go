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
	Name        string
	TeseraLink  url.URL
	VkLinkBNI   url.URL
	AvitoLink   url.URL
	YoutubeLink string
	GoogleLink  url.URL
	BGGLink     string
	Info        Info
}

//type Info struct {
//	Name                       string
//	Description                string
//	ImageUrl                   string
//	RecommendedAge             string
//	NumberOfPlayers            string
//	RecommendedNumberOfPlayers string
//	GameTime                   string
//}

type Info struct {
	Name                       string `json:"name,omitempty"`
	YearPublished              int    `json:"year_published,omitempty"`
	Description                string `json:"description,omitempty"`
	Image                      string `json:"image,omitempty"`
	MinAge                     string `json:"min_age,omitempty"`
	NumberOfPlayers            string
	RecommendedNumberOfPlayers string
	PlayTime                   string  `json:"play_time,omitempty"`
	AverageRate                float64 `json:"average_rate,omitempty"`
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
