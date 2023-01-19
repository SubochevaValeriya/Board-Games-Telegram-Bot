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
	random(min int, max int) int
}

func GifDiceRolling() (tgbotapi.FileBytes, error) {
	//	data, err := os.ReadFile("dice.gif") // from local file
	response, err := http.Get("https://i.pinimg.com/originals/48/ff/32/48ff322720bb1bb377da359adc66b0fb.gif")
	if err != nil {
		return tgbotapi.FileBytes{}, fmt.Errorf("error opening gif: %w", err)
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return tgbotapi.FileBytes{}, fmt.Errorf("error reading gif body: %w", err)
	}

	gif := tgbotapi.FileBytes{
		Name:  "dice.gif",
		Bytes: data,
	}

	return gif, err
}

func random(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	var result int
	result = rand.Intn(max-min+1) + min

	return result
}
