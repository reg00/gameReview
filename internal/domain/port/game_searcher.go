package port

import (
	"github.com/Reg00/gameReview/internal/domain/models"
)

type GameSearcher interface {
	GetGamesByName(offset int, limit int, name string) ([]models.Game, error)
	GetGameById(id int) (models.Game, error)
}
