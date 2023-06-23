package gorm

import (
	"fmt"
	"log"

	dto "github.com/Reg00/gameReview/internal/domain/dto"
	"github.com/Reg00/gameReview/internal/domain/dto/httperr"
	"github.com/Reg00/gameReview/internal/infrastructure/config"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormStorage struct {
	gormDB *gorm.DB
}

func Register(cfg *config.Configuration) (*GormStorage, error) {
	opts, err := newGormOptions(cfg.Storage.Options)
	if err != nil {
		return nil, fmt.Errorf("gorm: malformed options: %s", err)
	}
	log.Println(opts.dsn)
	var db = &gorm.DB{}
	switch {
	case opts.driver == "postgres":
		db, err = gorm.Open(postgres.Open(opts.dsn))
		if err != nil {
			return nil, fmt.Errorf("gorm: open db: %s", err)
		}
	default:
		return nil, errors.New("gorm: invalid storage driver")
	}
	db.AutoMigrate(&dto.Review{})
	return &GormStorage{
		gormDB: db,
	}, nil
}

func (storage *GormStorage) AddReview(review *dto.Review) (*dto.Review, error) {
	result := storage.gormDB.Create(review)
	if result.Error != nil {
		return nil, result.Error
	}
	return review, nil
}

func (storage *GormStorage) GetReviewById(id int) (*dto.Review, error) {
	var review dto.Review
	storage.gormDB.First(&review, "id = ?", id)

	if review == (dto.Review{}) {
		return &dto.Review{}, errors.Wrap(httperr.ErrNotFound, fmt.Sprintf("GORM: %s", errors.New(fmt.Sprintf("Cannot find review with %d", id))))
	}

	return &review, nil
}
