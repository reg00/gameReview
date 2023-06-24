package service

import (
	"github.com/Reg00/gameReview/internal/domain/models"
	"github.com/Reg00/gameReview/internal/domain/port"
)

type GameService struct {
	igdbClient port.GameSearcher
}

func NewGameService(
	gameSearcher *port.GameSearcher,
) *GameService {
	return &GameService{
		igdbClient: *gameSearcher,
	}
}

func (gs *GameService) GetGamesByName(offset int, limit int, name string) ([]models.Game, error) {
	games, err := gs.igdbClient.GetGamesByName(offset, limit, name)

	if err != nil {
		return []models.Game{}, err
	}

	return games, nil
}

func (gs *GameService) GetGameById(id int) (models.Game, error) {
	game, err := gs.igdbClient.GetGameById(id)

	if err != nil {
		return models.Game{}, err
	}

	return game, nil
}
