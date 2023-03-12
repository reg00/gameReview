package service

import (
	"github.com/Reg00/gameReview/internal/domain/dto"
	"github.com/Reg00/gameReview/internal/domain/port"
)

type GameReviewService struct {
	igdbClient port.GameSearcher
}

func NewGameReviewService(
	gameSearcher *port.GameSearcher,
) *GameReviewService {
	return &GameReviewService{
		igdbClient: *gameSearcher,
	}
}

func (grs *GameReviewService) GetGamesByName(offset int, limit int, name string) ([]dto.Game, error) {
	return grs.igdbClient.GetGamesByName(offset, limit, name)
}

func (grs *GameReviewService) GetGameById(id int) (dto.Game, error) {
	return grs.igdbClient.GetGameById(id)
}
