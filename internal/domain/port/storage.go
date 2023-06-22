package port

import "github.com/Reg00/gameReview/internal/domain/dto/storage"

type Storager interface {
	AddReview(review *storage.Review) (*storage.Review, error)
}
