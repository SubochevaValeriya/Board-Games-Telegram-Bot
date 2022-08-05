package internal

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"math/rand"
	"os"
	"time"
)

type diceRoller interface {
	GifDiceRolling() tgbotapi.FileBytes
	diceRolling() int
}

func GifDiceRolling() tgbotapi.FileBytes {
	data, _ := os.ReadFile("dice.gif")
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
