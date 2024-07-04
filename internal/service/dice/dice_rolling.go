package dice

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
}

type Service struct {
}

func New() *Service {
	return &Service{}
}

var catGifs = map[int]string{
	0: "https://media1.tenor.com/m/Igw1jy1xv7kAAAAC/orange-cat.gif",
	1: "https://media1.tenor.com/m/SjyEg0nYnqIAAAAC/rolling-cat-fat-cats.gif",
	2: "https://media1.tenor.com/m/ykqUd0TmnbIAAAAC/cat-anime.gif",
	3: "https://media1.tenor.com/m/pXnh-svQzWgAAAAC/v-roll-koreas-cat.gif",
	4: "https://media1.tenor.com/m/3UmraKZO-1QAAAAC/rolling-cat-cat-rolling.gif",
	5: "https://media1.tenor.com/m/wsobmzpjvugAAAAC/rolling-cat-cat-rolling.gif",
	6: "https://media1.tenor.com/m/JB86G9GUJMsAAAAC/rolling-cat-cat-rolling.gif",
	7: "https://media1.tenor.com/m/cra3fK-yPosAAAAC/siamese-cat-siamese.gif",
}

const baseGif = "https://i.pinimg.com/originals/48/ff/32/48ff322720bb1bb377da359adc66b0fb.gif"

func (s *Service) GifDiceRolling() (tgbotapi.FileBytes, error) {
	//	data, err := os.ReadFile("dice.gif") // from local file
	response, err := http.Get(catGifs[Random(0, 7)])
	if err != nil {
		return tgbotapi.FileBytes{}, fmt.Errorf("error opening gif: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		response, err = http.Get(baseGif)
		if err != nil {
			return tgbotapi.FileBytes{}, fmt.Errorf("error opening gif: %w", err)
		}
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

func Random(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	var result int
	result = rand.Intn(max-min+1) + min

	return result
}
