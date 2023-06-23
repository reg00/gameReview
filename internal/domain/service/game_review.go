package service

import (
	"github.com/Reg00/gameReview/internal/domain/models"
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

func (grs *GameReviewService) GetGamesByName(offset int, limit int, name string) ([]models.Game, error) {
	games, err := grs.igdbClient.GetGamesByName(offset, limit, name)

	if err != nil {
		return []models.Game{}, err
	}

	return games, nil
}

func (grs *GameReviewService) GetGameById(id int) (models.Game, error) {
	game, err := grs.igdbClient.GetGameById(id)

	if err != nil {
		return models.Game{}, err
	}

	return game, nil
}
