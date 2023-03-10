package internal

import (
	"context"
	"fmt"
	"github.com/fzerorubigd/gobgg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindTheGame(t *testing.T) {
	bgg := ConnectToBGGClient()
	result, _ := FindTheGame(bgg, "дюна")
	fmt.Println()
	results, _ := bgg.GetThings(context.TODO(), gobgg.GetThingIDs(3955))
	fmt.Println(results[0].Image)
	assert.Equal(t, "https://boardgamegeek.com/boardgame/3955", result)
}

func TestRandomGameLinkGenerator(t *testing.T) {
	bgg := ConnectToBGGClient()
	var result string
	for {
		result, err := RandomGame(bgg)
		fmt.Println(result, err)
		if result != "" {
			break
		}

	}
	fmt.Println(result)
}

func TestGameSearch(t *testing.T) {
	result, _ := GameSearch("дюн")
	//fmt.Println()
	//results, _ := bgg.GetThings(context.TODO(), gobgg.GetThingIDs(3955))
	//fmt.Println(results[0].Image)
	assert.Equal(t, "https://boardgamegeek.com/boardgame/3955", result)
}
