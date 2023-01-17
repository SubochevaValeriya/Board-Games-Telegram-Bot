package internal

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindTheGame(t *testing.T) {
	bgg := ConnectToBGGClient()
	result, _ := FindTheGame(bgg, "BANG!")

	assert.Equal(t, "https://boardgamegeek.com/boardgame/3955", result)
}

func TestRandomGameLinkGenerator(t *testing.T) {
	bgg := ConnectToBGGClient()
	var result string
	for {
		result, _ = RandomGame(bgg)
		if result != "" {
			fmt.Println(result)
			break
		}

	}
	fmt.Println(result)
}
