package internal

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io"
	"math/rand"
	"net/http"
	"time"
)

type diceRoller interface {
	GifDiceRolling() (tgbotapi.FileBytes, error)
	diceRolling() int
}

func GifDiceRolling() tgbotapi.FileBytes {
	//	data, _ := os.ReadFile("dice.gif")
	response, err := http.Get("https://media.giphy.com/media/l0IukZ771bIyeJcS4/giphy.gif")
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	data, _ := io.ReadAll(response.Body)

	gif := tgbotapi.FileBytes{
		Name:  "dice.gif",
		Bytes: data,
	}

	return gif
}

func diceRolling() int {
	rand.Seed(time.Now().UnixNano())
	var result int
	min := 1
	max := 6
	result = rand.Intn(max-min+1) + min

	return result
}