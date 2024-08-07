package links

import (
	"context"
	swagger "pwd/pkg/service/game/tesera-swagger"
)

type gameInfoService interface {
	GameInfo(ctx context.Context, search string) (*swagger.GameInfo, error)
}
