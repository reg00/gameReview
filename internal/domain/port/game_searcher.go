package port

import (
	"github.com/Reg00/gameReview/internal/domain/dto"
)

type GameSearcher interface {
	GetGamesByName(offset int, limit int, name string) ([]dto.Game, error)
}
