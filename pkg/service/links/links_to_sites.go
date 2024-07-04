package links

import (
	"context"
	"strconv"
)

type Service struct {
	gameInfoService gameInfoService
}

func New(gameInfoService gameInfoService) *Service {
	return &Service{gameInfoService: gameInfoService}
}

func (s *Service) AllLinks(ctx context.Context, search string) Links {
	links := Links{}
	links.googleLink(search)
	links.avitoLink(search)
	links.vkLinkBNI(search)
	links.youtubeLink(search)

	gameInfo, err := s.gameInfoService.GameInfo(ctx, search)
	if err == nil {
		links.TeseraLink = gameInfo.TeseraUrl
		links.BGGLink = "https://boardgamegeek.com/boardgame/" + strconv.FormatInt(gameInfo.BggId, 10)
	}

	return links
}
