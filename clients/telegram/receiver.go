package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type receiver interface {
	ReceiveRequest(bot *tgbotapi.BotAPI) tgbotapi.UpdatesChannel
}

func (s *Service) ReceiveRequest() tgbotapi.UpdatesChannel {
	s.bot.Debug = true

	updateConfig := tgbotapi.NewUpdate(0)

	updateConfig.Timeout = 30

	updates := s.bot.GetUpdatesChan(updateConfig)

	return updates
}
