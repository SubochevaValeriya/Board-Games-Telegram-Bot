package internal

import (
	"fmt"
)

const (
	MsgHello        = "Привет! Этот бот поможет найти информацию о настольной игре, а так же он может выполнить ещё несколько интересных функций 😼"
	MsgCommandList  = "Список команд: /hello - Приветственное сообщение /dice - Кинуть кубик"
	MsgGameNotFound = "Страница игры не найдена, но вы всё ещё можете поискать её в Google 🙃"
)

func MsgGameInfo(info GameInfo) string {
	return fmt.Sprintf(`
<a href="%s"> </a> 
Название: %s

Краткое описание: %s

🕑 Среднее время партии: %s
👨‍👩‍👧‍👦 Рекомендуемый возраст: %s
👥 Количество игроков: %s
👥👥 Рекомендуемое количество игроков: %s

<a href="%s">🔍 Поиск игры в Google </a> 
<a href="%s">♣️ Страница на Тесере </a> 
<a href="%s">⏯️ Youtube </a> 
<a href="%s">💸 Искать игру на Авито </a>
<a href="%s">📘 Искать игру в БНИ </a>`,

		info.InfoFromTesera.ImageUrl,
		info.InfoFromTesera.Name,
		info.InfoFromTesera.Description,
		info.InfoFromTesera.GameTime,
		info.InfoFromTesera.RecommendedAge,
		info.InfoFromTesera.NumberOfPlayers,
		info.InfoFromTesera.RecommendedNumberOfPlayers,
		info.GoogleLink.String(),
		info.TeseraLink.String(),
		info.YoutubeLink,
		info.AvitoLink.String(),
		info.VkLinkBNI.String(),
	)
}

func MsgShortInfo(info GameInfo) string {
	return fmt.Sprintf(`
<a href="%s">🔍 Поиск в Google </a>
<a href="%s">⏯ Youtube </a>
<a href="%s">💸 Искать на Авито </a>`,
		info.GoogleLink.String(),
		info.YoutubeLink,
		info.AvitoLink.String(),
	)
}

func MsgRollingDice() string {
	var msgDice string
	rollResult := diceRolling()
	switch rollResult {
	case 1:
		msgDice = "Критический провал! Повезёт в следующий раз ;)"
	case 6:
		msgDice = "А ты хорош 🍀"
	}

	return fmt.Sprintf(`🎲 Результат броска: %v 
%s`,
		rollResult, msgDice)
}
