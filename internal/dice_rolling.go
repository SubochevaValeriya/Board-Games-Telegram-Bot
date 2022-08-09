package internal

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"math/rand"
	"time"
)

type diceRoller interface {
	GifDiceRolling() (tgbotapi.FileBytes, error)
	diceRolling() int
}

func GifDiceRolling() tgbotapi.FileBytes {
	//data, _ := os.ReadFile("dice.gif")
	//var data []byte
	//response, err := http.Get("http://your.url.com/whatever.html")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//defer response.Body.Close()
	//
	//// Copy data from the response to standard output
	//n, err1 := io.Copy(data, response.Body) //use package "io" and "os"
	//if err != nil {
	//	fmt.Println(err1)
	//	return
	//}
	//
	//gif := tgbotapi.FileBytes{
	//	Name:  "dice.gif",
	//	Bytes: data,
	//}

	gif := tgbotapi.FileBytes{
		Name:  "",
		Bytes: nil,
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
