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
	cache   port.Cacher
}

func NewReviewService(
	storage *port.Storager,
	gs *port.GameSearcher,
	cache *port.Cacher,
) *ReviewService {
	return &ReviewService{
		storage: *storage,
		gs:      *gs,
		cache:   *cache,
	}
}

func (rs *ReviewService) AddReview(addReview *models.AddReview) (*models.GetReview, error) {

	if addReview.Rate < 0 || addReview.Rate > 10 {
		return nil, errors.Wrap(httperr.ErrBadRequest, fmt.Sprintf("ReviewService: %s", errors.New("Rate must be between 0 and 10")))
	}

	dtoReview, err := rs.storage.AddReview(dto.ConvertAddToDto(addReview))
	if err != nil {
		return nil, err
	}

	review := dtoReview.Convert()

	game, err := rs.verifyCache(dtoReview.ID)
	if err != nil {
		return nil, err
	}

	review.Game = game

	return review, nil
}

func (rs *ReviewService) UpdateReview(id int, updatReview *models.UpdateReview) (*models.GetReview, error) {

	if updatReview.Rate < 0 || updatReview.Rate > 10 {
		return nil, errors.Wrap(httperr.ErrBadRequest, fmt.Sprintf("ReviewService: %s", errors.New("Rate must be between 0 and 10")))
	}

	dtoReview, err := rs.storage.UpdateReview(id, dto.ConvertUpdateToDto(updatReview))
	if err != nil {
		return nil, err
	}

	review := dtoReview.Convert()
	game, err := rs.verifyCache(dtoReview.ID)
	if err != nil {
		return nil, err
	}

	review.Game = game

	return review, nil
}

func (rs *ReviewService) DeleteReview(id int) error {

	err := rs.storage.DeleteReview(id)
	if err != nil {
		return err
	}

	return nil
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
	game, err := rs.verifyCache(dtoReview.GameID)
	if err != nil {
		return nil, err
	}

	review.Game = game

	return review, nil
}

func (rs *ReviewService) verifyCache(id int) (models.Game, error) {
	game, err := rs.cache.GetGameById(id)
	if err == nil {
		return game, nil
	}

	game, err = rs.gs.GetGameById(id)
	if err != nil {
		return models.Game{}, err
	}

	err = rs.cache.SetGame(game)
	if err != nil {
		return models.Game{}, err
	}

	return game, nil
}
