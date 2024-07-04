package consumer

import (
	"context"
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"pwd/pkg/model"
	"pwd/pkg/service/game"
	"regexp"
	"strconv"
	"strings"
)

type Service struct {
	gameInfoService     gameInfoService
	linkService         linkService
	diceService         diceService
	googleSheetsService googleSheetsService
	bot                 *tgbotapi.BotAPI
}

func New(
	gameInfoService gameInfoService,
	linkService linkService,
	diceService diceService,
	googleSheetsService googleSheetsService,
	bot *tgbotapi.BotAPI,
) *Service {
	return &Service{
		gameInfoService:     gameInfoService,
		linkService:         linkService,
		diceService:         diceService,
		googleSheetsService: googleSheetsService,
		bot:                 bot,
	}
}

var (
	regexpNumberOfPlayer = regexp.MustCompile(`Выбранное количество игроков: \[(.*?)\]`)
	regexpTime           = regexp.MustCompile(`Выбранное время игры: \[(.*?)\]`)

	keyboardTime = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("30 минут", "30"),
			tgbotapi.NewInlineKeyboardButtonData("1 час", "60"),
			tgbotapi.NewInlineKeyboardButtonData("2 часа", "120"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("3 часа", "180"),
			tgbotapi.NewInlineKeyboardButtonData("4 часа", "240"),
			tgbotapi.NewInlineKeyboardButtonData("5 часов", "300"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Не имеет значения", "Не имеет значения"),
		),
	)

	keyboardPlayers = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("1", "1"),
			tgbotapi.NewInlineKeyboardButtonData("2", "2"),
			tgbotapi.NewInlineKeyboardButtonData("3", "3"),
			tgbotapi.NewInlineKeyboardButtonData("4", "4"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("5", "5"),
			tgbotapi.NewInlineKeyboardButtonData("6", "6"),
			tgbotapi.NewInlineKeyboardButtonData("7", "7"),
			tgbotapi.NewInlineKeyboardButtonData("8", "8"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("9 и более", "9"),
		),
	)

	keyboardOneMoreGame = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(msgOneMoreGame, msgOneMoreGame),
		),
	)
)

func (s *Service) ConsumeCallback(ctx context.Context, callbackQuery *tgbotapi.CallbackQuery) (tgbotapi.MessageConfig, error) {
	msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "")
	msg.ParseMode = "HTML"

	switch callbackQuery.Message.Text {
	case msgGameAskNumberOfPlayers:
		msg.Text = fmt.Sprintf(msgGameAskTimeForGameWithNumberOfPlayers, callbackQuery.Data)
		msg.ReplyMarkup = keyboardTime
	}

	if strings.Contains(callbackQuery.Message.Text, msgGameAskTimeForGame) {
		players := string(callbackQuery.Message.Text[strings.Index(callbackQuery.Message.Text, "[")+1])
		playersInt, err := strconv.ParseInt(players, 10, 64)
		if err != nil {
			return tgbotapi.MessageConfig{}, err
		}

		var timeToPlay *int64

		if callbackQuery.Data != "Не имеет значения" {
			time, err := strconv.ParseInt(callbackQuery.Data, 10, 64)
			if err != nil {
				return tgbotapi.MessageConfig{}, err
			}

			timeToPlay = &time
		}

		if msg.Text, err = s.gameAdviceMessage(ctx, playersInt, timeToPlay); err != nil {
			return tgbotapi.MessageConfig{}, err
		}

		msg.ReplyMarkup = keyboardOneMoreGame
	}

	if callbackQuery.Data == msgOneMoreGame {
		playersFirstIndex := strings.Index(callbackQuery.Message.Text, "Выбранное количество игроков: [")
		playersLastIndex := strings.Index(callbackQuery.Message.Text, "]")
		fmt.Println(callbackQuery.Message.Text[playersFirstIndex+len("Выбранное количество игроков: [") : playersLastIndex])
		playersInt, err := strconv.ParseInt(callbackQuery.Message.Text[playersFirstIndex+len("Выбранное количество игроков: ["):playersLastIndex], 10, 64)
		if err != nil {
			return tgbotapi.MessageConfig{}, err
		}

		var timeToPlay *int64
		timeToPlayFirstIndex := strings.LastIndex(callbackQuery.Message.Text, "Выбранное время игры: [")
		if timeToPlayFirstIndex != -1 {
			timeToPlayLastIndex := strings.LastIndex(callbackQuery.Message.Text, "]")
			time, err := strconv.ParseInt(callbackQuery.Message.Text[timeToPlayFirstIndex+len("Выбранное время игры: ["):timeToPlayLastIndex], 10, 64)
			if err != nil {
				return tgbotapi.MessageConfig{}, err
			}

			timeToPlay = &time
		}

		if msg.Text, err = s.gameAdviceMessage(ctx, playersInt, timeToPlay); err != nil {
			return tgbotapi.MessageConfig{}, err
		}

		msg.ReplyMarkup = keyboardOneMoreGame
	}

	return msg, nil
}

func (s *Service) Consume(ctx context.Context, message *tgbotapi.Message) (tgbotapi.MessageConfig, error) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "")
	msg.ParseMode = "HTML"
	switch message.Text {
	case "/hello":
		msg.Text = msgHello
	case "/help":
		msg.Text = msgCommandList
	case "/dice", "Бросить кубик!":
		gif, err := s.diceService.GifDiceRolling()
		if err != nil {
			return tgbotapi.MessageConfig{}, err
		}
		msgAnim := tgbotapi.NewAnimation(message.Chat.ID, gif)
		if _, err := s.bot.Send(msgAnim); err != nil {
			return tgbotapi.MessageConfig{}, err
		}
		msg.Text = msgRollingDice()
	case "/advise", "Во что поиграть?":
		msg.Text = msgGameAskNumberOfPlayers
		msg.ReplyMarkup = keyboardPlayers
	default:
		msg.ReplyToMessageID = message.MessageID
		msg.Text = s.gameInfoMessage(ctx, message.Text)
	}

	return msg, nil
}

func (s *Service) gameInfoMessage(ctx context.Context, title string) string {
	gameInfo := model.GameInfo{
		Title: title,
		Links: s.linkService.AllLinks(ctx, title),
	}

	info, err := s.gameInfoService.GameInfo(ctx, title)
	if err != nil || info == nil {
		return msgGameNotFound + msgShortInfo(gameInfo)
	}

	gameInfo.GameInfo = *info
	msg := msgGameInfo(gameInfo)

	anticafeGameInStock := s.googleSheetsService.IsGameInStockInAnticafe(*info)
	if len(anticafeGameInStock) != 0 {
		msg += msgAvailableInAnticafe
		for _, anticafe := range anticafeGameInStock {
			msg += fmt.Sprintf(msgAnticafeWithLink, anticafe.Link, anticafe.AnticafeName)
		}
		msg += "\n"
	}

	return msg
}

func (s *Service) gameAdviceMessage(ctx context.Context, numberOfPlayers int64, timeToPlay *int64) (string, error) {
	var (
		title string
		err   error
		text  string
	)

	for i := 0; i < 100; i++ {
		title, err = s.gameInfoService.RandomGame(ctx, &numberOfPlayers, timeToPlay)
		if err != nil && !errors.Is(err, game.ErrNoGameFound) {
			return "", err
		}
		if title != "" {
			break
		}
	}

	if title == "" {
		text += msgUserRequestNotFound
	} else {
		text += s.gameInfoMessage(ctx, title)
	}

	text += fmt.Sprintf(msgUserRequestNumberOfPlayers, numberOfPlayers)

	if timeToPlay != nil {
		text += fmt.Sprintf(msgUserRequestTimeOfPlay, *timeToPlay)
	}

	return text, nil
}
