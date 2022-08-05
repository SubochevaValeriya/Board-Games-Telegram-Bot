package internal

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
)
import "log"
import "fmt"

func (t *InfoFromTesera) TeseraParsing(webUrl string) {
	response, err := http.Get(webUrl)

	if err != nil {
		log.Fatalln(err)
	} else if response.StatusCode == 200 {
		fmt.Println("We can scrape this")
	} else {
		log.Fatalln("Do not scrape this")
	}

	document, err := goquery.NewDocumentFromReader(response.Body)

	if err != nil {
		log.Fatalln(err)
	}

	document.Find("title").Each(func(index int, selector *goquery.Selection) {
		t.Name = strings.Replace(selector.Text(), "| Tesera", "", -1)
	})

	document.Find(".raw_text_output").Each(func(index int, selector *goquery.Selection) {
		t.Description = strings.TrimSpace(selector.Text())
	})

	document.Find(".colleft").Each(func(index int, selector *goquery.Selection) {
		t.ImageUrl = "https://tesera.ru/" + selector.Find("img").AttrOr("src", "")
	})

	document.Find("li").Each(func(index int, selector *goquery.Selection) {
		switch selector.Find("img").AttrOr("title", "") {
		case "возраст":
			t.RecommendedAge = selector.Text()
		case "число игроков":
			t.NumberOfPlayers = selector.Text()
		case "рекомендуемое число игроков":
			t.RecommendedNumberOfPlayers = selector.Text()
		case "время партии":
			t.GameTime = selector.Text()
		}
	})

}

//if selector.Find("img").AttrOr("title", "") == "возраст" {
//t.RecommendedAge = selector.Text()
//}
//
//if selector.Find("img").AttrOr("title", "") == "число игроков" {
//t.RecommendedAge = selector.Text()
//}
//
//if selector.Find("img").AttrOr("title", "") == "рекомендуемое число игроков" {
//t.RecommendedAge = selector.Text()
//}
//
//if selector.Find("img").AttrOr("title", "") == "время партии" {
//t.RecommendedAge = selector.Text()
//}
