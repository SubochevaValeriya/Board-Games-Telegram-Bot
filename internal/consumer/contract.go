package consumer

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"pwd/internal/service/anticafe"
	swagger "pwd/internal/service/game/tesera-swagger"
	"pwd/internal/service/links"
)

type gameInfoService interface {
	GameInfo(ctx context.Context, search string) (*swagger.GameInfo, error)
	RandomGame(ctx context.Context, players *int64, timeToPlay *int64) (string, error)
}

type linkService interface {
	AllLinks(ctx context.Context, search string) links.Links
}

type diceService interface {
	GifDiceRolling() (tgbotapi.FileBytes, error)
}

type googleSheetsService interface {
	IsGameInStockInAnticafe(game swagger.GameInfo) []anticafe.AnticafeGameInStock
}
