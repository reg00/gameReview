package port

import "github.com/Reg00/gameReview/internal/domain/models"

type Cacher interface {
	SetGame(game models.Game) error
	GetGameById(id int) (models.Game, error)
}
