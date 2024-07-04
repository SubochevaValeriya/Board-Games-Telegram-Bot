package consumer

import (
	"fmt"
	"pwd/pkg/model"
	"pwd/pkg/service/dice"
)

const (
	msgHello                                 = "Привет! Этот бот поможет найти информацию о настольной игре, а так же он может выполнить ещё несколько интересных функций 😼"
	msgCommandList                           = "Список команд: /hello - Приветственное сообщение /dice - Кинуть кубик"
	msgGameNotFound                          = "Страница игры не найдена, но вы всё ещё можете поискать её в Google 🙃"
	msgGameAskTimeForGameWithNumberOfPlayers = "Выбранное количество игроков: [%s]. Сколько времени есть на игру?"
	msgGameAskTimeForGame                    = "Сколько времени есть на игру?"
	msgGameAskNumberOfPlayers                = "Количество игроков?"
	msgUserRequestNumberOfPlayers            = `
Выбранное количество игроков: [%d].`
	msgUserRequestTimeOfPlay = `
Выбранное время игры: [%d].
`
	msgUserRequestNotFound = "Простите, нам не удалось найти игру, подходящую под ваши параметры. Попробуйте их изменить."

	msgOneMoreGame         = "Посоветуй ещё игру!"
	msgAvailableInAnticafe = `
Наличие в антикафе:`
	msgAnticafeWithLink = `
<a href="%s"> %s </a>`
)

func msgGameInfo(info model.GameInfo) string {
	return fmt.Sprintf(`
<a href="%s"> </a> 
Название: %s

Краткое описание: %s

📆 Год выпуска: %d
🕑 Среднее время партии: %d мин.
⏳ Время на изучение правил: %d мин.
👨‍👩‍👧‍👦 Рекомендуемый возраст: %d
👥 Количество игроков: %d-%d
👥👥 Рекомендуемое количество игроков: %d-%d
📊 Рейтинг пользователей: %.2f/10

<a href="%s">🔍 Поиск игры в Google </a> 
<a href="%s">♣️ Страница на Тесере </a> 
<a href="%s">🕶️ Страница на BGG </a> 
<a href="%s">⏯️ Youtube </a> 
<a href="%s">💸 Искать игру на Авито </a>
<a href="%s">📘 Искать игру в БНИ </a>
`,

		info.PhotoUrl,
		info.Title,
		info.DescriptionShort,
		info.Year,
		info.PlaytimeMax,
		info.TimeToLearn,
		info.PlayersAgeMin,
		info.PlayersMin,
		info.PlayersMax,
		info.PlayersMinRecommend,
		info.PlayersMaxRecommend,
		info.RatingUser,
		info.Links.GoogleLink.String(),
		info.Links.TeseraLink,
		info.Links.BGGLink,
		info.Links.YoutubeLink,
		info.Links.AvitoLink.String(),
		info.Links.VKLinkBNI.String(),
	)
}

func msgShortInfo(info model.GameInfo) string {
	return fmt.Sprintf(`
<a href="%s">🔍 Поиск в Google </a>
<a href="%s">⏯ Youtube </a>
<a href="%s">💸 Искать на Авито </a>`,
		info.Links.GoogleLink.String(),
		info.Links.YoutubeLink,
		info.Links.AvitoLink.String(),
	)
}

func msgRollingDice() string {
	var msgDice string
	rollResult := dice.Random(1, 6)
	switch rollResult {
	case 1:
		msgDice = "Критический провал! Повезёт в следующий раз ;)"
	case 6:
		msgDice = "Удача на твоей стороне 🍀"
	}

	return fmt.Sprintf(`🎲 Результат броска: %v 
%s`,
		rollResult, msgDice)
}
