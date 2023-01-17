package internal

import (
	"context"
	"errors"
	"fmt"
	"github.com/fzerorubigd/gobgg"
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
	}

	return result.Name, nil
}
