package game

import (
	"context"
	"errors"
	"fmt"
	"github.com/antihax/optional"
	"pwd/internal/service/dice"
	swagger "pwd/internal/service/game/tesera-swagger"
	"time"
)

const (
	basePathAPI           = "https://api.tesera.ru/"
	gameCount             = 6000
	offs                  = 6000
	tooOld                = 2000
	minNumVotesForNewGame = 2
	minNumVotesForOldGame = 15
)

type Service struct {
	client *swagger.APIClient
}

func New() *Service {
	client := swagger.NewAPIClient(&swagger.Configuration{BasePath: basePathAPI})
	return &Service{
		client: client,
	}
}

var ErrNoGameFound = errors.New("no game found")

func (s *Service) RandomGame(ctx context.Context, players *int64, timeToPlay *int64) (string, error) {
	offset := dice.Random(0, offs)

	options := &swagger.GamesApiGetOpts{
		Offset:       optional.NewInt32(int32(offset)),
		EmptyPicture: optional.NewBool(false),
	}

	//if timeToPlay != nil {
	//	//options.PlayTimeMax = optional.NewInt64(*timeToPlay)
	//}

	games, _, err := s.client.GamesApi.Get(ctx, options)
	if err != nil {
		return "", fmt.Errorf("failed to get list of games: %w", err)
	}

	if len(games) == 0 {
		return "", ErrNoGameFound
	}

	for _, game := range games {
		gameYearCreated := int(game.Year)
		currentYear := time.Now().Year()
		if timeToPlay != nil {
			if game.PlaytimeMax > *timeToPlay {
				continue
			}
		}
		if (gameYearCreated < currentYear-2 && game.NumVotes < minNumVotesForOldGame) || game.NumVotes < minNumVotesForNewGame {
			continue
		}
		if gameYearCreated > currentYear || gameYearCreated < tooOld {
			continue
		}
		if players != nil {
			if game.PlayersMin <= *players && game.PlayersMax >= *players {
				return game.Title, nil
			}
		}
	}

	return "", nil
}

func (s *Service) GameInfo(ctx context.Context, search string) (*swagger.GameInfo, error) {
	searchInfo, _, err := s.client.SearchApi.Get(ctx, &swagger.SearchApiGetOpts{Query: optional.NewString(search)})
	if err != nil {
		return nil, fmt.Errorf("failed to get list of games: %w", err)
	}
	if len(searchInfo) == 0 {
		return nil, errors.New("failed to find game")
	}

	game, _, err := s.client.GamesApi.Get_13(ctx, searchInfo[0].Alias, nil)

	if err != nil {
		return nil, fmt.Errorf("failed to get game info: %w", err)
	}

	return game.Game, nil
}
