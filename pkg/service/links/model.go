package links

import "net/url"

type Links struct {
	TeseraLink  string
	VKLinkBNI   url.URL
	AvitoLink   url.URL
	YoutubeLink string
	GoogleLink  url.URL
	BGGLink     string
}

func (l *Links) googleLink(search string) {
	url, err := url.Parse("https://www.google.com/search?q=" + search + " настольная игра")
	if err != nil {
		return
	}

	l.GoogleLink = *url
}

func (l *Links) avitoLink(search string) {
	url, err := url.Parse("https://www.avito.ru/rossiya?q=" + search + " настольная игра")
	if err != nil {
		return
	}

	l.AvitoLink = *url
}

func (l *Links) vkLinkBNI(search string) {
	url, err := url.Parse("https://vk.com/wall-114967596?owners_only=1&q=" + url.QueryEscape(search))
	if err != nil {
		return
	}

	l.VKLinkBNI = *url
}

func (l *Links) youtubeLink(search string) {
	l.YoutubeLink = "https://www.youtube.com/results?search_query=" + search
}
