package port

import dto "github.com/Reg00/gameReview/internal/domain/dto"

type Storager interface {
	AddReview(review *dto.Review) (*dto.Review, error)
	UpdateReview(id int, updatReview *dto.Review) (*dto.Review, error)
	GetReviewById(id int) (*dto.Review, error)
	DeleteReview(id int) error
}
