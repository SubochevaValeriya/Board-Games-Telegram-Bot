package pkg

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
)
import "log"
import "fmt"

func (i *Info) TeseraParsing(webUrl string) error {
	response, err := http.Get(webUrl)
	if err != nil {
		log.Println(err)
		return err
	} else if response.StatusCode == 200 {
		fmt.Println("We can scrape this")
	} else {
		log.Println("Do not scrape this")
		return err
	}

	document, err := goquery.NewDocumentFromReader(response.Body)

	if err != nil {
		log.Println(err)
		return err
	}

	document.Find("title").Each(func(index int, selector *goquery.Selection) {
		i.Name = strings.Replace(selector.Text(), "| Tesera", "", -1)
	})

	document.Find(".raw_text_output").Each(func(index int, selector *goquery.Selection) {
		i.Description = strings.TrimSpace(selector.Text())
	})

	document.Find(".colleft").Each(func(index int, selector *goquery.Selection) {
		i.Image = "https://tesera.ru/" + selector.Find("img").AttrOr("src", "")
	})

	document.Find("li").Each(func(index int, selector *goquery.Selection) {
		switch selector.Find("img").AttrOr("title", "") {
		case "возраст":
			i.MinAge = selector.Text()
		case "число игроков":
			i.NumberOfPlayers = selector.Text()
		case "рекомендуемое число игроков":
			i.RecommendedNumberOfPlayers = selector.Text()
		case "время партии":
			i.PlayTime = selector.Text()
		}
	})

	return nil
}

//func ()  {
//	client := http.Client{}
//	params := url.Values{}
//	params.Add("param1", "value1")
//	params.Add("param2", "value2")
//	resp, err := client.Get("https://api.tesera.ru/games")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer resp.Body.Close()
//	io.Copy(os.Stdout, resp.Body)
//}
