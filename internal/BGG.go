package internal

import (
	"context"
	"errors"
	"fmt"
	"github.com/fzerorubigd/gobgg"
	"net/url"
	"strings"
)

const BGGLinkToGame = "https://boardgamegeek.com/boardgame/"

var gameNotFound = errors.New("game not found")

func ConnectToBGGClient() *gobgg.BGG {
	return gobgg.NewBGGClient()
}

func FindTheGame(bggClient *gobgg.BGG, name string) (string, error) {
	results, err := bggClient.Search(context.TODO(), name)
	if err != nil {
		return "", err
	}
	for _, result := range results {
		if strings.ToLower(result.Name) == strings.ToLower(name) {
			fmt.Println(result)
			return fmt.Sprintf("%s%v", BGGLinkToGame, result.ID), nil
		}
	}
	return "", gameNotFound
}

func RandomGame(bggClient *gobgg.BGG) (string, error) {
	id := random(1, 120000)
	results, err := bggClient.GetThings(context.TODO(), gobgg.GetThingIDs(int64(id)))

	if err != nil {
		return "", err
	}

	var result gobgg.ThingResult
	fmt.Println(results)
	for _, result = range results {
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
		if g.TeseraLinkM(result.Name) != nil {
			return "", errors.New("no info on Tesera")
		}
	}

	return result.Name, nil
}
