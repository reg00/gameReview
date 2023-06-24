package service

import (
	"fmt"

	"github.com/Reg00/gameReview/internal/domain/dto"
	"github.com/Reg00/gameReview/internal/domain/dto/httperr"
	"github.com/Reg00/gameReview/internal/domain/models"
	"github.com/Reg00/gameReview/internal/domain/port"
	"github.com/pkg/errors"
)

type ReviewService struct {
	storage port.Storager
	gs      port.GameSearcher
}

func NewReviewService(
	storage *port.Storager,
	gs *port.GameSearcher,
) *ReviewService {
	return &ReviewService{
		storage: *storage,
		gs:      *gs,
	}
}

func (rs *ReviewService) AddReview(addReview *models.AddReview) (*models.GetReview, error) {

	if addReview.Rate < 0 || addReview.Rate > 10 {
		return nil, errors.Wrap(httperr.ErrBadRequest, fmt.Sprintf("ReviewService: %s", errors.New("Rate must be between 0 and 10")))
	}

	dtoReview, err := rs.storage.AddReview(dto.ConvertToDto(addReview))
	if err != nil {
		return nil, err
	}

	review := dtoReview.Convert()
	game, err := rs.gs.GetGameById(dtoReview.GameID)
	if err != nil {
		return nil, err
	}

	review.Game = game

	return review, nil
}

func (rs *ReviewService) GetReviewById(id int) (*models.GetReview, error) {
	dtoReview, err := rs.storage.GetReviewById(id)
	if err != nil {
		return nil, err
	}

	if dtoReview == nil {
		return nil, errors.Wrap(httperr.ErrNotFound, fmt.Sprintf("GORM: %s", errors.New(fmt.Sprintf("Cannot find review with %d", id))))
	}

	review := dtoReview.Convert()
	game, err := rs.gs.GetGameById(dtoReview.GameID)
	if err != nil {
		return nil, err
	}

	review.Game = game

	return review, nil
}