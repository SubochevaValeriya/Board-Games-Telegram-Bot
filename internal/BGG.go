package internal

import (
	"context"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/fzerorubigd/gobgg"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const BGGLinkToGame = "https://boardgamegeek.com/boardgame/"

var gameNotFound = errors.New("game not found")

func ConnectToBGGClient() *gobgg.BGG {
	return gobgg.NewBGGClient()
}

func GameSearch(name string) (string, error) {
	response, err := http.Get("https://boardgamegeek.com/geeksearch.php?action=search&objecttype=boardgame&q=" + url.QueryEscape(name))

	if err != nil {
		log.Println(err)
	} else if response.StatusCode == 200 {
		fmt.Println("We can scrape this")
	} else {
		log.Fatalln("Do not scrape this")
	}

	document, err := goquery.NewDocumentFromReader(response.Body)

	if err != nil {
		log.Println(err)
	}

	document.Find("div").Each(func(index int, selector *goquery.Selection) {
		//fmt.Println(selector.Text())
		if selector.AttrOr("id", "") == "results_objectname1" {
			fmt.Println(selector.Find("a").AttrOr("href", ""))
			//fmt.Println(selector.Text())
		}
	})
	//	Each(func(index int, selector *goquery.Selection) {
	//		selector.Find("img").AttrOr("title", "")
	//		{
	//case "возраст":
	//	t.RecommendedAge = selector.Text()
	//	case "число игроков":
	//	t.NumberOfPlayers = selector.Text()
	//	case "рекомендуемое число игроков":
	//	t.RecommendedNumberOfPlayers = selector.Text()
	//	case "время партии":
	//	t.GameTime = selector.Text()
	//	}
	//	fmt.Println(selector.Text())
	//t.Name = strings.Replace(selector.Text(), "| Tesera", "", -1)
	//})

	//results, err := bggClient.Search(context.TODO(), name)
	//if err != nil {
	//	return "", err
	//}
	//for _, result := range results {
	//	if strings.ToLower(result.Name) == strings.ToLower(name) {
	//		return fmt.Sprintf("%s%v", BGGLinkToGame, result.ID), nil
	//	}
	//}

	return "", gameNotFound
}

func FindTheGame(bggClient *gobgg.BGG, name string) (string, error) {
	results, err := bggClient.Search(context.TODO(), name)
	if err != nil {
		return "", err
	}
	for _, result := range results {
		if strings.ToLower(result.Name) == strings.ToLower(name) {
			return fmt.Sprintf("%s%v", BGGLinkToGame, result.ID), nil
		}
	}

	return "", gameNotFound
}

func RandomGame(bggClient *gobgg.BGG) (string, error) {
	id := random(1, 320000)
	results, err := bggClient.GetThings(context.TODO(), gobgg.GetThingIDs(int64(id)))

	if err != nil {
		return "", err
	}
	log.Println(results)

	if len(results) == 0 {
		return "", errors.New("game not found")
	}

	var result gobgg.ThingResult
	for _, result = range results {
		log.Println(result.Name)
		if result.Type != gobgg.BoardGameType {
			return "", errors.New("it's not board game")
		}

		if result.YearPublished < 2000 {
			return "", errors.New("too old")
		}

		if result.UsersOwned < 100 {
			return "", errors.New("not famous")
		}

		g := GameInfo{
			Name:           result.Name,
			TeseraLink:     url.URL{},
			VkLinkBNI:      url.URL{},
			AvitoLink:      url.URL{},
			YoutubeLink:    "",
			GoogleLink:     url.URL{},
			BGGLink:        "",
			InfoFromTesera: InfoFromTesera{},
		}

		log.Println(result.Name)
		if g.TeseraLinkM(result.Name) != nil {
			return "", errors.New("no info on Tesera")
		} else {
		}
	}

	return result.Name, nil
}
